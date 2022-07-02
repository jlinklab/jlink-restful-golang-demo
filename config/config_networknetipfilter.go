package config

type NetWorkNetIPFilterResp struct {
	Name               string             `json:"Name"`
	NetWorkNetIPFilter NetWorkNetIPFilter `json:"NetWork.NetIPFilter"`
	Ret                int                `json:"Ret"`
	SessionID          string             `json:"SessionID"`
	RetMsg             string             `json:"RetMsg"`
}

type NetWorkNetIPFilter struct {
	Banned  []string `json:"Banned"`
	Enable  bool     `json:"Enable"`
	Trusted []string `json:"Trusted"`
}
