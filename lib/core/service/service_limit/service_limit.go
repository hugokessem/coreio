package servicelimit

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username     string
	Password     string
	ChannelLimit []ChannelLimit
}

type ServiceLimitParam struct {
	ChannelLimit []ChannelLimit
}

type ChannelLimit struct {
	UserChannelType string // USSD, APP
	ServiceLimits   []ServiceLimits
}

type ServiceLimits struct {
	ServiceType   string `xml:"SERVICETYPE"`
	ServiceMaxAmt string `xml:"SERVICEMAXAMT"`
	UserMaxCnt    string `xml:"USERMAXCNT"`
}

func NewServiceLimit(param Params) string {

	channelLimit := make([]string, 0, len(param.ChannelLimit))
	for index, value := range param.ChannelLimit {
		serviceLimits := make([]string, 0, len(value.ServiceLimits))
		for innerIndex, innterValue := range value.ServiceLimits {
			serviceLimits = append(serviceLimits, fmt.Sprintf(`
				<cus:SERVICETYPE s="%d">
					<cus:SERVICETYPE>%s</cus:SERVICETYPE>
					<cus:SERVICEMAXAMT>%s</cus:SERVICEMAXAMT>
					<cus:USERMAXCNT>%s</cus:USERMAXCNT>
				</cus:SERVICETYPE>
			`, innerIndex, innterValue.ServiceType, innterValue.ServiceMaxAmt, innterValue.UserMaxCnt))

		}
		channelLimit = append(channelLimit, fmt.Sprintf(`
			<cus:USERCHANNELTYPE>%s</cus:USERCHANNELTYPE>
			<cus:sgSERVICETYPE sg="%d">
				%s
			</cus:sgSERVICETYPE>
		`, value.UserChannelType, index, strings.Join(serviceLimits, "")))

	}

	return fmt.Sprintf(`
		<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
			xmlns:cbes="http://temenos.com/CBESUPERAPP"
			xmlns:cus="http://temenos.com/CUSTOMERLIMITCUSTOMERSERVICEAMEND">
		<soapenv:Header/>
		<soapenv:Body>
			<cbes:CustomerServiceLimitAmendment>
				<WebRequestCommon>
					<company/>
					<password>%s</password>
					<userName>%s</userName>
				</WebRequestCommon>
				<OfsFunction/>
				<CUSTOMERLIMITCUSTOMERSERVICEAMENDType id="1005">
					<cus:gUSERCHANNELTYPE g="1">
						%s
					</cus:gUSERCHANNELTYPE>
				</CUSTOMERLIMITCUSTOMERSERVICEAMENDType>
			</cbes:CustomerServiceLimitAmendment>
		</soapenv:Body>
	</soapenv:Envelope>
	`, param.Password, param.Username, strings.Join(channelLimit, ""))
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	CustomerServiceLimitAmendmentResponse *CustomerServiceLimitAmendmentResponse `xml:"CustomerServiceLimitAmendmentResponse"`
}

type CustomerServiceLimitAmendmentResponse struct {
	Status *struct {
		SuccessIndicator string   `xml:"successIndicator"`
		Messages         []string `xml:"messages"`
		Application      string   `xml:"application"`
		TransactionId    string   `xml:"transactionId"`
	} `xml:"Status"`
	CustomerLimitType *struct {
		Group *struct {
			Details []CustomerLimitDetail `xml:"mUSERCHANNELTYPE"`
		} `xml:"gUSERCHANNELTYPE"`
	} `xml:"CUSTOMERLIMITType"`
}

type CustomerLimitDetail struct {
	UserChannelType string        `xml:"USERCHANNELTYPE"`
	ServiceType     []ServiceType `xml:"sgSERVICETYPE"`
}

type ServiceType struct {
	ServiceLimits []ServiceLimits `xml:"SERVICETYPE"`
}

type ServiceLimitResult struct {
	Success  bool
	Messages []string
	Details  []CustomerLimitDetail
}

func ParseServiceLimitSOAP(xmlData string) (*ServiceLimitResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}
	if env.Body.CustomerServiceLimitAmendmentResponse != nil && env.Body.CustomerServiceLimitAmendmentResponse.Status != nil {
		resp := env.Body.CustomerServiceLimitAmendmentResponse.Status
		if strings.ToLower(resp.SuccessIndicator) != "success" {
			return &ServiceLimitResult{
				Success:  false,
				Messages: []string{"API returned failure"},
			}, nil
		}

		if env.Body.CustomerServiceLimitAmendmentResponse.CustomerLimitType == nil {
			return &ServiceLimitResult{
				Success:  false,
				Messages: []string{"Missing CustomerLimitType"},
			}, nil
		}

		if env.Body.CustomerServiceLimitAmendmentResponse.CustomerLimitType.Group == nil {
			return &ServiceLimitResult{
				Success: true,
				Details: []CustomerLimitDetail{},
			}, nil
		}

		return &ServiceLimitResult{
			Success: true,
			Details: env.Body.CustomerServiceLimitAmendmentResponse.CustomerLimitType.Group.Details,
		}, nil
	}
	return &ServiceLimitResult{
		Success:  false,
		Messages: []string{"Invalid response type"},
	}, nil

}
