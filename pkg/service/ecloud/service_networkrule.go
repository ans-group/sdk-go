package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetNetworkRules retrieves a list of network rules
func (s *Service) GetNetworkRules(parameters connection.APIRequestParameters) ([]NetworkRule, error) {
	return connection.InvokeRequestAll(s.GetNetworkRulesPaginated, parameters)
}

// GetNetworkRulesPaginated retrieves a paginated list of network rules
func (s *Service) GetNetworkRulesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[NetworkRule], error) {
	body, err := connection.Get[[]NetworkRule](s.connection, "/ecloud/v2/network-rules", parameters)
	return connection.NewPaginated(body, parameters, s.GetNetworkRulesPaginated), err
}

// GetNetworkRule retrieves a single rule by id
func (s *Service) GetNetworkRule(ruleID string) (NetworkRule, error) {
	if ruleID == "" {
		return NetworkRule{}, fmt.Errorf("invalid network rule id")
	}
	body, err := connection.Get[NetworkRule](s.connection, fmt.Sprintf("/ecloud/v2/network-rules/%s", ruleID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&NetworkRuleNotFoundError{ID: ruleID}))
	return body.Data, err
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

// GetNetworkRuleNetworkRulePorts retrieves a list of network rule ports
func (s *Service) GetNetworkRuleNetworkRulePorts(networkRuleID string, parameters connection.APIRequestParameters) ([]NetworkRulePort, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[NetworkRulePort], error) {
		return s.GetNetworkRuleNetworkRulePortsPaginated(networkRuleID, p)
	}, parameters)
}

// GetNetworkRuleNetworkRulePortsPaginated retrieves a paginated list of network rule ports
func (s *Service) GetNetworkRuleNetworkRulePortsPaginated(networkRuleID string, parameters connection.APIRequestParameters) (*connection.Paginated[NetworkRulePort], error) {
	body, err := connection.Get[[]NetworkRulePort](s.connection, fmt.Sprintf("/ecloud/v2/network-rules/%s/ports", networkRuleID), parameters)
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[NetworkRulePort], error) {
		return s.GetNetworkRuleNetworkRulePortsPaginated(networkRuleID, p)
	}), err
}
