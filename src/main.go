/*
Copyright 2019 The Cloud-Barista Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package main is the starting point of CM-Beetle
package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/fsnotify/fsnotify"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"

	"github.com/cloud-barista/cm-beetle/src/core/common"

	restServer "github.com/cloud-barista/cm-beetle/src/api/rest/server"
)

// init for main
func init() {
	profile := "cloud_conf"
	setConfig(profile)
}

// setConfig get cloud settings from a config file
func setConfig(profile string) {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf/")
	viper.AddConfigPath("../conf/")
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	err = viper.Unmarshal(&common.RuntimeConf)
	if err != nil {
		panic(err)
	}

	// const mrttArrayXMax = 300
	// const mrttArrayYMax = 300
	// common.RuntimeLatancyMap = make([][]string, mrttArrayXMax)

	// cloudlatencymap.csv
	file, fileErr := os.Open("../assets/cloudlatencymap.csv")
	defer file.Close()
	if fileErr != nil {
		common.CBLog.Error(fileErr)
		panic(fileErr)
	}
	rdr := csv.NewReader(bufio.NewReader(file))
	common.RuntimeLatancyMap, _ = rdr.ReadAll()

	for i, v := range common.RuntimeLatancyMap {
		if i == 0 {
			continue
		}
		if v[0] == "" {
			break
		}
		common.RuntimeLatancyMapIndex[v[0]] = i
	}

	//fmt.Printf("RuntimeLatancyMap: %v\n\n", common.RuntimeLatancyMap)
	//fmt.Printf("[RuntimeLatancyMapIndex]\n %v\n", common.RuntimeLatancyMapIndex)

}

// Main Body

// @title CM-Beetle REST API
// @version latest
// @description CM-Beetle REST API

// @contact.name API Support
// @contact.url http://cloud-barista.github.io
// @contact.email contact-to-cloud-barista@googlegroups.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /beetle

// @securityDefinitions.basic BasicAuth
func main() {
	fmt.Println("")

	// giving a default value of "1323"
	port := flag.String("port", "1323", "port number for the restapiserver to listen to")
	flag.Parse()

	// validate arguments from flag
	validationFlag := true
	// validation: port
	// set validationFlag to false if your number is not in [1-65535] range
	if portInt, err := strconv.Atoi(*port); err == nil {
		if portInt < 1 || portInt > 65535 {
			validationFlag = false
		}
	} else {
		validationFlag = false
	}
	if !validationFlag {
		fmt.Printf("%s is not a valid port number.\n", *port)
		fmt.Printf("Please retry with a valid port number (ex: -port=[1-65535]).\n")
		os.Exit(1)
	}

	common.SpiderRestUrl = common.NVL(os.Getenv("SPIDER_REST_URL"), "http://localhost:1024/spider")
	common.DragonflyRestUrl = common.NVL(os.Getenv("DRAGONFLY_REST_URL"), "http://localhost:9090/dragonfly")
	common.DBUrl = common.NVL(os.Getenv("DB_URL"), "localhost:3306")
	common.DBDatabase = common.NVL(os.Getenv("DB_DATABASE"), "cb_beetle")
	common.DBUser = common.NVL(os.Getenv("DB_USER"), "cb_beetle")
	common.DBPassword = common.NVL(os.Getenv("DB_PASSWORD"), "cb_beetle")
	common.AutocontrolDurationMs = common.NVL(os.Getenv("AUTOCONTROL_DURATION_MS"), "10000")

	// load the latest configuration from DB (if exist)
	fmt.Println("")
	fmt.Println("[Update system environment]")
	common.UpdateGlobalVariable(common.StrDragonflyRestUrl)
	common.UpdateGlobalVariable(common.StrSpiderRestUrl)
	common.UpdateGlobalVariable(common.StrAutocontrolDurationMs)

	// load config
	//masterConfigInfos = confighandler.GetMasterConfigInfos()

	//Setup database (meta_db/dat/cmbeetle.s3db)
	fmt.Println("")
	fmt.Println("[Setup SQL Database]")

	err := os.MkdirAll("../meta_db/dat/", os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}

	go func() {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
			err := viper.ReadInConfig()
			if err != nil { // Handle errors reading the config file
				panic(fmt.Errorf("fatal error config file: %w", err))
			}
			err = viper.Unmarshal(&common.RuntimeConf)
			if err != nil {
				panic(err)
			}
		})
	}()

	// Launch API servers (REST)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	// Start REST Server
	go func() {
		restServer.RunServer(*port)
		wg.Done()
	}()

	// Note: Deprecated gRPC server
	// Start gRPC Server
	// go func() {
	// 	grpcServer.RunServer()
	// 	wg.Done()
	// }()
	// fmt.Println("RuntimeConf: ", common.RuntimeConf.Cloud)

	wg.Wait()
}
