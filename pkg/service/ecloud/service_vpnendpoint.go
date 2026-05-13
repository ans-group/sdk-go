package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) vpnEndpointRes() *resource.Resource[VPNEndpoint, string] {
	return resource.NewStringResource[VPNEndpoint](s.connection, "/ecloud/v2/vpn-endpoints", "vpn endpoint", func(id string) error {
		return &VPNEndpointNotFoundError{ID: id}
	})
}

// GetVPNEndpoints retrieves a list of VPN endpoints
func (s *Service) GetVPNEndpoints(parameters connection.APIRequestParameters) ([]VPNEndpoint, error) {
	return s.vpnEndpointRes().List(parameters)
}

// GetVPNEndpointsPaginated retrieves a paginated list of VPN endpoints
func (s *Service) GetVPNEndpointsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VPNEndpoint], error) {
	return s.vpnEndpointRes().ListPaginated(parameters)
}

// GetVPNEndpoint retrieves a single VPN endpoint by id
func (s *Service) GetVPNEndpoint(endpointID string) (VPNEndpoint, error) {
	return s.vpnEndpointRes().Get(endpointID)
}

// CreateVPNEndpoint creates a new VPN endpoint
func (s *Service) CreateVPNEndpoint(req CreateVPNEndpointRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/vpn-endpoints", &req)
	return body.Data, err
}

// PatchVPNEndpoint patches a VPN endpoint
func (s *Service) PatchVPNEndpoint(endpointID string, req PatchVPNEndpointRequest) (TaskReference, error) {
	if endpointID == "" {
		return TaskReference{}, fmt.Errorf("invalid endpoint id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/vpn-endpoints/%s", endpointID), &req, connection.NotFoundResponseHandler(&VPNEndpointNotFoundError{ID: endpointID}))
	return body.Data, err
}

// DeleteVPNEndpoint deletes a VPN endpoint
func (s *Service) DeleteVPNEndpoint(endpointID string) (string, error) {
	if endpointID == "" {
		return "", fmt.Errorf("invalid endpoint id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/vpn-endpoints/%s", endpointID), nil, connection.NotFoundResponseHandler(&VPNEndpointNotFoundError{ID: endpointID}))
	return body.Data.TaskID, err
}

// GetVPNEndpointTasks retrieves a list of VPN endpoint tasks
func (s *Service) GetVPNEndpointTasks(endpointID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetVPNEndpointTasksPaginated(endpointID, p)
	}, parameters)
}

// GetVPNEndpointTasksPaginated retrieves a paginated list of VPN endpoint tasks
func (s *Service) GetVPNEndpointTasksPaginated(endpointID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	if endpointID == "" {
		return nil, fmt.Errorf("invalid vpn endpoint id")
	}
	body, err := connection.Get[[]Task](s.connection, fmt.Sprintf("/ecloud/v2/vpn-endpoints/%s/tasks", endpointID), parameters, connection.NotFoundResponseHandler(&VPNEndpointNotFoundError{ID: endpointID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetVPNEndpointTasksPaginated(endpointID, p)
	}), err
}
