package config

type NetWorkNetARSPResp struct {
	Name           string           `json:"Name"`
	NetWorkNetARSP []NetWorkNetARSP `json:"NetWork.NetARSP"`
	Ret            int              `json:"Ret"`
	SessionID      string           `json:"SessionID"`
	RetMsg         string           `json:"RetMsg"`
}

type NetWorkNetARSP struct {
	ARSPKey  string `json:"ARSPKey"`
	Enable   bool   `json:"Enable"`
	HTTPPort int    `json:"HttpPort"`
	Interval int    `json:"Interval"`
	Server   Server `json:"Server"`
	URL      string `json:"URL"`
}
