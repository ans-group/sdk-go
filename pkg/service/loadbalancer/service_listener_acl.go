package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetListenerACLs retrieves a list of ACLs
func (s *Service) GetListenerACLs(listenerID int, parameters connection.APIRequestParameters) ([]ACL, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[ACL], error) {
		return s.GetListenerACLsPaginated(listenerID, p)
	}, parameters)
}

// GetListenerACLsPaginated retrieves a paginated list of ACLs
func (s *Service) GetListenerACLsPaginated(listenerID int, parameters connection.APIRequestParameters) (*connection.Paginated[ACL], error) {
	if listenerID < 1 {
		return nil, fmt.Errorf("invalid listener id")
	}
	body, err := connection.Get[[]ACL](s.connection, fmt.Sprintf("/loadbalancers/v2/listeners/%d/acls", listenerID), parameters)
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[ACL], error) {
		return s.GetListenerACLsPaginated(listenerID, p)
	}), err
}
