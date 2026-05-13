package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetRegions retrieves a list of regions
func (s *Service) GetRegions(parameters connection.APIRequestParameters) ([]Region, error) {
	return connection.InvokeRequestAll(s.GetRegionsPaginated, parameters)
}

// GetRegionsPaginated retrieves a paginated list of regions
func (s *Service) GetRegionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Region], error) {
	body, err := connection.Get[[]Region](s.connection, "/ecloud/v2/regions", parameters)
	return connection.NewPaginated(body, parameters, s.GetRegionsPaginated), err
}

// GetRegion retrieves a single region by id
func (s *Service) GetRegion(regionID string) (Region, error) {
	if regionID == "" {
		return Region{}, fmt.Errorf("invalid region id")
	}
	body, err := connection.Get[Region](s.connection, fmt.Sprintf("/ecloud/v2/regions/%s", regionID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&RegionNotFoundError{ID: regionID}))
	return body.Data, err
}
