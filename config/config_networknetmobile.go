package config

type NetWorkNetMobileResp struct {
	Name             string      `json:"Name"`
	Ret              int         `json:"Ret"`
	SessionID        string      `json:"SessionID"`
	NetWorkNetMobile interface{} `json:"NetWork.NetMobile"`
	RetMsg           string      `json:"RetMsg"`
}
