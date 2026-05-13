package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) floatingIPRes() *resource.Resource[FloatingIP, string] {
	return resource.NewStringResource[FloatingIP](s.connection, "/ecloud/v2/floating-ips", "floating ip", func(id string) error {
		return &FloatingIPNotFoundError{ID: id}
	})
}

// GetFloatingIPs retrieves a list of floating ips
func (s *Service) GetFloatingIPs(parameters connection.APIRequestParameters) ([]FloatingIP, error) {
	return s.floatingIPRes().List(parameters)
}

// GetFloatingIPsPaginated retrieves a paginated list of floating ips
func (s *Service) GetFloatingIPsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[FloatingIP], error) {
	return s.floatingIPRes().ListPaginated(parameters)
}

// GetFloatingIP retrieves a single floating ip by id
func (s *Service) GetFloatingIP(fipID string) (FloatingIP, error) {
	return s.floatingIPRes().Get(fipID)
}

// CreateFloatingIP creates a new FloatingIP
func (s *Service) CreateFloatingIP(req CreateFloatingIPRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/floating-ips", &req)
	return body.Data, err
}

// PatchFloatingIP patches a floating IP
func (s *Service) PatchFloatingIP(fipID string, req PatchFloatingIPRequest) (TaskReference, error) {
	if fipID == "" {
		return TaskReference{}, fmt.Errorf("invalid floating IP id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/floating-ips/%s", fipID), &req, connection.NotFoundResponseHandler(&FloatingIPNotFoundError{ID: fipID}))
	return body.Data, err
}

// DeleteFloatingIP deletes a floating IP
func (s *Service) DeleteFloatingIP(fipID string) (string, error) {
	if fipID == "" {
		return "", fmt.Errorf("invalid floating IP id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/floating-ips/%s", fipID), nil, connection.NotFoundResponseHandler(&FloatingIPNotFoundError{ID: fipID}))
	return body.Data.TaskID, err
}

// AssignFloatingIP assigns a floating IP to a resource
func (s *Service) AssignFloatingIP(fipID string, req AssignFloatingIPRequest) (string, error) {
	if fipID == "" {
		return "", fmt.Errorf("invalid floating IP id")
	}
	body, err := connection.Post[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/floating-ips/%s/assign", fipID), &req, connection.NotFoundResponseHandler(&FloatingIPNotFoundError{ID: fipID}))
	return body.Data.TaskID, err
}

// UnassignFloatingIP unassigns a floating IP from a resource
func (s *Service) UnassignFloatingIP(fipID string) (string, error) {
	if fipID == "" {
		return "", fmt.Errorf("invalid floating IP id")
	}
	body, err := connection.Post[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/floating-ips/%s/unassign", fipID), nil, connection.NotFoundResponseHandler(&FloatingIPNotFoundError{ID: fipID}))
	return body.Data.TaskID, err
}

// GetFloatingIPTasks retrieves a list of FloatingIP tasks
func (s *Service) GetFloatingIPTasks(fipID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetFloatingIPTasksPaginated(fipID, p)
	}, parameters)
}

// GetFloatingIPTasksPaginated retrieves a paginated list of FloatingIP tasks
func (s *Service) GetFloatingIPTasksPaginated(fipID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	if fipID == "" {
		return nil, fmt.Errorf("invalid floating ip id")
	}
	body, err := connection.Get[[]Task](s.connection, fmt.Sprintf("/ecloud/v2/floating-ips/%s/tasks", fipID), parameters, connection.NotFoundResponseHandler(&FloatingIPNotFoundError{ID: fipID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetFloatingIPTasksPaginated(fipID, p)
	}), err
}
