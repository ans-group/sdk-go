package cloudflare

import (
	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetSubscriptions retrieves a list of subscriptions
func (s *Service) GetSubscriptions(parameters connection.APIRequestParameters) ([]Subscription, error) {
	return connection.InvokeRequestAll(s.GetSubscriptionsPaginated, parameters)
}

// GetSubscriptionsPaginated retrieves a paginated list of subscriptions
func (s *Service) GetSubscriptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Subscription], error) {
	body, err := connection.Get[[]Subscription](s.connection, "/cloudflare/v1/subscriptions", parameters)
	return connection.NewPaginated(body, parameters, s.GetSubscriptionsPaginated), err
}
