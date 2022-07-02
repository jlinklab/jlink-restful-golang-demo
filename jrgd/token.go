package jrgd

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	kTokenUrl = "https://tks.xmeye.net/v2/device/token/%s/%s.rs"
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
	Sn    string `json:"sn"`
	Token string `json:"token"`
}

// Get device token
func (jDevice *JLinkDevice) GetDeviceToken() (string, error) {
	tokenInfo := &tokenRequest{}
	tokenInfo.Sn = []string{jDevice.Sn}

	Tm := getTimeMillisUtil()
	Sign := getEncryptStr(jDevice.JClient.Uuid, jDevice.JClient.AppKey, jDevice.JClient.AppSecret, Tm, jDevice.JClient.Movecard)

	tokenUrl := fmt.Sprintf(kTokenUrl, Tm, Sign)

	reqs, _ := json.Marshal(tokenInfo)
	payload := strings.NewReader(string(reqs))
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
	msg, _ := url.PathUnescape(string(body))
	resp := &authResponse{}
	err = json.Unmarshal([]byte(msg), resp)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if resp.Code != 2000 {
		return "", errors.New(resp.Msg)
	}

	// fmt.Println(resp)
	jDevice.Jdtoken = resp.Data[0].Token
	return resp.Data[0].Token, nil
}

func getEncryptStr(uuid, appKey, secret, tm string, moved int) string {
	secret = uuid + appKey + secret + tm
	change := change(secret, moved)
	merge := mergeByte(secret, change)
	secret = fmt.Sprintf("%x", md5.Sum(merge))
	return secret
}

var counter int64 //
func getTimeMillisUtil() string {
	var value string
	counter++
	if counter < 10 {
		value = "000000"
	} else if counter < 100 {
		value = "00000"
	} else if counter < 1000 {
		value = "0000"
	} else if counter < 10000 {
		value = "000"
	} else if counter < 100000 {
		value = "00"
	} else if counter < 1000000 {
		value = "0"
	} else if counter >= 10000000 {
		counter = 1
		value = "000000"
	}

	return fmt.Sprintf("%s%d%d", value, counter, time.Now().UnixNano()/1000000)
}

func change(data string, mc int) string {
	src := []byte(data)
	length := len(src)
	for i := 0; i < length; i++ {
		temp := src[length-(i+1)]
		if (i % mc) > ((length - i) % mc) {
			temp = src[i]
		}

		src[i] = src[length-(i+1)]
		src[length-(i+1)] = temp
	}

	return string(src)
}

func mergeByte(encrypt, change string) []byte {
	length := len(encrypt)
	length2 := length * 2
	temp := make([]byte, length2)
	for i := 0; i < length; i++ {
		temp[i] = encrypt[i]
		temp[length2-1-i] = change[i]
	}

	return temp
}
