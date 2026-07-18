package customercreation

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username            string
	Password            string
	Company             string
	FirstName           string
	MiddleName          string
	LastName            string
	PhoneNumber         string
	Address             string
	PostalCode          string
	ISOCountryCode      string
	AccountOffice       string
	Industry            string
	ISONationalityCode  string
	ISOResidentCode     string
	UniqueID            string
	IssuesBy            string
	IssuedDate          string
	ExpiryDate          string
	Title               string
	Gender              string
	DateOfBirth         string
	MaritalStatus       string
	NoOfDependents      string
	Email               string
	EmploymentStatus    string
	Occupation          string
	EmployerName        string
	EmployerAddress     string
	EmployerBusiness    string
	CustomerCurrency    string
	Salary              string
	AnnualBonus         string
	NetMonthlyIncome    string
	NetMonthlyExpence   string
	TinNumber           string
	LegalDocumenetName  string
	CustomerOccupation  string
	EducationStatus     string
	MotherName          string
	FATCACompliant      string
	USPerson            string
	KebeleHNO           string
	CustomerSubSegment  string
	CustomerSegment     string
	GrandFatherName     string
	CustomerGroup       string
	Street              string
	NationalId          string
	TownCountry         string
	Menmonic            string
	Url                 string
	Header              map[string]string
}

type CreateCustomerParams struct {
	Company             string
	FirstName           string
	MiddleName          string
	LastName            string
	PhoneNumber         string
	Address             string
	PostalCode          string
	ISOCountryCode      string
	AccountOffice       string
	Industry            string
	ISONationalityCode  string
	ISOResidentCode     string
	UniqueID            string
	IssuesBy            string
	IssuedDate          string
	ExpiryDate          string
	Title               string
	Gender              string
	DateOfBirth         string
	MaritalStatus       string
	NoOfDependents      string
	Email               string
	EmploymentStatus    string
	Occupation          string
	EmployerName        string
	EmployerAddress     string
	EmployerBusiness    string
	LegalDocumenetName  string
	CustomerCurrency    string
	Salary              string
	AnnualBonus         string
	NetMonthlyIncome    string
	NetMonthlyExpence   string
	TownCountry         string
	TinNumber           string
	CustomerOccupation  string
	EducationStatus     string
	MotherName          string
	FATCACompliant      string
	USPerson            string
	KebeleHNO           string
	CustomerSubSegment  string
	CustomerSegment     string
	GrandFatherName     string
	CustomerGroup       string
	Street              string
	NationalId          string
	Menmonic            string
	Url                 string
	Header              map[string]string
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
	return fmt.Sprintf("%s%s", string(param.LastName[0]), lastNineDigit)
}

