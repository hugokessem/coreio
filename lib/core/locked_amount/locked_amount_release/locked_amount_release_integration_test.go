package lockedamountrelease

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReleaseLockedAmount(t *testing.T, lockedID string) {
	params := Params{
		Username:      "SUPERAPP",
		Password:      "123456",
		TransactionID: lockedID,
	}

	xmlRequest := NewReleaseLockedAmount(params)
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
	assert.NoError(t, err)
	assert.NotEmpty(t, responseData, "Expected response body to be non-empty")

	result, err := ParseCancleLockedAmountSOAP(string(responseData))

	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeeded
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		t.Logf(" TransactionID: %s, LockID: %s", result.Detail.TransactionID, lockedID)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}

func TestIntegrationReleaseLockedAmount(t *testing.T) {

	// lockedIDs := []string{
	// 	"ACLK21343QD6SP",
	// 	"ACLK21343GRD0Z",
	// 	"ACLK21343YHQB4",
	// 	"ACLK21343G4LM7",
	// 	"ACLK21343Y72GK",
	// 	"ACLK21343N7DVC",
	// 	"ACLK213431KQBM",
	// 	"ACLK21343774NF",
	// 	"ACLK21343CYG3T",
	// 	"ACLK21343KS6TP",
	// 	"ACLK21343QP6ZG",
	// 	"ACLK21343QYJ8C",
	// 	"ACLK21343TGQHN",
	// 	"ACLK21343ZTHRC",
	// 	"ACLK213430HNC6",
	// 	"ACLK21343BCFVJ",
	// 	"ACLK213431WH79",
	// 	"ACLK21343GJ3DP",
	// 	"ACLK21343J5F6H",
	// 	"ACLK21343JS1QG",
	// 	"ACLK21343MM346",
	// 	"ACLK21343PHVBG",
	// 	"ACLK21343PNXPW",
	// 	"ACLK21343XF8JD",
	// 	"ACLK2134307Y97",
	// 	"ACLK213436MPC6",
	// 	"ACLK213437WKCV",
	// 	"ACLK21343S09SM",
	// 	"ACLK21343ZQQ78",
	// 	"ACLK2134308HTT",
	// 	"ACLK213437W2JC",
	// 	"ACLK21343BQFHL",
	// 	"ACLK21343CHLFD",
	// 	"ACLK21343RJVVS",
	// 	"ACLK21343X8TXQ",
	// 	"ACLK2134375RYF",
	// 	"ACLK21343P8M07",
	// 	"ACLK21343P9QFK",
	// 	"ACLK2134303180",
	// 	"ACLK213433477Q",
	// 	"ACLK2134359CQY",
	// 	"ACLK21343BBXJC",
	// 	"ACLK21343KHZ2P",
	// 	"ACLK21343KQSCM",
	// 	"ACLK21343LQSJY",
	// 	"ACLK21343PF9M1",
	// 	"ACLK213439L85K",
	// 	"ACLK21343DMWK9",
	// 	"ACLK21343G4R42",
	// 	"ACLK21343HLYPN",
	// 	"ACLK21343LB2ZJ",
	// 	"ACLK21343MJH38",
	// 	"ACLK21343MKK8D",
	// 	"ACLK21343MN12V",
	// 	"ACLK21343NZVS6",
	// 	"ACLK21343QPNFR",
	// 	"ACLK21343R4GX9",
	// 	"ACLK21343VVN6C",
	// 	"ACLK21343BQBDP",
	// 	"ACLK21343F9C9B",
	// 	"ACLK213436KC4T",
	// 	"ACLK21343ZL40W",
	// 	"ACLK213434FCKL",
	// 	"ACLK21343WWNTY",
	// 	"ACLK21343K4TGK",
	// 	"ACLK21343LQSF9",
	// 	"ACLK213433SHFX",
	// 	"ACLK2134385F9G",
	// 	"ACLK21343DCSW5",
	// 	"ACLK213437Q5J5",
	// 	"ACLK21343JJSW9",
	// 	"ACLK21343SLR20",
	// 	"ACLK21343PTQND",
	// 	"ACLK21343XB1F5",
	// 	"ACLK21343T968K",
	// 	"ACLK21343S556L",
	// 	"ACLK21343TTZPZ",
	// 	"ACLK21343ZQGCG",
	// 	"ACLK21343WLB16",
	// 	"ACLK213430WFVQ",
	// 	"ACLK21343XF635",
	// 	"ACLK21343ZX3CV",
	// 	"ACLK21343356WM",
	// 	"ACLK21343GP290",
	// 	"ACLK21343N0W61",
	// 	"ACLK213436980C",
	// 	"ACLK213437ZL24",
	// 	"ACLK21343XTDXW",
	// 	"ACLK2134312FHM",
	// 	"ACLK21343JNMYG",
	// 	"ACLK21343KGKMW",
	// 	"ACLK21343LH48V",
	// 	"ACLK213430T24T",
	// 	"ACLK213434P4B3",
	// 	"ACLK21343PQDD8",
	// 	"ACLK213430HNWB",
	// 	"ACLK213435Y2BQ",
	// 	"ACLK21343CVWM5",
	// 	"ACLK21343L395C",
	// 	"ACLK21343YLJCS",
	// 	"ACLK2134354DXM",
	// 	"ACLK2134304K5Y",
	// 	"ACLK213432B3VD",
	// 	"ACLK213433HPXB",
	// 	"ACLK213433JN6G",
	// 	"ACLK21343FFYMV",
	// 	"ACLK21343NF0BZ",
	// 	"ACLK21343SWWZT",
	// 	"ACLK21343V2ZGM",
	// 	"ACLK21343Z05Z4",
	// }

	// for _, lockedID := range lockedIDs {
	// 	t.Logf("Testing release for LockedID: %s", lockedID)
	// 	testReleaseLockedAmount(t, lockedID)
	// }
	// return
	params := Params{
		Username:      "SUPERAPP",
		Password:      "123456",
		TransactionID: "ACLK213439PN1J",
	}

	xmlRequest := NewReleaseLockedAmount(params)
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
	assert.NoError(t, err)
	assert.NotEmpty(t, responseData, "Expected response body to be non-empty")

	result, err := ParseCancleLockedAmountSOAP(string(responseData))

	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeeded
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		assert.Equal(t, "1000000006924", result.Detail.AccountNumber)
		assert.Equal(t, "ACLK213436WRSG", result.Detail.TransactionID)
		assert.Equal(t, "4510.00", result.Detail.LockedAmount)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
