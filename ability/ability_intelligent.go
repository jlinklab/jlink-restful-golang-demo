package ability

type IntelligentResp struct {
	Intelligent Intelligent `json:"Intelligent"`
	Name        string      `json:"Name"`
	Ret         int         `json:"Ret"`
	SessionID   string      `json:"SessionID"`
	RetMsg      string      `json:"RetMsg"`
}

type Intelligent struct {
	AlgorithmAVD string      `json:"AlgorithmAVD"`
	AlgorithmCPC string      `json:"AlgorithmCPC"`
	AlgorithmPEA string      `json:"AlgorithmPEA"`
	IntelAVD     string      `json:"IntelAVD"`
	IntelCPC     string      `json:"IntelCPC"`
	IntelPEA     string      `json:"IntelPEA"`
	LimitOSC     LimitOSCDTO `json:"LimitOSC"`
	LimitPEA     LimitPEADTO `json:"LimitPEA"`
}

type LimitOSCDTO struct {
	AreaMaxLineNum int `json:"AreaMaxLineNum"`
	AreaNum        int `json:"AreaNum"`
	Rectangle      int `json:"Rectangle"`
}

type LimitPEADTO struct {
	AreaMaxLineNum int `json:"AreaMaxLineNum"`
	AreaNum        int `json:"AreaNum"`
	LineNum        int `json:"LineNum"`
	Rectangle      int `json:"Rectangle"`
}
