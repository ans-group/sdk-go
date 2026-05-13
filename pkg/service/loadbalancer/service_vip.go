package loadbalancer

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) vipRes() *resource.Resource[VIP, int] {
	return resource.NewIntResource[VIP](s.connection, "/loadbalancers/v2/vips", "vip",
		func(id int) error { return &VIPNotFoundError{ID: id} })
}

// GetVIPs retrieves a list of VIPs
func (s *Service) GetVIPs(parameters connection.APIRequestParameters) ([]VIP, error) {
	return s.vipRes().List(parameters)
}

// GetVIPsPaginated retrieves a paginated list of VIPs
func (s *Service) GetVIPsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VIP], error) {
	return s.vipRes().ListPaginated(parameters)
}

// GetVIP retrieves a single VIP by id
func (s *Service) GetVIP(vipID int) (VIP, error) {
	return s.vipRes().Get(vipID)
}
