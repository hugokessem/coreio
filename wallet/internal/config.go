// coreio target: wallet/internal/config.go
package internal

type Config struct {
	Url                         string `json:"Url"`
	Password                    string `json:"Password"`
	Authorization               string `json:"Authorization"`
	IIBAuthorization            string `json:"iib_authorization"`
	SecurityCredential          string `json:"SecurityCredential"` // KYC credential (lookups)
	ThirdPartyIdentifier        string `json:"ThirdPartyID"`
	KYCThirdPartyIdentifier     string `json:"KYCThirdPartyID"`           // InitiatorIdentifier for lookups
	PaymentThirdPartyIdentifier string `json:"PaymentThirdPartyID"`       // InitiatorIdentifier for transfers
	PaymentSecurityCredential   string `json:"PaymentSecurityCredential"` // payment credential (transfers)
	ShortCode                   string `json:"ShortCode"`                 // customer FT ShortCode
}

var coreAPI *Config

func SetConfig(
	url, password, authorization, iibAuthorization,
	securityCredential, thirdPartyIdentifier,
	kycThirdPartyIdentifier, paymentThirdPartyIdentifier,
	paymentSecurityCredential, shortCode string,
) *Config {
	coreAPI = &Config{
		Url:                         url,
		Password:                    password,
		Authorization:               authorization,
		IIBAuthorization:            iibAuthorization,
		SecurityCredential:          securityCredential,
		ThirdPartyIdentifier:        thirdPartyIdentifier,
		KYCThirdPartyIdentifier:     kycThirdPartyIdentifier,
		PaymentThirdPartyIdentifier: paymentThirdPartyIdentifier,
		PaymentSecurityCredential:   paymentSecurityCredential,
		ShortCode:                   shortCode,
	}
	return coreAPI
}

func GetConfig() *Config {
	if coreAPI == nil {
		panic("Wallet API not initialized. Please call SetConfig() first.")
	}
	return coreAPI
}
