package billing

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetRecurringCosts retrieves a list of costs
func (s *Service) GetRecurringCosts(parameters connection.APIRequestParameters) ([]RecurringCost, error) {
	return connection.InvokeRequestAll(s.GetRecurringCostsPaginated, parameters)
}

// GetRecurringCostsPaginated retrieves a paginated list of costs
func (s *Service) GetRecurringCostsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[RecurringCost], error) {
	body, err := connection.Get[[]RecurringCost](s.connection, "/billing/v1/recurring-costs", parameters)
	return connection.NewPaginated(body, parameters, s.GetRecurringCostsPaginated), err
}

// GetRecurringCost retrieves a single cost by id
func (s *Service) GetRecurringCost(costID int) (RecurringCost, error) {
	if costID < 1 {
		return RecurringCost{}, fmt.Errorf("invalid cost id")
	}
	body, err := connection.Get[RecurringCost](s.connection, fmt.Sprintf("/billing/v1/recurring-costs/%d", costID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&RecurringCostNotFoundError{ID: costID}))
	return body.Data, err
}
