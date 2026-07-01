package fayda

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username string
	Password string
	NID      string
}

type FaydaParam struct {
	NID string
}

func NewFayda(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:CustomerUniqueVerification>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <CUSTOMERVERIFYENQSUPERAPPType>
                <enquiryInputCollection>
                    <columnName>NID</columnName>
                    <criteriaValue>%s</criteriaValue>
                    <operand>EQ</operand>
                </enquiryInputCollection>
            </CUSTOMERVERIFYENQSUPERAPPType>
        </cbes:CustomerUniqueVerification>
    </soapenv:Body>
</soapenv:Envelope>
	`, param.Password, param.Username, param.NID)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	CustomerUniqueVerificationResponse *CustomerUniqueVerificationResponse `xml:"CustomerUniqueVerificationResponse"`
}

type CustomerUniqueVerificationResponse struct {
	Status *struct {
		SuccessIndicator string   `xml:"successIndicator"`
		Messages         []string `xml:"messages"`
	} `xml:"Status"`
	CUSTOMERVERIFYENQSUPERAPPType *struct {
		Group *struct {
			Details *FaydaDetail `xml:"mCUSTOMERVERIFYENQSUPERAPPDetailType"`
		} `xml:"gCUSTOMERVERIFYENQSUPERAPPDetailType"`
	} `xml:"CUSTOMERVERIFYENQSUPERAPPType"`
}

type FaydaDetail struct {
	CustomerID   string `xml:"CustomerID"`
	CustomerName string `xml:"SHORTNAME"`
	CustomerFlag string `xml:"CustomerFlag"`
}

type FaydaResult struct {
	Success bool
	Detail  *FaydaDetail
	Message []string
}

func ParseFaydaSOAP(xmlData string) (*FaydaResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}
	if env.Body.CustomerUniqueVerificationResponse != nil {
		resp := env.Body.CustomerUniqueVerificationResponse
		if resp.Status == nil {
			return &FaydaResult{
				Success: false,
				Message: []string{"missing status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &FaydaResult{
				Success: false,
				Message: resp.Status.Messages,
			}, nil
		}

		if resp.CUSTOMERVERIFYENQSUPERAPPType == nil || resp.CUSTOMERVERIFYENQSUPERAPPType.Group == nil || resp.CUSTOMERVERIFYENQSUPERAPPType.Group.Details == nil {
			return &FaydaResult{
				Success: false,
				Message: []string{"no details found"},
			}, nil
		}

		return &FaydaResult{
			Success: true,
			Detail: &FaydaDetail{
				CustomerID:   resp.CUSTOMERVERIFYENQSUPERAPPType.Group.Details.CustomerID,
				CustomerName: resp.CUSTOMERVERIFYENQSUPERAPPType.Group.Details.CustomerName,
				CustomerFlag: resp.CUSTOMERVERIFYENQSUPERAPPType.Group.Details.CustomerFlag,
			},
		}, nil
	}

	return &FaydaResult{
		Success: false,
		Message: []string{"Invalid response type"},
	}, nil
}
