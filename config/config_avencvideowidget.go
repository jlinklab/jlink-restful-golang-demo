package config

type AVEncVideoWidgetResp struct {
	AVEncVideoWidget []AVEncVideoWidget `json:"AVEnc.VideoWidget,omitempty"`
	Name             string             `json:"Name,omitempty"`
	Ret              int                `json:"Ret,omitempty"`
	SessionID        string             `json:"SessionID,omitempty"`
	RetMsg           string             `json:"RetMsg,omitempty"`
}

type ChannelTitle struct {
	Name     string `json:"Name,omitempty"`
	SerialNo string `json:"SerialNo,omitempty"`
}

type ChannelTitleAttribute struct {
	BackColor    string `json:"BackColor,omitempty"`
	EncodeBlend  bool   `json:"EncodeBlend,omitempty"`
	FrontColor   string `json:"FrontColor,omitempty"`
	PreviewBlend bool   `json:"PreviewBlend,omitempty"`
	RelativePos  []int  `json:"RelativePos,omitempty"`
}

type Covers struct {
	BackColor    string `json:"BackColor,omitempty"`
	EncodeBlend  bool   `json:"EncodeBlend,omitempty"`
	FrontColor   string `json:"FrontColor,omitempty"`
	PreviewBlend bool   `json:"PreviewBlend,omitempty"`
	RelativePos  []int  `json:"RelativePos,omitempty"`
}

type TimeTitleAttribute struct {
	BackColor    string `json:"BackColor,omitempty"`
	EncodeBlend  bool   `json:"EncodeBlend,omitempty"`
	FrontColor   string `json:"FrontColor,omitempty"`
	PreviewBlend bool   `json:"PreviewBlend,omitempty"`
	RelativePos  []int  `json:"RelativePos,omitempty"`
}

type AVEncVideoWidget struct {
	ChannelTitle          ChannelTitle          `json:"ChannelTitle,omitempty"`
	ChannelTitleAttribute ChannelTitleAttribute `json:"ChannelTitleAttribute,omitempty"`
	Covers                []Covers              `json:"Covers,omitempty"`
	CoversNum             int                   `json:"CoversNum,omitempty"`
	TimeTitleAttribute    TimeTitleAttribute    `json:"TimeTitleAttribute,omitempty"`
}
