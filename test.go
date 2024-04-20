package main

import (
	"fmt"
	"jlink-restful-golang-demo/config"
	"jlink-restful-golang-demo/utils"
	jrgd "jlink-restful-golang-demo/v3"
)

const (
	// test
	kDeviceId   = ""
	devUsername = ""
	devPassword = ""
	// test
	kUuid      = ""
	kAppKey    = ""
	KAppSecret = ""
	kMovecard  = 4
	userName   = ""
	password   = ""
)

func main() {
	// Initialize user platform fields
	jClient := jrgd.NewJLinkClient(kUuid, kAppKey, KAppSecret, kMovecard)
	jDevice := jrgd.NewJLinkDevice(jClient, kDeviceId, devUsername, devPassword)

	// // V3 版本接口

	// // 先绑定设备
	DeviceBind, err := jDevice.DeviceBind() // 绑定设备
	fmt.Println("DeviceBind", DeviceBind, err)
	// DeviceList, err := jDevice.DeviceList(1, 10) // 获取设备列表
	// fmt.Println(DeviceList, err)
	// 获取设备token
	DeviceToken, err := jDevice.GetDeviceTokenV3() // 获取设备token
	fmt.Println("DeviceToken", DeviceToken, err)
	// jDevice.LoginToken()
	// JDStatus, err := jDevice.DeviceStatus() // 获取设备状态
	// fmt.Println(JDStatus, err)
	// jDevice.DeviceWakeup("1") // 唤醒低功耗设备
	// // tailoredconfig.GetTailoredConfig(jDevice, []string{"bl"}) //设备离线配置获取

	// jDevice.DeviceLoginToken()    // 登录设备
	DeviceLogin, err := jDevice.DeviceLoginPassWord() // 登录设备
	fmt.Println("DeviceLogin", DeviceLogin, err)
	// jDevice.DeviceKeepalive()                          // 设备保活
	// jDevice.DeviceUsermanage()                         // 获取设备用户信息
	// jDevice.DeviceGetinfo(utils.SystemInfo)            // 获取设备配置 SystemInfo 4GInfo NetWork.RTMPALL StorageInfo
	// jDevice.DeviceAbility(utils.AbilitySystemFunction) // 获取设备能力集 SystemFunction EncodeCapability BlindCapability MotionArea Camera TalkAudioFormat MultiLanguage Intelligent
	jDConfig, err := jDevice.DeviceGetconfig(utils.ConfigNetWorkDAS) // 获取设备配置 AVEnc.Encode
	fmt.Println("jDConfig", jDConfig, err)

	if op, ok := jDConfig.(config.NetWorkDAS); ok {
		op.Enable = false
		nw := config.NetWorkDASResp{Name: utils.ConfigNetWorkDAS, NetWorkDAS: op}

		setConfigFlag, err := jDevice.SetConfig(nw)
		fmt.Println(setConfigFlag, err)
	}

	// // Equipment operation
	// // Take the PTZ direction control as an example
	// ptzDirectionControlDTO := new(opedev.PtzDirectionControlDTO)
	// ptzDirectionControlDTO.Command = utils.DirectionLeft
	// ptzDirectionControlDTO.Parameter.Channel = 0
	// ptzDirectionControlDTO.Parameter.Preset = -1
	// ptzDirectionControlDTO.Parameter.Step = 6
	// _, err = opedev.DeviceOperate(jDevice, ptzDirectionControlDTO)
	// fmt.Println(err)

	// jDevice.Capture(&jrgd.CaptureRequest{Channel: 0, PicType: 0})

	// doorLock.DoorLockRemoteUnlock(jDevice, "")	// 远程解锁

	// doorLock.DoorLockSetTempPassword(jDevice, []doorLock.TempPassword{ //临时密码
	// 	{
	// 		EndTime:   "2024-04-19 18:45:56",
	// 		Password:  "123456",
	// 		StartTime: "2024-04-19 10:15:56",
	// 		ValidNum:  1,
	// 	}})
	// STREAM_EXTRA := "1"
	// MEDIATYPE_HLS := "hls"
	// CHANNEL := "0"
	// PROTOCOL_TS := "ts"
	// Livestream, err := jDevice.Livestream(MEDIATYPE_HLS, CHANNEL, STREAM_EXTRA, PROTOCOL_TS)
	// fmt.Println(Livestream, err)

	jDevice.DeviceLogout() // 退出设备
	// DeviceUnBind, err := jDevice.DeviceUnBind() // 取消绑定
	// fmt.Println(DeviceUnBind, err)
	// fmt.Println(JDStatus, err)
}
