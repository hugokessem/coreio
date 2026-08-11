package internal

type Config struct {
	AuthURL     string
	BaseURL     string
	CustomerKey string
	AccessToken string
}

var masterCardAPI *Config

func SetConfig(authURL, baseURL, customerKey, accessToken string) *Config {
	masterCardAPI = &Config{
		AuthURL:     authURL,
		BaseURL:     baseURL,
		CustomerKey: customerKey,
		AccessToken: accessToken,
	}
	return masterCardAPI
}

func GetConfig() *Config {
	if masterCardAPI == nil {
		panic("MasterCardAPI not initialized. Please call SetConfig() first.")
	}
	return masterCardAPI
}
