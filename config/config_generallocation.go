package config

type GeneralLocationResp struct {
	GeneralLocation GeneralLocation `json:"General.Location"`
	Name            string          `json:"Name"`
	Ret             int             `json:"Ret"`
	SessionID       string          `json:"SessionID"`
	RetMsg          string          `json:"RetMsg"`
}

type DSTEnd struct {
	Day    int `json:"Day"`
	Hour   int `json:"Hour"`
	Minute int `json:"Minute"`
	Month  int `json:"Month"`
	Week   int `json:"Week"`
	Year   int `json:"Year"`
}

type DSTStart struct {
	Day    int `json:"Day"`
	Hour   int `json:"Hour"`
	Minute int `json:"Minute"`
	Month  int `json:"Month"`
	Week   int `json:"Week"`
	Year   int `json:"Year"`
}

type GeneralLocation struct {
	DSTEnd        DSTEnd   `json:"DSTEnd"`
	DSTRule       string   `json:"DSTRule"`
	DSTStart      DSTStart `json:"DSTStart"`
	DateFormat    string   `json:"DateFormat"`
	DateSeparator string   `json:"DateSeparator"`
	IranCalendar  int      `json:"IranCalendar"`
	Language      string   `json:"Language"`
	TimeFormat    string   `json:"TimeFormat"`
	VideoFormat   string   `json:"VideoFormat"`
	WorkDay       int      `json:"WorkDay"`
}
