package jrgd

import (
	"encoding/json"
	"errors"
	"jlink-restful-golang-demo/http"
	"log"
)

type Response struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// 	Low-power device wakeup--
func (jDevice *JLinkDevice) WakeUp() (bool, error) {
	payload := http.NewStatusRequest(jDevice.Jdtoken)
	resbody, err := http.HttpPost(dWakeupUrl, jDevice.Jdtoken, payload)
	if err != nil {
		log.Println("HttpPosterr:" + err.Error())
		return false, err
	}
	resp := &Response{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		log.Println(err)
		return false, err
	}

	if resp.Code == 2000 {
		return true, nil
	}
	log.Println("WakeUperr：" + resp.Msg)
	return false, errors.New(resp.Msg)
}
