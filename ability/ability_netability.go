package ability

type NetAbilityResp struct {
	Name       string     `json:"Name"`
	NetAbility NetAbility `json:"NetAbility"`
	Ret        int        `json:"Ret"`
	SessionID  string     `json:"SessionID"`
	RetMsg     string     `json:"RetMsg"`
}

type NetAbility struct {
	CurNatNum int `json:"CurNatNum"`
	CurTCPNum int `json:"CurTcpNum"`
	MaxNatNum int `json:"MaxNatNum"`
	MaxTCPNum int `json:"MaxTcpNum"`
}
