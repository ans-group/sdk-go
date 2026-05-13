package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) discountPlanRes() *resource.Resource[DiscountPlan, string] {
	return resource.NewStringResource[DiscountPlan](s.connection, "/ecloud/v2/discount-plans", "discount plan", func(id string) error {
		return &DiscountPlanNotFoundError{ID: id}
	})
}

// GetDiscountPlans retrieves a list of discount plans
func (s *Service) GetDiscountPlans(parameters connection.APIRequestParameters) ([]DiscountPlan, error) {
	return s.discountPlanRes().List(parameters)
}

// GetDiscountPlansPaginated retrieves a paginated list of discount plans
func (s *Service) GetDiscountPlansPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[DiscountPlan], error) {
	return s.discountPlanRes().ListPaginated(parameters)
}

// GetDiscountPlan retrieves a single discount plan by id
func (s *Service) GetDiscountPlan(discID string) (DiscountPlan, error) {
	return s.discountPlanRes().Get(discID)
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
