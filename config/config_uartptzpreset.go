package config

type UartPTZPresetResp struct {
	Name          string      `json:"Name"`
	Ret           int         `json:"Ret"`
	SessionID     string      `json:"SessionID"`
	UartPTZPreset interface{} `json:"Uart.PTZPreset"`
	RetMsg        string      `json:"RetMsg"`
}
