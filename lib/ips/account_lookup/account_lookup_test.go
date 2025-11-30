package accountlookup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountLookupGeneratedXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate Account Lookups",
			param: Params{
				CreditAccountNumber:  "1234567890",
				DebitBankBIC:         "CBETETAA",
				CreditBankBIC:        "ETSETAA",
				BizMessageIdentifier: "CBETETAA553572981",
				MessageIdentifier:    "CBETETAA847772981",
				CreditDateTime:       "2023-06-24T00:00:00.000+03:00",
				CreditDate:           "2023-06-24T00:00:00.000Z",
			},
			expect: []string{
				`<urn:BizMsgIdr>CBETETAA553572981</urn:BizMsgIdr>`,
				`<urn:CreDt>2023-06-24T00:00:00.000Z</urn:CreDt>`,
				`<urn1:MsgId>CBETETAA847772981</urn1:MsgId>`,
				`<urn1:Id>1234567890</urn1:Id>`,
			},
		},
		{
			name: "Validate Account Lookups with different values",
			param: Params{
				CreditAccountNumber:  "9876543210",
				DebitBankBIC:         "CBETETAA",
				CreditBankBIC:        "AWINETAA",
				BizMessageIdentifier: "CBETETAA999999999",
				MessageIdentifier:    "CBETETAA888888888",
				CreditDateTime:       "2024-01-15T10:30:00.000+03:00",
				CreditDate:           "2024-01-15T00:00:00.000Z",
			},
			expect: []string{
				`<urn:BizMsgIdr>CBETETAA999999999</urn:BizMsgIdr>`,
				`<urn:CreDt>2024-01-15T00:00:00.000Z</urn:CreDt>`,
				`<urn1:MsgId>CBETETAA888888888</urn1:MsgId>`,
				`<urn1:Id>9876543210</urn1:Id>`,
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewAccountLookup(tc.param)
			
			// Validate XML structure
			assert.Contains(t, xmlRequest, "<soapenv:Envelope")
			assert.Contains(t, xmlRequest, "<soapenv:Body>")
			assert.Contains(t, xmlRequest, "<mb:AccountVerfication>")
			
			// Validate all expected strings are present
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}
			
			// Validate XML is not empty
			assert.NotEmpty(t, xmlRequest)
		})
	}
}

