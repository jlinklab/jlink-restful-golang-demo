package config

type NetWorkPMSResp struct {
	Name       string     `json:"Name"`
	NetWorkPMS NetWorkPMS `json:"NetWork.PMS"`
	Ret        int        `json:"Ret"`
	SessionID  string     `json:"SessionID"`
	RetMsg     string     `json:"RetMsg"`
}

type NetWorkPMS struct {
	BoxID        string `json:"BoxID"`
	Enable       bool   `json:"Enable"`
	Port         int    `json:"Port"`
	PushInterval int    `json:"PushInterval"`
	ServName     string `json:"ServName"`
}
