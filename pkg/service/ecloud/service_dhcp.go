package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetDHCPs retrieves a list of dhcps
func (s *Service) GetDHCPs(parameters connection.APIRequestParameters) ([]DHCP, error) {
	return connection.InvokeRequestAll(s.GetDHCPsPaginated, parameters)
}

// GetDHCPsPaginated retrieves a paginated list of dhcps
func (s *Service) GetDHCPsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[DHCP], error) {
	body, err := connection.Get[[]DHCP](s.connection, "/ecloud/v2/dhcps", parameters)
	return connection.NewPaginated(body, parameters, s.GetDHCPsPaginated), err
}

// GetDHCP retrieves a single dhcp by id
func (s *Service) GetDHCP(dhcpID string) (DHCP, error) {
	if dhcpID == "" {
		return DHCP{}, fmt.Errorf("invalid dhcp id")
	}
	body, err := connection.Get[DHCP](s.connection, fmt.Sprintf("/ecloud/v2/dhcps/%s", dhcpID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&DHCPNotFoundError{ID: dhcpID}))
	return body.Data, err
}

// GetDHCPTasks retrieves a list of DHCP tasks
func (s *Service) GetDHCPTasks(dhcpID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetDHCPTasksPaginated(dhcpID, p)
	}, parameters)
}

// GetDHCPTasksPaginated retrieves a paginated list of DHCP tasks
func (s *Service) GetDHCPTasksPaginated(dhcpID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	if dhcpID == "" {
		return nil, fmt.Errorf("invalid dhcp id")
	}
	body, err := connection.Get[[]Task](s.connection, fmt.Sprintf("/ecloud/v2/dhcps/%s/tasks", dhcpID), parameters, connection.NotFoundResponseHandler(&DHCPNotFoundError{ID: dhcpID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetDHCPTasksPaginated(dhcpID, p)
	}), err
}
