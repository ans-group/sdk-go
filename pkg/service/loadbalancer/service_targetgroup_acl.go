package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) targetGroupACLRes() *resource.SubResourceList[ACL, int] {
	return resource.NewIntSubResourceList[ACL](s.connection,
		func(targetGroupID int) string {
			return fmt.Sprintf("/loadbalancers/v2/target-groups/%d/acls", targetGroupID)
		},
		"target group", "id", nil)
}

// GetTargetGroupACLs retrieves a list of ACLs
func (s *Service) GetTargetGroupACLs(targetGroupID int, parameters connection.APIRequestParameters) ([]ACL, error) {
	return s.targetGroupACLRes().List(targetGroupID, parameters)
}

// GetTargetGroupACLsPaginated retrieves a paginated list of ACLs
func (s *Service) GetTargetGroupACLsPaginated(targetGroupID int, parameters connection.APIRequestParameters) (*connection.Paginated[ACL], error) {
	return s.targetGroupACLRes().ListPaginated(targetGroupID, parameters)
}
