package acccountcreation

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username       string
	Password       string
	Customer       string
	Category       string
	Currency       string
	AccountOfficer string
}

type AccountCreationParams struct {
	Customer       string
	Category       string
	Currency       string
	AccountOfficer string
}

func NewAccountCreation(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP" xmlns:acc="http://temenos.com/ACCOUNTCREATEINDIVIDUAL">
		<soapenv:Header/>
		<soapenv:Body>
			<cbes:AccountOpeningSuperApp>
			<WebRequestCommon>
			<company/>
			<password>%s</password>
			<userName>%s</userName>
			</WebRequestCommon>
			<OfsFunction/>
			<ACCOUNTCREATEINDIVIDUALType id="">
				<acc:CUSTOMER>%s</acc:CUSTOMER>
				<acc:CATEGORY>%s</acc:CATEGORY>
				<acc:CURRENCY>%s</acc:CURRENCY>
				<acc:ACCOUNTOFFICER>%s</acc:ACCOUNTOFFICER>
			</ACCOUNTCREATEINDIVIDUALType>
			</cbes:AccountOpeningSuperApp>
		</soapenv:Body>
	</soapenv:Envelope>
	`, param.Password, param.Username, param.Customer, param.Category, param.Currency, param.AccountOfficer)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	AccountCreationResponse *AccountCreationResponse `xml:"AccountOpeningSuperAppResponse"`
}

type AccountCreationResponse struct {
	Status *struct {
		SuccessIndicator string   `xml:"successIndicator"`
		Messages         []string `xml:"messages"`
		TransactionId    string   `xml:"transactionId"`
		Application      string   `xml:"application"`
	} `xml:"Status"`
	AccountCreationDetail *AccountCreationDetail `xml:"ACCOUNTType"`
}

type AccountCreationDetail struct {
	XMLName            xml.Name `xml:"ACCOUNTType"`
	AccountNumber      string   `xml:"id,attr"`
	Customer           string   `xml:"CUSTOMER"`
	Category           string   `xml:"CATEGORY"`
	GlobalAccountTitle struct {
		AccountTitle string `xml:"ACCOUNTTITLE1"`
	} `xml:"gACCOUNTTITLE1"`
	GlobalShortTitle struct {
		ShortTitle string `xml:"SHORTTITLE"`
	} `xml:"gSHORTTITLE"`
	PositionType          string `xml:"POSITIONTYPE"`
	Currency              string `xml:"CURRENCY"`
	CurrencyMarket        string `xml:"CURRENCYMARKET"`
	AccountOfficer        string `xml:"ACCOUNTOFFICER"`
	GlobalPostingRestrict struct {
		PostingRestrict string `xml:"POSTINGRESTRICT"`
	} `xml:"gPOSTINGRESTRICT"`
	ConditionGroup      string `xml:"CONDITIONGROUP"`
	GlobalCapDateCharge struct {
		CapDateCharge string `xml:"CAPDATECHARGE"`
	} `xml:"gCAPDATECHARGE"`
	Passbook          string `xml:"PASSBOOK"`
	OpeningDate       string `xml:"OPENINGDATE"`
	OpenCategory      string `xml:"OPENCATEGORY"`
	ChargeCcy         string `xml:"CHARGECCY"`
	ChargeMkt         string `xml:"CHARGEMKT"`
	InterestCcy       string `xml:"INTERESTCCY"`
	InterestMkt       string `xml:"INTERESTMKT"`
	GlobalAltAcctType struct {
		MALTACCTTYPE []struct {
			AltAcctType string `xml:"ALTACCTTYPE"`
		} `xml:"mALTACCTTYPE"`
	} `xml:"gALTACCTTYPE"`
	AllowNetting   string `xml:"ALLOWNETTING"`
	SingleLimit    string `xml:"SINGLELIMIT"`
	CurrNo         string `xml:"CURRNO"`
	GlobalInputter struct {
		Inputter string `xml:"INPUTTER"`
	} `xml:"gINPUTTER"`
	GlobalDatetime struct {
		Datetime string `xml:"DATETIME"`
	} `xml:"gDATETIME"`
	Authoriser   string `xml:"AUTHORISER"`
	CoCode       string `xml:"COCODE"`
	DeptCode     string `xml:"DEPTCODE"`
	HasJointCust string `xml:"HASJOINTCUST"`
	ProductType  string `xml:"PRODUCTTYPE"`
}

type AccountCreationResult struct {
	Success  bool
	Detail   *AccountCreationDetail
	Messages []string
}

func ParseAccountCreationSOAP(xmlData string) (*AccountCreationResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.AccountCreationResponse != nil {
		resp := env.Body.AccountCreationResponse
		if resp.Status == nil {
			return &AccountCreationResult{
				Success:  false,
				Messages: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &AccountCreationResult{
				Success:  false,
				Messages: resp.Status.Messages,
			}, nil
		}

		if resp.AccountCreationDetail == nil {
			return &AccountCreationResult{
				Success:  false,
				Messages: []string{},
			}, nil
		}

		return &AccountCreationResult{
			Success:  true,
			Detail:   resp.AccountCreationDetail,
			Messages: resp.Status.Messages,
		}, nil
	}

	return &AccountCreationResult{
		Success:  false,
		Messages: []string{"Invalid Response!"},
	}, nil
}
