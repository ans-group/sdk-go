package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetLoadBalancerClusters retrieves a list of load balancer clusters
func (s *Service) GetLoadBalancerClusters(parameters connection.APIRequestParameters) ([]LoadBalancerCluster, error) {
	var sites []LoadBalancerCluster

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetLoadBalancerClustersPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, site := range response.(*PaginatedLoadBalancerCluster).Items {
			sites = append(sites, site)
		}
	}

	return sites, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetLoadBalancerClustersPaginated retrieves a paginated list of load balancer clusters
func (s *Service) GetLoadBalancerClustersPaginated(parameters connection.APIRequestParameters) (*PaginatedLoadBalancerCluster, error) {
	body, err := s.getLoadBalancerClustersPaginatedResponseBody(parameters)

	return NewPaginatedLoadBalancerCluster(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetLoadBalancerClustersPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getLoadBalancerClustersPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetLoadBalancerClusterSliceResponseBody, error) {
	body := &GetLoadBalancerClusterSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/lbcs", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetLoadBalancerCluster retrieves a single load balancer cluster by id
func (s *Service) GetLoadBalancerCluster(lbcsID string) (LoadBalancerCluster, error) {
	body, err := s.getLoadBalancerClusterResponseBody(lbcsID)

	return body.Data, err
}

func (s *Service) getLoadBalancerClusterResponseBody(lbcsID string) (*GetLoadBalancerClusterResponseBody, error) {
	body := &GetLoadBalancerClusterResponseBody{}

	if lbcsID == "" {
		return body, fmt.Errorf("invalid load balancer cluster id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/lbcs/%s", lbcsID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &LoadBalancerClusterNotFoundError{ID: lbcsID}
		}

		return nil
	})
}
