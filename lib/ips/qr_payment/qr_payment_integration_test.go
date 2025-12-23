package qrpayment

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

func TestIPSQrPayment(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(12345678)
	o := r.Intn(12345678)
	endtoend := fmt.Sprintf("CBETETAA%s", strconv.Itoa(r.Intn(12345678)))
	txid := fmt.Sprintf("CBETETAAFT%s", strconv.Itoa(r.Intn(12345678)))
	
	// Get current time for timestamps
	now := time.Now()
	creditDate := now.UTC().Format("2006-01-02T15:04:05.000Z")
	creditDateTime := now.Format("2006-01-02T15:04:05.000-07:00")
	acceptanceDateTime := now.Format("2006-01-02T15:04:05.000-07:00")

	params := Params{
		DebitBankBIC:              "CBETETAA",
		CreditBankBIC:             "ETSETAA",
		BizMessageIdentifier:      fmt.Sprintf("CBETETAA%s", strconv.Itoa(n)),
		CreditDate:                creditDate,
		MessageIdentifier:         fmt.Sprintf("CBETETAA%s", strconv.Itoa(o)),
		CreditDateTime:             creditDateTime,
		EndToEndIdentifier:         endtoend,
		TransactionIdentifier:     txid,
		InterBankSettlementAmount: "10.00",
		AccptanceDtatTime:         acceptanceDateTime,
		InstructedAmount:          "10.00",
		DebitAccountNumber:        "1234567890",
		CreditAccountNumber:       "1234567890",
		CreditAccountHolderName:   "Merchant Name",
		Narative:                  "QR Payment Test",
		DebiterInformation: DebiterInformation{
			Name:           "Test Debiter",
			StreetName:      "Main Street",
			BuildingNumber: "123",
			PostalCode:      "1000",
			TownName:        "Addis Ababa",
			Country:         "ET",
		},
		CreditInformation: CrediterInformation{
			Name:           "Merchant Name",
			StreetName:     "Tito St.",
			BuildingNumber: "17",
			PostalCode:     "18444",
			City:           "Addis Ababa",
			TownName:       "Addis Ababa",
			Country:        "ET",
			AddressLine:    "Kazanchis",
		},
	}

	xmlRequest := NewQrPayment(params)
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/cbe-dev/sandbox/mb_ips_soap?target=https://api-gw-uat-gateway-apic-nonprod.apps.cp4itest.cbe.local"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("MB_authorization", "Basic TUJVU0VSLUlQUzoxMjM0NTY=")
	req.Header.Set("username", "cbe")
	req.Header.Set("password", "cbe1")
	req.Header.Set("grant_type", "password")
	req.Header.Set("Jwt_Assertion", "eyJhbGciOiJSUzI1NiJ9.eyJpc3MiOiJDQkVURVRBQSIsImNlcnRfaXNzIjoiQ049VEVTVCBFVFMgSVBTIElzc3VpbmcgQ0EsIE89RXRoU3dpdGNoLCBDPUVUIiwiY2VydF9zbiI6IjQyMzcxNDE1OTEwNjI1MzI5NjM5NDAzNTQxMTM0NDcwNjU1Njk4MDYyNTQ3MiIsImp0aSI6IjExMjIzMzEyNDEyMzIxIiwiZXhwIjo0NjgzNDc2NjU3MDR9.HhTOwliC86XOhpXhNUwD0t_-S7tcSvAoJrs5fLnzQ7jjJHu3GrjZKyqjhzjg5E5DydsOiht8BONlYeuSjou9QD7ZMayzq1DATdo26TVsSzLrp4Ao_8c12xbCYV8yvGjI1xXOGTNF08ylxcznGj-Jiyp9QmywTQFIGPceJYEsi83TJePbO2dWiHIyQexT45dNivp1DAvxk8CD7W63q_R4bRgKW-F8thy9ER5NC-V5l_xWSxvPl0Iu_JyD1ig59Mpc5UjQ92fpe1D0vXBsRrDMmqCVWL5Axj9ZTKY9HZziu0kNQxgpxKB1ZXFs_Btoqni6LWE4sO_i9JV9uyPOFmy7vw")
	req.Header.Set("Authorization", "Bearer AAIgNTljMjFmZThhMDdhN2NiNmYzNjM2ZjZmMzExMjQ2NTMFmUEuhlkH_ozPL50d6smezJQaiVJrpWXs9rSIWi9cA8MAy6I4PiLA_N6jHXdGVxf12ft2CmgC_pJkZLUu06aBTsOer2AZRAyZbsbFcVn1PUda7XFD0eOsIrWIuAGd_pdyxsmK0fCMcgLAwKkS_9sY")

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

	result, err := ParseQrPaymentSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Log the result for debugging
	t.Logf("QR Payment Result: Success=%v, TransactionStatus=%s", result.Success, func() string {
		if result.Detail != nil {
			return result.Detail.TransactionStatus
		}
		return "N/A"
	}())

	// Check that we got a valid response
	assert.NotNil(t, result.Detail, "Expected Detail to be non-nil")

	if result.Detail != nil {
		// Verify transaction identifiers match
		assert.Equal(t, endtoend, result.Detail.OriginalEndToEndIdentifier, "EndToEndIdentifier should match")
		assert.Equal(t, txid, result.Detail.OriginalTransactionIdentifier, "TransactionIdentifier should match")

		// Log transaction details
		t.Logf("Transaction Details:")
		t.Logf("  Status: %s", result.Detail.TransactionStatus)
		t.Logf("  From Bank: %s", result.Detail.FromBankBIC)
		t.Logf("  To Bank: %s", result.Detail.ToBankBIC)
		t.Logf("  Amount: %s %s", result.Detail.InterBankSettlementAmount, result.Detail.InterBankSettlementCurrency)
		t.Logf("  Creditor: %s", result.Detail.CreditorName)
		t.Logf("  Creditor Account: %s", result.Detail.CreditorAccountNumber)

		// If transaction was successful, verify additional fields
		if result.Success {
			assert.Equal(t, "ACSC", result.Detail.TransactionStatus, "Transaction status should be ACSC for success")
			assert.NotEmpty(t, result.Detail.CreditorAccountNumber, "Creditor account number should not be empty")
			assert.NotEmpty(t, result.Detail.CreditorName, "Creditor name should not be empty")
			assert.NotEmpty(t, result.Detail.InterBankSettlementAmount, "InterBank settlement amount should not be empty")
		} else {
			// If transaction was rejected, verify status reason is present
			t.Logf("Transaction Rejected:")
			if result.Detail.StatusReason != "" {
				t.Logf("  Reason: %s", result.Detail.StatusReason)
			}
			if result.Detail.StatusReasonAdditionalInfo != "" {
				t.Logf("  Additional Info: %s", result.Detail.StatusReasonAdditionalInfo)
			}
			if len(result.Messages) > 0 {
				t.Logf("  Messages: %v", result.Messages)
			}
		}
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}



