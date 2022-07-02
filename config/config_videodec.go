package config

type VideoDecResp struct {
	Name      string      `json:"Name"`
	Ret       int         `json:"Ret"`
	SessionID string      `json:"SessionID"`
	VideoDec  interface{} `json:"VideoDec"`
	RetMsg    string      `json:"RetMsg"`
}
