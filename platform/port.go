package platform

import (
	customerdetail "github.com/hugokessem/coreio/lib/core/customer/customer_detail"
	servicedetail "github.com/hugokessem/coreio/lib/core/service/service_detail"
)

type CoreServices interface {
	ServiceDetail(param servicedetail.ServiceDetailParams) (*servicedetail.ServiceDetailResult, error)
	CustomerDetail(param customerdetail.CustomerDetailParam) (*customerdetail.CustomerDetailResult, error)
}
