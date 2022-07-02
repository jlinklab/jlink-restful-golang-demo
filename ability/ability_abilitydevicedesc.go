package ability

type AbilityDeviceDescResp struct {
	Name              string      `json:"Name"`
	Ret               int         `json:"Ret"`
	SessionID         string      `json:"SessionID"`
	AbilityDeviceDesc interface{} `json:"Ability.DeviceDesc"`
	RetMsg            string      `json:"RetMsg"`
}
