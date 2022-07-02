package ability

type UartProtocolResp struct {
	Name         string   `json:"Name"`
	Ret          int      `json:"Ret"`
	SessionID    string   `json:"SessionID"`
	UartProtocol []string `json:"UartProtocol"`
	RetMsg       string   `json:"RetMsg"`
}
