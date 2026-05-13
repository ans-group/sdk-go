package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetVPNServices retrieves a list of VPN services
func (s *Service) GetVPNServices(parameters connection.APIRequestParameters) ([]VPNService, error) {
	return connection.InvokeRequestAll(s.GetVPNServicesPaginated, parameters)
}

// GetVPNServicesPaginated retrieves a paginated list of VPN services
func (s *Service) GetVPNServicesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VPNService], error) {
	body, err := connection.Get[[]VPNService](s.connection, "/ecloud/v2/vpn-services", parameters)
	return connection.NewPaginated(body, parameters, s.GetVPNServicesPaginated), err
}

// GetVPNService retrieves a single VPN service by id
func (s *Service) GetVPNService(serviceID string) (VPNService, error) {
	if serviceID == "" {
		return VPNService{}, fmt.Errorf("invalid vpn service id")
	}
	body, err := connection.Get[VPNService](s.connection, fmt.Sprintf("/ecloud/v2/vpn-services/%s", serviceID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&VPNServiceNotFoundError{ID: serviceID}))
	return body.Data, err
}

// CreateVPNService creates a new VPN service
func (s *Service) CreateVPNService(req CreateVPNServiceRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/vpn-services", &req)
	return body.Data, err
}

// PatchVPNService patches a VPN service
func (s *Service) PatchVPNService(serviceID string, req PatchVPNServiceRequest) (TaskReference, error) {
	if serviceID == "" {
		return TaskReference{}, fmt.Errorf("invalid service id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/vpn-services/%s", serviceID), &req, connection.NotFoundResponseHandler(&VPNServiceNotFoundError{ID: serviceID}))
	return body.Data, err
}

// DeleteVPNService deletes a VPN service
func (s *Service) DeleteVPNService(serviceID string) (string, error) {
	if serviceID == "" {
		return "", fmt.Errorf("invalid service id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/vpn-services/%s", serviceID), nil, connection.NotFoundResponseHandler(&VPNServiceNotFoundError{ID: serviceID}))
	return body.Data.TaskID, err
}

// GetVPNServiceTasks retrieves a list of VPN service tasks
func (s *Service) GetVPNServiceTasks(serviceID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetVPNServiceTasksPaginated(serviceID, p)
	}, parameters)
}

// GetVPNServiceTasksPaginated retrieves a paginated list of VPN service tasks
func (s *Service) GetVPNServiceTasksPaginated(serviceID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	if serviceID == "" {
		return nil, fmt.Errorf("invalid vpn service id")
	}
	body, err := connection.Get[[]Task](s.connection, fmt.Sprintf("/ecloud/v2/vpn-services/%s/tasks", serviceID), parameters, connection.NotFoundResponseHandler(&VPNServiceNotFoundError{ID: serviceID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetVPNServiceTasksPaginated(serviceID, p)
	}), err
}
