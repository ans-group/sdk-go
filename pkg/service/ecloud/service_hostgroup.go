package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) hostGroupRes() *resource.Resource[HostGroup, string] {
	return resource.NewStringResource[HostGroup](s.connection, "/ecloud/v2/host-groups", "host group", func(id string) error {
		return &HostGroupNotFoundError{ID: id}
	})
}

// GetHostGroups retrieves a list of host groups
func (s *Service) GetHostGroups(parameters connection.APIRequestParameters) ([]HostGroup, error) {
	return s.hostGroupRes().List(parameters)
}

// GetHostGroupsPaginated retrieves a paginated list of host groups
func (s *Service) GetHostGroupsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[HostGroup], error) {
	return s.hostGroupRes().ListPaginated(parameters)
}

// GetHostGroup retrieves a single host group by id
func (s *Service) GetHostGroup(hostGroupID string) (HostGroup, error) {
	return s.hostGroupRes().Get(hostGroupID)
}

// CreateHostGroup creates a host group
func (s *Service) CreateHostGroup(req CreateHostGroupRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/host-groups", &req)
	return body.Data, err
}

// PatchHostGroup patches a host group
func (s *Service) PatchHostGroup(hostGroupID string, req PatchHostGroupRequest) (TaskReference, error) {
	if hostGroupID == "" {
		return TaskReference{}, fmt.Errorf("invalid host group id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/host-groups/%s", hostGroupID), &req, connection.NotFoundResponseHandler(&HostGroupNotFoundError{ID: hostGroupID}))
	return body.Data, err
}

// DeleteHostGroup deletes a host group
func (s *Service) DeleteHostGroup(hostGroupID string) (string, error) {
	if hostGroupID == "" {
		return "", fmt.Errorf("invalid host group id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/host-groups/%s", hostGroupID), nil, connection.NotFoundResponseHandler(&HostGroupNotFoundError{ID: hostGroupID}))
	return body.Data.TaskID, err
}

func (s *Service) hostGroupTasksRes() *resource.SubResourceList[Task, string] {
	return resource.NewStringSubResourceList[Task](s.connection,
		func(hostGroupID string) string { return fmt.Sprintf("/ecloud/v2/host-groups/%s/tasks", hostGroupID) },
		"host group", "id", func(hostGroupID string) error { return &HostGroupNotFoundError{ID: hostGroupID} })
}

// GetHostGroupTasks retrieves a list of HostGroup tasks
func (s *Service) GetHostGroupTasks(hostGroupID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return s.hostGroupTasksRes().List(hostGroupID, parameters)
}

// GetHostGroupTasksPaginated retrieves a paginated list of HostGroup tasks
func (s *Service) GetHostGroupTasksPaginated(hostGroupID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	return s.hostGroupTasksRes().ListPaginated(hostGroupID, parameters)
}
