package config

type FVideoOSDInfoResp struct {
	Name          string        `json:"Name"`
	Ret           int           `json:"Ret"`
	SessionID     string        `json:"SessionID"`
	FVideoOSDInfo FVideoOSDInfo `json:"fVideo.OSDInfo"`
	RetMsg        string        `json:"RetMsg"`
}

type OSDInfoWidget struct {
	BackColor    string `json:"BackColor"`
	EncodeBlend  bool   `json:"EncodeBlend"`
	FrontColor   string `json:"FrontColor"`
	PreviewBlend bool   `json:"PreviewBlend"`
	RelativePos  []int  `json:"RelativePos"`
}

type OSDInfo struct {
	Info          []string      `json:"Info"`
	OSDInfoWidget OSDInfoWidget `json:"OSDInfoWidget"`
}

type FVideoOSDInfo struct {
	Align   int       `json:"Align"`
	OSDInfo []OSDInfo `json:"OSDInfo"`
	StrEnc  string    `json:"strEnc"`
}
