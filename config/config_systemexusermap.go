package config

type SystemExUserMapResp struct {
	Name            string          `json:"Name"`
	Ret             int             `json:"Ret"`
	SessionID       string          `json:"SessionID"`
	SystemExUserMap SystemExUserMap `json:"System.ExUserMap"`
	RetMsg          string          `json:"RetMsg"`
}

type User struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

type SystemExUserMap struct {
	User    []User `json:"User"`
	UserNum int    `json:"UserNum"`
}
