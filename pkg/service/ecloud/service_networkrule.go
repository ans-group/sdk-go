package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) networkRuleRes() *resource.Resource[NetworkRule, string] {
	return resource.NewStringResource[NetworkRule](s.connection, "/ecloud/v2/network-rules", "network rule", func(id string) error {
		return &NetworkRuleNotFoundError{ID: id}
	})
}

// GetNetworkRules retrieves a list of network rules
func (s *Service) GetNetworkRules(parameters connection.APIRequestParameters) ([]NetworkRule, error) {
	return s.networkRuleRes().List(parameters)
}

// GetNetworkRulesPaginated retrieves a paginated list of network rules
func (s *Service) GetNetworkRulesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[NetworkRule], error) {
	return s.networkRuleRes().ListPaginated(parameters)
}

// GetNetworkRule retrieves a single rule by id
func (s *Service) GetNetworkRule(ruleID string) (NetworkRule, error) {
	return s.networkRuleRes().Get(ruleID)
}

// CreateNetworkRule creates a new NetworkRule
func (s *Service) CreateNetworkRule(req CreateNetworkRuleRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/network-rules", &req)
	return body.Data, err
}

// PatchNetworkRule patches a NetworkRule
func (s *Service) PatchNetworkRule(ruleID string, req PatchNetworkRuleRequest) (TaskReference, error) {
	if ruleID == "" {
		return TaskReference{}, fmt.Errorf("invalid network rule id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/network-rules/%s", ruleID), &req, connection.NotFoundResponseHandler(&NetworkRuleNotFoundError{ID: ruleID}))
	return body.Data, err
}

// DeleteNetworkRule deletes a NetworkRule
func (s *Service) DeleteNetworkRule(ruleID string) (string, error) {
	if ruleID == "" {
		return "", fmt.Errorf("invalid network rule id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/network-rules/%s", ruleID), nil, connection.NotFoundResponseHandler(&NetworkRuleNotFoundError{ID: ruleID}))
	return body.Data.TaskID, err
}

func (s *Service) networkRulePortsRes() *resource.SubResourceList[NetworkRulePort, string] {
	return resource.NewUncheckedStringSubResourceList[NetworkRulePort](s.connection,
		func(networkRuleID string) string {
			return fmt.Sprintf("/ecloud/v2/network-rules/%s/ports", networkRuleID)
		})
}

// GetNetworkRuleNetworkRulePorts retrieves a list of network rule ports
func (s *Service) GetNetworkRuleNetworkRulePorts(networkRuleID string, parameters connection.APIRequestParameters) ([]NetworkRulePort, error) {
	return s.networkRulePortsRes().List(networkRuleID, parameters)
}

// GetNetworkRuleNetworkRulePortsPaginated retrieves a paginated list of network rule ports
func (s *Service) GetNetworkRuleNetworkRulePortsPaginated(networkRuleID string, parameters connection.APIRequestParameters) (*connection.Paginated[NetworkRulePort], error) {
	return s.networkRulePortsRes().ListPaginated(networkRuleID, parameters)
}
