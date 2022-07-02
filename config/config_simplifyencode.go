package config

type SimplifyEncodeResp struct {
	Name           string           `json:"Name"`
	Ret            int              `json:"Ret"`
	SessionID      string           `json:"SessionID"`
	SimplifyEncode []SimplifyEncode `json:"Simplify.Encode"`
	RetMsg         string           `json:"RetMsg"`
}

type SimplifyEncode struct {
	ExtraFormat ExtraFormat `json:"ExtraFormat"`
	MainFormat  MainFormat  `json:"MainFormat"`
}
