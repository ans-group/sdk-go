package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) affinityRuleMemberRes() *resource.Resource[AffinityRuleMember, string] {
	return resource.NewStringResource[AffinityRuleMember](s.connection, "/ecloud/v2/affinity-rule-members", "affinity rule member", func(id string) error {
		return &AffinityRuleMemberNotFoundError{ID: id}
	})
}

func (s *Service) affinityRuleMembersRes() *resource.SubResourceList[AffinityRuleMember, string] {
	return resource.NewUncheckedStringSubResourceList[AffinityRuleMember](s.connection,
		func(affinityRuleID string) string {
			return fmt.Sprintf("/ecloud/v2/affinity-rules/%s/members", affinityRuleID)
		})
}

// GetAffinityRuleMembers retrieves a list of affinity rule members
func (s *Service) GetAffinityRuleMembers(affinityRuleID string, parameters connection.APIRequestParameters) ([]AffinityRuleMember, error) {
	return s.affinityRuleMembersRes().List(affinityRuleID, parameters)
}

// GetAffinityRuleMembersPaginated retrieves a paginated list of affinity rule members
func (s *Service) GetAffinityRuleMembersPaginated(affinityRuleID string, parameters connection.APIRequestParameters) (*connection.Paginated[AffinityRuleMember], error) {
	return s.affinityRuleMembersRes().ListPaginated(affinityRuleID, parameters)
}

// GetAffinityRuleMember retrieves a single AffinityRuleMember by id
func (s *Service) GetAffinityRuleMember(memberID string) (AffinityRuleMember, error) {
	return s.affinityRuleMemberRes().Get(memberID)
}

// CreateAffinityRuleMember creates a new AffinityRuleMember
func (s *Service) CreateAffinityRuleMember(req CreateAffinityRuleMemberRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/affinity-rule-members", &req)
	return body.Data, err
}

// DeleteAffinityRuleMember deletes a AffinityRuleMember
func (s *Service) DeleteAffinityRuleMember(memberID string) (string, error) {
	if memberID == "" {
		return "", fmt.Errorf("invalid affinity rule member id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/affinity-rule-members/%s", memberID), nil, connection.NotFoundResponseHandler(&AffinityRuleMemberNotFoundError{ID: memberID}))
	return body.Data.TaskID, err
}
