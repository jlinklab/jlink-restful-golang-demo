package ability

type BlindCapabilityResp struct {
	BlindCapability BlindCapability `json:"BlindCapability"`
	Name            string          `json:"Name"`
	Ret             int             `json:"Ret"`
	SessionID       string          `json:"SessionID"`
	RetMsg          string          `json:"RetMsg"`
}

type BlindCapability struct {
	BlindCoverNum int `json:"BlindCoverNum"`
}
