package servicelimitfetch

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

type ServiceLimitFetchParam struct {
	CustomerNumber string
}

func NewServiceLimitFetch(param Params) string {
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
	</soapenv:Envelope>
	`, param.Password, param.Username, param.CustomerNumber)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	CustomerLimitViewResponse *CustomerLimitViewResponse `xml:"CustomerLimitViewResponse"`
}

type CustomerLimitViewResponse struct {
	Status *struct {
		TransactionID    string   `xml:"transactionId"`
		Messages         []string `xml:"messages"`
		SuccessIndicator string   `xml:"successIndicator"`
		Application      string   `xml:"application"`
	} `xml:"Status"`
	CustomerLimitType *CustomerLimitType `xml:"CUSTOMERLIMITType"`
}

type CustomerLimitType struct {
	GlobalChannelType *struct {
		MultipleChannelType []MultipleChannelType `xml:"mUSERCHANNELTYPE"`
	} `xml:"gUSERCHANNELTYPE"`
}

type MultipleChannelType struct {
	Channel           string `xml:"USERCHANNELTYPE"`
	GlobalServiceType *struct {
		MultipleServiceType []ServiceType `xml:"SERVICETYPE"`
	} `xml:"sgSERVICETYPE"`
}

type ServiceType struct {
	ServiceType      string `xml:"SERVICETYPE"`
	MaximumAmount    string `xml:"SERVICEMAXAMT"`
	MaximumFrequency string `xml:"USERMAXCNT"`
}

type ServiceLimitFetchResult struct {
	Success  bool
	Detail   []MultipleChannelType
	Messages []string
}

func ParseServiceLimitFetchSOAP(soapResponse string) (*ServiceLimitFetchResult, error) {
	var envelope Envelope
	err := xml.Unmarshal([]byte(soapResponse), &envelope)
	if err != nil {
		return nil, err
	}

	if envelope.Body.CustomerLimitViewResponse != nil {
		resp := envelope.Body.CustomerLimitViewResponse
		if resp.Status == nil {
			return &ServiceLimitFetchResult{
				Success:  false,
				Messages: []string{"API response missing status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) == "true" {
			return &ServiceLimitFetchResult{
				Success:  true,
				Messages: resp.Status.Messages,
			}, nil
		}

		if resp.CustomerLimitType == nil &&
			resp.CustomerLimitType.GlobalChannelType == nil &&
			len(resp.CustomerLimitType.GlobalChannelType.MultipleChannelType) == 0 {
			return &ServiceLimitFetchResult{
				Success:  false,
				Messages: resp.Status.Messages,
			}, nil
		}

		channelResult := make([]MultipleChannelType, 0, len(resp.CustomerLimitType.GlobalChannelType.MultipleChannelType))
		for _, channel := range resp.CustomerLimitType.GlobalChannelType.MultipleChannelType {
			serviceResult := make([]ServiceType, 0, len(channel.GlobalServiceType.MultipleServiceType))
			for _, service := range channel.GlobalServiceType.MultipleServiceType {
				serviceResult = append(serviceResult, ServiceType{
					ServiceType:      service.ServiceType,
					MaximumAmount:    service.MaximumAmount,
					MaximumFrequency: service.MaximumFrequency,
				})
			}
			channelResult = append(channelResult, MultipleChannelType{
				Channel: channel.Channel,
				GlobalServiceType: &struct {
					MultipleServiceType []ServiceType `xml:"SERVICETYPE"`
				}{
					MultipleServiceType: serviceResult,
				},
			})
		}

		return &ServiceLimitFetchResult{
			Success:  true,
			Detail:   channelResult,
			Messages: resp.Status.Messages,
		}, nil
	}

	return &ServiceLimitFetchResult{
		Success:  false,
		Messages: []string{"Invalid response format"},
	}, nil
}
