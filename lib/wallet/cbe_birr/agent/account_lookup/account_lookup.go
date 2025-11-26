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

type AgentAccountLookupParams struct {
	OriginalConverstationIdentifier string
	Timestamp                       string
	PhoneNumber                     string
}

func NewAgentAccountLookup(param Params) string {
	return fmt.Sprintf(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:api="http://cps.huawei.com/synccpsinterface/api_requestmgr" xmlns:req="http://cps.huawei.com/synccpsinterface/request" xmlns:com="http://cps.huawei.com/synccpsinterface/common" xmlns:cus="http://cps.huawei.com/cpsinterface/customizedrequest">
   <soapenv:Header/>
   <soapenv:Body>
      <api:Request>
         <req:Header>
            <req:Version>1.0</req:Version>
            <req:CommandID>QueryOrganizationInfo</req:CommandID>
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
                  <req:IdentifierType>4</req:IdentifierType>
                  <req:Identifier>%s</req:Identifier>
               </req:ReceiverParty>
            </req:Identity>
            <req:QueryOrganizationInfoRequest/>
            <req:Remark>query</req:Remark>
         </req:Body>
      </api:Request>
   </soapenv:Body>
</soapenv:Envelope>
`, param.OriginalConverstationIdentifier, param.ThirdPartyIdentifier, param.Password, param.Timestamp, param.SecurityCredential, param.PhoneNumber)
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
	ResultType            string                 `xml:"ResultType"`
	ResultCode            string                 `xml:"ResultCode"`
	ResultDescription     string                 `xml:"ResultDesc"`
	QueryOrganizationInfo *OrganizationBasicData `xml:"QueryOrganizationInfoResult"`
}

type OrganizationBasicData struct {
	BOCompletedTime       string `xml:"BOCompletedTime"`
	OrganizationBasicData *struct {
		ShortCode           string `xml:"ShortCode"`
		OrganizationName    string `xml:"OrganizationName"`
		IdentityStatus      string `xml:"IdentityStatus"`
		CreationDate        string `xml:"CreationDate"`
		TrustLevel          string `xml:"TrustLevel"`
		TrustLevelName      string `xml:"TrustLevelName"`
		RuleProfileID       string `xml:"RuleProfileID"`
		RuleProfileName     string `xml:"RuleProfileName"`
		ChargeProfileID     string `xml:"ChargeProfileID"`
		ChargeProfileName   string `xml:"ChargeProfileName"`
		AggregatorAcctModel string `xml:"AggregatorAcctModel"`
		HierarchyLevel      string `xml:"HierarchyLevel"`
		HierarchyModel      string `xml:"HierarchyModel"`
	} `xml:"OrganizationBasicData"`
}

type AccountLookupResponse struct {
	Version                         string
	OriginalConverstationIdentifier string
	ConversationIdentifier          string
	OrganizationBasicData           OrganizationBasicData
}
type AgentAccountLookupResult struct {
	Success bool
	Detail  *AccountLookupResponse
	Message string
}

func ParseAgentLookupSOAP(xmlData string) (*AgentAccountLookupResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}

	if env.Body.Result != nil && env.Body.Result.Header != nil && env.Body.Result.ResultBody != nil {
		resp := env.Body.Result
		if resp.ResultBody.ResultCode != "0" {
			// 1001 - Credential Error
			// 1003 - Duplicate Converstation ID
			return &AgentAccountLookupResult{
				Success: false,
				Message: resp.ResultBody.ResultDescription,
			}, nil
		}

		if resp.ResultBody.QueryOrganizationInfo == nil {
			return &AgentAccountLookupResult{
				Success: false,
				Message: "Invalid Request!",
			}, nil
		}

		if resp.ResultBody.QueryOrganizationInfo.OrganizationBasicData == nil {
			return &AgentAccountLookupResult{
				Success: false,
				Message: "Invalid Request! Missing OrganizationBasicData",
			}, nil
		}

		return &AgentAccountLookupResult{
			Success: true,
			Detail: &AccountLookupResponse{
				Version:                         resp.Header.Version,
				OriginalConverstationIdentifier: resp.Header.OriginalConverstationIdentifier,
				ConversationIdentifier:          resp.Header.ConversationIdentifier,
				OrganizationBasicData:           *resp.ResultBody.QueryOrganizationInfo,
			},
		}, nil
	}

	return &AgentAccountLookupResult{
		Success: false,
		Message: "invalid request!",
	}, nil
}
