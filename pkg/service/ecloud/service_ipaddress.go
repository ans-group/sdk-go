package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetIPAddresses retrieves a list of ips
func (s *Service) GetIPAddresses(parameters connection.APIRequestParameters) ([]IPAddress, error) {
	return connection.InvokeRequestAll(s.GetIPAddressesPaginated, parameters)
}

// GetIPAddressesPaginated retrieves a paginated list of ips
func (s *Service) GetIPAddressesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[IPAddress], error) {
	body, err := connection.Get[[]IPAddress](s.connection, "/ecloud/v2/ip-addresses", parameters)
	return connection.NewPaginated(body, parameters, s.GetIPAddressesPaginated), err
}

// GetIPAddress retrieves a single ip by id
func (s *Service) GetIPAddress(ipID string) (IPAddress, error) {
	if ipID == "" {
		return IPAddress{}, fmt.Errorf("invalid ip address id")
	}
	body, err := connection.Get[IPAddress](s.connection, fmt.Sprintf("/ecloud/v2/ip-addresses/%s", ipID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&IPAddressNotFoundError{ID: ipID}))
	return body.Data, err
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