func NewCustomerCreation(param Params) string {
	menemoic := param.Menmonic
	fullName := FullName(param)
	grandFatherName := param.GrandFatherName
	if grandFatherName == "" {
		grandFatherName = param.LastName
	}
	noOfDependents := param.NoOfDependents
	if noOfDependents == "" {
		noOfDependents = "0"
	}

	return fmt.Sprintf(`
    <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:iib="http://temenos.com/IIBONBOARDING" xmlns:cus="http://temenos.com/CUSTOMERCREATEINDIVIDUAL">
    <soapenv:Header/>
    <soapenv:Body>
        <iib:CustomerOpening>
            <WebRequestCommon>
                <company>%s</company>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <OfsFunction>
                <noOfAuth>0</noOfAuth>
            </OfsFunction>
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
                    <cus:STREET>%s</cus:STREET>
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
                <cus:NATIONALITY>%s</cus:NATIONALITY>
                <cus:RESIDENCE>%s</cus:RESIDENCE>
                <cus:gLEGALID g="1">
                    <cus:mLEGALID m="1">
                        <cus:LEGALID>%s</cus:LEGALID>
                        <cus:LEGALDOCNAME>%s</cus:LEGALDOCNAME>
                        <cus:LEGALHOLDERNAME>%s</cus:LEGALHOLDERNAME>
                        <cus:LEGALISSAUTH>%s</cus:LEGALISSAUTH>
                        <cus:LEGALISSDATE>%s</cus:LEGALISSDATE>
                        <cus:LEGALEXPDATE>%s</cus:LEGALEXPDATE>
                    </cus:mLEGALID>
                </cus:gLEGALID>
                <cus:TITLE>%s</cus:TITLE>
                <cus:GIVENNAMES>%s</cus:GIVENNAMES>
                <cus:FAMILYNAME>%s</cus:FAMILYNAME>
                <cus:GENDER>%s</cus:GENDER>
                <cus:DATEOFBIRTH>%s</cus:DATEOFBIRTH>
                <cus:MARITALSTATUS>%s</cus:MARITALSTATUS>
                <cus:NOOFDEPENDENTS>%s</cus:NOOFDEPENDENTS>
                <cus:gPHONE1 g="1">
                    <cus:mPHONE1 m="1">
                        <cus:SMS1>%s</cus:SMS1>
                        <cus:EMAIL1>%s</cus:EMAIL1>
                    </cus:mPHONE1>
                </cus:gPHONE1>
                <cus:gEMPLOYMENTSTATUS g="1">
                    <cus:mEMPLOYMENTSTATUS m="1">
                        <cus:EMPLOYMENTSTATUS>%s</cus:EMPLOYMENTSTATUS>
                        <cus:OCCUPATION>%s</cus:OCCUPATION>
                        <cus:CUSTOMERCURRENCY>%s</cus:CUSTOMERCURRENCY>
                        <cus:SALARY>%s</cus:SALARY>
                    </cus:mEMPLOYMENTSTATUS>
                </cus:gEMPLOYMENTSTATUS>
                <cus:NETMONTHLYIN>%s</cus:NETMONTHLYIN>
                <cus:NETMONTHLYOUT>%s</cus:NETMONTHLYOUT>
                <cus:gTAXID g="1">
                    <cus:TAXID></cus:TAXID>
                </cus:gTAXID>
                <cus:gFORMERVISTYPE g="1">
                    <cus:FORMERVISTYPE></cus:FORMERVISTYPE>
                </cus:gFORMERVISTYPE>
                <cus:gRISKASSETTYPE g="1">
                    <cus:mRISKASSETTYPE m="1">
                        <cus:RISKASSETTYPE></cus:RISKASSETTYPE>
                        <cus:RISKLEVEL></cus:RISKLEVEL>
                        <cus:RISKTOLERANCE></cus:RISKTOLERANCE>
                        <cus:RISKFROMDATE></cus:RISKFROMDATE>
                    </cus:mRISKASSETTYPE>
                </cus:gRISKASSETTYPE>
                <cus:AMLCHECK></cus:AMLCHECK>
                <cus:AMLRESULT></cus:AMLRESULT>
                <cus:TinNumber>%s</cus:TinNumber>
                <cus:CustomerOccupation>%s</cus:CustomerOccupation>
                <cus:EduactionStatus>%s</cus:EduactionStatus>
                <cus:MotherName>%s</cus:MotherName>
                <cus:FATCACOMPLIANT>%s</cus:FATCACOMPLIANT>
                <cus:USPerson>%s</cus:USPerson>
                <cus:KebeleHNO>%s</cus:KebeleHNO>
                <cus:CustomerSubSegement>%s</cus:CustomerSubSegement>
                <cus:CustomerSegment>%s</cus:CustomerSegment>
                <cus:GrandFatherName>%s</cus:GrandFatherName>
                <cus:CustomerGroup>%s</cus:CustomerGroup>
                <cus:NationalId>%s</cus:NationalId>
            </CUSTOMERCREATEINDIVIDUALType>
        </iib:CustomerOpening>
    </soapenv:Body>
</soapenv:Envelope>
    `, param.Company, param.Password, param.Username, menemoic, fullName, fullName, fullName, param.Street, param.Address, param.TownCountry, param.PostalCode, param.ISOCountryCode, param.ISONationalityCode, param.ISOResidentCode, param.UniqueID, param.LegalDocumenetName, fullName, param.IssuesBy, param.IssuedDate, param.ExpiryDate, param.Title, param.FirstName, param.MiddleName, param.Gender, param.DateOfBirth, param.MaritalStatus, noOfDependents, param.PhoneNumber, param.Email, param.EmploymentStatus, param.Occupation, param.CustomerCurrency, param.Salary, param.NetMonthlyIncome, param.NetMonthlyExpence, param.TinNumber, param.CustomerOccupation, param.EducationStatus, param.MotherName, param.FATCACompliant, param.USPerson, param.KebeleHNO, param.CustomerSubSegment, param.CustomerSegment, grandFatherName, param.CustomerGroup, param.NationalId)
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
		Messages      []string `xml:"messages"`
	} `xml:"Status"`
	CustomerType CustomerType `xml:"CUSTOMERType"`
}

