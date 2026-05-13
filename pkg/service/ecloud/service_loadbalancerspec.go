package ecloud

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) loadBalancerSpecRes() *resource.Resource[LoadBalancerSpec, string] {
	return resource.NewStringResource[LoadBalancerSpec](s.connection, "/ecloud/v2/load-balancer-specs", "load balancer spec", func(id string) error {
		return &LoadBalancerSpecNotFoundError{ID: id}
	})
}

// GetLoadBalancerSpecs retrieves a list of load balancer specs
func (s *Service) GetLoadBalancerSpecs(parameters connection.APIRequestParameters) ([]LoadBalancerSpec, error) {
	return s.loadBalancerSpecRes().List(parameters)
}

// GetLoadBalancerSpecsPaginated retrieves a paginated list of load balancer specs
func (s *Service) GetLoadBalancerSpecsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[LoadBalancerSpec], error) {
	return s.loadBalancerSpecRes().ListPaginated(parameters)
}

// GetLoadBalancerSpec retrieves a single spec by id
func (s *Service) GetLoadBalancerSpec(lbSpecID string) (LoadBalancerSpec, error) {
	return s.loadBalancerSpecRes().Get(lbSpecID)
}
