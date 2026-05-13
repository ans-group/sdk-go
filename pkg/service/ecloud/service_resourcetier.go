package ecloud

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) resourceTierRes() *resource.Resource[ResourceTier, string] {
	return resource.NewStringResource[ResourceTier](s.connection, "/ecloud/v2/resource-tiers", "resource tier", func(id string) error {
		return &ResourceTierNotFoundError{ID: id}
	})
}

// GetResourceTiers retrieves a list of resource tiers
func (s *Service) GetResourceTiers(parameters connection.APIRequestParameters) ([]ResourceTier, error) {
	return s.resourceTierRes().List(parameters)
}

// GetResourceTiersPaginated retrieves a paginated list of resource tiers
func (s *Service) GetResourceTiersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ResourceTier], error) {
	return s.resourceTierRes().ListPaginated(parameters)
}

// GetResourceTier retrieves a single resource tier by id
func (s *Service) GetResourceTier(tierID string) (ResourceTier, error) {
	return s.resourceTierRes().Get(tierID)
}
