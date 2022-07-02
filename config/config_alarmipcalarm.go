package config

type AlarmIPCAlarmResp struct {
	Name          string      `json:"Name,omitempty"`
	Ret           int         `json:"Ret,omitempty"`
	SessionID     string      `json:"SessionID,omitempty"`
	AlarmIPCAlarm interface{} `json:"Alarm.IPCAlarm,omitempty"`
	RetMsg        string      `json:"RetMsg,omitempty"`
}
