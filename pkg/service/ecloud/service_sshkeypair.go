package ecloud

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) sshKeyPairRes() *resource.Resource[SSHKeyPair, string] {
	return resource.NewStringResource[SSHKeyPair](s.connection, "/ecloud/v2/ssh-key-pairs", "SSH key pair", func(id string) error {
		return &SSHKeyPairNotFoundError{ID: id}
	})
}

// GetSSHKeyPairs retrieves a list of keypairs
func (s *Service) GetSSHKeyPairs(parameters connection.APIRequestParameters) ([]SSHKeyPair, error) {
	return s.sshKeyPairRes().List(parameters)
}

// GetSSHKeyPairsPaginated retrieves a paginated list of keypairs
func (s *Service) GetSSHKeyPairsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[SSHKeyPair], error) {
	return s.sshKeyPairRes().ListPaginated(parameters)
}

// GetSSHKeyPair retrieves a single keypair by id
func (s *Service) GetSSHKeyPair(keypairID string) (SSHKeyPair, error) {
	return s.sshKeyPairRes().Get(keypairID)
}

// CreateSSHKeyPair creates a new SSHKeyPair
func (s *Service) CreateSSHKeyPair(req CreateSSHKeyPairRequest) (string, error) {
	data, err := s.sshKeyPairRes().Create(&req)
	return data.ID, err
}

// PatchSSHKeyPair patches a SSHKeyPair
func (s *Service) PatchSSHKeyPair(keypairID string, req PatchSSHKeyPairRequest) error {
	return s.sshKeyPairRes().Patch(keypairID, &req)
}

// DeleteSSHKeyPair deletes a SSHKeyPair
func (s *Service) DeleteSSHKeyPair(keypairID string) error {
	return s.sshKeyPairRes().Delete(keypairID)
}
