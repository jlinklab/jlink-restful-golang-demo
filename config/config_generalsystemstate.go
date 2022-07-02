package config

type GeneralSystemStateResp struct {
	Name               string      `json:"Name"`
	Ret                int         `json:"Ret"`
	SessionID          string      `json:"SessionID"`
	GeneralSystemState interface{} `json:"General.SystemState"`
	RetMsg             string      `json:"RetMsg"`
}