type CustomerType struct {
	XMLName        xml.Name `xml:"CUSTOMERType"`
	CustomerNumber string   `xml:"id,attr"`
	Menmonic       string   `xml:"MNEMONIC"`
	GShortName     struct {
		ShortName string `xml:"SHORTNAME"`
	} `xml:"gSHORTNAME"`
	GNameOne struct {
		NameOne string `xml:"NAME1"`
	} `xml:"gNAME1"`
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
			LegalDocName    string `xml:"LEGALDOCNAME"`
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
	Title          string `xml:"TITLE"`
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
			CustomerCurrency string `xml:"CUSTOMERCURRENCY"`
			Salary           string `xml:"SALARY"`
		} `xml:"mEMPLOYMENTSTATUS"`
	} `xml:"gEMPLOYMENTSTATUS"`
	CustomerStatusType     string `xml:"CUSTOMERTYPE"`
	AMLCheck               string `xml:"AMLCHECK"`
	AMLResult              string `xml:"AMLRESULT"`
	KYCComplete            string `xml:"KYCCOMPLETE"`
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
	Authoriser   string `xml:"AUTHORISER"`
	CoCode       string `xml:"COCODE"`
	DeptCode     string `xml:"DEPTCODE"`
	Ownership    string `xml:"Ownership"`
	GCorBanGroup struct {
		CorBanGroup []string `xml:"CORBANGROUP"`
	} `xml:"gCORBAN.GROUP"`
	CustOccupation     string `xml:"CUSTOCCUPATION"`
	CustEdu            string `xml:"CUSTEDU"`
	CommPre            string `xml:"COMMPRE"`
	CustMother         string `xml:"CUSTMOTHER"`
	DataCleanD         string `xml:"DATACLEAND"`
	FATCACompliant     string `xml:"FATCACOMPLIANT"`
	PEPStatus          string `xml:"PEPSTATUS"`
	USPerson           string `xml:"USPERSON"`
	HouseNo            string `xml:"HOUSENO"`
	CutSegment         string `xml:"CUTSEGEMENT"`
	MCustSegment       string `xml:"MCUSTSEGEMENT"`
	GrandFatherName    string `xml:"GFNAME"`
	CustGruop          string `xml:"CUSTGRUOP"`
	NationalId         string `xml:"NATIONALID"`
	CompVsInd          string `xml:"COMPVSIND"`
}

