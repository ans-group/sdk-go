package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetNetworkPolicies retrieves a list of network policies
func (s *Service) GetNetworkPolicies(parameters connection.APIRequestParameters) ([]NetworkPolicy, error) {
	return connection.InvokeRequestAll(s.GetNetworkPoliciesPaginated, parameters)
}

// GetNetworkPoliciesPaginated retrieves a paginated list of network policies
func (s *Service) GetNetworkPoliciesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[NetworkPolicy], error) {
	body, err := connection.Get[[]NetworkPolicy](s.connection, "/ecloud/v2/network-policies", parameters)
	return connection.NewPaginated(body, parameters, s.GetNetworkPoliciesPaginated), err
}

// GetNetworkPolicy retrieves a single network policy by id
func (s *Service) GetNetworkPolicy(policyID string) (NetworkPolicy, error) {
	if policyID == "" {
		return NetworkPolicy{}, fmt.Errorf("invalid network policy id")
	}
	body, err := connection.Get[NetworkPolicy](s.connection, fmt.Sprintf("/ecloud/v2/network-policies/%s", policyID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&NetworkPolicyNotFoundError{ID: policyID}))
	return body.Data, err
}

// CreateNetworkPolicy creates a new NetworkPolicy
func (s *Service) CreateNetworkPolicy(req CreateNetworkPolicyRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/network-policies", &req)
	return body.Data, err
}

// PatchNetworkPolicy patches a NetworkPolicy
func (s *Service) PatchNetworkPolicy(policyID string, req PatchNetworkPolicyRequest) (TaskReference, error) {
	if policyID == "" {
		return TaskReference{}, fmt.Errorf("invalid policy id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/network-policies/%s", policyID), &req, connection.NotFoundResponseHandler(&NetworkPolicyNotFoundError{ID: policyID}))
	return body.Data, err
}

// DeleteNetworkPolicy deletes a NetworkPolicy
func (s *Service) DeleteNetworkPolicy(policyID string) (string, error) {
	if policyID == "" {
		return "", fmt.Errorf("invalid policy id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/network-policies/%s", policyID), nil, connection.NotFoundResponseHandler(&NetworkPolicyNotFoundError{ID: policyID}))
	return body.Data.TaskID, err
}

// GetNetworkPolicyNetworkRules retrieves a list of network policy rules
func (s *Service) GetNetworkPolicyNetworkRules(policyID string, parameters connection.APIRequestParameters) ([]NetworkRule, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[NetworkRule], error) {
		return s.GetNetworkPolicyNetworkRulesPaginated(policyID, p)
	}, parameters)
}

// GetNetworkPolicyNetworkRulesPaginated retrieves a paginated list of network policy NetworkRules
func (s *Service) GetNetworkPolicyNetworkRulesPaginated(policyID string, parameters connection.APIRequestParameters) (*connection.Paginated[NetworkRule], error) {
	if policyID == "" {
		return nil, fmt.Errorf("invalid network policy id")
	}
	body, err := connection.Get[[]NetworkRule](s.connection, fmt.Sprintf("/ecloud/v2/network-policies/%s/network-rules", policyID), parameters, connection.NotFoundResponseHandler(&NetworkPolicyNotFoundError{ID: policyID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[NetworkRule], error) {
		return s.GetNetworkPolicyNetworkRulesPaginated(policyID, p)
	}), err
}

// GetNetworkPolicyTasks retrieves a list of NetworkPolicy tasks
func (s *Service) GetNetworkPolicyTasks(policyID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetNetworkPolicyTasksPaginated(policyID, p)
	}, parameters)
}

// GetNetworkPolicyTasksPaginated retrieves a paginated list of NetworkPolicy tasks
func (s *Service) GetNetworkPolicyTasksPaginated(policyID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	if policyID == "" {
		return nil, fmt.Errorf("invalid network policy id")
	}
	body, err := connection.Get[[]Task](s.connection, fmt.Sprintf("/ecloud/v2/network-policies/%s/tasks", policyID), parameters, connection.NotFoundResponseHandler(&NetworkPolicyNotFoundError{ID: policyID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetNetworkPolicyTasksPaginated(policyID, p)
	}), err
}
