package config

type NetWorkNetCommonResp struct {
	Name             string           `json:"Name"`
	NetWorkNetCommon NetWorkNetCommon `json:"NetWork.NetCommon"`
	Ret              int              `json:"Ret"`
	SessionID        string           `json:"SessionID"`
	RetMsg           string           `json:"RetMsg"`
}

type NetWorkNetCommon struct {
	GateWay       string `json:"GateWay"`
	HostIP        string `json:"HostIP"`
	HostName      string `json:"HostName"`
	HTTPPort      int    `json:"HttpPort"`
	MAC           string `json:"MAC"`
	MaxBps        int    `json:"MaxBps"`
	MonMode       string `json:"MonMode"`
	SSLPort       int    `json:"SSLPort"`
	Submask       string `json:"Submask"`
	TCPMaxConn    int    `json:"TCPMaxConn"`
	TCPPort       int    `json:"TCPPort"`
	TransferPlan  string `json:"TransferPlan"`
	UDPPort       int    `json:"UDPPort"`
	UseHSDownLoad bool   `json:"UseHSDownLoad"`
}
