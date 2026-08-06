package verifyaml

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Param struct {
	Password       string
	UserName       string
	CustomerNumber string
}

type VerifyAMLParam struct {
	CustomerNumber string
}

func NewVerifyAML(param Param) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
   <soapenv:Header/>
   <soapenv:Body>
      <cbes:AMLStatusCheckSuperApp>
         <WebRequestCommon>
            <company/>
            <password>%s</password>
            <userName>%s</userName>
         </WebRequestCommon>
         <AMLSTATUSSUPERAPPType>
            <enquiryInputCollection>
               <columnName>ID</columnName>
               <criteriaValue>%s</criteriaValue>
               <operand>CT</operand>
            </enquiryInputCollection>
         </AMLSTATUSSUPERAPPType>
      </cbes:AMLStatusCheckSuperApp>
   </soapenv:Body>
</soapenv:Envelope>
	`, param.Password, param.UserName, param.CustomerNumber)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	AMLStatusCheckSuperAppResponse *AMLStatusCheckSuperAppResponse `xml:"AMLStatusCheckSuperAppResponse"`
}

type AMLStatusCheckSuperAppResponse struct {
	Status                *Status                `xml:"Status"`
	AMLSTATUSSUPERAPPType *AMLSTATUSSUPERAPPType `xml:"AMLSTATUSSUPERAPPType"`
}

type Status struct {
	SuccessIndicator string   `xml:"successIndicator"`
	Messages         []string `xml:"messages"`
}

type AMLSTATUSSUPERAPPType struct {
	DetailGroup *struct {
		Detail *AMLStatusDetail `xml:"mAMLSTATUSSUPERAPPDetailType"`
	} `xml:"gAMLSTATUSSUPERAPPDetailType"`
}

type AMLStatusDetail struct {
	FCMStatus string `xml:"FCMSTATUS"`
}

type VerifyAMLResult struct {
	Success   bool
	FCMStatus string
	Messages  []string
}

func ParseVerifyAMLSOAP(xmlData string) (*VerifyAMLResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.AMLStatusCheckSuperAppResponse == nil {
		return &VerifyAMLResult{
			Success:  false,
			Messages: []string{"Invalid response type"},
		}, nil
	}

	resp := env.Body.AMLStatusCheckSuperAppResponse
	if resp.Status == nil {
		return &VerifyAMLResult{
			Success:  false,
			Messages: []string{"Missing Status"},
		}, nil
	}

	if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
		messages := resp.Status.Messages
		if len(messages) == 0 {
			messages = []string{"AML status check failed"}
		}
		return &VerifyAMLResult{
			Success:  false,
			Messages: messages,
		}, nil
	}

	if resp.AMLSTATUSSUPERAPPType == nil ||
		resp.AMLSTATUSSUPERAPPType.DetailGroup == nil ||
		resp.AMLSTATUSSUPERAPPType.DetailGroup.Detail == nil {
		messages := resp.Status.Messages
		if len(messages) == 0 {
			messages = []string{"No AML status details found"}
		}
		return &VerifyAMLResult{
			Success:  false,
			Messages: messages,
		}, nil
	}

	fcmStatus := strings.TrimSpace(resp.AMLSTATUSSUPERAPPType.DetailGroup.Detail.FCMStatus)
	return &VerifyAMLResult{
		Success:   true,
		FCMStatus: fcmStatus,
		Messages:  nil,
	}, nil
}
