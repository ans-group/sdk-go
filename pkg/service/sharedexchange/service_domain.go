package sharedexchange

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) domainRes() *resource.Resource[Domain, int] {
	return resource.NewIntResource[Domain](s.connection, "/shared-exchange/v1/domains", "domain",
		func(id int) error { return &DomainNotFoundError{ID: id} })
}

// GetDomains retrieves a list of domains
func (s *Service) GetDomains(parameters connection.APIRequestParameters) ([]Domain, error) {
	return s.domainRes().List(parameters)
}

// GetDomainsPaginated retrieves a paginated list of domains
func (s *Service) GetDomainsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Domain], error) {
	return s.domainRes().ListPaginated(parameters)
}

// GetDomain retrieves a single domain by id
func (s *Service) GetDomain(domainID int) (Domain, error) {
	return s.domainRes().Get(domainID)
}
