package ministatementbylimit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiniStatementGeneratedXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate Mini Statement",
			param: Params{
				Username:            "SUPERAPP",
				Password:            "123456",
				AccountNumber:       "1000030677308",
				NumberOfTransaction: "4",
			},
			expect: []string{
				`<criteriaValue>1000030677308</criteriaValue>`,
				`<criteriaValue>4</criteriaValue>`,
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewMiniStatement(tc.param)
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}
		})
	}
}
