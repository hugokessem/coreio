package customerlimitfetchbycif

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

type CustomerLimitFetchByCIFParam struct {
	CustomerNumber string
}

func NewCustomerLimitFetchByCIF(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
    <soapenv:Header/>
		<soapenv:Body>
			<cbes:CustomerLimitView>
				<WebRequestCommon>
					<company/>
					<password>%s</password>
					<userName>%s</userName>
				</WebRequestCommon>
				<CUSTOMERLIMITCUSTOMLIMITType>
					<transactionId>%s</transactionId>
				</CUSTOMERLIMITCUSTOMLIMITType>
			</cbes:CustomerLimitView>
		</soapenv:Body>
	</soapenv:Envelope>`, param.Password, param.Username, param.CustomerNumber)
}

// Envelope/Body follow common patterns in repo SOAP parsers
type EnvelopeCV struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    BodyCV   `xml:"Body"`
}

type BodyCV struct {
	CustomerLimitViewResponse *CustomerLimitViewResponse `xml:"CustomerLimitViewResponse"`
}

type CustomerLimitViewResponse struct {
	Status            *Status            `xml:"Status"`
	CustomerLimitType *CustomerLimitType `xml:"CUSTOMERLIMITType"`
}

type Status struct {
	TransactionID    string   `xml:"transactionId"`
	MessageID        string   `xml:"messageId"`
	SuccessIndicator string   `xml:"successIndicator"`
	Application      string   `xml:"application"`
	Messages         []string `xml:"messages"`
}

type CustomerLimitType struct {
	XMLName      xml.Name          `xml:"CUSTOMERLIMITType"`
	ID           string            `xml:"id,attr"`
	GUserChannel *GUserChannelType `xml:"gUSERCHANNELTYPE"`
	CURRNO       string            `xml:"CURRNO"`
	GInputter    *struct {
		Inputter string `xml:"INPUTTER"`
	} `xml:"gINPUTTER"`
	GDateTime *struct {
		DateTime string `xml:"DATETIME"`
	} `xml:"gDATETIME"`
	AUTHORISER string `xml:"AUTHORISER"`
	COCODE     string `xml:"COCODE"`
	DEPTCODE   string `xml:"DEPTCODE"`
}

type GUserChannelType struct {
	MUserChannel []MUserChannelType `xml:"mUSERCHANNELTYPE"`
}

type MUserChannelType struct {
	UserChannelType string         `xml:"USERCHANNELTYPE"`
	SGServiceType   *SGServiceType `xml:"sgSERVICETYPE"`
}

type SGServiceType struct {
	Services []ServiceType `xml:"SERVICETYPE"`
}

type ServiceType struct {
	Name          string `xml:"SERVICETYPE"`
	ServiceMaxAmt string `xml:"SERVICEMAXAMT"`
	UserMaxCnt    string `xml:"USERMAXCNT"`
}

type CustomerLimitViewResult struct {
	Success  bool
	Detail   *CustomerLimitType
	Messages []string
}

func ParseCustomerLimitFetchByCIFSOAP(xmlData string) (*CustomerLimitViewResult, error) {
	var env EnvelopeCV
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.CustomerLimitViewResponse != nil {
		resp := env.Body.CustomerLimitViewResponse
		if resp.Status == nil {
			return &CustomerLimitViewResult{
				Success:  false,
				Messages: []string{"missing status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			messages := make([]string, 0, len(resp.Status.Messages))
			for _, v := range resp.Status.Messages {
				messages = append(messages, v)
			}

			return &CustomerLimitViewResult{
				Success:  false,
				Messages: messages,
			}, nil
		}

		return &CustomerLimitViewResult{
			Success: true,
			Detail:  resp.CustomerLimitType,
		}, nil
	}

	return &CustomerLimitViewResult{
		Success:  false,
		Messages: []string{"invalid response"},
	}, nil
}
