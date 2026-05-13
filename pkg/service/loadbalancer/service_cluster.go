package loadbalancer

import (
	"errors"
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetClusters retrieves a list of clusters
func (s *Service) GetClusters(parameters connection.APIRequestParameters) ([]Cluster, error) {
	return connection.InvokeRequestAll(s.GetClustersPaginated, parameters)
}

// GetClustersPaginated retrieves a paginated list of clusters
func (s *Service) GetClustersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Cluster], error) {
	body, err := connection.Get[[]Cluster](s.connection, "/loadbalancers/v2/clusters", parameters)
	return connection.NewPaginated(body, parameters, s.GetClustersPaginated), err
}

// GetCluster retrieves a single cluster by id
func (s *Service) GetCluster(clusterID int) (Cluster, error) {
	if clusterID < 1 {
		return Cluster{}, fmt.Errorf("invalid cluster id")
	}
	body, err := connection.Get[Cluster](s.connection, fmt.Sprintf("/loadbalancers/v2/clusters/%d", clusterID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ClusterNotFoundError{ID: clusterID}))
	return body.Data, err
}

// PatchCluster patches a Cluster
func (s *Service) PatchCluster(clusterID int, req PatchClusterRequest) error {
	if clusterID < 1 {
		return fmt.Errorf("invalid cluster id")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/loadbalancers/v2/clusters/%d", clusterID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&ClusterNotFoundError{ID: clusterID}))
}

// DeployCluster deploys a Cluster
func (s *Service) DeployCluster(clusterID int) error {
	if clusterID < 1 {
		return fmt.Errorf("invalid cluster id")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/loadbalancers/v2/clusters/%d/deploy", clusterID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&ClusterNotFoundError{ID: clusterID}))
}

// ValidateCluster validates a cluster
func (s *Service) ValidateCluster(clusterID int) error {
	response := &connection.APIResponse{}

	if clusterID < 1 {
		return fmt.Errorf("invalid cluster id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/clusters/%d/validate", clusterID), connection.APIRequestParameters{})
	if err != nil {
		return err
	}

	if response.StatusCode == 422 {
		body := &validateClusterResponseBody{}

		return errors.New(body.Error())
	}

	return response.HandleResponse(&connection.APIResponseBody{}, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ClusterNotFoundError{ID: clusterID}
		}

		return nil
	})
}

type validateClusterResponseBody struct {
	connection.APIResponseBody
}

// GetClusterACLTemplates retrieves a single cluster by id
func (s *Service) GetClusterACLTemplates(clusterID int) (ACLTemplates, error) {
	if clusterID < 1 {
		return ACLTemplates{}, fmt.Errorf("invalid cluster id")
	}
	body, err := connection.Get[ACLTemplates](s.connection, fmt.Sprintf("/loadbalancers/v2/clusters/%d/acl-templates", clusterID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ClusterNotFoundError{ID: clusterID}))
	return body.Data, err
}
