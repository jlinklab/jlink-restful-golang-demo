package config

type StorageStorageNotExistResp struct {
	Name                   string                 `json:"Name"`
	Ret                    int                    `json:"Ret"`
	SessionID              string                 `json:"SessionID"`
	StorageStorageNotExist StorageStorageNotExist `json:"Storage.StorageNotExist"`
	RetMsg                 string                 `json:"RetMsg"`
}

type StorageStorageNotExist struct {
	Enable       bool         `json:"Enable"`
	EventHandler EventHandler `json:"EventHandler"`
}
