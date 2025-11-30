package fundtransfer

import (
	"crypto/tls"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIPSAccountLookup(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(12345678)
	o := r.Intn(12345678)
	endtoend := fmt.Sprintf("ENDTOEND%s", strconv.Itoa(r.Intn(12345678)))
	ft := fmt.Sprintf("FT%s", strconv.Itoa(r.Intn(12345678)))
	params := Params{
		CreditAccountNumber:       "1234567890",
		DebitBankBIC:              "CBETETAA",
		CreditBankBIC:             "ETSETAA",
		BizMessageIdentifier:      fmt.Sprintf("CBETETAA%s", strconv.Itoa(n)),
		MessageIdentifier:         fmt.Sprintf("CBETETAA%s", strconv.Itoa(o)),
		CreditDateTime:            "2023-06-24T00:00:00.000+03:00",
		CreditDate:                "2023-06-24T00:00:00.000Z",
		EndToEndIdentifier:        endtoend,
		TransactionIdentifier:     ft,
		InterBankSettlementAmount: "10",
		AccptanceDtatTime:         "2023-06-24T00:00:00.000+03:00",
		InstructedAmount:          "10",
		DebitAccountNumber:        "1234567890",
		CreditAccountHolderName:   "test",
		Narative:                  "Fule",
	}

	xmlRequest := NewFundTransfer(params)
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/cbe-dev/sandbox/mb_ips_soap?target=https://api-gw-uat-gateway-apic-nonprod.apps.cp4itest.cbe.local"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("MB_authorization", "Basic TUJVU0VSLUlQUzoxMjM0NTY=")
	req.Header.Set("username", "cbe")
	req.Header.Set("password", "cbe1")
	req.Header.Set("grant_type", "password")
	req.Header.Set("Jwt_Assertion", "eyJhbGciOiJSUzI1NiJ9.eyJpc3MiOiJDQkVURVRBQSIsImNlcnRfaXNzIjoiQ049VEVTVCBFVFMgSVBTIElzc3VpbmcgQ0EsIE89RXRoU3dpdGNoLCBDPUVUIiwiY2VydF9zbiI6IjQyMzcxNDE1OTEwNjI1MzI5NjM5NDAzNTQxMTM0NDcwNjU1Njk4MDYyNTQ3MiIsImp0aSI6IjExMjIzMzEyNDEyMzIxIiwiZXhwIjo0NjgzNDc2NjU3MDR9.HhTOwliC86XOhpXhNUwD0t_-S7tcSvAoJrs5fLnzQ7jjJHu3GrjZKyqjhzjg5E5DydsOiht8BONlYeuSjou9QD7ZMayzq1DATdo26TVsSzLrp4Ao_8c12xbCYV8yvGjI1xXOGTNF08ylxcznGj-Jiyp9QmywTQFIGPceJYEsi83TJePbO2dWiHIyQexT45dNivp1DAvxk8CD7W63q_R4bRgKW-F8thy9ER5NC-V5l_xWSxvPl0Iu_JyD1ig59Mpc5UjQ92fpe1D0vXBsRrDMmqCVWL5Axj9ZTKY9HZziu0kNQxgpxKB1ZXFs_Btoqni6LWE4sO_i9JV9uyPOFmy7vw")
	req.Header.Set("Authorization", "Bearer AAIgNTljMjFmZThhMDdhN2NiNmYzNjM2ZjZmMzExMjQ2NTM41xsB09BzTR1ZIx-FkZEoIrapx8mJqVBXgEv6KiSZWgg1fQXk3iItNcfymTl_SKBuKZVvSoWzKizoncd3g6T6UHsM9YJkuR9D9iAsmQab2j2bgJ0Wc6ZSVs4yc9cbIJPc9KyI7tzCirmpLPfr88n4")

	client := &http.Client{
		Timeout: 30 * time.Second,
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

	result, err := ParseFundTransferSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeeded
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)
	t.Log(result.Detail)

	if result.Detail != nil {
		assert.Equal(t, endtoend, result.Detail.OriginalEndtoEndIdentifier)
		assert.Equal(t, ft, result.Detail.OriginalTransactionIdentifier)
		t.Log("passed")
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
