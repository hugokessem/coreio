package exchangerate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewExchangeRate(t *testing.T) {
	tests := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate exchange rate XML generation",
			param: Params{
				Username: "SUPERAPP",
				Password: "123456",
			},
			expect: []string{
				`<password>123456</password>`,
				`<userName>SUPERAPP</userName>`,
				`<soapenv:Envelope`,
				`<soapenv:Body>`,
				`<cbes:ExchangeRateSuperApp>`,
				`<WebRequestCommon>`,
				`<EXCHANGERATESUPERAPPType>`,
				`<enquiryInputCollection>`,
			},
		},
		{
			name: "Validate exchange rate with different values",
			param: Params{
				Username: "TESTUSER",
				Password: "PASSWORD123",
			},
			expect: []string{
				`<password>PASSWORD123</password>`,
				`<userName>TESTUSER</userName>`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewExchangeRate(tc.param)

			// Validate XML structure
			assert.Contains(t, xmlRequest, "<soapenv:Envelope")
			assert.Contains(t, xmlRequest, "<soapenv:Body>")
			assert.Contains(t, xmlRequest, "<cbes:ExchangeRateSuperApp>")

			// Validate all expected strings are present
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}

			// Validate XML is not empty
			assert.NotEmpty(t, xmlRequest)
		})
	}
}

