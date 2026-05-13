package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) ipAddressRes() *resource.Resource[IPAddress, string] {
	return resource.NewStringResource[IPAddress](s.connection, "/ecloud/v2/ip-addresses", "ip address", func(id string) error {
		return &IPAddressNotFoundError{ID: id}
	})
}

// GetIPAddresses retrieves a list of ips
func (s *Service) GetIPAddresses(parameters connection.APIRequestParameters) ([]IPAddress, error) {
	return s.ipAddressRes().List(parameters)
}

// GetIPAddressesPaginated retrieves a paginated list of ips
func (s *Service) GetIPAddressesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[IPAddress], error) {
	return s.ipAddressRes().ListPaginated(parameters)
}

// GetIPAddress retrieves a single ip by id
func (s *Service) GetIPAddress(ipID string) (IPAddress, error) {
	return s.ipAddressRes().Get(ipID)
}

// CreateIPAddress creates a new IPAddress
func (s *Service) CreateIPAddress(req CreateIPAddressRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/ip-addresses", &req)
	return body.Data, err
}

// PatchIPAddress patches a IPAddress
func (s *Service) PatchIPAddress(ipID string, req PatchIPAddressRequest) (TaskReference, error) {
	if ipID == "" {
		return TaskReference{}, fmt.Errorf("invalid ip address id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/ip-addresses/%s", ipID), &req, connection.NotFoundResponseHandler(&IPAddressNotFoundError{ID: ipID}))
	return body.Data, err
}

// DeleteIPAddress deletes a IPAddress
func (s *Service) DeleteIPAddress(ipID string) (string, error) {
	if ipID == "" {
		return "", fmt.Errorf("invalid ip address id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/ip-addresses/%s", ipID), nil, connection.NotFoundResponseHandler(&IPAddressNotFoundError{ID: ipID}))
	return body.Data.TaskID, err
}
