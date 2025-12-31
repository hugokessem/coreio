package customercreation

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username           string
	Password           string
	FirstName          string
	MiddleName         string
	LastName           string
	PhoneNumber        string
	Address            string
	PostalCode         string
	ISOCountryCode     string
	AccountOffice      string
	Industry           string
	ISONationalityCode string
	ISOResidentCode    string
	UniqueID           string
	IssuesBy           string
	IssuedDate         string
	ExpiryDate         string
	Gender             string
	DateOfBirth        string
	MaritalStatus      string
	Email              string
	EmploymentStatus   string
	Occupation         string
	EmployerName       string
	EmployerAddress    string
	EmployerBusiness   string
	CustomerCurrency   string
	Salary             string
	AnnualBonus        string
	NetMonthlyIncome   string
	NetMonthlyExpence  string
	TinNumber          string
	MotherName         string
	CustomerGroup      string
}

func LastNineDigits(phone string) (string, bool) {
	if len(phone) < 9 {
		return "", false
	}

	return phone[len(phone)-9:], true
}

func FullName(param Params) string {
	return fmt.Sprintf("%s %s %s", param.FirstName, param.MiddleName, param.LastName)
}

func SetMenemoic(param Params) string {
	lastNineDigit, ok := LastNineDigits(param.PhoneNumber)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s%s", string(param.FirstName[0]), lastNineDigit)
}

