package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) listenerAccessIPRes() *resource.SubResource[AccessIP, int, int] {
	return resource.NewIntIntSubResource[AccessIP](s.connection,
		func(listenerID int) string {
			return fmt.Sprintf("/loadbalancers/v2/listeners/%d/access-ips", listenerID)
		},
		"listener", "id", func(listenerID int) error { return &AccessIPNotFoundError{ID: listenerID} },
		"access", "id", func(listenerID, _ int) error { return &AccessIPNotFoundError{ID: listenerID} })
}

// GetListenerAccessIPs retrieves a list of access IPs
func (s *Service) GetListenerAccessIPs(listenerID int, parameters connection.APIRequestParameters) ([]AccessIP, error) {
	return s.listenerAccessIPRes().List(listenerID, parameters)
}

// GetListenerAccessIPsPaginated retrieves a paginated list of access IPs
func (s *Service) GetListenerAccessIPsPaginated(listenerID int, parameters connection.APIRequestParameters) (*connection.Paginated[AccessIP], error) {
	return s.listenerAccessIPRes().ListPaginated(listenerID, parameters)
}

// GetListenerAccessIP retrieves a single access IP by id
func (s *Service) GetListenerAccessIP(listenerID int, accessID int) (AccessIP, error) {
	return s.listenerAccessIPRes().Get(listenerID, accessID)
}

// CreateListenerAccessIP creates an access IP
func (s *Service) CreateListenerAccessIP(listenerID int, req CreateAccessIPRequest) (int, error) {
	accessIP, err := s.listenerAccessIPRes().Create(listenerID, &req)
	return accessIP.ID, err
}
