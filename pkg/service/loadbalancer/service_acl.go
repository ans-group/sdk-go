package loadbalancer

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) aclRes() *resource.Resource[ACL, int] {
	return resource.NewIntResource[ACL](s.connection, "/loadbalancers/v2/acls", "acl",
		func(id int) error { return &ACLNotFoundError{ID: id} })
}

// GetACLs retrieves a list of ACLs
// Currently, a target_group_id or listener_id filter must be provided for this to return data
func (s *Service) GetACLs(parameters connection.APIRequestParameters) ([]ACL, error) {
	return s.aclRes().List(parameters)
}

// GetACLsPaginated retrieves a paginated list of ACLs
// Currently, a target_group_id or listener_id filter must be provided for this to return data
func (s *Service) GetACLsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ACL], error) {
	return s.aclRes().ListPaginated(parameters)
}

// GetACL retrieves a single ACL by id
func (s *Service) GetACL(aclID int) (ACL, error) {
	return s.aclRes().Get(aclID)
}

// CreateACL creates an ACL
func (s *Service) CreateACL(req CreateACLRequest) (int, error) {
	data, err := s.aclRes().Create(&req)
	return data.ID, err
}

// PatchACL patches an ACL
func (s *Service) PatchACL(aclID int, req PatchACLRequest) error {
	return s.aclRes().Patch(aclID, &req)
}

// DeleteACL deletes an ACL
func (s *Service) DeleteACL(aclID int) error {
	return s.aclRes().Delete(aclID)
}
