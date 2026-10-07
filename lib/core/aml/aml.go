package aml

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

type AMLParam struct {
	AccountNumber string
}

func NewAMLLookup(param Params) string {
	return fmt.Sprintf(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
   <soapenv:Header/>
   <soapenv:Body>
      <cbes:AccountNameLookup>
         <WebRequestCommon>
            <company/>
            <password>%s</password>
            <userName>%s</userName>
         </WebRequestCommon>
         <NAMELOOKUPSUPERAPPType>
            <enquiryInputCollection>
               <columnName>ID</columnName>
               <criteriaValue>%s</criteriaValue>
               <operand>EQ</operand>
            </enquiryInputCollection>
         </NAMELOOKUPSUPERAPPType>
      </cbes:AccountNameLookup>
   </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.AccountNumber)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	AccountNameLookupResponse *AccountNameLookupResponse `xml:"AccountNameLookupResponse"`
}

type AccountNameLookupResponse struct {
	Status *struct {
		SuccessIndicator string   `xml:"successIndicator"`
		Messages         []string `xml:"messages"`
	} `xml:"Status"`
	NameLookupSuperAppType *struct {
		Group *struct {
			Detail *AMLLookupDetail `xml:"mNAMELOOKUPSUPERAPPDetailType"`
		} `xml:"gNAMELOOKUPSUPERAPPDetailType"`
	} `xml:"NAMELOOKUPSUPERAPPType"`
}

type AMLLookupDetail struct {
	AccountNumber   string `xml:"AccountNumber"`
	AccountName     string `xml:"AccountName"`
	Currency        string `xml:"Currency"`
	RestrictionType string `xml:"RestrictionType"`
	InactiveFlag    string `xml:"InactiveFlag"`
	CustomerNumber  string `xml:"CIF"`
	Industry        string `xml:"Industry"`
	CustomerSegment string `xml:"CustomerSegment"`
	SubSegment      string `xml:"SubSegment"`
	CustomerGroup   string `xml:"CustomerGroup"`
	Category        string `xml:"Category"`
	Sector          string `xml:"Sector"`
	TinNumber       string `xml:"TINNo"`
}

type AMLLookupResult struct {
	Success  bool
	Detail   *AMLLookupDetail
	Messages []string
}

func ParseAMLLookupSOAP(xmlData string) (*AMLLookupResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.AccountNameLookupResponse == nil {
		return &AMLLookupResult{
			Success:  false,
			Messages: []string{"Invalid response type"},
		}, nil
	}

	resp := env.Body.AccountNameLookupResponse
	if resp.Status == nil {
		return &AMLLookupResult{
			Success:  false,
			Messages: []string{"Missing Status"},
		}, nil
	}

	if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
		messages := resp.Status.Messages
		if len(messages) == 0 {
			messages = []string{"AML lookup failed"}
		}
		return &AMLLookupResult{
			Success:  false,
			Messages: messages,
		}, nil
	}

	if resp.NameLookupSuperAppType == nil ||
		resp.NameLookupSuperAppType.Group == nil ||
		resp.NameLookupSuperAppType.Group.Detail == nil {
		messages := resp.Status.Messages
		if len(messages) == 0 {
			messages = []string{"No account details found"}
		}
		return &AMLLookupResult{
			Success:  false,
			Messages: messages,
		}, nil
	}

	detail := resp.NameLookupSuperAppType.Group.Detail
	return &AMLLookupResult{
		Success: true,
		Detail: &AMLLookupDetail{
			AccountNumber:   detail.AccountNumber,
			AccountName:     detail.AccountName,
			Currency:        detail.Currency,
			RestrictionType: detail.RestrictionType,
			InactiveFlag:    detail.InactiveFlag,
			CustomerNumber:  detail.CustomerNumber,
			Industry:        detail.Industry,
			CustomerSegment: detail.CustomerSegment,
			SubSegment:      detail.SubSegment,
			CustomerGroup:   detail.CustomerGroup,
			Category:        detail.Category,
			Sector:          detail.Sector,
			TinNumber:       detail.TinNumber,
		},
	}, nil
}
