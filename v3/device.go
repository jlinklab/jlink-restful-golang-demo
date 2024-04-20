package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"jlink-restful-golang-demo/utils"
	"log"
	"strconv"
	"strings"
)

// 绑定设备
func (jDevice *JLinkDevice) DeviceBind() (interface{}, error) {
	url := utils.GwpUrl + utils.DBind
	payload := strings.NewReader(`{"sn": "` + jDevice.Sn + `","username": "` + jDevice.DevUsername + `","password": "` + jDevice.DevPassword + `"}`)
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

// 解绑设备
func (jDevice *JLinkDevice) DeviceUnBind() (interface{}, error) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DUnbind, jDevice.Jdtoken)
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

// 设备列表
func (jDevice *JLinkDevice) DeviceList(page, limit int) (interface{}, error) {
	url := utils.GwpUrl + utils.DList
	payload := strings.NewReader(`{"page": ` + strconv.Itoa(page) + `,"limit": ` + strconv.Itoa(limit) + `}`)
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

// 设备状态
func (jDevice *JLinkDevice) DeviceStatus() (interface{}, error) {
	url := utils.GwpUrl + utils.DStatusUrl
	payload := strings.NewReader(`{"deviceTokenList": ["` + jDevice.Jdtoken + `"]}`)
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

// 设备保活
func (jDevice *JLinkDevice) DeviceKeepalive() (interface{}, error) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DKeepaliveUrl, jDevice.Jdtoken)
	payload := strings.NewReader(`{"Name":"KeepAlive"}`)
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

// 设备登出
func (jDevice *JLinkDevice) DeviceLogout() (interface{}, error) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DLogout, jDevice.Jdtoken)
	payload := strings.NewReader(`{"Name":"Logout"}`)
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

// 获取设备信息
func (jDevice *JLinkDevice) DeviceGetinfo(name string) (interface{}, error) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DGetInfoUrl, jDevice.Jdtoken)
	payload := strings.NewReader(`{"Name":"` + name + `"}`)
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

// 设备系统能力集
func (jDevice *JLinkDevice) DeviceAbility(name string) (interface{}, error) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DAbilityUrl, jDevice.Jdtoken)
	payload := strings.NewReader(`{"Name":"` + name + `"}`)
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

// 低功耗设备唤醒 hostedControl是否唤醒主控（空值或0：不唤醒主控；1：唤醒主控）
func (jDevice *JLinkDevice) DeviceWakeup(hostedControl string) (interface{}, error) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DWakeupUrl, jDevice.Jdtoken)
	payload := strings.NewReader(`{"sn":"` + jDevice.Sn + `","hostedControl":"` + hostedControl + `"}`)
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

// 设备登陆
func (jDevice *JLinkDevice) DeviceLoginPassWord() (interface{}, error) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DLoginUrl, jDevice.Jdtoken)
	payload := strings.NewReader(`{"EncryptType":"MD5","LoginType":"DVRIP-Web","Name":"generalinfo","UserName":"` + jDevice.DevUsername + `","PassWord":"` + jDevice.DevPassword + `"}`)
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

// 设备登陆
func (jDevice *JLinkDevice) DeviceLoginToken() (interface{}, error) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DLoginUrl, jDevice.Jdtoken)
	payload := strings.NewReader(`{"EncryptType":"TOKEN","LoginType":"DVRIP-Web","Name":"generalinfo","LoginToken":"` + jDevice.JdLoginToken + `"}`)
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

type LoginTokenResponse struct {
	Msg  string         `json:"msg"`
	Code int            `json:"code"`
	Data LoginTokenData `json:"data"`
}
type LoginTokenData struct {
	Sn         string `json:"sn"`
	LoginToken string `json:"loginToken"`
}

// 获取设备登陆token
func (jDevice *JLinkDevice) LoginToken() error {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DQLoginToken, jDevice.Jdtoken)
	payload := strings.NewReader(`{"sn":"` + jDevice.Sn + `"}`)
	body, err := jDevice.HttpPost(url, payload)
	if err != nil {
		return err
	}
	// jDevice.Jdtoken = string(body)

	resp := &LoginTokenResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return err
	}
	if resp.Code != 2000 {
		return errors.New(resp.Msg)
	}
	jDevice.JdLoginToken = resp.Data.LoginToken
	return nil
}

type CaptureRequest struct {
	Channel int
	PicType int
}

// 设备抓图
func (jDevice *JLinkDevice) Capture(captureRequest *CaptureRequest) (interface{}, error) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DcaptureUrl, jDevice.Jdtoken)
	parm := make(map[string]interface{})
	parm["Name"] = "OPSNAP"
	parm["OPSNAP"] = captureRequest
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