func NewCustomerCreation(param Params) string {
	menemoic := SetMenemoic(param)
	fullName := FullName(param)
	return fmt.Sprintf(`
    <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
xmlns:iib="http://temenos.com/IIBONBOARDING"
xmlns:cus="http://temenos.com/CUSTOMERCREATEINDIVIDUAL">
    <soapenv:Header/>
    <soapenv:Body>
        <iib:CustomerOpening>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <OfsFunction></OfsFunction>
            <CUSTOMERCREATEINDIVIDUALType id="">
                <cus:MNEMONIC>%s</cus:MNEMONIC>
                <cus:gSHORTNAME g="1">
                    <cus:SHORTNAME>%s</cus:SHORTNAME>
                </cus:gSHORTNAME>
                <cus:gNAME1 g="1">
                    <cus:NAME1>%s</cus:NAME1>
                </cus:gNAME1>
                <cus:gNAME2 g="1">
                    <cus:NAME2>%s</cus:NAME2>
                </cus:gNAME2>
                <cus:gSTREET g="1">
                    <cus:STREET>AM</cus:STREET>
                </cus:gSTREET>
                <cus:gLLADDRESS g="1">
                    <cus:mLLADDRESS m="1">
                        <cus:sgLLADDRESS sg="1">
                            <cus:ADDRESS s="1">
                                <cus:ADDRESS>%s</cus:ADDRESS>
                            </cus:ADDRESS>
                        </cus:sgLLADDRESS>
                    </cus:mLLADDRESS>
                </cus:gLLADDRESS>
                <cus:gTOWNCOUNTRY g="1">
                    <cus:TOWNCOUNTRY>%s</cus:TOWNCOUNTRY>
                </cus:gTOWNCOUNTRY>
                <cus:gPOSTCODE g="1">
                    <cus:POSTCODE>%s</cus:POSTCODE>
                </cus:gPOSTCODE>
                <cus:gCOUNTRY g="1">
                    <cus:COUNTRY>%s</cus:COUNTRY>
                </cus:gCOUNTRY>
                <cus:gRELATIONCODE g="1">
                    <cus:mRELATIONCODE m="1">
                        <cus:RELATIONCODE/>
                        <cus:RELCUSTOMER/>
                        <cus:sgRELDELIVOPT sg="1">
                            <cus:RELDELIVOPT s="1">
                                <cus:RELDELIVOPT/>
                                <cus:ROLE/>
                                <cus:ROLEMOREINFO/>
                                <cus:ROLENOTES/>
                            </cus:RELDELIVOPT>
                        </cus:sgRELDELIVOPT>
                    </cus:mRELATIONCODE>
                </cus:gRELATIONCODE>
                <cus:ACCOUNTOFFICER>%s</cus:ACCOUNTOFFICER>
                <cus:INDUSTRY>%s</cus:INDUSTRY>
                <cus:NATIONALITY>%s</cus:NATIONALITY>
                <cus:RESIDENCE>%s</cus:RESIDENCE>
                <cus:CONTACTDATE/>
                <cus:INTRODUCER/>
                <cus:gLEGALID g="1">
                    <cus:mLEGALID m="1">
                        <cus:LEGALID>%s</cus:LEGALID>
                        <cus:LEGALDOCNAME/>
                        <cus:LEGALHOLDERNAME>%s</cus:LEGALHOLDERNAME>
                        <cus:LEGALISSAUTH>%s</cus:LEGALISSAUTH>
                        <cus:LEGALISSDATE>%s</cus:LEGALISSDATE>
                        <cus:LEGALEXPDATE>%s</cus:LEGALEXPDATE>
                    </cus:mLEGALID>
                </cus:gLEGALID>
                <cus:BIRTHINCORPDATE/>
                <cus:CUSTOMERLIABILITY/>
                <cus:LANGUAGE>1</cus:LANGUAGE>
                <cus:gPOSTINGRESTRICT g="1">
                    <cus:POSTINGRESTRICT/>
                </cus:gPOSTINGRESTRICT>
                <cus:COMPANYBOOK/>
                <cus:CONFIDTXT/>
                <cus:ISSUECHEQUES/>
                <cus:gCUSTOMERRATING g="1">
                    <cus:CUSTOMERRATING/>
                </cus:gCUSTOMERRATING>
                <cus:NOUPDATECRM/>
                <cus:GIVENNAMES>%s</cus:GIVENNAMES>
                <cus:FAMILYNAME>%s</cus:FAMILYNAME>
                <cus:GENDER>%s</cus:GENDER>
                <cus:DATEOFBIRTH>%s</cus:DATEOFBIRTH>
                <cus:MARITALSTATUS>%s</cus:MARITALSTATUS>
                <cus:NOOFDEPENDENTS>0</cus:NOOFDEPENDENTS>
                <cus:SMS1>%s</cus:SMS1>
                <cus:EMAIL1>%s</cus:EMAIL1>
                <cus:gEMPLOYMENTSTATUS g="1">
                    <cus:mEMPLOYMENTSTATUS m="1">
                        <cus:EMPLOYMENTSTATUS>%s</cus:EMPLOYMENTSTATUS>
                        <cus:OCCUPATION>%s</cus:OCCUPATION>
                        <cus:JOBTITLE/>
                        <cus:EMPLOYERSNAME>%s</cus:EMPLOYERSNAME>
                        <cus:sgEMPLOYERSADD sg="1">
                            <cus:EMPLOYERSADD s="1">
                                <cus:EMPLOYERSADD>%s</cus:EMPLOYERSADD>
                            </cus:EMPLOYERSADD>
                        </cus:sgEMPLOYERSADD>
                        <cus:EMPLOYERSBUSS>%s</cus:EMPLOYERSBUSS>
                        <cus:EMPLOYMENTSTART/>
                        <cus:CUSTOMERCURRENCY>%s</cus:CUSTOMERCURRENCY>
                        <cus:SALARY>%s</cus:SALARY>
                        <cus:ANNUALBONUS>%s</cus:ANNUALBONUS>
                        <cus:SALARYDATEFREQ/>
                    </cus:mEMPLOYMENTSTATUS>
                </cus:gEMPLOYMENTSTATUS>
                <cus:NETMONTHLYIN>%s</cus:NETMONTHLYIN>
                <cus:NETMONTHLYOUT>%s</cus:NETMONTHLYOUT>
                <cus:gRESIDENCESTATUS g="1">
                    <cus:mRESIDENCESTATUS m="1">
                        <cus:RESIDENCESTATUS/>
                        <cus:RESIDENCETYPE/>
                        <cus:RESIDENCESINCE/>
                        <cus:RESIDENCEVALUE/>
                        <cus:MORTGAGEAMT/>
                    </cus:mRESIDENCESTATUS>
                </cus:gRESIDENCESTATUS>
                <cus:gOTHERFINREL g="1">
                    <cus:mOTHERFINREL m="1">
                        <cus:OTHERFINREL/>
                        <cus:OTHERFININST/>
                    </cus:mOTHERFINREL>
                </cus:gOTHERFINREL>
                <cus:gCOMMTYPE g="1">
                    <cus:mCOMMTYPE m="1">
                        <cus:COMMTYPE/>
                        <cus:PREFCHANNEL/>
                    </cus:mCOMMTYPE>
                </cus:gCOMMTYPE>
                <cus:ALLOWBULKPROCESS/>
                <cus:gINTERESTS g="1">
                    <cus:INTERESTS/>
                </cus:gINTERESTS>
                <cus:CUSTOMERSINCE/>
                <cus:CUSTOMERTYPE/>
                <cus:gPASTIMES g="1">
                    <cus:PASTIMES/>
                </cus:gPASTIMES>
                <cus:gFURTHERDETAILS g="1">
                    <cus:FURTHERDETAILS/>
                </cus:gFURTHERDETAILS>
                <cus:DOMICILE/>
                <cus:gOTHERNATIONALITY g="1">
                    <cus:OTHERNATIONALITY/>
                </cus:gOTHERNATIONALITY>
                <cus:CALCRISKCLASS/>
                <cus:MANUALRISKCLASS/>
                <cus:gOVERRIDEREASON g="1">
                    <cus:OVERRIDEREASON/>
                </cus:gOVERRIDEREASON>
                <cus:TinNumber>%s</cus:TinNumber>
                <cus:MotherName>%s</cus:MotherName>
                <cus:CustomerGroup>%s</cus:CustomerGroup>
                <cus:NationalId>%s</cus:NationalId>
            </CUSTOMERCREATEINDIVIDUALType>
        </iib:CustomerOpening>
    </soapenv:Body>
</soapenv:Envelope>
    `, param.Password, param.Username, menemoic, fullName, fullName, fullName, param.Address, param.Address, param.PostalCode, param.ISOCountryCode, param.AccountOffice, param.Industry, param.ISONationalityCode, param.ISOResidentCode, param.UniqueID, fullName, param.IssuesBy, param.IssuedDate, param.ExpiryDate, param.FirstName, param.LastName, param.Gender, param.DateOfBirth, param.MaritalStatus, param.PhoneNumber, param.Email, param.EmploymentStatus, param.Occupation, param.EmployerName, param.EmployerAddress, param.EmployerBusiness, param.CustomerCurrency, param.Salary, param.AnnualBonus, param.NetMonthlyIncome, param.NetMonthlyExpence, param.TinNumber, param.MotherName, param.CustomerGroup, param.UniqueID)
}

