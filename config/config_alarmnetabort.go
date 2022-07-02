package config

type AlarmNetAbortResp struct {
	AlarmNetAbort AlarmNetAbort `json:"Alarm.NetAbort,omitempty"`
	Name          string        `json:"Name,omitempty"`
	Ret           int           `json:"Ret,omitempty"`
	SessionID     string        `json:"SessionID,omitempty"`
	RetMsg        string        `json:"RetMsg,omitempty"`
}

type AlarmNetAbort struct {
	Enable       bool         `json:"Enable,omitempty"`
	EventHandler EventHandler `json:"EventHandler,omitempty"`
}
