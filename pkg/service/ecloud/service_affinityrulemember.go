package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetAffinityRuleMembers retrieves a list of affinity rules
func (s *Service) GetAffinityRuleMembers(affinityRuleID string, parameters connection.APIRequestParameters) ([]AffinityRuleMember, error) {
	return connection.InvokeRequestAll(
		func(p connection.APIRequestParameters) (*connection.Paginated[AffinityRuleMember], error) {
			return s.GetAffinityRuleMembersPaginated(affinityRuleID, p)
		}, parameters)
}

// GetAffinityRuleMembersPaginated retrieves a paginated list of lbs
func (s *Service) GetAffinityRuleMembersPaginated(affinityRuleID string, parameters connection.APIRequestParameters) (*connection.Paginated[AffinityRuleMember], error) {
	body, err := s.getAffinityRuleMembersPaginatedResponseBody(affinityRuleID, parameters)
	return connection.NewPaginated(
		body,
		parameters,
		func(p connection.APIRequestParameters) (*connection.Paginated[AffinityRuleMember], error) {
			return s.GetAffinityRuleMembersPaginated(affinityRuleID, p)
		}), err
}

func (s *Service) getAffinityRuleMembersPaginatedResponseBody(affinityRuleID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]AffinityRuleMember], error) {
	body := &connection.APIResponseBodyData[[]AffinityRuleMember]{}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/affinity-rules/%s/members", affinityRuleID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetAffinityRuleMember retrieves a single lb by id
func (s *Service) GetAffinityRuleMember(affinityRuleID string, memberID string) (AffinityRuleMember, error) {
	body, err := s.getAffinityRuleMemberResponseBody(affinityRuleID, memberID)

	return body.Data, err
}

func (s *Service) getAffinityRuleMemberResponseBody(affinityRuleID string, memberID string) (*connection.APIResponseBodyData[AffinityRuleMember], error) {
	body := &connection.APIResponseBodyData[AffinityRuleMember]{}

	if affinityRuleID == "" {
		return body, fmt.Errorf("invalid affinity rule id")
	}

	if memberID == "" {
		return body, fmt.Errorf("invalid affinity rule member id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/affinity-rules/%s/members/%s", affinityRuleID, memberID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &AffinityRuleMemberNotFoundError{ID: memberID}
		}

		return nil
	})
}

// CreateAffinityRuleMember creates a new AffinityRuleMember
func (s *Service) CreateAffinityRuleMember(affinityRuleID string, req CreateAffinityRuleMemberRequest) (TaskReference, error) {
	body, err := s.createAffinityRuleMemberResponseBody(affinityRuleID, req)

	return body.Data, err
}

func (s *Service) createAffinityRuleMemberResponseBody(affinityRuleID string, req CreateAffinityRuleMemberRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v2/affinity-rules/%s/members", affinityRuleID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// DeleteAffinityRuleMember deletes a AffinityRuleMember
func (s *Service) DeleteAffinityRuleMember(affinityRuleID string, memberID string) (string, error) {
	body, err := s.deleteAffinityRuleMemberResponseBody(affinityRuleID, memberID)

	return body.Data.TaskID, err
}

func (s *Service) deleteAffinityRuleMemberResponseBody(affinityRuleID string, memberID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if affinityRuleID == "" {
		return body, fmt.Errorf("invalid affinity rule id")
	}

	if memberID == "" {
		return body, fmt.Errorf("invalid affinity rule member id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/affinity-rules/%s/members/%s", affinityRuleID, memberID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &AffinityRuleMemberNotFoundError{ID: memberID}
		}

		return nil
	})
}
