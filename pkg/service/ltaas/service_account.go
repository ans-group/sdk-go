package ltaas

// CreateAccount creates a new account
func (s *Service) CreateAccount() (string, error) {
	body, err := s.createAccountResponseBody()

	return body.Data.ID, err
}

func (s *Service) createAccountResponseBody() (*GetAccountResponseBody, error) {
	body := &GetAccountResponseBody{}

	response, err := s.connection.Post("/ltaas/v1/accounts", nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
