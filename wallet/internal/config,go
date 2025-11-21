package internal

type Config struct {
	Url                  string `json:"Url"`
	Password             string `json:"Password"`
	Authorization        string `json:"Authorization"`
	IIBAuthorization     string `json:"iib_authorization"`
	SecurityCredential   string `json:"SecurityCredential"`
	ThirdPartyIdentifier string `json:"ThirdPartyID"`
}

var coreAPI *Config

func SetConfig(url, password, authorization, iibAuthorization, securityCredential, thirdPartyIdentifier string) *Config {
	coreAPI = &Config{
		Url:                  url,
		Password:             password,
		Authorization:        authorization,
		IIBAuthorization:     iibAuthorization,
		SecurityCredential:   securityCredential,
		ThirdPartyIdentifier: thirdPartyIdentifier,
	}
	return coreAPI
}

func GetConfig() *Config {
	if coreAPI == nil {
		panic("Wallet API not initialized. Please call SetConfig() first.")
	}
	return coreAPI
}