type CustomerTypeDetail struct {
	CustomerNumber     string
	Menmonic           string
	FullName           string
	Title              string
	GivenNames         string
	FamilyName         string
	Street             string
	Address            string
	TownCountry        string
	PostalCode         string
	Country            string
	AccountOfficer     string
	Industry           string
	Nationality        string
	LegalID            string
	LegalDocName       string
	IssuedDate         string
	ExpiryDate         string
	CompanyBook        string
	Gender             string
	DateOfBirth        string
	MaritalStatus      string
	NoOfDependents     string
	PhoneNumber        string
	Email              string
	EmploymentStatus   string
	Occupation         string
	Salary             string
	Currency           string
	CustomerType       string
	AMLCheck           string
	AMLResult          string
	KYCComplete        string
	CustomerOccupation string
	EducationStatus    string
	MotherName         string
	FATCACompliant     string
	PEPStatus          string
	USPerson           string
	KebeleHNO          string
	CustomerSegment    string
	CustomerSubSegment string
	GrandFatherName    string
	CustomerGroup      string
	Ownership          string
	NationalId         string
	Cocode             string
	AccountNumber      string
	Override           []string
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
				Messages: resp.Status.Messages,
			}, nil
		}

		detail := &CustomerTypeDetail{
			CustomerNumber:     resp.CustomerType.CustomerNumber,
			Menmonic:           resp.CustomerType.Menmonic,
			FullName:           resp.CustomerType.GShortName.ShortName,
			Title:              resp.CustomerType.Title,
			GivenNames:         resp.CustomerType.GivenNames,
			FamilyName:         resp.CustomerType.FamilyName,
			Street:             resp.CustomerType.GStreet.Street,
			Address:            resp.CustomerType.GLLAddress.MLLAddress.SGLLAddress.Address.Address,
			TownCountry:        resp.CustomerType.GTownCountry.TownCountry,
			PostalCode:         resp.CustomerType.GPostCode.PostCode,
			Country:            resp.CustomerType.GCountry.Country,
			AccountOfficer:     resp.CustomerType.AccountOfficer,
			Industry:           resp.CustomerType.Industry,
			Nationality:        resp.CustomerType.Nationality,
			LegalID:            resp.CustomerType.GLegalID.MLegalID.LegalID,
			LegalDocName:       resp.CustomerType.GLegalID.MLegalID.LegalDocName,
			IssuedDate:         resp.CustomerType.GLegalID.MLegalID.LegalIssDate,
			ExpiryDate:         resp.CustomerType.GLegalID.MLegalID.LegalExpDate,
			CompanyBook:        resp.CustomerType.CompanyBook,
			Gender:             resp.CustomerType.Gender,
			DateOfBirth:        resp.CustomerType.DateOfBirth,
			MaritalStatus:      resp.CustomerType.MaritalStatus,
			NoOfDependents:     resp.CustomerType.NoOfDependents,
			PhoneNumber:        resp.CustomerType.GPhoneOne.MPhoneOne.SMSOne,
			Email:              resp.CustomerType.GPhoneOne.MPhoneOne.EmailOne,
			EmploymentStatus:   resp.CustomerType.GEmploymentStatus.MEmploymentStatus.EmploymentStatus,
			Occupation:         resp.CustomerType.GEmploymentStatus.MEmploymentStatus.Occupation,
			Salary:             resp.CustomerType.GEmploymentStatus.MEmploymentStatus.Salary,
			Currency:           resp.CustomerType.GEmploymentStatus.MEmploymentStatus.CustomerCurrency,
			CustomerType:       resp.CustomerType.CustomerStatusType,
			AMLCheck:           resp.CustomerType.AMLCheck,
			AMLResult:          resp.CustomerType.AMLResult,
			KYCComplete:        resp.CustomerType.KYCComplete,
			CustomerOccupation: resp.CustomerType.CustOccupation,
			EducationStatus:    resp.CustomerType.CustEdu,
			MotherName:         resp.CustomerType.CustMother,
			FATCACompliant:     resp.CustomerType.FATCACompliant,
			PEPStatus:          resp.CustomerType.PEPStatus,
			USPerson:           resp.CustomerType.USPerson,
			KebeleHNO:          resp.CustomerType.HouseNo,
			CustomerSegment:    resp.CustomerType.MCustSegment,
			CustomerSubSegment: resp.CustomerType.CutSegment,
			GrandFatherName:    resp.CustomerType.GrandFatherName,
			CustomerGroup:      resp.CustomerType.CustGruop,
			Ownership:          resp.CustomerType.Ownership,
			NationalId:         resp.CustomerType.NationalId,
			Cocode:             resp.CustomerType.CoCode,
			Override:           resp.CustomerType.GOverride.Override,
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
