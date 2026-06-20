package acccountcreation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const sampleAccountCreationResponse = `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns6:AccountOpeningSuperAppResponse xmlns:ns2="http://temenos.com/CUSTOMERCREATEINDIVIDUAL" xmlns:ns3="http://temenos.com/CUSTOMER" xmlns:ns4="http://temenos.com/ACCOUNTCREATEINDIVIDUAL" xmlns:ns5="http://temenos.com/ACCOUNT" xmlns:ns6="http://temenos.com/IIBONBOARDING">
            <Status>
                <transactionId>1000446112639</transactionId>
                <messageId></messageId>
                <successIndicator>Success</successIndicator>
                <application>ACCOUNT</application>
            </Status>
            <ACCOUNTType id="1000446112639">
                <ns5:CUSTOMER>1027958756</ns5:CUSTOMER>
                <ns5:CATEGORY>6501</ns5:CATEGORY>
                <ns5:gACCOUNTTITLE1>
                    <ns5:ACCOUNTTITLE1>ALEMNESH ZINABI DIKO</ns5:ACCOUNTTITLE1>
                </ns5:gACCOUNTTITLE1>
                <ns5:gSHORTTITLE>
                    <ns5:SHORTTITLE>ALEMNESH ZINABI DIKO</ns5:SHORTTITLE>
                </ns5:gSHORTTITLE>
                <ns5:POSITIONTYPE>TR</ns5:POSITIONTYPE>
                <ns5:CURRENCY>ETB</ns5:CURRENCY>
                <ns5:CURRENCYMARKET>1</ns5:CURRENCYMARKET>
                <ns5:ACCOUNTOFFICER>7016</ns5:ACCOUNTOFFICER>
                <ns5:gPOSTINGRESTRICT>
                    <ns5:POSTINGRESTRICT>14</ns5:POSTINGRESTRICT>
                </ns5:gPOSTINGRESTRICT>
                <ns5:CONDITIONGROUP>14</ns5:CONDITIONGROUP>
                <ns5:gCAPDATECHARGE>
                    <ns5:CAPDATECHARGE>20211231</ns5:CAPDATECHARGE>
                </ns5:gCAPDATECHARGE>
                <ns5:PASSBOOK>NO</ns5:PASSBOOK>
                <ns5:OPENINGDATE>20211209</ns5:OPENINGDATE>
                <ns5:OPENCATEGORY>6501</ns5:OPENCATEGORY>
                <ns5:CHARGECCY>ETB</ns5:CHARGECCY>
                <ns5:CHARGEMKT>1</ns5:CHARGEMKT>
                <ns5:INTERESTCCY>ETB</ns5:INTERESTCCY>
                <ns5:INTERESTMKT>1</ns5:INTERESTMKT>
                <ns5:gALTACCTTYPE>
                    <ns5:mALTACCTTYPE>
                        <ns5:ALTACCTTYPE>BANKMASTER</ns5:ALTACCTTYPE>
                    </ns5:mALTACCTTYPE>
                    <ns5:mALTACCTTYPE>
                        <ns5:ALTACCTTYPE>SMARTBANK</ns5:ALTACCTTYPE>
                    </ns5:mALTACCTTYPE>
                </ns5:gALTACCTTYPE>
                <ns5:ALLOWNETTING>NO</ns5:ALLOWNETTING>
                <ns5:SINGLELIMIT>Y</ns5:SINGLELIMIT>
                <ns5:CURRNO>1</ns5:CURRNO>
                <ns5:gINPUTTER>
                    <ns5:INPUTTER>2230_SUPERAPP.1__OFS_GCS</ns5:INPUTTER>
                </ns5:gINPUTTER>
                <ns5:gDATETIME>
                    <ns5:DATETIME>2512241644</ns5:DATETIME>
                </ns5:gDATETIME>
                <ns5:AUTHORISER>2230_SUPERAPP.1_OFS_GCS</ns5:AUTHORISER>
                <ns5:COCODE>ET0010001</ns5:COCODE>
                <ns5:DEPTCODE>1</ns5:DEPTCODE>
                <ns5:HASJOINTCUST>NO</ns5:HASJOINTCUST>
                <ns5:PRODUCTTYPE>PLASSTIC</ns5:PRODUCTTYPE>
            </ACCOUNTType>
        </ns6:AccountOpeningSuperAppResponse>
    </S:Body>
</S:Envelope>`

