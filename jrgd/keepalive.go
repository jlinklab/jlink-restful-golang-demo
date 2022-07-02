package jrgd

import (
	"encoding/json"
	"errors"
	"fmt"
	"jlink-restful-golang-demo/http"
	"log"
)

// Keepalive
func (jDevice *JLinkDevice) Keepalive() (bool, error) {
	parm := make(map[string]interface{})
	parm["Name"] = KeepAlive
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return false, err
	}
	resbody, err := http.HttpPost(dKeepaliveUrl, jDevice.Jdtoken, postData)
	if err != nil {
		log.Println("HttpPost err:" + err.Error())
		return false, err
	}
	fmt.Println(string(resbody))
	resp := &XYResponse{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		log.Println(err)
		return false, err
	}

	if resp.Code == 2000 && resp.Data.Ret == 100 {
		return true, nil
	}
	log.Println("Keepalive err：" + resp.Msg)
	return false, errors.New(resp.Msg)

}
