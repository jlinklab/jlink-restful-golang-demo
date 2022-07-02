package config

type FVideoVideoOutPriorityResp struct {
	Name                   string                 `json:"Name"`
	Ret                    int                    `json:"Ret"`
	SessionID              string                 `json:"SessionID"`
	FVideoVideoOutPriority FVideoVideoOutPriority `json:"fVideo.VideoOutPriority"`
	RetMsg                 string                 `json:"RetMsg"`
}

type FVideoVideoOutPriority struct {
	Mode int `json:"Mode"`
}
