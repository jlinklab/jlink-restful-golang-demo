package ability

type MaxPreRecordResp struct {
	Name         string      `json:"Name"`
	Ret          int         `json:"Ret"`
	SessionID    string      `json:"SessionID"`
	MaxPreRecord interface{} `json:"MaxPreRecord"`
	RetMsg       string      `json:"RetMsg"`
}
