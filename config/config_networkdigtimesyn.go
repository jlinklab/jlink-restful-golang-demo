package config

type NetWorkDigTimeSynResp struct {
	Name              string      `json:"Name"`
	Ret               int         `json:"Ret"`
	SessionID         string      `json:"SessionID"`
	NetWorkDigTimeSyn interface{} `json:"NetWork.DigTimeSyn"`
	RetMsg            string      `json:"RetMsg"`
}
