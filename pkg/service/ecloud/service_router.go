package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetRouters retrieves a list of routers
func (s *Service) GetRouters(parameters connection.APIRequestParameters) ([]Router, error) {
	var sites []Router

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRoutersPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, site := range response.(*PaginatedRouter).Items {
			sites = append(sites, site)
		}
	}

	return sites, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetRoutersPaginated retrieves a paginated list of routers
func (s *Service) GetRoutersPaginated(parameters connection.APIRequestParameters) (*PaginatedRouter, error) {
	body, err := s.getRoutersPaginatedResponseBody(parameters)

	return NewPaginatedRouter(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRoutersPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getRoutersPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetRouterSliceResponseBody, error) {
	body := &GetRouterSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/routers", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetRouter retrieves a single router by id
func (s *Service) GetRouter(routerID string) (Router, error) {
	body, err := s.getRouterResponseBody(routerID)

	return body.Data, err
}

func (s *Service) getRouterResponseBody(routerID string) (*GetRouterResponseBody, error) {
	body := &GetRouterResponseBody{}

	if routerID == "" {
		return body, fmt.Errorf("invalid router id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/routers/%s", routerID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &RouterNotFoundError{ID: routerID}
		}

		return nil
	})
}
