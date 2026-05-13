package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetAffinityRules retrieves a list of affinity rules
func (s *Service) GetAffinityRules(parameters connection.APIRequestParameters) ([]AffinityRule, error) {
	return connection.InvokeRequestAll(s.GetAffinityRulesPaginated, parameters)
}

// GetAffinityRulesPaginated retrieves a paginated list of affinity rules
func (s *Service) GetAffinityRulesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[AffinityRule], error) {
	body, err := connection.Get[[]AffinityRule](s.connection, "/ecloud/v2/affinity-rules", parameters)
	return connection.NewPaginated(body, parameters, s.GetAffinityRulesPaginated), err
}

// GetAffinityRule retrieves a single AffinityRule by id
func (s *Service) GetAffinityRule(affinityruleID string) (AffinityRule, error) {
	if affinityruleID == "" {
		return AffinityRule{}, fmt.Errorf("invalid affinity rule id")
	}
	body, err := connection.Get[AffinityRule](s.connection, fmt.Sprintf("/ecloud/v2/affinity-rules/%s", affinityruleID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&AffinityRuleNotFoundError{ID: affinityruleID}))
	return body.Data, err
}

// CreateAffinityRule creates a new AffinityRule
func (s *Service) CreateAffinityRule(req CreateAffinityRuleRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/affinity-rules", &req)
	return body.Data, err
}

// PatchAffinityRule patches a AffinityRule
func (s *Service) PatchAffinityRule(affinityruleID string, req PatchAffinityRuleRequest) (TaskReference, error) {
	if affinityruleID == "" {
		return TaskReference{}, fmt.Errorf("invalid affinity rule id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/affinity-rules/%s", affinityruleID), &req, connection.NotFoundResponseHandler(&AffinityRuleNotFoundError{ID: affinityruleID}))
	return body.Data, err
}

// DeleteAffinityRule deletes a AffinityRule
func (s *Service) DeleteAffinityRule(affinityruleID string) (string, error) {
	if affinityruleID == "" {
		return "", fmt.Errorf("invalid affinity rule id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/affinity-rules/%s", affinityruleID), nil, connection.NotFoundResponseHandler(&AffinityRuleNotFoundError{ID: affinityruleID}))
	return body.Data.TaskID, err
}
