package ability

type AbilityDigitalRealResp struct {
	Name               string      `json:"Name"`
	Ret                int         `json:"Ret"`
	SessionID          string      `json:"SessionID"`
	AbilityDigitalReal interface{} `json:"Ability.DigitalReal"`
	RetMsg             string      `json:"RetMsg"`
}
