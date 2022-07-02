package config

type StorageStorageLowSpaceResp struct {
	Name                   string                 `json:"Name"`
	Ret                    int                    `json:"Ret"`
	SessionID              string                 `json:"SessionID"`
	StorageStorageLowSpace StorageStorageLowSpace `json:"Storage.StorageLowSpace"`
	RetMsg                 string                 `json:"RetMsg"`
}

type StorageStorageLowSpace struct {
	Enable           bool         `json:"Enable"`
	EnableType       int          `json:"EnableType"`
	EventHandler     EventHandler `json:"EventHandler"`
	LowerLimit       int          `json:"LowerLimit"`
	LowerLimitSpace  int          `json:"LowerLimitSpace"`
	RecordTime       int          `json:"RecordTime"`
	RecordTimeEnable bool         `json:"RecordTimeEnable"`
}
