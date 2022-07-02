package config

type NetWorkNet3GResp struct {
	Name         string       `json:"Name"`
	NetWorkNet3G NetWorkNet3G `json:"NetWork.Net3G"`
	Ret          int          `json:"Ret"`
	SessionID    string       `json:"SessionID"`
	RetMsg       string       `json:"RetMsg"`
}

type NetWorkNet3G struct {
	APN            string `json:"APN"`
	DialNum        string `json:"DialNum"`
	Enable         bool   `json:"Enable"`
	NetType        string `json:"NetType"`
	OperatorsValue string `json:"OperatorsValue"`
	Password       string `json:"Password"`
	UserName       string `json:"UserName"`
}
