package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) targetGroupTargetRes() *resource.SubResource[Target, int, int] {
	return resource.NewIntIntSubResource[Target](s.connection,
		func(groupID int) string { return fmt.Sprintf("/loadbalancers/v2/target-groups/%d/targets", groupID) },
		"target group", "id", func(groupID int) error { return &TargetNotFoundError{ID: groupID} },
		"target", "id", func(groupID, _ int) error { return &TargetNotFoundError{ID: groupID} })
}

// GetTargetGroupTargets retrieves a list of targets
func (s *Service) GetTargetGroupTargets(groupID int, parameters connection.APIRequestParameters) ([]Target, error) {
	return s.targetGroupTargetRes().List(groupID, parameters)
}

// GetTargetGroupTargetsPaginated retrieves a paginated list of targets
func (s *Service) GetTargetGroupTargetsPaginated(groupID int, parameters connection.APIRequestParameters) (*connection.Paginated[Target], error) {
	return s.targetGroupTargetRes().ListPaginated(groupID, parameters)
}

// GetTargetGroupTarget retrieves a single target by id
func (s *Service) GetTargetGroupTarget(groupID int, targetID int) (Target, error) {
	return s.targetGroupTargetRes().Get(groupID, targetID)
}

// CreateTargetGroupTarget creates a target
func (s *Service) CreateTargetGroupTarget(groupID int, req CreateTargetRequest) (int, error) {
	target, err := s.targetGroupTargetRes().Create(groupID, &req)
	return target.ID, err
}

// PatchTargetGroupTarget patches a target
func (s *Service) PatchTargetGroupTarget(groupID int, targetID int, req PatchTargetRequest) error {
	return s.targetGroupTargetRes().Patch(groupID, targetID, &req)
}

// DeleteTargetGroupTarget deletes a target
func (s *Service) DeleteTargetGroupTarget(groupID int, targetID int) error {
	return s.targetGroupTargetRes().Delete(groupID, targetID)
}
