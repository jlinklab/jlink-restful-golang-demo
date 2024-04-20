package opedev

import (
	"encoding/json"
	"errors"
	"fmt"
	"jlink-restful-golang-demo/utils"
	v3 "jlink-restful-golang-demo/v3"
	"log"
	"strings"
)

type StorageInfoReq struct {
}

type StorageInfoResponse struct {
	Code int             `json:"msg"`
	Msg  string          `json:"code"`
	Data StorageInfoData `json:"data"`
}

type StorageInfoData struct {
	Name        string            `json:"name"`
	Ret         int               `json:"ret"`
	StorageInfo []StorageInfoResp `json:"StorageInfo"`
}

type StorageInfoResp struct {
	PlysicalNo   int       `json:"PlysicalNo"`
	SerialNumber string    `json:"SerialNumber"`
	ModelNumber  string    `json:"ModelNumber"`
	PartNumber   int       `json:"PartNumber"`
	Partition    Partition `json:"Partition"`
}

type Partition struct {
	DirverType    int    `json:"DirverType"`
	TotalSpace    int    `json:"TotalSpace"`
	RemainSpace   int    `json:"RemainSpace"`
	IsCurrent     bool   `json:"IsCurrent"`
	LogicSerialNo int    `json:"LogicSerialNo"`
	NewEndTime    string `json:"NewEndTime"`
	NewStartTime  string `json:"NewStartTime"`
	OldEndTime    string `json:"OldEndTime"`
	OldStartTime  string `json:"OldStartTime"`
	Status        int    `json:"Status"`
}

func OpeDevStorageInfo(jDevice *v3.JLinkDevice, pdcd *StorageInfoReq) (bool, error) {
	parm := make(map[string]interface{})
	parm["Name"] = utils.StorageInfo
	// parm["OPPTZControl"] = pdcd
	// fmt.Println(parm)
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return false, err
	}
	// fmt.Println(string(postData))
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DOpdevUrl, jDevice.Jdtoken)
	resbody, err := jDevice.HttpPost(url, strings.NewReader(string(postData)))
	if err != nil {
		log.Println("HttpPost err:" + err.Error())
		return false, err
	}
	// fmt.Println(string(resbody))
	resp := &utils.XYResponse{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return false, err
	}

	if resp.Code == 2000 && resp.Data.Ret == 100 {
		return true, nil
	}

	log.Println("DeviceOperate err：" + resp.Msg)
	return false, errors.New(resp.Msg)
}