func TestParseExchangeRateSOAP(t *testing.T) {
	tests := []struct {
		name            string
		xmlData         string
		expectedSuccess bool
		expectedError   bool
		expectedDetail  bool
		expectedCount   int
		expectedMessage string
	}{
		{
			name: "Parse successful response with single exchange rate",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:ExchangeRateSuperAppResponse xmlns:ns18="http://temenos.com/EXCHANGERATESUPERAPP" xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <EXCHANGERATESUPERAPPType>
                <ns18:gEXCHANGERATESUPERAPPDetailType>
                    <ns18:mEXCHANGERATESUPERAPPDetailType>
                        <ns18:ID>USD</ns18:ID>
                        <ns18:NUMERICCCYCODE>840</ns18:NUMERICCCYCODE>
                        <ns18:CCYNAME>US Dollar</ns18:CCYNAME>
                        <ns18:CURRENCYMARKET>Transaction</ns18:CURRENCYMARKET>
                        <ns18:BUYRATE>55.5000</ns18:BUYRATE>
                        <ns18:SELLRATE>56.2000</ns18:SELLRATE>
                        <ns18:MIDRATE>55.8500</ns18:MIDRATE>
                    </ns18:mEXCHANGERATESUPERAPPDetailType>
                </ns18:gEXCHANGERATESUPERAPPDetailType>
            </EXCHANGERATESUPERAPPType>
        </ns32:ExchangeRateSuperAppResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: true,
			expectedError:   false,
			expectedDetail:  true,
			expectedCount:   1,
		},
		{
			name: "Parse successful response with multiple exchange rates",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:ExchangeRateSuperAppResponse xmlns:ns18="http://temenos.com/EXCHANGERATESUPERAPP" xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <EXCHANGERATESUPERAPPType>
                <ns18:gEXCHANGERATESUPERAPPDetailType>
                    <ns18:mEXCHANGERATESUPERAPPDetailType>
                        <ns18:ID>USD</ns18:ID>
                        <ns18:NUMERICCCYCODE>840</ns18:NUMERICCCYCODE>
                        <ns18:CCYNAME>US Dollar</ns18:CCYNAME>
                        <ns18:CURRENCYMARKET>Transaction</ns18:CURRENCYMARKET>
                        <ns18:BUYRATE>55.5000</ns18:BUYRATE>
                        <ns18:SELLRATE>56.2000</ns18:SELLRATE>
                        <ns18:MIDRATE>55.8500</ns18:MIDRATE>
                    </ns18:mEXCHANGERATESUPERAPPDetailType>
                    <ns18:mEXCHANGERATESUPERAPPDetailType>
                        <ns18:ID>EUR</ns18:ID>
                        <ns18:NUMERICCCYCODE>978</ns18:NUMERICCCYCODE>
                        <ns18:CCYNAME>Euro</ns18:CCYNAME>
                        <ns18:CURRENCYMARKET>Transaction</ns18:CURRENCYMARKET>
                        <ns18:BUYRATE>60.0000</ns18:BUYRATE>
                        <ns18:SELLRATE>61.0000</ns18:SELLRATE>
                        <ns18:MIDRATE>60.5000</ns18:MIDRATE>
                    </ns18:mEXCHANGERATESUPERAPPDetailType>
                </ns18:gEXCHANGERATESUPERAPPDetailType>
            </EXCHANGERATESUPERAPPType>
        </ns32:ExchangeRateSuperAppResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: true,
			expectedError:   false,
			expectedDetail:  true,
			expectedCount:   2,
		},
		{
			name: "Parse response with failure status",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:ExchangeRateSuperAppResponse xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Failure</successIndicator>
            </Status>
        </ns32:ExchangeRateSuperAppResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetail:  false,
			expectedMessage: "API returned failure",
		},
		{
			name: "Parse invalid XML",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<InvalidXML>
    <Broken>
</InvalidXML>`,
			expectedSuccess: false,
			expectedError:   true,
			expectedDetail:  false,
		},
		{
			name: "Parse response without Status",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:ExchangeRateSuperAppResponse xmlns:ns32="http://temenos.com/CBESUPERAPP">
        </ns32:ExchangeRateSuperAppResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetail:  false,
			expectedMessage: "Missing Status",
		},
		{
			name: "Parse response with no exchange rate type",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:ExchangeRateSuperAppResponse xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
        </ns32:ExchangeRateSuperAppResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetail:  false,
		},
		{
			name: "Parse response with empty exchange rate details",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:ExchangeRateSuperAppResponse xmlns:ns18="http://temenos.com/EXCHANGERATESUPERAPP" xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <EXCHANGERATESUPERAPPType>
                <ns18:gEXCHANGERATESUPERAPPDetailType>
                </ns18:gEXCHANGERATESUPERAPPDetailType>
            </EXCHANGERATESUPERAPPType>
        </ns32:ExchangeRateSuperAppResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: true,
			expectedError:   false,
			expectedDetail:  true,
			expectedCount:   0,
		},
		{
			name: "Parse response with case-insensitive success indicator",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:ExchangeRateSuperAppResponse xmlns:ns18="http://temenos.com/EXCHANGERATESUPERAPP" xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>SUCCESS</successIndicator>
            </Status>
            <EXCHANGERATESUPERAPPType>
                <ns18:gEXCHANGERATESUPERAPPDetailType>
                    <ns18:mEXCHANGERATESUPERAPPDetailType>
                        <ns18:ID>USD</ns18:ID>
                        <ns18:NUMERICCCYCODE>840</ns18:NUMERICCCYCODE>
                        <ns18:CCYNAME>US Dollar</ns18:CCYNAME>
                        <ns18:CURRENCYMARKET>Transaction</ns18:CURRENCYMARKET>
                        <ns18:BUYRATE>55.5000</ns18:BUYRATE>
                        <ns18:SELLRATE>56.2000</ns18:SELLRATE>
                        <ns18:MIDRATE>55.8500</ns18:MIDRATE>
                    </ns18:mEXCHANGERATESUPERAPPDetailType>
                </ns18:gEXCHANGERATESUPERAPPDetailType>
            </EXCHANGERATESUPERAPPType>
        </ns32:ExchangeRateSuperAppResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: true,
			expectedError:   false,
			expectedDetail:  true,
			expectedCount:   1,
		},
		{
			name: "Parse response with invalid response structure",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <OtherResponse xmlns:ns32="http://temenos.com/CBESUPERAPP">
        </OtherResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetail:  false,
			expectedMessage: "Invalid response type",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseExchangeRateSOAP(tc.xmlData)

			if tc.expectedError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tc.expectedSuccess, result.Success)

				if tc.expectedDetail {
					assert.NotNil(t, result.Detail)
					if result.Detail != nil {
						assert.Equal(t, tc.expectedCount, len(result.Detail), "Expected %d exchange rates, got %d", tc.expectedCount, len(result.Detail))
					}
				} else {
					if result != nil {
						if tc.expectedMessage != "" {
							assert.NotEmpty(t, result.Message)
							if len(result.Message) > 0 {
								assert.Contains(t, result.Message[0], tc.expectedMessage)
							}
						}
					}
				}
			}
		})
	}
}

func TestParseExchangeRateSOAP_DetailFields(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:ExchangeRateSuperAppResponse xmlns:ns18="http://temenos.com/EXCHANGERATESUPERAPP" xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <EXCHANGERATESUPERAPPType>
                <ns18:gEXCHANGERATESUPERAPPDetailType>
                    <ns18:mEXCHANGERATESUPERAPPDetailType>
                        <ns18:ID>USD</ns18:ID>
                        <ns18:NUMERICCCYCODE>840</ns18:NUMERICCCYCODE>
                        <ns18:CCYNAME>US Dollar</ns18:CCYNAME>
                        <ns18:CURRENCYMARKET>Transaction</ns18:CURRENCYMARKET>
                        <ns18:BUYRATE>55.5000</ns18:BUYRATE>
                        <ns18:SELLRATE>56.2000</ns18:SELLRATE>
                        <ns18:MIDRATE>55.8500</ns18:MIDRATE>
                    </ns18:mEXCHANGERATESUPERAPPDetailType>
                    <ns18:mEXCHANGERATESUPERAPPDetailType>
                        <ns18:ID>EUR</ns18:ID>
                        <ns18:NUMERICCCYCODE>978</ns18:NUMERICCCYCODE>
                        <ns18:CCYNAME>Euro</ns18:CCYNAME>
                        <ns18:CURRENCYMARKET>Cash</ns18:CURRENCYMARKET>
                        <ns18:BUYRATE>60.0000</ns18:BUYRATE>
                        <ns18:SELLRATE>61.0000</ns18:SELLRATE>
                        <ns18:MIDRATE>60.5000</ns18:MIDRATE>
                    </ns18:mEXCHANGERATESUPERAPPDetailType>
                </ns18:gEXCHANGERATESUPERAPPDetailType>
            </EXCHANGERATESUPERAPPType>
        </ns32:ExchangeRateSuperAppResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseExchangeRateSOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		assert.Equal(t, 2, len(result.Detail), "Expected exactly 2 exchange rates")

		// Check first exchange rate
		if len(result.Detail) > 0 {
			firstRate := result.Detail[0]
			assert.Equal(t, "USD", firstRate.ID)
			assert.Equal(t, "840", firstRate.NumericCCYCode)
			assert.Equal(t, "US Dollar", firstRate.CCYName)
			assert.Equal(t, "Transaction", firstRate.CurrencyMarket)
			assert.Equal(t, "55.5000", firstRate.BuyRate)
			assert.Equal(t, "56.2000", firstRate.SellRate)
			assert.Equal(t, "55.8500", firstRate.MidRate)
		}

		// Check second exchange rate
		if len(result.Detail) > 1 {
			secondRate := result.Detail[1]
			assert.Equal(t, "EUR", secondRate.ID)
			assert.Equal(t, "978", secondRate.NumericCCYCode)
			assert.Equal(t, "Euro", secondRate.CCYName)
			assert.Equal(t, "Cash", secondRate.CurrencyMarket)
			assert.Equal(t, "60.0000", secondRate.BuyRate)
			assert.Equal(t, "61.0000", secondRate.SellRate)
			assert.Equal(t, "60.5000", secondRate.MidRate)
		}
	}
}

func TestParseExchangeRateSOAP_MultipleRates(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:ExchangeRateSuperAppResponse xmlns:ns18="http://temenos.com/EXCHANGERATESUPERAPP" xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <EXCHANGERATESUPERAPPType>
                <ns18:gEXCHANGERATESUPERAPPDetailType>
                    <ns18:mEXCHANGERATESUPERAPPDetailType>
                        <ns18:ID>AED</ns18:ID>
                        <ns18:NUMERICCCYCODE>784</ns18:NUMERICCCYCODE>
                        <ns18:CCYNAME>United Arab Emirates Dirhams</ns18:CCYNAME>
                        <ns18:CURRENCYMARKET>Transaction</ns18:CURRENCYMARKET>
                        <ns18:BUYRATE>37.7845</ns18:BUYRATE>
                        <ns18:SELLRATE>38.5402</ns18:SELLRATE>
                        <ns18:MIDRATE>38.1624</ns18:MIDRATE>
                    </ns18:mEXCHANGERATESUPERAPPDetailType>
                    <ns18:mEXCHANGERATESUPERAPPDetailType>
                        <ns18:ID>AUD</ns18:ID>
                        <ns18:NUMERICCCYCODE>36</ns18:NUMERICCCYCODE>
                        <ns18:CCYNAME>Australian Dollars</ns18:CCYNAME>
                        <ns18:CURRENCYMARKET>Transaction</ns18:CURRENCYMARKET>
                        <ns18:BUYRATE>34.1743</ns18:BUYRATE>
                        <ns18:SELLRATE>34.8578</ns18:SELLRATE>
                        <ns18:MIDRATE>34.5161</ns18:MIDRATE>
                    </ns18:mEXCHANGERATESUPERAPPDetailType>
                    <ns18:mEXCHANGERATESUPERAPPDetailType>
                        <ns18:ID>CAD</ns18:ID>
                        <ns18:NUMERICCCYCODE>124</ns18:NUMERICCCYCODE>
                        <ns18:CCYNAME>Canadian Dollar</ns18:CCYNAME>
                        <ns18:CURRENCYMARKET>Transaction</ns18:CURRENCYMARKET>
                        <ns18:BUYRATE>40.5000</ns18:BUYRATE>
                        <ns18:SELLRATE>41.2000</ns18:SELLRATE>
                        <ns18:MIDRATE>40.8500</ns18:MIDRATE>
                    </ns18:mEXCHANGERATESUPERAPPDetailType>
                </ns18:gEXCHANGERATESUPERAPPDetailType>
            </EXCHANGERATESUPERAPPType>
        </ns32:ExchangeRateSuperAppResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseExchangeRateSOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		assert.Equal(t, 3, len(result.Detail), "Expected exactly 3 exchange rates")

		// Verify we can find the expected currencies
		currencyIDs := make(map[string]bool)
		for _, rate := range result.Detail {
			currencyIDs[rate.ID] = true
		}

		assert.True(t, currencyIDs["AED"], "Should contain AED")
		assert.True(t, currencyIDs["AUD"], "Should contain AUD")
		assert.True(t, currencyIDs["CAD"], "Should contain CAD")

		// Verify rate fields are populated
		for _, rate := range result.Detail {
			if rate.ID == "AED" && rate.CurrencyMarket == "Transaction" {
				assert.Equal(t, "37.7845", rate.BuyRate)
				assert.Equal(t, "38.5402", rate.SellRate)
				assert.Equal(t, "38.1624", rate.MidRate)
			}
		}
	}
}

