package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetTargetGroupACLs retrieves a list of ACLs
func (s *Service) GetTargetGroupACLs(targetGroupID int, parameters connection.APIRequestParameters) ([]ACL, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[ACL], error) {
		return s.GetTargetGroupACLsPaginated(targetGroupID, p)
	}, parameters)
}

// GetTargetGroupACLsPaginated retrieves a paginated list of ACLs
func (s *Service) GetTargetGroupACLsPaginated(targetGroupID int, parameters connection.APIRequestParameters) (*connection.Paginated[ACL], error) {
	if targetGroupID < 1 {
		return nil, fmt.Errorf("invalid target group id")
	}
	body, err := connection.Get[[]ACL](s.connection, fmt.Sprintf("/loadbalancers/v2/target-groups/%d/acls", targetGroupID), parameters)
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[ACL], error) {
		return s.GetTargetGroupACLsPaginated(targetGroupID, p)
	}), err
}
