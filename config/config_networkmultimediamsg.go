package config

type NetWorkMultimediaMsgResp struct {
	Name                 string      `json:"Name"`
	Ret                  int         `json:"Ret"`
	SessionID            string      `json:"SessionID"`
	NetWorkMultimediaMsg interface{} `json:"NetWork.MultimediaMsg"`
	RetMsg               string      `json:"RetMsg"`
}
