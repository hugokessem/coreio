package exchangerate

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username string
	Password string
}

func NewExchangeRate(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
xmlns:cbes="http://temenos.com/CBESUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:ExchangeRateSuperApp>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <EXCHANGERATESUPERAPPType>
                <enquiryInputCollection>
                    <columnName/>
                    <criteriaValue/>
                    <operand/>
                </enquiryInputCollection>
            </EXCHANGERATESUPERAPPType>
        </cbes:ExchangeRateSuperApp>
    </soapenv:Body>
</soapenv:Envelope>
	`, param.Password, param.Username)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	ExchangeRateSuperAppResponse *ExchangeRateSuperAppResponse `xml:"ExchangeRateSuperAppResponse"`
}

type ExchangeRateSuperAppResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
	} `xml:"Status"`
	ExchangeRateSuperAppType *struct {
		Group *struct {
			Details []ExchangeRateDetail `xml:"mEXCHANGERATESUPERAPPDetailType"`
		} `xml:"gEXCHANGERATESUPERAPPDetailType"`
	} `xml:"EXCHANGERATESUPERAPPType"`
}

type ExchangeRateDetail struct {
	ID             string `xml:"ID"`
	NumericCCYCode string `xml:"NUMERICCCYCODE"`
	CCYName        string `xml:"CCYNAME"`
	CurrencyMarket string `xml:"CURRENCYMARKET"`
	BuyRate        string `xml:"BUYRATE"`
	SellRate       string `xml:"SELLRATE"`
	MidRate        string `xml:"MIDRATE"`
}

type ExchangeRateResult struct {
	Success bool
	Detail  []ExchangeRateDetail
	Message []string
}

func ParseExchangeRateSOAP(xmlData string) (*ExchangeRateResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}

	if env.Body.ExchangeRateSuperAppResponse != nil {
		resp := env.Body.ExchangeRateSuperAppResponse
		if resp.Status == nil {
			return &ExchangeRateResult{
				Success: false,
				Message: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &ExchangeRateResult{
				Success: false,
				Message: []string{"API returned failure"},
			}, nil
		}

		if resp.ExchangeRateSuperAppType == nil || resp.ExchangeRateSuperAppType.Group == nil {
			return &ExchangeRateResult{
				Success: false,
				Message: []string{},
			}, nil
		}
		details := resp.ExchangeRateSuperAppType.Group.Details
		detailsList := make([]ExchangeRateDetail, len(details))
		for i, detail := range details {
			detailsList[i] = ExchangeRateDetail{
				ID:             detail.ID,
				NumericCCYCode: detail.NumericCCYCode,
				CCYName:        detail.CCYName,
				CurrencyMarket: detail.CurrencyMarket,
				BuyRate:        detail.BuyRate,
				SellRate:       detail.SellRate,
				MidRate:        detail.MidRate,
			}
		}

		return &ExchangeRateResult{
			Success: true,
			Detail:  detailsList,
		}, nil
	}

	return &ExchangeRateResult{
		Success: false,
		Message: []string{"Invalid response type"},
	}, nil
}
