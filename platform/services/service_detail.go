package services

import (
	customerdetail "github.com/hugokessem/coreio/lib/core/customer/customer_detail"
	servicedetail "github.com/hugokessem/coreio/lib/core/service/service_detail"
)

type Params struct {
	CustomerNumber string
}

type CustomerServiceDetailResult struct {
	Success  bool
	Detail   []ServiceDetail
	Messages []string
}

type ServiceDetail struct {
	ServiceName        string
	ServiceDescription string
	MaximumAmount      string
	MaximumFrequency   string
}

type serviceDetail struct {
}

// func mergedDetail(serviceDetailResult *servicedetail.ServiceDetailResult, customerlimitfetchbycif *customerlimitfetchbycif.CustomerLimitViewResponse) *CustomerServiceDetailResult {
// 	if !serviceDetailResult.Success {
// 		return &CustomerServiceDetailResult{
// 			Success:  false,
// 			Messages: serviceDetailResult.Messages,
// 		}
// 	}

// 	if !customerlimitfetchbycif.Success {
// 		return &CustomerServiceDetailResult{
// 			Success:  false,
// 			Messages: customerlimitfetchbycif.Message,
// 		}
// 	}

// 	detail := make([]ServiceDetail, len(serviceDetailResult.Detail))
// 	for i, d := range serviceDetailResult.Detail {
// 		for j, k := range customerlimitfetchbycif.CustomerInfos {
// 			if d.ID == k.CustomerID {
// 				detail[i] = ServiceDetail{
// 					ServiceName:        d.ServiceDescription,
// 					ServiceDescription: d.Description,
// 					MaximumAmount:      d.MaximumAmount,
// 					MaximumFrequency:   d.MaximumFrequency,
// 				}

// 			}
// 		}
// 	}

// 	return &CustomerServiceDetailResult{
// 		Success:  true,
// 		Detail:   detail,
// 		Messages: []string{},
// 	}
// }

func (s *serviceDetail) ServiceDetail(param servicedetail.ServiceDetailParams) (*servicedetail.ServiceDetailResult, error) {
	result, err := s.ServiceDetail(param)
	if err != nil {
		return nil, err
	}

	return &servicedetail.ServiceDetailResult{
		Success:  result.Success,
		Detail:   result.Detail,
		Messages: result.Messages,
	}, nil
}

func (s *serviceDetail) CustomerDetail(param customerdetail.CustomerDetailParam) (*customerdetail.CustomerDetailResult, error) {
	result, err := s.CustomerDetail(param)
	if err != nil {
		return nil, err
	}

	return &customerdetail.CustomerDetailResult{
		Success:       result.Success,
		CustomerInfos: result.CustomerInfos,
		Message:       result.Message,
	}, nil
}
