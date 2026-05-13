package billing

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) recurringCostRes() *resource.Resource[RecurringCost, int] {
	return resource.NewIntResource[RecurringCost](s.connection, "/billing/v1/recurring-costs", "cost",
		func(id int) error { return &RecurringCostNotFoundError{ID: id} })
}

// GetRecurringCosts retrieves a list of costs
func (s *Service) GetRecurringCosts(parameters connection.APIRequestParameters) ([]RecurringCost, error) {
	return s.recurringCostRes().List(parameters)
}

// GetRecurringCostsPaginated retrieves a paginated list of costs
func (s *Service) GetRecurringCostsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[RecurringCost], error) {
	return s.recurringCostRes().ListPaginated(parameters)
}

// GetRecurringCost retrieves a single cost by id
func (s *Service) GetRecurringCost(costID int) (RecurringCost, error) {
	return s.recurringCostRes().Get(costID)
}
