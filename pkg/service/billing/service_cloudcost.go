package billing

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetCloudCosts retrieves a list of costs
func (s *Service) GetCloudCosts(parameters connection.APIRequestParameters) ([]CloudCost, error) {
	return connection.InvokeRequestAll(s.GetCloudCostsPaginated, parameters)
}

// GetCloudCostsPaginated retrieves a paginated list of costs
func (s *Service) GetCloudCostsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CloudCost], error) {
	body, err := connection.Get[[]CloudCost](s.connection, "/billing/v1/cloud-costs", parameters)
	return connection.NewPaginated(body, parameters, s.GetCloudCostsPaginated), err
}

// GetCloudCost retrieves a single cost by id
func (s *Service) GetCloudCost(costID int) (CloudCost, error) {
	if costID < 1 {
		return CloudCost{}, fmt.Errorf("invalid cost id")
	}
	body, err := connection.Get[CloudCost](s.connection, fmt.Sprintf("/billing/v1/cloud-costs/%d", costID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&CloudCostNotFoundError{ID: costID}))
	return body.Data, err
}
