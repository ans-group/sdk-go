package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) hostRes() *resource.Resource[Host, string] {
	return resource.NewStringResource[Host](s.connection, "/ecloud/v2/hosts", "host", func(id string) error {
		return &HostNotFoundError{ID: id}
	})
}

// GetHosts retrieves a list of hosts
func (s *Service) GetHosts(parameters connection.APIRequestParameters) ([]Host, error) {
	return s.hostRes().List(parameters)
}

// GetHostsPaginated retrieves a paginated list of hosts
func (s *Service) GetHostsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Host], error) {
	return s.hostRes().ListPaginated(parameters)
}

// GetHost retrieves a single host by id
func (s *Service) GetHost(hostID string) (Host, error) {
	return s.hostRes().Get(hostID)
}

// CreateHost creates a host
func (s *Service) CreateHost(req CreateHostRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/hosts", &req)
	return body.Data, err
}

// PatchHost patches a host
func (s *Service) PatchHost(hostID string, req PatchHostRequest) (TaskReference, error) {
	if hostID == "" {
		return TaskReference{}, fmt.Errorf("invalid host id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/hosts/%s", hostID), &req, connection.NotFoundResponseHandler(&HostNotFoundError{ID: hostID}))
	return body.Data, err
}

// DeleteHost deletes a host
func (s *Service) DeleteHost(hostID string) (string, error) {
	if hostID == "" {
		return "", fmt.Errorf("invalid host id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/hosts/%s", hostID), nil, connection.NotFoundResponseHandler(&HostNotFoundError{ID: hostID}))
	return body.Data.TaskID, err
}

func (s *Service) hostTasksRes() *resource.SubResourceList[Task, string] {
	return resource.NewStringSubResourceList[Task](s.connection,
		func(hostID string) string { return fmt.Sprintf("/ecloud/v2/hosts/%s/tasks", hostID) },
		"host", "id", func(hostID string) error { return &HostNotFoundError{ID: hostID} })
}

// GetHostTasks retrieves a list of Host tasks
func (s *Service) GetHostTasks(hostID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return s.hostTasksRes().List(hostID, parameters)
}

// GetHostTasksPaginated retrieves a paginated list of Host tasks
func (s *Service) GetHostTasksPaginated(hostID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	return s.hostTasksRes().ListPaginated(hostID, parameters)
}
