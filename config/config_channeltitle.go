package config

type ChannelTitleResp struct {
	ChannelTitle []string `json:"ChannelTitle"`
	Name         string   `json:"Name"`
	Ret          int      `json:"Ret"`
	SessionID    string   `json:"SessionID"`
	RetMsg       string   `json:"RetMsg"`
}
