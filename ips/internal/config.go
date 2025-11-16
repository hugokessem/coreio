package internal

type Config struct {
	Username        string
	Password        string
	GrantType       string
	JwtAssertion    string
	MBAuthorization string
	Authorization   string
	Url             string
}

var coreAPI *Config

func SetConfig(username, password, granttype, jwtAssertion, mbAuthorization, authorization, url string) *Config {
	coreAPI = &Config{
		Username:        username,
		Password:        password,
		GrantType:       granttype,
		JwtAssertion:    jwtAssertion,
		MBAuthorization: mbAuthorization,
		Authorization:   authorization,
		Url:             url,
	}
	return coreAPI
}

func GetConfig() *Config {
	if coreAPI == nil {
		panic("IPSAPI not initialized. Please call SetConfig() first.")
	}
	return coreAPI
}
