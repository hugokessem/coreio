package accountlookup

import (
	"encoding/xml"
	"fmt"
)

type Params struct {
	OriginalConverstationIdentifier string
	ThirdPartyIdentifier            string
	Password                        string
	Timestamp                       string
	SecurityCredential              string
	PhoneNumber                     string
}
type CustomerAccountLookupParams struct {
	OriginalConverstationIdentifier string
	Timestamp                       string
	PhoneNumber                     string
}

func NewCustomerAccountLookup(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:api="http://cps.huawei.com/synccpsinterface/api_requestmgr" xmlns:req="http://cps.huawei.com/synccpsinterface/request" xmlns:com="http://cps.huawei.com/synccpsinterface/common" xmlns:cus="http://cps.huawei.com/cpsinterface/customizedrequest">
   <soapenv:Header/>
   <soapenv:Body>
      <api:Request>
         <req:Header>
            <req:Version>1.0</req:Version>
            <req:CommandID>QueryCustomerKYC</req:CommandID>
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
               <req:ReceiverParty>
                  <req:IdentifierType>1</req:IdentifierType>
                  <req:Identifier>%s</req:Identifier>
               </req:ReceiverParty>
            </req:Identity>
            <req:QueryCustomerKYCRequest/>
         </req:Body>
      </api:Request>
   </soapenv:Body>
</soapenv:Envelope>`, param.OriginalConverstationIdentifier, param.ThirdPartyIdentifier, param.Password, param.Timestamp, param.SecurityCredential, param.PhoneNumber)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	Result *struct {
		Header     *Header     `xml:"Header"`
		ResultBody *ResultBody `xml:"Body"`
	} `xml:"Result"`
}

type Header struct {
	Version                         string `xml:"Version"`
	OriginalConverstationIdentifier string `xml:"OriginatorConversationID"`
	ConversationIdentifier          string `xml:"ConversationID"`
}

type ResultBody struct {
	ResultType           string           `xml:"ResultType"`
	ResultCode           string           `xml:"ResultCode"`
	ResultDescription    string           `xml:"ResultDesc"`
	QueryCustomerKYCDate *CustomerKYCData `xml:"QueryCustomerKYCResult"`
}

type CustomerKYCData struct {
	BOCompletedTime string     `xml:"BOCompletedTime"`
	SimpleKYCData   []KycField `xml:"SimpleKYCData"`
	IDDetailsData   *struct {
		IDRecord *struct {
			IDTypeValue  string `xml:"IDTypeValue"`
			IDNumber     string `xml:"IDNumber"`
			IDExpiryDate string `xml:"IDExpiryDate"`
			IssuedBy     string `xml:"IssuedBy"`
		} `xml:"IDRecord"`
	} `xml:"IDDetailsData"`
}

type KycField struct {
	KYCName  string `xml:"KYCName"`
	KYCValue string `xml:"KYCValue"`
}

type CustomerAccountLookupResponse struct {
	Version                         string
	OriginalConverstationIdentifier string
	ConversationIdentifier          string
	CustomerKYCData                 CustomerKYCData
}
type CustomerAccountLookupResult struct {
	Success bool
	Detail  *CustomerAccountLookupResponse
	Message string
}

func ParseCustomerLookupSOAP(xmlData string) (*CustomerAccountLookupResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}

	if env.Body.Result.Header != nil && env.Body.Result.ResultBody != nil {
		resp := env.Body.Result
		if resp.ResultBody.ResultCode != "0" {
			// 1001 - Credential Error
			// 1003 - Duplicate Converstation ID
			return &CustomerAccountLookupResult{
				Success: false,
				Message: resp.ResultBody.ResultDescription,
			}, nil
		}

		if resp.ResultBody.QueryCustomerKYCDate == nil {
			return &CustomerAccountLookupResult{
				Success: false,
				Message: "API returned failure!",
			}, nil
		}

		return &CustomerAccountLookupResult{
			Success: true,
			Detail: &CustomerAccountLookupResponse{
				Version:                         resp.Header.Version,
				OriginalConverstationIdentifier: resp.Header.OriginalConverstationIdentifier,
				ConversationIdentifier:          resp.Header.ConversationIdentifier,
				CustomerKYCData:                 *resp.ResultBody.QueryCustomerKYCDate,
			},
		}, nil
	}

	return &CustomerAccountLookupResult{
		Success: false,
		Message: "invalid request!",
	}, nil
}
