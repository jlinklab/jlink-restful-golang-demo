package ability

type ComProtocolResp struct {
	ComProtocol []string `json:"ComProtocol"`
	Name        string   `json:"Name"`
	Ret         int      `json:"Ret"`
	SessionID   string   `json:"SessionID"`
	RetMsg      string   `json:"RetMsg"`
}
