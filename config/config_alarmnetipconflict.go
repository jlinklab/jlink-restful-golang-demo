package config

type AlarmNetIPConflictResp struct {
	AlarmNetIPConflict AlarmNetIPConflict `json:"Alarm.NetIPConflict,omitempty"`
	Name               string             `json:"Name,omitempty"`
	Ret                int                `json:"Ret,omitempty"`
	SessionID          string             `json:"SessionID,omitempty"`
	RetMsg             string             `json:"RetMsg,omitempty"`
}

type AlarmNetIPConflict struct {
	Enable       bool         `json:"Enable"`
	EventHandler EventHandler `json:"EventHandler,omitempty"`
}
