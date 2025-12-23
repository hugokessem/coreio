package customercreation

import "fmt"

type Params struct {
	Username         string
	Password         string
	Minmonic         string
	ShortName        string
	Name1            string
	Name2            string
	Street           string
	Address          string
	TownCountry      string
	Postcode         string
	Country          string
	RelationCode     string
	AccountOfficer   string
	Industry         string
	Nationality      string
	Residence        string
	LegalId          string
	LegalHolderName  string
	LegalIssAuth     string
	LegalIssDate     string
	LegalExpDate     string
	Language         string
	GivenNames       string
	FamilyName       string
	Gender           string
	DateOfBirth      string
	MaritalStatus    string
	NoOfDependents   string
	Sms1             string
	Email1           string
	EmploymentStatus string
	Occupation       string
	EmployersName    string
	EmployersAdd     string
	EmployersBus     string
	EmployersStart   string
	CustomerCurrency string
	Salary           string
	AnnualBonus      string
	NetMonthlyIn     string
	NetMonthlyOut    string
	TinNumber        string
	MotherName       string
	CustomerGroup    string
	NationalId       string
}

func NewCustomerCreation(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
xmlns:cbes="http://temenos.com/CBESUPERAPP"
xmlns:cus="http://temenos.com/CUSTOMERCREATEINDIVIDUAL">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:CustomerOpening>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <OfsFunction/>
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
                <cus:LANGUAGE>%s</cus:LANGUAGE>
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
                <cus:NOOFDEPENDENTS>%s</cus:NOOFDEPENDENTS>
                <cus:SMS1>+%s</cus:SMS1>
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
        </cbes:CustomerOpening>
    </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.Minmonic, param.ShortName, param.Name1, param.Name2, param.Street, param.Address, param.TownCountry, param.Postcode, param.Country, param.AccountOfficer, param.Industry, param.Nationality, param.Residence, param.LegalId, param.LegalHolderName, param.LegalIssAuth, param.LegalIssDate, param.LegalExpDate, param.Language, param.GivenNames, param.FamilyName, param.Gender, param.DateOfBirth, param.MaritalStatus, param.NoOfDependents, param.Sms1, param.Email1, param.EmploymentStatus, param.Occupation, param.EmployersName, param.EmployersAdd, param.EmployersBus, param.CustomerCurrency, param.Salary, param.AnnualBonus, param.NetMonthlyIn, param.NetMonthlyOut, param.TinNumber, param.MotherName, param.CustomerGroup, param.NationalId)
}
