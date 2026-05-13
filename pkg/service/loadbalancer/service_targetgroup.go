package loadbalancer

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) targetGroupRes() *resource.Resource[TargetGroup, int] {
	return resource.NewIntResource[TargetGroup](s.connection, "/loadbalancers/v2/target-groups", "target group",
		func(id int) error { return &TargetGroupNotFoundError{ID: id} })
}

// GetTargetGroups retrieves a list of target groups
func (s *Service) GetTargetGroups(parameters connection.APIRequestParameters) ([]TargetGroup, error) {
	return s.targetGroupRes().List(parameters)
}

// GetTargetGroupsPaginated retrieves a paginated list of target groups
func (s *Service) GetTargetGroupsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[TargetGroup], error) {
	return s.targetGroupRes().ListPaginated(parameters)
}

// GetTargetGroup retrieves a single target group by id
func (s *Service) GetTargetGroup(groupID int) (TargetGroup, error) {
	return s.targetGroupRes().Get(groupID)
}

// CreateTargetGroup creates a target group
func (s *Service) CreateTargetGroup(req CreateTargetGroupRequest) (int, error) {
	data, err := s.targetGroupRes().Create(&req)
	return data.ID, err
}

// PatchTargetGroup patches a target group
func (s *Service) PatchTargetGroup(groupID int, req PatchTargetGroupRequest) error {
	return s.targetGroupRes().Patch(groupID, &req)
}

// DeleteTargetGroup deletes a target group
func (s *Service) DeleteTargetGroup(groupID int) error {
	return s.targetGroupRes().Delete(groupID)
}
