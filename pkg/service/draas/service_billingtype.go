package draas

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetBillingTypes retrieves a list of solutions
func (s *Service) GetBillingTypes(parameters connection.APIRequestParameters) ([]BillingType, error) {
	return connection.InvokeRequestAll(s.GetBillingTypesPaginated, parameters)
}

// GetBillingTypesPaginated retrieves a paginated list of solutions
func (s *Service) GetBillingTypesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[BillingType], error) {
	body, err := connection.Get[[]BillingType](s.connection, "/draas/v1/billing-types", parameters)
	return connection.NewPaginated(body, parameters, s.GetBillingTypesPaginated), err
}

// GetBillingType retrieves a single solution by id
func (s *Service) GetBillingType(billingTypeID string) (BillingType, error) {
	if billingTypeID == "" {
		return BillingType{}, fmt.Errorf("invalid billing type id")
	}
	body, err := connection.Get[BillingType](s.connection, fmt.Sprintf("/draas/v1/billing-types/%s", billingTypeID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&BillingTypeNotFoundError{ID: billingTypeID}))
	return body.Data, err
}
