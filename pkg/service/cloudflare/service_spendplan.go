package cloudflare

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) spendPlanRes() *resource.Resource[SpendPlan, string] {
	return resource.NewStringResource[SpendPlan](s.connection, "/cloudflare/v1/spend-plans", "spend plan",
		func(id string) error { return fmt.Errorf("spend plan not found: %s", id) })
}

// GetSpendPlans retrieves a list of spend plans
func (s *Service) GetSpendPlans(parameters connection.APIRequestParameters) ([]SpendPlan, error) {
	return s.spendPlanRes().List(parameters)
}

// GetSpendPlansPaginated retrieves a paginated list of spend plans
func (s *Service) GetSpendPlansPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[SpendPlan], error) {
	return s.spendPlanRes().ListPaginated(parameters)
}
