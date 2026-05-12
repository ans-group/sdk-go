package cloudflare

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetAccounts retrieves a list of accounts
func (s *Service) GetAccounts(parameters connection.APIRequestParameters) ([]Account, error) {
	return connection.InvokeRequestAll(s.GetAccountsPaginated, parameters)
}

// GetAccountsPaginated retrieves a paginated list of accounts
func (s *Service) GetAccountsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Account], error) {
	body, err := connection.Get[[]Account](s.connection, "/cloudflare/v1/accounts", parameters)
	return connection.NewPaginated(body, parameters, s.GetAccountsPaginated), err
}

// GetAccount retrieves a single account by id
func (s *Service) GetAccount(accountID string) (Account, error) {
	if accountID == "" {
		return Account{}, fmt.Errorf("invalid account id")
	}
	body, err := connection.Get[Account](s.connection, fmt.Sprintf("/cloudflare/v1/accounts/%s", accountID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&AccountNotFoundError{ID: accountID}))
	return body.Data, err
}

// CreateAccount creates a new account
func (s *Service) CreateAccount(req CreateAccountRequest) (string, error) {
	body, err := connection.Post[Account](s.connection, "/cloudflare/v1/accounts", &req)
	return body.Data.ID, err
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
