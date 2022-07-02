package config

type UartRS485Resp struct {
	Name      string      `json:"Name"`
	Ret       int         `json:"Ret"`
	SessionID string      `json:"SessionID"`
	UartRS485 []UartRS485 `json:"Uart.RS485"`
	RetMsg    string      `json:"RetMsg"`
}

type UartRS485 struct {
	Attribute       []interface{} `json:"Attribute"`
	DeviceNo        int           `json:"DeviceNo"`
	NumberInMatrixs int           `json:"NumberInMatrixs"`
	PortNo          int           `json:"PortNo"`
	ProtocolName    string        `json:"ProtocolName"`
}
