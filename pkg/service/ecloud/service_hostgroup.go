package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetHostGroups retrieves a list of host groups
func (s *Service) GetHostGroups(parameters connection.APIRequestParameters) ([]HostGroup, error) {
	var hostGroups []HostGroup

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetHostGroupsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, hostGroup := range response.(*PaginatedHostGroup).Items {
			hostGroups = append(hostGroups, hostGroup)
		}
	}

	return hostGroups, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetHostGroupsPaginated retrieves a paginated list of host groups
func (s *Service) GetHostGroupsPaginated(parameters connection.APIRequestParameters) (*PaginatedHostGroup, error) {
	body, err := s.getHostGroupsPaginatedResponseBody(parameters)

	return NewPaginatedHostGroup(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetHostGroupsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getHostGroupsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetHostGroupSliceResponseBody, error) {
	body := &GetHostGroupSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/host-groups", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetHostGroup retrieves a single host group by id
func (s *Service) GetHostGroup(hostGroupID string) (HostGroup, error) {
	body, err := s.getHostGroupResponseBody(hostGroupID)

	return body.Data, err
}

func (s *Service) getHostGroupResponseBody(hostGroupID string) (*GetHostGroupResponseBody, error) {
	body := &GetHostGroupResponseBody{}

	if hostGroupID == "" {
		return body, fmt.Errorf("invalid host group id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/host-groups/%s", hostGroupID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HostGroupNotFoundError{ID: hostGroupID}
		}

		return nil
	})
}

// CreateHostGroup creates a host group
func (s *Service) CreateHostGroup(req CreateHostGroupRequest) (string, error) {
	body, err := s.createHostGroupResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createHostGroupResponseBody(req CreateHostGroupRequest) (*GetHostGroupResponseBody, error) {
	body := &GetHostGroupResponseBody{}

	response, err := s.connection.Post("/ecloud/v2/host-groups", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchHostGroup patches a host group
func (s *Service) PatchHostGroup(hostGroupID string, req PatchHostGroupRequest) error {
	_, err := s.patchHostGroupResponseBody(hostGroupID, req)

	return err
}

func (s *Service) patchHostGroupResponseBody(hostGroupID string, req PatchHostGroupRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if hostGroupID == "" {
		return body, fmt.Errorf("invalid host group id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/host-groups/%s", hostGroupID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HostGroupNotFoundError{ID: hostGroupID}
		}

		return nil
	})
}

// DeleteHostGroup deletes a host group
func (s *Service) DeleteHostGroup(hostGroupID string) error {
	_, err := s.deleteHostGroupResponseBody(hostGroupID)

	return err
}

func (s *Service) deleteHostGroupResponseBody(hostGroupID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if hostGroupID == "" {
		return body, fmt.Errorf("invalid host group id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/host-groups/%s", hostGroupID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HostGroupNotFoundError{ID: hostGroupID}
		}

		return nil
	})
}
