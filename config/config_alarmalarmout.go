package config

type AlarmAlarmOutResp struct {
	AlarmAlarmOut interface{} `json:"Alarm.AlarmOut,omitempty"`
	Name          string      `json:"Name,omitempty"`
	Ret           int         `json:"Ret,omitempty"`
	SessionID     string      `json:"SessionID,omitempty"`
	RetMsg        string      `json:"RetMsg,omitempty"`
}
