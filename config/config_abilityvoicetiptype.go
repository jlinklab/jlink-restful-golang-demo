package config

type AbilityVoiceTipTypeResp struct {
	AbilityVoiceTipType []AbilityVoiceTipType `json:"Ability.VoiceTipType,omitempty"`
	Name                string                `json:"Name,omitempty"`
	Ret                 int                   `json:"Ret,omitempty"`
	SessionID           string                `json:"SessionID,omitempty"`
	RetMsg              string                `json:"RetMsg,omitempty"`
}

type AbilityVoiceTipType struct {
	VoiceTip []VoiceTip `json:"VoiceTip,omitempty"`
}

type VoiceTip struct {
	VoiceEnum int    `json:"VoiceEnum,omitempty"`
	VoiceText string `json:"VoiceText,omitempty"`
}
