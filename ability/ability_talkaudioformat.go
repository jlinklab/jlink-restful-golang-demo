package ability

type TalkAudioFormatResp struct {
	Name            string            `json:"Name"`
	Ret             int               `json:"Ret"`
	SessionID       string            `json:"SessionID"`
	TalkAudioFormat []TalkAudioFormat `json:"TalkAudioFormat"`
	RetMsg          string            `json:"RetMsg"`
}

type TalkAudioFormat struct {
	BitRate    int    `json:"BitRate"`
	EncodeType string `json:"EncodeType"`
	SampleBit  int    `json:"SampleBit"`
	SampleRate int    `json:"SampleRate"`
}
