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

// Package common is to include common methods for managing multi-cloud infra
package common

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/cloud-barista/cm-beetle/pkg/lkvstore"
// 	"github.com/rs/zerolog/log"
// )

// // swagger:request ConfigReq
// type ConfigReq struct {
// 	Name  string `json:"name" example:"SPIDER_REST_URL"`
// 	Value string `json:"value" example:"http://localhost:1024/spider"`
// }

// // swagger:response ConfigInfo
// type ConfigInfo struct {
// 	Id    string `json:"id" example:"SPIDER_REST_URL"`
// 	Name  string `json:"name" example:"SPIDER_REST_URL"`
// 	Value string `json:"value" example:"http://localhost:1024/spider"`
// }

// func UpdateConfig(u *ConfigReq) (ConfigInfo, error) {

// 	if u.Name == "" {
// 		return ConfigInfo{}, fmt.Errorf("The provided name is empty.")
// 	}

// 	content := ConfigInfo{}
// 	content.Id = u.Name
// 	content.Name = u.Name
// 	content.Value = u.Value

// 	key := "/config/" + content.Id
// 	//mapA := map[string]string{"name": content.Name, "description": content.Description}
// 	val, _ := json.Marshal(content)
// 	lkvstore.Put(key, string(val))
// 	if err != nil {
// 		log.Error().Err(err).Msg("")
// 		return content, err
// 	}
// 	keyValue, exists := lkvstore.GetKv(key)
// 	if !exists {
// 		err := fmt.Errorf("Failed to put the config with key: " + key)
// 		log.Error().Err(err).Msg("")
// 		return content, err
// 	}

// 	log.Info().Msgf("UpdateConfig(); Key: " + keyValue.Key + "\nValue: " + keyValue.Value)

// 	UpdateGlobalVariable(content.Id)

// 	return content, nil
// }

// func UpdateGlobalVariable(id string) error {

// 	configInfo, err := GetConfig(id)
// 	if err != nil {
// 		//CBLog.Error(err)
// 		return err
// 	}

// 	switch id {
// 	// case StrSpiderRestUrl:
// 	// 	SpiderRestUrl = configInfo.Value
// 	// 	fmt.Println("<SPIDER_REST_URL> " + SpiderRestUrl)
// 	// case StrDragonflyRestUrl:
// 	// 	DragonflyRestUrl = configInfo.Value
// 	// 	fmt.Println("<DRAGONFLY_REST_URL> " + DragonflyRestUrl)
// 	case StrTumblebugRestUrl:
// 		TumblebugRestUrl = configInfo.Value
// 		fmt.Println("<TUMBELBUG_REST_URL> " + TumblebugRestUrl)
// 	case StrDBUrl:
// 		DBUrl = configInfo.Value
// 		fmt.Println("<DB_URL> " + DBUrl)
// 	case StrDBDatabase:
// 		DBDatabase = configInfo.Value
// 		fmt.Println("<DB_DATABASE> " + DBDatabase)
// 	case StrDBUser:
// 		DBUser = configInfo.Value
// 		fmt.Println("<DB_USER> " + DBUser)
// 	case StrDBPassword:
// 		DBPassword = configInfo.Value
// 		fmt.Println("<DB_PASSWORD> " + DBPassword)
// 	case StrAutocontrolDurationMs:
// 		AutocontrolDurationMs = configInfo.Value
// 		fmt.Println("<AUTOCONTROL_DURATION_MS> " + AutocontrolDurationMs)
// 	default:

// 	}

// 	return nil
// }

// func InitConfig(id string) error {

// 	switch id {
// 	// case StrSpiderRestUrl:
// 	// 	SpiderRestUrl = NVL(os.Getenv("BEETLE_SPIDER_REST_URL"), "http://localhost:1024/spider")
// 	// 	fmt.Println("<SPIDER_REST_URL> " + SpiderRestUrl)
// 	// case StrDragonflyRestUrl:
// 	// 	DragonflyRestUrl = NVL(os.Getenv("BEETLE_DRAGONFLY_REST_URL"), "http://localhost:9090/dragonfly")
// 	// 	fmt.Println("<DRAGONFLY_REST_URL> " + DragonflyRestUrl)
// 	case StrTumblebugRestUrl:
// 		TumblebugRestUrl = NVL(os.Getenv("BEETLE_TUMBLEBUG_REST_URL"), "http://localhost:1323/tumblebug")
// 		fmt.Println("<BEETLE_TUMBLEBUG_REST_URL> " + TumblebugRestUrl)
// 	case StrDBUrl:
// 		DBUrl = NVL(os.Getenv("BEETLE_SQLITE_URL"), "localhost:3306")
// 		fmt.Println("<BEETLE_SQLITE_URL> " + DBUrl)
// 	case StrDBDatabase:
// 		DBDatabase = NVL(os.Getenv("BEETLE_SQLITE_DATABASE"), "cm_beetle")
// 		fmt.Println("<BEETLE_SQLITE_DATABASE> " + DBDatabase)
// 	case StrDBUser:
// 		DBUser = NVL(os.Getenv("BEETLE_SQLITE_USER"), "cm_beetle")
// 		fmt.Println("<BEETLE_SQLITE_USER> " + DBUser)
// 	case StrDBPassword:
// 		DBPassword = NVL(os.Getenv("BEETLE_SQLITE_PASSWORD"), "cm_beetle")
// 		fmt.Println("<BEETLE_SQLITE_PASSWORD> " + DBPassword)
// 	case StrAutocontrolDurationMs:
// 		AutocontrolDurationMs = NVL(os.Getenv("BEETLE_AUTOCONTROL_DURATION_MS"), "10000")
// 		fmt.Println("<BEETLE_AUTOCONTROL_DURATION_MS> " + AutocontrolDurationMs)
// 	default:

