package opedev

const (
	OPMachine        = "OPMachine"
	OPDefaultConfig  = "OPDefaultConfig"
	StorageInfo      = "StorageInfo"
	OPStorageManager = "OPStorageManager"
	OPFileQuery      = "OPFileQuery"
	OPTimeQuery      = "OPTimeQuery"
	OPSystemConfig   = "OPSystemConfig"
	OPTimeSetting    = "OPTimeSetting"
	OPPTZControl     = "OPPTZControl"
	OPLogQuery       = "OPLogQuery"
	// OPVersionList                 = "OPVersionList"
	// OPReqVersion                  = "OPReqVersion"
	// OPCloudUpgradeStart           = "OPCloudUpgradeStart"
	// OPCloudUpgradeStop            = "OPCloudUpgradeStop"
	// OPCloudFileGetDownloadState   = "OPCloudFileGetDownloadState"
	// OPCloudUpgradeGetBurnProgress = "OPCloudUpgradeGetBurnProgress"
	dOpdevUrl = "/v2/rtc/device/opdev/"
)

type XYResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data Data   `json:"data"`
}

type Data struct {
	Name string `json:"name"`
	Ret  int    `json:"ret"`
}
