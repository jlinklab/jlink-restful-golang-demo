package ability

type EncodeCapabilityResp struct {
	EncodeCapability EncodeCapability `json:"EncodeCapability"`
	Name             string           `json:"Name"`
	Ret              int              `json:"Ret"`
	SessionID        string           `json:"SessionID"`
	RetMsg           string           `json:"RetMsg"`
}

type CombEncodeInfo struct {
	CompressionMask string `json:"CompressionMask"`
	Enable          bool   `json:"Enable"`
	HaveAudio       bool   `json:"HaveAudio"`
	ResolutionMask  string `json:"ResolutionMask"`
	StreamType      string `json:"StreamType"`
}

type EncodeInfo struct {
	CompressionMask string `json:"CompressionMask"`
	Enable          bool   `json:"Enable"`
	HaveAudio       bool   `json:"HaveAudio"`
	ResolutionMask  string `json:"ResolutionMask"`
	StreamType      string `json:"StreamType"`
}

type EncodeCapability struct {
	ChannelMaxSetSync        int              `json:"ChannelMaxSetSync"`
	CombEncodeInfo           []CombEncodeInfo `json:"CombEncodeInfo"`
	Compression              string           `json:"Compression"`
	EncodeInfo               []EncodeInfo     `json:"EncodeInfo"`
	ExImageSizePerChannel    []string         `json:"ExImageSizePerChannel"`
	ExImageSizePerChannelEx  [][]string       `json:"ExImageSizePerChannelEx"`
	FourthStreamImageSize    []string         `json:"FourthStreamImageSize"`
	ImageSizePerChannel      []string         `json:"ImageSizePerChannel"`
	MaxBitrate               int              `json:"MaxBitrate"`
	MaxEncodePower           int              `json:"MaxEncodePower"`
	MaxEncodePowerPerChannel []string         `json:"MaxEncodePowerPerChannel"`
	ThirdStreamImageSize     []string         `json:"ThirdStreamImageSize"`
}
