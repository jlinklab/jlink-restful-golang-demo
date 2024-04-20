package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"jlink-restful-golang-demo/utils"
	"log"
	"strings"
)

func (jDevice *JLinkDevice) GetPullStreamConnectionsCount() (interface{}, error) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DGetStreamCount, jDevice.Jdtoken)
	payload := strings.NewReader(`{"sn":"` + jDevice.Sn + `"}`)
	resbody, err := jDevice.HttpPost(url, payload)
	if err != nil {
		return nil, err
	}
	// fmt.Println(string(resbody))
	resp := &utils.Response{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		return nil, err
	}

	if resp.Code == 2000 {
		return resp.Data, nil
	}
	return false, errors.New(resp.Msg)
}

func (jDevice *JLinkDevice) Livestream(mediaType, channel, stream, protocol string) (interface{}, error) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DlivestreamUrl, jDevice.Jdtoken)

	parm := make(map[string]interface{})
	parm["mediaType"] = mediaType
	parm["channel"] = channel
	parm["stream"] = stream
	parm["protocol"] = protocol
	parm["username"] = jDevice.DevUsername
	parm["password"] = jDevice.DevPassword
	// parm["OPPTZControl"] = pdcd

	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return nil, err
	}
	resbody, err := jDevice.HttpPost(url, strings.NewReader(string(postData)))
	if err != nil {
		return nil, err
	}
	// fmt.Println(string(resbody))
	resp := &utils.Response{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		return nil, err
	}

	if resp.Code == 2000 {
		return resp.Data, nil
	}
	return false, errors.New(resp.Msg)
}

func (jDevice *JLinkDevice) CloseLivestream(mediaType, channel, stream, protocol string) (interface{}, error) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DCloseLivestream, jDevice.Jdtoken)

	parm := make(map[string]interface{})
	parm["channel"] = channel
	parm["stream"] = stream
	parm["username"] = jDevice.DevUsername
	parm["password"] = jDevice.DevPassword

	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return nil, err
	}
	resbody, err := jDevice.HttpPost(url, strings.NewReader(string(postData)))

	if err != nil {
		return nil, err
	}
	// fmt.Println(string(resbody))
	resp := &utils.Response{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		return nil, err
	}

	if resp.Code == 2000 {
		return resp.Data, nil
	}
	return false, errors.New(resp.Msg)
}
func (jDevice *JLinkDevice) TalkbackUrl(channel string) (interface{}, error) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DTalkbackUrl, jDevice.Jdtoken)

	parm := make(map[string]interface{})
	parm["mediaType"] = "rtmp"
	parm["channel"] = channel
	parm["username"] = jDevice.DevUsername
	parm["password"] = jDevice.DevPassword
	parm["audioCode"] = "aac"

	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return nil, err
	}
	resbody, err := jDevice.HttpPost(url, strings.NewReader(string(postData)))

	if err != nil {
		return nil, err
	}
	// fmt.Println(string(resbody))
	resp := &utils.Response{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		return nil, err
	}

	if resp.Code == 2000 {
		return resp.Data, nil
	}
	return false, errors.New(resp.Msg)
}

func (jDevice *JLinkDevice) PlaybackUrl(channel, streamType, startTime, endTime, filename string) (interface{}, error) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DPlaybackUrl, jDevice.Jdtoken)
	parm := make(map[string]interface{})
	parm["channel"] = channel
	parm["streamType"] = streamType
	parm["startTime"] = startTime
	parm["endTime"] = endTime
	parm["username"] = jDevice.DevUsername
	parm["devPwd"] = jDevice.DevPassword
	parm["FileName"] = filename

	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return nil, err
	}
	resbody, err := jDevice.HttpPost(url, strings.NewReader(string(postData)))

	if err != nil {
		return nil, err
	}
	// fmt.Println(string(resbody))
	resp := &utils.Response{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		return nil, err
	}

	if resp.Code == 2000 {
		return resp.Data, nil
	}
	return false, errors.New(resp.Msg)
}
func (jDevice *JLinkDevice) GetVideoList(startTime, stopTime string) (interface{}, error) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DGetVideoList, jDevice.Jdtoken)
	parm := make(map[string]interface{})
	parm["startTime"] = startTime
	parm["stopTime"] = stopTime
	parm["sn"] = jDevice.Sn

	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return nil, err
	}
	resbody, err := jDevice.HttpPost(url, strings.NewReader(string(postData)))

	if err != nil {
		return nil, err
	}
	// fmt.Println(string(resbody))
	resp := &utils.Response{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		return nil, err
	}

	if resp.Code == 2000 {
		return resp.Data, nil
	}
	return false, errors.New(resp.Msg)
}

// GetVideoUrl
func (jDevice *JLinkDevice) GetVideoUrl(startTime, stopTime string) (interface{}, error) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DGetVideoUrl, jDevice.Jdtoken)
	parm := make(map[string]interface{})
	parm["startTime"] = startTime
	parm["stopTime"] = stopTime
	parm["sn"] = jDevice.Sn

	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return nil, err
	}
	resbody, err := jDevice.HttpPost(url, strings.NewReader(string(postData)))

	if err != nil {
		return nil, err
	}
	// fmt.Println(string(resbody))
	resp := &utils.Response{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		return nil, err
	}

	if resp.Code == 2000 {
		return resp.Data, nil
	}
	return false, errors.New(resp.Msg)
}
func (jDevice *JLinkDevice) GetPicUrl(channel string, alarmIds []string) (interface{}, error) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DGetPicUrl, jDevice.Jdtoken)
	parm := make(map[string]interface{})
	parm["channel"] = channel
	parm["alarmIds"] = alarmIds
	parm["sn"] = jDevice.Sn

	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return nil, err
	}
	resbody, err := jDevice.HttpPost(url, strings.NewReader(string(postData)))

	if err != nil {
		return nil, err
	}
	// fmt.Println(string(resbody))
	resp := &utils.Response{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		return nil, err
	}

	if resp.Code == 2000 {
		return resp.Data, nil
	}
	return false, errors.New(resp.Msg)
}
