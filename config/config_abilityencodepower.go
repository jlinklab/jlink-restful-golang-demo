package config

type AbilityEncodePowerResp struct {
	AbilityEncodePower AbilityEncodePower `json:"Ability.EncodePower,omitempty"`
	Name               string             `json:"Name,omitempty"`
	Ret                int                `json:"Ret,omitempty"`
	SessionID          string             `json:"SessionID,omitempty"`
	RetMsg             string             `json:"RetMsg,omitempty"`
}

type AbilityEncodePower struct {
	MaxResolution []string `json:"MaxResolution"`
}
