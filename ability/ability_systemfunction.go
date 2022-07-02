package ability

type SystemFunctionResp struct {
	Name           string         `json:"Name"`
	Ret            int            `json:"Ret"`
	SessionID      string         `json:"SessionID"`
	SystemFunction SystemFunction `json:"SystemFunction"`
	RetMsg         string         `json:"RetMsg"`
}

type AlarmFunction struct {
	Four33Alarm       bool `json:"433Alarm"`
	AlarmConfig       bool `json:"AlarmConfig"`
	BlindDetect       bool `json:"BlindDetect"`
	BlurCheck         bool `json:"BlurCheck"`
	Consumer433Alarm  bool `json:"Consumer433Alarm"`
	ConsumerRemote    bool `json:"ConsumerRemote"`
	IPCAlarm          bool `json:"IPCAlarm"`
	LossDetect        bool `json:"LossDetect"`
	MotionDetect      bool `json:"MotionDetect"`
	NetAbort          bool `json:"NetAbort"`
	NetAbortExtend    bool `json:"NetAbortExtend"`
	NetAlarm          bool `json:"NetAlarm"`
	NetIPConflict     bool `json:"NetIpConflict"`
	NewVideoAnalyze   bool `json:"NewVideoAnalyze"`
	SensorAlarmCenter bool `json:"SensorAlarmCenter"`
	SerialAlarm       bool `json:"SerialAlarm"`
	StorageFailure    bool `json:"StorageFailure"`
	StorageLowSpace   bool `json:"StorageLowSpace"`
	StorageNotExist   bool `json:"StorageNotExist"`
	VideoAnalyze      bool `json:"VideoAnalyze"`
}

type CommFunction struct {
	CommRS232 bool `json:"CommRS232"`
	CommRS485 bool `json:"CommRS485"`
}

type EncodeFunction struct {
	CombineStream     bool `json:"CombineStream"`
	DoubleStream      bool `json:"DoubleStream"`
	IFrameRange       bool `json:"IFrameRange"`
	IntelligentEncode bool `json:"IntelligentEncode"`
	LowBitRate        bool `json:"LowBitRate"`
	SmartH264         bool `json:"SmartH264"`
	SnapStream        bool `json:"SnapStream"`
	WaterMark         bool `json:"WaterMark"`
}

type InputMethod struct {
	NoSupportChinese bool `json:"NoSupportChinese"`
}

type MobileDVR struct {
	CarPlateSet    bool `json:"CarPlateSet"`
	DVRBootType    bool `json:"DVRBootType"`
	DelaySet       bool `json:"DelaySet"`
	GpsTiming      bool `json:"GpsTiming"`
	StatusExchange bool `json:"StatusExchange"`
}

type NetServerFunction struct {
	IPAdaptive            bool `json:"IPAdaptive"`
	MACProtocol           bool `json:"MACProtocol"`
	MonitorPlatform       bool `json:"MonitorPlatform"`
	NATProtocol           bool `json:"NATProtocol"`
	Net3G                 bool `json:"Net3G"`
	Net4G                 bool `json:"Net4G"`
	NetARSP               bool `json:"NetARSP"`
	NetAlarmCenter        bool `json:"NetAlarmCenter"`
	NetAnJuP2P            bool `json:"NetAnJuP2P"`
	NetBaiduCloud         bool `json:"NetBaiduCloud"`
	NetBjlThy             bool `json:"NetBjlThy"`
	NetDAS                bool `json:"NetDAS"`
	NetDDNS               bool `json:"NetDDNS"`
	NetDHCP               bool `json:"NetDHCP"`
	NetDNS                bool `json:"NetDNS"`
	NetDataLink           bool `json:"NetDataLink"`
	NetEmail              bool `json:"NetEmail"`
	NetFTP                bool `json:"NetFTP"`
	NetGlobalEyes         bool `json:"NetGlobalEyes"`
	NetGodEyeAlarm        bool `json:"NetGodEyeAlarm"`
	NetIPFilter           bool `json:"NetIPFilter"`
	NetIPv6               bool `json:"NetIPv6"`
	NetKaiCong            bool `json:"NetKaiCong"`
	NetKeyboard           bool `json:"NetKeyboard"`
	NetLocalSdkPlatform   bool `json:"NetLocalSdkPlatform"`
	NetMobile             bool `json:"NetMobile"`
	NetMobileWatch        bool `json:"NetMobileWatch"`
	NetMutliCast          bool `json:"NetMutliCast"`
	NetNTP                bool `json:"NetNTP"`
	NetNat                bool `json:"NetNat"`
	NetOpenVPN            bool `json:"NetOpenVPN"`
	NetPMS                bool `json:"NetPMS"`
	NetPMSV2              bool `json:"NetPMSV2"`
	NetPPPoE              bool `json:"NetPPPoE"`
	NetPhoneMultimediaMsg bool `json:"NetPhoneMultimediaMsg"`
	NetPhoneShortMsg      bool `json:"NetPhoneShortMsg"`
	NetPlatMega           bool `json:"NetPlatMega"`
	NetPlatShiSou         bool `json:"NetPlatShiSou"`
	NetPlatVVEye          bool `json:"NetPlatVVEye"`
	NetPlatXingWang       bool `json:"NetPlatXingWang"`
	NetRTSP               bool `json:"NetRTSP"`
	NetSPVMN              bool `json:"NetSPVMN"`
	NetSPVMNSIP           bool `json:"NetSPVMNSIP"`
	NetTUTKIOTC           bool `json:"NetTUTKIOTC"`
	NetUPNP               bool `json:"NetUPNP"`
	NetVPN                bool `json:"NetVPN"`
	NetWifi               bool `json:"NetWifi"`
	NetWifiMode           bool `json:"NetWifiMode"`
	OnvifPwdCheckout      bool `json:"OnvifPwdCheckout"`
	PlatFormGBeyes        bool `json:"PlatFormGBeyes"`
	RTMP                  bool `json:"RTMP"`
	XMHeartBeat           bool `json:"XMHeartBeat"`
}

