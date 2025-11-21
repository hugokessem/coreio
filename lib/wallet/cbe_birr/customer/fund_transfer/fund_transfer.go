package fundtransfer

import (
	"encoding/xml"
	"fmt"
)

type Params struct {
	FTNumber               string
	Password               string
	Timestamp              string
	SecurityCredential     string
	ThirdPartyIdentifier   string
	PrimaryParty           string
	ReceiverParty          string
	Amount                 string
	Currency               string
	Narative               string
	DebitAccountNumber     string
	DebitAccountHolderName string
}
type CustomerFundTransferParams struct {
	FTNumber               string
	Timestamp              string
	PrimaryParty           string
	ReceiverParty          string
	Amount                 string
	Currency               string
	Narative               string
	DebitAccountNumber     string
	DebitAccountHolderName string
}

func NewCustomerFundTransfer(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:api="http://cps.huawei.com/synccpsinterface/api_requestmgr" xmlns:req="http://cps.huawei.com/synccpsinterface/request" xmlns:com="http://cps.huawei.com/synccpsinterface/common" xmlns:cus="http://cps.huawei.com/cpsinterface/customizedrequest">
    <soapenv:Header/>
    <soapenv:Body>
        <api:Request>
            <req:Header>
                <req:Version>1.0</req:Version>
                <req:CommandID>InitTrans_MB E-Money Creation</req:CommandID>
                <req:OriginatorConversationID>%s</req:OriginatorConversationID> 
                <req:Caller>
                    <req:CallerType>2</req:CallerType>
                    <req:ThirdPartyID>%s</req:ThirdPartyID>
                    <req:Password>%s</req:Password>
                </req:Caller>
                <req:KeyOwner>1</req:KeyOwner>
                <req:Timestamp>%s</req:Timestamp>
            </req:Header>
            <req:Body>
                <req:Identity>
                    <req:Initiator>
                        <req:IdentifierType>14</req:IdentifierType>
                        <req:Identifier>Anamail</req:Identifier>
                        <req:SecurityCredential>%s</req:SecurityCredential>
                    </req:Initiator>
                    <req:PrimaryParty>
                        <req:IdentifierType>4</req:IdentifierType>
                        <req:Identifier>%s</req:Identifier>
                    </req:PrimaryParty>
                    <req:ReceiverParty>
                        <req:IdentifierType>1</req:IdentifierType>
                        <req:Identifier>%s</req:Identifier>
                    </req:ReceiverParty>
                </req:Identity>
                <req:TransactionRequest>
                    <req:Parameters>
                        <req:Amount>%s</req:Amount>
                        <req:Currency>%s</req:Currency>
                        <req:ReasonType>%s</req:ReasonType>
                    </req:Parameters>
                </req:TransactionRequest>
                <req:ReferenceData>
                    <req:ReferenceItem>
                        <com:Key>Debited shortcode</com:Key>
                        <com:Value>%s</com:Value>
                    </req:ReferenceItem>
                    <req:ReferenceItem>
                        <com:Key>Debited Customer Name</com:Key>
                        <com:Value>%s</com:Value>
                    </req:ReferenceItem>
                    <req:ReferenceItem>
                        <com:Key>Debited Acct</com:Key>
                        <com:Value>%s</com:Value>
                    </req:ReferenceItem>
                    <req:ReferenceItem>
                        <com:Key>MB txnID</com:Key>
                        <com:Value>%s</com:Value>
                    </req:ReferenceItem>
                </req:ReferenceData>
            </req:Body>
        </api:Request>
    </soapenv:Body>
</soapenv:Envelope>
	`, param.FTNumber, param.ThirdPartyIdentifier, param.Password, param.Timestamp, param.SecurityCredential, param.PrimaryParty, param.ReceiverParty, param.Amount, param.Currency, param.Narative, param.PrimaryParty, param.DebitAccountHolderName, param.DebitAccountNumber, param.FTNumber)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	Result Result `xml:"Result"`
}

type Result struct {
	Header     *Header     `xml:"Header"`
	ResultBody *ResultBody `xml:"Body"`
}
type Header struct {
	Version                         string `xml:"Version"`
	OriginalConverstationIdentifier string `xml:"OriginatorConversationID"`
	ConversationIdentifier          string `xml:"ConversationID"`
}

type ResultBody struct {
	ResultType        string `xml:"ResultType"`
	ResultCode        string `xml:"ResultCode"`
	ResultDescription string `xml:"ResultDesc"`
	TransactionResult *struct {
		TrasnactionId string `xml:"TransactionID"`
	} `xml:"TransactionResult"`
	ReferenceData *struct {
		Details []ReferenceDetail `xml:"ReferenceItem"`
	} `xml:"ReferenceData"`
}

type ReferenceDetail struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

type CustomerFundTransferDetail struct {
	FTNumber                string
	ConverstationIdentifier string
	ReferenceDetail         []ReferenceDetail
}

type CustomerFundTransferResult struct {
	Status  bool
	Detail  CustomerFundTransferDetail
	Message string
}

func ParserCustomreFundTransfer(xmlDate string) (*CustomerFundTransferResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlDate), &env)
	if err != nil {
		return nil, err
	}

	if env.Body.Result.Header != nil && env.Body.Result.ResultBody != nil {
		resp := env.Body.Result
		if resp.ResultBody.ResultCode != "0" {
			return &CustomerFundTransferResult{
				Status:  false,
				Message: resp.ResultBody.ResultDescription,
			}, nil
		}

		if resp.ResultBody.TransactionResult == nil || resp.ResultBody.ReferenceData == nil {
			return &CustomerFundTransferResult{
				Status:  false,
				Message: "API returned failure!",
			}, nil
		}

		return &CustomerFundTransferResult{
			Status: true,
			Detail: CustomerFundTransferDetail{
				FTNumber:                resp.ResultBody.TransactionResult.TrasnactionId,
				ConverstationIdentifier: resp.Header.ConversationIdentifier,
				ReferenceDetail:         resp.ResultBody.ReferenceData.Details,
			},
		}, nil
	}

	return &CustomerFundTransferResult{
		Status:  false,
		Message: "invalid request",
	}, nil

}
