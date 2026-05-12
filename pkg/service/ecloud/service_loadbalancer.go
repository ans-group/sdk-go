package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetLoadBalancers retrieves a list of load balancers
func (s *Service) GetLoadBalancers(parameters connection.APIRequestParameters) ([]LoadBalancer, error) {
	return connection.InvokeRequestAll(s.GetLoadBalancersPaginated, parameters)
}

// GetLoadBalancersPaginated retrieves a paginated list of lbs
func (s *Service) GetLoadBalancersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[LoadBalancer], error) {
	body, err := connection.Get[[]LoadBalancer](s.connection, "/ecloud/v2/load-balancers", parameters)
	return connection.NewPaginated(body, parameters, s.GetLoadBalancersPaginated), err
}

// GetLoadBalancer retrieves a single lb by id
func (s *Service) GetLoadBalancer(loadbalancerID string) (LoadBalancer, error) {
	if loadbalancerID == "" {
		return LoadBalancer{}, fmt.Errorf("invalid load balancer id")
	}
	body, err := connection.Get[LoadBalancer](s.connection, fmt.Sprintf("/ecloud/v2/load-balancers/%s", loadbalancerID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&LoadBalancerNotFoundError{ID: loadbalancerID}))
	return body.Data, err
}

// CreateLoadBalancer creates a new LoadBalancer
func (s *Service) CreateLoadBalancer(req CreateLoadBalancerRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/load-balancers", &req)
	return body.Data, err
}

// PatchLoadBalancer patches a LoadBalancer
func (s *Service) PatchLoadBalancer(loadbalancerID string, req PatchLoadBalancerRequest) (TaskReference, error) {
	if loadbalancerID == "" {
		return TaskReference{}, fmt.Errorf("invalid load balancer id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/load-balancers/%s", loadbalancerID), &req, connection.NotFoundResponseHandler(&LoadBalancerNotFoundError{ID: loadbalancerID}))
	return body.Data, err
}

// DeleteLoadBalancer deletes a LoadBalancer
func (s *Service) DeleteLoadBalancer(loadbalancerID string) (string, error) {
	if loadbalancerID == "" {
		return "", fmt.Errorf("invalid load balancer id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/load-balancers/%s", loadbalancerID), nil, connection.NotFoundResponseHandler(&LoadBalancerNotFoundError{ID: loadbalancerID}))
	return body.Data.TaskID, err
}
