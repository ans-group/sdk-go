package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) firewallRuleRes() *resource.Resource[FirewallRule, string] {
	return resource.NewStringResource[FirewallRule](s.connection, "/ecloud/v2/firewall-rules", "firewall rule", func(id string) error {
		return &FirewallRuleNotFoundError{ID: id}
	})
}

// GetFirewallRules retrieves a list of firewall rules
func (s *Service) GetFirewallRules(parameters connection.APIRequestParameters) ([]FirewallRule, error) {
	return s.firewallRuleRes().List(parameters)
}

// GetFirewallRulesPaginated retrieves a paginated list of firewall rules
func (s *Service) GetFirewallRulesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[FirewallRule], error) {
	return s.firewallRuleRes().ListPaginated(parameters)
}

// GetFirewallRule retrieves a single rule by id
func (s *Service) GetFirewallRule(ruleID string) (FirewallRule, error) {
	return s.firewallRuleRes().Get(ruleID)
}

// CreateFirewallRule creates a new FirewallRule
func (s *Service) CreateFirewallRule(req CreateFirewallRuleRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/firewall-rules", &req)
	return body.Data, err
}

// PatchFirewallRule patches a FirewallRule
func (s *Service) PatchFirewallRule(ruleID string, req PatchFirewallRuleRequest) (TaskReference, error) {
	if ruleID == "" {
		return TaskReference{}, fmt.Errorf("invalid firewall rule id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/firewall-rules/%s", ruleID), &req, connection.NotFoundResponseHandler(&FirewallRuleNotFoundError{ID: ruleID}))
	return body.Data, err
}

// DeleteFirewallRule deletes a FirewallRule
func (s *Service) DeleteFirewallRule(ruleID string) (string, error) {
	if ruleID == "" {
		return "", fmt.Errorf("invalid firewall rule id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/firewall-rules/%s", ruleID), nil, connection.NotFoundResponseHandler(&FirewallRuleNotFoundError{ID: ruleID}))
	return body.Data.TaskID, err
}

// GetFirewallRuleFirewallRulePorts retrieves a list of firewall rule ports
func (s *Service) GetFirewallRuleFirewallRulePorts(firewallRuleID string, parameters connection.APIRequestParameters) ([]FirewallRulePort, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[FirewallRulePort], error) {
		return s.GetFirewallRuleFirewallRulePortsPaginated(firewallRuleID, p)
	}, parameters)
}

// GetFirewallRuleFirewallRulePortsPaginated retrieves a paginated list of firewall rule ports
func (s *Service) GetFirewallRuleFirewallRulePortsPaginated(firewallRuleID string, parameters connection.APIRequestParameters) (*connection.Paginated[FirewallRulePort], error) {
	body, err := connection.Get[[]FirewallRulePort](s.connection, fmt.Sprintf("/ecloud/v2/firewall-rules/%s/ports", firewallRuleID), parameters)
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[FirewallRulePort], error) {
		return s.GetFirewallRuleFirewallRulePortsPaginated(firewallRuleID, p)
	}), err
}
