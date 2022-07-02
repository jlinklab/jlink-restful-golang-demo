package opedev

import (
	"encoding/json"
	"errors"
	"jlink-restful-golang-demo/http"
	"log"
)

type OPSystemConfigReq struct {
	OPSystemConfig string                    `json:"OPSystemConfig"`
	Action         string                    `json:"Action"`
	Parameter      []OPSystemConfigParameter `json:"Parameter"`
}
type OPSystemConfigParameter struct {
	Type string `json:"Type"`
}

func OpeDevOPSystemConfig(pdcd *OPSystemConfigReq, token string) (bool, error) {
	parm := make(map[string]interface{})
	parm["Name"] = OPSystemConfig
	parm["OPSystemConfig"] = pdcd
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

	log.Println("OpeDevOPSystemConfig err：" + resp.Msg)
	return false, errors.New(resp.Msg)
}
