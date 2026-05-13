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

func (s *Service) dhcpTasksRes() *resource.SubResourceList[Task, string] {
	return resource.NewStringSubResourceList[Task](s.connection,
		func(dhcpID string) string { return fmt.Sprintf("/ecloud/v2/dhcps/%s/tasks", dhcpID) },
		"dhcp", "id", func(dhcpID string) error { return &DHCPNotFoundError{ID: dhcpID} })
}

// GetDHCPTasks retrieves a list of DHCP tasks
func (s *Service) GetDHCPTasks(dhcpID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return s.dhcpTasksRes().List(dhcpID, parameters)
}

// GetDHCPTasksPaginated retrieves a paginated list of DHCP tasks
func (s *Service) GetDHCPTasksPaginated(dhcpID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	return s.dhcpTasksRes().ListPaginated(dhcpID, parameters)
}
