package opedev

import (
	"encoding/json"
	"errors"
	"jlink-restful-golang-demo/http"
	"log"
)

type OPStorageManagerReq struct {
	Action        string          `json:"Action"`
	SerialNo      int             `json:"SerialNo"`
	PartitionSize []PartitionSize `json:"PartitionSize"`
}
type PartitionSize struct {
	Record   int `json:"Record,omitempty"`
	SnapShot int `json:"SnapShot,omitempty"`
}

func OpeDevPartitionSize(pdcd *OPStorageManagerReq, token string) (bool, error) {
	parm := make(map[string]interface{})
	parm["Name"] = OPStorageManager
	parm["OPStorageManager"] = pdcd
	// fmt.Println(parm)
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return false, err
	}
	// fmt.Println(string(postData))
	resbody, err := http.HttpPost(dOpdevUrl, token, postData)
	if err != nil {
		log.Println("HttpPost err:" + err.Error())
		return false, err
	}
	// fmt.Println(string(resbody))
	resp := &XYResponse{}
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
