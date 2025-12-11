package exchangerate

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationExchangeRate(t *testing.T) {
	params := Params{
		Username: "SUPERAPP",
		Password: "123456",
	}

	xmlRequest := NewExchangeRate(params)
	endpoint := "https://devopscbe.eaglelionsystems.com/superapp/parser/proxy/CBESUPERAPP/services?target=http%3A%2F%2F10.1.15.195%3A8080&wsdl=null"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseData, "Expected response body to be non-empty")

	result, err := ParseExchangeRateSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeeded
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)
	t.Logf("result: %v", result.Detail)

	if result.Detail != nil {
		assert.Greater(t, len(result.Detail), 0, "Expected at least one exchange rate")
		t.Logf("Exchange Rate result: Found %d exchange rates", len(result.Detail))

		// Log first few exchange rates
		maxLog := 5
		if len(result.Detail) < maxLog {
			maxLog = len(result.Detail)
		}

		for i := 0; i < maxLog; i++ {
			rate := result.Detail[i]
			t.Logf("  [%d] ID=%s, CCYName=%s, Market=%s, BuyRate=%s, SellRate=%s, MidRate=%s",
				i+1, rate.ID, rate.CCYName, rate.CurrencyMarket,
				rate.BuyRate, rate.SellRate, rate.MidRate)
		}

		// Validate that exchange rates have required fields
		for i, rate := range result.Detail {
			if rate.ID != "" {
				assert.NotEmpty(t, rate.ID, "Exchange rate %d should have ID", i+1)
				t.Logf("Exchange rate %d: %s - Buy: %s, Sell: %s, Mid: %s",
					i+1, rate.ID, rate.BuyRate, rate.SellRate, rate.MidRate)
			}
		}

		t.Log("Integration test passed")
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}

