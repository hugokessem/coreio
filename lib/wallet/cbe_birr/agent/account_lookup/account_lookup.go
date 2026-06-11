// coreio target: lib/wallet/cbe_birr/agent/account_lookup/account_lookup.go
package accountlookup

import (
	"crypto/rand"
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

type Params struct {
	ThirdPartyIdentifier            string
	Password                        string
	InitiatorIdentifier             string // KYC third-party ID (was hardcoded "Anamail")
	SecurityCredential              string
	OriginalConverstationIdentifier string
	Timestamp                       string
	PhoneNumber                     string // agent code (field name kept for wallet/init.go compat)
}

type AgentAccountLookupParams struct {
	OriginalConverstationIdentifier string
	Timestamp                       string
	PhoneNumber                     string // agent code
}

// GenerateOriginatorConversationID matches platform/cbebirr generateUniqueOriginatorConversationID.
func GenerateOriginatorConversationID() string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	randomBytes := make([]byte, 4)
	rand.Read(randomBytes)
	return fmt.Sprintf("S_%d_%08x", timestamp, randomBytes)
}

func NewAgentAccountLookup(param Params) string {
	xe := escapeXML
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
                  <req:Identifier>%s</req:Identifier>
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
</soapenv:Envelope>`, xe(param.OriginalConverstationIdentifier), xe(param.ThirdPartyIdentifier), xe(param.Password), xe(param.Timestamp), xe(param.InitiatorIdentifier), xe(param.SecurityCredential), xe(param.PhoneNumber))
}

type OrganizationBasicData struct {
	ShortCode        string `xml:"http://cps.huawei.com/synccpsinterface/result ShortCode"`
	OrganizationName string `xml:"http://cps.huawei.com/synccpsinterface/result OrganizationName"`
}

type AccountLookupResponse struct {
	Version                         string
	OriginalConverstationIdentifier string
	ConversationIdentifier          string
	ResultCode                      string
	ResultDesc                      string
	OrganizationName                string
	OrganizationBasicData           OrganizationBasicData
}

type AgentAccountLookupResult struct {
	Success bool
	Detail  *AccountLookupResponse
	Message string
}

type envelope struct {
	Body struct {
		Result struct {
			Header struct {
				Version                  string `xml:"http://cps.huawei.com/synccpsinterface/result Version"`
				OriginatorConversationID string `xml:"http://cps.huawei.com/synccpsinterface/result OriginatorConversationID"`
				ConversationID           string `xml:"http://cps.huawei.com/synccpsinterface/result ConversationID"`
			} `xml:"http://cps.huawei.com/synccpsinterface/result Header"`
			ResultBody struct {
				ResultCode                  string `xml:"http://cps.huawei.com/synccpsinterface/result ResultCode"`
				ResultDesc                  string `xml:"http://cps.huawei.com/synccpsinterface/result ResultDesc"`
				QueryOrganizationInfoResult *struct {
					OrganizationBasicData *OrganizationBasicData `xml:"http://cps.huawei.com/synccpsinterface/result OrganizationBasicData"`
				} `xml:"http://cps.huawei.com/synccpsinterface/result QueryOrganizationInfoResult"`
			} `xml:"http://cps.huawei.com/synccpsinterface/result Body"`
		} `xml:"http://cps.huawei.com/synccpsinterface/api_requestmgr Result"`
	} `xml:"Body"`
}

func ParseAgentLookupSOAP(xmlData string) (*AgentAccountLookupResult, error) {
	var env envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	rb := env.Body.Result.ResultBody
	if rb.ResultCode != "0" {
		return &AgentAccountLookupResult{
			Success: false,
			Message: rb.ResultDesc,
			Detail: &AccountLookupResponse{
				ResultCode: rb.ResultCode,
				ResultDesc: rb.ResultDesc,
			},
		}, nil
	}

	org := rb.QueryOrganizationInfoResult
	if org == nil || org.OrganizationBasicData == nil {
		return &AgentAccountLookupResult{
			Success: false,
			Message: "API returned failure!",
		}, nil
	}

	return &AgentAccountLookupResult{
		Success: true,
		Detail: &AccountLookupResponse{
			Version:                         env.Body.Result.Header.Version,
			OriginalConverstationIdentifier: env.Body.Result.Header.OriginatorConversationID,
			ConversationIdentifier:          env.Body.Result.Header.ConversationID,
			ResultCode:                      rb.ResultCode,
			ResultDesc:                      rb.ResultDesc,
			OrganizationName:                strings.TrimSpace(org.OrganizationBasicData.OrganizationName),
			OrganizationBasicData:           *org.OrganizationBasicData,
		},
	}, nil
}

func escapeXML(s string) string {
	var b strings.Builder
	_ = xml.EscapeText(&b, []byte(s))
	return b.String()
}
