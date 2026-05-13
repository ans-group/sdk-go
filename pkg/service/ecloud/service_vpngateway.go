package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) vpnGatewayRes() *resource.Resource[VPNGateway, string] {
	return resource.NewStringResource[VPNGateway](s.connection, "/ecloud/v2/vpn-gateways", "vpn gateway", func(id string) error {
		return &VPNGatewayNotFoundError{ID: id}
	})
}

// GetVPNGateways retrieves a list of VPN gateways
func (s *Service) GetVPNGateways(parameters connection.APIRequestParameters) ([]VPNGateway, error) {
	return s.vpnGatewayRes().List(parameters)
}

// GetVPNGatewaysPaginated retrieves a paginated list of VPN gateways
func (s *Service) GetVPNGatewaysPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VPNGateway], error) {
	return s.vpnGatewayRes().ListPaginated(parameters)
}

// GetVPNGateway retrieves a single VPN gateway by ID
func (s *Service) GetVPNGateway(gatewayID string) (VPNGateway, error) {
	return s.vpnGatewayRes().Get(gatewayID)
}

// CreateVPNGateway creates a new VPN gateway
func (s *Service) CreateVPNGateway(req CreateVPNGatewayRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/vpn-gateways", &req)
	return body.Data, err
}

// PatchVPNGateway patches a VPN gateway
func (s *Service) PatchVPNGateway(gatewayID string, req PatchVPNGatewayRequest) (TaskReference, error) {
	if gatewayID == "" {
		return TaskReference{}, fmt.Errorf("invalid gateway id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/vpn-gateways/%s", gatewayID), &req, connection.NotFoundResponseHandler(&VPNGatewayNotFoundError{ID: gatewayID}))
	return body.Data, err
}

// DeleteVPNGateway deletes a VPN gateway
func (s *Service) DeleteVPNGateway(gatewayID string) (string, error) {
	if gatewayID == "" {
		return "", fmt.Errorf("invalid gateway id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/vpn-gateways/%s", gatewayID), nil, connection.NotFoundResponseHandler(&VPNGatewayNotFoundError{ID: gatewayID}))
	return body.Data.TaskID, err
}

func (s *Service) vpnGatewayTasksRes() *resource.SubResourceList[Task, string] {
	return resource.NewStringSubResourceList[Task](s.connection,
		func(gatewayID string) string { return fmt.Sprintf("/ecloud/v2/vpn-gateways/%s/tasks", gatewayID) },
		"vpn gateway", "id", func(gatewayID string) error { return &VPNGatewayNotFoundError{ID: gatewayID} })
}

// GetVPNGatewayTasks retrieves a list of VPN gateway tasks
func (s *Service) GetVPNGatewayTasks(gatewayID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return s.vpnGatewayTasksRes().List(gatewayID, parameters)
}

// GetVPNGatewayTasksPaginated retrieves a paginated list of VPN gateway tasks
func (s *Service) GetVPNGatewayTasksPaginated(gatewayID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	return s.vpnGatewayTasksRes().ListPaginated(gatewayID, parameters)
}
