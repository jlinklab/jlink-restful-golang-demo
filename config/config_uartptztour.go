package config

type UartPTZTourResp struct {
	Name        string      `json:"Name"`
	Ret         int         `json:"Ret"`
	SessionID   string      `json:"SessionID"`
	UartPTZTour interface{} `json:"Uart.PTZTour"`
	RetMsg      string      `json:"RetMsg"`
}
