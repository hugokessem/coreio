package customerlimitfetch

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

type CustomerLimitFetchParam struct {
	CustomerNumber string
}

func NewCustomerLimitFetch(param Params) string {
	return fmt.Sprintf(`
		<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
	xmlns:cbes="http://temenos.com/CBESUPERAPP">
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
	</soapenv:Envelope>
	`, param.Password, param.Username, param.CustomerNumber)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	CustomerLimitResponse *CustomerLimitResponse `xml:"CustomerLimitViewResponse"`
}

type CustomerLimitResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
		MessageId        string `xml:"messageId"`
		Application      string `xml:"application"`
		TransactionId    string `xml:"transactionId"`
	} `xml:"Status"`
	CustomerLimitDetail *CustomerLimitDetail `xml:"CUSTOMERLIMITType"`
}

type CustomerLimitDetail struct {
	UserChannelType *struct {
		Details []UserChannelDetail `xml:"mUSERCHANNELTYPE"`
	} `xml:"gUSERCHANNELTYPE"`
	CurrNo         string `xml:"CURRNO"`
	GlobalInputter *struct {
		Inputter string `xml:"INPUTTER"`
	} `xml:"gINPUTTER"`
	GlobalDatetime *struct {
		Datetime string `xml:"DATETIME"`
	} `xml:"gDATETIME"`
	Authoriser string `xml:"AUTHORISER"`
	CoCode     string `xml:"COCODE"`
	DeptCode   string `xml:"DEPTCODE"`
}

type UserChannelDetail struct {
	UserChannelType string `xml:"USERCHANNELTYPE"`
	UserMaxLimit    string `xml:"USERMAXLIMIT"`
}

type CustomerLimitFetchResult struct {
	Success bool
	Detail  *CustomerLimitDetail
	Message string
}

func ParseCustomerLimitFetchSOAP(xmlData string) (*CustomerLimitFetchResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}

	if env.Body.CustomerLimitResponse != nil {
		resp := env.Body.CustomerLimitResponse
		if resp.Status == nil {
			return &CustomerLimitFetchResult{
				Success: false,
				Message: "missing status",
			}, nil
		}
		if resp.CustomerLimitDetail == nil {
			return &CustomerLimitFetchResult{
				Success: false,
				Message: "missing customer limit detail",
			}, nil
		}
		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &CustomerLimitFetchResult{
				Success: false,
				Message: "API returned failure",
			}, nil
		}

		return &CustomerLimitFetchResult{
			Success: true,
			Detail:  resp.CustomerLimitDetail,
		}, nil
	}
	return &CustomerLimitFetchResult{
		Success: false,
		Message: "invalid response type",
	}, nil
}
