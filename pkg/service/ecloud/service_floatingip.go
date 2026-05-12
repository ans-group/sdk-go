package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetFloatingIPs retrieves a list of floating ips
func (s *Service) GetFloatingIPs(parameters connection.APIRequestParameters) ([]FloatingIP, error) {
	return connection.InvokeRequestAll(s.GetFloatingIPsPaginated, parameters)
}

// GetFloatingIPsPaginated retrieves a paginated list of floating ips
func (s *Service) GetFloatingIPsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[FloatingIP], error) {
	body, err := connection.Get[[]FloatingIP](s.connection, "/ecloud/v2/floating-ips", parameters)
	return connection.NewPaginated(body, parameters, s.GetFloatingIPsPaginated), err
}

// GetFloatingIP retrieves a single floating ip by id
func (s *Service) GetFloatingIP(fipID string) (FloatingIP, error) {
	if fipID == "" {
		return FloatingIP{}, fmt.Errorf("invalid floating ip id")
	}
	body, err := connection.Get[FloatingIP](s.connection, fmt.Sprintf("/ecloud/v2/floating-ips/%s", fipID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&FloatingIPNotFoundError{ID: fipID}))
	return body.Data, err
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
