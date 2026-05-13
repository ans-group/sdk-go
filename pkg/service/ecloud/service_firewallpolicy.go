package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) firewallPolicyRes() *resource.Resource[FirewallPolicy, string] {
	return resource.NewStringResource[FirewallPolicy](s.connection, "/ecloud/v2/firewall-policies", "firewall policy", func(id string) error {
		return &FirewallPolicyNotFoundError{ID: id}
	})
}

// GetFirewallPolicies retrieves a list of firewall policies
func (s *Service) GetFirewallPolicies(parameters connection.APIRequestParameters) ([]FirewallPolicy, error) {
	return s.firewallPolicyRes().List(parameters)
}

// GetFirewallPoliciesPaginated retrieves a paginated list of firewall policies
func (s *Service) GetFirewallPoliciesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[FirewallPolicy], error) {
	return s.firewallPolicyRes().ListPaginated(parameters)
}

// GetFirewallPolicy retrieves a single firewall policy by id
func (s *Service) GetFirewallPolicy(policyID string) (FirewallPolicy, error) {
	return s.firewallPolicyRes().Get(policyID)
}

// CreateFirewallPolicy creates a new FirewallPolicy
func (s *Service) CreateFirewallPolicy(req CreateFirewallPolicyRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/firewall-policies", &req)
	return body.Data, err
}

// PatchFirewallPolicy patches a FirewallPolicy
func (s *Service) PatchFirewallPolicy(policyID string, req PatchFirewallPolicyRequest) (TaskReference, error) {
	if policyID == "" {
		return TaskReference{}, fmt.Errorf("invalid policy id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/firewall-policies/%s", policyID), &req, connection.NotFoundResponseHandler(&FirewallPolicyNotFoundError{ID: policyID}))
	return body.Data, err
}

// DeleteFirewallPolicy deletes a FirewallPolicy
func (s *Service) DeleteFirewallPolicy(policyID string) (string, error) {
	if policyID == "" {
		return "", fmt.Errorf("invalid policy id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/firewall-policies/%s", policyID), nil, connection.NotFoundResponseHandler(&FirewallPolicyNotFoundError{ID: policyID}))
	return body.Data.TaskID, err
}

func (s *Service) firewallPolicyRuleRes() *resource.SubResourceList[FirewallRule, string] {
	return resource.NewStringSubResourceList[FirewallRule](s.connection,
		func(policyID string) string {
			return fmt.Sprintf("/ecloud/v2/firewall-policies/%s/firewall-rules", policyID)
		},
		"firewall policy", "id", func(policyID string) error { return &FirewallPolicyNotFoundError{ID: policyID} })
}

// GetFirewallPolicyFirewallRules retrieves a list of firewall policy rules
func (s *Service) GetFirewallPolicyFirewallRules(policyID string, parameters connection.APIRequestParameters) ([]FirewallRule, error) {
	return s.firewallPolicyRuleRes().List(policyID, parameters)
}

// GetFirewallPolicyFirewallRulesPaginated retrieves a paginated list of firewall policy FirewallRules
func (s *Service) GetFirewallPolicyFirewallRulesPaginated(policyID string, parameters connection.APIRequestParameters) (*connection.Paginated[FirewallRule], error) {
	return s.firewallPolicyRuleRes().ListPaginated(policyID, parameters)
}

func (s *Service) firewallPolicyTasksRes() *resource.SubResourceList[Task, string] {
	return resource.NewStringSubResourceList[Task](s.connection,
		func(policyID string) string { return fmt.Sprintf("/ecloud/v2/firewall-policies/%s/tasks", policyID) },
		"firewall policy", "id", func(policyID string) error { return &FirewallPolicyNotFoundError{ID: policyID} })
}

// GetFirewallPolicyTasks retrieves a list of FirewallPolicy tasks
func (s *Service) GetFirewallPolicyTasks(policyID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return s.firewallPolicyTasksRes().List(policyID, parameters)
}

// GetFirewallPolicyTasksPaginated retrieves a paginated list of FirewallPolicy tasks
func (s *Service) GetFirewallPolicyTasksPaginated(policyID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	return s.firewallPolicyTasksRes().ListPaginated(policyID, parameters)
}
