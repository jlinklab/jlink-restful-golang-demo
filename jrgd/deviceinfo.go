package jrgd

import (
	"encoding/json"
	"errors"
	"jlink-restful-golang-demo/http"
	"log"
)

type DevicesInfo struct {
	Name       string      `json:"name"`
	Ret        int         `json:"ret"`
	SessionID  string      `json:"sessionID"`
	SystemInfo SystemInfos `json:"systemInfo"`
}

type SystemInfos struct {
	AlarmInChannel  int    `json:"alarmInChannel"`
	AlarmOutChannel int    `json:"alarmOutChannel"`
	AudioInChannel  int    `json:"audioInChannel"`
	BuildTime       string `json:"buildTime"`
	CombineSwitch   int    `json:"combineSwitch"`
	DeviceRunTime   string `json:"deviceRunTime"`
	DigChannel      int    `json:"digChannel"`
	EncryptVersion  string `json:"encryptVersion"`
	ExtraChannel    int    `json:"extraChannel"`
	HardWare        string `json:"hardWare"`
	HardWareVersion string `json:"hardWareVersion"`
	SerialNo        string `json:"serialNo"`
	SoftWareVersion string `json:"softWareVersion"`
	TalkInChannel   int    `json:"talkInChannel"`
	TalkOutChannel  int    `json:"talkOutChannel"`
	UpdataTime      string `json:"updataTime"`
	UpdataType      string `json:"updataType"`
	VideoInChannel  int    `json:"videoInChannel"`
	VideoOutChannel int    `json:"videoOutChannel"`
}

// get DeviceInfo
func (jDevice *JLinkDevice) DeviceInfo(name string) (di *DevicesInfo, err error) {
	parm := make(map[string]interface{})
	parm["Name"] = name
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return
	}
	resbody, err := http.HttpPost(dGetInfoUrl, jDevice.Jdtoken, postData)
	if err != nil {
		log.Println("HttpPost err:" + err.Error())
		return
	}

	resp := &Response{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		log.Println("Unmarshal err:" + err.Error())
		return
	}

	if resp.Code != 2000 {
		return nil, errors.New(resp.Msg)
	}

	resByre, resByteErr := json.Marshal(resp.Data)
	if resByteErr != nil {
		log.Println("Marshal err:" + resByteErr.Error())
		return nil, resByteErr
	}

	jsonRes := json.Unmarshal(resByre, &di)
	if jsonRes != nil {
		log.Println("Unmarshal err:" + jsonRes.Error())
		return nil, jsonRes
	}

	return di, nil
}
