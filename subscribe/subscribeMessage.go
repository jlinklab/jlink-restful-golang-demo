package subscribe

import (
	"encoding/json"
	"fmt"
	"log"

	"jlink-restful-golang-demo/utils"
	v3 "jlink-restful-golang-demo/v3"
	"strings"
)

func SubscribeMessage(jDevice *v3.JLinkDevice, JUser *v3.JUser, callbackUrl, serviceType string) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DSubserviceMessage, jDevice.Jdtoken)
	parm := make(map[string]interface{})
	parm["sn"] = jDevice.Sn
	parm["callbackUrl"] = callbackUrl
	parm["userToken"] = JUser.Jutoken
	parm["serviceType"] = serviceType

	// fmt.Println(parm)
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
	}

	jDevice.HttpPost(url, strings.NewReader(string(postData)))
}

func UnSubscribeMessage(jDevice *v3.JLinkDevice, serviceType string) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DUnSubserviceMessage, jDevice.Jdtoken)
	parm := make(map[string]interface{})
	parm["sn"] = jDevice.Sn
	parm["serviceType"] = serviceType

	// fmt.Println(parm)
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
	}

	jDevice.HttpPost(url, strings.NewReader(string(postData)))
}

func GetDeviceAlarmList(jDevice *v3.JLinkDevice, startTime, endTime string) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DGetDeviceAlarmList, jDevice.Jdtoken)
	parm := make(map[string]interface{})
	parm["sn"] = jDevice.Sn
	parm["startTime"] = startTime
	parm["endTime"] = endTime

	// fmt.Println(parm)
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
	}

	jDevice.HttpPost(url, strings.NewReader(string(postData)))
}
