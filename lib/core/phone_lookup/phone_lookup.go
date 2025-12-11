package phonelookup

import (
	"encoding/xml"
	"fmt"
)

type Params struct {
	Username    string
	Password    string
	PhoneNumber string
}

type PhoneLookupParam struct {
	PhoneNumber string
}

func NewPhoneLookup(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
xmlns:cbes="http://temenos.com/CBESUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:GetCustomerPhoneNo>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <GETPHONECUSTOMERType>
                <enquiryInputCollection>
                    <columnName>MNEMONIC</columnName>
                    <criteriaValue>%s</criteriaValue>
                    <operand>EQ</operand>
                </enquiryInputCollection>
            </GETPHONECUSTOMERType>
        </cbes:GetCustomerPhoneNo>
    </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.PhoneNumber)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	GetCustomerPhoneNoResponse *GetCustomerPhoneNoResponse `xml:"GetCustomerPhoneNoResponse"`
}

type GetCustomerPhoneNoResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
	} `xml:"Status"`
	GetCByPhoneNumberType *struct {
		Group *struct {
			Details *GetCustomerPhoneNoDetail `xml:"mGETPHONECUSTOMERDetailType"`
		} `xml:"gGETPHONECUSTOMERDetailType"`
	} `xml:"GETPHONECUSTOMERType"`
}

type GetCustomerPhoneNoDetail struct {
	CustomerID  string `xml:"CustomerID"`
	PhoneNumber string `xml:"PhoneNumber"`
	Email       string `xml:"Email"`
	FullName    string `xml:"FullName"`
}

type PhoneLookupResult struct {
	Success bool
	Detail  *GetCustomerPhoneNoDetail
	Message string
}

func ParsePhoneLookupSOAP(xmlData string) (*PhoneLookupResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}
	if env.Body.GetCustomerPhoneNoResponse != nil {
		resp := env.Body.GetCustomerPhoneNoResponse
		if resp.Status == nil {
			return &PhoneLookupResult{
				Success: false,
				Message: "missing status",
			}, nil
		}

		if resp.Status.SuccessIndicator != "Success" {
			return &PhoneLookupResult{
				Success: false,
				Message: "API returned failure",
			}, nil
		}

		if resp.GetCByPhoneNumberType == nil || resp.GetCByPhoneNumberType.Group == nil || resp.GetCByPhoneNumberType.Group.Details == nil {
			return &PhoneLookupResult{
				Success: true,
				Message: "no details found",
			}, nil
		}

		return &PhoneLookupResult{
			Success: true,
			Detail: &GetCustomerPhoneNoDetail{
				CustomerID:  resp.GetCByPhoneNumberType.Group.Details.CustomerID,
				PhoneNumber: resp.GetCByPhoneNumberType.Group.Details.PhoneNumber,
				Email:       resp.GetCByPhoneNumberType.Group.Details.Email,
				FullName:    resp.GetCByPhoneNumberType.Group.Details.FullName,
			},
		}, nil
	}

	return &PhoneLookupResult{
		Success: false,
		Message: "no details found",
	}, nil

}
