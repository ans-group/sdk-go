package cloudflare

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) accountRes() *resource.Resource[Account, string] {
	return resource.NewStringResource[Account](s.connection, "/cloudflare/v1/accounts", "account",
		func(id string) error { return &AccountNotFoundError{ID: id} })
}

// GetAccounts retrieves a list of accounts
func (s *Service) GetAccounts(parameters connection.APIRequestParameters) ([]Account, error) {
	return s.accountRes().List(parameters)
}

// GetAccountsPaginated retrieves a paginated list of accounts
func (s *Service) GetAccountsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Account], error) {
	return s.accountRes().ListPaginated(parameters)
}

// GetAccount retrieves a single account by id
func (s *Service) GetAccount(accountID string) (Account, error) {
	return s.accountRes().Get(accountID)
}

// CreateAccount creates a new account
func (s *Service) CreateAccount(req CreateAccountRequest) (string, error) {
	data, err := s.accountRes().Create(&req)
	return data.ID, err
}

// PatchAccount updates an account
func (s *Service) PatchAccount(accountID string, req PatchAccountRequest) error {
	if accountID == "" {
		return fmt.Errorf("invalid account id")
	}
	_, err := connection.Post[struct{}](s.connection, fmt.Sprintf("/cloudflare/v1/accounts/%s", accountID), &req)
	return err
}

// CreateAccount creates a new account member
func (s *Service) CreateAccountMember(accountID string, req CreateAccountMemberRequest) error {
	if accountID == "" {
		return fmt.Errorf("invalid account id")
	}
	_, err := connection.Post[struct{}](s.connection, fmt.Sprintf("/cloudflare/v1/accounts/%s/members", accountID), &req)
	return err
}
