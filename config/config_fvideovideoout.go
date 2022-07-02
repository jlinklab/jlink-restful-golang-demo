package config

type FVideoVideoOutResp struct {
	Name           string         `json:"Name"`
	Ret            int            `json:"Ret"`
	SessionID      string         `json:"SessionID"`
	FVideoVideoOut FVideoVideoOut `json:"fVideo.VideoOut"`
	RetMsg         string         `json:"RetMsg"`
}

type Mode struct {
	Height int `json:"Height"`
	Width  int `json:"Width"`
}

type FVideoVideoOut struct {
	Mode Mode `json:"Mode"`
}
