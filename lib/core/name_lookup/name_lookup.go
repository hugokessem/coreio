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
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
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
	NameLookupDetailResult *struct {
		NameLookup *struct {
			NameLookupResultDetail *NameLookupResultDetail `xml:"mNAMELOOKUPSUPERAPPDetailType"`
		} `xml:"gNAMELOOKUPSUPERAPPDetailType"`
	} `xml:"NAMELOOKUPSUPERAPPType"`
}

type NameLookupResultDetail struct {
	AccountNumber   string `xml:"AccountNumber"`
	AccountName     string `xml:"AccountName"`
	Currency        string `xml:"Currency"`
	RestrictionType string `xml:"RestrictionType"`
	InactiveFlag    string `xml:"InactiveFlag"`
}

type NameLookupResult struct {
	Success  bool
	Detail   *NameLookupResultDetail
	Messages []string
}

func ParseNameLookupSOAP(xmlData string) (*NameLookupResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}

	if env.Body.AccountNameLookupResponse != nil {
		resp := env.Body.AccountNameLookupResponse
		if resp.Status == nil {
			return &NameLookupResult{
				Success:  true,
				Messages: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &NameLookupResult{
				Success:  false,
				Detail:   nil,
				Messages: []string{"API returned failure"},
			}, nil
		}

		if resp.NameLookupDetailResult.NameLookup.NameLookupResultDetail == nil {
			return &NameLookupResult{
				Success:  false,
				Detail:   nil,
				Messages: resp.Status.Messages,
			}, nil
		}

		return &NameLookupResult{
			Success: true,
			Detail: &NameLookupResultDetail{
				AccountNumber:   resp.NameLookupDetailResult.NameLookup.NameLookupResultDetail.AccountNumber,
				AccountName:     resp.NameLookupDetailResult.NameLookup.NameLookupResultDetail.AccountName,
				Currency:        resp.NameLookupDetailResult.NameLookup.NameLookupResultDetail.Currency,
				RestrictionType: resp.NameLookupDetailResult.NameLookup.NameLookupResultDetail.RestrictionType,
				InactiveFlag:    resp.NameLookupDetailResult.NameLookup.NameLookupResultDetail.InactiveFlag,
			},
		}, nil

	}
	return &NameLookupResult{
		Success:  false,
		Detail:   nil,
		Messages: []string{"Invalid response type"},
	}, nil
}
