package servicedetail

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
type ServiceDetailParams struct {
	ServiceCode string
}

func NewServiceDetail(param Params) string {
	var columnName, operand string
	if param.ServiceCode != "" {
		columnName = "@ID"
		operand = "EQ"
	}
	return fmt.Sprintf(`
		<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
		<soapenv:Header/>
		<soapenv:Body>
			<cbes:ServiceTypeList>
				<WebRequestCommon>
					<company></company>
					<password>%s</password>
					<userName>%s</userName>
				</WebRequestCommon>
				<SERVICELISTVIEWSUPERAPPType>
					<enquiryInputCollection>
					<columnName>%s</columnName>
					<criteriaValue>%s</criteriaValue>
					<operand>%s</operand>
					</enquiryInputCollection>
				</SERVICELISTVIEWSUPERAPPType>
			</cbes:ServiceTypeList>
		</soapenv:Body>
		</soapenv:Envelope>
	`, param.Password, param.Username, columnName, param.ServiceCode, operand)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	ServiceTypeListResponse *ServiceTypeListResponse `xml:"ServiceTypeListResponse"`
}

type ServiceTypeListResponse struct {
	Status *struct {
		SuccessIndicator string   `xml:"successIndicator"`
		Messages         []string `xml:"messages"`
	} `xml:"Status"`
	ServiceTypeList *struct {
		Group *struct {
			ServiceDetails []ServiceDetail `xml:"mSERVICELISTVIEWSUPERAPPDetailType"`
		} `xml:"gSERVICELISTVIEWSUPERAPPDetailType"`
	} `xml:"SERVICELISTVIEWSUPERAPPType"`
}

type ServiceDetail struct {
	ID                 string `xml:"ID"`
	ServiceDescription string `xml:"ServiceDesc"`
	CBEChargeName      string `xml:"CBEChargeName"`
	CBEChargeCode      string `xml:"CBEChargeCode"`
	ParentChargeCode   string `xml:"PartnerChargeCode"`
	PartnerChargeName  string `xml:"PartnerChargeName"`
	MaximumAmount      string `xml:"MaxAmount"`
	MaximumFrequency   string `xml:"MaxFreq"`
	Description        string `xml:"DESCRIPTION"`
}

type ServiceDetailResult struct {
	Success  bool
	Detail   []ServiceDetail
	Messages []string
}

func ParseServiceDetailSOAP(response string) (*ServiceDetailResult, error) {
	var envelope Envelope
	err := xml.Unmarshal([]byte(response), &envelope)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal SOAP response: %w", err)
	}

	if envelope.Body.ServiceTypeListResponse != nil {
		resp := envelope.Body.ServiceTypeListResponse
		if resp.Status == nil {
			return &ServiceDetailResult{
				Success:  false,
				Messages: []string{"Failed to parse service detail response"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &ServiceDetailResult{
				Success:  false,
				Messages: resp.Status.Messages,
			}, nil
		}

		if resp.ServiceTypeList != nil && resp.ServiceTypeList.Group != nil {
			return &ServiceDetailResult{
				Success:  true,
				Detail:   resp.ServiceTypeList.Group.ServiceDetails,
				Messages: []string{},
			}, nil
		}
	}

	return &ServiceDetailResult{
		Success:  false,
		Messages: []string{"Failed to parse service detail response"},
	}, nil
}
