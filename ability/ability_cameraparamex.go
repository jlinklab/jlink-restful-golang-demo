package ability

type CameraParamExResp struct {
	Name          string      `json:"Name"`
	Ret           int         `json:"Ret"`
	SessionID     string      `json:"SessionID"`
	CameraParamEx interface{} `json:"Camera.ParamEx"`
	RetMsg        string      `json:"RetMsg"`
}
