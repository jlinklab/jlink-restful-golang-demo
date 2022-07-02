package config

type MediaDecodeParamResp struct {
	Name             string      `json:"Name"`
	Ret              int         `json:"Ret"`
	SessionID        string      `json:"SessionID"`
	MediaDecodeParam interface{} `json:"Media.DecodeParam"`
	RetMsg           string      `json:"RetMsg"`
}
