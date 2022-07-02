package config

type DetectHumanDetection0Resp struct {
	DetectHumanDetection0 DetectHumanDetection `json:"Detect.HumanDetection.[0]"`
	Name                  string               `json:"Name"`
	Ret                   int                  `json:"Ret"`
	SessionID             string               `json:"SessionID"`
	RetMsg                string               `json:"RetMsg"`
}
