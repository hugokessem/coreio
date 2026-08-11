package internal

type Config struct {
	AuthURL     string
	BaseURL     string
	CustomerKey string
}

var masterCardAPI *Config

func SetConfig(authURL, baseURL, customerKey string) *Config {
	masterCardAPI = &Config{
		AuthURL:     authURL,
		BaseURL:     baseURL,
		CustomerKey: customerKey,
	}
	return masterCardAPI
}
