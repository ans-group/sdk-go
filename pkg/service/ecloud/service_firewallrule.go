package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetFirewallRules retrieves a list of firewall rules
func (s *Service) GetFirewallRules(parameters connection.APIRequestParameters) ([]FirewallRule, error) {
	var rules []FirewallRule

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFirewallRulesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, rule := range response.(*PaginatedFirewallRule).Items {
			rules = append(rules, rule)
		}
	}

	return rules, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
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
func (s *Service) GetFirewallRule(ruleID string) (FirewallRule, error) {
	body, err := s.getFirewallRuleResponseBody(ruleID)

	return body.Data, err
}

func (s *Service) getFirewallRuleResponseBody(ruleID string) (*GetFirewallRuleResponseBody, error) {
	body := &GetFirewallRuleResponseBody{}

	if ruleID == "" {
		return body, fmt.Errorf("invalid firewall rule id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/firewall-rules/%s", ruleID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FirewallRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}
