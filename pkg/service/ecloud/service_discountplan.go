package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetDiscountPlans retrieves a list of discount plans
func (s *Service) GetDiscountPlans(parameters connection.APIRequestParameters) ([]DiscountPlan, error) {
	return connection.InvokeRequestAll(s.GetDiscountPlansPaginated, parameters)
}

// GetDiscountPlansPaginated retrieves a paginated list of discount plans
func (s *Service) GetDiscountPlansPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[DiscountPlan], error) {
	body, err := connection.Get[[]DiscountPlan](s.connection, "/ecloud/v2/discount-plans", parameters)
	return connection.NewPaginated(body, parameters, s.GetDiscountPlansPaginated), err
}

// GetDiscountPlan retrieves a single discount plan by id
func (s *Service) GetDiscountPlan(discID string) (DiscountPlan, error) {
	if discID == "" {
		return DiscountPlan{}, fmt.Errorf("invalid discount plan id")
	}
	body, err := connection.Get[DiscountPlan](s.connection, fmt.Sprintf("/ecloud/v2/discount-plans/%s", discID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&DiscountPlanNotFoundError{ID: discID}))
	return body.Data, err
}

// ApproveDiscountPlan approves a floating IP to a resource
func (s *Service) ApproveDiscountPlan(discID string) error {
	if discID == "" {
		return fmt.Errorf("invalid floating IP id")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/ecloud/v2/discount-plans/%s/approve", discID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DiscountPlanNotFoundError{ID: discID}))
}

// RejectDiscountPlan rejects a floating IP from a resource
func (s *Service) RejectDiscountPlan(discID string) error {
	if discID == "" {
		return fmt.Errorf("invalid floating IP id")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/ecloud/v2/discount-plans/%s/reject", discID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DiscountPlanNotFoundError{ID: discID}))
}
