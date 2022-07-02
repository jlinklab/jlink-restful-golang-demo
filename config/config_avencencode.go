package config

type AVEncEncodeResp struct {
	AVEncEncode []AVEncEncode `json:"AVEnc.Encode,omitempty"`
	Name        string        `json:"Name,omitempty"`
	Ret         int           `json:"Ret,omitempty"`
	SessionID   string        `json:"SessionID,omitempty"`
	RetMsg      string        `json:"RetMsg,omitempty"`
}

type Audio struct {
	BitRate    int `json:"BitRate,omitempty"`
	MaxVolume  int `json:"MaxVolume,omitempty"`
	SampleRate int `json:"SampleRate,omitempty"`
}

type Video struct {
	BitRate        int    `json:"BitRate,omitempty"`
	BitRateControl string `json:"BitRateControl,omitempty"`
	Compression    string `json:"Compression,omitempty"`
	FPS            int    `json:"FPS,omitempty"`
	GOP            int    `json:"GOP,omitempty"`
	Quality        int    `json:"Quality,omitempty"`
	Resolution     string `json:"Resolution,omitempty"`
}

type ExtraFormat struct {
	Audio       Audio `json:"Audio,omitempty"`
	AudioEnable bool  `json:"AudioEnable,omitempty"`
	Video       Video `json:"Video,omitempty"`
	VideoEnable bool  `json:"VideoEnable,omitempty"`
}

type MainFormat struct {
	Audio       Audio `json:"Audio,omitempty"`
	AudioEnable bool  `json:"AudioEnable,omitempty"`
	Video       Video `json:"Video,omitempty"`
	VideoEnable bool  `json:"VideoEnable,omitempty"`
}

type SnapFormat struct {
	Audio       Audio `json:"Audio,omitempty"`
	AudioEnable bool  `json:"AudioEnable,omitempty"`
	Video       Video `json:"Video,omitempty"`
	VideoEnable bool  `json:"VideoEnable,omitempty"`
}

type AVEncEncode struct {
	ExtraFormat []ExtraFormat `json:"ExtraFormat,omitempty"`
	MainFormat  []MainFormat  `json:"MainFormat,omitempty"`
	SnapFormat  []SnapFormat  `json:"SnapFormat,omitempty"`
}
