package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetFirewallRules retrieves a list of firewall rules
func (s *Service) GetFirewallRules(parameters connection.APIRequestParameters) ([]FirewallRule, error) {
	var sites []FirewallRule

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFirewallRulesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, site := range response.(*PaginatedFirewallRule).Items {
			sites = append(sites, site)
		}
	}

	return sites, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetFirewallRulesPaginated retrieves a paginated list of firewall rules
func (s *Service) GetFirewallRulesPaginated(parameters connection.APIRequestParameters) (*PaginatedFirewallRule, error) {
	body, err := s.getFirewallRulesPaginatedResponseBody(parameters)

	return NewPaginatedFirewallRule(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFirewallRulesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getFirewallRulesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetFirewallRuleSliceResponseBody, error) {
	body := &GetFirewallRuleSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/firewall-rules", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetFirewallRule retrieves a single firewall rule by id
func (s *Service) GetFirewallRule(firewallIPID string) (FirewallRule, error) {
	body, err := s.getFirewallRuleResponseBody(firewallIPID)

	return body.Data, err
}

func (s *Service) getFirewallRuleResponseBody(firewallIPID string) (*GetFirewallRuleResponseBody, error) {
	body := &GetFirewallRuleResponseBody{}

	if firewallIPID == "" {
		return body, fmt.Errorf("invalid firewall rule id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/firewall-rules/%s", firewallIPID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FirewallRuleNotFoundError{ID: firewallIPID}
		}

		return nil
	})
}
