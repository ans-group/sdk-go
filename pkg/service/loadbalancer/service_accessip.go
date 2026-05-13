package loadbalancer

import (
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) accessIPRes() *resource.Resource[AccessIP, int] {
	return resource.NewIntResource[AccessIP](s.connection, "/loadbalancers/v2/access-ips", "access",
		func(id int) error { return &AccessIPNotFoundError{ID: id} })
}

// GetAccessIP retrieves a single access IP by id
func (s *Service) GetAccessIP(accessID int) (AccessIP, error) {
	return s.accessIPRes().Get(accessID)
}

// PatchAccessIP patches an access IP
func (s *Service) PatchAccessIP(accessID int, req PatchAccessIPRequest) error {
	return s.accessIPRes().Patch(accessID, &req)
}

// DeleteAccessIP deletes an access IP
func (s *Service) DeleteAccessIP(accessID int) error {
	return s.accessIPRes().Delete(accessID)
}
