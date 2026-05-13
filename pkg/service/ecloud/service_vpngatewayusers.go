package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) vpnGatewayUserRes() *resource.Resource[VPNGatewayUser, string] {
	return resource.NewStringResource[VPNGatewayUser](s.connection, "/ecloud/v2/vpn-gateway-users", "vpn gateway user", func(id string) error {
		return &VPNGatewayUserNotFoundError{ID: id}
	})
}

// GetVPNGatewayUsers retrieves a list of VPN gateway users
func (s *Service) GetVPNGatewayUsers(parameters connection.APIRequestParameters) ([]VPNGatewayUser, error) {
	return s.vpnGatewayUserRes().List(parameters)
}

// GetVPNGatewayUsersPaginated retrieves a paginated list of VPN gateway users
func (s *Service) GetVPNGatewayUsersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VPNGatewayUser], error) {
	return s.vpnGatewayUserRes().ListPaginated(parameters)
}

// GetVPNGatewayUser retrieves a single VPN gateway user by ID
func (s *Service) GetVPNGatewayUser(userID string) (VPNGatewayUser, error) {
	return s.vpnGatewayUserRes().Get(userID)
}

// CreateVPNGatewayUser creates a new VPN gateway user
func (s *Service) CreateVPNGatewayUser(req CreateVPNGatewayUserRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/vpn-gateway-users", &req)
	return body.Data, err
}

// PatchVPNGatewayUser patches a VPN gateway user
func (s *Service) PatchVPNGatewayUser(userID string, req PatchVPNGatewayUserRequest) (TaskReference, error) {
	if userID == "" {
		return TaskReference{}, fmt.Errorf("invalid user id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/vpn-gateway-users/%s", userID), &req, connection.NotFoundResponseHandler(&VPNGatewayUserNotFoundError{ID: userID}))
	return body.Data, err
}

// DeleteVPNGatewayUser deletes a VPN gateway user
func (s *Service) DeleteVPNGatewayUser(userID string) (string, error) {
	if userID == "" {
		return "", fmt.Errorf("invalid user id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/vpn-gateway-users/%s", userID), nil, connection.NotFoundResponseHandler(&VPNGatewayUserNotFoundError{ID: userID}))
	return body.Data.TaskID, err
}
