package tailoredconfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"jlink-restful-golang-demo/utils"
	v3 "jlink-restful-golang-demo/v3"
	"log"
	"strings"
)

type SetTailoredConfigReq struct {
	SystemSleep struct {
		LightSleep struct {
			Mode  int `json:"Mode"`
			Param []struct {
				Value int `json:"Value"`
				Delay int `json:"Delay"`
			} `json:"Param"`
		} `json:"LightSleep"`
	} `json:"System.Sleep"`
}

func SetTailoredConfig(jDevice *v3.JLinkDevice, stcr *SetTailoredConfigReq) (bool, error) {
	parm := make(map[string]interface{})
	parm["sn"] = jDevice.Sn
	parm["data"] = stcr
	// fmt.Println(parm)
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return false, err
	}
	// fmt.Println(string(postData))
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DsetTailoredConfig, jDevice.Jdtoken)
	resbody, err := jDevice.HttpPost(url, strings.NewReader(string(postData)))
	if err != nil {
		log.Println("HttpPost err:" + err.Error())
		return false, err
	}
	fmt.Println(string(resbody))
	resp := &utils.XYResponse{}
	err = json.Unmarshal(resbody, resp)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
		return false, err
	}

	if resp.Code == 2000 {
		return true, nil
	}

	log.Println("SetTailoredConfig err：" + resp.Msg)
	return false, errors.New(resp.Msg)
}
