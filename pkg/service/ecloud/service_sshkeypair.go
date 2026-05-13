package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetSSHKeyPairs retrieves a list of keypairs
func (s *Service) GetSSHKeyPairs(parameters connection.APIRequestParameters) ([]SSHKeyPair, error) {
	return connection.InvokeRequestAll(s.GetSSHKeyPairsPaginated, parameters)
}

// GetSSHKeyPairsPaginated retrieves a paginated list of keypairs
func (s *Service) GetSSHKeyPairsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[SSHKeyPair], error) {
	body, err := connection.Get[[]SSHKeyPair](s.connection, "/ecloud/v2/ssh-key-pairs", parameters)
	return connection.NewPaginated(body, parameters, s.GetSSHKeyPairsPaginated), err
}

// GetSSHKeyPair retrieves a single keypair by id
func (s *Service) GetSSHKeyPair(keypairID string) (SSHKeyPair, error) {
	if keypairID == "" {
		return SSHKeyPair{}, fmt.Errorf("invalid SSH key pair id")
	}
	body, err := connection.Get[SSHKeyPair](s.connection, fmt.Sprintf("/ecloud/v2/ssh-key-pairs/%s", keypairID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&SSHKeyPairNotFoundError{ID: keypairID}))
	return body.Data, err
}

// CreateSSHKeyPair creates a new SSHKeyPair
func (s *Service) CreateSSHKeyPair(req CreateSSHKeyPairRequest) (string, error) {
	body, err := connection.Post[SSHKeyPair](s.connection, "/ecloud/v2/ssh-key-pairs", &req)
	return body.Data.ID, err
}

// PatchSSHKeyPair patches a SSHKeyPair
func (s *Service) PatchSSHKeyPair(keypairID string, req PatchSSHKeyPairRequest) error {
	if keypairID == "" {
		return fmt.Errorf("invalid SSH key pair id")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/ecloud/v2/ssh-key-pairs/%s", keypairID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&SSHKeyPairNotFoundError{ID: keypairID}))
}

// DeleteSSHKeyPair deletes a SSHKeyPair
func (s *Service) DeleteSSHKeyPair(keypairID string) error {
	if keypairID == "" {
		return fmt.Errorf("invalid SSH key pair id")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ecloud/v2/ssh-key-pairs/%s", keypairID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&SSHKeyPairNotFoundError{ID: keypairID}))
}
