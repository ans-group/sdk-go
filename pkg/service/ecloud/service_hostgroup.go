package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetHostGroups retrieves a list of host groups
func (s *Service) GetHostGroups(parameters connection.APIRequestParameters) ([]HostGroup, error) {
	return connection.InvokeRequestAll(s.GetHostGroupsPaginated, parameters)
}

// GetHostGroupsPaginated retrieves a paginated list of host groups
func (s *Service) GetHostGroupsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[HostGroup], error) {
	body, err := connection.Get[[]HostGroup](s.connection, "/ecloud/v2/host-groups", parameters)
	return connection.NewPaginated(body, parameters, s.GetHostGroupsPaginated), err
}

// GetHostGroup retrieves a single host group by id
func (s *Service) GetHostGroup(hostGroupID string) (HostGroup, error) {
	if hostGroupID == "" {
		return HostGroup{}, fmt.Errorf("invalid host group id")
	}
	body, err := connection.Get[HostGroup](s.connection, fmt.Sprintf("/ecloud/v2/host-groups/%s", hostGroupID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&HostGroupNotFoundError{ID: hostGroupID}))
	return body.Data, err
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

// GetHostGroupTasks retrieves a list of HostGroup tasks
func (s *Service) GetHostGroupTasks(hostGroupID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetHostGroupTasksPaginated(hostGroupID, p)
	}, parameters)
}

// GetHostGroupTasksPaginated retrieves a paginated list of HostGroup tasks
func (s *Service) GetHostGroupTasksPaginated(hostGroupID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	if hostGroupID == "" {
		return nil, fmt.Errorf("invalid host group id")
	}
	body, err := connection.Get[[]Task](s.connection, fmt.Sprintf("/ecloud/v2/host-groups/%s/tasks", hostGroupID), parameters, connection.NotFoundResponseHandler(&HostGroupNotFoundError{ID: hostGroupID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetHostGroupTasksPaginated(hostGroupID, p)
	}), err
}
