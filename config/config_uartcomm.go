package config

type UartCommResp struct {
	Name      string     `json:"Name"`
	Ret       int        `json:"Ret"`
	SessionID string     `json:"SessionID"`
	UartComm  []UartComm `json:"Uart.Comm"`
	RetMsg    string     `json:"RetMsg"`
}

type UartComm struct {
	Attribute    []interface{} `json:"Attribute"`
	PortNo       int           `json:"PortNo"`
	ProtocolName string        `json:"ProtocolName"`
}
