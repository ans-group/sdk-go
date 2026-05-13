package draas

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) billingTypeRes() *resource.Resource[BillingType, string] {
	return resource.NewStringResource[BillingType](s.connection, "/draas/v1/billing-types", "billing type",
		func(id string) error { return &BillingTypeNotFoundError{ID: id} })
}

// GetBillingTypes retrieves a list of solutions
func (s *Service) GetBillingTypes(parameters connection.APIRequestParameters) ([]BillingType, error) {
	return s.billingTypeRes().List(parameters)
}

// GetBillingTypesPaginated retrieves a paginated list of solutions
func (s *Service) GetBillingTypesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[BillingType], error) {
	return s.billingTypeRes().ListPaginated(parameters)
}

// GetBillingType retrieves a single solution by id
func (s *Service) GetBillingType(billingTypeID string) (BillingType, error) {
	return s.billingTypeRes().Get(billingTypeID)
}
