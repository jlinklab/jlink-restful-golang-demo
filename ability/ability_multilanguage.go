package ability

type MultiLanguageResp struct {
	MultiLanguage []string `json:"MultiLanguage"`
	Name          string   `json:"Name"`
	Ret           int      `json:"Ret"`
	SessionID     string   `json:"SessionID"`
	RetMsg        string   `json:"RetMsg"`
}
