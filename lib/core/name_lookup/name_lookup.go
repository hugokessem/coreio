package namelookup

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

type NameLookupParam struct {
	AccountNumber string
}

func NewNameLookup(param Params) string {
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
	NameLookupResponse         *NameLookupResponse         `xml:"AccountLookupResponse"`
	NameBalanceInquiryResponse *NameBalanceInquiryResponse `xml:"AccountBalanceInquiryResponse"`
}

type NameLookupResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
	} `xml:"Status"`
	AccountEnquirySuperappType *struct {
		Group *struct {
			Details *NameLookupDetail `xml:"mACCOUNTENQUIRYSUPERAPPDetailType"`
		} `xml:"gACCOUNTENQUIRYSUPERAPPDetailType"`
	} `xml:"ACCOUNTENQUIRYSUPERAPPType"`
}

type NameLookupDetail struct {
	AccountNumber   string `xml:"AccountNumber"`
	CustomerName    string `xml:"CustomerName"`
	Restriction     string `xml:"Restriction"`
	Currency        string `xml:"Currency"`
	WorkingBalance  string `xml:"WorkingBalance"`
	CustomerID      string `xml:"CustomerID"`
	BranchName      string `xml:"BranchName"`
	Category        string `xml:"Category"`
	AccountType     string `xml:"AccountType"`
	PhoneNumber     string `xml:"PhoneNo"`
	BirthOfDate     string `xml:"DOB"`
	Gender          string `xml:"Gender"`
	BranchCode      string `xml:"BranchCode"`
	DistrictName    string `xml:"DistrictName"`
	Industry        string `xml:"Industry"`
	Sector          string `xml:"Sector"`
	Ownership       string `xml:"Ownership"`
	CustomerSegment string `xml:"CustomerSegment"`
	Target          string `xml:"Target"`
	TinNO           string `xml:"TinNO"`
	RestrictionType string `xml:"RestrictionType"`
	AccountName     string `xml:"AccountName"`
	SubSegment      string `xml:"SubSegment"`
	CustomerGroup   string `xml:"CustomerGroup"`
	AccountInactive string `xml:"InactiveFlag"`
	DAOCode         string `xml:"DAOCode"`
	DAOName         string `xml:"DAOName"`
}

// ----------------- Failure or no records -----------------
type NameBalanceInquiryResponse struct {
	Status *struct {
		SuccessIndicator string   `xml:"successIndicator"`
		Messages         []string `xml:"messages"`
	} `xml:"Status"`
}

// ----------------- Parser -----------------
type NameLookupResult struct {
	Success  bool
	Detail   *NameLookupDetail
	Messages []string
}

func ParseNameLookupSOAP(xmlData string) (*NameLookupResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}

	// Case 1: NameLookupResponse
	if env.Body.NameLookupResponse != nil {
		resp := env.Body.NameLookupResponse
		if resp.Status == nil {
			return &NameLookupResult{
				Success:  false,
				Messages: []string{"Missing Status"},
			}, nil
		}
		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &NameLookupResult{
				Success:  false,
				Messages: []string{"API returned failure"},
			}, nil
		}
		if resp.AccountEnquirySuperappType == nil ||
			resp.AccountEnquirySuperappType.Group == nil ||
			resp.AccountEnquirySuperappType.Group.Details == nil {
			return &NameLookupResult{
				Success:  false,
				Messages: []string{"No account details found"},
			}, nil
		}
		return &NameLookupResult{
			Success: true,
			Detail: &NameLookupDetail{
				AccountNumber:   resp.AccountEnquirySuperappType.Group.Details.AccountNumber,
				CustomerName:    resp.AccountEnquirySuperappType.Group.Details.CustomerName,
				Restriction:     resp.AccountEnquirySuperappType.Group.Details.Restriction,
				Currency:        resp.AccountEnquirySuperappType.Group.Details.Currency,
				WorkingBalance:  resp.AccountEnquirySuperappType.Group.Details.WorkingBalance,
				CustomerID:      resp.AccountEnquirySuperappType.Group.Details.CustomerID,
				AccountType:     resp.AccountEnquirySuperappType.Group.Details.AccountType,
				PhoneNumber:     resp.AccountEnquirySuperappType.Group.Details.PhoneNumber,
				BirthOfDate:     resp.AccountEnquirySuperappType.Group.Details.BirthOfDate,
				Gender:          resp.AccountEnquirySuperappType.Group.Details.Gender,
				RestrictionType: resp.AccountEnquirySuperappType.Group.Details.RestrictionType,
				BranchName:      resp.AccountEnquirySuperappType.Group.Details.BranchName,
				BranchCode:      resp.AccountEnquirySuperappType.Group.Details.BranchCode,
				DistrictName:    resp.AccountEnquirySuperappType.Group.Details.DistrictName,
				Industry:        resp.AccountEnquirySuperappType.Group.Details.Industry,
				Sector:          resp.AccountEnquirySuperappType.Group.Details.Sector,
				Ownership:       resp.AccountEnquirySuperappType.Group.Details.Ownership,
				CustomerSegment: resp.AccountEnquirySuperappType.Group.Details.CustomerSegment,
				Target:          resp.AccountEnquirySuperappType.Group.Details.Target,
				TinNO:           resp.AccountEnquirySuperappType.Group.Details.TinNO,
				Category:        resp.AccountEnquirySuperappType.Group.Details.Category,
				AccountName:     resp.AccountEnquirySuperappType.Group.Details.AccountName,
				SubSegment:      resp.AccountEnquirySuperappType.Group.Details.SubSegment,
				CustomerGroup:   resp.AccountEnquirySuperappType.Group.Details.CustomerGroup,
				AccountInactive: resp.AccountEnquirySuperappType.Group.Details.AccountInactive,
				DAOName:         resp.AccountEnquirySuperappType.Group.Details.DAOName,
				DAOCode:         resp.AccountEnquirySuperappType.Group.Details.DAOCode,
			},
		}, nil
	}

	// Case 2: NameBalanceInquiryResponse (failure / no records)
	if env.Body.NameBalanceInquiryResponse != nil {
		resp := env.Body.NameBalanceInquiryResponse
		messages := []string{}
		if resp.Status != nil && len(resp.Status.Messages) > 0 {
			messages = resp.Status.Messages
		}
		success := false
		if resp.Status != nil && strings.ToLower(resp.Status.SuccessIndicator) == "success" {
			success = true
		}
		return &NameLookupResult{
			Success:  success,
			Messages: messages,
		}, nil
	}

	// Unknown response
	return &NameLookupResult{
		Success:  false,
		Messages: []string{"Invalid response type"},
	}, nil
}
