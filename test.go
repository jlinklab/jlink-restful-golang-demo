package main

import (
	"fmt"
	"jlink-restful-golang-demo/jrgd"
)

const (
	// test
	kDeviceId   = ""
	kUuid       = ""
	kAppKey     = ""
	KAppSecret  = ""
	kMovecard   = 2
	devUsername = ""
	devPassword = ""
	userName    = ""
	password    = ""
)

func main() {
	// Initialize user platform fields
	jClient := jrgd.NewJLinkClient(kUuid, kAppKey, KAppSecret, kMovecard)
	jDevice := jrgd.NewJLinkDevice(jClient, kDeviceId, devUsername, devPassword)

	// GetDeviceToken
	JDToken, err := jDevice.GetDeviceToken()
	fmt.Println(JDToken, err)

	//DeviceStatus
	JDStatus, err := jDevice.DeviceStatus()
	fmt.Println(JDStatus, err)

	//WakeUp
	// wakeupFlag := jDevice.WakeUp()
	// fmt.Println(wakeupFlag)

	// Login
	// 1. Account password login
	jDLogin, err := jDevice.Login()
	fmt.Println(jDLogin, err)
	// 2.Login with loginToken does not require account password verification
	// loginToken := ""
	// jDLogin, err = jDevice.DeviceLoginByToken(loginToken)
	// fmt.Println(jDLogin)

	// DeviceInfo
	// SystemInfo
	// devInfo, err := jDevice.DeviceInfo(jrgd.SystemInfo)
	// fmt.Println(devInfo, err)

	// DeviceAbility
	// After the device is successfully logged in, the device capability set is obtained. Take the acquisition of the system capability set as an example
	// jDAbility, err := jDevice.DeviceAbility(jrgd.AbilityChannelSystemFunctionSmartH264)
	// fmt.Println(jDAbility, err)

	// GetConfig
	// Take obtaining the device active registration service DAS configuration as an example
	jDConfig, err := jDevice.GetConfig(jrgd.ConfigNetWorkDAS)
	fmt.Println(jDConfig)

	// // SetConfig
	// // Take the device actively registering the DAS service as an example enable: enable serverAddr: server address Port: server port

	// if op, ok := jDConfig.(config.NetWorkDAS); ok {
	// 	op.Enable = false
	// 	nw := config.NetWorkDASResp{Name: jrgd.ConfigNetWorkDAS, NetWorkDAS: op}
	// 	setConfigFlag, err := jDevice.SetConfig(nw)
	// 	fmt.Println(setConfigFlag, err)
	// }

	// nw := config.NetWorkDASResp{Name: jrgd.ConfigNetWorkDAS, NetWorkDAS: config.NetWorkDAS{Enable: true}}
	// setConfigFlag, err := jDevice.SetConfig(nw)
	// fmt.Println(setConfigFlag, err)

	// Keepalive
	// keepalive, err := jDevice.Keepalive()
	// fmt.Println(keepalive)

	// Equipment operation
	// Take the PTZ direction control as an example
	// ptzDirectionControlDTO := new(opedev.PtzDirectionControlDTO)
	// ptzDirectionControlDTO.Command = jrgd.DirectionLeft
	// ptzDirectionControlDTO.Parameter.Channel = 0
	// ptzDirectionControlDTO.Parameter.Preset = 65535
	// ptzDirectionControlDTO.Parameter.Step = 6
	// _, err = jDevice.DeviceOperate(ptzDirectionControlDTO)
	// fmt.Println(err)

	// Get live
	// Take obtaining the mainstream live broadcast address of the device in hls format as an example; rs Account represents the user registered on the open platform through the RS interface, and rs Pass represents the user password registered on the open platform through the RS interface
	// jUser := jrgd.NewJLinkUser(jClient, userName, password)
	// usertoken, err := jUser.GetUserToken()
	// fmt.Println(usertoken)
	// STREAM_EXTRA := "1"
	// MEDIATYPE_HLS := "hls"
	// CHANNEL := "0"
	// PROTOCOL_TS := "ts"
	// liveStream, err := jDevice.DeviceLivePlayUrl(STREAM_EXTRA, MEDIATYPE_HLS, PROTOCOL_TS, CHANNEL, jUser)
	// fmt.Println(liveStream)
	// channel := 0
	// captureUrl, err := jDevice.Capture(channel)
	// fmt.Println(captureUrl)
}
