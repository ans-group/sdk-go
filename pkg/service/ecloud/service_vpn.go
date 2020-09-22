package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetVPNs retrieves a list of vpns
func (s *Service) GetVPNs(parameters connection.APIRequestParameters) ([]VPN, error) {
	var sites []VPN

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVPNsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, site := range response.(*PaginatedVPN).Items {
			sites = append(sites, site)
		}
	}

	return sites, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetVPNsPaginated retrieves a paginated list of vpns
func (s *Service) GetVPNsPaginated(parameters connection.APIRequestParameters) (*PaginatedVPN, error) {
	body, err := s.getVPNsPaginatedResponseBody(parameters)

	return NewPaginatedVPN(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVPNsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getVPNsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetVPNSliceResponseBody, error) {
	body := &GetVPNSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/vpns", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetVPN retrieves a single vpn by id
func (s *Service) GetVPN(vpnID string) (VPN, error) {
	body, err := s.getVPNResponseBody(vpnID)

	return body.Data, err
}

func (s *Service) getVPNResponseBody(vpnID string) (*GetVPNResponseBody, error) {
	body := &GetVPNResponseBody{}

	if vpnID == "" {
		return body, fmt.Errorf("invalid vpn id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/vpns/%s", vpnID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNNotFoundError{ID: vpnID}
		}

		return nil
	})
}

// CreateVPN creates a new VPN
func (s *Service) CreateVPN(req CreateVPNRequest) (string, error) {
	body, err := s.createVPNResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createVPNResponseBody(req CreateVPNRequest) (*GetVPNResponseBody, error) {
	body := &GetVPNResponseBody{}

	response, err := s.connection.Post("/ecloud/v2/vpns", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// DeleteVPN deletes a VPN
func (s *Service) DeleteVPN(vpnID string) error {
	_, err := s.deleteVPNResponseBody(vpnID)

	return err
}

func (s *Service) deleteVPNResponseBody(vpnID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if vpnID == "" {
		return body, fmt.Errorf("invalid vpn id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/vpns/%s", vpnID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNNotFoundError{ID: vpnID}
		}

		return nil
	})
}
