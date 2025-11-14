package accountlookup

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username      string
	Password      string
	AccountNumber string
}

type AccountLookupParam struct {
	AccountNumber string
}

func NewAccountLookup(param Params) string {
	return fmt.Sprintf(
		`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
   <soapenv:Header/>
   <soapenv:Body>
      <cbes:AccountLookup>
         <WebRequestCommon>
            <company/>
            <password>%s</password>
            <userName>%s</userName>
         </WebRequestCommon>
         <ACCOUNTENQUIRYSUPERAPPType>
            <enquiryInputCollection>
               <columnName>ID</columnName>
               <criteriaValue>%s</criteriaValue>
               <operand>EQ</operand>
            </enquiryInputCollection>
         </ACCOUNTENQUIRYSUPERAPPType>
      </cbes:AccountLookup>
   </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.AccountNumber)
}

// ----------------- Generic SOAP Envelope -----------------
type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	AccountLookupResponse         *AccountLookupResponse         `xml:"AccountLookupResponse"`
	AccountBalanceInquiryResponse *AccountBalanceInquiryResponse `xml:"AccountBalanceInquiryResponse"`
}

// ----------------- Account Lookup Success -----------------
type AccountLookupResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
	} `xml:"Status"`
	AccountEnquirySuperappType *struct {
		Group *struct {
			Details *AccountLookupDetail `xml:"mACCOUNTENQUIRYSUPERAPPDetailType"`
		} `xml:"gACCOUNTENQUIRYSUPERAPPDetailType"`
	} `xml:"ACCOUNTENQUIRYSUPERAPPType"`
}

type AccountLookupDetail struct {
	AccountNumber  string `xml:"AccountNumber"`
	CustomerName   string `xml:"CustomerName"`
	Restriction    string `xml:"Restriction"`
	Currency       string `xml:"Currency"`
	WorkingBalance string `xml:"WorkingBalance"`
	CustomerID     string `xml:"CustomerID"`
	AccountType    string `xml:"AccountType"`
}

// ----------------- Failure or no records -----------------
type AccountBalanceInquiryResponse struct {
	Status *struct {
		SuccessIndicator string   `xml:"successIndicator"`
		Messages         []string `xml:"messages"`
	} `xml:"Status"`
}

// ----------------- Parser -----------------
type AccountLookupResult struct {
	Success  bool
	Detail   *AccountLookupDetail
	Messages []string
}

func ParseAccountLookupSOAP(xmlData string) (*AccountLookupResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}

	// Case 1: AccountLookupResponse
	if env.Body.AccountLookupResponse != nil {
		resp := env.Body.AccountLookupResponse
		if resp.Status == nil {
			return &AccountLookupResult{
				Success:  false,
				Messages: []string{"Missing Status"},
			}, nil
		}
		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &AccountLookupResult{
				Success:  false,
				Messages: []string{"API returned failure"},
			}, nil
		}
		if resp.AccountEnquirySuperappType == nil ||
			resp.AccountEnquirySuperappType.Group == nil ||
			resp.AccountEnquirySuperappType.Group.Details == nil {
			return &AccountLookupResult{
				Success:  false,
				Messages: []string{"No account details found"},
			}, nil
		}
		return &AccountLookupResult{
			Success: true,
			Detail: &AccountLookupDetail{
				AccountNumber:  resp.AccountEnquirySuperappType.Group.Details.AccountNumber,
				CustomerName:   resp.AccountEnquirySuperappType.Group.Details.CustomerName,
				Restriction:    resp.AccountEnquirySuperappType.Group.Details.Restriction,
				Currency:       resp.AccountEnquirySuperappType.Group.Details.Currency,
				WorkingBalance: resp.AccountEnquirySuperappType.Group.Details.WorkingBalance,
				CustomerID:     resp.AccountEnquirySuperappType.Group.Details.CustomerID,
				AccountType:    resp.AccountEnquirySuperappType.Group.Details.AccountType,
			},
		}, nil
	}

	// Case 2: AccountBalanceInquiryResponse (failure / no records)
	if env.Body.AccountBalanceInquiryResponse != nil {
		resp := env.Body.AccountBalanceInquiryResponse
		messages := []string{}
		if resp.Status != nil && len(resp.Status.Messages) > 0 {
			messages = resp.Status.Messages
		}
		success := false
		if resp.Status != nil && strings.ToLower(resp.Status.SuccessIndicator) == "success" {
			success = true
		}
		return &AccountLookupResult{
			Success:  success,
			Messages: messages,
		}, nil
	}

	// Unknown response
	return &AccountLookupResult{
		Success:  false,
		Messages: []string{"Invalid response type"},
	}, nil
}
