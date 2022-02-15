package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetLoadBalancerNetworks retrieves a list of load balancer networks
func (s *Service) GetLoadBalancerNetworks(parameters connection.APIRequestParameters) ([]LoadBalancerNetwork, error) {
	var networks []LoadBalancerNetwork

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetLoadBalancerNetworksPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, network := range response.(*PaginatedLoadBalancerNetwork).Items {
			networks = append(networks, network)
		}
	}

	return networks, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetLoadBalancerNetworksPaginated retrieves a paginated list of load balancer networks
func (s *Service) GetLoadBalancerNetworksPaginated(parameters connection.APIRequestParameters) (*PaginatedLoadBalancerNetwork, error) {
	body, err := s.getLoadBalancerNetworksPaginatedResponseBody(parameters)

	return NewPaginatedLoadBalancerNetwork(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetLoadBalancerNetworksPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getLoadBalancerNetworksPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetLoadBalancerNetworkSliceResponseBody, error) {
	body := &GetLoadBalancerNetworkSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/load-balancer-networks", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetLoadBalancerNetwork retrieves a single network by id
func (s *Service) GetLoadBalancerNetwork(lbNetworkID string) (LoadBalancerNetwork, error) {
	body, err := s.getLoadBalancerNetworkResponseBody(lbNetworkID)

	return body.Data, err
}

func (s *Service) getLoadBalancerNetworkResponseBody(lbNetworkID string) (*GetLoadBalancerNetworkResponseBody, error) {
	body := &GetLoadBalancerNetworkResponseBody{}

	if lbNetworkID == "" {
		return body, fmt.Errorf("invalid load balancer network id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/load-balancer-networks/%s", lbNetworkID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &LoadBalancerNetworkNotFoundError{ID: lbNetworkID}
		}

		return nil
	})
}

// CreateLoadBalancerNetwork creates a new LoadBalancerNetwork
func (s *Service) CreateLoadBalancerNetwork(req CreateLoadBalancerNetworkRequest) (TaskReference, error) {
	body, err := s.createLoadBalancerNetworkResponseBody(req)

	return body.Data, err
}

func (s *Service) createLoadBalancerNetworkResponseBody(req CreateLoadBalancerNetworkRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	response, err := s.connection.Post("/ecloud/v2/load-balancer-networks", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchLoadBalancerNetwork patches a LoadBalancerNetwork
func (s *Service) PatchLoadBalancerNetwork(lbNetworkID string, req PatchLoadBalancerNetworkRequest) (TaskReference, error) {
	body, err := s.patchLoadBalancerNetworkResponseBody(lbNetworkID, req)

	return body.Data, err
}

func (s *Service) patchLoadBalancerNetworkResponseBody(lbNetworkID string, req PatchLoadBalancerNetworkRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if lbNetworkID == "" {
		return body, fmt.Errorf("invalid load balancer network id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/load-balancer-networks/%s", lbNetworkID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &LoadBalancerNetworkNotFoundError{ID: lbNetworkID}
		}

		return nil
	})
}

// DeleteLoadBalancerNetwork deletes a LoadBalancerNetwork
func (s *Service) DeleteLoadBalancerNetwork(lbNetworkID string) (string, error) {
	body, err := s.deleteLoadBalancerNetworkResponseBody(lbNetworkID)

	return body.Data.TaskID, err
}

func (s *Service) deleteLoadBalancerNetworkResponseBody(lbNetworkID string) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if lbNetworkID == "" {
		return body, fmt.Errorf("invalid load balancer network id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/load-balancer-networks/%s", lbNetworkID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &LoadBalancerNetworkNotFoundError{ID: lbNetworkID}
		}

		return nil
	})
}
