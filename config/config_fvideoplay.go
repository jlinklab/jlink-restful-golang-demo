package config

type FVideoPlayResp struct {
	Name       string     `json:"Name"`
	Ret        int        `json:"Ret"`
	SessionID  string     `json:"SessionID"`
	FVideoPlay FVideoPlay `json:"fVideo.Play"`
	RetMsg     string     `json:"RetMsg"`
}

type FVideoPlay struct {
	Channels   int    `json:"Channels"`
	Continue   bool   `json:"Continue"`
	PlayMode   int    `json:"PlayMode"`
	SearchType string `json:"SearchType"`
	Volume     []int  `json:"Volume"`
}
