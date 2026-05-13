package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) natOverloadRuleRes() *resource.Resource[NATOverloadRule, string] {
	return resource.NewStringResource[NATOverloadRule](s.connection, "/ecloud/v2/nat-overload-rules", "nat overload rule", func(id string) error {
		return &NATOverloadRuleNotFoundError{ID: id}
	})
}

// GetNATOverloadRules retrieves a list of NAT overload rules
func (s *Service) GetNATOverloadRules(parameters connection.APIRequestParameters) ([]NATOverloadRule, error) {
	return s.natOverloadRuleRes().List(parameters)
}

// GetNATOverloadRulesPaginated retrieves a paginated list of NAT overload rules
func (s *Service) GetNATOverloadRulesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[NATOverloadRule], error) {
	return s.natOverloadRuleRes().ListPaginated(parameters)
}

// GetNATOverloadRule retrieves a single NAT overload rule by id
func (s *Service) GetNATOverloadRule(ruleID string) (NATOverloadRule, error) {
	return s.natOverloadRuleRes().Get(ruleID)
}

// CreateNATOverloadRule creates a new NAT overload
func (s *Service) CreateNATOverloadRule(req CreateNATOverloadRuleRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/nat-overload-rules", &req)
	return body.Data, err
}

// PatchNATOverloadRule patches a NAT overload
func (s *Service) PatchNATOverloadRule(ruleID string, req PatchNATOverloadRuleRequest) (TaskReference, error) {
	if ruleID == "" {
		return TaskReference{}, fmt.Errorf("invalid nat overload rule id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/nat-overload-rules/%s", ruleID), &req, connection.NotFoundResponseHandler(&NATOverloadRuleNotFoundError{ID: ruleID}))
	return body.Data, err
}

// DeleteNATOverloadRule deletes a NAT overload
func (s *Service) DeleteNATOverloadRule(ruleID string) (string, error) {
	if ruleID == "" {
		return "", fmt.Errorf("invalid nat overload rule id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/nat-overload-rules/%s", ruleID), nil, connection.NotFoundResponseHandler(&NATOverloadRuleNotFoundError{ID: ruleID}))
	return body.Data.TaskID, err
}
