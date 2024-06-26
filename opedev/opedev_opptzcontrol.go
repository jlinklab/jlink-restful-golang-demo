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

type PtzDirectionControlDTO struct {
	Command   string    `json:"Command"`
	Parameter Parameter `json:"Parameter"`
}
type Parameter struct {
	Preset  int `json:"Preset"`
	Channel int `json:"Channel"`
	Step    int `json:"Step"`
}

func OpeDevOPPTZControl(jDevice *v3.JLinkDevice, pdcd *PtzDirectionControlDTO) (bool, error) {
	parm := make(map[string]interface{})
	parm["Name"] = utils.OPPTZControl
	parm["OPPTZControl"] = pdcd
	// fmt.Println(parm)
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return false, err
	}

	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DOpdevUrl, jDevice.Jdtoken)
	resbody, err := jDevice.HttpPost(url, strings.NewReader(string(postData)))
	if err != nil {
		log.Println("HttpPost err:" + err.Error())
		return false, err
	}

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
