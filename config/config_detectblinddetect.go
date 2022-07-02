package config

type DetectBlindDetectResp struct {
	DetectBlindDetect []DetectBlindDetect `json:"Detect.BlindDetect"`
	Name              string              `json:"Name"`
	Ret               int                 `json:"Ret"`
	SessionID         string              `json:"SessionID"`
	RetMsg            string              `json:"RetMsg"`
}

type EventHandler struct {
	AlarmInfo           string          `json:"AlarmInfo"`
	AlarmOutEnable      bool            `json:"AlarmOutEnable"`
	AlarmOutLatch       int             `json:"AlarmOutLatch"`
	AlarmOutMask        string          `json:"AlarmOutMask"`
	BeepEnable          bool            `json:"BeepEnable"`
	EventLatch          int             `json:"EventLatch"`
	FTPEnable           bool            `json:"FTPEnable"`
	LogEnable           bool            `json:"LogEnable"`
	MailEnable          bool            `json:"MailEnable"`
	MatrixEnable        bool            `json:"MatrixEnable"`
	MatrixMask          string          `json:"MatrixMask"`
	MessageEnable       bool            `json:"MessageEnable"`
	MsgtoNetEnable      bool            `json:"MsgtoNetEnable"`
	MultimediaMsgEnable bool            `json:"MultimediaMsgEnable"`
	PtzEnable           bool            `json:"PtzEnable"`
	PtzLink             [][]interface{} `json:"PtzLink"`
	RecordEnable        bool            `json:"RecordEnable"`
	RecordLatch         int             `json:"RecordLatch"`
	RecordMask          string          `json:"RecordMask"`
	ShortMsgEnable      bool            `json:"ShortMsgEnable"`
	ShowInfo            bool            `json:"ShowInfo"`
	ShowInfoMask        string          `json:"ShowInfoMask"`
	SnapEnable          bool            `json:"SnapEnable"`
	SnapShotMask        string          `json:"SnapShotMask"`
	TimeSection         [][]string      `json:"TimeSection"`
	TipEnable           bool            `json:"TipEnable"`
	TourEnable          bool            `json:"TourEnable"`
	TourMask            string          `json:"TourMask"`
	VoiceEnable         bool            `json:"VoiceEnable"`
}

type DetectBlindDetect struct {
	Enable       bool         `json:"Enable"`
	EventHandler EventHandler `json:"EventHandler"`
	Level        int          `json:"Level"`
}
