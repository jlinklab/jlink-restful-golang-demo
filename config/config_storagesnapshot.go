package config

type StorageSnapshotResp struct {
	Name            string            `json:"Name"`
	Ret             int               `json:"Ret"`
	SessionID       string            `json:"SessionID"`
	StorageSnapshot []StorageSnapshot `json:"Storage.Snapshot"`
	RetMsg          string            `json:"RetMsg"`
}

type StorageSnapshot struct {
	Mask        [][]string `json:"Mask"`
	PreSnap     int        `json:"PreSnap"`
	Redundancy  bool       `json:"Redundancy"`
	SnapMode    string     `json:"SnapMode"`
	TimeSection [][]string `json:"TimeSection"`
}
