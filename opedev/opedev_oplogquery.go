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

func OpeDevOPLogQuery(jDevice *v3.JLinkDevice, pdcd *OPLogQueryReq) (bool, error) {
	parm := make(map[string]interface{})
	parm["Name"] = utils.OPLogQuery
	parm["OPLogQuery"] = pdcd
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

	log.Println("OpeDevOPTimeSetting err：" + resp.Msg)
	return false, errors.New(resp.Msg)
}
