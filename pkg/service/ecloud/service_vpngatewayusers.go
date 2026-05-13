package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetVPNGatewayUsers retrieves a list of VPN gateway users
func (s *Service) GetVPNGatewayUsers(parameters connection.APIRequestParameters) ([]VPNGatewayUser, error) {
	return connection.InvokeRequestAll(s.GetVPNGatewayUsersPaginated, parameters)
}

// GetVPNGatewayUsersPaginated retrieves a paginated list of VPN gateway users
func (s *Service) GetVPNGatewayUsersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VPNGatewayUser], error) {
	body, err := connection.Get[[]VPNGatewayUser](s.connection, "/ecloud/v2/vpn-gateway-users", parameters)
	return connection.NewPaginated(body, parameters, s.GetVPNGatewayUsersPaginated), err
}

// GetVPNGatewayUser retrieves a single VPN gateway user by ID
func (s *Service) GetVPNGatewayUser(userID string) (VPNGatewayUser, error) {
	if userID == "" {
		return VPNGatewayUser{}, fmt.Errorf("invalid vpn gateway user id")
	}
	body, err := connection.Get[VPNGatewayUser](s.connection, fmt.Sprintf("/ecloud/v2/vpn-gateway-users/%s", userID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&VPNGatewayUserNotFoundError{ID: userID}))
	return body.Data, err
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
