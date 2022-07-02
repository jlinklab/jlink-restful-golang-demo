package config

type FVideoVolumeResp struct {
	Name         string         `json:"Name"`
	Ret          int            `json:"Ret"`
	SessionID    string         `json:"SessionID"`
	FVideoVolume []FVideoVolume `json:"fVideo.Volume"`
	RetMsg       string         `json:"RetMsg"`
}

type FVideoVolume struct {
	AudioMode   string `json:"AudioMode"`
	LeftVolume  int    `json:"LeftVolume"`
	RightVolume int    `json:"RightVolume"`
}
