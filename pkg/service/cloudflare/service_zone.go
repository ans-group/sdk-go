package cloudflare

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) zoneRes() *resource.Resource[Zone, string] {
	return resource.NewStringResource[Zone](s.connection, "/cloudflare/v1/zones", "zone",
		func(id string) error { return &ZoneNotFoundError{ID: id} })
}

// GetZones retrieves a list of zones
func (s *Service) GetZones(parameters connection.APIRequestParameters) ([]Zone, error) {
	return s.zoneRes().List(parameters)
}

// GetZonesPaginated retrieves a paginated list of zones
func (s *Service) GetZonesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Zone], error) {
	return s.zoneRes().ListPaginated(parameters)
}

// GetZone retrieves a single zone by id
func (s *Service) GetZone(zoneID string) (Zone, error) {
	return s.zoneRes().Get(zoneID)
}

// CreateZone creates a new zone
func (s *Service) CreateZone(req CreateZoneRequest) (string, error) {
	data, err := s.zoneRes().Create(&req)
	return data.ID, err
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
	return s.zoneRes().Delete(zoneID)
}
