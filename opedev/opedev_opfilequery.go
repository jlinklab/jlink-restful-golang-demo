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

type OPFileQueryReq struct {
	Channel        int    `json:"Channel"`
	Type           string `json:"Type"`
	Event          string `json:"Event"`
	DriverTypeMask string `json:"DriverTypeMask"`
	BeginTime      string `json:"BeginTime"`
	EndTime        string `json:"EndTime"`
	StreamType     string `json:"StreamType"`
}

type OPFileQueryResponse struct {
	Code int             `json:"code"`
	Data OPFileQueryData `json:"data"`
	Msg  string          `json:"msg"`
}

type OPFileQueryResp struct {
	BeginTime  string `json:"BeginTime"`
	DiskNo     int    `json:"DiskNo"`
	EndTime    string `json:"EndTime"`
	FileLength string `json:"FileLength"`
	FileName   string `json:"FileName"`
	SerialNo   int    `json:"SerialNo"`
}

type OPFileQueryData struct {
	Name        string            `json:"Name"`
	OPFileQuery []OPFileQueryResp `json:"OPFileQuery"`
	Ret         int               `json:"Ret"`
	SessionID   string            `json:"SessionID"`
}

func OpeDevOPFileQuery(jDevice *v3.JLinkDevice, pdcd *OPFileQueryReq) ([]OPFileQueryResp, error) {
	parm := make(map[string]interface{})
	parm["Name"] = utils.OPFileQuery
	parm["OPFileQuery"] = pdcd
	// fmt.Println(parm)
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return nil, err
	}
	// fmt.Println(string(postData))
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DOpdevUrl, jDevice.Jdtoken)
	resbody, err := jDevice.HttpPost(url, strings.NewReader(string(postData)))
	if err != nil {
		log.Println("HttpPost err:" + err.Error())
		return nil, err
	}
	// fmt.Println(string(resbody))
	resp := &OPFileQueryResponse{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return nil, err
	}

	if resp.Code == 2000 && resp.Data.Ret == 100 {
		return resp.Data.OPFileQuery, nil
	}

	log.Println("OpeDevOPDefaultConfig err：" + resp.Msg)
	return nil, errors.New(resp.Msg)
}
