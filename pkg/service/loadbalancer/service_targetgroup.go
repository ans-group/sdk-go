package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetTargetGroups retrieves a list of target groups
func (s *Service) GetTargetGroups(parameters connection.APIRequestParameters) ([]TargetGroup, error) {
	return connection.InvokeRequestAll(s.GetTargetGroupsPaginated, parameters)
}

// GetTargetGroupsPaginated retrieves a paginated list of target groups
func (s *Service) GetTargetGroupsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[TargetGroup], error) {
	body, err := connection.Get[[]TargetGroup](s.connection, "/loadbalancers/v2/target-groups", parameters)
	return connection.NewPaginated(body, parameters, s.GetTargetGroupsPaginated), err
}

// GetTargetGroup retrieves a single target group by id
func (s *Service) GetTargetGroup(groupID int) (TargetGroup, error) {
	if groupID < 1 {
		return TargetGroup{}, fmt.Errorf("invalid target group id")
	}
	body, err := connection.Get[TargetGroup](s.connection, fmt.Sprintf("/loadbalancers/v2/target-groups/%d", groupID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&TargetGroupNotFoundError{ID: groupID}))
	return body.Data, err
}

// CreateTargetGroup creates a target group
func (s *Service) CreateTargetGroup(req CreateTargetGroupRequest) (int, error) {
	body, err := connection.Post[TargetGroup](s.connection, "/loadbalancers/v2/target-groups", &req)
	return body.Data.ID, err
}

// PatchTargetGroup patches a target group
func (s *Service) PatchTargetGroup(groupID int, req PatchTargetGroupRequest) error {
	if groupID < 1 {
		return fmt.Errorf("invalid target group id")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/loadbalancers/v2/target-groups/%d", groupID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&TargetGroupNotFoundError{ID: groupID}))
}

// DeleteTargetGroup deletes a target group
func (s *Service) DeleteTargetGroup(groupID int) error {
	if groupID < 1 {
		return fmt.Errorf("invalid target group id")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/loadbalancers/v2/target-groups/%d", groupID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&TargetGroupNotFoundError{ID: groupID}))
}
