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

type OPTimeQueryReq struct {
}
type OPTimeQueryResponse struct {
	Msg  string          `json:"msg"`
	Code int             `json:"code"`
	Data OPTimeQueryData `json:"data"`
}

type OPTimeQueryData struct {
	Name        string `json:"name"`
	Ret         int    `json:"ret"`
	OPTimeQuery string `json:"OPTimeQuery"`
}

func OpeDevOPTimeQuery(jDevice *v3.JLinkDevice, pdcd *OPTimeQueryReq) (string, error) {
	parm := make(map[string]interface{})
	parm["Name"] = utils.OPTimeQuery
	// fmt.Println(parm)
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return "", err
	}
	// fmt.Println(string(postData))
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DOpdevUrl, jDevice.Jdtoken)
	resbody, err := jDevice.HttpPost(url, strings.NewReader(string(postData)))
	if err != nil {
		log.Println("HttpPost err:" + err.Error())
		return "", err
	}
	// fmt.Println(string(resbody))
	resp := &OPTimeQueryResponse{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return "", err
	}

	if resp.Code == 2000 && resp.Data.Ret == 100 {
		return resp.Data.OPTimeQuery, nil
	}

	log.Println("OpeDevOPTimeQuery err：" + resp.Msg)
	return "", errors.New(resp.Msg)
}
