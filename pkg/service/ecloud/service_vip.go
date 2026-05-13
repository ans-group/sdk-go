package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) vipRes() *resource.Resource[VIP, string] {
	return resource.NewStringResource[VIP](s.connection, "/ecloud/v2/vips", "vip", func(id string) error {
		return &VIPNotFoundError{ID: id}
	})
}

// GetVIPs retrieves a list of vips
func (s *Service) GetVIPs(parameters connection.APIRequestParameters) ([]VIP, error) {
	return s.vipRes().List(parameters)
}

// GetVIPsPaginated retrieves a paginated list of vips
func (s *Service) GetVIPsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VIP], error) {
	return s.vipRes().ListPaginated(parameters)
}

// GetVIP retrieves a single vip by id
func (s *Service) GetVIP(vipID string) (VIP, error) {
	return s.vipRes().Get(vipID)
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
