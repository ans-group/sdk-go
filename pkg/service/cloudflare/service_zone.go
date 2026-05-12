package cloudflare

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetZones retrieves a list of zones
func (s *Service) GetZones(parameters connection.APIRequestParameters) ([]Zone, error) {
	return connection.InvokeRequestAll(s.GetZonesPaginated, parameters)
}

// GetZonesPaginated retrieves a paginated list of zones
func (s *Service) GetZonesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Zone], error) {
	body, err := connection.Get[[]Zone](s.connection, "/cloudflare/v1/zones", parameters)
	return connection.NewPaginated(body, parameters, s.GetZonesPaginated), err
}

// GetZone retrieves a single zone by id
func (s *Service) GetZone(zoneID string) (Zone, error) {
	if zoneID == "" {
		return Zone{}, fmt.Errorf("invalid zone id")
	}
	body, err := connection.Get[Zone](s.connection, fmt.Sprintf("/cloudflare/v1/zones/%s", zoneID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ZoneNotFoundError{ID: zoneID}))
	return body.Data, err
}

// CreateZone creates a new zone
func (s *Service) CreateZone(req CreateZoneRequest) (string, error) {
	body, err := connection.Post[Zone](s.connection, "/cloudflare/v1/zones", &req)
	return body.Data.ID, err
}

// PatchZone updates a zone
func (s *Service) PatchZone(zoneID string, req PatchZoneRequest) error {
	if zoneID == "" {
		return fmt.Errorf("invalid zone id")
	}
	_, err := connection.Post[struct{}](s.connection, fmt.Sprintf("/cloudflare/v1/zones/%s", zoneID), &req)
	return err
}

// DeleteZone removes a single zone by id
func (s *Service) DeleteZone(zoneID string) error {
	if zoneID == "" {
		return fmt.Errorf("invalid zone id")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/cloudflare/v1/zones/%s", zoneID), connection.APIRequestParameters{}, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&ZoneNotFoundError{ID: zoneID}))
}
