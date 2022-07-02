package opedev

import (
	"encoding/json"
	"errors"
	"jlink-restful-golang-demo/http"
	"log"
)

type OPTimeSettingReq struct {
	OPTimeSetting string `json:"OPTimeSetting"`
}

func OpeDevOPTimeSetting(pdcd *OPTimeSettingReq, token string) (bool, error) {
	parm := make(map[string]interface{})
	parm["Name"] = OPTimeSetting
	parm["OPTimeSetting"] = pdcd.OPTimeSetting
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

	log.Println("OpeDevOPTimeSetting err：" + resp.Msg)
	return false, errors.New(resp.Msg)
}
