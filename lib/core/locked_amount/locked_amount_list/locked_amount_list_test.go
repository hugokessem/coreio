package lockedamountlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLockedAmountGeneratedXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate List Locked Amount",
			param: Params{
				Username:      "SUPERAPP",
				Password:      "123456",
				AccountNumber: "10009876543",
			},
			expect: []string{},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewListLockedAmount(tc.param)
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}
		})
	}
}
