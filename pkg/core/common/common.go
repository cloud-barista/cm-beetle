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

import (
	"database/sql"
	"fmt"

	icbs "github.com/cloud-barista/cb-store/interfaces"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"xorm.io/xorm"
)

type KeyValue struct {
	Key   string
	Value string
}

type IdList struct {
	IdList []string `json:"output"`
}

// SystemReady is global variable for checking SystemReady status
var SystemReady bool

// CB-Store
var CBLog *logrus.Logger
var CBStore icbs.Store

// var SpiderRestUrl string
// var DragonflyRestUrl string
var TumblebugRestUrl string
var DBUrl string
var DBDatabase string
var DBUser string
var DBPassword string
var AutocontrolDurationMs string
var MyDB *sql.DB
var err error
var ORM *xorm.Engine

const (
	// StrSpiderRestUrl              string = "SPIDER_REST_URL"
	// StrDragonflyRestUrl           string = "DRAGONFLY_REST_URL"
	StrTumblebugRestUrl           string = "TUMBLEBUG_REST_URL"
	StrDBUrl                      string = "DB_URL"
	StrDBDatabase                 string = "DB_DATABASE"
	StrDBUser                     string = "DB_USER"
	StrDBPassword                 string = "DB_PASSWORD"
	StrAutocontrolDurationMs      string = "AUTOCONTROL_DURATION_MS"
	CbStoreKeyNotFoundErrorString string = "key not found"
	StrAdd                        string = "add"
	StrDelete                     string = "delete"
	StrSSHKey                     string = "sshKey"
	StrImage                      string = "image"
	StrCustomImage                string = "customImage"
	StrSecurityGroup              string = "securityGroup"
	StrSpec                       string = "spec"
	StrVNet                       string = "vNet"
	StrSubnet                     string = "subnet"
	StrDataDisk                   string = "dataDisk"
	StrNLB                        string = "nlb"
	StrVM                         string = "vm"
	StrMCI                       string = "mci"
	StrDefaultResourceName        string = "-systemdefault-"
	// StrFirewallRule               string = "firewallRule"

	// SystemCommonNs is const for SystemCommon NameSpace ID
	SystemCommonNs string = "system-purpose-common-ns"
)

var StartTime string

func init() {
	// CBLog = config.Cblogger
	// CBStore = cbstore.GetStore()

	// StartTime = time.Now().Format("2006.01.02 15:04:05 Mon")
}

// Spider 2020-03-30 https://github.com/cloud-barista/cb-spider/blob/master/cloud-control-manager/cloud-driver/interfaces/resources/IId.go
type IID struct {
	NameId   string // NameID by user
	SystemId string // SystemID by CloudOS
}

type SpiderConnectionName struct {
	ConnectionName string `json:"ConnectionName"`
}

func OpenSQL(path string) error {
	/*
		common.MYDB, err = sql.Open("mysql", //"root:pwd@tcp(127.0.0.1:3306)/testdb")
			common.DB_USER+":"+
				common.DB_PASSWORD+"@tcp("+
				common.DB_URL+")/"+
				common.DB_DATABASE)
	*/

	fullPathString := "file:" + path
	MyDB, err = sql.Open("sqlite3", fullPathString)
	return err
}

func SelectDatabase(database string) error {
	query := "USE " + database + ";"
	_, err = MyDB.Exec(query)
	return err
}

