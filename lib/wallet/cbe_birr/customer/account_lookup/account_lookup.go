// coreio target: lib/wallet/cbe_birr/customer/account_lookup/account_lookup.go
package accountlookup

import (
	"crypto/rand"
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

type Params struct {
	OriginalConverstationIdentifier string
	ThirdPartyIdentifier            string
	Password                        string
	Timestamp                       string
	InitiatorIdentifier             string // KYC third-party ID (was hardcoded "Anamail")
	SecurityCredential              string
	PhoneNumber                     string
}

type CustomerAccountLookupParams struct {
	OriginalConverstationIdentifier string
	Timestamp                       string
	PhoneNumber                     string
}

func NewCustomerAccountLookup(param Params) string {
	xe := escapeXML
	return fmt.Sprintf(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:api="http://cps.huawei.com/synccpsinterface/api_requestmgr" xmlns:req="http://cps.huawei.com/synccpsinterface/request" xmlns:com="http://cps.huawei.com/synccpsinterface/common" xmlns:cus="http://cps.huawei.com/cpsinterface/customizedrequest">
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
                  <req:Identifier>%s</req:Identifier>
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
</soapenv:Envelope>`, xe(param.OriginalConverstationIdentifier), xe(param.ThirdPartyIdentifier), xe(param.Password), xe(param.Timestamp), xe(param.InitiatorIdentifier), xe(param.SecurityCredential), xe(param.PhoneNumber))
}

type CustomerAccountLookupResponse struct {
	Version                         string
	OriginalConverstationIdentifier string
	ConversationIdentifier          string
	ResultCode                      string
	ResultDesc                      string
	FullName                        string
	FirstName                       string
	MiddleName                      string
	LastName                        string
	PhoneNumber                     string
	AccountStatus                   string
}

type CustomerAccountLookupResult struct {
	Success bool
	Detail  *CustomerAccountLookupResponse
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
				ResultCode             string `xml:"http://cps.huawei.com/synccpsinterface/result ResultCode"`
				ResultDesc             string `xml:"http://cps.huawei.com/synccpsinterface/result ResultDesc"`
				QueryCustomerKYCResult *struct {
					SimpleKYCData *struct {
						KYCFields []struct {
							KYCName  string `xml:"http://cps.huawei.com/synccpsinterface/common KYCName"`
							KYCValue string `xml:"http://cps.huawei.com/synccpsinterface/common KYCValue"`
						} `xml:"http://cps.huawei.com/synccpsinterface/common KYCField"`
					} `xml:"http://cps.huawei.com/synccpsinterface/result SimpleKYCData"`
				} `xml:"http://cps.huawei.com/synccpsinterface/result QueryCustomerKYCResult"`
			} `xml:"http://cps.huawei.com/synccpsinterface/result Body"`
		} `xml:"http://cps.huawei.com/synccpsinterface/api_requestmgr Result"`
	} `xml:"Body"`
}

// GenerateOriginatorConversationID matches platform/cbebirr generateUniqueOriginatorConversationID.
func GenerateOriginatorConversationID() string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	randomBytes := make([]byte, 4)
	rand.Read(randomBytes)
	return fmt.Sprintf("S_%d_%08x", timestamp, randomBytes)
}

// ParseCustomerLookupSOAP parses SOAP response. phoneNumber is the request phone (echoed in Detail.PhoneNumber).
func ParseCustomerLookupSOAP(xmlData, phoneNumber string) (*CustomerAccountLookupResult, error) {
	var env envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	rb := env.Body.Result.ResultBody
	if rb.ResultCode != "0" {
		return &CustomerAccountLookupResult{
			Success: false,
			Message: rb.ResultDesc,
			Detail: &CustomerAccountLookupResponse{
				ResultCode: rb.ResultCode,
				ResultDesc: rb.ResultDesc,
			},
		}, nil
	}

	if rb.QueryCustomerKYCResult == nil || rb.QueryCustomerKYCResult.SimpleKYCData == nil {
		return &CustomerAccountLookupResult{
			Success: false,
			Message: "API returned failure!",
		}, nil
	}

	detail := &CustomerAccountLookupResponse{
		Version:                         env.Body.Result.Header.Version,
		OriginalConverstationIdentifier: env.Body.Result.Header.OriginatorConversationID,
		ConversationIdentifier:          env.Body.Result.Header.ConversationID,
		ResultCode:                      rb.ResultCode,
		ResultDesc:                      rb.ResultDesc,
		PhoneNumber:                     strings.TrimSpace(phoneNumber),
	}

	var firstName, secondName, lastName string
	for _, kycField := range rb.QueryCustomerKYCResult.SimpleKYCData.KYCFields {
		kycName := strings.TrimSpace(kycField.KYCName)
		kycValue := strings.TrimSpace(kycField.KYCValue)
		if kycValue == "" || kycValue == "true" {
			continue
		}
		switch kycName {
		case "[KYC][Personal Details][First Name]":
			firstName = kycValue
			detail.FirstName = kycValue
		case "[KYC][Personal Details][Second Name]":
			secondName = kycValue
			detail.MiddleName = kycValue
		case "[KYC][Personal Details][Last Name]":
			lastName = kycValue
			detail.LastName = kycValue
		default:
			lower := strings.ToLower(kycName)
			if detail.PhoneNumber == "" && (strings.Contains(lower, "phone") || strings.Contains(lower, "msisdn") || strings.Contains(lower, "mobile")) {
				detail.PhoneNumber = kycValue
			}
		}
		if detail.AccountStatus == "" {
			lower := strings.ToLower(kycName)
			if strings.Contains(lower, "status") ||
				strings.Contains(lower, "[account status]") ||
				strings.Contains(lower, "[account_status]") {
				detail.AccountStatus = kycValue
			}
		}
	}

	parts := []string{}
	for _, p := range []string{firstName, secondName, lastName} {
		if p != "" {
			parts = append(parts, p)
		}
	}
	if len(parts) > 0 {
		detail.FullName = strings.Join(parts, " ")
	}

	if isAccountStatusBlocked(detail.AccountStatus) {
		return &CustomerAccountLookupResult{
			Success: false,
			Message: detail.AccountStatus,
			Detail:  detail,
		}, nil
	}

	return &CustomerAccountLookupResult{
		Success: true,
		Detail:  detail,
	}, nil
}

func isAccountStatusBlocked(statusCode string) bool {
	switch strings.TrimSpace(statusCode) {
	case "-", ",00", "00", ",05", "05", ",06", "06", ",07", "07":
		return true
	default:
		return false
	}
}

func escapeXML(s string) string {
	var b strings.Builder
	_ = xml.EscapeText(&b, []byte(s))
	return b.String()
}
