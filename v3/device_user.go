package v3

import (
	"fmt"
	"jlink-restful-golang-demo/utils"
	"strings"
)

// 获取设备用户信息
func (jDevice *JLinkDevice) DeviceUsermanage() {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DUsermanage, jDevice.Jdtoken)
	payload := strings.NewReader(`{"Name":"USERS"}`)
	jDevice.HttpPost(url, payload)
}
