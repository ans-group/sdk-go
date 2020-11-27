package loadbalancer

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetClusters retrieves a list of clusters
func (s *Service) GetClusters(parameters connection.APIRequestParameters) ([]Cluster, error) {
	var sites []Cluster

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetClustersPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, site := range response.(*PaginatedCluster).Items {
			sites = append(sites, site)
		}
	}

	return sites, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetClustersPaginated retrieves a paginated list of clusters
func (s *Service) GetClustersPaginated(parameters connection.APIRequestParameters) (*PaginatedCluster, error) {
	body, err := s.getClustersPaginatedResponseBody(parameters)

	return NewPaginatedCluster(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetClustersPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getClustersPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetClusterSliceResponseBody, error) {
	body := &GetClusterSliceResponseBody{}

	response, err := s.connection.Get("/loadbalancers/v2/clusters", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetCluster retrieves a single cluster by id
func (s *Service) GetCluster(clusterID string) (Cluster, error) {
	body, err := s.getClusterResponseBody(clusterID)

	return body.Data, err
}

func (s *Service) getClusterResponseBody(clusterID string) (*GetClusterResponseBody, error) {
	body := &GetClusterResponseBody{}

	if clusterID == "" {
		return body, fmt.Errorf("invalid cluster id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/clusters/%s", clusterID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ClusterNotFoundError{ID: clusterID}
		}

		return nil
	})
}

// PatchCluster patches a Cluster
func (s *Service) PatchCluster(clusterID string, req PatchClusterRequest) error {
	_, err := s.patchClusterResponseBody(clusterID, req)

	return err
}

func (s *Service) patchClusterResponseBody(clusterID string, req PatchClusterRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if clusterID == "" {
		return body, fmt.Errorf("invalid cluster id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/loadbalancers/v2/clusters/%s", clusterID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ClusterNotFoundError{ID: clusterID}
		}

		return nil
	})
}

// DeleteCluster deletes a Cluster
func (s *Service) DeleteCluster(clusterID string) error {
	_, err := s.deleteClusterResponseBody(clusterID)

	return err
}

func (s *Service) deleteClusterResponseBody(clusterID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if clusterID == "" {
		return body, fmt.Errorf("invalid cluster id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/loadbalancers/v2/clusters/%s", clusterID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ClusterNotFoundError{ID: clusterID}
		}

		return nil
	})
}
