package config

type NetWorkRemoteDeviceResp struct {
	Name                string      `json:"Name"`
	Ret                 int         `json:"Ret"`
	SessionID           string      `json:"SessionID"`
	NetWorkRemoteDevice interface{} `json:"NetWork.RemoteDevice"`
	RetMsg              string      `json:"RetMsg"`
}
