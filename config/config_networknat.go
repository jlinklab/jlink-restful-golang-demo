package config

type NetWorkNatResp struct {
	Name       string     `json:"Name"`
	NetWorkNat NetWorkNat `json:"NetWork.Nat"`
	Ret        int        `json:"Ret"`
	SessionID  string     `json:"SessionID"`
	RetMsg     string     `json:"RetMsg"`
}

type NetWorkNat struct {
	Addr       string `json:"Addr"`
	DNSServer1 string `json:"DnsServer1"`
	DNSServer2 string `json:"DnsServer2"`
	NatEnable  bool   `json:"NatEnable"`
	Port       int    `json:"Port"`
	XMeyeMTU   int    `json:"XMeyeMTU"`
}
