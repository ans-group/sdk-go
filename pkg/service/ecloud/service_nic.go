package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetNICs retrieves a list of nics
func (s *Service) GetNICs(parameters connection.APIRequestParameters) ([]NIC, error) {
	var nics []NIC

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetNICsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, nic := range response.(*PaginatedNIC).Items {
			nics = append(nics, nic)
		}
	}

	return nics, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetNICsPaginated retrieves a paginated list of nics
func (s *Service) GetNICsPaginated(parameters connection.APIRequestParameters) (*PaginatedNIC, error) {
	body, err := s.getNICsPaginatedResponseBody(parameters)

	return NewPaginatedNIC(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetNICsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getNICsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetNICSliceResponseBody, error) {
	body := &GetNICSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/nics", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetNIC retrieves a single nic by id
func (s *Service) GetNIC(nicID string) (NIC, error) {
	body, err := s.getNICResponseBody(nicID)

	return body.Data, err
}

func (s *Service) getNICResponseBody(nicID string) (*GetNICResponseBody, error) {
	body := &GetNICResponseBody{}

	if nicID == "" {
		return body, fmt.Errorf("invalid nic id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/nics/%s", nicID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NICNotFoundError{ID: nicID}
		}

		return nil
	})
}

// GetNICTasks retrieves a list of NIC tasks
func (s *Service) GetNICTasks(nicID string, parameters connection.APIRequestParameters) ([]Task, error) {
	var tasks []Task

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetNICTasksPaginated(nicID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, task := range response.(*PaginatedTask).Items {
			tasks = append(tasks, task)
		}
	}

	return tasks, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetNICTasksPaginated retrieves a paginated list of NIC tasks
func (s *Service) GetNICTasksPaginated(nicID string, parameters connection.APIRequestParameters) (*PaginatedTask, error) {
	body, err := s.getNICTasksPaginatedResponseBody(nicID, parameters)

	return NewPaginatedTask(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetNICTasksPaginated(nicID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getNICTasksPaginatedResponseBody(nicID string, parameters connection.APIRequestParameters) (*GetTaskSliceResponseBody, error) {
	body := &GetTaskSliceResponseBody{}

	if nicID == "" {
		return body, fmt.Errorf("invalid nic id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/nics/%s/tasks", nicID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NICNotFoundError{ID: nicID}
		}

		return nil
	})
}

// GetNICIPAddress retrieves a list of NIC IP addresses
func (s *Service) GetNICIPAddresses(nicID string, parameters connection.APIRequestParameters) ([]IPAddress, error) {
	var ips []IPAddress

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetNICIPAddressesPaginated(nicID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, ip := range response.(*PaginatedIPAddress).Items {
			ips = append(ips, ip)
		}
	}

	return ips, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetNICIPAddressPaginated retrieves a paginated list of NIC IP addresses
func (s *Service) GetNICIPAddressesPaginated(nicID string, parameters connection.APIRequestParameters) (*PaginatedIPAddress, error) {
	body, err := s.getNICIPAddressesPaginatedResponseBody(nicID, parameters)

	return NewPaginatedIPAddress(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetNICIPAddressesPaginated(nicID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getNICIPAddressesPaginatedResponseBody(nicID string, parameters connection.APIRequestParameters) (*GetIPAddressSliceResponseBody, error) {
	body := &GetIPAddressSliceResponseBody{}

	if nicID == "" {
		return body, fmt.Errorf("invalid nic id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/nics/%s/ip-addresses", nicID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NICNotFoundError{ID: nicID}
		}

		return nil
	})
}

func (s *Service) AssignNICIPAddress(nicID string, req AssignIPAddressRequest) (string, error) {
	body, err := s.assignNICIPAddressResponseBody(nicID, req)

	return body.Data.TaskID, err
}

func (s *Service) assignNICIPAddressResponseBody(nicID string, req AssignIPAddressRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if nicID == "" {
		return body, fmt.Errorf("invalid nic id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v2/nics/%s/ip-addresses", nicID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NICNotFoundError{ID: nicID}
		}

		return nil
	})
}

// UnassignNICIPAddress unassigns an IP Address from a resource
func (s *Service) UnassignNICIPAddress(nicID string, ipID string) (string, error) {
	body, err := s.unassignNICIPAddressResponseBody(nicID, ipID)

	return body.Data.TaskID, err
}

func (s *Service) unassignNICIPAddressResponseBody(nicID string, ipID string) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if nicID == "" {
		return body, fmt.Errorf("invalid nic id")
	}

	if ipID == "" {
		return body, fmt.Errorf("invalid ip address id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/nics/%s/ip-addresses/%s", nicID, ipID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NICNotFoundError{ID: nicID}
		}

		return nil
	})
}