// 	}

// 	check, err := CheckConfig(id)

// 	if check && err == nil {
// 		fmt.Println("[Init config] " + id)
// 		key := "/config/" + id

// 		lkvstore.Delete(key)
// 		// if err != nil {
// 		// 	CBLog.Error(err)
// 		// 	return err
// 		// }
// 	}

// 	return nil
// }

// func GetConfig(id string) (ConfigInfo, error) {

// 	res := ConfigInfo{}

// 	check, err := CheckConfig(id)

// 	if !check {
// 		errMsg := fmt.Errorf("dose not exist, the config ID: %s", id)
// 		log.Error().Err(errMsg).Msg("")
// 		return res, err
// 	}

// 	if err != nil {
// 		temp := ConfigInfo{}
// 		log.Error().Err(err)
// 		return temp, err
// 	}

// 	fmt.Println("[Get config] " + id)
// 	key := "/config/" + id

// 	keyValue, exist := lkvstore.GetKv(key)
// 	if !exist {
// 		errMsg := fmt.Errorf("dose not exist, the config ID: %s", id)
// 		log.Error().Err(errMsg).Msg("")
// 		return res, errMsg
// 	}

// 	log.Debug().Msgf("GetConfig(); Key: " + keyValue.Key + "\nValue: " + keyValue.Value)

// 	err = json.Unmarshal([]byte(keyValue.Value), &res)
// 	if err != nil {
// 		log.Error().Err(err).Msg("")
// 		return res, err
// 	}
// 	return res, nil
// }

// func ListConfig() ([]ConfigInfo, error) {
// 	fmt.Println("[List config]")
// 	key := "/config"
// 	fmt.Println(key)

// 	keyValue, err := CBStore.GetList(key, true)
// 	keyValue = cbstore_utils.GetChildList(keyValue, key)

// 	if err != nil {
// 		CBLog.Error(err)
// 		return nil, err
// 	}
// 	if keyValue != nil {
// 		res := []ConfigInfo{}
// 		for _, v := range keyValue {
// 			tempObj := ConfigInfo{}
// 			err = json.Unmarshal([]byte(v.Value), &tempObj)
// 			if err != nil {
// 				CBLog.Error(err)
// 				return nil, err
// 			}
// 			res = append(res, tempObj)
// 		}
// 		return res, nil
// 		//return true, nil
// 	}
// 	return nil, nil // When err == nil && keyValue == nil
// }

// func ListConfigId() []string {

// 	fmt.Println("[List config]")
// 	key := "/config"
// 	fmt.Println(key)

// 	keyValue, _ := CBStore.GetList(key, true)

// 	var configList []string
// 	for _, v := range keyValue {
// 		configList = append(configList, strings.TrimPrefix(v.Key, "/config/"))
// 	}
// 	for _, v := range configList {
// 		fmt.Println("<" + v + "> \n")
// 	}
// 	fmt.Println("===============================================")
// 	return configList
// }

/*
func DelAllConfig() error {
	fmt.Printf("DelAllConfig() called;")

	key := "/config"
	fmt.Println(key)
	keyValue, _ := CBStore.GetList(key, true)

	if len(keyValue) == 0 {
		return nil
	}

	for _, v := range keyValue {
		err = CBStore.Delete(v.Key)
		if err != nil {
			return err
		}
	}
	return nil
}
*/

// func InitAllConfig() error {
// 	fmt.Printf("InitAllConfig() called;")

// 	configIdList := ListConfigId()

// 	for _, v := range configIdList {
// 		InitConfig(v)
// 	}
// 	return nil
// }

// func CheckConfig(id string) (bool, error) {

// 	if id == "" {
// 		err := fmt.Errorf("invalid config ID: %s.", id)
// 		log.Error().Err(err).Msg("")
// 		return false, err
// 	}

// 	key := "/config/" + id

// 	_, exist := lkvstore.GetKv(key)

// 	return exist, nil
// }
