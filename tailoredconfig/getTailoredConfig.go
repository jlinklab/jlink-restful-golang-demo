package tailoredconfig

import (
	"encoding/json"
	"fmt"
	"jlink-restful-golang-demo/utils"
	v3 "jlink-restful-golang-demo/v3"
	"log"
	"strings"
)

func GetTailoredConfig(jDevice *v3.JLinkDevice, configs []string) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DgetTailoredConfig, jDevice.Jdtoken)

	parm := make(map[string]interface{})
	parm["sn"] = jDevice.Sn
	parm["configs"] = configs
	// fmt.Println(parm)
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
	}

	jDevice.HttpPost(url, strings.NewReader(string(postData)))
}
