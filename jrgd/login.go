package jrgd

import (
	"encoding/json"
	"errors"
	"jlink-restful-golang-demo/http"
	"log"
)

// Notice DeviceType space after
type DeviceLoginData struct {
	AliveInterval int    `json:"AliveInterval"`
	ChannelNum    int    `json:"ChannelNum"`
	DeviceType    string `json:"DeviceType "`
	ExtraChannel  int    `json:"ExtraChannel"`
	Ret           int    `json:"Ret"`
	SessionID     string `json:"SessionID"`
}

// Account password login --
func (jDevice *JLinkDevice) Login() (dld *DeviceLoginData, err error) {
	parm := make(map[string]interface{})
	parm["UserName"] = jDevice.DevUsername
	parm["PassWord"] = jDevice.DevPassword
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return
	}
	resbody, err := http.HttpPost(dLoginUrl, jDevice.Jdtoken, postData)
	if err != nil {
		log.Println("HttpPost err:" + err.Error())
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		log.Println("Unmarshal err:" + err.Error())
		return nil, err
	}

	if resp.Code != 2000 {
		return nil, errors.New(resp.Msg)
	}
	resByre, resByteErr := json.Marshal(resp.Data)
	if resByteErr != nil {
		log.Println("Marshal err:" + resByteErr.Error())
		return nil, resByteErr
	}

	jsonRes := json.Unmarshal(resByre, &dld)
	if jsonRes != nil {
		log.Println("Unmarshal err:" + jsonRes.Error())
		return nil, jsonRes
	}
	return dld, nil
}

// Login with loginToken does not require account password verification
func (jDevice *JLinkDevice) DeviceLoginByToken(loginToken string) (dld *DeviceLoginData, err error) {
	parm := make(map[string]interface{})
	parm["LoginToken"] = loginToken
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return
	}
	resbody, err := http.HttpPost(dLoginUrl, jDevice.Jdtoken, postData)
	if err != nil {
		log.Println("HttpPost err:" + err.Error())
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return nil, err
	}

	if resp.Code != 2000 {
		return nil, errors.New(resp.Msg)
	}

	resByre, resByteErr := json.Marshal(resp.Data)
	if resByteErr != nil {
		log.Println("Marshal err:" + resByteErr.Error())
		return nil, resByteErr
	}

	jsonRes := json.Unmarshal(resByre, &dld)
	if jsonRes != nil {
		log.Println("Unmarshal err:" + jsonRes.Error())
		return nil, jsonRes
	}
	return dld, nil
}
