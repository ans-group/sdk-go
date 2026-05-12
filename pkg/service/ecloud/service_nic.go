package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetNICs retrieves a list of nics
func (s *Service) GetNICs(parameters connection.APIRequestParameters) ([]NIC, error) {
	return connection.InvokeRequestAll(s.GetNICsPaginated, parameters)
}

// GetNICsPaginated retrieves a paginated list of nics
func (s *Service) GetNICsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[NIC], error) {
	body, err := connection.Get[[]NIC](s.connection, "/ecloud/v2/nics", parameters)
	return connection.NewPaginated(body, parameters, s.GetNICsPaginated), err
}

// GetNIC retrieves a single nic by id
func (s *Service) GetNIC(nicID string) (NIC, error) {
	if nicID == "" {
		return NIC{}, fmt.Errorf("invalid nic id")
	}
	body, err := connection.Get[NIC](s.connection, fmt.Sprintf("/ecloud/v2/nics/%s", nicID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&NICNotFoundError{ID: nicID}))
	return body.Data, err
}

// GetNICTasks retrieves a list of NIC tasks
func (s *Service) GetNICTasks(nicID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetNICTasksPaginated(nicID, p)
	}, parameters)
}

// GetNICTasksPaginated retrieves a paginated list of NIC tasks
func (s *Service) GetNICTasksPaginated(nicID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	if nicID == "" {
		return nil, fmt.Errorf("invalid nic id")
	}
	body, err := connection.Get[[]Task](s.connection, fmt.Sprintf("/ecloud/v2/nics/%s/tasks", nicID), parameters, connection.NotFoundResponseHandler(&NICNotFoundError{ID: nicID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetNICTasksPaginated(nicID, p)
	}), err
}

// GetNICIPAddress retrieves a list of NIC IP addresses
func (s *Service) GetNICIPAddresses(nicID string, parameters connection.APIRequestParameters) ([]IPAddress, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[IPAddress], error) {
		return s.GetNICIPAddressesPaginated(nicID, p)
	}, parameters)
}

// GetNICIPAddressPaginated retrieves a paginated list of NIC IP addresses
func (s *Service) GetNICIPAddressesPaginated(nicID string, parameters connection.APIRequestParameters) (*connection.Paginated[IPAddress], error) {
	if nicID == "" {
		return nil, fmt.Errorf("invalid nic id")
	}
	body, err := connection.Get[[]IPAddress](s.connection, fmt.Sprintf("/ecloud/v2/nics/%s/ip-addresses", nicID), parameters, connection.NotFoundResponseHandler(&NICNotFoundError{ID: nicID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[IPAddress], error) {
		return s.GetNICIPAddressesPaginated(nicID, p)
	}), err
}

func (s *Service) AssignNICIPAddress(nicID string, req AssignIPAddressRequest) (string, error) {
	if nicID == "" {
		return "", fmt.Errorf("invalid nic id")
	}
	body, err := connection.Post[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/nics/%s/ip-addresses", nicID), &req, connection.NotFoundResponseHandler(&NICNotFoundError{ID: nicID}))
	return body.Data.TaskID, err
}

// UnassignNICIPAddress unassigns an IP Address from a resource
func (s *Service) UnassignNICIPAddress(nicID string, ipID string) (string, error) {
	if nicID == "" {
		return "", fmt.Errorf("invalid nic id")
	}
	if ipID == "" {
		return "", fmt.Errorf("invalid ip address id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/nics/%s/ip-addresses/%s", nicID, ipID), nil, connection.NotFoundResponseHandler(&NICNotFoundError{ID: nicID}))
	return body.Data.TaskID, err
}

// CreateNIC creates a new NIC
func (s *Service) CreateNIC(req CreateNICRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/nics", &req)
	return body.Data, err
}

// PatchNIC patches a NIC
func (s *Service) PatchNIC(nicID string, req PatchNICRequest) (TaskReference, error) {
	if nicID == "" {
		return TaskReference{}, fmt.Errorf("invalid nic id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/nics/%s", nicID), &req, connection.NotFoundResponseHandler(&NICNotFoundError{ID: nicID}))
	return body.Data, err
}

// DeleteNIC deletes a NIC
func (s *Service) DeleteNIC(NICID string) (string, error) {
	if NICID == "" {
		return "", fmt.Errorf("invalid nic id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/nics/%s", NICID), nil, connection.NotFoundResponseHandler(&NICNotFoundError{ID: NICID}))
	return body.Data.TaskID, err
}
