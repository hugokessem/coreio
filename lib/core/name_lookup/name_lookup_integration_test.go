package namelookup

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var accountNumbers = []string{
	"1000485754923",
	"1000139436605",
	"1000419162624",
	"1000187614825",
	"1000445347276",
	"1000425404947",
	"1000387584092",
	"1000349971608",
	"1000420695587",
	"1000269518182",
	"1000396430472",
	"1000305718819",
	"1000298095649",
	"1000517052152",
}

func TestIntegrationNameLookupForMultipleAccounts(t *testing.T) {
	var listOfResults []map[string]string
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPP/services?target=http://10.1.15.195%3A8080&wsdl=null"
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		},
	}

	for _, accountNumber := range accountNumbers {
		accountNumber := accountNumber
		t.Run(accountNumber, func(t *testing.T) {
			params := Params{
				Username:      "SUPERAPP",
				Password:      "123456",
				AccountNumber: accountNumber,
			}
			xmlRequest := NewNameLookup(params)
			t.Logf("xmlRequest %v", xmlRequest)

			req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))
			require.NoError(t, err)
			req.Header.Add("Content-Type", "text/xml; charset=utf-8")

			resp, err := client.Do(req)
			if err != nil {
				t.Skipf("network error (endpoint may be unreachable): %v", err)
			}
			require.NotNil(t, resp)
			defer resp.Body.Close()

			responseData, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			require.NotEmpty(t, responseData)

			result, err := ParseNameLookupSOAP(string(responseData))
			require.NoError(t, err)
			require.NotNil(t, result)

			entry := map[string]string{
				"accountNumber":  accountNumber,
				"customerName":   result.Detail.CustomerNumber,
				"customerNumber": result.Detail.CustomerNumber,
			}
			if result.Detail != nil {
				entry["customerName"] = result.Detail.AccountName
				entry["currency"] = result.Detail.Currency
			}
			if len(result.Messages) > 0 {
				entry["messages"] = strings.Join(result.Messages, ", ")
			}

			listOfResults = append(listOfResults, entry)
			t.Logf("result: success=%v detail=%+v messages=%v", result.Success, result.Detail, result.Messages)

			if !result.Success {
				t.Logf("lookup failed for account %s: %v", accountNumber, result.Messages)
				return
			}

			require.NotNil(t, result.Detail)
			assert.Equal(t, accountNumber, result.Detail.AccountNumber)
			assert.NotEmpty(t, result.Detail.AccountName)
			time.Sleep(3 * time.Second)
		})
	}

	t.Logf("collected %d results: %+v", len(listOfResults), listOfResults)
}

func TestIntegrationAccountLookup(t *testing.T) {
	params := Params{
		Username:      "SUPERAPP",
		Password:      "123456",
		AccountNumber: "1000197649848",
	}

	xmlRequest := NewNameLookup(params)
	t.Logf("xmlRequest %v", xmlRequest)
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPP/services?target=http://10.1.15.195%3A8080&wsdl=null"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		},
	}

	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseData, "Expected response body to be non-empty")

	result, err := ParseNameLookupSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeeded
	// if result == nil {
	// 	t.Fatal("Expected result to be non-nil")
	// }

	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)
	t.Logf("Details: %+v", result.Detail)
	// t.Logf("Currency: %+v", result.Detail.Currency)
	// t.Logf("Customer Number: %+v", result.Detail.CustomerNumber)
	// t.Log("resultDetail", result)

	// if result.Detail != nil {
	// 	assert.Equal(t, "1000517052152", result.Detail.AccountNumber)
	// 	assert.Equal(t, "ETB", result.Detail.Currency)
	// } else {
	// 	t.Error("Expected Detail to be non-nil")
	// }
}
