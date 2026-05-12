package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetVIPs retrieves a list of VIPs
func (s *Service) GetVIPs(parameters connection.APIRequestParameters) ([]VIP, error) {
	return connection.InvokeRequestAll(s.GetVIPsPaginated, parameters)
}

// GetVIPsPaginated retrieves a paginated list of VIPs
func (s *Service) GetVIPsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VIP], error) {
	body, err := connection.Get[[]VIP](s.connection, "/loadbalancers/v2/vips", parameters)
	return connection.NewPaginated(body, parameters, s.GetVIPsPaginated), err
}

// GetVIP retrieves a single VIP by id
func (s *Service) GetVIP(vipID int) (VIP, error) {
	if vipID < 1 {
		return VIP{}, fmt.Errorf("invalid vip id")
	}
	body, err := connection.Get[VIP](s.connection, fmt.Sprintf("/loadbalancers/v2/vips/%d", vipID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&VIPNotFoundError{ID: vipID}))
	return body.Data, err
}
