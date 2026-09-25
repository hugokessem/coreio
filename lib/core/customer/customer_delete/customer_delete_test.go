package customerdelete

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCustomerDelete(t *testing.T) {
	xmlRequest := NewCustomerDelete(Param{
		Username:       "SUPERAPP",
		Password:       "123456",
		CustomerNumber: "1419395586",
	})

	assert.Contains(t, xmlRequest, "<iib:DeleteCustomerIDSuperApp>")
	assert.Contains(t, xmlRequest, "<password>123456</password>")
	assert.Contains(t, xmlRequest, "<userName>SUPERAPP</userName>")
	assert.Contains(t, xmlRequest, "<transactionId>1419395586</transactionId>")
}

func TestParseCustomerDeleteSOAP_Success(t *testing.T) {
	xmlData := `<?xml version="1.0" encoding="UTF-8"?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
	<S:Body>
		<ns6:DeleteCustomerIDSuperAppResponse
			xmlns:ns6="http://temenos.com/IIBONBOARDING"
			xmlns:ns5="http://temenos.com/ACCOUNT"
			xmlns:ns4="http://temenos.com/ACCOUNTCREATEINDIVIDUAL"
			xmlns:ns3="http://temenos.com/CUSTOMER"
			xmlns:ns2="http://temenos.com/CUSTOMERCREATEINDIVIDUAL">
			<Status>
				<transactionId>1419395586</transactionId>
				<messageId></messageId>
				<successIndicator>Success</successIndicator>
				<application>CUSTOMER</application>
			</Status>
			<CUSTOMERType id="1419395586">
				<ns3:MNEMONIC>M2567890</ns3:MNEMONIC>
				<ns3:gSHORTNAME>
					<ns3:SHORTNAME>MELESE TESFAYE KIFLE</ns3:SHORTNAME>
				</ns3:gSHORTNAME>
				<ns3:gNAME1>
					<ns3:NAME1>MELESE TESFAYE KIFLE</ns3:NAME1>
				</ns3:gNAME1>
				<ns3:gNAME2>
					<ns3:NAME2>MELESE TESFAYE KIFLE</ns3:NAME2>
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
				<ns3:GIVENNAMES>MELESE</ns3:GIVENNAMES>
				<ns3:FAMILYNAME>TESFAYE</ns3:FAMILYNAME>
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
				<ns3:CUSTOMERTYPE>PROSPECT</ns3:CUSTOMERTYPE>
				<ns3:gFURTHERDETAILS>
					<ns3:FURTHERDETAILS>SUPERAPP</ns3:FURTHERDETAILS>
				</ns3:gFURTHERDETAILS>
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
					<ns3:OVERRIDE>KEBELEID/CUS*100 FROM 1419395586 NOT RECEIVED</ns3:OVERRIDE>
				</ns3:gOVERRIDE>
				<ns3:RECORDSTATUS>IDEL</ns3:RECORDSTATUS>
				<ns3:CURRNO>1</ns3:CURRNO>
				<ns3:gINPUTTER>
					<ns3:INPUTTER>72467_SUPERAPP.A__OFS_GCS</ns3:INPUTTER>
				</ns3:gINPUTTER>
				<ns3:gDATETIME>
					<ns3:DATETIME>2609251500</ns3:DATETIME>
				</ns3:gDATETIME>
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
				<ns3:CUSTMOTHER>TIGIST ALEMU</ns3:CUSTMOTHER>
				<ns3:DATACLEAND>YES</ns3:DATACLEAND>
				<ns3:FATCACOMPLIANT>NO</ns3:FATCACOMPLIANT>
				<ns3:PEPSTATUS>NO</ns3:PEPSTATUS>
				<ns3:USPERSON>NO</ns3:USPERSON>
				<ns3:HOUSENO>1544/02</ns3:HOUSENO>
				<ns3:CUTSEGEMENT>MASS</ns3:CUTSEGEMENT>
				<ns3:MCUSTSEGEMENT>MASS</ns3:MCUSTSEGEMENT>
				<ns3:CUSTGRUOP>RETAIL</ns3:CUSTGRUOP>
				<ns3:COMPVSIND>INDIVIDUAL</ns3:COMPVSIND>
				<ns3:SALESPERSON>888888</ns3:SALESPERSON>
				<ns3:FAYDAVERIFIED>YES</ns3:FAYDAVERIFIED>
				<ns3:USTINNO>10002520520</ns3:USTINNO>
			</CUSTOMERType>
		</ns6:DeleteCustomerIDSuperAppResponse>
	</S:Body>
</S:Envelope>`

	result, err := ParseCustomerDeleteSOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.True(t, result.Success)
	require.NotNil(t, result.Detail)

	d := result.Detail
	assert.Equal(t, "1419395586", d.CustomerNumber)
	assert.Equal(t, "1419395586", d.TransactionID)
	assert.Equal(t, "CUSTOMER", d.Application)
	assert.Equal(t, "M2567890", d.Mnemonic)
	assert.Equal(t, "MELESE TESFAYE KIFLE", d.ShortName)
	assert.Equal(t, "MELESE TESFAYE KIFLE", d.Name1)
	assert.Equal(t, "MELESE TESFAYE KIFLE", d.Name2)
	assert.Equal(t, "AM", d.Street)
	assert.Equal(t, "ADDIS ABABA", d.Address)
	assert.Equal(t, "ADDIS ABABA", d.TownCountry)
	assert.Equal(t, "4125", d.PostalCode)
	assert.Equal(t, "ET", d.Country)
	assert.Equal(t, "1000", d.Sector)
	assert.Equal(t, "7858", d.AccountOfficer)
	assert.Equal(t, "1499", d.Industry)
	assert.Equal(t, "4", d.Target)
	assert.Equal(t, "ET", d.Nationality)
	assert.Equal(t, "1", d.CustomerStatus)
	assert.Equal(t, "ET", d.Residence)
	assert.Equal(t, "WS55651461632DSS", d.LegalID)
	assert.Equal(t, "NATIONAL.ID", d.LegalDocName)
	assert.Equal(t, "MELESE TESFAYE KIFLE", d.LegalHolderName)
	assert.Equal(t, "FAYDA", d.LegalIssAuth)
	assert.Equal(t, "20211209", d.LegalIssDate)
	assert.Equal(t, "20270808", d.LegalExpDate)
	assert.Equal(t, "1", d.Language)
	assert.Equal(t, "ET0011859", d.CompanyBook)
	assert.Equal(t, "NO", d.CLSCParty)
	assert.Equal(t, "VALUED.CUSTOMER", d.CRProfileType)
	assert.Equal(t, "14", d.CRProfile)
	assert.Equal(t, "MELESE", d.GivenNames)
	assert.Equal(t, "TESFAYE", d.FamilyName)
	assert.Equal(t, "MALE", d.Gender)
	assert.Equal(t, "19910215", d.DateOfBirth)
	assert.Equal(t, "MARRIED", d.MaritalStatus)
	assert.Equal(t, "1", d.NoOfDependents)
	assert.Equal(t, "+251913323635", d.PhoneNumber)
	assert.Equal(t, "sampleTest@gmail.com", d.Email)
	assert.Equal(t, "EMPLOYED", d.EmploymentStatus)
	assert.Equal(t, "ACCOUNTANT", d.Occupation)
	assert.Equal(t, "ETB", d.CustomerCurrency)
	assert.Equal(t, "28000.00", d.Salary)
	assert.Equal(t, "PROSPECT", d.CustomerType)
	assert.Equal(t, "SUPERAPP", d.FurtherDetails)
	assert.Equal(t, "NULL", d.AMLCheck)
	assert.Equal(t, "NULL", d.AMLResult)
	assert.Equal(t, "YES", d.KYCComplete)
	assert.Equal(t, "NULL", d.InternetBankingService)
	assert.Equal(t, "NULL", d.MobileBankingService)
	assert.Equal(t, "VALUED.CUSTOMER", d.CRUserProfileType)
	assert.Equal(t, "14", d.CRCalcProfile)
	assert.Equal(t, "14", d.CRUserProfile)
	assert.Equal(t, "NO", d.Reserved01)
	assert.Equal(t, []string{"KEBELEID/CUS*100 FROM 1419395586 NOT RECEIVED"}, d.Override)
	assert.Equal(t, "IDEL", d.RecordStatus)
	assert.Equal(t, "1", d.CurrNo)
	assert.Equal(t, "72467_SUPERAPP.A__OFS_GCS", d.Inputter)
	assert.Equal(t, "2609251500", d.DateTime)
	assert.Equal(t, "ET0010001", d.CoCode)
	assert.Equal(t, "1", d.DeptCode)
	assert.Equal(t, "3000", d.Ownership)
	assert.Equal(t, []string{"AC.ALERTS", "FT.ALERTS", "TT.ALERTS"}, d.CorBanGroup)
	assert.Equal(t, "Banker", d.CustomerOccupation)
	assert.Equal(t, "First Degree", d.EducationStatus)
	assert.Equal(t, "SMS", d.CommunicationPref)
	assert.Equal(t, "TIGIST ALEMU", d.MotherName)
	assert.Equal(t, "YES", d.DataCleaned)
	assert.Equal(t, "NO", d.FATCACompliant)
	assert.Equal(t, "NO", d.PEPStatus)
	assert.Equal(t, "NO", d.USPerson)
	assert.Equal(t, "1544/02", d.HouseNo)
	assert.Equal(t, "MASS", d.CustomerSubSegment)
	assert.Equal(t, "MASS", d.CustomerSegment)
	assert.Equal(t, "RETAIL", d.CustomerGroup)
	assert.Equal(t, "INDIVIDUAL", d.CompVsInd)
	assert.Equal(t, "888888", d.SalesPerson)
	assert.Equal(t, "YES", d.FaydaVerified)
	assert.Equal(t, "10002520520", d.USTinNo)
}

