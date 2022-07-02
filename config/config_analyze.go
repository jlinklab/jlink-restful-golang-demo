package config

type AnalyzeResp struct {
	Name      string      `json:"Name,omitempty"`
	Ret       int         `json:"Ret,omitempty"`
	SessionID string      `json:"SessionID,omitempty"`
	Analyze   interface{} `json:"Analyze,omitempty"`
	RetMsg    string      `json:"RetMsg,omitempty"`
}
