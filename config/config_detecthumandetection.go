package config

type DetectHumanDetectionResp struct {
	DetectHumanDetection []DetectHumanDetection `json:"Detect.HumanDetection"`
	Name                 string                 `json:"Name"`
	Ret                  int                    `json:"Ret"`
	SessionID            string                 `json:"SessionID"`
	RetMsg               string                 `json:"RetMsg"`
}

type Pts struct {
	StartX int `json:"StartX"`
	StartY int `json:"StartY"`
	StopX  int `json:"StopX"`
	StopY  int `json:"StopY"`
}

type RuleLine struct {
	AlarmDirect int `json:"AlarmDirect"`
	Pts         Pts `json:"Pts"`
}

type RuleRegion struct {
	AlarmDirect int   `json:"AlarmDirect"`
	Pts         []Pts `json:"Pts"`
	PtsNum      int   `json:"PtsNum"`
}

type PedRule struct {
	Enable     bool       `json:"Enable"`
	RuleLine   RuleLine   `json:"RuleLine"`
	RuleRegion RuleRegion `json:"RuleRegion"`
	RuleType   int        `json:"RuleType"`
}

type DetectHumanDetection struct {
	Enable       bool      `json:"Enable"`
	ObjectType   int       `json:"ObjectType"`
	PedFdrAlg    int       `json:"PedFdrAlg"`
	PedRule      []PedRule `json:"PedRule"`
	PushInterval int       `json:"PushInterval"`
	Sensitivity  int       `json:"Sensitivity"`
	ShowRule     int       `json:"ShowRule"`
	ShowTrack    int       `json:"ShowTrack"`
}
