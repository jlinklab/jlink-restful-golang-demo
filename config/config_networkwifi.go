package config

type NetWorkWifiResp struct {
	Name        string      `json:"Name"`
	NetWorkWifi NetWorkWifi `json:"NetWork.Wifi"`
	Ret         int         `json:"Ret"`
	SessionID   string      `json:"SessionID"`
	RetMsg      string      `json:"RetMsg"`
}

type NetWorkWifi struct {
	Auth       string `json:"Auth"`
	Channel    int    `json:"Channel"`
	Enable     bool   `json:"Enable"`
	EncrypType string `json:"EncrypType"`
	GateWay    string `json:"GateWay"`
	HostIP     string `json:"HostIP"`
	KeyType    int    `json:"KeyType"`
	Keys       string `json:"Keys"`
	NetType    string `json:"NetType"`
	SSID       string `json:"SSID"`
	Submask    string `json:"Submask"`
}
