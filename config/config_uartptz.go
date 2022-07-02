package config

type UartPTZResp struct {
	Name      string    `json:"Name"`
	Ret       int       `json:"Ret"`
	SessionID string    `json:"SessionID"`
	UartPTZ   []UartPTZ `json:"Uart.PTZ"`
	RetMsg    string    `json:"RetMsg"`
}

type UartPTZ struct {
	Attribute       []interface{} `json:"Attribute"`
	DeviceNo        int           `json:"DeviceNo"`
	NumberInMatrixs int           `json:"NumberInMatrixs"`
	PortNo          int           `json:"PortNo"`
	ProtocolName    string        `json:"ProtocolName"`
}
