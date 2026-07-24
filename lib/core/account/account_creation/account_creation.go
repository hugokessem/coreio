package acccountcreation

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username       string
	Password       string
	CustomerNumber string
	Category       string
	Currency       string
	AccountOfficer string
	Url            string
	Header         map[string]string
}

type AccountCreationParams struct {
	CustomerNumber string
	Category       string
	Currency       string
	AccountOfficer string
	Url            string
	Header         map[string]string
}

func NewAccountCreation(param Params) string {
	return fmt.Sprintf(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
xmlns:iib="http://temenos.com/IIBONBOARDING"
xmlns:acc="http://temenos.com/ACCOUNTCREATEINDIVIDUAL">
    <soapenv:Header/>
    <soapenv:Body>
        <iib:AccountOpeningSuperApp>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <OfsFunction></OfsFunction>
            <ACCOUNTCREATEINDIVIDUALType id="">
                <acc:CUSTOMER>%s</acc:CUSTOMER>
                <acc:CATEGORY>%s</acc:CATEGORY>
                <acc:CURRENCY>%s</acc:CURRENCY>
                <acc:ACCOUNTOFFICER>%s</acc:ACCOUNTOFFICER>
            </ACCOUNTCREATEINDIVIDUALType>
        </iib:AccountOpeningSuperApp>
    </soapenv:Body>
</soapenv:Envelope>`,
		param.Password,
		param.Username,
		param.CustomerNumber,
		param.Category,
		param.Currency,
		param.AccountOfficer,
	)
}

type Envelope struct {
	Body Body `xml:"Body"`
}

type Body struct {
	AccountCreationResponse *AccountCreationResponse `xml:"AccountOpeningSuperAppResponse"`
}

type AccountCreationResponse struct {
	Status *struct {
		TransactionId string   `xml:"transactionId"`
		Success       string   `xml:"successIndicator"`
		Application   string   `xml:"application"`
		Messages      []string `xml:"messages"`
	} `xml:"Status"`
	AccountType AccountType `xml:"ACCOUNTType"`
}

type AccountType struct {
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

type AccountTypeDetail struct {
	AccountNumber   string
	Customer        string
	Category        string
	AccountTitle    string
	ShortTitle      string
	PositionType    string
	Currency        string
	CurrencyMarket  string
	AccountOfficer  string
	PostingRestrict string
	ConditionGroup  string
	CapDateCharge   string
	Passbook        string
	OpeningDate     string
	OpenCategory    string
	ChargeCcy       string
	ChargeMkt       string
	InterestCcy     string
	InterestMkt     string
	AltAcctTypes    []string
	AllowNetting    string
	SingleLimit     string
	CurrNo          string
	Inputter        string
	Datetime        string
	Authoriser      string
	CoCode          string
	DeptCode        string
	HasJointCust    string
	ProductType     string
}

type AccountCreationResult struct {
	Success  bool
	Detail   *AccountTypeDetail
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

		if strings.ToLower(resp.Status.Success) != "success" {
			return &AccountCreationResult{
				Success:  false,
				Messages: resp.Status.Messages,
			}, nil
		}

		if resp.AccountType.AccountNumber == "" {
			return &AccountCreationResult{
				Success:  true,
				Messages: resp.Status.Messages,
			}, nil
		}

		altAcctTypes := make([]string, 0, len(resp.AccountType.GlobalAltAcctType.MALTACCTTYPE))
		for _, altAcctType := range resp.AccountType.GlobalAltAcctType.MALTACCTTYPE {
			altAcctTypes = append(altAcctTypes, altAcctType.AltAcctType)
		}

		detail := &AccountTypeDetail{
			AccountNumber:   resp.AccountType.AccountNumber,
			Customer:        resp.AccountType.Customer,
			Category:        resp.AccountType.Category,
			AccountTitle:    resp.AccountType.GlobalAccountTitle.AccountTitle,
			ShortTitle:      resp.AccountType.GlobalShortTitle.ShortTitle,
			PositionType:    resp.AccountType.PositionType,
			Currency:        resp.AccountType.Currency,
			CurrencyMarket:  resp.AccountType.CurrencyMarket,
			AccountOfficer:  resp.AccountType.AccountOfficer,
			PostingRestrict: resp.AccountType.GlobalPostingRestrict.PostingRestrict,
			ConditionGroup:  resp.AccountType.ConditionGroup,
			CapDateCharge:   resp.AccountType.GlobalCapDateCharge.CapDateCharge,
			Passbook:        resp.AccountType.Passbook,
			OpeningDate:     resp.AccountType.OpeningDate,
			OpenCategory:    resp.AccountType.OpenCategory,
			ChargeCcy:       resp.AccountType.ChargeCcy,
			ChargeMkt:       resp.AccountType.ChargeMkt,
			InterestCcy:     resp.AccountType.InterestCcy,
			InterestMkt:     resp.AccountType.InterestMkt,
			AltAcctTypes:    altAcctTypes,
			AllowNetting:    resp.AccountType.AllowNetting,
			SingleLimit:     resp.AccountType.SingleLimit,
			CurrNo:          resp.AccountType.CurrNo,
			Inputter:        resp.AccountType.GlobalInputter.Inputter,
			Datetime:        resp.AccountType.GlobalDatetime.Datetime,
			Authoriser:      resp.AccountType.Authoriser,
			CoCode:          resp.AccountType.CoCode,
			DeptCode:        resp.AccountType.DeptCode,
			HasJointCust:    resp.AccountType.HasJointCust,
			ProductType:     resp.AccountType.ProductType,
		}

		return &AccountCreationResult{
			Success:  true,
			Detail:   detail,
			Messages: resp.Status.Messages,
		}, nil
	}

	return &AccountCreationResult{
		Success:  false,
		Messages: []string{"Invalid response type"},
	}, nil
}
