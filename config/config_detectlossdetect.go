package config

type DetectLossDetectResp struct {
	DetectLossDetect []DetectLossDetect `json:"Detect.LossDetect"`
	Name             string             `json:"Name"`
	Ret              int                `json:"Ret"`
	SessionID        string             `json:"SessionID"`
	RetMsg           string             `json:"RetMsg"`
}

type DetectLossDetect struct {
	Enable       bool         `json:"Enable"`
	EventHandler EventHandler `json:"EventHandler"`
}
