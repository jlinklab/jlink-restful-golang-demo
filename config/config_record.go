package config

type RecordResp struct {
	Name      string   `json:"Name"`
	Record    []Record `json:"Record"`
	Ret       int      `json:"Ret"`
	SessionID string   `json:"SessionID"`
	RetMsg    string   `json:"RetMsg"`
}

type Record struct {
	Mask         [][]string `json:"Mask"`
	PacketLength int        `json:"PacketLength"`
	PreRecord    int        `json:"PreRecord"`
	RecordMode   string     `json:"RecordMode"`
	Redundancy   bool       `json:"Redundancy"`
	TimeSection  [][]string `json:"TimeSection"`
}
