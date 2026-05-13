package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetTargetGroupTargets retrieves a list of targets
func (s *Service) GetTargetGroupTargets(groupID int, parameters connection.APIRequestParameters) ([]Target, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Target], error) {
		return s.GetTargetGroupTargetsPaginated(groupID, p)
	}, parameters)
}

// GetTargetGroupTargetsPaginated retrieves a paginated list of targets
func (s *Service) GetTargetGroupTargetsPaginated(groupID int, parameters connection.APIRequestParameters) (*connection.Paginated[Target], error) {
	if groupID < 1 {
		return nil, fmt.Errorf("invalid target group id")
	}
	body, err := connection.Get[[]Target](s.connection, fmt.Sprintf("/loadbalancers/v2/target-groups/%d/targets", groupID), parameters)
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Target], error) {
		return s.GetTargetGroupTargetsPaginated(groupID, p)
	}), err
}

// GetTargetGroupTarget retrieves a single target by id
func (s *Service) GetTargetGroupTarget(groupID int, targetID int) (Target, error) {
	if groupID < 1 {
		return Target{}, fmt.Errorf("invalid target group id")
	}
	if targetID < 1 {
		return Target{}, fmt.Errorf("invalid target id")
	}
	body, err := connection.Get[Target](s.connection, fmt.Sprintf("/loadbalancers/v2/target-groups/%d/targets/%d", groupID, targetID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&TargetNotFoundError{ID: groupID}))
	return body.Data, err
}

// CreateTargetGroupTarget creates a target
func (s *Service) CreateTargetGroupTarget(groupID int, req CreateTargetRequest) (int, error) {
	if groupID < 1 {
		return 0, fmt.Errorf("invalid target group id")
	}
	body, err := connection.Post[Target](s.connection, fmt.Sprintf("/loadbalancers/v2/target-groups/%d/targets", groupID), &req, connection.NotFoundResponseHandler(&TargetNotFoundError{ID: groupID}))
	return body.Data.ID, err
}

// PatchTargetGroupTarget patches a target
func (s *Service) PatchTargetGroupTarget(groupID int, targetID int, req PatchTargetRequest) error {
	if groupID < 1 {
		return fmt.Errorf("invalid target group id")
	}
	if targetID < 1 {
		return fmt.Errorf("invalid target id")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/loadbalancers/v2/target-groups/%d/targets/%d", groupID, targetID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&TargetNotFoundError{ID: groupID}))
}

// DeleteTargetGroupTarget deletes a target
func (s *Service) DeleteTargetGroupTarget(groupID int, targetID int) error {
	if groupID < 1 {
		return fmt.Errorf("invalid target group id")
	}
	if targetID < 1 {
		return fmt.Errorf("invalid target id")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/loadbalancers/v2/target-groups/%d/targets/%d", groupID, targetID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&TargetNotFoundError{ID: groupID}))
}
