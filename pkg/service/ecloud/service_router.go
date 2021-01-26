package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetRouters retrieves a list of routers
func (s *Service) GetRouters(parameters connection.APIRequestParameters) ([]Router, error) {
	var routers []Router

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRoutersPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, router := range response.(*PaginatedRouter).Items {
			routers = append(routers, router)
		}
	}

	return routers, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
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

// CreateRouter creates a new Router
func (s *Service) CreateRouter(req CreateRouterRequest) (string, error) {
	body, err := s.createRouterResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createRouterResponseBody(req CreateRouterRequest) (*GetRouterResponseBody, error) {
	body := &GetRouterResponseBody{}

	response, err := s.connection.Post("/ecloud/v2/routers", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchRouter patches a Router
func (s *Service) PatchRouter(routerID string, req PatchRouterRequest) error {
	_, err := s.patchRouterResponseBody(routerID, req)

	return err
}

func (s *Service) patchRouterResponseBody(routerID string, req PatchRouterRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if routerID == "" {
		return body, fmt.Errorf("invalid router id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/routers/%s", routerID), &req)
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

// DeleteRouter deletes a Router
func (s *Service) DeleteRouter(routerID string) error {
	_, err := s.deleteRouterResponseBody(routerID)

	return err
}

func (s *Service) deleteRouterResponseBody(routerID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if routerID == "" {
		return body, fmt.Errorf("invalid router id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/routers/%s", routerID), nil)
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
