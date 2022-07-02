package config

type OPUTCTimeSettingResp struct {
	Ret              int         `json:"Ret"`
	Name             string      `json:"Name"`
	SessionID        string      `json:"SessionID"`
	OPUTCTimeSetting interface{} `json:"OPUTCTimeSetting"`
	RetMsg           string      `json:"RetMsg"`
}
