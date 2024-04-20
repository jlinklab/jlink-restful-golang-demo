package doorlock

import (
	"encoding/json"
	"fmt"
	"jlink-restful-golang-demo/utils"
	v3 "jlink-restful-golang-demo/v3"
	"log"
	"strings"
)

type OPDoorLockProCmdReq struct {
	Cmd string `json:"Cmd"`
}

func DoorLockTransparent(jDevice *v3.JLinkDevice, OPDoorLockProCmdReq OPDoorLockProCmdReq) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DDoorLockTransparent, jDevice.Jdtoken)
	parm := make(map[string]interface{})
	parm["Name"] = "OPDoorLockProCmd"
	parm["OPDoorLockProCmd"] = OPDoorLockProCmdReq

	// fmt.Println(parm)
	postData, err := json.Marshal(parm)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
	}

	jDevice.HttpPost(url, strings.NewReader(string(postData)))
}

type DoorLockReq struct {
	Sn    string `json:"sn"`
	Props Props  `json:"props"`
}
type Props struct {
	DoorLock DoorLock `json:"doorLock"`
}
type DoorLock struct {
	// RemoteUnlock RemoteUnlock   `json:"remoteUnlock,omitempty"`
	TempPassword []TempPassword `json:"tempPassword,omitempty"`
}
type RemoteUnlock struct {
	Password string `json:"password"`
}

func DoorLockRemoteUnlock(jDevice *v3.JLinkDevice, password string) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DDoorLockRemoteUnlock, jDevice.Jdtoken)
	doorlockreq := DoorLockReq{
		Sn: jDevice.Sn,
		Props: Props{
			DoorLock: DoorLock{
				// RemoteUnlock: RemoteUnlock{
				// 	Password: password,
				// },
			},
		},
	}

	postData, err := json.Marshal(doorlockreq)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
	}
	fmt.Println(string(postData))
	jDevice.HttpPost(url, strings.NewReader(string(postData)))
}

type TempPassword struct {
	EndTime   string `json:"endTime"`
	Password  string `json:"password"`
	StartTime string `json:"startTime"`
	ValidNum  int    `json:"validNum"`
}

func DoorLockSetTempPassword(jDevice *v3.JLinkDevice, TempPassword []TempPassword) {
	url := fmt.Sprintf("%s%s", utils.GwpUrl+utils.DDoorLockSetTempPassword, jDevice.Jdtoken)
	doorlockreq := DoorLockReq{
		Sn: jDevice.Sn,
		Props: Props{
			DoorLock: DoorLock{
				TempPassword: TempPassword,
			},
		},
	}
	postData, err := json.Marshal(doorlockreq)
	if err != nil {
		log.Println("Marshal err:" + err.Error())
	}
	fmt.Println(string(postData))
	jDevice.HttpPost(url, strings.NewReader(string(postData)))
}
