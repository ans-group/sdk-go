package registrar

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) domainRes() *resource.Resource[Domain, string] {
	return resource.NewStringResourceWithIdentifier[Domain](s.connection, "/registrar/v1/domains", "domain", "name",
		func(id string) error { return &DomainNotFoundError{Name: id} })
}

// GetDomains retrieves a list of domains
func (s *Service) GetDomains(parameters connection.APIRequestParameters) ([]Domain, error) {
	return s.domainRes().List(parameters)
}

// GetDomainsPaginated retrieves a paginated list of domains
func (s *Service) GetDomainsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Domain], error) {
	return s.domainRes().ListPaginated(parameters)
}

// GetDomain retrieves a single domain by name
func (s *Service) GetDomain(domainName string) (Domain, error) {
	return s.domainRes().Get(domainName)
}

// GetDomainNameservers retrieves the nameservers for a domain
func (s *Service) GetDomainNameservers(domainName string) ([]Nameserver, error) {
	if domainName == "" {
		return []Nameserver{}, fmt.Errorf("invalid domain name")
	}
	body, err := connection.Get[[]Nameserver](s.connection, fmt.Sprintf("/registrar/v1/domains/%s/nameservers", domainName), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
	return body.Data, err
}
