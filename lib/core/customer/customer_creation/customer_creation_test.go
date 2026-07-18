package customercreation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const sampleCustomerCreationResponse = `<?xml version="1.0" encoding="UTF-8"?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns6:CustomerOpeningResponse xmlns:ns2="http://temenos.com/CUSTOMERCREATEINDIVIDUAL" xmlns:ns3="http://temenos.com/CUSTOMER" xmlns:ns4="http://temenos.com/ACCOUNTCREATEINDIVIDUAL" xmlns:ns5="http://temenos.com/ACCOUNT" xmlns:ns6="http://temenos.com/IIBONBOARDING">
            <Status>
                <transactionId>1202724929</transactionId>
                <messageId></messageId>
                <successIndicator>Success</successIndicator>
                <application>CUSTOMER</application>
            </Status>
            <CUSTOMERType id="1202724929">
                <ns3:MNEMONIC>M913323635</ns3:MNEMONIC>
                <ns3:gSHORTNAME>
                    <ns3:SHORTNAME>Mr Melese Tesfaye Kifle</ns3:SHORTNAME>
                </ns3:gSHORTNAME>
                <ns3:gNAME1>
                    <ns3:NAME1>Mr Melese Tesfaye Kifle</ns3:NAME1>
                </ns3:gNAME1>
                <ns3:gSTREET>
                    <ns3:STREET>AM</ns3:STREET>
                </ns3:gSTREET>
                <ns3:gLLADDRESS>
                    <ns3:mLLADDRESS>
                        <ns3:sgLLADDRESS>
                            <ns3:ADDRESS>
                                <ns3:ADDRESS>ADDIS ABABA</ns3:ADDRESS>
                            </ns3:ADDRESS>
                        </ns3:sgLLADDRESS>
                    </ns3:mLLADDRESS>
                </ns3:gLLADDRESS>
                <ns3:gTOWNCOUNTRY>
                    <ns3:TOWNCOUNTRY>ADDIS ABABA</ns3:TOWNCOUNTRY>
                </ns3:gTOWNCOUNTRY>
                <ns3:gPOSTCODE>
                    <ns3:POSTCODE>4125</ns3:POSTCODE>
                </ns3:gPOSTCODE>
                <ns3:gCOUNTRY>
                    <ns3:COUNTRY>ET</ns3:COUNTRY>
                </ns3:gCOUNTRY>
                <ns3:SECTOR>1000</ns3:SECTOR>
                <ns3:ACCOUNTOFFICER>7858</ns3:ACCOUNTOFFICER>
                <ns3:INDUSTRY>1499</ns3:INDUSTRY>
                <ns3:TARGET>4</ns3:TARGET>
                <ns3:NATIONALITY>ET</ns3:NATIONALITY>
                <ns3:CUSTOMERSTATUS>1</ns3:CUSTOMERSTATUS>
                <ns3:RESIDENCE>ET</ns3:RESIDENCE>
                <ns3:gLEGALID>
                    <ns3:mLEGALID>
                        <ns3:LEGALID>WS55651461632DSS</ns3:LEGALID>
                        <ns3:LEGALDOCNAME>NATIONAL.ID</ns3:LEGALDOCNAME>
                        <ns3:LEGALHOLDERNAME>MELESE TESFAYE KIFLE</ns3:LEGALHOLDERNAME>
                        <ns3:LEGALISSAUTH>FAYDA</ns3:LEGALISSAUTH>
                        <ns3:LEGALISSDATE>20211209</ns3:LEGALISSDATE>
                        <ns3:LEGALEXPDATE>20270808</ns3:LEGALEXPDATE>
                    </ns3:mLEGALID>
                </ns3:gLEGALID>
                <ns3:LANGUAGE>1</ns3:LANGUAGE>
                <ns3:COMPANYBOOK>ET0011859</ns3:COMPANYBOOK>
                <ns3:CLSCPARTY>NO</ns3:CLSCPARTY>
                <ns3:gCRPROFILETYPE>
                    <ns3:mCRPROFILETYPE>
                        <ns3:CRPROFILETYPE>VALUED.CUSTOMER</ns3:CRPROFILETYPE>
                        <ns3:CRPROFILE>14</ns3:CRPROFILE>
                    </ns3:mCRPROFILETYPE>
                </ns3:gCRPROFILETYPE>
                <ns3:TITLE>MR</ns3:TITLE>
                <ns3:GIVENNAMES>Melese</ns3:GIVENNAMES>
                <ns3:FAMILYNAME>Tesfaye</ns3:FAMILYNAME>
                <ns3:GENDER>MALE</ns3:GENDER>
                <ns3:DATEOFBIRTH>19910215</ns3:DATEOFBIRTH>
                <ns3:MARITALSTATUS>MARRIED</ns3:MARITALSTATUS>
                <ns3:NOOFDEPENDENTS>1</ns3:NOOFDEPENDENTS>
                <ns3:gPHONE1>
                    <ns3:mPHONE1>
                        <ns3:SMS1>+251913323635</ns3:SMS1>
                        <ns3:EMAIL1>sampleTest@gmail.com</ns3:EMAIL1>
                    </ns3:mPHONE1>
                </ns3:gPHONE1>
                <ns3:gEMPLOYMENTSTATUS>
                    <ns3:mEMPLOYMENTSTATUS>
                        <ns3:EMPLOYMENTSTATUS>EMPLOYED</ns3:EMPLOYMENTSTATUS>
                        <ns3:OCCUPATION>ACCOUNTANT</ns3:OCCUPATION>
                        <ns3:CUSTOMERCURRENCY>ETB</ns3:CUSTOMERCURRENCY>
                        <ns3:SALARY>28000.00</ns3:SALARY>
                    </ns3:mEMPLOYMENTSTATUS>
                </ns3:gEMPLOYMENTSTATUS>
                <ns3:CUSTOMERTYPE>ACTIVE</ns3:CUSTOMERTYPE>
                <ns3:AMLCHECK>NULL</ns3:AMLCHECK>
                <ns3:AMLRESULT>NULL</ns3:AMLRESULT>
                <ns3:KYCCOMPLETE>YES</ns3:KYCCOMPLETE>
                <ns3:INTERNETBANKINGSERVICE>NULL</ns3:INTERNETBANKINGSERVICE>
                <ns3:MOBILEBANKINGSERVICE>NULL</ns3:MOBILEBANKINGSERVICE>
                <ns3:gCRUSERPROFILETY>
                    <ns3:mCRUSERPROFILETY>
                        <ns3:CRUSERPROFILETYPE>VALUED.CUSTOMER</ns3:CRUSERPROFILETYPE>
                        <ns3:CRCALCPROFILE>14</ns3:CRCALCPROFILE>
                        <ns3:CRUSERPROFILE>14</ns3:CRUSERPROFILE>
                    </ns3:mCRUSERPROFILETY>
                </ns3:gCRUSERPROFILETY>
                <ns3:RESERVED01>NO</ns3:RESERVED01>
                <ns3:gOVERRIDE>
                    <ns3:OVERRIDE>VL-VL.CONT.SENT.AML}Contract sent to AML{{{{{{{{{1</ns3:OVERRIDE>
                    <ns3:OVERRIDE>KEBELEID/CUS*100 FROM 1202724929 NOT RECEIVED</ns3:OVERRIDE>
                </ns3:gOVERRIDE>
                <ns3:RECORDSTATUS>INAO</ns3:RECORDSTATUS>
                <ns3:CURRNO>1</ns3:CURRNO>
                <ns3:gINPUTTER>
                    <ns3:INPUTTER>54603_SUPERAPP.1__OFS_GCS</ns3:INPUTTER>
                </ns3:gINPUTTER>
                <ns3:gDATETIME>
                    <ns3:DATETIME>2607181409</ns3:DATETIME>
                </ns3:gDATETIME>
                <ns3:AUTHORISER>54603_SUPERAPP.1</ns3:AUTHORISER>
                <ns3:COCODE>ET0010001</ns3:COCODE>
                <ns3:DEPTCODE>1</ns3:DEPTCODE>
                <ns3:Ownership>3000</ns3:Ownership>
                <ns3:gCORBAN.GROUP>
                    <ns3:CORBANGROUP>AC.ALERTS</ns3:CORBANGROUP>
                    <ns3:CORBANGROUP>FT.ALERTS</ns3:CORBANGROUP>
                    <ns3:CORBANGROUP>TT.ALERTS</ns3:CORBANGROUP>
                </ns3:gCORBAN.GROUP>
                <ns3:CUSTOCCUPATION>Banker</ns3:CUSTOCCUPATION>
                <ns3:CUSTEDU>First Degree</ns3:CUSTEDU>
                <ns3:COMMPRE>SMS</ns3:COMMPRE>
                <ns3:CUSTMOTHER>Tigist Alemu </ns3:CUSTMOTHER>
                <ns3:DATACLEAND>YES</ns3:DATACLEAND>
                <ns3:FATCACOMPLIANT>NO</ns3:FATCACOMPLIANT>
                <ns3:PEPSTATUS>NO</ns3:PEPSTATUS>
                <ns3:USPERSON>NO</ns3:USPERSON>
                <ns3:HOUSENO>1544/02</ns3:HOUSENO>
                <ns3:CUTSEGEMENT>MASS</ns3:CUTSEGEMENT>
                <ns3:MCUSTSEGEMENT>MASS</ns3:MCUSTSEGEMENT>
                <ns3:GFNAME>Kifle</ns3:GFNAME>
                <ns3:CUSTGRUOP>RETAIL</ns3:CUSTGRUOP>
                <ns3:NATIONALID>351913961480002346398580898992284483</ns3:NATIONALID>
                <ns3:COMPVSIND>INDIVIDUAL</ns3:COMPVSIND>
            </CUSTOMERType>
        </ns6:CustomerOpeningResponse>
    </S:Body>
</S:Envelope>`

