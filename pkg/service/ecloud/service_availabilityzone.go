package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetAvailabilityZones retrieves a list of azs
func (s *Service) GetAvailabilityZones(parameters connection.APIRequestParameters) ([]AvailabilityZone, error) {
	return connection.InvokeRequestAll(s.GetAvailabilityZonesPaginated, parameters)
}

// GetAvailabilityZonesPaginated retrieves a paginated list of azs
func (s *Service) GetAvailabilityZonesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[AvailabilityZone], error) {
	body, err := connection.Get[[]AvailabilityZone](s.connection, "/ecloud/v2/availability-zones", parameters)
	return connection.NewPaginated(body, parameters, s.GetAvailabilityZonesPaginated), err
}

// GetAvailabilityZone retrieves a single az by id
func (s *Service) GetAvailabilityZone(azID string) (AvailabilityZone, error) {
	if azID == "" {
		return AvailabilityZone{}, fmt.Errorf("invalid az id")
	}
	body, err := connection.Get[AvailabilityZone](s.connection, fmt.Sprintf("/ecloud/v2/availability-zones/%s", azID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&AvailabilityZoneNotFoundError{ID: azID}))
	return body.Data, err
}

// GetAvailabilityZones retrieves a list of azs
func (s *Service) GetAvailabilityZoneIOPSTiers(azID string, parameters connection.APIRequestParameters) ([]IOPSTier, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[IOPSTier], error) {
		return s.GetAvailabilityZoneIOPSTiersPaginated(azID, p)
	}, parameters)
}

// GetAvailabilityZonesPaginated retrieves a paginated list of azs
func (s *Service) GetAvailabilityZoneIOPSTiersPaginated(azID string, parameters connection.APIRequestParameters) (*connection.Paginated[IOPSTier], error) {
	if azID == "" {
		return nil, fmt.Errorf("invalid az id")
	}
	body, err := connection.Get[[]IOPSTier](s.connection, fmt.Sprintf("/ecloud/v2/availability-zones/%s/iops", azID), parameters, connection.NotFoundResponseHandler(&AvailabilityZoneNotFoundError{ID: azID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[IOPSTier], error) {
		return s.GetAvailabilityZoneIOPSTiersPaginated(azID, p)
	}), err
}
