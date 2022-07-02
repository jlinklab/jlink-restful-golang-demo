package config

type NetWorkShortMsgResp struct {
	Name            string      `json:"Name"`
	Ret             int         `json:"Ret"`
	SessionID       string      `json:"SessionID"`
	NetWorkShortMsg interface{} `json:"NetWork.ShortMsg"`
	RetMsg          string      `json:"RetMsg"`
}
