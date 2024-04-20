package utils

const (
	GwpUrl = "https://api.jftechws.com/gwp/v3/rtc"
)

const (
	DStatusUrl               = "/device/status"
	DWakeupUrl               = "/device/wakeup/"
	DLoginUrl                = "/device/login/"
	DGetInfoUrl              = "/device/getinfo/"
	DAbilityUrl              = "/device/getability/"
	DGetconfigUrl            = "/device/getconfig/"
	DSetconfigUrl            = "/device/setconfig/"
	DKeepaliveUrl            = "/device/keepalive/"
	DOpdevUrl                = "/device/opdev/"
	DcaptureUrl              = "/device/capture/"
	DlivestreamUrl           = "/device/livestream/"
	DgetTailoredConfig       = "/device/getTailoredConfig/"
	DsetTailoredConfig       = "/device/setTailoredConfig/"
	DUsermanage              = "/device/usermanage/"
	DBind                    = "/device/bind"
	DUnbind                  = "/device/unbind/"
	DList                    = "/device/list"
	DLogout                  = "/device/logout/"
	DQLoginToken             = "/device/queryLoginToken/"
	DGetStreamCount          = "/device/getPullStreamConnectionsCount/"
	DCloseLivestream         = "/device/closeLivestream/"
	DTalkbackUrl             = "/device/talkbackUrl/"
	DPlaybackUrl             = "/device/playbackUrl/"
	DGetVideoList            = "/device/getVideoList/"
	DGetVideoUrl             = "/device/getVideoUrl/"
	DGetPicUrl               = "/device/getPicUrl/"
	DSubserviceMessage       = "/device/subscribeMessage/"
	DUnSubserviceMessage     = "/device/unsubscribeMessage/"
	DGetDeviceAlarmList      = "/device/getDeviceAlarmList/"
	DDoorLockTransparent     = "/device/doorLockTransparent/"
	DDoorLockRemoteUnlock    = "/device/doorLockRemoteUnlock/"
	DDoorLockSetTempPassword = "/device/doorLockSetTempPassword/"
)

const (
	SystemInfo    = "SystemInfo"
	KeepAlive     = "KeepAlive"
	Logout        = "Logout"
	DirectionLeft = "DirectionLeft"
	OPSNAP        = "OPSNAP"
)

// Device Capability Set Related Protocols name
const (
	AbilitySystemFunction                            = "SystemFunction"
	AbilityEncodeCapability                          = "EncodeCapability"
	AbilityBlindCapability                           = "BlindCapability"
	AbilityMotionArea                                = "MotionArea"
	AbilityDDNSService                               = "DDNSService"
	AbilityComProtocol                               = "ComProtocol"
	AbilityCamera                                    = "Camera"
	AbilityTalkAudioFormat                           = "TalkAudioFormat"
	AbilityMultiLanguage                             = "MultiLanguage"
	AbilityUartProtocol                              = "UartProtocol"
	AbilityIntelligent                               = "Intelligent"
	AbilityCameraParamEx                             = "Camera.ParamEx"
	AbilityAbilitySupportVstd                        = "Ability.SupportVstd"
	AbilityAbilityDeviceDesc                         = "Ability.DeviceDesc"
	AbilityAbilityVGAResolution                      = "Ability.VGAResolution"
	AbilityAbilityGUIThemeList                       = "Ability.GUIThemeList"
	AbilityAbilityDigitalReal                        = "Ability.DigitalReal"
	AbilityVideoAbility                              = "VideoAbility"
	AbilityNetAbility                                = "NetAbility"
	AbilityEncStaticParam                            = "EncStaticParam"
	AbilityMaxPreRecord                              = "MaxPreRecord"
	AbilityMultiChannel                              = "MultiChannel"
	AbilityChannelSystemFunctionSupportPeaInHumanPed = "ChannelSystemFunction@SupportPeaInHumanPed"
	AbilityChannelSystemFunctionSupportFaceDetectV2  = "ChannelSystemFunction@SupportFaceDetectV2"
	AbilityChannelSystemFunctionSmartH264            = "ChannelSystemFunction@SmartH264"
)

