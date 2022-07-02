package config

type NetWorkNetNTPResp struct {
	Name          string        `json:"Name"`
	NetWorkNetNTP NetWorkNetNTP `json:"NetWork.NetNTP"`
	Ret           int           `json:"Ret"`
	SessionID     string        `json:"SessionID"`
	RetMsg        string        `json:"RetMsg"`
}

type Server struct {
	Address   string `json:"Address"`
	Anonymity bool   `json:"Anonymity"`
	Name      string `json:"Name"`
	Password  string `json:"Password"`
	Port      int    `json:"Port"`
	UserName  string `json:"UserName"`
}

type NetWorkNetNTP struct {
	Enable       bool   `json:"Enable"`
	Server       Server `json:"Server"`
	TimeZone     int    `json:"TimeZone"`
	UpdatePeriod int    `json:"UpdatePeriod"`
}