func TestParseCustomerDeleteSOAP_RecordNotFound(t *testing.T) {
	xmlData := `<?xml version="1.0" encoding="UTF-8"?><S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/"><S:Body><ns6:DeleteCustomerIDSuperAppResponse xmlns:ns6="http://temenos.com/IIBONBOARDING" xmlns:ns5="http://temenos.com/ACCOUNT" xmlns:ns4="http://temenos.com/ACCOUNTCREATEINDIVIDUAL" xmlns:ns3="http://temenos.com/CUSTOMER" xmlns:ns2="http://temenos.com/CUSTOMERCREATEINDIVIDUAL"><Status><transactionId>1419395586</transactionId><messageId></messageId><successIndicator>T24Error</successIndicator><application>CUSTOMER</application><messages>@ID:1:1=UNAUTH. RECORD MISSING</messages></Status></ns6:DeleteCustomerIDSuperAppResponse></S:Body></S:Envelope>`

	result, err := ParseCustomerDeleteSOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Nil(t, result.Detail)
	assert.Equal(t, []string{"@ID:1:1=UNAUTH. RECORD MISSING"}, result.Messages)
}

func TestParseCustomerDeleteSOAP_InvalidResponseType(t *testing.T) {
	xmlData := `<?xml version="1.0" encoding="UTF-8"?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
  <S:Body>
    <SomeOtherResponse/>
  </S:Body>
</S:Envelope>`

	result, err := ParseCustomerDeleteSOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, []string{"Invalid response type"}, result.Messages)
}

func TestParseCustomerDeleteSOAP_MissingStatus(t *testing.T) {
	xmlData := `<?xml version="1.0" encoding="UTF-8"?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
  <S:Body>
    <ns6:DeleteCustomerIDSuperAppResponse xmlns:ns6="http://temenos.com/IIBONBOARDING">
      <CUSTOMERType id="1419395586"/>
    </ns6:DeleteCustomerIDSuperAppResponse>
  </S:Body>
</S:Envelope>`

	result, err := ParseCustomerDeleteSOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, []string{"Missing Status"}, result.Messages)
}

func TestParseCustomerDeleteSOAP_SOAPFault(t *testing.T) {
	xmlData := `<?xml version="1.0" encoding="UTF-8"?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
  <S:Body>
    <S:Fault>
      <faultcode>S:Client</faultcode>
      <faultstring>Invalid request</faultstring>
    </S:Fault>
  </S:Body>
</S:Envelope>`

	result, err := ParseCustomerDeleteSOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, []string{"Invalid request"}, result.Messages)
}
