package jrgd

import (
	"encoding/json"
	"errors"
	"jlink-restful-golang-demo/http"
	"log"
)

type DeviceStatusResponse struct {
	Msg  string        `json:"msg"`
	Code int           `json:"code"`
	Data []interface{} `json:"data"`
}
type DeviceStatusData struct {
	Uuid         string      `json:"uuid"`
	Status       string      `json:"status"`
	Channel      interface{} `json:"channel"`
	Main         int         `json:"main"`
	Extra        int         `json:"extra"`
	WakeUpStatus string      `json:"wakeUpStatus"`
	WakeUpEnable string      `json:"wakeUpEnable"`
	WanIp        string      `json:"wanIp"`
}

// DeviceStatus--
func (jDevice *JLinkDevice) DeviceStatus() (dsd *DeviceStatusData, err error) {
	payload := http.NewStatusRequest(jDevice.Jdtoken)
	resbody, err := http.HttpPost(dStatusUrl, "", payload)
	if err != nil {
		log.Println("HttpPost err:" + err.Error())
		return nil, err
	}

	resp := &DeviceStatusResponse{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		log.Println("Unmarshal err:" + err.Error())
		return nil, err
	}

	if resp.Code != 2000 {
		return nil, errors.New(resp.Msg)
	}

	resByre, resByteErr := json.Marshal(resp.Data[0])
	if resByteErr != nil {
		log.Println("Marshal err:" + resByteErr.Error())
		return nil, resByteErr
	}

	err = json.Unmarshal(resByre, &dsd)
	if err != nil {
		log.Println("Unmarshal err:" + err.Error())
		return
	}

	return dsd, nil
}
