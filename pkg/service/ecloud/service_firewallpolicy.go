package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetFirewallPolicies retrieves a list of firewall policies
func (s *Service) GetFirewallPolicies(parameters connection.APIRequestParameters) ([]FirewallPolicy, error) {
	var policys []FirewallPolicy

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFirewallPoliciesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, policy := range response.(*PaginatedFirewallPolicy).Items {
			policys = append(policys, policy)
		}
	}

	return policys, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetFirewallPoliciesPaginated retrieves a paginated list of firewall policies
func (s *Service) GetFirewallPoliciesPaginated(parameters connection.APIRequestParameters) (*PaginatedFirewallPolicy, error) {
	body, err := s.getFirewallPoliciesPaginatedResponseBody(parameters)

	return NewPaginatedFirewallPolicy(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFirewallPoliciesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getFirewallPoliciesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetFirewallPolicySliceResponseBody, error) {
	body := &GetFirewallPolicySliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/firewall-policies", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetFirewallPolicy retrieves a single firewall policy by id
func (s *Service) GetFirewallPolicy(policyID string) (FirewallPolicy, error) {
	body, err := s.getFirewallPolicyResponseBody(policyID)

	return body.Data, err
}

func (s *Service) getFirewallPolicyResponseBody(policyID string) (*GetFirewallPolicyResponseBody, error) {
	body := &GetFirewallPolicyResponseBody{}

	if policyID == "" {
		return body, fmt.Errorf("invalid firewall policy id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/firewall-policies/%s", policyID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FirewallPolicyNotFoundError{ID: policyID}
		}

		return nil
	})
}

// CreateFirewallPolicy creates a new FirewallPolicy
func (s *Service) CreateFirewallPolicy(req CreateFirewallPolicyRequest) (string, error) {
	body, err := s.createFirewallPolicyResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createFirewallPolicyResponseBody(req CreateFirewallPolicyRequest) (*GetFirewallPolicyResponseBody, error) {
	body := &GetFirewallPolicyResponseBody{}

	response, err := s.connection.Post("/ecloud/v2/firewall-policies", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchFirewallPolicy patches a FirewallPolicy
func (s *Service) PatchFirewallPolicy(policyID string, req PatchFirewallPolicyRequest) error {
	_, err := s.patchFirewallPolicyResponseBody(policyID, req)

	return err
}

func (s *Service) patchFirewallPolicyResponseBody(policyID string, req PatchFirewallPolicyRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if policyID == "" {
		return body, fmt.Errorf("invalid policy id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/firewall-policies/%s", policyID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FirewallPolicyNotFoundError{ID: policyID}
		}

		return nil
	})
}

// DeleteFirewallPolicy deletes a FirewallPolicy
func (s *Service) DeleteFirewallPolicy(policyID string) error {
	_, err := s.deleteFirewallPolicyResponseBody(policyID)

	return err
}

func (s *Service) deleteFirewallPolicyResponseBody(policyID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if policyID == "" {
		return body, fmt.Errorf("invalid policy id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/firewall-policies/%s", policyID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FirewallPolicyNotFoundError{ID: policyID}
		}

		return nil
	})
}
