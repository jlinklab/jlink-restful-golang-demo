package ability

type DDNSServiceResp struct {
	DDNSService []string `json:"DDNSService"`
	Name        string   `json:"Name"`
	Ret         int      `json:"Ret"`
	SessionID   string   `json:"SessionID"`
	RetMsg      string   `json:"RetMsg"`
}
