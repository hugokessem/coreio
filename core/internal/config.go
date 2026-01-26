package internal

type Config struct {
	Username       string
	Password       string
	Url            string
	FraudAPIConfig FraudAPIConfig
}

type FraudAPIConfig struct {
	Authorization string
	ForwardHost   string
	Url           string
}

var coreAPI *Config

func SetConfig(username, password, url, authorization, fraud_url, forward_host string) *Config {
	coreAPI = &Config{
		Username: username,
		Password: password,
		Url:      url,
		FraudAPIConfig: FraudAPIConfig{
			Authorization: authorization,
			ForwardHost:   forward_host,
			Url:           fraud_url,
		},
	}
	return coreAPI
}

func GetConfig() *Config {
	if coreAPI == nil {
		panic("CBECoreAPI not initialized. Please call SetConfig() first.")
	}
	return coreAPI
}
