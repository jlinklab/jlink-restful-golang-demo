package config

type VideoChannelResp struct {
	Name         string         `json:"Name"`
	Ret          int            `json:"Ret"`
	SessionID    string         `json:"SessionID"`
	VideoChannel []VideoChannel `json:"VideoChannel"`
	RetMsg       string         `json:"RetMsg"`
}

type VideoChannel struct {
	Channel int    `json:"Channel"`
	Mode    string `json:"Mode"`
}
