package ability

type AbilitySupportVstdResp struct {
	Name               string      `json:"Name"`
	Ret                int         `json:"Ret"`
	SessionID          string      `json:"SessionID"`
	AbilitySupportVstd interface{} `json:"Ability.SupportVstd"`
	RetMsg             string      `json:"RetMsg"`
}
