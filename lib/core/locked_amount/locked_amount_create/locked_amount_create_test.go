package lockedamountcreate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLockedAmountGeneratedXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate Create Locked Amount",
			param: Params{
				Username:      "SUPERAPP",
				Password:      "123456",
				AccountNumber: "1000000006924",
				Description:   "3 Click",
				From:          "20251108",
				To:            "20251111",
				LockedAmount:  "1200",
			},
			expect: []string{
				`<acl:ACCOUNTNUMBER>1000000006924</acl:ACCOUNTNUMBER>`,
				`<acl:DESCRIPTION>3 Click</acl:DESCRIPTION>`,
				`<acl:FROMDATE>20251108</acl:FROMDATE>`,
				`<acl:TODATE>20251111</acl:TODATE>`,
				`<acl:LOCKEDAMOUNT>1200</acl:LOCKEDAMOUNT>`,
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewCreateLockedAmount(tc.param)
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}
		})
	}
}
