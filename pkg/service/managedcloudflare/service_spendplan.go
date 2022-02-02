package managedcloudflare

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetSpendPlans retrieves a list of spend plans
func (s *Service) GetSpendPlans(parameters connection.APIRequestParameters) ([]SpendPlan, error) {
	var plans []SpendPlan

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSpendPlansPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		plans = append(plans, response.(*PaginatedSpendPlan).Items...)
	}

	return plans, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetSpendPlansPaginated retrieves a paginated list of spend plans
func (s *Service) GetSpendPlansPaginated(parameters connection.APIRequestParameters) (*PaginatedSpendPlan, error) {
	body, err := s.getSpendPlansPaginatedResponseBody(parameters)

	return NewPaginatedSpendPlan(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSpendPlansPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getSpendPlansPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetSpendPlanSliceResponseBody, error) {
	body := &GetSpendPlanSliceResponseBody{}

	response, err := s.connection.Get("/managed-cloudflare/v2/spend-plans", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
