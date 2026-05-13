package cloudflare

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) subscriptionRes() *resource.Resource[Subscription, string] {
	return resource.NewStringResource[Subscription](s.connection, "/cloudflare/v1/subscriptions", "subscription",
		func(id string) error { return fmt.Errorf("subscription not found: %s", id) })
}

// GetSubscriptions retrieves a list of subscriptions
func (s *Service) GetSubscriptions(parameters connection.APIRequestParameters) ([]Subscription, error) {
	return s.subscriptionRes().List(parameters)
}

// GetSubscriptionsPaginated retrieves a paginated list of subscriptions
func (s *Service) GetSubscriptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Subscription], error) {
	return s.subscriptionRes().ListPaginated(parameters)
}
