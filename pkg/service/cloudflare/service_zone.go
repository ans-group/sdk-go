package cloudflare

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetZones retrieves a list of zones
func (s *Service) GetZones(parameters connection.APIRequestParameters) ([]Zone, error) {
	var zones []Zone

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetZonesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		zones = append(zones, response.(*PaginatedZone).Items...)
	}

	return zones, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetZonesPaginated retrieves a paginated list of zones
func (s *Service) GetZonesPaginated(parameters connection.APIRequestParameters) (*PaginatedZone, error) {
	body, err := s.getZonesPaginatedResponseBody(parameters)

	return NewPaginatedZone(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetZonesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getZonesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetZoneSliceResponseBody, error) {
	body := &GetZoneSliceResponseBody{}

	response, err := s.connection.Get("/cloudflare/v1/zones", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetZone retrieves a single zone by id
func (s *Service) GetZone(zoneID string) (Zone, error) {
	body, err := s.getZoneResponseBody(zoneID)

	return body.Data, err
}

func (s *Service) getZoneResponseBody(zoneID string) (*GetZoneResponseBody, error) {
	body := &GetZoneResponseBody{}

	if zoneID == "" {
		return body, fmt.Errorf("invalid zone id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/cloudflare/v1/zones/%s", zoneID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ZoneNotFoundError{ID: zoneID}
		}

		return nil
	})
}

// CreateZone creates a new zone
func (s *Service) CreateZone(req CreateZoneRequest) (string, error) {
	body, err := s.createZoneResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createZoneResponseBody(req CreateZoneRequest) (*GetZoneResponseBody, error) {
	body := &GetZoneResponseBody{}

	response, err := s.connection.Post("/cloudflare/v1/zones", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchZone updates a zone
func (s *Service) PatchZone(zoneID string, req PatchZoneRequest) error {
	_, err := s.patchZoneResponseBody(zoneID, req)

	return err
}

func (s *Service) patchZoneResponseBody(zoneID string, req PatchZoneRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if zoneID == "" {
		return body, fmt.Errorf("invalid zone id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/cloudflare/v1/zones/%s", zoneID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// DeleteZone removes a single zone by id
func (s *Service) DeleteZone(zoneID string) error {
	_, err := s.deleteZoneResponseBody(zoneID)

	return err
}

func (s *Service) deleteZoneResponseBody(zoneID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if zoneID == "" {
		return body, fmt.Errorf("invalid zone id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/cloudflare/v1/zones/%s", zoneID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ZoneNotFoundError{ID: zoneID}
		}

		return nil
	})
}