package ability

type VideoAbilityResp struct {
	Name         string      `json:"Name"`
	Ret          int         `json:"Ret"`
	SessionID    string      `json:"SessionID"`
	VideoAbility interface{} `json:"VideoAbility"`
	RetMsg       string      `json:"RetMsg"`
}
