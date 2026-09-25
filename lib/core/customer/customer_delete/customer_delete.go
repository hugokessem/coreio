package customerdelete

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Param struct {
	Username       string
	Password       string
	CustomerNumber string
}

type CustomerDeleteParam struct {
	CustomerNumber string
}

func NewCustomerDelete(param Param) string {
	return fmt.Sprintf(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:iib="http://temenos.com/IIBONBOARDING">
	<soapenv:Header/>
	<soapenv:Body>
		<iib:DeleteCustomerIDSuperApp>
			<WebRequestCommon>
				<company/>
				<password>%s</password>
				<userName>%s</userName>
			</WebRequestCommon>
			<CUSTOMERCREATEINDIVIDUALType>
				<transactionId>%s</transactionId>
			</CUSTOMERCREATEINDIVIDUALType>
		</iib:DeleteCustomerIDSuperApp>
	</soapenv:Body>
	</soapenv:Envelope>`, param.Password, param.Username, param.CustomerNumber)
}

type Envelope struct {
	Body Body `xml:"Body"`
}

type Body struct {
	DeleteCustomerResponse *DeleteCustomerResponse `xml:"DeleteCustomerIDSuperAppResponse"`
	Fault                  *SOAPFault              `xml:"Fault"`
}

type SOAPFault struct {
	FaultCode   string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
}

type DeleteCustomerResponse struct {
	Status *struct {
		TransactionID    string   `xml:"transactionId"`
		MessageID        string   `xml:"messageId"`
		SuccessIndicator string   `xml:"successIndicator"`
		Application      string   `xml:"application"`
		Messages         []string `xml:"messages"`
	} `xml:"Status"`
	CustomerType *CustomerType `xml:"CUSTOMERType"`
}

type CustomerType struct {
	XMLName        xml.Name `xml:"CUSTOMERType"`
	CustomerNumber string   `xml:"id,attr"`
	Mnemonic       string   `xml:"MNEMONIC"`
	GShortName     struct {
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
	CustomerStatusType string `xml:"CUSTOMERTYPE"`
	GFurtherDetails    struct {
		FurtherDetails string `xml:"FURTHERDETAILS"`
	} `xml:"gFURTHERDETAILS"`
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
	CoCode       string `xml:"COCODE"`
	DeptCode     string `xml:"DEPTCODE"`
	Ownership    string `xml:"Ownership"`
	GCorBanGroup struct {
		CorBanGroup []string `xml:"CORBANGROUP"`
	} `xml:"gCORBAN.GROUP"`
	CustOccupation string `xml:"CUSTOCCUPATION"`
	CustEdu        string `xml:"CUSTEDU"`
	CommPre        string `xml:"COMMPRE"`
	CustMother     string `xml:"CUSTMOTHER"`
	DataCleanD     string `xml:"DATACLEAND"`
	FATCACompliant string `xml:"FATCACOMPLIANT"`
	PEPStatus      string `xml:"PEPSTATUS"`
	USPerson       string `xml:"USPERSON"`
	HouseNo        string `xml:"HOUSENO"`
	CutSegment     string `xml:"CUTSEGEMENT"`
	MCustSegment   string `xml:"MCUSTSEGEMENT"`
	CustGroup      string `xml:"CUSTGRUOP"`
	CompVsInd      string `xml:"COMPVSIND"`
	SalesPerson    string `xml:"SALESPERSON"`
	FaydaVerified  string `xml:"FAYDAVERIFIED"`
	USTinNo        string `xml:"USTINNO"`
}

type CustomerDeleteDetail struct {
	CustomerNumber         string
	TransactionID          string
	MessageID              string
	Application            string
	Mnemonic               string
	ShortName              string
	Name1                  string
	Name2                  string
	Street                 string
	Address                string
	TownCountry            string
	PostalCode             string
	Country                string
	Sector                 string
	AccountOfficer         string
	Industry               string
	Target                 string
	Nationality            string
	CustomerStatus         string
	Residence              string
	LegalID                string
	LegalDocName           string
	LegalHolderName        string
	LegalIssAuth           string
	LegalIssDate           string
	LegalExpDate           string
	Language               string
	CompanyBook            string
	CLSCParty              string
	CRProfileType          string
	CRProfile              string
	GivenNames             string
	FamilyName             string
	Gender                 string
	DateOfBirth            string
	MaritalStatus          string
	NoOfDependents         string
	PhoneNumber            string
	Email                  string
	EmploymentStatus       string
	Occupation             string
	CustomerCurrency       string
	Salary                 string
	CustomerType           string
	FurtherDetails         string
	AMLCheck               string
	AMLResult              string
	KYCComplete            string
	InternetBankingService string
	MobileBankingService   string
	CRUserProfileType      string
	CRCalcProfile          string
	CRUserProfile          string
	Reserved01             string
	Override               []string
	RecordStatus           string
	CurrNo                 string
	Inputter               string
	DateTime               string
	CoCode                 string
	DeptCode               string
	Ownership              string
	CorBanGroup            []string
	CustomerOccupation     string
	EducationStatus        string
	CommunicationPref      string
	MotherName             string
	DataCleaned            string
	FATCACompliant         string
	PEPStatus              string
	USPerson               string
	HouseNo                string
	CustomerSubSegment     string
	CustomerSegment        string
	CustomerGroup          string
	CompVsInd              string
	SalesPerson            string
	FaydaVerified          string
	USTinNo                string
}

type CustomerDeleteResult struct {
	Success  bool
	Detail   *CustomerDeleteDetail
	Messages []string
}

func ParseCustomerDeleteSOAP(xmlData string) (*CustomerDeleteResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.Fault != nil {
		message := strings.TrimSpace(env.Body.Fault.FaultString)
		if message == "" {
			message = strings.TrimSpace(env.Body.Fault.FaultCode)
		}
		if message == "" {
			message = "SOAP Fault"
		}
		return &CustomerDeleteResult{
			Success:  false,
			Messages: []string{message},
		}, nil
	}

	if env.Body.DeleteCustomerResponse == nil {
		return &CustomerDeleteResult{
			Success:  false,
			Messages: []string{"Invalid response type"},
		}, nil
	}

	resp := env.Body.DeleteCustomerResponse
	if resp.Status == nil {
		return &CustomerDeleteResult{
			Success:  false,
			Messages: []string{"Missing Status"},
		}, nil
	}

	// Record not found / unauthorized missing record example:
	// successIndicator=T24Error, messages=@ID:1:1=UNAUTH. RECORD MISSING
	if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
		messages := resp.Status.Messages
		if len(messages) == 0 {
			messages = []string{"Customer delete failed"}
		}
		return &CustomerDeleteResult{
			Success:  false,
			Messages: messages,
		}, nil
	}

	detail := &CustomerDeleteDetail{
		TransactionID: resp.Status.TransactionID,
		MessageID:     resp.Status.MessageID,
		Application:   resp.Status.Application,
	}

	if resp.CustomerType != nil {
		ct := resp.CustomerType
		detail.CustomerNumber = ct.CustomerNumber
		detail.Mnemonic = ct.Mnemonic
		detail.ShortName = ct.GShortName.ShortName
		detail.Name1 = ct.GNameOne.NameOne
		detail.Name2 = ct.GNameTwo.NameTwo
		detail.Street = ct.GStreet.Street
		detail.Address = ct.GLLAddress.MLLAddress.SGLLAddress.Address.Address
		detail.TownCountry = ct.GTownCountry.TownCountry
		detail.PostalCode = ct.GPostCode.PostCode
		detail.Country = ct.GCountry.Country
		detail.Sector = ct.Sector
		detail.AccountOfficer = ct.AccountOfficer
		detail.Industry = ct.Industry
		detail.Target = ct.Target
		detail.Nationality = ct.Nationality
		detail.CustomerStatus = ct.CustomerStatus
		detail.Residence = ct.Residence
		detail.LegalID = ct.GLegalID.MLegalID.LegalID
		detail.LegalDocName = ct.GLegalID.MLegalID.LegalDocName
		detail.LegalHolderName = ct.GLegalID.MLegalID.LegalHolderName
		detail.LegalIssAuth = ct.GLegalID.MLegalID.LegalIssAuth
		detail.LegalIssDate = ct.GLegalID.MLegalID.LegalIssDate
		detail.LegalExpDate = ct.GLegalID.MLegalID.LegalExpDate
		detail.Language = ct.Language
		detail.CompanyBook = ct.CompanyBook
		detail.CLSCParty = ct.CLSCParty
		detail.CRProfileType = ct.GCRProfileType.MCRProfileType.CRProfileType
		detail.CRProfile = ct.GCRProfileType.MCRProfileType.CRProfile
		detail.GivenNames = ct.GivenNames
		detail.FamilyName = ct.FamilyName
		detail.Gender = ct.Gender
		detail.DateOfBirth = ct.DateOfBirth
		detail.MaritalStatus = ct.MaritalStatus
		detail.NoOfDependents = ct.NoOfDependents
		detail.PhoneNumber = ct.GPhoneOne.MPhoneOne.SMSOne
		detail.Email = ct.GPhoneOne.MPhoneOne.EmailOne
		detail.EmploymentStatus = ct.GEmploymentStatus.MEmploymentStatus.EmploymentStatus
		detail.Occupation = ct.GEmploymentStatus.MEmploymentStatus.Occupation
		detail.CustomerCurrency = ct.GEmploymentStatus.MEmploymentStatus.CustomerCurrency
		detail.Salary = ct.GEmploymentStatus.MEmploymentStatus.Salary
		detail.CustomerType = ct.CustomerStatusType
		detail.FurtherDetails = ct.GFurtherDetails.FurtherDetails
		detail.AMLCheck = ct.AMLCheck
		detail.AMLResult = ct.AMLResult
		detail.KYCComplete = ct.KYCComplete
		detail.InternetBankingService = ct.InternetBankingService
		detail.MobileBankingService = ct.MobileBankingService
		detail.CRUserProfileType = ct.GCRUserProfileTy.MCRUserProfileTy.CRUserProfileType
		detail.CRCalcProfile = ct.GCRUserProfileTy.MCRUserProfileTy.CRCalcProfile
		detail.CRUserProfile = ct.GCRUserProfileTy.MCRUserProfileTy.CRUserProfile
		detail.Reserved01 = ct.Reserved01
		detail.Override = ct.GOverride.Override
		detail.RecordStatus = ct.RecordStatus
		detail.CurrNo = ct.CurrNo
		detail.Inputter = ct.GInputter.Inputter
		detail.DateTime = ct.GDateTime.DateTime
		detail.CoCode = ct.CoCode
		detail.DeptCode = ct.DeptCode
		detail.Ownership = ct.Ownership
		detail.CorBanGroup = ct.GCorBanGroup.CorBanGroup
		detail.CustomerOccupation = ct.CustOccupation
		detail.EducationStatus = ct.CustEdu
		detail.CommunicationPref = ct.CommPre
		detail.MotherName = ct.CustMother
		detail.DataCleaned = ct.DataCleanD
		detail.FATCACompliant = ct.FATCACompliant
		detail.PEPStatus = ct.PEPStatus
		detail.USPerson = ct.USPerson
		detail.HouseNo = ct.HouseNo
		detail.CustomerSubSegment = ct.CutSegment
		detail.CustomerSegment = ct.MCustSegment
		detail.CustomerGroup = ct.CustGroup
		detail.CompVsInd = ct.CompVsInd
		detail.SalesPerson = ct.SalesPerson
		detail.FaydaVerified = ct.FaydaVerified
		detail.USTinNo = ct.USTinNo

		if detail.CustomerNumber == "" {
			detail.CustomerNumber = resp.Status.TransactionID
		}
	}

	return &CustomerDeleteResult{
		Success: true,
		Detail:  detail,
	}, nil
}
