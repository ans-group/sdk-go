package ecloud

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) activeDirectoryDomainRes() *resource.Resource[ActiveDirectoryDomain, int] {
	return resource.NewIntResource[ActiveDirectoryDomain](s.connection, "/ecloud/v1/active-directory/domains", "domain", func(id int) error {
		return &ActiveDirectoryDomainNotFoundError{ID: id}
	})
}

// GetActiveDirectoryDomains retrieves a list of Active Directory Domains
func (s *Service) GetActiveDirectoryDomains(parameters connection.APIRequestParameters) ([]ActiveDirectoryDomain, error) {
	return s.activeDirectoryDomainRes().List(parameters)
}

// GetActiveDirectoryDomainsPaginated retrieves a paginated list of Active Directory Domains
func (s *Service) GetActiveDirectoryDomainsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ActiveDirectoryDomain], error) {
	return s.activeDirectoryDomainRes().ListPaginated(parameters)
}

// GetActiveDirectoryDomain retrieves a single domain by ID
func (s *Service) GetActiveDirectoryDomain(domainID int) (ActiveDirectoryDomain, error) {
	return s.activeDirectoryDomainRes().Get(domainID)
}
