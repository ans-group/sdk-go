package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetVIPs retrieves a list of vips
func (s *Service) GetVIPs(parameters connection.APIRequestParameters) ([]VIP, error) {
	return connection.InvokeRequestAll(s.GetVIPsPaginated, parameters)
}

// GetVIPsPaginated retrieves a paginated list of vips
func (s *Service) GetVIPsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VIP], error) {
	body, err := connection.Get[[]VIP](s.connection, "/ecloud/v2/vips", parameters)
	return connection.NewPaginated(body, parameters, s.GetVIPsPaginated), err
}

// GetVIP retrieves a single vip by id
func (s *Service) GetVIP(vipID string) (VIP, error) {
	if vipID == "" {
		return VIP{}, fmt.Errorf("invalid vip id")
	}
	body, err := connection.Get[VIP](s.connection, fmt.Sprintf("/ecloud/v2/vips/%s", vipID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&VIPNotFoundError{ID: vipID}))
	return body.Data, err
}

// CreateVIP creates a new VIP
func (s *Service) CreateVIP(req CreateVIPRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/vips", &req)
	return body.Data, err
}

// PatchVIP patches a VIP
func (s *Service) PatchVIP(vipID string, req PatchVIPRequest) (TaskReference, error) {
	if vipID == "" {
		return TaskReference{}, fmt.Errorf("invalid vip id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/vips/%s", vipID), &req, connection.NotFoundResponseHandler(&VIPNotFoundError{ID: vipID}))
	return body.Data, err
}

// DeleteVIP deletes a VIP
func (s *Service) DeleteVIP(vipID string) (string, error) {
	if vipID == "" {
		return "", fmt.Errorf("invalid vip id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/vips/%s", vipID), nil, connection.NotFoundResponseHandler(&VIPNotFoundError{ID: vipID}))
	return body.Data.TaskID, err
}
