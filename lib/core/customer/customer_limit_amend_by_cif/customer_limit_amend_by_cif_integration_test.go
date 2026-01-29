package customerlimitamendbycif

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationCustomerLimitAmendByCIF(t *testing.T) {
	params := Params{
		Username:       "SUPERAPP",
		Password:       "123456",
		CustomerNumber: "111111",
		ChannelLimit: []ChannelLimit{
			{
				Channel: "APP",
				ServiceLimits: []ServiceLimit{
					{
						ServiceType:           "AAPARKING",
						ServiceMaximumAmount:  "23000",
						UserMaximumDebitCount: "5",
					},
					{
						ServiceType:           "TELEBIRR",
						ServiceMaximumAmount:  "25000",
						UserMaximumDebitCount: "5",
					},
				},
			},
			{
				Channel: "USSD",
				ServiceLimits: []ServiceLimit{
					{
						ServiceType:           "AAPARKING",
						ServiceMaximumAmount:  "50100",
						UserMaximumDebitCount: "10",
					},
				},
			},
		},
	}

	xmlRequest := NewCustomerLimitAmendByCIF(params)
	t.Log("XML Request:", xmlRequest)
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPP/services?target=http://10.1.15.195%3A8080&wsdl=null"

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
	assert.NoError(t, err)
	assert.NotEmpty(t, responseData, "Expected response body to be non-empty")

	result, err := ParseCustomerLimitAmendByCIFSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the amendment succeeded
	assert.True(t, result.Success)
	t.Log("resultMessage", result.Messages)
	t.Log("resultDetail", result.Detail.ID)

	for _, channelLimit := range result.Detail.GUserChannel.MUserChannel {
		t.Log("Channel:", channelLimit.UserChannelType)
	}
}