/*
func CreateSpecTable() error {
	stmt, err := MYDB.Prepare("CREATE Table IF NOT EXISTS spec(" +
		"namespace varchar(50) NOT NULL," +
		"id varchar(50) NOT NULL," +
		"connectionName varchar(50) NOT NULL," +
		"cspSpecName varchar(50) NOT NULL," +
		"name varchar(50)," +
		"osType varchar(50)," +
		"numvCPU SMALLINT," + // SMALLINT: -32768 ~ 32767
		"numcore SMALLINT," + // SMALLINT: -32768 ~ 32767
		"memGiB SMALLINT," + // SMALLINT: -32768 ~ 32767
		"storageGiB MEDIUMINT," + // MEDIUMINT: -8388608 to 8388607
		"description varchar(50)," +
		"costPerHour FLOAT," +
		"numAtorage SMALLINT," + // SMALLINT: -32768 ~ 32767
		"maxNumStorage SMALLINT," + // SMALLINT: -32768 ~ 32767
		"maxTotalStorage_TiB SMALLINT," + // SMALLINT: -32768 ~ 32767
		"netBwGbps SMALLINT," + // SMALLINT: -32768 ~ 32767
		"ebsBwMbps MEDIUMINT," + // MEDIUMINT: -8388608 to 8388607
		"gpuModel varchar(50)," +
		"numGpu SMALLINT," + // SMALLINT: -32768 ~ 32767
		"gpumemGiB SMALLINT," + // SMALLINT: -32768 ~ 32767
		"gpuP2p varchar(50)," +
		"orderInFilteredResult SMALLINT," + // SMALLINT: -32768 ~ 32767
		"evaluationStatus varchar(50)," +
		"evaluationScore01 FLOAT," +
		"evaluationScore02 FLOAT," +
		"evaluationScore03 FLOAT," +
		"evaluationScore04 FLOAT," +
		"evaluationScore05 FLOAT," +
		"evaluationScore06 FLOAT," +
		"evaluationScore07 FLOAT," +
		"evaluationScore08 FLOAT," +
		"evaluationScore09 FLOAT," +
		"evaluationScore10 FLOAT," +
		"CONSTRAINT PK_Spec PRIMARY KEY (namespace, id));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()

	return err
}

func CreateImageTable() error {
	stmt, err := MYDB.Prepare("CREATE Table IF NOT EXISTS image(" +
		"namespace varchar(50) NOT NULL," +
		"id varchar(50) NOT NULL," +
		"name varchar(50)," +
		"connectionName varchar(50) NOT NULL," +
		"cspImageId varchar(400) NOT NULL," +
		"cspImageName varchar(400) NOT NULL," +
		"creationDate varchar(50) NOT NULL," +
		"description varchar(400) NOT NULL," +
		"guestOS varchar(50) NOT NULL," +
		"status varchar(50) NOT NULL," +
		"CONSTRAINT PK_Image PRIMARY KEY (namespace, id));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()

	return err
}
*/

var DefaulNamespaceId = "ns-mig01"

// NsReq is struct for namespace creation
type NsReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// NsInfo is struct for namespace information
type NsInfo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Response structure for RestGetAllNs
type RestGetAllNsResponse struct {
	//Name string     `json:"name"`
	Ns []NsInfo `json:"ns"`
}

func CreateNamespace(nsInfo NsReq) (NsInfo, error) {
	// Initialize resty client with basic auth
	client := resty.New()
	apiUser := viper.GetString("beetle.tumblebug.api.username")
	apiPass := viper.GetString("beetle.tumblebug.api.password")
	client.SetBasicAuth(apiUser, apiPass)

	// set endpoint
	epTumblebug := TumblebugRestUrl

	// check readyz
	method := "GET"
	url := fmt.Sprintf("%s/readyz", epTumblebug)
	reqReadyz := NoBody
	resReadyz := new(SimpleMsg)

	err := ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		SetUseBody(reqReadyz),
		&reqReadyz,
		resReadyz,
		VeryShortDuration,
	)

	if err != nil {
		log.Err(err).Msg("")
		return NsInfo{}, err
	}
	log.Debug().Msgf("resReadyz: %+v", resReadyz.Message)

	// check readyz
	method = "POST"
	url = fmt.Sprintf("%s/ns", epTumblebug)
	reqNs := nsInfo
	resNs := new(NsInfo)

	err = ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		SetUseBody(reqNs),
		&reqNs,
		resNs,
		VeryShortDuration,
	)

	if err != nil {
		log.Err(err).Msg("")
		return NsInfo{}, err
	}
	log.Debug().Msgf("resNs: %+v", resNs)

	return *resNs, nil
}

func GetAllNamespaces() (RestGetAllNsResponse, error) {
	// Initialize resty client with basic auth
	client := resty.New()
	apiUser := viper.GetString("beetle.tumblebug.api.username")
	apiPass := viper.GetString("beetle.tumblebug.api.password")
	client.SetBasicAuth(apiUser, apiPass)

	// set endpoint
	epTumblebug := TumblebugRestUrl

	// check readyz
	method := "GET"
	url := fmt.Sprintf("%s/readyz", epTumblebug)
	reqReadyz := NoBody
	resReadyz := new(SimpleMsg)

	err := ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		SetUseBody(reqReadyz),
		&reqReadyz,
		resReadyz,
		VeryShortDuration,
	)

	if err != nil {
		log.Err(err).Msg("")
		return RestGetAllNsResponse{}, err
	}
	log.Debug().Msgf("resReadyz: %+v", resReadyz.Message)

	// check readyz
	method = "GET"
	url = fmt.Sprintf("%s/ns", epTumblebug)
	reqNs := NoBody
	resAllNs := new(RestGetAllNsResponse)

	err = ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		SetUseBody(reqNs),
		&reqNs,
		resAllNs,
		VeryShortDuration,
	)

	if err != nil {
		log.Err(err).Msg("")
		return RestGetAllNsResponse{}, err
	}
	log.Debug().Msgf("resAllNs: %+v", resAllNs)

	return *resAllNs, nil
}

