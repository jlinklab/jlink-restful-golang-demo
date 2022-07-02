package config

type NetWorkUpnpResp struct {
	Name        string      `json:"Name"`
	NetWorkUpnp NetWorkUpnp `json:"NetWork.Upnp"`
	Ret         int         `json:"Ret"`
	SessionID   string      `json:"SessionID"`
	RetMsg      string      `json:"RetMsg"`
}

type NetWorkUpnp struct {
	Enable     bool `json:"Enable"`
	HTTPPort   int  `json:"HTTPPort"`
	MediaPort  int  `json:"MediaPort"`
	MobilePort int  `json:"MobilePort"`
	State      bool `json:"State"`
}
