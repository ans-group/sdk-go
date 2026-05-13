package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) nicRes() *resource.Resource[NIC, string] {
	return resource.NewStringResource[NIC](s.connection, "/ecloud/v2/nics", "nic", func(id string) error {
		return &NICNotFoundError{ID: id}
	})
}

// GetNICs retrieves a list of nics
func (s *Service) GetNICs(parameters connection.APIRequestParameters) ([]NIC, error) {
	return s.nicRes().List(parameters)
}

// GetNICsPaginated retrieves a paginated list of nics
func (s *Service) GetNICsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[NIC], error) {
	return s.nicRes().ListPaginated(parameters)
}

// GetNIC retrieves a single nic by id
func (s *Service) GetNIC(nicID string) (NIC, error) {
	return s.nicRes().Get(nicID)
}

func (s *Service) nicTasksRes() *resource.SubResourceList[Task, string] {
	return resource.NewStringSubResourceList[Task](s.connection,
		func(nicID string) string { return fmt.Sprintf("/ecloud/v2/nics/%s/tasks", nicID) },
		"nic", "id", func(nicID string) error { return &NICNotFoundError{ID: nicID} })
}

// GetNICTasks retrieves a list of NIC tasks
func (s *Service) GetNICTasks(nicID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return s.nicTasksRes().List(nicID, parameters)
}

// GetNICTasksPaginated retrieves a paginated list of NIC tasks
func (s *Service) GetNICTasksPaginated(nicID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	return s.nicTasksRes().ListPaginated(nicID, parameters)
}

func (s *Service) nicIPAddressRes() *resource.SubResourceList[IPAddress, string] {
	return resource.NewStringSubResourceList[IPAddress](s.connection,
		func(nicID string) string { return fmt.Sprintf("/ecloud/v2/nics/%s/ip-addresses", nicID) },
		"nic", "id", func(nicID string) error { return &NICNotFoundError{ID: nicID} })
}

// GetNICIPAddress retrieves a list of NIC IP addresses
func (s *Service) GetNICIPAddresses(nicID string, parameters connection.APIRequestParameters) ([]IPAddress, error) {
	return s.nicIPAddressRes().List(nicID, parameters)
}

// GetNICIPAddressPaginated retrieves a paginated list of NIC IP addresses
func (s *Service) GetNICIPAddressesPaginated(nicID string, parameters connection.APIRequestParameters) (*connection.Paginated[IPAddress], error) {
	return s.nicIPAddressRes().ListPaginated(nicID, parameters)
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
