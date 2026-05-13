package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) affinityRuleRes() *resource.Resource[AffinityRule, string] {
	return resource.NewStringResource[AffinityRule](s.connection, "/ecloud/v2/affinity-rules", "affinity rule", func(id string) error {
		return &AffinityRuleNotFoundError{ID: id}
	})
}

// GetAffinityRules retrieves a list of affinity rules
func (s *Service) GetAffinityRules(parameters connection.APIRequestParameters) ([]AffinityRule, error) {
	return s.affinityRuleRes().List(parameters)
}

// GetAffinityRulesPaginated retrieves a paginated list of affinity rules
func (s *Service) GetAffinityRulesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[AffinityRule], error) {
	return s.affinityRuleRes().ListPaginated(parameters)
}

// GetAffinityRule retrieves a single AffinityRule by id
func (s *Service) GetAffinityRule(affinityruleID string) (AffinityRule, error) {
	return s.affinityRuleRes().Get(affinityruleID)
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