func GetNamespace(nsId string) (NsInfo, error) {
	// Initialize resty client with basic auth
	client := resty.New()
	apiUser := viper.GetString("beetle.tumblebug.api.username")
	apiPass := viper.GetString("beetle.tumblebug.api.password")
	client.SetBasicAuth(apiUser, apiPass)

	// set endpoint
	epTumblebug := TumblebugRestUrl

	// check readyz
	method := "GET"
	url := fmt.Sprintf("%s/readyz", epTumblebug)
	reqReadyz := NoBody
	resReadyz := new(SimpleMsg)

	err := ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		SetUseBody(reqReadyz),
		&reqReadyz,
		resReadyz,
		VeryShortDuration,
	)

	if err != nil {
		log.Err(err).Msg("")
		return NsInfo{}, err
	}
	log.Debug().Msgf("resReadyz: %+v", resReadyz.Message)

	// check readyz
	method = "GET"
	url = fmt.Sprintf("%s/ns/%s", epTumblebug, nsId)
	reqNs := NoBody
	resNsInfo := new(NsInfo)

	err = ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		SetUseBody(reqNs),
		&reqNs,
		resNsInfo,
		VeryShortDuration,
	)

	if err != nil {
		log.Err(err).Msg("")
		return NsInfo{}, err
	}
	log.Debug().Msgf("resNsInfo: %+v", resNsInfo)

	return *resNsInfo, nil
}

func DeleteNamespace(nsId string) (SimpleMsg, error) {
	// Initialize resty client with basic auth
	client := resty.New()
	apiUser := viper.GetString("beetle.tumblebug.api.username")
	apiPass := viper.GetString("beetle.tumblebug.api.password")
	client.SetBasicAuth(apiUser, apiPass)

	// set endpoint
	epTumblebug := TumblebugRestUrl

	// check readyz
	method := "GET"
	url := fmt.Sprintf("%s/readyz", epTumblebug)
	reqReadyz := NoBody
	resReadyz := new(SimpleMsg)

	err := ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		SetUseBody(reqReadyz),
		&reqReadyz,
		resReadyz,
		VeryShortDuration,
	)

	if err != nil {
		log.Err(err).Msg("")
		return SimpleMsg{}, err
	}
	log.Debug().Msgf("resReadyz: %+v", resReadyz.Message)

	// check readyz
	method = "DELETE"
	url = fmt.Sprintf("%s/ns/%s", epTumblebug, nsId)
	reqNs := NoBody
	resMsg := new(SimpleMsg)

	err = ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		SetUseBody(reqNs),
		&reqNs,
		resMsg,
		VeryShortDuration,
	)

	if err != nil {
		log.Err(err).Msg("")
		return SimpleMsg{}, err
	}
	log.Debug().Msgf("resMsg: %+v", resMsg)

	return *resMsg, nil
}

// func NsValidation() echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			fmt.Printf("%v\n", "[Handle API Request]")
// 			nsId := c.Param("nsId")
// 			if nsId == "" {
// 				return next(c)
// 			}

// 			err := CheckString(nsId)
// 			if err != nil {
// 				return echo.NewHTTPError(http.StatusNotFound, "The first character of name must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.")
// 			}

// 			check, err := CheckNs(nsId)

// 			if !check || err != nil {
// 				return echo.NewHTTPError(http.StatusNotFound, "Not valid namespace")
// 			}
// 			return next(c)
// 		}
// 	}
// }

// func CreateNs(u *NsReq) (NsInfo, error) {
// 	err := CheckString(u.Name)
// 	if err != nil {
// 		temp := NsInfo{}
// 		CBLog.Error(err)
// 		return temp, err
// 	}

// 	check, err := CheckNs(u.Name)

// 	if check {
// 		temp := NsInfo{}
// 		err := fmt.Errorf("CreateNs(); The namespace " + u.Name + " already exists.")
// 		return temp, err
// 	}

// 	if err != nil {
// 		temp := NsInfo{}
// 		CBLog.Error(err)
// 		return temp, err
// 	}

// 	content := NsInfo{}
// 	//content.Id = GenUid()
// 	content.Id = u.Name
// 	content.Name = u.Name
// 	content.Description = u.Description

// 	// TODO here: implement the logic