func TestParseAccountLookupSOAP(t *testing.T) {
	tests := []struct {
		name           string
		xmlData        string
		expectedResult *AccountLookupResult
		expectError    bool
	}{
		{
			name: "Successful account lookup response",
			xmlData: `<?xml version="1.0" encoding="UTF-8"?>
<Envelope>
	<Body>
		<AccountVerficationResponse>
			<output1>
				<AppHdr>
					<Fr>
						<FIId>
							<FinInstnId>
								<Othr>
									<Id>CBETETAA</Id>
								</Othr>
							</FinInstnId>
						</FIId>
					</Fr>
					<To>
						<FIId>
							<FinInstnId>
								<Othr>
									<Id>ETSETAA</Id>
								</Othr>
							</FinInstnId>
						</FIId>
					</To>
					<BizMsgIdr>CBETETAA553572981</BizMsgIdr>
					<CreDt>2023-06-24T00:00:00.000Z</CreDt>
				</AppHdr>
				<Document>
					<IdVrfctnRpt>
						<Assgnmt>
							<MsgId>CBETETAA847772981</MsgId>
							<CreDtTm>2023-06-24T00:00:00.000+03:00</CreDtTm>
						</Assgnmt>
						<Rpt>
							<OrgnlId>CBETETAA847772981</OrgnlId>
							<Vrfctn>true</Vrfctn>
							<Rsn>
								<Prtry>ACCOUNT_FOUND</Prtry>
							</Rsn>
							<OrgnlPtyAndAcctId>
								<Acct>
									<Id>
										<Othr>
											<Id>1234567890</Id>
											<SchmeNm>
												<Prtry>ACCT</Prtry>
											</SchmeNm>
										</Othr>
									</Id>
								</Acct>
							</OrgnlPtyAndAcctId>
							<UpdtdPtyAndAcctId>
								<Pty>
									<Nm>John Doe</Nm>
								</Pty>
								<Acct>
									<Id>
										<Othr>
											<Id>1234567890</Id>
											<SchmeNm>
												<Prtry>ACCT</Prtry>
											</SchmeNm>
										</Othr>
									</Id>
								</Acct>
							</UpdtdPtyAndAcctId>
						</Rpt>
					</IdVrfctnRpt>
				</Document>
			</output1>
		</AccountVerficationResponse>
	</Body>
</Envelope>`,
			expectedResult: &AccountLookupResult{
				Success: true,
				Detail: &AccountVerficationDetail{
					OriginalIdentifier:      "CBETETAA847772981",
					CreditAccountNumber:     "1234567890",
					CreditAccountHolderName: "John Doe",
				},
			},
			expectError: false,
		},
		{
			name: "Account not found response",
			xmlData: `<?xml version="1.0" encoding="UTF-8"?>
<Envelope>
	<Body>
		<AccountVerficationResponse>
			<output1>
				<AppHdr>
					<Fr>
						<FIId>
							<FinInstnId>
								<Othr>
									<Id>CBETETAA</Id>
								</Othr>
							</FinInstnId>
						</FIId>
					</Fr>
					<To>
						<FIId>
							<FinInstnId>
								<Othr>
									<Id>ETSETAA</Id>
								</Othr>
							</FinInstnId>
						</FIId>
					</To>
				</AppHdr>
				<Document>
					<IdVrfctnRpt>
						<Rpt>
							<OrgnlId>CBETETAA847772981</OrgnlId>
							<Vrfctn>false</Vrfctn>
							<Rsn>
								<Prtry>ACCOUNT_NOT_FOUND</Prtry>
							</Rsn>
						</Rpt>
					</IdVrfctnRpt>
				</Document>
			</output1>
		</AccountVerficationResponse>
	</Body>
</Envelope>`,
			expectedResult: &AccountLookupResult{
				Success:  false,
				Detail:   nil,
				Messages: []string{"Account Not Found!"},
			},
			expectError: false,
		},
		{
			name: "Missing report data",
			xmlData: `<?xml version="1.0" encoding="UTF-8"?>
<Envelope>
	<Body>
		<AccountVerficationResponse>
			<output1>
				<AppHdr>
					<Fr>
						<FIId>
							<FinInstnId>
								<Othr>
									<Id>CBETETAA</Id>
								</Othr>
							</FinInstnId>
						</FIId>
					</Fr>
				</AppHdr>
				<Document>
					<IdVrfctnRpt>
					</IdVrfctnRpt>
				</Document>
			</output1>
		</AccountVerficationResponse>
	</Body>
</Envelope>`,
			expectedResult: &AccountLookupResult{
				Success:  false,
				Detail:   nil,
				Messages: []string{"Invalid Response: Missing report data"},
			},
			expectError: false,
		},
		{
			name: "Invalid response structure",
			xmlData: `<?xml version="1.0" encoding="UTF-8"?>
<Envelope>
	<Body>
		<InvalidResponse>
		</InvalidResponse>
	</Body>
</Envelope>`,
			expectedResult: &AccountLookupResult{
				Success:  false,
				Detail:   nil,
				Messages: []string{"Invalid Response!"},
			},
			expectError: false,
		},
		{
			name:        "Invalid XML",
			xmlData:     `<invalid xml>`,
			expectError: true,
		},
		{
			name: "Verification with uppercase TRUE",
			xmlData: `<?xml version="1.0" encoding="UTF-8"?>
<Envelope>
	<Body>
		<AccountVerficationResponse>
			<output1>
				<AppHdr>
					<To>
						<FIId>
							<FinInstnId>
								<Othr>
									<Id>ETSETAA</Id>
								</Othr>
							</FinInstnId>
						</FIId>
					</To>
				</AppHdr>
				<Document>
					<IdVrfctnRpt>
						<Rpt>
							<OrgnlId>CBETETAA847772981</OrgnlId>
							<Vrfctn>TRUE</Vrfctn>
							<OrgnlPtyAndAcctId>
								<Acct>
									<Id>
										<Othr>
											<Id>9876543210</Id>
										</Othr>
									</Id>
								</Acct>
							</OrgnlPtyAndAcctId>
							<UpdtdPtyAndAcctId>
								<Pty>
									<Nm>Jane Smith</Nm>
								</Pty>
							</UpdtdPtyAndAcctId>
						</Rpt>
					</IdVrfctnRpt>
				</Document>
			</output1>
		</AccountVerficationResponse>
	</Body>
</Envelope>`,
			expectedResult: &AccountLookupResult{
				Success: true,
				Detail: &AccountVerficationDetail{
					OriginalIdentifier:      "CBETETAA847772981",
					CreditAccountNumber:     "9876543210",
					CreditAccountHolderName: "Jane Smith",
				},
			},
			expectError: false,
		},
		{
			name: "Fallback to From field when To is empty",
			xmlData: `<?xml version="1.0" encoding="UTF-8"?>
<Envelope>
	<Body>
		<AccountVerficationResponse>
			<output1>
				<AppHdr>
					<Fr>
						<FIId>
							<FinInstnId>
								<Othr>
									<Id>AWINETAA</Id>
								</Othr>
							</FinInstnId>
						</FIId>
					</Fr>
					<To>
						<FIId>
							<FinInstnId>
								<Othr>
									<Id></Id>
								</Othr>
							</FinInstnId>
						</FIId>
					</To>
				</AppHdr>
				<Document>
					<IdVrfctnRpt>
						<Rpt>
							<OrgnlId>CBETETAA847772981</OrgnlId>
							<Vrfctn>true</Vrfctn>
							<OrgnlPtyAndAcctId>
								<Acct>
									<Id>
										<Othr>
											<Id>1111111111</Id>
										</Othr>
									</Id>
								</Acct>
							</OrgnlPtyAndAcctId>
							<UpdtdPtyAndAcctId>
								<Pty>
									<Nm>Test User</Nm>
								</Pty>
							</UpdtdPtyAndAcctId>
						</Rpt>
					</IdVrfctnRpt>
				</Document>
			</output1>
		</AccountVerficationResponse>
	</Body>
</Envelope>`,
			expectedResult: &AccountLookupResult{
				Success: true,
				Detail: &AccountVerficationDetail{
					OriginalIdentifier:      "CBETETAA847772981",
					CreditAccountNumber:     "1111111111",
					CreditAccountHolderName: "Test User",
				},
			},
			expectError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseAccountLookupSOAP(tc.xmlData)

			if tc.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				
				if tc.expectedResult != nil {
					assert.Equal(t, tc.expectedResult.Success, result.Success)
					assert.Equal(t, tc.expectedResult.Messages, result.Messages)
					
					if tc.expectedResult.Detail != nil {
						assert.NotNil(t, result.Detail)
						assert.Equal(t, tc.expectedResult.Detail.OriginalIdentifier, result.Detail.OriginalIdentifier)
						assert.Equal(t, tc.expectedResult.Detail.CreditAccountNumber, result.Detail.CreditAccountNumber)
						assert.Equal(t, tc.expectedResult.Detail.CreditAccountHolderName, result.Detail.CreditAccountHolderName)
					} else {
						assert.Nil(t, result.Detail)
					}
				}
			}
		})
	}
}
