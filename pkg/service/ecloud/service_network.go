package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetNetworks retrieves a list of networks
func (s *Service) GetNetworks(parameters connection.APIRequestParameters) ([]Network, error) {
	return connection.InvokeRequestAll(s.GetNetworksPaginated, parameters)
}

// GetNetworksPaginated retrieves a paginated list of networks
func (s *Service) GetNetworksPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Network], error) {
	body, err := connection.Get[[]Network](s.connection, "/ecloud/v2/networks", parameters)
	return connection.NewPaginated(body, parameters, s.GetNetworksPaginated), err
}

// GetNetwork retrieves a single network by id
func (s *Service) GetNetwork(networkID string) (Network, error) {
	if networkID == "" {
		return Network{}, fmt.Errorf("invalid network id")
	}
	body, err := connection.Get[Network](s.connection, fmt.Sprintf("/ecloud/v2/networks/%s", networkID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&NetworkNotFoundError{ID: networkID}))
	return body.Data, err
}

// CreateNetwork creates a new Network
func (s *Service) CreateNetwork(req CreateNetworkRequest) (string, error) {
	body, err := connection.Post[Network](s.connection, "/ecloud/v2/networks", &req)
	return body.Data.ID, err
}

// PatchNetwork patches a Network
func (s *Service) PatchNetwork(networkID string, req PatchNetworkRequest) error {
	if networkID == "" {
		return fmt.Errorf("invalid network id")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/ecloud/v2/networks/%s", networkID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&NetworkNotFoundError{ID: networkID}))
}

// DeleteNetwork deletes a Network
func (s *Service) DeleteNetwork(networkID string) error {
	if networkID == "" {
		return fmt.Errorf("invalid network id")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ecloud/v2/networks/%s", networkID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&NetworkNotFoundError{ID: networkID}))
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
