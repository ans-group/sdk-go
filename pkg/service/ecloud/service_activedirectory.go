package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetActiveDirectoryDomains retrieves a list of Active Directory Domains
func (s *Service) GetActiveDirectoryDomains(parameters connection.APIRequestParameters) ([]ActiveDirectoryDomain, error) {
	return connection.InvokeRequestAll(s.GetActiveDirectoryDomainsPaginated, parameters)
}

// GetActiveDirectoryDomainsPaginated retrieves a paginated list of Active Directory Domains
func (s *Service) GetActiveDirectoryDomainsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ActiveDirectoryDomain], error) {
	body, err := connection.Get[[]ActiveDirectoryDomain](s.connection, "/ecloud/v1/active-directory/domains", parameters)
	return connection.NewPaginated(body, parameters, s.GetActiveDirectoryDomainsPaginated), err
}

// GetActiveDirectoryDomain retrieves a single domain by ID
func (s *Service) GetActiveDirectoryDomain(domainID int) (ActiveDirectoryDomain, error) {
	if domainID < 1 {
		return ActiveDirectoryDomain{}, fmt.Errorf("invalid domain id")
	}
	body, err := connection.Get[ActiveDirectoryDomain](s.connection, fmt.Sprintf("/ecloud/v1/active-directory/domains/%d", domainID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ActiveDirectoryDomainNotFoundError{ID: domainID}))
	return body.Data, err
}
