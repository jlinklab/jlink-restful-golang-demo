package config

type StorageStorageFailureResp struct {
	Name                  string                `json:"Name"`
	Ret                   int                   `json:"Ret"`
	SessionID             string                `json:"SessionID"`
	StorageStorageFailure StorageStorageFailure `json:"Storage.StorageFailure"`
	RetMsg                string                `json:"RetMsg"`
}

type StorageStorageFailure struct {
	Enable       bool         `json:"Enable"`
	EventHandler EventHandler `json:"EventHandler"`
	IsRebooted   int          `json:"IsRebooted"`
	NoBitRateCnt int          `json:"NoBitRateCnt"`
	RebootEnable bool         `json:"RebootEnable"`
}
