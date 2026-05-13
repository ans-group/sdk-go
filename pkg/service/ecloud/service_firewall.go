package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetFirewalls retrieves a list of firewalls
func (s *Service) GetFirewalls(parameters connection.APIRequestParameters) ([]Firewall, error) {
	return connection.InvokeRequestAll(s.GetFirewallsPaginated, parameters)
}

// GetFirewallsPaginated retrieves a paginated list of firewalls
func (s *Service) GetFirewallsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Firewall], error) {
	body, err := connection.Get[[]Firewall](s.connection, "/ecloud/v1/firewalls", parameters)
	return connection.NewPaginated(body, parameters, s.GetFirewallsPaginated), err
}

// GetFirewall retrieves a single firewall by ID
func (s *Service) GetFirewall(firewallID int) (Firewall, error) {
	if firewallID < 1 {
		return Firewall{}, fmt.Errorf("invalid firewall id")
	}
	body, err := connection.Get[Firewall](s.connection, fmt.Sprintf("/ecloud/v1/firewalls/%d", firewallID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&FirewallNotFoundError{ID: firewallID}))
	return body.Data, err
}

// GetFirewallConfig retrieves a single firewall config by ID
func (s *Service) GetFirewallConfig(firewallID int) (FirewallConfig, error) {
	if firewallID < 1 {
		return FirewallConfig{}, fmt.Errorf("invalid firewall id")
	}
	body, err := connection.Get[FirewallConfig](s.connection, fmt.Sprintf("/ecloud/v1/firewalls/%d/config", firewallID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&FirewallNotFoundError{ID: firewallID}))
	return body.Data, err
}
