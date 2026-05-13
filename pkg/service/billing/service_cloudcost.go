package billing

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) cloudCostRes() *resource.Resource[CloudCost, int] {
	return resource.NewIntResource[CloudCost](s.connection, "/billing/v1/cloud-costs", "cost",
		func(id int) error { return &CloudCostNotFoundError{ID: id} })
}

// GetCloudCosts retrieves a list of costs
func (s *Service) GetCloudCosts(parameters connection.APIRequestParameters) ([]CloudCost, error) {
	return s.cloudCostRes().List(parameters)
}

// GetCloudCostsPaginated retrieves a paginated list of costs
func (s *Service) GetCloudCostsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CloudCost], error) {
	return s.cloudCostRes().ListPaginated(parameters)
}

// GetCloudCost retrieves a single cost by id
func (s *Service) GetCloudCost(costID int) (CloudCost, error) {
	return s.cloudCostRes().Get(costID)
}
