package config

type FVideoTourResp struct {
	Name       string       `json:"Name"`
	Ret        int          `json:"Ret"`
	SessionID  string       `json:"SessionID"`
	FVideoTour []FVideoTour `json:"fVideo.Tour"`
	RetMsg     string       `json:"RetMsg"`
}

type FVideoTour struct {
	Enable   bool     `json:"Enable"`
	Interval int      `json:"Interval"`
	Mask     []string `json:"Mask"`
	Return   bool     `json:"Return"`
	Type     string   `json:"Type"`
}
