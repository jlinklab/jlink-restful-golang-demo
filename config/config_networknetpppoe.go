package config

type NetWorkNetPPPOEResp struct {
	Name            string      `json:"Name"`
	Ret             int         `json:"Ret"`
	SessionID       string      `json:"SessionID"`
	NetWorkNetPPPOE interface{} `json:"NetWork.NetPPPOE"`
	RetMsg          string      `json:"RetMsg"`
}
