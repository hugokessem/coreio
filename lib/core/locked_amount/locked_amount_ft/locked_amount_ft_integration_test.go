package lockedamountft

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationLockedAmountFt(t *testing.T) {
	uuid, _ := uuid.NewV7()
	params := Params{
		Username: "SUPERAPP",
		Password: "123456",
		// DebitAccountNumber:  "1000000006924",
		// DebitCurrency:       "ETB",
		// DebitAmount:         "240.00",
		// DebitAccountNumber: "1000446113608",
		// DebitCurrency:      "USD",
		// CreditCurrency:      "ETB",
		// CreditAmount:        "200",
		DebitAccountNumber:  "1000259876854",
		DebitCurrency:       "ETB",
		DebitAmount:         "1000",
		CreditCurrency:      "ETB",
		CreditAccountNumber: "1000382499388",
		CreditReference:     "Credit reference",
		DebitReference:      "Debit reference",
		ClientReference:     uuid.String(),
		ServiceCode:         "IPS",
		LockID:              "ACLK2134368NJV",
		PaymentDetails:      "Apple",
		CustomerRole:        "MASS",
		ChannelType:         "APP",
		SuperappUserCode:    "8a9s08d0f9a98sdf0as80d9f80a9sd8f00",
	}

	xmlRequest := NewLockedAmountFt(params)
	t.Log("Request XML:", xmlRequest)
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPP/services?target=http%3A%2F%2F10.1.15.195%3A8080&wsdl=null"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		},
	}

	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	t.Logf("---resp: %+v", resp.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseData, "Expected response body to be non-empty")

	result, err := ParseLockedAmountFTSOAP(string(responseData))
	t.Logf("--result: %+v", result)

	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")
	t.Log("details", result.Detail)

	if result.Success {
		assert.True(t, result.Success)
		assert.NotEmpty(t, result.Detail.TransactionId)
		assert.NotEmpty(t, result.Detail.TransactionType)
		assert.NotEmpty(t, result.Detail.DebitAccountNumber)
		assert.NotEmpty(t, result.Detail.DebitCurrency)
		assert.NotEmpty(t, result.Detail.DebitAmount)
		assert.NotEmpty(t, result.Detail.CreditAccountNumber)
		assert.NotEmpty(t, result.Detail.CreditCurrenct)
		assert.NotEmpty(t, result.Detail.CreditAmountWithCurrency)
		assert.NotEmpty(t, result.Detail.DebitAmountWithCurrency)
		assert.NotEmpty(t, result.Detail.LockId)
	} else {
		t.Log("Error Messages:", result.Messages)
		assert.Fail(t, "Locked Amount FT operation failed")
	}
}
