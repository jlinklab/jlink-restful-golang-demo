package internal

const (
	// Iot 服务
	KRestfulApiIotFeeder               = "/v2/rtc/device/feeder/"
	KRestfulApiIotPropGet              = "/v2/rtc/device/iotPropGet/"
	KRestfulApiIotPropSet              = "/v2/rtc/device/iotPropSet/"
	KRestfulApiIotTimer                = "/v2/rtc/device/iotTimer/"
	KRestfulApiIotPropInterlock        = "/v2/rtc/device/iotPropInterlock/"
	KRestfulApiDoorLockSetTempPassword = "/v2/rtc/device/doorLockSetTempPassword/"
	KRestfulApiDoorLockRemoteUnlock    = "/v2/rtc/device/doorLockRemoteUnlock/"
	KRestfulApiDoorLockTransparent     = "/v2/rtc/device/doorLockTransparent/"
	KRestfulApiIotLightSourceSet       = "/v2/rtc/device/iotLightSourceSet/"
	KRestfulApiIotLightSourceGet       = "/v2/rtc/device/iotLightSourceGet/"
	KRestfulApiIotLightSourceControl   = "/v2/rtc/device/iotLightSourceControl/"
)

// 协议网关错误码
const (
	KCodeSuccess      = 2000
	KCodeAlreadyLogin = 2004
	KCodeParamErr     = 4000
	KCodeDevTokenErr  = 4001
	KCodeDevTokenExp  = 4002
	KCodeUserTokenErr = 4003
	KCodeUserTokenExp = 4004

	KCodeIllegalRequest            = 4100
	KCodeDeviceNotOnline           = 4101
	KCodeCaptureFailed             = 4102
	KCodeWakeUpFailed              = 4103
	KCodeTimeout                   = 4104
	KCodeGwmFailed                 = 4105
	KCodeVoesFailed                = 4106
	KCodeSubFailed                 = 4107
	KCodeSmsFailed                 = 4108
	KCodeAudioFailed               = 4109
	KCodeCloudServiceNotSupport    = 4110
	KCodeCloudServiceNotEnable     = 4111
	KCodeCloudServiceExpiration    = 4112
	KCodeCloudServiceFailed        = 4113
	KCodeCloudServiceInsufficient  = 4114
	KCodeImageGetFailed            = 4115
	KCodeDeviceNotFound            = 4116
	KCodeCloudServiceNotPermission = 4117
	kCodeDeviceConnectTimeout      = 4118
	KCodeServerFailed              = 5000
)

const (
	KDeviceRetSuccess = 100
)
