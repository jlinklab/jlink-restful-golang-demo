package opedev

import (
	"encoding/json"
	"errors"
	"jlink-restful-golang-demo/http"
	"log"
)

type OPFileQueryReq struct {
	Channel        int    `json:"Channel"`
	Type           string `json:"Type"`
	Event          string `json:"Event"`
	DriverTypeMask string `json:"DriverTypeMask"`
	BeginTime      string `json:"BeginTime"`
	EndTime        string `json:"EndTime"`
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

func OpeDevOPFileQuery(pdcd *OPFileQueryReq, token string) ([]OPFileQueryResp, error) {
	parm := make(map[string]interface{})
	parm["Name"] = OPFileQuery
	parm["OPFileQuery"] = pdcd
	// fmt.Println(parm)
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return nil, err
	}
	// fmt.Println(string(postData))
	resbody, err := http.HttpPost(dOpdevUrl, token, postData)
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
