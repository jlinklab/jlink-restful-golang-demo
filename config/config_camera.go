package config

type CameraResp struct {
	Camera    Camera `json:"Camera"`
	Name      string `json:"Name"`
	Ret       int    `json:"Ret"`
	SessionID string `json:"SessionID"`
	RetMsg    string `json:"RetMsg"`
}

type ClearFog struct {
	Enable int `json:"enable"`
	Level  int `json:"level"`
}

type DistortionCorrect struct {
	Lenstype int `json:"Lenstype"`
	Version  int `json:"Version"`
}

type FishLensParam struct {
	CenterOffsetX int    `json:"CenterOffsetX"`
	CenterOffsetY int    `json:"CenterOffsetY"`
	ImageHeight   int    `json:"ImageHeight"`
	ImageWidth    int    `json:"ImageWidth"`
	LensType      int    `json:"LensType"`
	PCMac         string `json:"PCMac"`
	Radius        int    `json:"Radius"`
	Version       int    `json:"Version"`
	ViewAngle     int    `json:"ViewAngle"`
	ViewMode      int    `json:"ViewMode"`
	Zoom          int    `json:"Zoom"`
}

type FishViCut struct {
	ImgHeight int `json:"ImgHeight"`
	ImgWidth  int `json:"ImgWidth"`
	Xoffset   int `json:"Xoffset"`
	Yoffset   int `json:"Yoffset"`
}

type ExposureParam struct {
	LeastTime string `json:"LeastTime"`
	Level     int    `json:"Level"`
	MostTime  string `json:"MostTime"`
}

type GainParam struct {
	AutoGain int `json:"AutoGain"`
	Gain     int `json:"Gain"`
}

type Param struct {
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

type BroadTrends struct {
	AutoGain int `json:"AutoGain"`
	Gain     int `json:"Gain"`
}

type ParamEx struct {
	BroadTrends        BroadTrends `json:"BroadTrends"`
	CorridorMode       int         `json:"CorridorMode"`
	ExposureTime       string      `json:"ExposureTime"`
	LightRestrainLevel int         `json:"LightRestrainLevel"`
	LowLuxMode         int         `json:"LowLuxMode"`
	Style              string      `json:"Style"`
}

type WorkPeriod struct {
	EHour   int `json:"EHour"`
	EMinute int `json:"EMinute"`
	Enable  int `json:"Enable"`
	SHour   int `json:"SHour"`
	SMinute int `json:"SMinute"`
}

type WhiteLight struct {
	WorkMode   string     `json:"WorkMode"`
	WorkPeriod WorkPeriod `json:"WorkPeriod"`
}

type Camera struct {
	ClearFog          []ClearFog        `json:"ClearFog"`
	DistortionCorrect DistortionCorrect `json:"DistortionCorrect"`
	FishLensParam     []FishLensParam   `json:"FishLensParam"`
	FishViCut         []FishViCut       `json:"FishViCut"`
	Param             []Param           `json:"Param"`
	ParamEx           []ParamEx         `json:"ParamEx"`
	WhiteLight        WhiteLight        `json:"WhiteLight"`
}
