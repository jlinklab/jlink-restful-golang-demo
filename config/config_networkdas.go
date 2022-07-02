package config

type NetWorkDASResp struct {
	Name       string     `json:"Name,omitempty"`
	NetWorkDAS NetWorkDAS `json:"NetWork.DAS,omitempty"`
	Ret        int        `json:"Ret,omitempty"`
	SessionID  string     `json:"SessionID,omitempty"`
	RetMsg     string     `json:"RetMsg,omitempty"`
}

type NetWorkDAS struct {
	DeviceID   string `json:"DeviceID,omitempty"`
	Enable     bool   `json:"Enable"`
	Password   string `json:"Password,omitempty"`
	Port       int    `json:"Port,omitempty"`
	ServerAddr string `json:"ServerAddr,omitempty"`
	UserName   string `json:"UserName,omitempty"`
}