type Envelope struct {
	Body Body `xml:"Body"`
}

type Body struct {
	CustomerCreationResponse *CustomerCreationResponse `xml:"CustomerOpeningResponse"`
}

type CustomerCreationResponse struct {
	Status *struct {
		TransactionId string   `xml:"transactionId"`
		Success       string   `xml:"successIndicator"`
		Application   string   `xml:"application"`
		Messages      []string `xml:"messagesId"`
	} `xml:"Status"`
	CustomerType CustomerType `xml:"CUSTOMERType"`
}

type CustomerType struct {
	Menmonic   string `xml:"MNEMONIC"`
	GShortName struct {
		ShortName string `xml:"SHORTNAME"`
	} `xml:"gSHORTNAME"`
	GNameOne struct {
		NameOne string `xml:"NAME1"`
	} `xml:"gNAME1"`
	GNameTwo struct {
		NameTwo string `xml:"NAME2"`
	} `xml:"gNAME2"`
	GStreet struct {
		Street string `xml:"STREET"`
	} `xml:"gSTREET"`
	GLLAddress struct {
		MLLAddress struct {
			SGLLAddress struct {
				Address struct {
					Address string `xml:"ADDRESS"`
				} `xml:"ADDRESS"`
			} `xml:"sgLLADDRESS"`
		} `xml:"mLLADDRESS"`
	} `xml:"gLLADDRESS"`
	GTownCountry struct {
		TownCountry string `xml:"TOWNCOUNTRY"`
	} `xml:"gTOWNCOUNTRY"`
	GPostCode struct {
		PostCode string `xml:"POSTCODE"`
	} `xml:"gPOSTCODE"`
	GCountry struct {
		Country string `xml:"COUNTRY"`
	} `xml:"gCOUNTRY"`
	Sector         string `xml:"SECTOR"`
	AccountOfficer string `xml:"ACCOUNTOFFICER"`
	Industry       string `xml:"INDUSTRY"`
	Target         string `xml:"TARGET"`
	Nationality    string `xml:"NATIONALITY"`
	CustomerStatus string `xml:"CUSTOMERSTATUS"`
	Residence      string `xml:"RESIDENCE"`
	GLegalID       struct {
		MLegalID struct {
			LegalID         string `xml:"LEGALID"`
			LegalHolderName string `xml:"LEGALHOLDERNAME"`
			LegalIssAuth    string `xml:"LEGALISSAUTH"`
			LegalIssDate    string `xml:"LEGALISSDATE"`
			LegalExpDate    string `xml:"LEGALEXPDATE"`
		} `xml:"mLEGALID"`
	} `xml:"gLEGALID"`
	Language       string `xml:"LANGUAGE"`
	CompanyBook    string `xml:"COMPANYBOOK"`
	CLSCParty      string `xml:"CLSCPARTY"`
	GCRProfileType struct {
		MCRProfileType struct {
			CRProfileType string `xml:"CRPROFILETYPE"`
			CRProfile     string `xml:"CRPROFILE"`
		} `xml:"mCRPROFILETYPE"`
	} `xml:"gCRPROFILETYPE"`
	GivenNames     string `xml:"GIVENNAMES"`
	FamilyName     string `xml:"FAMILYNAME"`
	Gender         string `xml:"GENDER"`
	DateOfBirth    string `xml:"DATEOFBIRTH"`
	MaritalStatus  string `xml:"MARITALSTATUS"`
	NoOfDependents string `xml:"NOOFDEPENDENTS"`
	GPhoneOne      struct {
		MPhoneOne struct {
			SMSOne   string `xml:"SMS1"`
			EmailOne string `xml:"EMAIL1"`
		} `xml:"mPHONE1"`
	} `xml:"gPHONE1"`
	GEmploymentStatus struct {
		MEmploymentStatus struct {
			EmploymentStatus string `xml:"EMPLOYMENTSTATUS"`
			Occupation       string `xml:"OCCUPATION"`
			EmployersName    string `xml:"EMPLOYERSNAME"`
			SGEmployersAdd   struct {
				EmployersAdd struct {
					EmployersAdd string `xml:"EMPLOYERSADD"`
				} `xml:"EMPLOYERSADD"`
			} `xml:"sgEMPLOYERSADD"`
			EmployersBuss    string `xml:"EMPLOYERSBUSS"`
			CustomerCurrency string `xml:"CUSTOMERCURRENCY"`
			Salary           string `xml:"SALARY"`
			AnnualBonus      string `xml:"ANNUALBONUS"`
		} `xml:"mEMPLOYMENTSTATUS"`
	} `xml:"gEMPLOYMENTSTATUS"`
	NetMonthlyIn           string `xml:"NETMONTHLYIN"`
	NetMonthlyOut          string `xml:"NETMONTHLYOUT"`
	AMLCheck               string `xml:"AMLCHECK"`
	AMLResult              string `xml:"AMLRESULT"`
	InternetBankingService string `xml:"INTERNETBANKINGSERVICE"`
	MobileBankingService   string `xml:"MOBILEBANKINGSERVICE"`
	GCRUserProfileTy       struct {
		MCRUserProfileTy struct {
			CRUserProfileType string `xml:"CRUSERPROFILETYPE"`
			CRCalcProfile     string `xml:"CRCALCPROFILE"`
			CRUserProfile     string `xml:"CRUSERPROFILE"`
		} `xml:"mCRUSERPROFILETY"`
	} `xml:"gCRUSERPROFILETY"`
	Reserved01 string `xml:"RESERVED01"`
	GOverride  struct {
		Override []string `xml:"OVERRIDE"`
	} `xml:"gOVERRIDE"`
	RecordStatus string `xml:"RECORDSTATUS"`
	CurrNo       string `xml:"CURRNO"`
	GInputter    struct {
		Inputter string `xml:"INPUTTER"`
	} `xml:"gINPUTTER"`
	GDateTime struct {
		DateTime string `xml:"DATETIME"`
	} `xml:"gDATETIME"`
	Authoriser string `xml:"AUTHORISER"`
	CoCode     string `xml:"COCODE"`
	DeptCode   string `xml:"DEPTCODE"`
	Ownership  string `xml:"Ownership"`
	ETMBTinNo  string `xml:"ETMBTINNO"`
	CommPre    string `xml:"COMMPRE"`
	CustMother string `xml:"CUSTMOTHER"`
	DataCleanD string `xml:"DATACLEAND"`
	CustGruop  string `xml:"CUSTGRUOP"`
	NationalId string `xml:"NATIONALID"`
	CompVsInd  string `xml:"COMPVSIND"`
}

