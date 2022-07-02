package ability

type MultiChannelResp struct {
	Name         string      `json:"Name"`
	Ret          int         `json:"Ret"`
	SessionID    string      `json:"SessionID"`
	MultiChannel interface{} `json:"MultiChannel"`
	RetMsg       string      `json:"RetMsg"`
}
