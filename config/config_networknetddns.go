package config

type NetWorkNetDDNSResp struct {
	Name           string           `json:"Name"`
	NetWorkNetDDNS []NetWorkNetDDNS `json:"NetWork.NetDDNS"`
	Ret            int              `json:"Ret"`
	SessionID      string           `json:"SessionID"`
	RetMsg         string           `json:"RetMsg"`
}

type NetWorkNetDDNS struct {
	DDNSKey  string `json:"DDNSKey"`
	Enable   bool   `json:"Enable"`
	HostName string `json:"HostName"`
	Server   Server `json:"Server"`
}
