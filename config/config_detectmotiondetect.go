package config

type DetectMotionDetectResp struct {
	DetectMotionDetect []DetectMotionDetect `json:"Detect.MotionDetect"`
	Name               string               `json:"Name"`
	Ret                int                  `json:"Ret"`
	SessionID          string               `json:"SessionID"`
	RetMsg             string               `json:"RetMsg"`
}

type PirTimeSectionOne struct {
	Enable   bool `json:"Enable"`
	WeekMask int  `json:"WeekMask"`
}

type PirTimeSectionTwo struct {
	Enable   bool `json:"Enable"`
	WeekMask int  `json:"WeekMask"`
}

type PirTimeSection struct {
	PirTimeSectionOne PirTimeSectionOne `json:"PirTimeSectionOne"`
	PirTimeSectionTwo PirTimeSectionTwo `json:"PirTimeSectionTwo"`
}

type DetectMotionDetect struct {
	AlarmType      int            `json:"AlarmType"`
	Enable         bool           `json:"Enable"`
	EventHandler   EventHandler   `json:"EventHandler"`
	Level          int            `json:"Level"`
	PIRCheckTime   int            `json:"PIRCheckTime"`
	PirSensitive   int            `json:"PirSensitive"`
	PirTimeSection PirTimeSection `json:"PirTimeSection"`
	Region         []string       `json:"Region"`
}
