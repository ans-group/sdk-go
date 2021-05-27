package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetHosts retrieves a list of hosts
func (s *Service) GetHosts(parameters connection.APIRequestParameters) ([]Host, error) {
	var hosts []Host

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetHostsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, host := range response.(*PaginatedHost).Items {
			hosts = append(hosts, host)
		}
	}

	return hosts, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetHostsPaginated retrieves a paginated list of hosts
func (s *Service) GetHostsPaginated(parameters connection.APIRequestParameters) (*PaginatedHost, error) {
	body, err := s.getHostsPaginatedResponseBody(parameters)

	return NewPaginatedHost(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetHostsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getHostsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetHostSliceResponseBody, error) {
	body := &GetHostSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/hosts", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetHost retrieves a single host by id
func (s *Service) GetHost(hostID string) (Host, error) {
	body, err := s.getHostResponseBody(hostID)

	return body.Data, err
}

func (s *Service) getHostResponseBody(hostID string) (*GetHostResponseBody, error) {
	body := &GetHostResponseBody{}

	if hostID == "" {
		return body, fmt.Errorf("invalid host id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/hosts/%s", hostID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HostNotFoundError{ID: hostID}
		}

		return nil
	})
}

// CreateHost creates a host
func (s *Service) CreateHost(req CreateHostRequest) (string, error) {
	body, err := s.createHostResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createHostResponseBody(req CreateHostRequest) (*GetHostResponseBody, error) {
	body := &GetHostResponseBody{}

	response, err := s.connection.Post("/ecloud/v2/hosts", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchHost patches a host
func (s *Service) PatchHost(hostID string, req PatchHostRequest) error {
	_, err := s.patchHostResponseBody(hostID, req)

	return err
}

func (s *Service) patchHostResponseBody(hostID string, req PatchHostRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if hostID == "" {
		return body, fmt.Errorf("invalid host id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/hosts/%s", hostID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HostNotFoundError{ID: hostID}
		}

		return nil
	})
}

// DeleteHost deletes a host
func (s *Service) DeleteHost(hostID string) error {
	_, err := s.deleteHostResponseBody(hostID)

	return err
}

func (s *Service) deleteHostResponseBody(hostID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if hostID == "" {
		return body, fmt.Errorf("invalid host id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/hosts/%s", hostID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HostNotFoundError{ID: hostID}
		}

		return nil
	})
}

// GetHostTasks retrieves a list of Host tasks
func (s *Service) GetHostTasks(hostID string, parameters connection.APIRequestParameters) ([]Task, error) {
	var tasks []Task

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetHostTasksPaginated(hostID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, task := range response.(*PaginatedTask).Items {
			tasks = append(tasks, task)
		}
	}

	return tasks, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetHostTasksPaginated retrieves a paginated list of Host tasks
func (s *Service) GetHostTasksPaginated(hostID string, parameters connection.APIRequestParameters) (*PaginatedTask, error) {
	body, err := s.getHostTasksPaginatedResponseBody(hostID, parameters)

	return NewPaginatedTask(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetHostTasksPaginated(hostID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getHostTasksPaginatedResponseBody(hostID string, parameters connection.APIRequestParameters) (*GetTaskSliceResponseBody, error) {
	body := &GetTaskSliceResponseBody{}

	if hostID == "" {
		return body, fmt.Errorf("invalid host id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/hosts/%s/tasks", hostID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HostNotFoundError{ID: hostID}
		}

		return nil
	})
}