// 	fmt.Println("CreateNs();")
// 	Key := "/ns/" + content.Id
// 	//mapA := map[string]string{"name": content.Name, "description": content.Description}
// 	Val, _ := json.Marshal(content)
// 	err = CBStore.Put(Key, string(Val))
// 	if err != nil {
// 		CBLog.Error(err)
// 		return content, err
// 	}
// 	keyValue, _ := CBStore.Get(Key)
// 	fmt.Println("CreateNs(); ===========================")
// 	fmt.Println("CreateNs(); Key: " + keyValue.Key + "\nValue: " + keyValue.Value)
// 	fmt.Println("CreateNs(); ===========================")
// 	return content, nil
// }

// // UpdateNs is func to update namespace info
// func UpdateNs(id string, u *NsReq) (NsInfo, error) {

// 	res := NsInfo{}
// 	emptyInfo := NsInfo{}

// 	err := CheckString(id)
// 	if err != nil {
// 		CBLog.Error(err)
// 		return emptyInfo, err
// 	}
// 	check, err := CheckNs(id)

// 	if !check {
// 		errString := "The namespace " + id + " does not exist."
// 		err := fmt.Errorf(errString)
// 		return emptyInfo, err
// 	}

// 	if err != nil {
// 		CBLog.Error(err)
// 		return emptyInfo, err
// 	}

// 	key := "/ns/" + id
// 	keyValue, err := CBStore.Get(key)
// 	if err != nil {
// 		CBLog.Error(err)
// 		return emptyInfo, err
// 	}

// 	err = json.Unmarshal([]byte(keyValue.Value), &res)
// 	if err != nil {
// 		CBLog.Error(err)
// 		return emptyInfo, err
// 	}

// 	res.Id = id
// 	res.Name = u.Name
// 	res.Description = u.Description

// 	Key := "/ns/" + id
// 	//mapA := map[string]string{"name": content.Name, "description": content.Description}
// 	Val, err := json.Marshal(res)
// 	if err != nil {
// 		CBLog.Error(err)
// 		return emptyInfo, err
// 	}
// 	err = CBStore.Put(Key, string(Val))
// 	if err != nil {
// 		CBLog.Error(err)
// 		return emptyInfo, err
// 	}
// 	keyValue, err = CBStore.Get(Key)
// 	if err != nil {
// 		CBLog.Error(err)
// 		return emptyInfo, err
// 	}
// 	err = json.Unmarshal([]byte(keyValue.Value), &res)
// 	if err != nil {
// 		CBLog.Error(err)
// 		return emptyInfo, err
// 	}
// 	return res, nil
// }

// func GetNs(id string) (NsInfo, error) {

// 	res := NsInfo{}

// 	err := CheckString(id)
// 	if err != nil {
// 		temp := NsInfo{}
// 		CBLog.Error(err)
// 		return temp, err
// 	}
// 	check, err := CheckNs(id)

// 	if !check {
// 		errString := "The namespace " + id + " does not exist."
// 		//mapA := map[string]string{"message": errString}
// 		//mapB, _ := json.Marshal(mapA)
// 		err := fmt.Errorf(errString)
// 		return res, err
// 	}

// 	if err != nil {
// 		temp := NsInfo{}
// 		CBLog.Error(err)
// 		return temp, err
// 	}

// 	fmt.Println("[Get namespace] " + id)
// 	key := "/ns/" + id
// 	fmt.Println(key)

// 	keyValue, err := CBStore.Get(key)
// 	if err != nil {
// 		CBLog.Error(err)
// 		return res, err
// 	}

// 	fmt.Println("<" + keyValue.Key + "> \n" + keyValue.Value)
// 	fmt.Println("===============================================")

// 	err = json.Unmarshal([]byte(keyValue.Value), &res)
// 	if err != nil {
// 		CBLog.Error(err)
// 		return res, err
// 	}
// 	return res, nil
// }

// func ListNs() ([]NsInfo, error) {
// 	fmt.Println("[List namespace]")
// 	key := "/ns"
// 	fmt.Println(key)

// 	keyValue, err := CBStore.GetList(key, true)
// 	keyValue = cbstore_utils.GetChildList(keyValue, key)

// 	if err != nil {
// 		CBLog.Error(err)
// 		return nil, err
// 	}
// 	if keyValue != nil {
// 		res := []NsInfo{}
// 		for _, v := range keyValue {
// 			tempObj := NsInfo{}
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

// func AppendIfMissing(slice []string, i string) []string {
// 	for _, ele := range slice {
// 		if ele == i {
// 			return slice
// 		}
// 	}
// 	return append(slice, i)
// }

// func ListNsId() ([]string, error) {

// 	key := "/ns"

// 	var nsList []string

// 	// Implementation Option 1
// 	// keyValue, _ := CBStore.GetList(key, true)

