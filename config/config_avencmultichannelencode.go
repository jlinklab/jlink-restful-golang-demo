package config

type AVEncMultiChannelEncodeResp struct {
	AVEncMultiChannelEncode interface{} `json:"AVEnc.MultiChannelEncode,omitempty"`
	Name                    string      `json:"Name,omitempty"`
	Ret                     int         `json:"Ret,omitempty"`
	SessionID               string      `json:"SessionID,omitempty"`
	RetMsg                  string      `json:"RetMsg,omitempty"`
}
