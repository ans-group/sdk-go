package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetLoadBalancerClusters retrieves a list of load balancer clusters
func (s *Service) GetLoadBalancerClusters(parameters connection.APIRequestParameters) ([]LoadBalancerCluster, error) {
	var lbcs []LoadBalancerCluster

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetLoadBalancerClustersPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, lbc := range response.(*PaginatedLoadBalancerCluster).Items {
			lbcs = append(lbcs, lbc)
		}
	}

	return lbcs, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
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
func (s *Service) GetLoadBalancerCluster(lbcID string) (LoadBalancerCluster, error) {
	body, err := s.getLoadBalancerClusterResponseBody(lbcID)

	return body.Data, err
}

func (s *Service) getLoadBalancerClusterResponseBody(lbcID string) (*GetLoadBalancerClusterResponseBody, error) {
	body := &GetLoadBalancerClusterResponseBody{}

	if lbcID == "" {
		return body, fmt.Errorf("invalid load balancer cluster id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/lbcs/%s", lbcID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &LoadBalancerClusterNotFoundError{ID: lbcID}
		}

		return nil
	})
}

// CreateLoadBalancerCluster creates a new LoadBalancerCluster
func (s *Service) CreateLoadBalancerCluster(req CreateLoadBalancerClusterRequest) (string, error) {
	body, err := s.createLoadBalancerClusterResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createLoadBalancerClusterResponseBody(req CreateLoadBalancerClusterRequest) (*GetLoadBalancerClusterResponseBody, error) {
	body := &GetLoadBalancerClusterResponseBody{}

	response, err := s.connection.Post("/ecloud/v2/lbcs", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchLoadBalancerCluster patches a LoadBalancerCluster
func (s *Service) PatchLoadBalancerCluster(lbcID string, req PatchLoadBalancerClusterRequest) error {
	_, err := s.patchLoadBalancerClusterResponseBody(lbcID, req)

	return err
}

func (s *Service) patchLoadBalancerClusterResponseBody(lbcID string, req PatchLoadBalancerClusterRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if lbcID == "" {
		return body, fmt.Errorf("invalid load balancer cluster id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/lbcs/%s", lbcID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &LoadBalancerClusterNotFoundError{ID: lbcID}
		}

		return nil
	})
}

// DeleteLoadBalancerCluster deletes a LoadBalancerCluster
func (s *Service) DeleteLoadBalancerCluster(lbcID string) error {
	_, err := s.deleteLoadBalancerClusterResponseBody(lbcID)

	return err
}

func (s *Service) deleteLoadBalancerClusterResponseBody(lbcID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if lbcID == "" {
		return body, fmt.Errorf("invalid load balancer cluster id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/lbcs/%s", lbcID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &LoadBalancerClusterNotFoundError{ID: lbcID}
		}

		return nil
	})
}
