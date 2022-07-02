package config

type CameraParamExResp struct {
	CameraParamEx []CameraParamEx `json:"Camera.ParamEx"`
	Name          string          `json:"Name"`
	Ret           int             `json:"Ret"`
	SessionID     string          `json:"SessionID"`
	RetMsg        string          `json:"RetMsg"`
}

type CameraParamEx struct {
	BroadTrends        BroadTrends `json:"BroadTrends"`
	CorridorMode       int         `json:"CorridorMode"`
	ExposureTime       string      `json:"ExposureTime"`
	LightRestrainLevel int         `json:"LightRestrainLevel"`
	LowLuxMode         int         `json:"LowLuxMode"`
	Style              string      `json:"Style"`
}
