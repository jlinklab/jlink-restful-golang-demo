package ability

type CameraResp struct {
	Camera    Camera `json:"Camera"`
	Name      string `json:"Name"`
	Ret       int    `json:"Ret"`
	SessionID string `json:"SessionID"`
	RetMsg    string `json:"RetMsg"`
}

type Camera struct {
	Count      int      `json:"Count"`
	ElecLevel  int      `json:"ElecLevel"`
	IsFishLens int      `json:"IsFishLens"`
	Luminance  int      `json:"Luminance"`
	Speeds     []string `json:"Speeds"`
	Status     int      `json:"Status"`
	Version    string   `json:"Version"`
}
