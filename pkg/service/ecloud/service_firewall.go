package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) firewallRes() *resource.Resource[Firewall, int] {
	return resource.NewIntResource[Firewall](s.connection, "/ecloud/v1/firewalls", "firewall", func(id int) error {
		return &FirewallNotFoundError{ID: id}
	})
}

// GetFirewalls retrieves a list of firewalls
func (s *Service) GetFirewalls(parameters connection.APIRequestParameters) ([]Firewall, error) {
	return s.firewallRes().List(parameters)
}

// GetFirewallsPaginated retrieves a paginated list of firewalls
func (s *Service) GetFirewallsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Firewall], error) {
	return s.firewallRes().ListPaginated(parameters)
}

// GetFirewall retrieves a single firewall by ID
func (s *Service) GetFirewall(firewallID int) (Firewall, error) {
	return s.firewallRes().Get(firewallID)
}

// GetFirewallConfig retrieves a single firewall config by ID
func (s *Service) GetFirewallConfig(firewallID int) (FirewallConfig, error) {
	if firewallID < 1 {
		return FirewallConfig{}, fmt.Errorf("invalid firewall id")
	}
	body, err := connection.Get[FirewallConfig](s.connection, fmt.Sprintf("/ecloud/v1/firewalls/%d/config", firewallID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&FirewallNotFoundError{ID: firewallID}))
	return body.Data, err
}
