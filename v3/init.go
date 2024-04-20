package v3

type JLinkClient struct {
	Uuid      string `json:"uuid"`
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
	Movecard  int    `json:"movecard"`
}

type JLinkDevice struct {
	JClient      *JLinkClient `json:"jClient"`
	Sn           string       `json:"sn"`
	DevUsername  string       `json:"devUsername"`
	DevPassword  string       `json:"devPassword"`
	Jdtoken      string       `json:"jdtoken"`
	JdLoginToken string       `json:"jdLoginToken"`
}

type JUser struct {
	JClient  *JLinkClient `json:"jClient"`
	Username string       `json:"devUsername"`
	Password string       `json:"devPassword"`
	Jutoken  string       `json:"jutoken"`
}

// uuid,appKey,appSecret,movecard Obtained from (open.jftech.com)
func NewJLinkClient(uuid, appKey, appSecret string, movecard int) *JLinkClient {
	jlinkclient := &JLinkClient{
		Uuid:      uuid,
		AppKey:    appKey,
		AppSecret: appSecret,
		Movecard:  movecard,
	}

	return jlinkclient
}

// Sn is device serial number devUsername Indicates the device login user name devPassword Indicates the device login password
func NewJLinkDevice(jClient *JLinkClient, sn, devUsername, devPassword string) *JLinkDevice {
	jlinkdevice := &JLinkDevice{
		JClient:     jClient,
		Sn:          sn,
		DevUsername: devUsername,
		DevPassword: devPassword,
	}
	return jlinkdevice
}

func NewJLinkUser(jClient *JLinkClient, username, password string) *JUser {
	juser := &JUser{
		JClient:  jClient,
		Username: username,
		Password: password,
	}
	return juser

}
