package config

type CameraParamResp struct {
	CameraParam []CameraParam `json:"Camera.Param"`
	Name        string        `json:"Name"`
	Ret         int           `json:"Ret"`
	SessionID   string        `json:"SessionID"`
	RetMsg      string        `json:"RetMsg"`
}

type CameraParam struct {
	AeSensitivity int           `json:"AeSensitivity"`
	ApertureMode  string        `json:"ApertureMode"`
	BLCMode       string        `json:"BLCMode"`
	DayNightColor string        `json:"DayNightColor"`
	DayNfLevel    int           `json:"Day_nfLevel"`
	DncThr        int           `json:"DncThr"`
	ElecLevel     int           `json:"ElecLevel"`
	EsShutter     string        `json:"EsShutter"`
	ExposureParam ExposureParam `json:"ExposureParam"`
	GainParam     GainParam     `json:"GainParam"`
	IRCUTMode     int           `json:"IRCUTMode"`
	IrcutSwap     int           `json:"IrcutSwap"`
	NightNfLevel  int           `json:"Night_nfLevel"`
	PictureFlip   string        `json:"PictureFlip"`
	PictureMirror string        `json:"PictureMirror"`
	RejectFlicker int           `json:"RejectFlicker"`
	WhiteBalance  string        `json:"WhiteBalance"`
}
