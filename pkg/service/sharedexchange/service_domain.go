package sharedexchange

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
	body, err := connection.Get[[]Domain](s.connection, "/shared-exchange/v1/domains", parameters)
	return connection.NewPaginated(body, parameters, s.GetDomainsPaginated), err
}

// GetDomain retrieves a single domain by id
func (s *Service) GetDomain(domainID int) (Domain, error) {
	if domainID < 1 {
		return Domain{}, fmt.Errorf("invalid domain id")
	}
	body, err := connection.Get[Domain](s.connection, fmt.Sprintf("/shared-exchange/v1/domains/%d", domainID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&DomainNotFoundError{ID: domainID}))
	return body.Data, err
}
