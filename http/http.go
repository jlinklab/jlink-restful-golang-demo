package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	kUrl = "https://rds.bcloud365.net"
)

func HttpPost(method, token string, payload []byte) ([]byte, error) {
	client := &http.Client{Timeout: time.Second * 10}
	restfulUrl := fmt.Sprintf("%s%s%s", kUrl, method, token)
	if token == "" {
		restfulUrl = fmt.Sprintf("%s%s", kUrl, method)
	}
	fmt.Println(restfulUrl)

	req, _ := http.NewRequest("POST", restfulUrl, bytes.NewReader(payload))
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}

func NewStatusRequest(token string) []byte {
	queryList := &struct {
		Token []string `json:"token"`
	}{
		Token: []string{token},
	}
	payload, _ := json.Marshal(queryList)

	return payload
}
