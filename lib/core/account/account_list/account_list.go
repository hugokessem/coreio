package accountlist

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username      string
	Password      string
	ColumnName    string
	CriteriaValue string
}

type AccountListParams struct {
	ColumnName    string
	CriteriaValue string
}

func NewAccountList(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:AccountListByCIF>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <ACCOUNTINFOSUPERAPPType>
                <enquiryInputCollection>
                    <columnName>%s</columnName>
                    <criteriaValue>%s</criteriaValue>
                    <operand>EQ</operand>
                </enquiryInputCollection>
                <enquiryInputCollection>
                    <columnName>CUS.ID</columnName>
                    <criteriaValue></criteriaValue>
                    <operand>EQ</operand>
                </enquiryInputCollection>
            </ACCOUNTINFOSUPERAPPType>
        </cbes:AccountListByCIF>
    </soapenv:Body>
</soapenv:Envelope>
	`, param.Password, param.Username, param.ColumnName, param.CriteriaValue)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	AccountListByCIFResponse *AccountListByCIFResponse `xml:"AccountListByCIFResponse"`
}

type AccountListByCIFResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
	} `xml:"Status"`
	AccountListByCIFType *struct {
		Group *struct {
			Details []AccountListByCIFDetail `xml:"mACCOUNTINFOSUPERAPPDetailType"`
		} `xml:"gACCOUNTINFOSUPERAPPDetailType"`
	} `xml:"ACCOUNTINFOSUPERAPPType"`
}

type AccountListByCIFDetail struct {
	AccountNumber   string `xml:"AccountNumber"`
	CustomerName    string `xml:"CustomerName"`
	Restriction     string `xml:"Restriction"`
	Currency        string `xml:"Currency"`
	CustomerID      string `xml:"CustomerID"`
	Category        string `xml:"Category"`
	AccountType     string `xml:"AccountType"`
	BranchCode      string `xml:"BranchCode"`
	BranchName      string `xml:"BranchName"`
	DistrictName    string `xml:"DistrictName"`
	PhoneNo         string `xml:"PhoneNo"`
	Industry        string `xml:"Industry"`
	Sector          string `xml:"Sector"`
	Ownership       string `xml:"Ownership"`
	CustomerSegment string `xml:"CustomerSegment"`
	Target          string `xml:"Target"`
}

type AccountListResult struct {
	Success bool
	Details []AccountListByCIFDetail
	Message []string
}

func ParseAccountListSOAP(xmlData string) (*AccountListResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}

	if env.Body.AccountListByCIFResponse != nil {
		resp := env.Body.AccountListByCIFResponse
		if resp.Status == nil {
			return &AccountListResult{
				Success: false,
				Message: []string{"Missing Status"},
			}, nil
		}
		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &AccountListResult{
				Success: false,
				Message: []string{"API returned failure"},
			}, nil
		}
		if resp.AccountListByCIFType == nil ||
			resp.AccountListByCIFType.Group == nil ||
			len(resp.AccountListByCIFType.Group.Details) == 0 {
			return &AccountListResult{
				Success: true,
				Message: []string{"No account details found"},
			}, nil
		}

		details := resp.AccountListByCIFType.Group.Details
		detailsList := make([]AccountListByCIFDetail, len(details))
		for i, detail := range details {
			detailsList[i] = AccountListByCIFDetail{
				AccountNumber:   detail.AccountNumber,
				CustomerName:    detail.CustomerName,
				Restriction:     detail.Restriction,
				Currency:        detail.Currency,
				CustomerID:      detail.CustomerID,
				Category:        detail.Category,
				AccountType:     detail.AccountType,
				BranchCode:      detail.BranchCode,
				BranchName:      detail.BranchName,
				DistrictName:    detail.DistrictName,
				PhoneNo:         detail.PhoneNo,
				Industry:        detail.Industry,
				Sector:          detail.Sector,
				Ownership:       detail.Ownership,
				CustomerSegment: detail.CustomerSegment,
				Target:          detail.Target,
			}
		}
		return &AccountListResult{
			Success: true,
			Details: detailsList,
			Message: []string{},
		}, nil
	}
	return &AccountListResult{
		Success: false,
		Message: []string{"Invalid response format"},
	}, nil
}
