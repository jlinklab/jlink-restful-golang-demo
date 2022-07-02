package config

type AVEncSmartH264Resp struct {
	AVEncSmartH264 []AVEncSmartH264 `json:"AVEnc.SmartH264,omitempty"`
	Name           string           `json:"Name,omitempty"`
	Ret            int              `json:"Ret,omitempty"`
	SessionID      string           `json:"SessionID,omitempty"`
	RetMsg         string           `json:"RetMsg,omitempty"`
}

type AVEncSmartH264 struct {
	SmartH264 bool `json:"SmartH264,omitempty"`
}
