package config

type AVEncVideoColorResp struct {
	AVEncVideoColor [][]AVEncVideoColor `json:"AVEnc.VideoColor,omitempty"`
	Name            string              `json:"Name,omitempty"`
	Ret             int                 `json:"Ret,omitempty"`
	SessionID       string              `json:"SessionID,omitempty"`
	RetMsg          string              `json:"RetMsg,omitempty"`
}

type VideoColorParam struct {
	Acutance     int `json:"Acutance,omitempty"`
	Brightness   int `json:"Brightness,omitempty"`
	Contrast     int `json:"Contrast,omitempty"`
	Gain         int `json:"Gain,omitempty"`
	Hue          int `json:"Hue,omitempty"`
	Saturation   int `json:"Saturation,omitempty"`
	Whitebalance int `json:"Whitebalance,omitempty"`
}

type AVEncVideoColor struct {
	Enable          bool            `json:"Enable"`
	TimeSection     string          `json:"TimeSection,omitempty"`
	VideoColorParam VideoColorParam `json:"VideoColorParam,omitempty"`
}
