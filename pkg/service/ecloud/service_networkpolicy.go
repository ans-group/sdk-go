package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) networkPolicyRes() *resource.Resource[NetworkPolicy, string] {
	return resource.NewStringResource[NetworkPolicy](s.connection, "/ecloud/v2/network-policies", "network policy", func(id string) error {
		return &NetworkPolicyNotFoundError{ID: id}
	})
}

// GetNetworkPolicies retrieves a list of network policies
func (s *Service) GetNetworkPolicies(parameters connection.APIRequestParameters) ([]NetworkPolicy, error) {
	return s.networkPolicyRes().List(parameters)
}

// GetNetworkPoliciesPaginated retrieves a paginated list of network policies
func (s *Service) GetNetworkPoliciesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[NetworkPolicy], error) {
	return s.networkPolicyRes().ListPaginated(parameters)
}

// GetNetworkPolicy retrieves a single network policy by id
func (s *Service) GetNetworkPolicy(policyID string) (NetworkPolicy, error) {
	return s.networkPolicyRes().Get(policyID)
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

func (s *Service) networkPolicyRuleRes() *resource.SubResourceList[NetworkRule, string] {
	return resource.NewStringSubResourceList[NetworkRule](s.connection,
		func(policyID string) string {
			return fmt.Sprintf("/ecloud/v2/network-policies/%s/network-rules", policyID)
		},
		"network policy", "id", func(policyID string) error { return &NetworkPolicyNotFoundError{ID: policyID} })
}

// GetNetworkPolicyNetworkRules retrieves a list of network policy rules
func (s *Service) GetNetworkPolicyNetworkRules(policyID string, parameters connection.APIRequestParameters) ([]NetworkRule, error) {
	return s.networkPolicyRuleRes().List(policyID, parameters)
}

// GetNetworkPolicyNetworkRulesPaginated retrieves a paginated list of network policy NetworkRules
func (s *Service) GetNetworkPolicyNetworkRulesPaginated(policyID string, parameters connection.APIRequestParameters) (*connection.Paginated[NetworkRule], error) {
	return s.networkPolicyRuleRes().ListPaginated(policyID, parameters)
}

func (s *Service) networkPolicyTasksRes() *resource.SubResourceList[Task, string] {
	return resource.NewStringSubResourceList[Task](s.connection,
		func(policyID string) string { return fmt.Sprintf("/ecloud/v2/network-policies/%s/tasks", policyID) },
		"network policy", "id", func(policyID string) error { return &NetworkPolicyNotFoundError{ID: policyID} })
}

// GetNetworkPolicyTasks retrieves a list of NetworkPolicy tasks
func (s *Service) GetNetworkPolicyTasks(policyID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return s.networkPolicyTasksRes().List(policyID, parameters)
}

// GetNetworkPolicyTasksPaginated retrieves a paginated list of NetworkPolicy tasks
func (s *Service) GetNetworkPolicyTasksPaginated(policyID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	return s.networkPolicyTasksRes().ListPaginated(policyID, parameters)
}
