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

type OPSystemConfigReq struct {
	OPSystemConfig string                    `json:"OPSystemConfig"`
	Action         string                    `json:"Action"`
	Parameter      []OPSystemConfigParameter `json:"Parameter"`
}
type OPSystemConfigParameter struct {
	Type string `json:"Type"`
}

func OpeDevOPSystemConfig(jDevice *v3.JLinkDevice, pdcd *OPSystemConfigReq) (bool, error) {
	parm := make(map[string]interface{})
	parm["Name"] = utils.OPSystemConfig
	parm["OPSystemConfig"] = pdcd
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

	log.Println("OpeDevOPSystemConfig err：" + resp.Msg)
	return false, errors.New(resp.Msg)
}
