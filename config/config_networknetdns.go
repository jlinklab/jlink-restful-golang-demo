package config

type NetWorkNetDNSResp struct {
	Name          string        `json:"Name"`
	NetWorkNetDNS NetWorkNetDNS `json:"NetWork.NetDNS"`
	Ret           int           `json:"Ret"`
	SessionID     string        `json:"SessionID"`
	RetMsg        string        `json:"RetMsg"`
}

type NetWorkNetDNS struct {
	Address      string `json:"Address"`
	SpareAddress string `json:"SpareAddress"`
}
