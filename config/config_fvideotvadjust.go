package config

type FVideoTVAdjustResp struct {
	Name           string         `json:"Name"`
	Ret            int            `json:"Ret"`
	SessionID      string         `json:"SessionID"`
	FVideoTVAdjust FVideoTVAdjust `json:"fVideo.TVAdjust"`
	RetMsg         string         `json:"RetMsg"`
}

type FVideoTVAdjust struct {
	AntiDither  int   `json:"AntiDither"`
	BlackMargin []int `json:"BlackMargin"`
	Brightness  int   `json:"Brightness"`
	Contrast    int   `json:"Contrast"`
	Margin      []int `json:"Margin"`
}