type OtherFunction struct {
	AlterDigitalName             bool `json:"AlterDigitalName"`
	DownLoadPause                bool `json:"DownLoadPause"`
	HddLowSpaceUseMB             bool `json:"HddLowSpaceUseMB"`
	HideDigital                  bool `json:"HideDigital"`
	MusicFilePlay                bool `json:"MusicFilePlay"`
	NOHDDRECORD                  bool `json:"NOHDDRECORD"`
	NotSupportAH                 bool `json:"NotSupportAH"`
	NotSupportAV                 bool `json:"NotSupportAV"`
	NotSupportTalk               bool `json:"NotSupportTalk"`
	SDsupportRecord              bool `json:"SDsupportRecord"`
	ShowAlarmLevelRegion         bool `json:"ShowAlarmLevelRegion"`
	ShowFalseCheckTime           bool `json:"ShowFalseCheckTime"`
	SupportAbnormitySendMail     bool `json:"SupportAbnormitySendMail"`
	SupportAlarmVoiceTips        bool `json:"SupportAlarmVoiceTips"`
	SupportBT                    bool `json:"SupportBT"`
	SupportBallCameraTrackDetect bool `json:"SupportBallCameraTrackDetect"`
	SupportBulbAlarmLightOn      bool `json:"SupportBulbAlarmLightOn"`
	SupportC7Platform            bool `json:"SupportC7Platform"`
	SupportCamareStyle           bool `json:"SupportCamareStyle"`
	SupportCameraMotorCtrl       bool `json:"SupportCameraMotorCtrl"`
	SupportCameraWhiteLight      bool `json:"SupportCameraWhiteLight"`
	SupportCfgCloudupgrade       bool `json:"SupportCfgCloudupgrade"`
	SupportCloudUpgrade          bool `json:"SupportCloudUpgrade"`
	SupportCommDataUpload        bool `json:"SupportCommDataUpload"`
	SupportConsSensorAlarmLink   bool `json:"SupportConsSensorAlarmLink"`
	SupportContinueUpgrade       bool `json:"SupportContinueUpgrade"`
	SupportCorridorMode          bool `json:"SupportCorridorMode"`
	SupportCustomOemInfo         bool `json:"SupportCustomOemInfo"`
	SupportDeviceInfoNew         bool `json:"SupportDeviceInfoNew"`
	SupportDigitalEncode         bool `json:"SupportDigitalEncode"`
	SupportDigitalPre            bool `json:"SupportDigitalPre"`
	SupportDimenCode             bool `json:"SupportDimenCode"`
	SupportEncodeAddBeep         bool `json:"SupportEncodeAddBeep"`
	SupportFTPTest               bool `json:"SupportFTPTest"`
	SupportFaceDetect            bool `json:"SupportFaceDetect"`
	SupportFishEye               bool `json:"SupportFishEye"`
	SupportImpRecord             bool `json:"SupportImpRecord"`
	SupportIntelligentPlayBack   bool `json:"SupportIntelligentPlayBack"`
	SupportLowLuxMode            bool `json:"SupportLowLuxMode"`
	SupportMailTest              bool `json:"SupportMailTest"`
	SupportMaxPlayback           bool `json:"SupportMaxPlayback"`
	SupportModifyFrontcfg        bool `json:"SupportModifyFrontcfg"`
	SupportNVR                   bool `json:"SupportNVR"`
	SupportNetLocalSearch        bool `json:"SupportNetLocalSearch"`
	SupportNetWorkMode           bool `json:"SupportNetWorkMode"`
	SupportOSDInfo               bool `json:"SupportOSDInfo"`
	SupportOnvifClient           bool `json:"SupportOnvifClient"`
	SupportPOS                   bool `json:"SupportPOS"`
	SupportPWDSafety             bool `json:"SupportPWDSafety"`
	SupportPlateDetect           bool `json:"SupportPlateDetect"`
	SupportPlayBackExactSeek     bool `json:"SupportPlayBackExactSeek"`
	SupportPlaybackLocate        bool `json:"SupportPlaybackLocate"`
	SupportPtzIdleState          bool `json:"SupportPtzIdleState"`
	SupportRPSVideo              bool `json:"SupportRPSVideo"`
	SupportRTSPClient            bool `json:"SupportRTSPClient"`
	SupportResumePtzState        bool `json:"SupportResumePtzState"`
	SupportSPVMNNasServer        bool `json:"SupportSPVMNNasServer"`
	SupportSafetyEmail           bool `json:"SupportSafetyEmail"`
	SupportSetDigIP              bool `json:"SupportSetDigIP"`
	SupportSetHardwareAbility    bool `json:"SupportSetHardwareAbility"`
	SupportSetPTZPresetAttribute bool `json:"SupportSetPTZPresetAttribute"`
	SupportShowConnectStatus     bool `json:"SupportShowConnectStatus"`
	SupportShowProductType       bool `json:"SupportShowProductType"`
	SupportSlowMotion            bool `json:"SupportSlowMotion"`
	SupportSmallChnTitleFont     bool `json:"SupportSmallChnTitleFont"`
	SupportSnapCfg               bool `json:"SupportSnapCfg"`
	SupportSnapSchedule          bool `json:"SupportSnapSchedule"`
	SupportSnapV2Stream          bool `json:"SupportSnapV2Stream"`
	SupportSplitControl          bool `json:"SupportSplitControl"`
	SupportStatusLed             bool `json:"SupportStatusLed"`
	SupportStorageFailReboot     bool `json:"SupportStorageFailReboot"`
	SupportStorageNAS            bool `json:"SupportStorageNAS"`
	SupportSwitchResolution      bool `json:"SupportSwitchResolution"`
	SupportTextPassword          bool `json:"SupportTextPassword"`
	SupportTimeZone              bool `json:"SupportTimeZone"`
	SupportUserProgram           bool `json:"SupportUserProgram"`
	SupportWIFINVR               bool `json:"SupportWIFINVR"`
	SupportWriteLog              bool `json:"SupportWriteLog"`
	Supportonviftitle            bool `json:"Supportonviftitle"`
	SuppportChangeOnvifPort      bool `json:"SuppportChangeOnvifPort"`
	TitleAndStateUpload          bool `json:"TitleAndStateUpload"`
	USBsupportRecord             bool `json:"USBsupportRecord"`
}

type PreviewFunction struct {
	GUISet bool `json:"GUISet"`
	Tour   bool `json:"Tour"`
}

type TipShow struct {
	NoBeepTipShow           bool `json:"NoBeepTipShow"`
	NoDiskManagerButtonShow bool `json:"NoDiskManagerButtonShow"`
	NoEmailTipShow          bool `json:"NoEmailTipShow"`
	NoFTPTipShow            bool `json:"NoFTPTipShow"`
}

type SystemFunction struct {
	AlarmFunction     AlarmFunction     `json:"AlarmFunction"`
	CommFunction      CommFunction      `json:"CommFunction"`
	EncodeFunction    EncodeFunction    `json:"EncodeFunction"`
	InputMethod       InputMethod       `json:"InputMethod"`
	MobileDVR         MobileDVR         `json:"MobileDVR"`
	NetServerFunction NetServerFunction `json:"NetServerFunction"`
	OtherFunction     OtherFunction     `json:"OtherFunction"`
	PreviewFunction   PreviewFunction   `json:"PreviewFunction"`
	TipShow           TipShow           `json:"TipShow"`
}
