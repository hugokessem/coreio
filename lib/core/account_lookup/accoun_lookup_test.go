package accountlookup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate account lookup",
			param: Params{
				Username:      "SUPERAPP",
				Password:      "123456",
				AccountNumber: "1000000006924",
			},
			expect: []string{
				"<criteriaValue>1000000006924</criteriaValue>",
				"<password>123456</password>",
				"<userName>SUPERAPP</userName>",
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewAccountLookup(tc.param)
			for _, expectStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectStr)
			}
		})
	}
}
