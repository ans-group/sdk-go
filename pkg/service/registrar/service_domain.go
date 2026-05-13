package registrar

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetDomains retrieves a list of domains
func (s *Service) GetDomains(parameters connection.APIRequestParameters) ([]Domain, error) {
	return connection.InvokeRequestAll(s.GetDomainsPaginated, parameters)
}

// GetDomainsPaginated retrieves a paginated list of domains
func (s *Service) GetDomainsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Domain], error) {
	body, err := connection.Get[[]Domain](s.connection, "/registrar/v1/domains", parameters)
	return connection.NewPaginated(body, parameters, s.GetDomainsPaginated), err
}

// GetDomain retrieves a single domain by name
func (s *Service) GetDomain(domainName string) (Domain, error) {
	if domainName == "" {
		return Domain{}, fmt.Errorf("invalid domain name")
	}
	body, err := connection.Get[Domain](s.connection, fmt.Sprintf("/registrar/v1/domains/%s", domainName), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
	return body.Data, err
}

// GetDomainNameservers retrieves the nameservers for a domain
func (s *Service) GetDomainNameservers(domainName string) ([]Nameserver, error) {
	if domainName == "" {
		return []Nameserver{}, fmt.Errorf("invalid domain name")
	}
	body, err := connection.Get[[]Nameserver](s.connection, fmt.Sprintf("/registrar/v1/domains/%s/nameservers", domainName), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
	return body.Data, err
}
