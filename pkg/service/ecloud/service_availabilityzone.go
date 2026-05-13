package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) availabilityZoneRes() *resource.Resource[AvailabilityZone, string] {
	return resource.NewStringResource[AvailabilityZone](s.connection, "/ecloud/v2/availability-zones", "az", func(id string) error {
		return &AvailabilityZoneNotFoundError{ID: id}
	})
}

// GetAvailabilityZones retrieves a list of azs
func (s *Service) GetAvailabilityZones(parameters connection.APIRequestParameters) ([]AvailabilityZone, error) {
	return s.availabilityZoneRes().List(parameters)
}

// GetAvailabilityZonesPaginated retrieves a paginated list of azs
func (s *Service) GetAvailabilityZonesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[AvailabilityZone], error) {
	return s.availabilityZoneRes().ListPaginated(parameters)
}

// GetAvailabilityZone retrieves a single az by id
func (s *Service) GetAvailabilityZone(azID string) (AvailabilityZone, error) {
	return s.availabilityZoneRes().Get(azID)
}

// GetAvailabilityZoneIOPSTiers retrieves a list of azs
func (s *Service) GetAvailabilityZoneIOPSTiers(azID string, parameters connection.APIRequestParameters) ([]IOPSTier, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[IOPSTier], error) {
		return s.GetAvailabilityZoneIOPSTiersPaginated(azID, p)
	}, parameters)
}

// GetAvailabilityZoneIOPSTiersPaginated retrieves a paginated list of azs
func (s *Service) GetAvailabilityZoneIOPSTiersPaginated(azID string, parameters connection.APIRequestParameters) (*connection.Paginated[IOPSTier], error) {
	if azID == "" {
		return nil, fmt.Errorf("invalid az id")
	}
	body, err := connection.Get[[]IOPSTier](s.connection, fmt.Sprintf("/ecloud/v2/availability-zones/%s/iops", azID), parameters, connection.NotFoundResponseHandler(&AvailabilityZoneNotFoundError{ID: azID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[IOPSTier], error) {
		return s.GetAvailabilityZoneIOPSTiersPaginated(azID, p)
	}), err
}
