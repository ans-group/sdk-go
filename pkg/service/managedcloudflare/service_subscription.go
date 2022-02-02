package managedcloudflare

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetSubscriptions retrieves a list of subscriptions
func (s *Service) GetSubscriptions(parameters connection.APIRequestParameters) ([]Subscription, error) {
	var subscriptions []Subscription

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSubscriptionsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		subscriptions = append(subscriptions, response.(*PaginatedSubscription).Items...)
	}

	return subscriptions, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetSubscriptionsPaginated retrieves a paginated list of subscriptions
func (s *Service) GetSubscriptionsPaginated(parameters connection.APIRequestParameters) (*PaginatedSubscription, error) {
	body, err := s.getSubscriptionsPaginatedResponseBody(parameters)

	return NewPaginatedSubscription(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSubscriptionsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getSubscriptionsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetSubscriptionSliceResponseBody, error) {
	body := &GetSubscriptionSliceResponseBody{}

	response, err := s.connection.Get("/managed-cloudflare/v1/subscriptions", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
