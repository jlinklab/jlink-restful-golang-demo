package config

type NetWorkRemoteDeviceV3Resp struct {
	Name                  string      `json:"Name"`
	Ret                   int         `json:"Ret"`
	SessionID             string      `json:"SessionID"`
	NetWorkRemoteDeviceV3 interface{} `json:"NetWork.RemoteDeviceV3"`
	RetMsg                string      `json:"RetMsg"`
}
