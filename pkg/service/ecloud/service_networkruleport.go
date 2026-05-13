package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) networkRulePortRes() *resource.Resource[NetworkRulePort, string] {
	return resource.NewStringResource[NetworkRulePort](s.connection, "/ecloud/v2/network-rule-ports", "network rule", func(id string) error {
		return &NetworkRulePortNotFoundError{ID: id}
	})
}

// GetNetworkRulePorts retrieves a list of network rules
func (s *Service) GetNetworkRulePorts(parameters connection.APIRequestParameters) ([]NetworkRulePort, error) {
	return s.networkRulePortRes().List(parameters)
}

// GetNetworkRulePortsPaginated retrieves a paginated list of network rules
func (s *Service) GetNetworkRulePortsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[NetworkRulePort], error) {
	return s.networkRulePortRes().ListPaginated(parameters)
}

// GetNetworkRulePort retrieves a single rule by id
func (s *Service) GetNetworkRulePort(ruleID string) (NetworkRulePort, error) {
	return s.networkRulePortRes().Get(ruleID)
}

// CreateNetworkRulePort creates a new NetworkRulePort
func (s *Service) CreateNetworkRulePort(req CreateNetworkRulePortRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/network-rule-ports", &req)
	return body.Data, err
}

// PatchNetworkRulePort patches a NetworkRulePort
func (s *Service) PatchNetworkRulePort(ruleID string, req PatchNetworkRulePortRequest) (TaskReference, error) {
	if ruleID == "" {
		return TaskReference{}, fmt.Errorf("invalid network rule id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/network-rule-ports/%s", ruleID), &req, connection.NotFoundResponseHandler(&NetworkRulePortNotFoundError{ID: ruleID}))
	return body.Data, err
}

// DeleteNetworkRulePort deletes a NetworkRulePort
func (s *Service) DeleteNetworkRulePort(ruleID string) (string, error) {
	if ruleID == "" {
		return "", fmt.Errorf("invalid network rule id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/network-rule-ports/%s", ruleID), nil, connection.NotFoundResponseHandler(&NetworkRulePortNotFoundError{ID: ruleID}))
	return body.Data.TaskID, err
}
