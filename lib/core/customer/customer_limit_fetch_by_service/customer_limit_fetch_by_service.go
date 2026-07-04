package customerlimitfetchbyservice

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username    string
	Password    string
	ServiceCode string
}

type CustomerLimitFetchByServiceParam struct {
	ServiceCode string
}

func NewCustomerLimitFetchByService(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:GlobalLimitView>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <GLOBALLIMITVIEWSUPERAPPType>
                <enquiryInputCollection>
                    <columnName>ID</columnName>
                    <criteriaValue>%s</criteriaValue>
                    <operand>EQ</operand>
                </enquiryInputCollection>
            </GLOBALLIMITVIEWSUPERAPPType>
        </cbes:GlobalLimitView>
    </soapenv:Body>
</soapenv:Envelope>
	`, param.Password, param.Username, param.ServiceCode)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	GenericLimitViewResponse *GenericLimitViewResponse `xml:"GlobalLimitViewResponse"`
}

type GenericLimitViewResponse struct {
	Status            *Status           `xml:"Status"`
	SuperappLimitView SuperappLimitView `xml:"GLOBALLIMITVIEWSUPERAPPType"`
}

type SuperappLimitView struct {
	CustomerLimits []CustomerLimit `xml:"gGLOBALLIMITVIEWSUPERAPPDetailType>mGLOBALLIMITVIEWSUPERAPPDetailType"`
}

type Status struct {
	TransactionID    string   `xml:"transactionId"`
	MessageID        string   `xml:"messageId"`
	SuccessIndicator string   `xml:"successIndicator"`
	Application      string   `xml:"application"`
	Messages         []string `xml:"messages"`
}
type CustomerLimit struct {
	ID               string `xml:"ID"`
	ChannelType      string `xml:"CHANNELTYPE"`
	ServiceCode      string `xml:"ServiceCode"`
	ServiceName      string `xml:"ServiceName"`
	ServiceMaxAmount string `xml:"CHANNELMAXLIMIT"`
	ServiceMinAmount string `xml:"CHANNELMINLIMIT"`
	ServiceCount     string `xml:"CHANNELCOUNT"`
}

type CustomerLimitFetchResult struct {
	Success bool
	Detail  *SuperappLimitView
	Message []string
}

func ParseCustomerLimitFetchByServiceSOAP(xmlData string) (*CustomerLimitFetchResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.GenericLimitViewResponse != nil {
		resp := env.Body.GenericLimitViewResponse
		if resp.Status == nil {
			return &CustomerLimitFetchResult{
				Success: false,
				Message: []string{"missing status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &CustomerLimitFetchResult{
				Success: false,
				Message: resp.Status.Messages,
			}, nil
		}

		return &CustomerLimitFetchResult{
			Success: true,
			Detail:  &resp.SuperappLimitView,
		}, nil
	}

	return &CustomerLimitFetchResult{
		Success: false,
		Message: []string{"invalid response"},
	}, nil
}