// 	// r, _ := regexp.Compile("/ns/[a-z]([-a-z0-9]*[a-z0-9])?$")

// 	// for _, v := range keyValue {

// 	// 	if v.Key == "" {
// 	// 		continue
// 	// 	}

// 	// 	filtered := r.FindString(v.Key)

// 	// 	if filtered != v.Key {
// 	// 		continue
// 	// 	} else {
// 	// 		trimmedString := strings.TrimPrefix(v.Key, "/ns/")
// 	// 		nsList = AppendIfMissing(nsList, trimmedString)
// 	// 	}
// 	// }
// 	// EOF of Implementation Option 1

// 	// Implementation Option 2
// 	keyValue, err := CBStore.GetList(key, true)
// 	keyValue = cbstore_utils.GetChildList(keyValue, key)

// 	if err != nil {
// 		CBLog.Error(err)
// 		return nil, err
// 	}
// 	if keyValue != nil {
// 		for _, v := range keyValue {
// 			trimmedString := strings.TrimPrefix(v.Key, "/ns/")
// 			nsList = append(nsList, trimmedString)
// 		}
// 	}
// 	// EOF of Implementation Option 2

// 	return nsList, nil

// }

// func DelNs(id string) error {

// 	err := CheckString(id)
// 	if err != nil {
// 		CBLog.Error(err)
// 		return err
// 	}

// 	check, err := CheckNs(id)

// 	if !check {
// 		errString := "The namespace " + id + " does not exist."
// 		err := fmt.Errorf(errString)
// 		return err
// 	}

// 	if err != nil {
// 		CBLog.Error(err)
// 		return err
// 	}

// 	fmt.Println("[Delete ns] " + id)
// 	key := "/ns/" + id
// 	fmt.Println(key)

// 	mcisList := GetChildIdList(key + "/mcis")
// 	imageList := GetChildIdList(key + "/resources/image")
// 	vNetList := GetChildIdList(key + "/resources/vNet")
// 	//subnetList := GetChildIdList(key + "/resources/subnet")
// 	//publicIpList := GetChildIdList(key + "/resources/publicIp")
// 	securityGroupList := GetChildIdList(key + "/resources/securityGroup")
// 	specList := GetChildIdList(key + "/resources/spec")
// 	sshKeyList := GetChildIdList(key + "/resources/sshKey")
// 	//vNicList := GetChildIdList(key + "/resources/vNic")

// 	if len(mcisList)+
// 		len(imageList)+
// 		len(vNetList)+
// 		//len(subnetList)
// 		len(securityGroupList)+
// 		len(specList)+
// 		len(sshKeyList) > 0 {
// 		errString := "Cannot delete NS " + id + ", which is not empty. There exists at least one MCIS or one of resources."
// 		errString += " \n len(mcisList): " + strconv.Itoa(len(mcisList))
// 		errString += " \n len(imageList): " + strconv.Itoa(len(imageList))
// 		errString += " \n len(vNetList): " + strconv.Itoa(len(vNetList))
// 		//errString += " \n len(publicIpList): " + strconv.Itoa(len(publicIpList))
// 		errString += " \n len(securityGroupList): " + strconv.Itoa(len(securityGroupList))
// 		errString += " \n len(specList): " + strconv.Itoa(len(specList))
// 		errString += " \n len(sshKeyList): " + strconv.Itoa(len(sshKeyList))
// 		//errString += " \n len(subnetList): " + strconv.Itoa(len(subnetList))
// 		//errString += " \n len(vNicList): " + strconv.Itoa(len(vNicList))

// 		err := fmt.Errorf(errString)
// 		CBLog.Error(err)
// 		return err
// 	}

// 	// delete ns info
// 	err = CBStore.Delete(key)
// 	if err != nil {
// 		CBLog.Error(err)
// 		return err
// 	}

// 	return nil
// }

// func DelAllNs() error {
// 	fmt.Printf("DelAllNs() called;")

// 	nsIdList, err := ListNsId()
// 	if err != nil {
// 		return err
// 	}

// 	if len(nsIdList) == 0 {
// 		return nil
// 	}

// 	for _, v := range nsIdList {
// 		err := DelNs(v)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func CheckNs(id string) (bool, error) {

// 	if id == "" {
// 		err := fmt.Errorf("CheckNs failed; nsId given is null.")
// 		return false, err
// 	}

// 	err := CheckString(id)
// 	if err != nil {
// 		CBLog.Error(err)
// 		return false, err
// 	}

// 	key := "/ns/" + id

// 	keyValue, _ := CBStore.Get(key)
// 	if keyValue != nil {
// 		return true, nil
// 	}
// 	return false, nil
// }