func TestParseAccountCreationSOAP_Success(t *testing.T) {
	result, err := ParseAccountCreationSOAP(sampleAccountCreationResponse)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	detail := result.Detail
	assert.Equal(t, "1000446112639", detail.AccountNumber)
	assert.Equal(t, "1027958756", detail.Customer)
	assert.Equal(t, "6501", detail.Category)
	assert.Equal(t, "ALEMNESH ZINABI DIKO", detail.AccountTitle)
	assert.Equal(t, "ALEMNESH ZINABI DIKO", detail.ShortTitle)
	assert.Equal(t, "TR", detail.PositionType)
	assert.Equal(t, "ETB", detail.Currency)
	assert.Equal(t, "1", detail.CurrencyMarket)
	assert.Equal(t, "7016", detail.AccountOfficer)
	assert.Equal(t, "14", detail.PostingRestrict)
	assert.Equal(t, "14", detail.ConditionGroup)
	assert.Equal(t, "20211231", detail.CapDateCharge)
	assert.Equal(t, "NO", detail.Passbook)
	assert.Equal(t, "20211209", detail.OpeningDate)
	assert.Equal(t, "6501", detail.OpenCategory)
	assert.Equal(t, "ETB", detail.ChargeCcy)
	assert.Equal(t, "1", detail.ChargeMkt)
	assert.Equal(t, "ETB", detail.InterestCcy)
	assert.Equal(t, "1", detail.InterestMkt)
	assert.Equal(t, []string{"BANKMASTER", "SMARTBANK"}, detail.AltAcctTypes)
	assert.Equal(t, "NO", detail.AllowNetting)
	assert.Equal(t, "Y", detail.SingleLimit)
	assert.Equal(t, "1", detail.CurrNo)
	assert.Equal(t, "2230_SUPERAPP.1__OFS_GCS", detail.Inputter)
	assert.Equal(t, "2512241644", detail.Datetime)
	assert.Equal(t, "2230_SUPERAPP.1_OFS_GCS", detail.Authoriser)
	assert.Equal(t, "ET0010001", detail.CoCode)
	assert.Equal(t, "1", detail.DeptCode)
	assert.Equal(t, "NO", detail.HasJointCust)
	assert.Equal(t, "PLASSTIC", detail.ProductType)
}

func TestNewAccountCreation_SampleRequest(t *testing.T) {
	params := Params{
		Username:       "SUPERAPP",
		Password:       "123456",
		CustomerNumber: "1195875233",
		Category:       "6501",
		Currency:       "ETB",
		AccountOfficer: "7016",
	}

	xmlRequest := NewAccountCreation(params)

	assert.Contains(t, xmlRequest, `xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"`)
	assert.Contains(t, xmlRequest, `xmlns:iib="http://temenos.com/IIBONBOARDING"`)
	assert.Contains(t, xmlRequest, `xmlns:acc="http://temenos.com/ACCOUNTCREATEINDIVIDUAL"`)
	assert.Contains(t, xmlRequest, "<iib:AccountOpeningSuperApp>")
	assert.Contains(t, xmlRequest, "<company/>")
	assert.Contains(t, xmlRequest, "<userName>SUPERAPP</userName>")
	assert.Contains(t, xmlRequest, "<password>123456</password>")
	assert.Contains(t, xmlRequest, "<acc:CUSTOMER>1195875233</acc:CUSTOMER>")
	assert.Contains(t, xmlRequest, "<acc:CATEGORY>6501</acc:CATEGORY>")
	assert.Contains(t, xmlRequest, "<acc:CURRENCY>ETB</acc:CURRENCY>")
	assert.Contains(t, xmlRequest, "<acc:ACCOUNTOFFICER>7016</acc:ACCOUNTOFFICER>")
}
