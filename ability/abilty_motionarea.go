package ability

type MotionAreaResp struct {
	MotionArea MotionArea `json:"MotionArea"`
	Name       string     `json:"Name"`
	Ret        int        `json:"Ret"`
	SessionID  string     `json:"SessionID"`
	RetMsg     string     `json:"RetMsg"`
}

type MotionArea struct {
	GridColumn int `json:"GridColumn"`
	GridRow    int `json:"GridRow"`
}
