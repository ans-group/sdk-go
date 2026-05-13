package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetFirewallPolicies retrieves a list of firewall policies
func (s *Service) GetFirewallPolicies(parameters connection.APIRequestParameters) ([]FirewallPolicy, error) {
	return connection.InvokeRequestAll(s.GetFirewallPoliciesPaginated, parameters)
}

// GetFirewallPoliciesPaginated retrieves a paginated list of firewall policies
func (s *Service) GetFirewallPoliciesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[FirewallPolicy], error) {
	body, err := connection.Get[[]FirewallPolicy](s.connection, "/ecloud/v2/firewall-policies", parameters)
	return connection.NewPaginated(body, parameters, s.GetFirewallPoliciesPaginated), err
}

// GetFirewallPolicy retrieves a single firewall policy by id
func (s *Service) GetFirewallPolicy(policyID string) (FirewallPolicy, error) {
	if policyID == "" {
		return FirewallPolicy{}, fmt.Errorf("invalid firewall policy id")
	}
	body, err := connection.Get[FirewallPolicy](s.connection, fmt.Sprintf("/ecloud/v2/firewall-policies/%s", policyID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&FirewallPolicyNotFoundError{ID: policyID}))
	return body.Data, err
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

// GetFirewallPolicyFirewallRules retrieves a list of firewall policy rules
func (s *Service) GetFirewallPolicyFirewallRules(policyID string, parameters connection.APIRequestParameters) ([]FirewallRule, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[FirewallRule], error) {
		return s.GetFirewallPolicyFirewallRulesPaginated(policyID, p)
	}, parameters)
}

// GetFirewallPolicyFirewallRulesPaginated retrieves a paginated list of firewall policy FirewallRules
func (s *Service) GetFirewallPolicyFirewallRulesPaginated(policyID string, parameters connection.APIRequestParameters) (*connection.Paginated[FirewallRule], error) {
	if policyID == "" {
		return nil, fmt.Errorf("invalid firewall policy id")
	}
	body, err := connection.Get[[]FirewallRule](s.connection, fmt.Sprintf("/ecloud/v2/firewall-policies/%s/firewall-rules", policyID), parameters, connection.NotFoundResponseHandler(&FirewallPolicyNotFoundError{ID: policyID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[FirewallRule], error) {
		return s.GetFirewallPolicyFirewallRulesPaginated(policyID, p)
	}), err
}

// GetFirewallPolicyTasks retrieves a list of FirewallPolicy tasks
func (s *Service) GetFirewallPolicyTasks(policyID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetFirewallPolicyTasksPaginated(policyID, p)
	}, parameters)
}

// GetFirewallPolicyTasksPaginated retrieves a paginated list of FirewallPolicy tasks
func (s *Service) GetFirewallPolicyTasksPaginated(policyID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	if policyID == "" {
		return nil, fmt.Errorf("invalid firewall policy id")
	}
	body, err := connection.Get[[]Task](s.connection, fmt.Sprintf("/ecloud/v2/firewall-policies/%s/tasks", policyID), parameters, connection.NotFoundResponseHandler(&FirewallPolicyNotFoundError{ID: policyID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetFirewallPolicyTasksPaginated(policyID, p)
	}), err
}
