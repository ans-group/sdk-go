package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetVPNProfileGroups retrieves a list of VPN sessions
func (s *Service) GetVPNProfileGroups(parameters connection.APIRequestParameters) ([]VPNProfileGroup, error) {
	var sessions []VPNProfileGroup

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVPNProfileGroupsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, session := range response.(*PaginatedVPNProfileGroup).Items {
			sessions = append(sessions, session)
		}
	}

	return sessions, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetVPNProfileGroupsPaginated retrieves a paginated list of VPN sessions
func (s *Service) GetVPNProfileGroupsPaginated(parameters connection.APIRequestParameters) (*PaginatedVPNProfileGroup, error) {
	body, err := s.getVPNProfileGroupsPaginatedResponseBody(parameters)

	return NewPaginatedVPNProfileGroup(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVPNProfileGroupsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getVPNProfileGroupsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetVPNProfileGroupSliceResponseBody, error) {
	body := &GetVPNProfileGroupSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/vpn-profile-groups", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetVPNProfileGroup retrieves a single VPN session by id
func (s *Service) GetVPNProfileGroup(sessionID string) (VPNProfileGroup, error) {
	body, err := s.getVPNProfileGroupResponseBody(sessionID)

	return body.Data, err
}

func (s *Service) getVPNProfileGroupResponseBody(sessionID string) (*GetVPNProfileGroupResponseBody, error) {
	body := &GetVPNProfileGroupResponseBody{}

	if sessionID == "" {
		return body, fmt.Errorf("invalid vpn session id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/vpn-profile-groups/%s", sessionID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNProfileGroupNotFoundError{ID: sessionID}
		}

		return nil
	})
}
