// coreio target: lib/wallet/cbe_birr/customer/fund_transfer/fund_transfer.go
package fundtransfer

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	FTNumber               string
	Password               string
	Timestamp              string
	ThirdPartyIdentifier   string
	InitiatorIdentifier    string // payment third-party ID (was hardcoded "Anamail")
	SecurityCredential     string // payment security credential
	ShortCode              string
	ReceiverParty          string
	Amount                 string
	Currency               string
	DebitAccountNumber     string
	DebitAccountHolderName string
}

type CustomerFundTransferParams struct {
	FTNumber               string
	Timestamp              string
	PrimaryParty           string // short code
	ReceiverParty          string
	Amount                 string
	Currency               string
	Narative               string
	DebitAccountNumber     string
	DebitAccountHolderName string
}

func NewCustomerFundTransfer(param Params) string {
	xe := escapeXML
	return fmt.Sprintf(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:api="http://cps.huawei.com/synccpsinterface/api_requestmgr" xmlns:req="http://cps.huawei.com/synccpsinterface/request" xmlns:com="http://cps.huawei.com/synccpsinterface/common" xmlns:cus="http://cps.huawei.com/cpsinterface/customizedrequest">
    <soapenv:Header/>
    <soapenv:Body>
        <api:Request>
            <req:Header>
                <req:Version>1.0</req:Version>
                <req:CommandID>InitTrans_MB EMC</req:CommandID>
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
                        <req:IdentifierType>11</req:IdentifierType>
                        <req:Identifier>%s</req:Identifier>
                        <req:SecurityCredential>%s</req:SecurityCredential>
                        <req:ShortCode>%s</req:ShortCode>
                    </req:Initiator>
                    <req:ReceiverParty>
                        <req:IdentifierType>1</req:IdentifierType>
                        <req:Identifier>%s</req:Identifier>
                    </req:ReceiverParty>
                </req:Identity>
                <req:TransactionRequest>
                    <req:Parameters>
                        <req:Amount>%s</req:Amount>
                        <req:Currency>%s</req:Currency>
                        <req:ReasonType>Transfer from CBE SuperAPP to CBEBIRR</req:ReasonType>
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
</soapenv:Envelope>`, xe(param.FTNumber), xe(param.ThirdPartyIdentifier), xe(param.Password), xe(param.Timestamp),
		xe(param.InitiatorIdentifier), xe(param.SecurityCredential), xe(param.ShortCode),
		xe(param.ReceiverParty), xe(param.Amount), xe(param.Currency),
		xe(param.ShortCode), xe(param.DebitAccountHolderName), xe(param.DebitAccountNumber), xe(param.FTNumber))
}

type ReferenceDetail struct {
	Key   string `xml:"http://cps.huawei.com/synccpsinterface/common Key"`
	Value string `xml:"http://cps.huawei.com/synccpsinterface/common Value"`
}

type CustomerFundTransferDetail struct {
	FTNumber                string
	ConverstationIdentifier string
	TransactionID           string
	ResultCode              string
	ResultDesc              string
	ReferenceDetail         []ReferenceDetail
}

type CustomerFundTransferResult struct {
	Status  bool
	Detail  *CustomerFundTransferDetail
	Message string
}

type envelope struct {
	Body struct {
		Result struct {
			Header struct {
				ConversationID string `xml:"http://cps.huawei.com/synccpsinterface/result ConversationID"`
			} `xml:"http://cps.huawei.com/synccpsinterface/result Header"`
			ResultBody struct {
				ResultCode        string `xml:"http://cps.huawei.com/synccpsinterface/result ResultCode"`
				ResultDesc        string `xml:"http://cps.huawei.com/synccpsinterface/result ResultDesc"`
				TransactionResult *struct {
					TransactionID string `xml:"http://cps.huawei.com/synccpsinterface/result TransactionID"`
				} `xml:"http://cps.huawei.com/synccpsinterface/result TransactionResult"`
				ReferenceData *struct {
					Details []ReferenceDetail `xml:"http://cps.huawei.com/synccpsinterface/common ReferenceItem"`
				} `xml:"http://cps.huawei.com/synccpsinterface/result ReferenceData"`
			} `xml:"http://cps.huawei.com/synccpsinterface/result Body"`
		} `xml:"http://cps.huawei.com/synccpsinterface/api_requestmgr Result"`
	} `xml:"Body"`
}

func ParserCustomreFundTransfer(xmlData string) (*CustomerFundTransferResult, error) {
	var env envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	rb := env.Body.Result.ResultBody
	if rb.ResultCode != "0" {
		return &CustomerFundTransferResult{
			Status:  false,
			Message: rb.ResultDesc,
			Detail: &CustomerFundTransferDetail{
				ResultCode: rb.ResultCode,
				ResultDesc: rb.ResultDesc,
			},
		}, nil
	}

	detail := &CustomerFundTransferDetail{
		ConverstationIdentifier: env.Body.Result.Header.ConversationID,
		ResultCode:              rb.ResultCode,
		ResultDesc:              rb.ResultDesc,
	}
	if rb.TransactionResult != nil {
		detail.TransactionID = rb.TransactionResult.TransactionID
		detail.FTNumber = rb.TransactionResult.TransactionID
	}
	if rb.ReferenceData != nil {
		detail.ReferenceDetail = rb.ReferenceData.Details
	}

	return &CustomerFundTransferResult{
		Status: true,
		Detail: detail,
	}, nil
}

func escapeXML(s string) string {
	var b strings.Builder
	_ = xml.EscapeText(&b, []byte(s))
	return b.String()
}
