package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetLoadBalancerSpecs retrieves a list of load balancer specs
func (s *Service) GetLoadBalancerSpecs(parameters connection.APIRequestParameters) ([]LoadBalancerSpec, error) {
	return connection.InvokeRequestAll(s.GetLoadBalancerSpecsPaginated, parameters)
}

// GetLoadBalancerSpecsPaginated retrieves a paginated list of load balancer specs
func (s *Service) GetLoadBalancerSpecsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[LoadBalancerSpec], error) {
	body, err := connection.Get[[]LoadBalancerSpec](s.connection, "/ecloud/v2/load-balancer-specs", parameters)
	return connection.NewPaginated(body, parameters, s.GetLoadBalancerSpecsPaginated), err
}

// GetLoadBalancerSpec retrieves a single spec by id
func (s *Service) GetLoadBalancerSpec(lbSpecID string) (LoadBalancerSpec, error) {
	if lbSpecID == "" {
		return LoadBalancerSpec{}, fmt.Errorf("invalid load balancer spec id")
	}
	body, err := connection.Get[LoadBalancerSpec](s.connection, fmt.Sprintf("/ecloud/v2/load-balancer-specs/%s", lbSpecID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&LoadBalancerSpecNotFoundError{ID: lbSpecID}))
	return body.Data, err
}
