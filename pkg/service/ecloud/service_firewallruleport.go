package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) firewallRulePortRes() *resource.Resource[FirewallRulePort, string] {
	return resource.NewStringResource[FirewallRulePort](s.connection, "/ecloud/v2/firewall-rule-ports", "firewall rule", func(id string) error {
		return &FirewallRulePortNotFoundError{ID: id}
	})
}

// GetFirewallRulePorts retrieves a list of firewall rules
func (s *Service) GetFirewallRulePorts(parameters connection.APIRequestParameters) ([]FirewallRulePort, error) {
	return s.firewallRulePortRes().List(parameters)
}

// GetFirewallRulePortsPaginated retrieves a paginated list of firewall rules
func (s *Service) GetFirewallRulePortsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[FirewallRulePort], error) {
	return s.firewallRulePortRes().ListPaginated(parameters)
}

// GetFirewallRulePort retrieves a single rule by id
func (s *Service) GetFirewallRulePort(ruleID string) (FirewallRulePort, error) {
	return s.firewallRulePortRes().Get(ruleID)
}

// CreateFirewallRulePort creates a new FirewallRulePort
func (s *Service) CreateFirewallRulePort(req CreateFirewallRulePortRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/firewall-rule-ports", &req)
	return body.Data, err
}

// PatchFirewallRulePort patches a FirewallRulePort
func (s *Service) PatchFirewallRulePort(ruleID string, req PatchFirewallRulePortRequest) (TaskReference, error) {
	if ruleID == "" {
		return TaskReference{}, fmt.Errorf("invalid firewall rule id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/firewall-rule-ports/%s", ruleID), &req, connection.NotFoundResponseHandler(&FirewallRulePortNotFoundError{ID: ruleID}))
	return body.Data, err
}

// DeleteFirewallRulePort deletes a FirewallRulePort
func (s *Service) DeleteFirewallRulePort(ruleID string) (string, error) {
	if ruleID == "" {
		return "", fmt.Errorf("invalid firewall rule id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/firewall-rule-ports/%s", ruleID), nil, connection.NotFoundResponseHandler(&FirewallRulePortNotFoundError{ID: ruleID}))
	return body.Data.TaskID, err
}
