package config

type GeneralGeneralResp struct {
	GeneralGeneral GeneralGeneral `json:"General.General"`
	Name           string         `json:"Name"`
	Ret            int            `json:"Ret"`
	SessionID      string         `json:"SessionID"`
	RetMsg         string         `json:"RetMsg"`
}

type GeneralGeneral struct {
	AutoLogout         int    `json:"AutoLogout"`
	FontSize           int    `json:"FontSize"`
	IranCalendarEnable int    `json:"IranCalendarEnable"`
	LocalNo            int    `json:"LocalNo"`
	MachineName        string `json:"MachineName"`
	OverWrite          string `json:"OverWrite"`
	ScreenAutoShutdown int    `json:"ScreenAutoShutdown"`
	ScreenSaveTime     int    `json:"ScreenSaveTime"`
	VideoOutPut        string `json:"VideoOutPut"`
}
