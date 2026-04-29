package customerdetail

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username       string
	Password       string
	CustomerNumber string
}

type CustomerDetailParam struct {
	CustomerNumber string
}

func NewCustomerDetail(param Params) string {
	return fmt.Sprintf(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:CustomerLookup>
            <WebRequestCommon>
                <company></company>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <CUSTOMERLOOKUPSUPERAPPType>
                <enquiryInputCollection>
                    <columnName>@ID</columnName>
                    <criteriaValue>%s</criteriaValue>
                    <operand>EQ</operand>
                </enquiryInputCollection>
            </CUSTOMERLOOKUPSUPERAPPType>
        </cbes:CustomerLookup>
    </soapenv:Body>
</soapenv:Envelope>
	`, param.Password, param.Username, param.CustomerNumber)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	CustomerLookupResponse CustomerLookupResponse `xml:"CustomerLookupResponse"`
}

type CustomerLookupResponse struct {
	Status *struct {
		SuccessIndicator string   `xml:"successIndicator"`
		Messages         []string `xml:"messages"`
	} `xml:"Status"`
	CustomerLookupInformation *struct {
		GlobalCustomerDetails struct {
			CustomerDetail *CustomerDetail `xml:"mCUSTOMERLOOKUPSUPERAPPDetailType"`
		} `xml:"gCUSTOMERLOOKUPSUPERAPPDetailType"`
	} `xml:"CUSTOMERLOOKUPSUPERAPPType"`
}

type CustomerDetail struct {
	CustomerID         string `xml:"CustomerID"`
	BranchCode         string `xml:"BranchCode"`
	BranchName         string `xml:"BranchName"`
	Restriction        string `xml:"Restriction"`
	CustomerName       string `xml:"CustomerName"`
	Gender             string `xml:"Gender"`
	DateOfBirth        string `xml:"DOB"`
	Email              string `xml:"EMail"`
	CustomerPhone      string `xml:"CustomerPhone"`
	CustomerGroup      string `xml:"CustomerGroup"`
	CustomerSegment    string `xml:"CustomerSegment"`
	CustomerSubSegment string `xml:"SubSegement"`
}

type CustomerDetailResult struct {
	Success       bool
	CustomerInfos *CustomerDetail
	Message       []string
}

func ParseCustomerDetailSOAP(response string) (*CustomerDetailResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(response), &env)
	if err != nil {
		return nil, err
	}

	if env.Body.CustomerLookupResponse.CustomerLookupInformation != nil {
		resp := env.Body.CustomerLookupResponse

		if resp.Status == nil {
			return &CustomerDetailResult{
				Success: false,
				Message: []string{"Invalid response structure"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &CustomerDetailResult{
				Success: false,
				Message: resp.Status.Messages,
			}, nil
		}

		if resp.CustomerLookupInformation.GlobalCustomerDetails.CustomerDetail == nil {
			return &CustomerDetailResult{
				Success: false,
				Message: resp.Status.Messages,
			}, nil
		}

		return &CustomerDetailResult{
			Success:       true,
			CustomerInfos: resp.CustomerLookupInformation.GlobalCustomerDetails.CustomerDetail,
		}, nil
	}

	return &CustomerDetailResult{
		Success: false,
		Message: []string{"Invalid response type"},
	}, nil
}
