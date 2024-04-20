package v3

import (
	"fmt"
	"io/ioutil"
	"jlink-restful-golang-demo/utils"
	"net/http"
	"strings"
)

func (jDevice *JLinkDevice) HttpPost(url string, payload *strings.Reader) ([]byte, error) {
	method := "POST"
	Tm := utils.GetTimeMillisUtil()
	Sign := utils.GetEncryptStr(jDevice.JClient.Uuid, jDevice.JClient.AppKey, jDevice.JClient.AppSecret, Tm, jDevice.JClient.Movecard)
	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("uuid", jDevice.JClient.Uuid)
	req.Header.Add("appKey", jDevice.JClient.AppKey)
	req.Header.Add("timeMillis", Tm)
	req.Header.Add("signature", Sign)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// fmt.Println(string(body))
	return body, nil

}
