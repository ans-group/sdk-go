package account

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) clientRes() *resource.Resource[Client, int] {
	return resource.NewIntResource[Client](s.connection, "/account/v1/clients", "client",
		func(id int) error { return &ClientNotFoundError{ID: id} })
}

// GetClients retrieves a list of clients
func (s *Service) GetClients(parameters connection.APIRequestParameters) ([]Client, error) {
	return s.clientRes().List(parameters)
}

// GetClientsPaginated retrieves a paginated list of clients
func (s *Service) GetClientsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Client], error) {
	return s.clientRes().ListPaginated(parameters)
}

// GetClient retrieves a single client by id
func (s *Service) GetClient(clientID int) (Client, error) {
	return s.clientRes().Get(clientID)
}

// CreateClient creates a new client
func (s *Service) CreateClient(req CreateClientRequest) (int, error) {
	data, err := s.clientRes().Create(&req)
	return data.ID, err
}

// PatchClient patches a client
func (s *Service) PatchClient(clientID int, patch PatchClientRequest) error {
	return s.clientRes().Patch(clientID, &patch)
}

// DeleteClient removes a client
func (s *Service) DeleteClient(clientID int) error {
	return s.clientRes().Delete(clientID)
}
