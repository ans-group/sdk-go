package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetResourceTiers retrieves a list of resource tiers
func (s *Service) GetResourceTiers(parameters connection.APIRequestParameters) ([]ResourceTier, error) {
	return connection.InvokeRequestAll(s.GetResourceTiersPaginated, parameters)
}

// GetResourceTiersPaginated retrieves a paginated list of resource tiers
func (s *Service) GetResourceTiersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ResourceTier], error) {
	body, err := connection.Get[[]ResourceTier](s.connection, "/ecloud/v2/resource-tiers", parameters)
	return connection.NewPaginated(body, parameters, s.GetResourceTiersPaginated), err
}

// GetResourceTier retrieves a single resource tier by id
func (s *Service) GetResourceTier(tierID string) (ResourceTier, error) {
	if tierID == "" {
		return ResourceTier{}, fmt.Errorf("invalid resource tier id")
	}
	body, err := connection.Get[ResourceTier](s.connection, fmt.Sprintf("/ecloud/v2/resource-tiers/%s", tierID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ResourceTierNotFoundError{ID: tierID}))
	return body.Data, err
}
