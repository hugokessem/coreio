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
			<cbes:GenericLimitView>
				<WebRequestCommon>
					<company/>
					<password>%s</password>
					<userName>%s</userName>
				</WebRequestCommon>
				<CUSTOMERLIMITVIEWType>
					<transactionId>%s</transactionId>
				</CUSTOMERLIMITVIEWType>
			</cbes:GenericLimitView>
		</soapenv:Body>
	</soapenv:Envelope>
	`, param.Password, param.Username, param.ServiceCode)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	GenericLimitViewResponse *GenericLimitViewResponse `xml:"GenericLimitViewResponse"`
}

type GenericLimitViewResponse struct {
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
	XMLName      xml.Name      `xml:"CUSTOMERLIMITType"`
	ID           string        `xml:"id,attr"`
	GChannelType *GChannelType `xml:"gCHANNELTYPE"`
	CURRNO       string        `xml:"CURRNO"`
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

type GChannelType struct {
	MChannelType []MChannelType `xml:"mCHANNELTYPE"`
}

type MChannelType struct {
	ChannelType    string         `xml:"CHANNELTYPE"`
	SGServiceTypes *SGServiceType `xml:"sgGSERVICETYPE"`
}

type SGServiceType struct {
	GServiceType []GServiceType `xml:"GSERVICETYPE"`
}

type GServiceType struct {
	Name            string `xml:"GSERVICETYPE"`
	CHANNELMAXLIMIT string `xml:"CHANNELMAXLIMIT"`
	CHANNELMINLIMIT string `xml:"CHANNELMINLIMIT"`
	CHANNELCOUNT    string `xml:"CHANNELCOUNT"`
}

type CustomerLimitFetchResult struct {
	Success bool
	Detail  *CustomerLimitType
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
				Message: []string{resp.Status.Application},
			}, nil
		}

		return &CustomerLimitFetchResult{
			Success: true,
			Detail:  resp.CustomerLimitType,
		}, nil
	}

	return &CustomerLimitFetchResult{
		Success: false,
		Message: []string{"invalid response"},
	}, nil
}
