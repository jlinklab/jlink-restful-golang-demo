package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"jlink-restful-golang-demo/utils"

	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	kTokenV3Url = "https://api.jftechws.com/ams/v3/deviceToken/%s/%s.rs"
)

type tokenRequest struct {
	Sn     []string `json:"sns"`
	UserId string   `json:"userId"`
}
type authResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data []data `json:"data"`
}
type data struct {
	Sn         string `json:"sn"`
	Token      string `json:"token"`
	LoginToken string `json:"loginToken"`
}

// Get device token
func (jDevice *JLinkDevice) GetDeviceTokenV3() (string, error) {
	tokenInfo := &tokenRequest{}
	tokenInfo.Sn = []string{jDevice.Sn}

	Tm := utils.GetTimeMillisUtil()
	Sign := utils.GetEncryptStr(jDevice.JClient.Uuid, jDevice.JClient.AppKey, jDevice.JClient.AppSecret, Tm, jDevice.JClient.Movecard)

	tokenUrl := fmt.Sprintf(kTokenV3Url, Tm, Sign)

	reqs, _ := json.Marshal(tokenInfo)
	key := utils.KeyFilter(Tm, jDevice.JClient.AppSecret)
	encryptedText, err := utils.AesEncrypt(string(reqs), key)
	if err != nil {
		return "", err
	}
	payload := strings.NewReader(encryptedText)
	req, _ := http.NewRequest("POST", tokenUrl, payload)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("uuid", jDevice.JClient.Uuid)
	req.Header.Add("appKey", jDevice.JClient.AppKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}

	defer res.Body.Close()

	var body, _ = ioutil.ReadAll(res.Body)
	decryptedText, err := utils.AesDecrypt(string(body), key)
	if err != nil {
		log.Println(err)
		return "", err
	}

	msg, _ := url.PathUnescape(string(decryptedText))
	fmt.Println(msg)
	resp := &authResponse{}
	err = json.Unmarshal([]byte(msg), resp)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if resp.Code != 2000 {
		return "", errors.New(resp.Msg)
	}
	if len(resp.Data) <= 0 {
		return "", errors.New("resp.Data is nil")
	}
	// fmt.Println(resp)
	jDevice.Jdtoken = resp.Data[0].Token
	return resp.Data[0].Token, nil
}