// Device configuration related protocols name
const (
	ConfigAVEncEncode             = "AVEnc.Encode"
	ConfigAVEncVideoWidget        = "AVEnc.VideoWidget"
	ConfigAVEncVideoColor         = "AVEnc.VideoColor"
	ConfigAVEncSmartH264          = "AVEnc.SmartH264"
	ConfigfVideoVolume            = "fVideo.Volume"
	ConfigfVideoOSDInfo           = "fVideo.OSDInfo"
	ConfigRecord                  = "Record"
	ConfigStorageSnapshot         = "Storage.Snapshot"
	ConfigDetectBlindDetect       = "Detect.BlindDetect"
	ConfigDetectMotionDetect      = "Detect.MotionDetect"
	ConfigDetectLossDetect        = "Detect.LossDetect"
	ConfigChannelTitle            = "ChannelTitle"
	ConfigCamera                  = "Camera"
	ConfigDetectHumanDetection    = "Detect.HumanDetection"
	ConfigDetectHumanDetection0   = "Detect.HumanDetection.[0]"
	ConfigAbilityVoiceTipType     = "Ability.VoiceTipType"
	ConfigUartComm                = "Uart.Comm"
	ConfigUartPTZ                 = "Uart.PTZ"
	ConfigUartRS485               = "Uart.RS485"
	ConfigNetWorkNetCommon        = "NetWork.NetCommon"
	ConfigNetWorkNetDHCP          = "NetWork.NetDHCP"
	ConfigNetWorkNetEmail         = "NetWork.NetEmail"
	ConfigNetWorkNetNTP           = "NetWork.NetNTP"
	ConfigNetWorkNetDNS           = "NetWork.NetDNS"
	ConfigNetWorkNetFTP           = "NetWork.NetFTP"
	ConfigAlarmNetIPConflict      = "Alarm.NetIPConflict"
	ConfigAlarmNetAbort           = "Alarm.NetAbort"
	ConfigNetWorkUpnp             = "NetWork.Upnp"
	ConfigNetWorkNetIPFilter      = "NetWork.NetIPFilter"
	ConfigOPUTCTimeSetting        = "OPUTCTimeSetting"
	ConfigNetWorkNetOrder         = "NetWork.NetOrder"
	ConfigNetWorkDAS              = "NetWork.DAS"
	ConfigNetWorkPMS              = "NetWork.PMS"
	ConfigNetWorkNat              = "NetWork.Nat"
	ConfigNetWorkWifi             = "NetWork.Wifi"
	ConfigFVideoGUISet            = "fVideo.GUISet"
	ConfigStorageStoragePosition  = "Storage.StoragePosition"
	ConfigStorageStorageNotExist  = "Storage.StorageNotExist"
	ConfigStorageStorageLowSpace  = "Storage.StorageLowSpace"
	ConfigStorageStorageFailure   = "Storage.StorageFailure"
	ConfigGeneralGeneral          = "General.General"
	ConfigGeneralLocation         = "General.Location"
	ConfigCameraParam             = "Camera.Param"
	ConfigCameraParamEx           = "Camera.ParamEx"
	ConfigSimplifyEncode          = "Simplify.Encode"
	ConfigAVEncMultiChannelEncode = "AVEnc.MultiChannelEncode"
	ConfigfVideoTVAdjust          = "fVideo.TVAdjust"
	ConfigfVideoVideoOut          = "fVideo.VideoOut"
	ConfigfVideoPlay              = "fVideo.Play"
	ConfigfVideoAudioInFormat     = "fVideo.AudioInFormat"
	ConfigfVideoTour              = "fVideo.Tour"
	ConfigfVideoVideoOutPriority  = "fVideo.VideoOutPriority"
	ConfigAnalyze                 = "Analyze"
	ConfigVideoDec                = "VideoDec"
	ConfigVideoChannel            = "VideoChannel"
	ConfigAbilityEncodePower      = "Ability.EncodePower"
	ConfigGeneralSystemState      = "General.SystemState"
	ConfigMediaDecodeParam        = "Media.DecodeParam"
	ConfigAlarmLocalAlarm         = "Alarm.LocalAlarm"
	ConfigAlarmAlarmOut           = "Alarm.AlarmOut"
	ConfigUartPTZPreset           = "Uart.PTZPreset"
	ConfigUartPTZTour             = "Uart.PTZTour"
	ConfigNetWorkRemoteDeviceV3   = "NetWork.RemoteDeviceV3"
	ConfigNetWorkNetDDNS          = "NetWork.NetDDNS"
	ConfigNetWorkNetPPPOE         = "NetWork.NetPPPOE"
	ConfigNetWorkNetARSP          = "NetWork.NetARSP"
	ConfigNetWorkRemoteDevice     = "NetWork.RemoteDevice"
	ConfigNetWorkNet3G            = "NetWork.Net3G"
	ConfigNetWorkNetMobile        = "NetWork.NetMobile"
	ConfigNetWorkDigTimeSyn       = "NetWork.DigTimeSyn"
	ConfigNetWorkShortMsg         = "NetWork.ShortMsg"
	ConfigNetWorkMultimediaMsg    = "NetWork.MultimediaMsg"
	ConfigAlarmIPCAlarm           = "Alarm.IPCAlarm"
	ConfigGeneralAutoMaintain     = "General.AutoMaintain"
	ConfigUsers                   = "Users"
	ConfigGroups                  = "Groups"
	ConfigSystemExUserMap         = "System.ExUserMap"
	ConfigDevLP4GLedParameter     = "Dev.LP4GLedParameter"
	ConfigLPDevWorkMode           = "LPDev.WorkMode"
)

// Operation related agreements name
const (
	OPMachine                     = "OPMachine"
	OPDefaultConfig               = "OPDefaultConfig"
	StorageInfo                   = "StorageInfo"
	OPStorageManager              = "OPStorageManager"
	OPFileQuery                   = "OPFileQuery"
	OPTimeQuery                   = "OPTimeQuery"
	OPSystemConfig                = "OPSystemConfig"
	OPTimeSetting                 = "OPTimeSetting"
	OPPTZControl                  = "OPPTZControl"
	OPLogQuery                    = "OPLogQuery"
	OPVersionList                 = "OPVersionList"
	OPReqVersion                  = "OPReqVersion"
	OPCloudUpgradeStart           = "OPCloudUpgradeStart"
	OPCloudUpgradeStop            = "OPCloudUpgradeStop"
	OPCloudFileGetDownloadState   = "OPCloudFileGetDownloadState"
	OPCloudUpgradeGetBurnProgress = "OPCloudUpgradeGetBurnProgress"
)