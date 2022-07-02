package config

type GeneralAutoMaintainResp struct {
	GeneralAutoMaintain GeneralAutoMaintain `json:"General.AutoMaintain"`
	Name                string              `json:"Name"`
	Ret                 int                 `json:"Ret"`
	SessionID           string              `json:"SessionID"`
	RetMsg              string              `json:"RetMsg"`
}

type GeneralAutoMaintain struct {
	AutoDeleteFilesDays int    `json:"AutoDeleteFilesDays"`
	AutoRebootDay       string `json:"AutoRebootDay"`
	AutoRebootHour      int    `json:"AutoRebootHour"`
}
