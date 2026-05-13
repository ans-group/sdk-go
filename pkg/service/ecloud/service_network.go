package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) networkRes() *resource.Resource[Network, string] {
	return resource.NewStringResource[Network](s.connection, "/ecloud/v2/networks", "network", func(id string) error {
		return &NetworkNotFoundError{ID: id}
	})
}

// GetNetworks retrieves a list of networks
func (s *Service) GetNetworks(parameters connection.APIRequestParameters) ([]Network, error) {
	return s.networkRes().List(parameters)
}

// GetNetworksPaginated retrieves a paginated list of networks
func (s *Service) GetNetworksPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Network], error) {
	return s.networkRes().ListPaginated(parameters)
}

// GetNetwork retrieves a single network by id
func (s *Service) GetNetwork(networkID string) (Network, error) {
	return s.networkRes().Get(networkID)
}

// CreateNetwork creates a new Network
func (s *Service) CreateNetwork(req CreateNetworkRequest) (string, error) {
	data, err := s.networkRes().Create(&req)
	return data.ID, err
}

// PatchNetwork patches a Network
func (s *Service) PatchNetwork(networkID string, req PatchNetworkRequest) error {
	return s.networkRes().Patch(networkID, &req)
}

// DeleteNetwork deletes a Network
func (s *Service) DeleteNetwork(networkID string) error {
	return s.networkRes().Delete(networkID)
}

// GetNetworkNICs retrieves a list of firewall rule nics
func (s *Service) GetNetworkNICs(networkID string, parameters connection.APIRequestParameters) ([]NIC, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[NIC], error) {
		return s.GetNetworkNICsPaginated(networkID, p)
	}, parameters)
}

// GetNetworkNICsPaginated retrieves a paginated list of firewall rule nics
func (s *Service) GetNetworkNICsPaginated(networkID string, parameters connection.APIRequestParameters) (*connection.Paginated[NIC], error) {
	body, err := connection.Get[[]NIC](s.connection, fmt.Sprintf("/ecloud/v2/networks/%s/nics", networkID), parameters)
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[NIC], error) {
		return s.GetNetworkNICsPaginated(networkID, p)
	}), err
}

// GetNetworkTasks retrieves a list of Network tasks
func (s *Service) GetNetworkTasks(networkID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetNetworkTasksPaginated(networkID, p)
	}, parameters)
}

// GetNetworkTasksPaginated retrieves a paginated list of Network tasks
func (s *Service) GetNetworkTasksPaginated(networkID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	if networkID == "" {
		return nil, fmt.Errorf("invalid network id")
	}
	body, err := connection.Get[[]Task](s.connection, fmt.Sprintf("/ecloud/v2/networks/%s/tasks", networkID), parameters, connection.NotFoundResponseHandler(&NetworkNotFoundError{ID: networkID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetNetworkTasksPaginated(networkID, p)
	}), err
}
