package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) dhcpRes() *resource.Resource[DHCP, string] {
	return resource.NewStringResource[DHCP](s.connection, "/ecloud/v2/dhcps", "dhcp", func(id string) error {
		return &DHCPNotFoundError{ID: id}
	})
}

// GetDHCPs retrieves a list of dhcps
func (s *Service) GetDHCPs(parameters connection.APIRequestParameters) ([]DHCP, error) {
	return s.dhcpRes().List(parameters)
}

// GetDHCPsPaginated retrieves a paginated list of dhcps
func (s *Service) GetDHCPsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[DHCP], error) {
	return s.dhcpRes().ListPaginated(parameters)
}

// GetDHCP retrieves a single dhcp by id
func (s *Service) GetDHCP(dhcpID string) (DHCP, error) {
	return s.dhcpRes().Get(dhcpID)
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
