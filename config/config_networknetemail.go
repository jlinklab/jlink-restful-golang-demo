package config

type NetWorkNetEmailResp struct {
	Name            string          `json:"Name"`
	NetWorkNetEmail NetWorkNetEmail `json:"NetWork.NetEmail"`
	Ret             int             `json:"Ret"`
	SessionID       string          `json:"SessionID"`
	RetMsg          string          `json:"RetMsg"`
}

type MailServer struct {
	Address   string `json:"Address"`
	Anonymity bool   `json:"Anonymity"`
	Name      string `json:"Name"`
	Password  string `json:"Password"`
	Port      int    `json:"Port"`
	UserName  string `json:"UserName"`
}

type NetWorkNetEmail struct {
	Enable     bool       `json:"Enable"`
	MailServer MailServer `json:"MailServer"`
	Port       int        `json:"Port"`
	Recievers  []string   `json:"Recievers"`
	Schedule   []string   `json:"Schedule"`
	SendAddr   string     `json:"SendAddr"`
	Title      string     `json:"Title"`
	UseSSL     bool       `json:"UseSSL"`
}
