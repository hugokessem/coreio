package customerfetch

import (
	"testing"

	valueobject "github.com/hugokessem/coreio/lib/core/customer/customer_fetch/value_object"
)

func TestCustomerFetch(t *testing.T) {
	param := Params{
		Username:       "SUPERAPP",
		Password:       "123456",
		FetchBy:        valueobject.FetchByCustomerNumber.String(),
		CustomerNumber: "1045384696",
	}

	t.Logf("param: %+v", param)
	xmlRequest := NewCustomerFetch(param)
	t.Logf("XML Request: %s", xmlRequest)
}
