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

func (s *Service) availabilityZoneIOPSTierRes() *resource.SubResourceList[IOPSTier, string] {
	return resource.NewStringSubResourceList[IOPSTier](s.connection,
		func(azID string) string { return fmt.Sprintf("/ecloud/v2/availability-zones/%s/iops", azID) },
		"az", "id", func(azID string) error { return &AvailabilityZoneNotFoundError{ID: azID} })
}

// GetAvailabilityZoneIOPSTiers retrieves a list of azs
func (s *Service) GetAvailabilityZoneIOPSTiers(azID string, parameters connection.APIRequestParameters) ([]IOPSTier, error) {
	return s.availabilityZoneIOPSTierRes().List(azID, parameters)
}

// GetAvailabilityZoneIOPSTiersPaginated retrieves a paginated list of azs
func (s *Service) GetAvailabilityZoneIOPSTiersPaginated(azID string, parameters connection.APIRequestParameters) (*connection.Paginated[IOPSTier], error) {
	return s.availabilityZoneIOPSTierRes().ListPaginated(azID, parameters)
}
