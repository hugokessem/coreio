package lockedamountlist

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username      string
	Password      string
	AccountNumber string
}
type ListLockedAmountParam struct {
	AccountNumber string
}

func NewListLockedAmount(param Params) string {
	return fmt.Sprintf(
		`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
	<soapenv:Header/>
	<soapenv:Body>
	<cbes:LockedAmountInquiry>
		<WebRequestCommon>
		<company/>
		<password>%s</password>
		<userName>%s</userName>
	</WebRequestCommon>
		<ACCTLOCKEDAMOUNTSSUPERAPPType>
			<enquiryInputCollection>
				<columnName>ACCOUNT.NUMBER</columnName>
				<criteriaValue>%s</criteriaValue>
				<operand>EQ</operand>
				</enquiryInputCollection>
		</ACCTLOCKEDAMOUNTSSUPERAPPType>
	</cbes:LockedAmountInquiry>
	</soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.AccountNumber)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	LockedAmounResponse             *ListLockedAmountResponse        `xml:"LockedAmountResponse"`
	ListLockedAmountInquiryResponse *ListLockedAmountInquiryResponse `xml:"LockedAmountInquiryResponse"`
}

type ListLockedAmountInquiryResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
	} `xml:"Status"`
	AccountLockedAmountsSuperappType *struct {
		Group *struct {
			Details []ListLockedAmountDetail `xml:"mACCTLOCKEDAMOUNTSSUPERAPPDetailType"`
		} `xml:"gACCTLOCKEDAMOUNTSSUPERAPPDetailType"`
	} `xml:"ACCTLOCKEDAMOUNTSSUPERAPPType"`
}

type ListLockedAmountDetail struct {
	AccountNumber string `xml:"AccountNumber"`
	Remark        string `xml:"Remark"`
	LockedAmount  string `xml:"LockedAmount"`
	LockedDate    string `xml:"LockedDate"`
}

type ListLockedAmountResponse struct {
	Status *struct {
		SuccessIndicator string   `xml:"successIndicator"`
		Message          []string `xml:"message"`
	} `xml:"Status"`
}

type ListLockedAmountResult struct {
	Success bool
	Details []ListLockedAmountDetail
	Message []string
}

func ParseListLockedAmountSOAP(xmlData string) (*ListLockedAmountResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}

	// Check for success in ListLockedAmountInquiryResponse
	if env.Body.ListLockedAmountInquiryResponse != nil {
		resp := env.Body.ListLockedAmountInquiryResponse
		if resp.Status == nil {
			return &ListLockedAmountResult{
				Success: false,
				Message: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &ListLockedAmountResult{
				Success: false,
				Message: []string{"API returned failure"},
			}, nil
		}

		if resp.AccountLockedAmountsSuperappType == nil ||
			resp.AccountLockedAmountsSuperappType.Group == nil ||
			len(resp.AccountLockedAmountsSuperappType.Group.Details) == 0 {
			return &ListLockedAmountResult{
				Success: true,
				Message: []string{"No Locked Amount Found!"},
			}, nil
		}

		details := resp.AccountLockedAmountsSuperappType.Group.Details
		detailsList := make([]ListLockedAmountDetail, len(details))
		for i, detail := range details {
			detailsList[i] = ListLockedAmountDetail{
				AccountNumber: detail.AccountNumber,
				Remark:        detail.Remark,
				LockedAmount:  detail.LockedAmount,
				LockedDate:    detail.LockedDate,
			}
		}

		return &ListLockedAmountResult{
			Success: true,
			Details: detailsList,
		}, nil
	}

	if env.Body.LockedAmounResponse != nil {
		resp := env.Body.LockedAmounResponse
		messages := []string{}
		if resp.Status != nil && len(resp.Status.Message) > 0 {
			messages = resp.Status.Message
		}
		success := false
		if resp.Status != nil && strings.ToLower(resp.Status.SuccessIndicator) == "success" {
			success = true
		}
		return &ListLockedAmountResult{
			Success: success,
			Message: messages,
		}, nil
	}

	return &ListLockedAmountResult{
		Success: false,
		Message: []string{"Invalid response format"},
	}, nil
}
