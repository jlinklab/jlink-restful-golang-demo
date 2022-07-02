package ability

type AbilityVGAResolutionResp struct {
	Name                 string      `json:"Name"`
	Ret                  int         `json:"Ret"`
	SessionID            string      `json:"SessionID"`
	AbilityVGAResolution interface{} `json:"Ability.VGAResolution"`
	RetMsg               string      `json:"RetMsg"`
}
