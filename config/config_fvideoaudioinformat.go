package config

type FVideoAudioInFormatResp struct {
	Name                string                `json:"Name"`
	Ret                 int                   `json:"Ret"`
	SessionID           string                `json:"SessionID"`
	FVideoAudioInFormat []FVideoAudioInFormat `json:"fVideo.AudioInFormat"`
	RetMsg              string                `json:"RetMsg"`
}

type FVideoAudioInFormat struct {
	BitRate    int    `json:"BitRate"`
	EncodeType string `json:"EncodeType"`
	SampleBit  int    `json:"SampleBit"`
	SampleRate int    `json:"SampleRate"`
}