type CustomerTypeDetail struct {
	Menmonic         string
	FullName         string
	Address          string
	PostalCode       string
	Country          string
	AccountOfficer   string
	Industry         string
	Nationality      string
	IssuedDate       string
	ExpiryDate       string
	CompanyBook      string
	Gender           string
	DateOfBirth      string
	MaritalStatus    string
	PhoneNumber      string
	Email            string
	EmploymentStatus string
	Salary           string
	Customer         string
	AnnualBonus      string
	Currency         string
	TinNumber        string
	MotherName       string
	CustomerGroup    string
	Ownership        string
	NationalId       string
	Cocode           string
}

type CustomerCreationResult struct {
	Success  bool
	Detail   *CustomerTypeDetail
	Messages []string
}

func ParseCustomerCreationSOAP(xmlData string) (*CustomerCreationResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.CustomerCreationResponse != nil {
		resp := env.Body.CustomerCreationResponse
		if resp.Status == nil {
			return &CustomerCreationResult{
				Success:  false,
				Messages: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.Success) != "success" {
			return &CustomerCreationResult{
				Success:  false,
				Messages: resp.Status.Messages,
			}, nil
		}

		if resp.CustomerType.Menmonic == "" {
			return &CustomerCreationResult{
				Success:  true,
				Messages: []string{},
			}, nil
		}

		detail := &CustomerTypeDetail{
			Menmonic:         resp.CustomerType.Menmonic,
			FullName:         resp.CustomerType.GShortName.ShortName,
			Address:          resp.CustomerType.GLLAddress.MLLAddress.SGLLAddress.Address.Address,
			PostalCode:       resp.CustomerType.GPostCode.PostCode,
			Country:          resp.CustomerType.GCountry.Country,
			AccountOfficer:   resp.CustomerType.AccountOfficer,
			Industry:         resp.CustomerType.Industry,
			Nationality:      resp.CustomerType.Nationality,
			IssuedDate:       resp.CustomerType.GLegalID.MLegalID.LegalIssDate,
			ExpiryDate:       resp.CustomerType.GLegalID.MLegalID.LegalExpDate,
			CompanyBook:      resp.CustomerType.CompanyBook,
			Gender:           resp.CustomerType.Gender,
			DateOfBirth:      resp.CustomerType.DateOfBirth,
			MaritalStatus:    resp.CustomerType.MaritalStatus,
			PhoneNumber:      resp.CustomerType.GPhoneOne.MPhoneOne.SMSOne,
			Email:            resp.CustomerType.GPhoneOne.MPhoneOne.EmailOne,
			EmploymentStatus: resp.CustomerType.GEmploymentStatus.MEmploymentStatus.EmploymentStatus,
			Salary:           resp.CustomerType.GEmploymentStatus.MEmploymentStatus.Salary,
			Customer:         resp.CustomerType.GEmploymentStatus.MEmploymentStatus.EmployersName,
			AnnualBonus:      resp.CustomerType.GEmploymentStatus.MEmploymentStatus.AnnualBonus,
			Currency:         resp.CustomerType.GEmploymentStatus.MEmploymentStatus.CustomerCurrency,
			TinNumber:        resp.CustomerType.ETMBTinNo,
			MotherName:       resp.CustomerType.CustMother,
			CustomerGroup:    resp.CustomerType.CustGruop,
			Ownership:        resp.CustomerType.Ownership,
			NationalId:       resp.CustomerType.NationalId,
			Cocode:           resp.CustomerType.CoCode,
		}

		return &CustomerCreationResult{
			Success: true,
			Detail:  detail,
		}, nil
	}
	return &CustomerCreationResult{
		Success:  false,
		Messages: []string{"Invalid response type"},
	}, nil
}
