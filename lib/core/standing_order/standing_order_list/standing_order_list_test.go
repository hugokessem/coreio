package standingorderlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListStandingOrderGeneratedXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate List Standing Order",
			param: Params{
				Username:      "SUPERAPP",
				Password:      "123456",
				AccountNumber: "10009876543",
			},
			expect: []string{
				"<criteriaValue>10009876543</criteriaValue>",
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewListStandingOrder(tc.param)
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}
		})
	}
}