func TestParseCustomerCreationSOAP_Success(t *testing.T) {
	result, err := ParseCustomerCreationSOAP(sampleCustomerCreationResponse)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	detail := result.Detail
	assert.Equal(t, "1202724929", detail.CustomerNumber)
	assert.Equal(t, "M913323635", detail.Menmonic)
	assert.Equal(t, "Mr Melese Tesfaye Kifle", detail.FullName)
	assert.Equal(t, "MR", detail.Title)
	assert.Equal(t, "Melese", detail.GivenNames)
	assert.Equal(t, "Tesfaye", detail.FamilyName)
	assert.Equal(t, "AM", detail.Street)
	assert.Equal(t, "ADDIS ABABA", detail.Address)
	assert.Equal(t, "ADDIS ABABA", detail.TownCountry)
	assert.Equal(t, "4125", detail.PostalCode)
	assert.Equal(t, "ET", detail.Country)
	assert.Equal(t, "7858", detail.AccountOfficer)
	assert.Equal(t, "1499", detail.Industry)
	assert.Equal(t, "ET", detail.Nationality)
	assert.Equal(t, "WS55651461632DSS", detail.LegalID)
	assert.Equal(t, "NATIONAL.ID", detail.LegalDocName)
	assert.Equal(t, "20211209", detail.IssuedDate)
	assert.Equal(t, "20270808", detail.ExpiryDate)
	assert.Equal(t, "ET0011859", detail.CompanyBook)
	assert.Equal(t, "MALE", detail.Gender)
	assert.Equal(t, "19910215", detail.DateOfBirth)
	assert.Equal(t, "MARRIED", detail.MaritalStatus)
	assert.Equal(t, "1", detail.NoOfDependents)
	assert.Equal(t, "+251913323635", detail.PhoneNumber)
	assert.Equal(t, "sampleTest@gmail.com", detail.Email)
	assert.Equal(t, "EMPLOYED", detail.EmploymentStatus)
	assert.Equal(t, "ACCOUNTANT", detail.Occupation)
	assert.Equal(t, "28000.00", detail.Salary)
	assert.Equal(t, "ETB", detail.Currency)
	assert.Equal(t, "ACTIVE", detail.CustomerType)
	assert.Equal(t, "YES", detail.KYCComplete)
	assert.Equal(t, "Banker", detail.CustomerOccupation)
	assert.Equal(t, "First Degree", detail.EducationStatus)
	assert.Equal(t, "Tigist Alemu ", detail.MotherName)
	assert.Equal(t, "NO", detail.FATCACompliant)
	assert.Equal(t, "NO", detail.PEPStatus)
	assert.Equal(t, "NO", detail.USPerson)
	assert.Equal(t, "1544/02", detail.KebeleHNO)
	assert.Equal(t, "MASS", detail.CustomerSubSegment)
	assert.Equal(t, "MASS", detail.CustomerSegment)
	assert.Equal(t, "Kifle", detail.GrandFatherName)
	assert.Equal(t, "RETAIL", detail.CustomerGroup)
	assert.Equal(t, "351913961480002346398580898992284483", detail.NationalId)
	assert.Equal(t, "3000", detail.Ownership)
	assert.Equal(t, "ET0010001", detail.Cocode)
	assert.Len(t, detail.Override, 2)
}

