package config

type NetWorkNetOrderResp struct {
	Name            string          `json:"Name"`
	NetWorkNetOrder NetWorkNetOrder `json:"NetWork.NetOrder"`
	Ret             int             `json:"Ret"`
	SessionID       string          `json:"SessionID"`
	RetMsg          string          `json:"RetMsg"`
}

type NetWorkNetOrder struct {
	Enable   bool       `json:"Enable"`
	NetCount int        `json:"NetCount"`
	NetOrder []NetOrder `json:"NetOrder"`
}
type NetOrder struct {
	NetOrder int `json:"NetOrder"`
	NetType  int `json:"NetType"`
}
