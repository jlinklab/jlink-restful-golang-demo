package config

type UsersResp struct {
	Name      string      `json:"Name"`
	Ret       int         `json:"Ret"`
	SessionID string      `json:"SessionID"`
	Users     interface{} `json:"Users"`
	RetMsg    string      `json:"RetMsg"`
}
