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
		Username:            "SUPERAPP",
		Password:            "123456",
		DebitAccountNumber:  "1000199986857",
		DebitCurrency:       "ETB",
		DebitAmount:         "6",
		CreditCurrency:      "ETB",
		CreditAccountNumber: "1000498200943",
		CreditReference:     "Credit reference",
		DebitReference:      "Debit reference",
		ClientReference:     uuid.String(),
		ServiceCode:         "SITOTA",
		LockID:              "ACLK21343NQQP0",
		PaymentDetails:      "Apple",
		CustomerRole:        "MASS",
		ChannelType:         "APP",
		SuperappUserCode:    "SA1771239173",
	}

	xmlRequest := NewLockedAmountFt(params)
	t.Logf("Request XML: %s", xmlRequest)
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
	t.Logf("---resp: %v", resp)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	t.Logf("---resp.Body: %+v", resp.Body)
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
