package config

type FVideoGUISetResp struct {
	Name         string       `json:"Name"`
	Ret          int          `json:"Ret"`
	SessionID    string       `json:"SessionID"`
	FVideoGUISet FVideoGUISet `json:"fVideo.GUISet"`
	RetMsg       string       `json:"RetMsg"`
}

type FVideoGUISet struct {
	AlarmStateEnable       bool `json:"AlarmStateEnable"`
	CarInfo                bool `json:"CarInfo"`
	ChanStateBitRateEnable bool `json:"ChanStateBitRateEnable"`
	ChanStateLckEnable     bool `json:"ChanStateLckEnable"`
	ChanStateMtdEnable     bool `json:"ChanStateMtdEnable"`
	ChanStateVlsEnable     bool `json:"ChanStateVlsEnable"`
	ChannelTitleEnable     bool `json:"ChannelTitleEnable"`
	Deflick                bool `json:"Deflick"`
	GPSInfo                bool `json:"GPSInfo"`
	RecordStateEnable      bool `json:"RecordStateEnable"`
	RemoteEnable           bool `json:"RemoteEnable"`
	TimeTitleEnable        bool `json:"TimeTitleEnable"`
	WindowAlpha            int  `json:"WindowAlpha"`
}
