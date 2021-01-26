package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetFloatingIPs retrieves a list of floating ips
func (s *Service) GetFloatingIPs(parameters connection.APIRequestParameters) ([]FloatingIP, error) {
	var fips []FloatingIP

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFloatingIPsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, fip := range response.(*PaginatedFloatingIP).Items {
			fips = append(fips, fip)
		}
	}

	return fips, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetFloatingIPsPaginated retrieves a paginated list of floating ips
func (s *Service) GetFloatingIPsPaginated(parameters connection.APIRequestParameters) (*PaginatedFloatingIP, error) {
	body, err := s.getFloatingIPsPaginatedResponseBody(parameters)

	return NewPaginatedFloatingIP(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFloatingIPsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getFloatingIPsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetFloatingIPSliceResponseBody, error) {
	body := &GetFloatingIPSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/floating-ips", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetFloatingIP retrieves a single floating ip by id
func (s *Service) GetFloatingIP(floatingIPID string) (FloatingIP, error) {
	body, err := s.getFloatingIPResponseBody(floatingIPID)

	return body.Data, err
}

func (s *Service) getFloatingIPResponseBody(floatingIPID string) (*GetFloatingIPResponseBody, error) {
	body := &GetFloatingIPResponseBody{}

	if floatingIPID == "" {
		return body, fmt.Errorf("invalid floating ip id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/floating-ips/%s", floatingIPID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FloatingIPNotFoundError{ID: floatingIPID}
		}

		return nil
	})
}
