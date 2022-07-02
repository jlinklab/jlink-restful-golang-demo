package opedev

import (
	"encoding/json"
	"errors"
	"jlink-restful-golang-demo/http"
	"log"
)

type OPLogQueryReq struct {
	BeginTime   string `json:"BeginTime"`
	EndTime     string `json:"EndTime"`
	LogPosition string `json:"LogPosition"`
	Type        string `json:"Type"`
}

type OPLogQueryResp struct {
	Name       string           `json:"Name"`
	Ret        string           `json:"Ret"`
	OPLogQuery []OPLogQueryData `json:"OPLogQuery"`
}
type OPLogQueryData struct {
	Data     string `json:"Data"`
	Position int    `json:"Position"`
	Time     string `json:"Time"`
	Type     string `json:"Type"`
	User     string `json:"User"`
}

func OpeDevOPLogQuery(pdcd *OPLogQueryReq, token string) (bool, error) {
	parm := make(map[string]interface{})
	parm["Name"] = OPLogQuery
	parm["OPLogQuery"] = pdcd
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
