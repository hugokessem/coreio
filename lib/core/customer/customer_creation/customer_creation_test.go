package customercreation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const sampleCustomerCreationResponse = `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns6:CustomerOpeningResponse xmlns:ns2="http://temenos.com/CUSTOMERCREATEINDIVIDUAL" xmlns:ns3="http://temenos.com/CUSTOMER" xmlns:ns4="http://temenos.com/ACCOUNTCREATEINDIVIDUAL" xmlns:ns5="http://temenos.com/ACCOUNT" xmlns:ns6="http://temenos.com/IIBONBOARDING">
            <Status>
                <transactionId>1195875233</transactionId>
                <messageId></messageId>
                <successIndicator>Success</successIndicator>
                <application>CUSTOMER</application>
            </Status>
            <CUSTOMERType id="1195875233">
                <ns3:MNEMONIC>K912895609</ns3:MNEMONIC>
                <ns3:gSHORTNAME>
                    <ns3:SHORTNAME>KETEM HAILU TAYE</ns3:SHORTNAME>
                </ns3:gSHORTNAME>
                <ns3:gNAME1>
                    <ns3:NAME1>KETEM HAILU TAYE</ns3:NAME1>
                </ns3:gNAME1>
                <ns3:gNAME2>
                    <ns3:NAME2>KETEM HAILU TAYE</ns3:NAME2>
                </ns3:gNAME2>
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
                    <ns3:POSTCODE>4144</ns3:POSTCODE>
                </ns3:gPOSTCODE>
                <ns3:gCOUNTRY>
                    <ns3:COUNTRY>ET</ns3:COUNTRY>
                </ns3:gCOUNTRY>
                <ns3:SECTOR>1000</ns3:SECTOR>
                <ns3:ACCOUNTOFFICER>7124</ns3:ACCOUNTOFFICER>
                <ns3:INDUSTRY>1499</ns3:INDUSTRY>
                <ns3:TARGET>4</ns3:TARGET>
                <ns3:NATIONALITY>ET</ns3:NATIONALITY>
                <ns3:CUSTOMERSTATUS>1</ns3:CUSTOMERSTATUS>
                <ns3:RESIDENCE>ET</ns3:RESIDENCE>
                <ns3:gLEGALID>
                    <ns3:mLEGALID>
                        <ns3:LEGALID>123456779</ns3:LEGALID>
                        <ns3:LEGALHOLDERNAME>KETEM HAILU TAYE</ns3:LEGALHOLDERNAME>
                        <ns3:LEGALISSAUTH>FAYDA</ns3:LEGALISSAUTH>
                        <ns3:LEGALISSDATE>20210504</ns3:LEGALISSDATE>
                        <ns3:LEGALEXPDATE>20500101</ns3:LEGALEXPDATE>
                    </ns3:mLEGALID>
                </ns3:gLEGALID>
                <ns3:LANGUAGE>1</ns3:LANGUAGE>
                <ns3:COMPANYBOOK>ET0010001</ns3:COMPANYBOOK>
                <ns3:CLSCPARTY>NO</ns3:CLSCPARTY>
                <ns3:gCRPROFILETYPE>
                    <ns3:mCRPROFILETYPE>
                        <ns3:CRPROFILETYPE>VALUED.CUSTOMER</ns3:CRPROFILETYPE>
                        <ns3:CRPROFILE>14</ns3:CRPROFILE>
                    </ns3:mCRPROFILETYPE>
                </ns3:gCRPROFILETYPE>
                <ns3:GIVENNAMES>Ketem</ns3:GIVENNAMES>
                <ns3:FAMILYNAME>Hailu</ns3:FAMILYNAME>
                <ns3:GENDER>MALE</ns3:GENDER>
                <ns3:DATEOFBIRTH>19900310</ns3:DATEOFBIRTH>
                <ns3:MARITALSTATUS>SINGLE</ns3:MARITALSTATUS>
                <ns3:NOOFDEPENDENTS>0</ns3:NOOFDEPENDENTS>
                <ns3:gPHONE1>
                    <ns3:mPHONE1>
                        <ns3:SMS1>+251912895689</ns3:SMS1>
                        <ns3:EMAIL1>sampletet@gmail.com</ns3:EMAIL1>
                    </ns3:mPHONE1>
                </ns3:gPHONE1>
                <ns3:gEMPLOYMENTSTATUS>
                    <ns3:mEMPLOYMENTSTATUS>
                        <ns3:EMPLOYMENTSTATUS>EMPLOYED</ns3:EMPLOYMENTSTATUS>
                        <ns3:OCCUPATION>HIRED</ns3:OCCUPATION>
                        <ns3:EMPLOYERSNAME>MIDRO</ns3:EMPLOYERSNAME>
                        <ns3:sgEMPLOYERSADD>
                            <ns3:EMPLOYERSADD>
                                <ns3:EMPLOYERSADD>ADDIS ABABA</ns3:EMPLOYERSADD>
                            </ns3:EMPLOYERSADD>
                        </ns3:sgEMPLOYERSADD>
                        <ns3:EMPLOYERSBUSS>SHARE COMPANY</ns3:EMPLOYERSBUSS>
                        <ns3:CUSTOMERCURRENCY>ETB</ns3:CUSTOMERCURRENCY>
                        <ns3:SALARY>75000.00</ns3:SALARY>
                        <ns3:ANNUALBONUS>50000.00</ns3:ANNUALBONUS>
                    </ns3:mEMPLOYMENTSTATUS>
                </ns3:gEMPLOYMENTSTATUS>
                <ns3:NETMONTHLYIN>55000.00</ns3:NETMONTHLYIN>
                <ns3:NETMONTHLYOUT>42000.00</ns3:NETMONTHLYOUT>
                <ns3:AMLCHECK>NULL</ns3:AMLCHECK>
                <ns3:AMLRESULT>NULL</ns3:AMLRESULT>
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
                    <ns3:OVERRIDE>KEBELEID/CUS*100 FROM 1195875233 NOT RECEIVED</ns3:OVERRIDE>
                </ns3:gOVERRIDE>
                <ns3:RECORDSTATUS>INAO</ns3:RECORDSTATUS>
                <ns3:CURRNO>1</ns3:CURRNO>
                <ns3:gINPUTTER>
                    <ns3:INPUTTER>40531_SUPERAPP.1__OFS_GCS</ns3:INPUTTER>
                </ns3:gINPUTTER>
                <ns3:gDATETIME>
                    <ns3:DATETIME>2512241604</ns3:DATETIME>
                </ns3:gDATETIME>
                <ns3:AUTHORISER>40531_SUPERAPP.1</ns3:AUTHORISER>
                <ns3:COCODE>ET0010001</ns3:COCODE>
                <ns3:DEPTCODE>1</ns3:DEPTCODE>
                <ns3:Ownership>3000</ns3:Ownership>
                <ns3:ETMBTINNO>7775854855128587</ns3:ETMBTINNO>
                <ns3:COMMPRE>SMS</ns3:COMMPRE>
                <ns3:CUSTMOTHER>TIGIST ADAM FEKADU</ns3:CUSTMOTHER>
                <ns3:DATACLEAND>YES</ns3:DATACLEAND>
                <ns3:CUSTGRUOP>RETAIL</ns3:CUSTGRUOP>
                <ns3:NATIONALID>4455567887777455</ns3:NATIONALID>
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
	assert.Equal(t, "K912895609", detail.Menmonic)
	assert.Equal(t, "KETEM HAILU TAYE", detail.FullName)
	assert.Equal(t, "ADDIS ABABA", detail.Address)
	assert.Equal(t, "4144", detail.PostalCode)
	assert.Equal(t, "ET", detail.Country)
	assert.Equal(t, "7124", detail.AccountOfficer)
	assert.Equal(t, "1499", detail.Industry)
	assert.Equal(t, "ET", detail.Nationality)
	assert.Equal(t, "20210504", detail.IssuedDate)
	assert.Equal(t, "20500101", detail.ExpiryDate)
	assert.Equal(t, "ET0010001", detail.CompanyBook)
	assert.Equal(t, "MALE", detail.Gender)
	assert.Equal(t, "19900310", detail.DateOfBirth)
	assert.Equal(t, "SINGLE", detail.MaritalStatus)
	assert.Equal(t, "+251912895689", detail.PhoneNumber)
	assert.Equal(t, "sampletet@gmail.com", detail.Email)
	assert.Equal(t, "EMPLOYED", detail.EmploymentStatus)
	assert.Equal(t, "75000.00", detail.Salary)
	assert.Equal(t, "MIDRO", detail.Customer)
	assert.Equal(t, "50000.00", detail.AnnualBonus)
	assert.Equal(t, "ETB", detail.Currency)
	assert.Equal(t, "7775854855128587", detail.TinNumber)
	assert.Equal(t, "TIGIST ADAM FEKADU", detail.MotherName)
	assert.Equal(t, "RETAIL", detail.CustomerGroup)
	assert.Equal(t, "4455567887777455", detail.NationalId)
	assert.Equal(t, "3000", detail.Ownership)
	assert.Equal(t, "ET0010001", detail.Cocode)
}

func TestNewCustomerCreation_SampleRequest(t *testing.T) {
	params := Params{
		Username:           "SUPERAPP",
		Password:           "123456",
		FirstName:          "KETEM",
		MiddleName:         "HAILU",
		LastName:           "TAYE",
		PhoneNumber:        "+251912895689",
		Address:            "ADDIS ABABA",
		PostalCode:         "4144",
		ISOCountryCode:     "ET",
		AccountOffice:      "7124",
		Industry:           "1499",
		ISONationalityCode: "ET",
		ISOResidentCode:    "ET",
		UniqueID:           "123456779",
		IssuesBy:           "FAYDA",
		IssuedDate:         "20210504",
		ExpiryDate:         "20500101",
		Gender:             "MALE",
		DateOfBirth:        "19900310",
		MaritalStatus:      "SINGLE",
		Email:              "sampletet@gmail.com",
		EmploymentStatus:   "EMPLOYED",
		Occupation:         "HIRED",
		EmployerName:       "MIDRO",
		EmployerAddress:    "ADDIS ABABA",
		EmployerBusiness:   "SHARE COMPANY",
		CustomerCurrency:   "ETB",
		Salary:             "75000",
		AnnualBonus:        "50000",
		NetMonthlyIncome:   "55000",
		NetMonthlyExpence:  "42000",
		TinNumber:          "7775854855128587",
		MotherName:         "TIGIST ADAM FEKADU",
		CustomerGroup:      "RETAIL",
		NationalId:         "4455567887777455",
	}

	xmlRequest := NewCustomerCreation(params)

	assert.Contains(t, xmlRequest, "<cus:MNEMONIC>K912895689</cus:MNEMONIC>")
	assert.Contains(t, xmlRequest, "<cus:SHORTNAME>KETEM HAILU TAYE</cus:SHORTNAME>")
	assert.Contains(t, xmlRequest, "<cus:LEGALID>123456779</cus:LEGALID>")
	assert.Contains(t, xmlRequest, "<cus:NationalId>4455567887777455</cus:NationalId>")
	assert.Contains(t, xmlRequest, "<cus:SMS1>+251912895689</cus:SMS1>")
	assert.Contains(t, xmlRequest, "<cus:EMAIL1>sampletet@gmail.com</cus:EMAIL1>")
	assert.Contains(t, xmlRequest, "<cus:ACCOUNTOFFICER>7124</cus:ACCOUNTOFFICER>")
	assert.Contains(t, xmlRequest, "<cus:CustomerGroup>RETAIL</cus:CustomerGroup>")
}
