package loadbalancer

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetGroups retrieves a list of groups
func (s *Service) GetGroups(parameters connection.APIRequestParameters) ([]Group, error) {
	var sites []Group

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetGroupsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, site := range response.(*PaginatedGroup).Items {
			sites = append(sites, site)
		}
	}

	return sites, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetGroupsPaginated retrieves a paginated list of groups
func (s *Service) GetGroupsPaginated(parameters connection.APIRequestParameters) (*PaginatedGroup, error) {
	body, err := s.getGroupsPaginatedResponseBody(parameters)

	return NewPaginatedGroup(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetGroupsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getGroupsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetGroupSliceResponseBody, error) {
	body := &GetGroupSliceResponseBody{}

	response, err := s.connection.Get("/loadbalancers/v2/groups", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetGroup retrieves a single group by id
func (s *Service) GetGroup(groupID string) (Group, error) {
	body, err := s.getGroupResponseBody(groupID)

	return body.Data, err
}

func (s *Service) getGroupResponseBody(groupID string) (*GetGroupResponseBody, error) {
	body := &GetGroupResponseBody{}

	if groupID == "" {
		return body, fmt.Errorf("invalid group id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/groups/%s", groupID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &GroupNotFoundError{ID: groupID}
		}

		return nil
	})
}
