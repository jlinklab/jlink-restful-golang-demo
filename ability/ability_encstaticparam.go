package ability

type EncStaticParamResp struct {
	Name           string      `json:"Name"`
	Ret            int         `json:"Ret"`
	SessionID      string      `json:"SessionID"`
	EncStaticParam interface{} `json:"EncStaticParam"`
	RetMsg         string      `json:"RetMsg"`
}
