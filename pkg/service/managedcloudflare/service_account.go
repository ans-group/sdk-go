package managedcloudflare

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetAccounts retrieves a list of accounts
func (s *Service) GetAccounts(parameters connection.APIRequestParameters) ([]Account, error) {
	var accounts []Account

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetAccountsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		accounts = append(accounts, response.(*PaginatedAccount).Items...)
	}

	return accounts, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetAccountsPaginated retrieves a paginated list of accounts
func (s *Service) GetAccountsPaginated(parameters connection.APIRequestParameters) (*PaginatedAccount, error) {
	body, err := s.getAccountsPaginatedResponseBody(parameters)

	return NewPaginatedAccount(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetAccountsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getAccountsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetAccountSliceResponseBody, error) {
	body := &GetAccountSliceResponseBody{}

	response, err := s.connection.Get("/managed-cloudflare/v2/accounts", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetAccount retrieves a single account by id
func (s *Service) GetAccount(accountID string) (Account, error) {
	body, err := s.getAccountResponseBody(accountID)

	return body.Data, err
}

func (s *Service) getAccountResponseBody(accountID string) (*GetAccountResponseBody, error) {
	body := &GetAccountResponseBody{}

	if accountID == "" {
		return body, fmt.Errorf("invalid account id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/managed-cloudflare/v2/accounts/%s", accountID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &AccountNotFoundError{ID: accountID}
		}

		return nil
	})
}

// CreateAccount creates a new account
func (s *Service) CreateAccount(req CreateAccountRequest) error {
	_, err := s.createAccountResponseBody(req)

	return err
}

func (s *Service) createAccountResponseBody(req CreateAccountRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	response, err := s.connection.Post("/managed-cloudflare/v2/accounts", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// CreateAccount creates a new account member
func (s *Service) CreateAccountMember(accountID string, req CreateAccountMemberRequest) error {
	_, err := s.createAccountMemberResponseBody(accountID, req)

	return err
}

func (s *Service) createAccountMemberResponseBody(accountID string, req CreateAccountMemberRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if accountID == "" {
		return body, fmt.Errorf("invalid account id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/managed-cloudflare/v2/accounts/%s/members", accountID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
