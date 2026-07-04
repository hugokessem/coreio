package customerfetch

import (
	"encoding/xml"
	"fmt"
	"strings"

	valueobject "github.com/hugokessem/coreio/lib/core/customer/customer_fetch/value_object"
)

type Params struct {
	Username       string
	Password       string
	FetchBy        string
	CustomerNumber string
}

type CustomerFetchParam struct {
	FetchBy        string
	CustomerNumber string
}

func NewCustomerFetch(param Params) string {
	if !valueobject.FetchBy(param.FetchBy).IsValid() {
		return fmt.Sprintf("invalid fetch by: %s", param.FetchBy)
	}

	var fetch strings.Builder
	if param.CustomerNumber == valueobject.FetchByAccountNumber.String() {
		fetch.WriteString(`
			<enquiryInputCollection>
				<columnName>ACCT.ID</columnName>
				<criteriaValue>`)
		fetch.WriteString(param.CustomerNumber)
		fetch.WriteString(`</criteriaValue>
		<operand>EQ</operand>
	</enquiryInputCollection>
	`)
	} else {
		fetch.WriteString(`
			<enquiryInputCollection>
				<columnName>CUS.ID</columnName>
				<criteriaValue>`)
		fetch.WriteString(param.CustomerNumber)
		fetch.WriteString(`</criteriaValue>
		<operand>EQ</operand>
	</enquiryInputCollection>
	`)
	}

	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
   <soapenv:Header/>
   <soapenv:Body>
      <cbes:CIFAccountList>
         <WebRequestCommon>
          <company/>
            <password>%s</password>
            <userName>%s</userName>
         </WebRequestCommon>
         <CIFINFOSUPERAPPType>
		 %s
         </CIFINFOSUPERAPPType>
      </cbes:CIFAccountList>
   </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, fetch.String())

}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	CIFAccountListResponse *CIFAccountListResponse `xml:"CIFAccountListResponse"`
}

type CIFAccountListResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
	} `xml:"Status"`
	CIFAccountListType *struct {
		Group *struct {
			Details []CustomerFetchDetail `xml:"mCIFINFOSUPERAPPDetailType"`
		} `xml:"gCIFINFOSUPERAPPDetailType"`
	} `xml:"CIFINFOSUPERAPPType"`
}

type CustomerFetchDetail struct {
	AccountNumber      string `xml:"AccountNo"`
	AccountName        string `xml:"AccountName"`
	Currency           string `xml:"Currency"`
	Category           string `xml:"Category"`
	AccountType        string `xml:"AccountType"`
	BranchCode         string `xml:"BranchCode"`
	BranchName         string `xml:"BranchName"`
	Balance            string `xml:"Balance"`
	RestrictionDesc    string `xml:"RestrictionDesc"`
	RestrictionType    string `xml:"RestrictionType"`
	InactiveFlag       string `xml:"InactiveFlag"`
	CustomerID         string `xml:"CustomerID"`
	CIFBranchCode      string `xml:"CIFBranchCode"`
	CIFBranchName      string `xml:"CIFBranchName"`
	CustomerName       string `xml:"CustomerName"`
	Gender             string `xml:"Gender"`
	DOB                string `xml:"DOB"`
	Mail               string `xml:"Mail"`
	Phone              string `xml:"Phone"`
	CIFRestrictType    string `xml:"CIFRestrictType"`
	CIFRestrictionDesc string `xml:"CIFRestrictionDesc"`
	CustomerGroup      string `xml:"CustomerGroup"`
	Segement           string `xml:"Segement"`
	SubSegement        string `xml:"SubSegement"`
	Industry           string `xml:"Industry"`
	Sector             string `xml:"Sector"`
	Ownership          string `xml:"Ownership"`
	Target             string `xml:"Target"`
}

type CustomerFetchResult struct {
	Success  bool
	Details  []CustomerFetchDetail
	Messages []string
}

func ParseCustomerFetchSOAP(response string) (*CustomerFetchResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(response), &env)
	if err != nil {
		return nil, err
	}

	if env.Body.CIFAccountListResponse != nil {
		resp := env.Body.CIFAccountListResponse
		if resp.Status == nil {
			return &CustomerFetchResult{
				Success:  false,
				Messages: []string{"Missing status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &CustomerFetchResult{
				Success:  false,
				Messages: []string{resp.Status.SuccessIndicator},
			}, nil
		}

		if resp.CIFAccountListType == nil ||
			resp.CIFAccountListType.Group == nil ||
			len(resp.CIFAccountListType.Group.Details) == 0 {
			return &CustomerFetchResult{
				Success:  true,
				Messages: []string{resp.Status.SuccessIndicator},
			}, nil
		}

		return &CustomerFetchResult{
			Success:  true,
			Details:  resp.CIFAccountListType.Group.Details,
			Messages: []string{resp.Status.SuccessIndicator},
		}, nil
	}

	return &CustomerFetchResult{
		Success:  false,
		Messages: []string{"Missing response"},
	}, nil

}
