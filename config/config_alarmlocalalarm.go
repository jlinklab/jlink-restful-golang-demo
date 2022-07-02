package config

type AlarmLocalAlarmResp struct {
	AlarmLocalAlarm []AlarmLocalAlarm `json:"Alarm.LocalAlarm,omitempty"`
	Name            string            `json:"Name,omitempty"`
	Ret             int               `json:"Ret,omitempty"`
	SessionID       string            `json:"SessionID,omitempty"`
	RetMsg          string            `json:"RetMsg,omitempty"`
}

type AlarmLocalAlarm struct {
	Enable       bool         `json:"Enable,omitempty"`
	EventHandler EventHandler `json:"EventHandler,omitempty"`
	SensorType   string       `json:"SensorType,omitempty"`
}
