package ecloud

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) regionRes() *resource.Resource[Region, string] {
	return resource.NewStringResource[Region](s.connection, "/ecloud/v2/regions", "region", func(id string) error {
		return &RegionNotFoundError{ID: id}
	})
}

// GetRegions retrieves a list of regions
func (s *Service) GetRegions(parameters connection.APIRequestParameters) ([]Region, error) {
	return s.regionRes().List(parameters)
}

// GetRegionsPaginated retrieves a paginated list of regions
func (s *Service) GetRegionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Region], error) {
	return s.regionRes().ListPaginated(parameters)
}

// GetRegion retrieves a single region by id
func (s *Service) GetRegion(regionID string) (Region, error) {
	return s.regionRes().Get(regionID)
}
