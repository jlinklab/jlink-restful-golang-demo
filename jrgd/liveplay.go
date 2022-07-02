package jrgd

import (
	"encoding/json"
	"errors"
	"jlink-restful-golang-demo/http"
	"log"
)

type LivePlayResponse struct {
	Msg  string       `json:"msg"`
	Code int          `json:"code"`
	Data LivePlayData `json:"data"`
}

type LivePlayData struct {
	Ret    int    `json:"Ret"`
	Url    string `json:"url"`
	RetMsg string `json:"retMsg"`
}

// pull video stream
func (jDevice *JLinkDevice) DeviceLivePlayUrl(stream, mediatype, protocol, channel string, jUser *JUser) (string, error) {
	parm := make(map[string]interface{})
	parm["mediaType"] = mediatype
	parm["channel"] = channel
	parm["stream"] = stream
	parm["protocol"] = protocol
	parm["username"] = jDevice.DevUsername
	parm["devPwd"] = jDevice.DevPassword
	parm["userToken"] = jUser.Jutoken
	// fmt.Println(parm)
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return "", err
	}
	// fmt.Println(string(postData))
	resbody, err := http.HttpPost(dlivestreamUrl, jDevice.Jdtoken, postData)
	if err != nil {
		log.Println("HttpPost err:" + err.Error())
		return "", err
	}
	// fmt.Println(string(resbody))
	resp := &LivePlayResponse{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		log.Println("Unmarshal err:" + err.Error())
		return "", err
	}

	if resp.Code == 2000 && resp.Data.Ret == 100 {
		return resp.Data.Url, nil
	}

	log.Println("DeviceLivePlayUrl err：" + resp.Msg)
	return "", errors.New(resp.Msg)
}
