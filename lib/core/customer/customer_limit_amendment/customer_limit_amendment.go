package customerlimitamendment

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username         string
	Password         string
	CustomerID       string
	AppUserMaxLimit  string
	USSDUserMaxLimit string
}

type CustomerLimitAmendmentParam struct {
	CustomerID       string
	AppUserMaxLimit  string
	USSDUserMaxLimit string
}

func NewCustomerLimitAmendment(param Params) string {
	return fmt.Sprintf(`
		<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
xmlns:cbes="http://temenos.com/CBESUPERAPP"
xmlns:cus="http://temenos.com/CUSTOMERLIMITCUSTOMLIMIT">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:CustomerLimitAmendment>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <OfsFunction/>
            <CUSTOMERLIMITCUSTOMLIMITType id="%s">
                <cus:gUSERCHANNELTYPE g="1">
                    <cus:mUSERCHANNELTYPE m="1">
                        <cus:USERCHANNELTYPE>APP</cus:USERCHANNELTYPE>
                        <cus:ACCOUNT/>
                        <cus:USERCUSTOMERID/>
                        <cus:USERMAXLIMIT>%s</cus:USERMAXLIMIT>
                    </cus:mUSERCHANNELTYPE>
                    <cus:mUSERCHANNELTYPE m="2">
                        <cus:USERCHANNELTYPE>USSD</cus:USERCHANNELTYPE>
                        <cus:ACCOUNT/>
                        <cus:USERCUSTOMERID/>
                        <cus:USERMAXLIMIT>%s</cus:USERMAXLIMIT>
                    </cus:mUSERCHANNELTYPE>
                </cus:gUSERCHANNELTYPE>
            </CUSTOMERLIMITCUSTOMLIMITType>
        </cbes:CustomerLimitAmendment>
    </soapenv:Body>
</soapenv:Envelope>
	`, param.Password, param.Username, param.CustomerID, param.AppUserMaxLimit, param.USSDUserMaxLimit)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	CustomerLimitAmendmentResponse *CustomerLimitAmendmentResponse `xml:"CustomerLimitAmendmentResponse"`
}

type CustomerLimitAmendmentResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
		MessageId        string `xml:"messageId"`
		Application      string `xml:"application"`
		TransactionId    string `xml:"transactionId"`
	} `xml:"Status"`
	CustomerLimitType *struct {
		CustomLimitType *CustomLimitType `xml:"gUSERCHANNELTYPE"`
	} `xml:"CUSTOMERLIMITType"`
}

type CustomLimitType struct {
	UserChannelType []UserChannelType `xml:"mUSERCHANNELTYPE"`
	UserMaxLimit    string            `xml:"USERMAXLIMIT"`
	UserCustomerId  string            `xml:"USERCUSTOMERID"`
	Account         string            `xml:"ACCOUNT"`
	Currency        string            `xml:"CURRENCY"`
	Inputter        string            `xml:"INPUTTER"`
	Datetime        string            `xml:"DATETIME"`
	Authoriser      string            `xml:"AUTHORISER"`
	Cocode          string            `xml:"COCODE"`
	Deptcode        string            `xml:"DEPTCODE"`
}

type UserChannelType struct {
	UserChannelType string `xml:"USERCHANNELTYPE"`
	UserMaxLimit    string `xml:"USERMAXLIMIT"`
}

type CustomerLimitAmendmentResult struct {
	Success bool
	Detail  *CustomLimitType
	Message []string
}

func ParseCustomerLimitAmendmentSOAP(xmlData string) (*CustomerLimitAmendmentResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}

	if env.Body.CustomerLimitAmendmentResponse != nil {
		resp := env.Body.CustomerLimitAmendmentResponse
		if resp.Status == nil {
			return &CustomerLimitAmendmentResult{
				Success: false,
				Message: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &CustomerLimitAmendmentResult{
				Success: false,
				Message: []string{"API returned failure"},
			}, nil
		}

		if resp.CustomerLimitType == nil {
			return &CustomerLimitAmendmentResult{
				Success: false,
				Message: []string{},
			}, nil
		}

		return &CustomerLimitAmendmentResult{
			Success: true,
			Detail: &CustomLimitType{
				UserChannelType: resp.CustomerLimitType.CustomLimitType.UserChannelType,
				UserMaxLimit:    resp.CustomerLimitType.CustomLimitType.UserMaxLimit,
				UserCustomerId:  resp.CustomerLimitType.CustomLimitType.UserCustomerId,
				Account:         resp.CustomerLimitType.CustomLimitType.Account,
				Currency:        resp.CustomerLimitType.CustomLimitType.Currency,
				Inputter:        resp.CustomerLimitType.CustomLimitType.Inputter,
				Datetime:        resp.CustomerLimitType.CustomLimitType.Datetime,
				Authoriser:      resp.CustomerLimitType.CustomLimitType.Authoriser,
				Cocode:          resp.CustomerLimitType.CustomLimitType.Cocode,
				Deptcode:        resp.CustomerLimitType.CustomLimitType.Deptcode,
			},
		}, nil
	}

	return &CustomerLimitAmendmentResult{
		Success: false,
		Message: []string{"Invalid response type"},
	}, nil
}
