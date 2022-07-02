package config

type GroupsResp struct {
	Name      string      `json:"Name"`
	Ret       int         `json:"Ret"`
	SessionID string      `json:"SessionID"`
	Groups    interface{} `json:"Groups"`
	RetMsg    string      `json:"RetMsg"`
}
