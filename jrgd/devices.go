package jrgd

import (
	"encoding/json"
	"errors"
	"fmt"
	"jlink-restful-golang-demo/http"
	"log"
)

type CaptureResponse struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data CaptureData `json:"data"`
}
type CaptureData struct {
	Ret   int    `json:"Ret"`
	Image string `json:"image"`
}

type OPSNAPReq struct {
	Channel int `json:"Channel"`
	PicType int `json:"picType"`
}

// Capture
func (jDevice *JLinkDevice) Capture(channel int) (string, error) {
	parm := make(map[string]interface{})
	parm["Name"] = OPSNAP
	parm["OPSNAP"] = OPSNAPReq{Channel: channel, PicType: 0}
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return "", err
	}
	resbody, err := http.HttpPost(dcaptureUrl, jDevice.Jdtoken, postData)
	if err != nil {
		log.Println("HttpPost err:" + err.Error())
		return "", err
	}
	fmt.Println(string(resbody))
	resp := &CaptureResponse{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		log.Println("Unmarshal err:" + err.Error())
		return "", err
	}

	if resp.Code == 2000 && resp.Data.Ret == 100 {
		return resp.Data.Image, nil
	}
	log.Println("Capture err：" + resp.Msg)
	return "", errors.New(resp.Msg)

}
