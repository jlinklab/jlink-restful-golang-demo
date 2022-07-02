package opedev

import (
	"encoding/json"
	"errors"
	"jlink-restful-golang-demo/http"
	"log"
)

type OPDefaultConfigReq struct {
	Account      bool `json:"Account"`
	XMModeSwitch bool `json:"XMModeSwitch"`
	Preview      bool `json:"Preview"`
	NetServer    bool `json:"NetServer"`
	Encode       bool `json:"Encode"`
	NetCommon    bool `json:"NetCommon"`
	CameraPARAM  bool `json:"CameraPARAM"`
	CommPtz      bool `json:"CommPtz"`
	Alarm        bool `json:"Alarm"`
	General      bool `json:"General"`
	Record       bool `json:"Record"`
}

func OpeDevOPDefaultConfig(pdcd *OPDefaultConfigReq, token string) (bool, error) {
	parm := make(map[string]interface{})
	parm["Name"] = OPDefaultConfig
	parm["OPDefaultConfig"] = pdcd
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

	log.Println("OpeDevOPDefaultConfig err：" + resp.Msg)
	return false, errors.New(resp.Msg)
}
