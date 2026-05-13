package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) listenerACLRes() *resource.SubResourceList[ACL, int] {
	return resource.NewIntSubResourceList[ACL](s.connection,
		func(listenerID int) string { return fmt.Sprintf("/loadbalancers/v2/listeners/%d/acls", listenerID) },
		"listener", "id", nil)
}

// GetListenerACLs retrieves a list of ACLs
func (s *Service) GetListenerACLs(listenerID int, parameters connection.APIRequestParameters) ([]ACL, error) {
	return s.listenerACLRes().List(listenerID, parameters)
}

// GetListenerACLsPaginated retrieves a paginated list of ACLs
func (s *Service) GetListenerACLsPaginated(listenerID int, parameters connection.APIRequestParameters) (*connection.Paginated[ACL], error) {
	return s.listenerACLRes().ListPaginated(listenerID, parameters)
}
