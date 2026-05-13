package cloudflare

import (
	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetSpendPlans retrieves a list of spend plans
func (s *Service) GetSpendPlans(parameters connection.APIRequestParameters) ([]SpendPlan, error) {
	return connection.InvokeRequestAll(s.GetSpendPlansPaginated, parameters)
}

// GetSpendPlansPaginated retrieves a paginated list of spend plans
func (s *Service) GetSpendPlansPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[SpendPlan], error) {
	body, err := connection.Get[[]SpendPlan](s.connection, "/cloudflare/v1/spend-plans", parameters)
	return connection.NewPaginated(body, parameters, s.GetSpendPlansPaginated), err
}