func TestParseCustomerCreationSOAP_FailureMessages(t *testing.T) {
	xmlData := `<?xml version="1.0" encoding="UTF-8"?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns6:CustomerOpeningResponse xmlns:ns6="http://temenos.com/IIBONBOARDING">
            <Status>
                <transactionId>1376589572</transactionId>
                <messageId></messageId>
                <successIndicator>T24Error</successIndicator>
                <application>CUSTOMER</application>
                <messages>NATIONAL.ID:1:1=National ID is duplicated for customer 1069797513</messages>
            </Status>
        </ns6:CustomerOpeningResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseCustomerCreationSOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Nil(t, result.Detail)
	assert.Contains(t, result.Messages, "NATIONAL.ID:1:1=National ID is duplicated for customer 1069797513")
}

func TestNewCustomerCreation_SampleRequest(t *testing.T) {
	params := Params{
		Username:           "SUPERAPP",
		Password:           "123456",
		Company:            "ET0011859",
		FirstName:          "MELESE",
		MiddleName:         "TESFAYE",
		LastName:           "KIFLE",
		Menmonic:           "M913323635",
		PhoneNumber:        "+251913323635",
		Address:            "ADDIS ABABA",
		PostalCode:         "4125",
		ISOCountryCode:     "ET",
		ISONationalityCode: "ET",
		ISOResidentCode:    "ET",
		UniqueID:           "WS55651461632DSS",
		LegalDocumenetName: "NATIONAL.ID",
		IssuesBy:           "FAYDA",
		IssuedDate:         "20211209",
		ExpiryDate:         "20270808",
		Title:              "MR",
		Gender:             "MALE",
		DateOfBirth:        "19910215",
		MaritalStatus:      "MARRIED",
		NoOfDependents:     "1",
		Email:              "sampleTest@gmail.com",
		EmploymentStatus:   "EMPLOYED",
		Occupation:         "ACCOUNTANT",
		CustomerCurrency:   "ETB",
		Salary:             "28000",
		Street:             "AM",
		TownCountry:        "ADDIS ABABA",
		CustomerOccupation: "Banker",
		EducationStatus:    "First Degree",
		MotherName:         "TIGIST ALEMU",
		FATCACompliant:     "NO",
		USPerson:           "NO",
		KebeleHNO:          "1544/02",
		CustomerSubSegment: "MASS",
		CustomerSegment:    "MASS",
		GrandFatherName:    "KIFLE",
		CustomerGroup:      "RETAIL",
		NationalId:         "351913961480002346398580898992284483",
	}

	xmlRequest := NewCustomerCreation(params)

	assert.Contains(t, xmlRequest, `<company>ET0011859</company>`)
	assert.Contains(t, xmlRequest, `<noOfAuth>0</noOfAuth>`)
	assert.Contains(t, xmlRequest, "<cus:MNEMONIC>M913323635</cus:MNEMONIC>")
	assert.Contains(t, xmlRequest, "<cus:SHORTNAME>MELESE TESFAYE KIFLE</cus:SHORTNAME>")
	assert.Contains(t, xmlRequest, "<cus:LEGALID>WS55651461632DSS</cus:LEGALID>")
	assert.Contains(t, xmlRequest, "<cus:LEGALDOCNAME>NATIONAL.ID</cus:LEGALDOCNAME>")
	assert.Contains(t, xmlRequest, "<cus:TITLE>MR</cus:TITLE>")
	assert.Contains(t, xmlRequest, "<cus:GIVENNAMES>MELESE</cus:GIVENNAMES>")
	assert.Contains(t, xmlRequest, "<cus:FAMILYNAME>TESFAYE</cus:FAMILYNAME>")
	assert.Contains(t, xmlRequest, "<cus:NOOFDEPENDENTS>1</cus:NOOFDEPENDENTS>")
	assert.Contains(t, xmlRequest, "<cus:gPHONE1 g=\"1\">")
	assert.Contains(t, xmlRequest, "<cus:SMS1>+251913323635</cus:SMS1>")
	assert.Contains(t, xmlRequest, "<cus:EMAIL1>sampleTest@gmail.com</cus:EMAIL1>")
	assert.Contains(t, xmlRequest, "<cus:CustomerOccupation>Banker</cus:CustomerOccupation>")
	assert.Contains(t, xmlRequest, "<cus:EduactionStatus>First Degree</cus:EduactionStatus>")
	assert.Contains(t, xmlRequest, "<cus:FATCACOMPLIANT>NO</cus:FATCACOMPLIANT>")
	assert.Contains(t, xmlRequest, "<cus:USPerson>NO</cus:USPerson>")
	assert.Contains(t, xmlRequest, "<cus:KebeleHNO>1544/02</cus:KebeleHNO>")
	assert.Contains(t, xmlRequest, "<cus:CustomerSubSegement>MASS</cus:CustomerSubSegement>")
	assert.Contains(t, xmlRequest, "<cus:CustomerSegment>MASS</cus:CustomerSegment>")
	assert.Contains(t, xmlRequest, "<cus:GrandFatherName>KIFLE</cus:GrandFatherName>")
	assert.Contains(t, xmlRequest, "<cus:CustomerGroup>RETAIL</cus:CustomerGroup>")
	assert.Contains(t, xmlRequest, "<cus:NationalId>351913961480002346398580898992284483</cus:NationalId>")
	assert.NotContains(t, xmlRequest, "<cus:ACCOUNTOFFICER>")
	assert.NotContains(t, xmlRequest, "<cus:gRELATIONCODE")
}
