package config

type NetWorkNetDHCPResp struct {
	Name           string           `json:"Name"`
	NetWorkNetDHCP []NetWorkNetDHCP `json:"NetWork.NetDHCP"`
	Ret            int              `json:"Ret"`
	SessionID      string           `json:"SessionID"`
	RetMsg         string           `json:"RetMsg"`
}

type NetWorkNetDHCP struct {
	Enable    bool   `json:"Enable"`
	Interface string `json:"Interface"`
}
