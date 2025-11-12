package internal

type Config struct {
	Username string
	Password string
	Url      string
}

var coreAPI *Config

func SetConfig(username, password, url string) *Config {
	coreAPI = &Config{
		Username: username,
		Password: password,
		Url:      url,
	}
	return coreAPI
}

func GetConfig() *Config {
	if coreAPI == nil {
		panic("CBECoreAPI not initialized. Please call SetConfig() first.")
	}
	return coreAPI
}
