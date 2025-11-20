package accountlookup

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationFundTransfer(t *testing.T) {
	params := Params{
		CreditAccountNumber:  "5122867049011",
		DebitBankBIC:         "CBETETAA",
		CreditBankBIC:        "DASHETAA",
		BizMessageIdentifier: "BIZMSGID20251120000123",
		MessageIdentifier:    "MSGID-FT-20251120-98763",
		CreditDateTime:       "2025-11-20T10:45:32",
		CreditDate:           "2025-11-20",
	}

	xmlRequest := NewAccountLookup(params)
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/cbe-dev/sandbox/mb_ips_soap?target=https://api-gw-uat-gateway-apic-nonprod.apps.cp4itest.cbe.local"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)

	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("MB_authorization", "Basic TUJVU0VSLUlQUzoxMjM0NTY=")
	req.Header.Add("username", "cbe")
	req.Header.Add("password", "cbe1")
	req.Header.Add("grant_type", "password")
	req.Header.Add("Jwt_Assertion", "eyJhbGciOiJSUzI1NiJ9.eyJpc3MiOiJDQkVURVRBQSIsImNlcnRfaXNzIjoiQ049VEVTVCBFVFMgSVBTIElzc3VpbmcgQ0EsIE89RXRoU3dpdGNoLCBDPUVUIiwiY2VydF9zbiI6IjQyMzcxNDE1OTEwNjI1MzI5NjM5NDAzNTQxMTM0NDcwNjU1Njk4MDYyNTQ3MiIsImp0aSI6IjExMjIzMzEyNDEyMzIxIiwiZXhwIjo0NjgzNDc2NjU3MDR9.HhTOwliC86XOhpXhNUwD0t_-S7tcSvAoJrs5fLnzQ7jjJHu3GrjZKyqjhzjg5E5DydsOiht8BONlYeuSjou9QD7ZMayzq1DATdo26TVsSzLrp4Ao_8c12xbCYV8yvGjI1xXOGTNF08ylxcznGj-Jiyp9QmywTQFIGPceJYEsi83TJePbO2dWiHIyQexT45dNivp1DAvxk8CD7W63q_R4bRgKW-F8thy9ER5NC-V5l_xWSxvPl0Iu_JyD1ig59Mpc5UjQ92fpe1D0vXBsRrDMmqCVWL5Axj9ZTKY9HZziu0kNQxgpxKB1ZXFs_Btoqni6LWE4sO_i9JV9uyPOFmy7vw")
	req.Header.Add("Authorization", "Bearer AAIgNTljMjFmZThhMDdhN2NiNmYzNjM2ZjZmMzExMjQ2NTPd9v993dzvIpyRyqwcULGmDqh-WvQ-lWN5g2Ro2NSpp981SFIgct2s5OIUmgTRFrS_HnMVynDhG1gc8Rb9UKwDFifVELuvsOpzKJxL2yxM3DYw-d1mXg-i5qCFXpwQ9jGFocFC75DY0Ot3IhW0-0Qs")
	req.Header.Add("Cookie", "681823e6349145585941f22ad5359e75=813e88cb63d01adf8f48b487a0ab6355; f27f73ac186db59cbd4936452d5d0df3=45d9258ae90c56656649465881d0bfbb")
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

	result, err := ParseAccountLookupSOAP(string(responseData))
	t.Log(result.Messages)
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeeded
	assert.True(t, result.Success)

	if result.Detail != nil {
		assert.Equal(t, "5122867049011", result.Detail.CreditAccountNumber)
		// assert.Equal(t, "ABIY HAILEYESUS MENGISTU", result.Detail.CreditAccountHolderName)
		assert.Equal(t, "DASHETAA", result.Detail.CreditBankBIC)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
