package customerlimitamendbycif

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type Params struct {
	Username       string
	Password       string
	CustomerNumber string
	ChannelLimit   []ChannelLimit
}

type Channel string

const (
	APP  Channel = "APP"
	USSD Channel = "USSD"
)

type ChannelLimit struct {
	Channel       Channel
	ServiceLimits []ServiceLimit
}

type ServiceLimit struct {
	ServiceType           string
	ServiceMaximumAmount  string
	UserMaximumDebitCount string
}

type CustomerLimitAmendByCIFParam struct {
	CustomerNumber string
	ChannelLimit   []ChannelLimit
}

func NewCustomerLimitAmendByCIF(param Params) string {
	limit := make([]string, 0, len(param.ChannelLimit))
	for channelIndex, channelLimit := range param.ChannelLimit {
		serviceLimits := make([]string, 0, len(channelLimit.ServiceLimits))
		for serviceIndex, serviceLimit := range channelLimit.ServiceLimits {
			serviceLimits = append(serviceLimits, fmt.Sprintf(`
				<cus:SERVICETYPE s="%s">
				<cus:SERVICETYPE>%s</cus:SERVICETYPE>
				<cus:SERVICEMAXAMT>%s</cus:SERVICEMAXAMT>
				<cus:USERMAXCNT>%s</cus:USERMAXCNT>
				</cus:SERVICETYPE>
			`, strconv.Itoa(serviceIndex+1), serviceLimit.ServiceType, serviceLimit.ServiceMaximumAmount, serviceLimit.UserMaximumDebitCount))
		}
		limit = append(limit, fmt.Sprintf(`
               <cus:mUSERCHANNELTYPE m="%s">
                  <cus:USERCHANNELTYPE>%s</cus:USERCHANNELTYPE>
                  <cus:sgSERVICETYPE sg="1">%s</cus:sgSERVICETYPE>
               </cus:mUSERCHANNELTYPE>
			`, strconv.Itoa(channelIndex+1), channelLimit.Channel, strings.Join(serviceLimits, "")))
	}

	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP" xmlns:cus="http://temenos.com/CUSTOMERLIMITCUSTOMERSERVICEAMEND">
   <soapenv:Header/>
   <soapenv:Body>
      <cbes:CustomerServiceLimitAmendment>
         <WebRequestCommon>
            <company/>
            <password>%s</password>
            <userName>%s</userName>
         </WebRequestCommon>
         <OfsFunction>
         </OfsFunction>
         <CUSTOMERLIMITCUSTOMERSERVICEAMENDType id="%s">
            <cus:gUSERCHANNELTYPE g="1">%s</cus:gUSERCHANNELTYPE>
         </CUSTOMERLIMITCUSTOMERSERVICEAMENDType>
      </cbes:CustomerServiceLimitAmendment>
   </soapenv:Body>
</soapenv:Envelope>
`, param.Password, param.Username, param.CustomerNumber, strings.Join(limit, ""))
}

// ---- Parser for CustomerServiceLimitAmendmentResponse ----

type envelopeCSLAM struct {
	XMLName xml.Name  `xml:"Envelope"`
	Body    bodyCSLAM `xml:"Body"`
}

type bodyCSLAM struct {
	Response *customerServiceLimitAmendmentResponse `xml:"CustomerServiceLimitAmendmentResponse"`
}

type customerServiceLimitAmendmentResponse struct {
	Status            *statusType        `xml:"Status"`
	CustomerLimitType *CustomerLimitType `xml:"CUSTOMERLIMITType"`
}

type statusType struct {
	TransactionID    string   `xml:"transactionId"`
	MessageID        string   `xml:"messageId"`
	SuccessIndicator string   `xml:"successIndicator"`
	Application      string   `xml:"application"`
	Messages         []string `xml:"messages"`
}

type CustomerLimitType struct {
	XMLName      xml.Name          `xml:"CUSTOMERLIMITType"`
	ID           string            `xml:"id,attr"`
	ACCTCOUNT    string            `xml:"ACCTCOUNT"`
	GUserChannel *gUserChannelType `xml:"gUSERCHANNELTYPE"`
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

type gUserChannelType struct {
	MUserChannel []mUserChannelType `xml:"mUSERCHANNELTYPE"`
}

type mUserChannelType struct {
	UserChannelType string         `xml:"USERCHANNELTYPE"`
	SGServiceType   *sGServiceType `xml:"sgSERVICETYPE"`
}

type sGServiceType struct {
	Services []serviceType `xml:"SERVICETYPE"`
}

type serviceType struct {
	Name          string `xml:"SERVICETYPE"`
	ServiceMaxAmt string `xml:"SERVICEMAXAMT"`
	UserMaxCnt    string `xml:"USERMAXCNT"`
}

type CustomerLimitAmendResult struct {
	Success  bool
	Detail   *CustomerLimitType
	Messages []string
}

func ParseCustomerLimitAmendByCIFSOAP(xmlData string) (*CustomerLimitAmendResult, error) {
	var env envelopeCSLAM
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.Response != nil {
		resp := env.Body.Response
		if resp.Status == nil {
			return &CustomerLimitAmendResult{
				Success:  false,
				Messages: []string{"missing status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			messages := make([]string, 0, len(resp.Status.Messages))
			for _, v := range resp.Status.Messages {
				messages = append(messages, v)
			}

			return &CustomerLimitAmendResult{
				Success:  false,
				Detail:   resp.CustomerLimitType,
				Messages: messages,
			}, nil
		}

		return &CustomerLimitAmendResult{
			Success: true,
			Detail:  resp.CustomerLimitType,
		}, nil
	}

	return &CustomerLimitAmendResult{
		Success: false, Messages: []string{"invalid response"},
	}, nil
}
