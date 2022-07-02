package config

type NetWorkNetFTPResp struct {
	Name          string        `json:"Name"`
	NetWorkNetFTP NetWorkNetFTP `json:"NetWork.NetFTP"`
	Ret           int           `json:"Ret"`
	SessionID     string        `json:"SessionID"`
	RetMsg        string        `json:"RetMsg"`
}

type NetWorkNetFTP struct {
	Directory  string `json:"Directory"`
	Enable     bool   `json:"Enable"`
	MaxFileLen int    `json:"MaxFileLen"`
	Server     Server `json:"Server"`
}
