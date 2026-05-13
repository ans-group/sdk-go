package account

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetClients retrieves a list of clients
func (s *Service) GetClients(parameters connection.APIRequestParameters) ([]Client, error) {
	return connection.InvokeRequestAll(s.GetClientsPaginated, parameters)
}

// GetClientsPaginated retrieves a paginated list of clients
func (s *Service) GetClientsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Client], error) {
	body, err := connection.Get[[]Client](s.connection, "/account/v1/clients", parameters)
	return connection.NewPaginated(body, parameters, s.GetClientsPaginated), err
}

// GetClient retrieves a single client by id
func (s *Service) GetClient(clientID int) (Client, error) {
	if clientID < 1 {
		return Client{}, fmt.Errorf("invalid client id")
	}
	body, err := connection.Get[Client](s.connection, fmt.Sprintf("/account/v1/clients/%d", clientID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ClientNotFoundError{ID: clientID}))
	return body.Data, err
}

// CreateClient creates a new client
func (s *Service) CreateClient(req CreateClientRequest) (int, error) {
	body, err := connection.Post[Client](s.connection, "/account/v1/clients", &req)
	return body.Data.ID, err
}

// PatchClient patches a client
func (s *Service) PatchClient(clientID int, patch PatchClientRequest) error {
	if clientID < 1 {
		return fmt.Errorf("invalid client id")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/account/v1/clients/%d", clientID), &patch, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&ClientNotFoundError{ID: clientID}))
}

// DeleteClient removes a client
func (s *Service) DeleteClient(clientID int) error {
	if clientID < 1 {
		return fmt.Errorf("invalid client id")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/account/v1/clients/%d", clientID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&ClientNotFoundError{ID: clientID}))
}
