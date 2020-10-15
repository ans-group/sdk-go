package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetInstances retrieves a list of instances
func (s *Service) GetInstances(parameters connection.APIRequestParameters) ([]Instance, error) {
	var sites []Instance

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetInstancesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, site := range response.(*PaginatedInstance).Items {
			sites = append(sites, site)
		}
	}

	return sites, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetInstancesPaginated retrieves a paginated list of instances
func (s *Service) GetInstancesPaginated(parameters connection.APIRequestParameters) (*PaginatedInstance, error) {
	body, err := s.getInstancesPaginatedResponseBody(parameters)

	return NewPaginatedInstance(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetInstancesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getInstancesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetInstanceSliceResponseBody, error) {
	body := &GetInstanceSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/instances", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetInstance retrieves a single instance by id
func (s *Service) GetInstance(instanceID string) (Instance, error) {
	body, err := s.getInstanceResponseBody(instanceID)

	return body.Data, err
}

func (s *Service) getInstanceResponseBody(instanceID string) (*GetInstanceResponseBody, error) {
	body := &GetInstanceResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/instances/%s", instanceID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// CreateInstance creates a new instance
func (s *Service) CreateInstance(req CreateInstanceRequest) (string, error) {
	body, err := s.createInstanceResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createInstanceResponseBody(req CreateInstanceRequest) (*GetInstanceResponseBody, error) {
	body := &GetInstanceResponseBody{}

	response, err := s.connection.Post("/ecloud/v2/instances", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// DeleteInstance removes a virtual machine
func (s *Service) DeleteInstance(instanceID string) error {
	_, err := s.deleteInstanceResponseBody(instanceID)

	return err
}

func (s *Service) deleteInstanceResponseBody(instanceID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/instances/%s", instanceID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}
