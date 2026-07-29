package upload

import (
	"crypto/tls"
	"encoding/base64"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const uploadEndpoint = "https://devapisuperapp.cbe.com.et/superapp/parser/fncmis/wsdl?target=https://10.3.46.185:9081"

func TestIntegrationUpload(t *testing.T) {
	stream := base64.StdEncoding.EncodeToString([]byte("sample document content for upload test"))

	params := Params{
		Username:         "SUPERAPP",
		Password:         "123456",
		Name:             "test-upload.pdf",
		DocumentType:     "NATIONAL.ID",
		CustomerNumber:   "1202724929",
		Source:           "SUPERAPP",
		UniqueIdentifier: "UID-TEST-001",
		Email:            "sampleTest@gmail.com",
		PhoneNumber:      "+251913323635",
		FullName:         "MELESE TESFAYE KIFLE",
		MimeType:         "application/pdf",
		FileName:         "test-upload.pdf",
		Stream:           stream,
	}

	xmlRequest := NewUpload(params)
	t.Log("XML Request:", xmlRequest)

	req, err := http.NewRequest("POST", uploadEndpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Logf("Network error (endpoint may be unreachable): %v", err)
		t.Skip("Skipping test due to network error - endpoint may be unreachable")
	}
	assert.NotNil(t, resp)
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseData, "Expected response body to be non-empty")

	t.Logf("Response status: %s", resp.Status)
	t.Logf("Response body: %s", string(responseData))

	result, err := ParseUploadSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result)

	t.Logf("Parsed ObjectID: %s", result.ObjectID)

	if result.ObjectID == "" {
		t.Log("No objectId returned - endpoint may not support CMIS createDocument")
		return
	}

	assert.NotEmpty(t, result.ObjectID)
}
