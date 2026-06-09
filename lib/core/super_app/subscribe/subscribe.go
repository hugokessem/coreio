package subscribe

import (
	"encoding/xml"
	"fmt"
)

type Params struct {
	Username     string
	Password     string
	UserCode     string
	CustomerName string
	CustomerId   string
	PhoneNumber  string
	Email        string
	Channel      string
	Description  string
	BranchCode   string
	SalesId      string
}

type SubscribeParams struct {
	CustomerName string
	CustomerId   string
	PhoneNumber  string
	Email        string
	Channel      string
	Description  string
	BranchCode   string
	SalesId      string
}

func NewSubscribe(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
xmlns:cbes="http://temenos.com/CBESUPERAPP" xmlns:sup="http://temenos.com/SUPERAPPUSERCREATE">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:RegisterSuperAppUser>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <OfsFunction/>
            <SUPERAPPUSERCREATEType id="%s">
                <sup:CUSTOMERNAME>%s</sup:CUSTOMERNAME>
                <sup:CUSTOMERID>%s</sup:CUSTOMERID>
                <sup:PHONENO>%s</sup:PHONENO>
                <sup:EMAIL>%s</sup:EMAIL>
                <sup:CHANNEL>%s</sup:CHANNEL>
                <sup:DESCRIPTION>%s</sup:DESCRIPTION>
                <sup:BRANCHCODE>%s</sup:BRANCHCODE>
                <sup:SALESID>%s</sup:SALESID>
            </SUPERAPPUSERCREATEType>
        </cbes:RegisterSuperAppUser>
    </soapenv:Body>
</soapenv:Envelope>
`, param.Password, param.Username, param.UserCode, param.CustomerName, param.CustomerId, param.PhoneNumber, param.Email, param.Channel, param.Description, param.BranchCode, param.SalesId)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	RegisterSuperAppUserResponse *RegisterSuperAppUserResponse `xml:"RegisterSuperAppUserResponse"`
}

type RegisterSuperAppUserResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
		TransactionId    string `xml:"transactionId"`
		MessageId        string `xml:"messageId"`
		Application      string `xml:"application"`
	}
	SuperappUserType *SuperappUserType `xml:"SUPERAPPUSERType"`
}

type SuperappUserType struct {
	CustomerName    string `xml:"CUSTOMERNAME"`
	CustomerId      string `xml:"CUSTOMERID"`
	Status          string `xml:"STATUS"`
	PhoneNumber     string `xml:"PHONENO"`
	Email           string `xml:"EMAIL"`
	Channel         string `xml:"CHANNEL"`
	Description     string `xml:"DESCRIPTION"`
	LastUpdatedDate string `xml:"LASTUPDATEDDATE"`
	BranchCode      string `xml:"BRANCHCODE"`
	SalesId         string `xml:"SALESID"`
	CurrNo          string `xml:"CURRNO"`
	Datetime        struct {
		DateTime string `xml:"DATETIME"`
	} `xml:"DATETIME"`
	Authoriser string `xml:"AUTHORISER"`
	Code       string `xml:"COCODE"`
	DeptCode   string `xml:"DEPTCODE"`
}

type SubscribeResult struct {
	Success  bool
	Detail   *SuperappUserType
	Messages []string
}

func ParseSubscribeResponseSOAP(xmlData string) (*SubscribeResult, error) {
	var envelope Envelope
	err := xml.Unmarshal([]byte(xmlData), &envelope)
	if err != nil {
		return nil, err
	}
	if envelope.Body.RegisterSuperAppUserResponse == nil {
		return &SubscribeResult{
			Success:  false,
			Messages: []string{"RegisterSuperAppUserResponse is nil"},
		}, nil
	}

	if envelope.Body.RegisterSuperAppUserResponse.Status == nil {
		return &SubscribeResult{
			Success:  false,
			Messages: []string{"Status is nil"},
		}, nil
	}

	if envelope.Body.RegisterSuperAppUserResponse.Status.SuccessIndicator != "Success" {
		return &SubscribeResult{
			Success:  false,
			Messages: []string{"SuccessIndicator is not Success"},
		}, nil
	}

	if envelope.Body.RegisterSuperAppUserResponse.SuperappUserType == nil {
		return &SubscribeResult{
			Success:  false,
			Messages: []string{"SuperappUserType is nil"},
		}, nil
	}

	return &SubscribeResult{
		Success: true,
		Detail:  envelope.Body.RegisterSuperAppUserResponse.SuperappUserType,
	}, nil
}
