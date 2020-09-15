package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetNetworks retrieves a list of networks
func (s *Service) GetNetworks(parameters connection.APIRequestParameters) ([]Network, error) {
	var sites []Network

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetNetworksPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, site := range response.(*PaginatedNetwork).Items {
			sites = append(sites, site)
		}
	}

	return sites, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetNetworksPaginated retrieves a paginated list of networks
func (s *Service) GetNetworksPaginated(parameters connection.APIRequestParameters) (*PaginatedNetwork, error) {
	body, err := s.getNetworksPaginatedResponseBody(parameters)

	return NewPaginatedNetwork(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetNetworksPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getNetworksPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetNetworkSliceResponseBody, error) {
	body := &GetNetworkSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/networks", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetNetwork retrieves a single network by id
func (s *Service) GetNetwork(networkID string) (Network, error) {
	body, err := s.getNetworkResponseBody(networkID)

	return body.Data, err
}

func (s *Service) getNetworkResponseBody(networkID string) (*GetNetworkResponseBody, error) {
	body := &GetNetworkResponseBody{}

	if networkID == "" {
		return body, fmt.Errorf("invalid network id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/networks/%s", networkID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NetworkNotFoundError{ID: networkID}
		}

		return nil
	})
}
