package ability

type AbilityGUIThemeListResp struct {
	Name                string      `json:"Name"`
	Ret                 int         `json:"Ret"`
	SessionID           string      `json:"SessionID"`
	AbilityGUIThemeList interface{} `json:"Ability.GUIThemeList"`
	RetMsg              string      `json:"RetMsg"`
}
