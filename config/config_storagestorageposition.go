package config

type StorageStoragePositionResp struct {
	Name                   string                 `json:"Name"`
	Ret                    int                    `json:"Ret"`
	SessionID              string                 `json:"SessionID"`
	StorageStoragePosition StorageStoragePosition `json:"Storage.StoragePosition"`
	RetMsg                 string                 `json:"RetMsg"`
}

type StorageStoragePosition struct {
	DVD  bool `json:"DVD"`
	SATA bool `json:"SATA"`
	SD   bool `json:"SD"`
	USB  bool `json:"USB"`
}
