package customerlookup

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username           string
	Password           string
	CustomerIdentifier string
}

type CustomerLookupParam struct {
	CustomerIdentifier string
}

func NewCustomerLookup(param Params) string {
	return fmt.Sprintf(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
xmlns:cbes="http://temenos.com/CBESUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:CustomerInformationDetails>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <CUSTOMERINFOSUPERAPPType>
                <enquiryInputCollection>
                    <columnName>ID</columnName>
                    <criteriaValue>%s</criteriaValue>
                    <operand>EQ</operand>
                </enquiryInputCollection>
            </CUSTOMERINFOSUPERAPPType>
        </cbes:CustomerInformationDetails>
    </soapenv:Body>
</soapenv:Envelope>
`, param.Password, param.Username, param.CustomerIdentifier)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	CustomerInformationDetailsResponse CustomerInformationDetailsResponse `xml:"CustomerInformationDetailsResponse"`
}

type CustomerInformationDetailsResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
	} `xml:"Status"`
	CustomerInformation *struct {
		CustomerDetails struct {
			CustomerLookupDetail *CustomerLookupDetail `xml:"mCUSTOMERINFOSUPERAPPDetailType"`
		} `xml:"gCUSTOMERINFOSUPERAPPDetailType"`
	} `xml:"CUSTOMERINFOSUPERAPPType"`
}

type CustomerLookupDetail struct {
	FullName      string `xml:"FullName"`
	BirthDate     string `xml:"BirthDate"`
	Gender        string `xml:"Gender"`
	Address       string `xml:"Address"`
	PhoneNumber   string `xml:"PhoneNumber"`
	Email         string `xml:"Email"`
	City          string `xml:"City"`
	Nationality   string `xml:"Nationality"`
	MaritalStatus string `xml:"MaritalStatus"`
	PostalCode    string `xml:"PostalCode"`
	IDDocument    string `xml:"IDDocument"`
	Title         string `xml:"Title"`
	FirstName     string `xml:"FirstName"`
	LastName      string `xml:"LastName"`
}

type CustomerLookupResult struct {
	Success       bool
	CustomerInfos *CustomerLookupDetail
	Message       string
}

func ParseCustomerLookupSOAP(response string) (*CustomerLookupResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(response), &env)
	if err != nil {
		return nil, err
	}

	if env.Body.CustomerInformationDetailsResponse.CustomerInformation != nil {
		resp := env.Body.CustomerInformationDetailsResponse

		if resp.Status == nil {
			return &CustomerLookupResult{
				Success: false,
				Message: "Invalid response structure",
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &CustomerLookupResult{
				Success: false,
				Message: "Customer lookup failed",
			}, nil
		}

		if resp.CustomerInformation.CustomerDetails.CustomerLookupDetail == nil {
			return &CustomerLookupResult{
				Success: false,
				Message: "No customer details found",
			}, nil
		}

		return &CustomerLookupResult{
			Success: true,
			CustomerInfos: &CustomerLookupDetail{
				FullName:      resp.CustomerInformation.CustomerDetails.CustomerLookupDetail.FullName,
				Title:         resp.CustomerInformation.CustomerDetails.CustomerLookupDetail.Title,
				Address:       resp.CustomerInformation.CustomerDetails.CustomerLookupDetail.Address,
				PhoneNumber:   resp.CustomerInformation.CustomerDetails.CustomerLookupDetail.PhoneNumber,
				Email:         resp.CustomerInformation.CustomerDetails.CustomerLookupDetail.Email,
				City:          resp.CustomerInformation.CustomerDetails.CustomerLookupDetail.City,
				FirstName:     resp.CustomerInformation.CustomerDetails.CustomerLookupDetail.FirstName,
				LastName:      resp.CustomerInformation.CustomerDetails.CustomerLookupDetail.LastName,
				BirthDate:     resp.CustomerInformation.CustomerDetails.CustomerLookupDetail.BirthDate,
				Nationality:   resp.CustomerInformation.CustomerDetails.CustomerLookupDetail.Nationality,
				PostalCode:    resp.CustomerInformation.CustomerDetails.CustomerLookupDetail.PostalCode,
				IDDocument:    resp.CustomerInformation.CustomerDetails.CustomerLookupDetail.IDDocument,
				MaritalStatus: resp.CustomerInformation.CustomerDetails.CustomerLookupDetail.MaritalStatus,
			},
		}, nil
	}

	return &CustomerLookupResult{
		Success: false,
		Message: "Invalid response type",
	}, nil
}
