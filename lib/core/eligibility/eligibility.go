package eligibility

import (
	"encoding/xml"
	"fmt"
	"strings"

	valueobject "github.com/hugokessem/coreio/lib/core/eligibility/value_object"
)

type Param struct {
	Username      string
	Password      string
	FetchBy       string
	CriticalValue string
}

type EligibilityParam struct {
	FetchBy       string
	CriticalValue string
}

func NewEligibility(param Param) string {
	var accountNumber, customerNumber string
	if valueobject.FetchBy(param.FetchBy).IsValid() {
		if valueobject.FetchByAccountNumber.Equal(param.FetchBy) {
			accountNumber = param.CriticalValue
		} else {
			customerNumber = param.CriticalValue
		}
	}
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
   <soapenv:Header/>
   <soapenv:Body>
      <cbes:CustomerEligibilityforOnboarding>
         <WebRequestCommon>
            <company/>
            <password>%s</password>
            <userName>%s</userName>
         </WebRequestCommon>
         <ACCOUNTINFOSUPERAPPRESTRICTType>
            <enquiryInputCollection>
               <columnName>CUS.ID</columnName>
               <criteriaValue>%s</criteriaValue>
               <operand>EQ</operand>
            </enquiryInputCollection>
            <enquiryInputCollection>
               <columnName>ACCT.ID</columnName>
               <criteriaValue>%s</criteriaValue>
               <operand>EQ</operand>
            </enquiryInputCollection>
         </ACCOUNTINFOSUPERAPPRESTRICTType>
      </cbes:CustomerEligibilityforOnboarding>
   </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, customerNumber, accountNumber)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	CustomerEligibilityforOnboardingResponse *CustomerEligibilityforOnboardingResponse `xml:"CustomerEligibilityforOnboardingResponse"`
}

type CustomerEligibilityforOnboardingResponse struct {
	Status *struct {
		SuccessIndicator string   `xml:"successIndicator"`
		Messages         []string `xml:"messages"`
	} `xml:"Status"`
	ACCOUNTINFOSUPERAPPRESTRICTType *struct {
		Group *struct {
			Details []CustomerEligibilityDetail `xml:"mACCOUNTINFOSUPERAPPRESTRICTDetailType"`
		} `xml:"gACCOUNTINFOSUPERAPPRESTRICTDetailType"`
	} `xml:"ACCOUNTINFOSUPERAPPRESTRICTType"`
}

type CustomerEligibilityDetail struct {
	AccountNumber   string `xml:"AccountNumber"`
	CustomerName    string `xml:"CustomerName"`
	Restriction     string `xml:"Restriction"`
	Currency        string `xml:"Currency"`
	CustomerID      string `xml:"CustomerID"`
	Category        string `xml:"Category"`
	AccountType     string `xml:"AccountType"`
	BranchCode      string `xml:"BranchCode"`
	BranchName      string `xml:"BranchName"`
	DistrictName    string `xml:"DistrictName"`
	PhoneNo         string `xml:"PhoneNo"`
	Industry        string `xml:"Industry"`
	Sector          string `xml:"Sector"`
	Ownership       string `xml:"Ownership"`
	CustomerSegment string `xml:"CustomerSegment"`
	Target          string `xml:"Target"`
	Gender          string `xml:"Gender"`
	DateofBirth     string `xml:"DateofBirth"`
	Email           string `xml:"Email"`
	RestrictionType string `xml:"RestrictionType"`
}

type CustomerEligibilityResult struct {
	Success  bool
	Details  []CustomerEligibilityDetail
	Messages []string
}

func ParseCustomerEligibilitySOAP(response string) (*CustomerEligibilityResult, error) {
	var envelope Envelope
	err := xml.Unmarshal([]byte(response), &envelope)
	if err != nil {
		return nil, err
	}

	if envelope.Body.CustomerEligibilityforOnboardingResponse != nil {
		resp := envelope.Body.CustomerEligibilityforOnboardingResponse
		if resp.Status == nil {
			return &CustomerEligibilityResult{
				Success:  false,
				Details:  nil,
				Messages: []string{"Missing status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &CustomerEligibilityResult{
				Success:  false,
				Details:  nil,
				Messages: resp.Status.Messages,
			}, nil
		}

		if resp.ACCOUNTINFOSUPERAPPRESTRICTType == nil ||
			resp.ACCOUNTINFOSUPERAPPRESTRICTType.Group == nil ||
			len(resp.ACCOUNTINFOSUPERAPPRESTRICTType.Group.Details) == 0 {
			return &CustomerEligibilityResult{
				Success:  false,
				Details:  nil,
				Messages: []string{"Missing details"},
			}, nil
		}

		return &CustomerEligibilityResult{
			Success:  true,
			Details:  resp.ACCOUNTINFOSUPERAPPRESTRICTType.Group.Details,
			Messages: resp.Status.Messages,
		}, nil
	}

	return &CustomerEligibilityResult{
		Success:  false,
		Details:  nil,
		Messages: []string{"Failed to parse eligibility response"},
	}, nil
}
