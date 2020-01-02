package ltaas

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetDomains retrieves a list of domains
func (s *Service) GetDomains(parameters connection.APIRequestParameters) ([]Domain, error) {
	var sites []Domain

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetDomainsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, site := range response.(*PaginatedDomain).Items {
			sites = append(sites, site)
		}
	}

	return sites, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetDomainsPaginated retrieves a paginated list of domains
func (s *Service) GetDomainsPaginated(parameters connection.APIRequestParameters) (*PaginatedDomain, error) {
	body, err := s.getDomainsPaginatedResponseBody(parameters)

	return NewPaginatedDomain(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetDomainsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getDomainsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetDomainsResponseBody, error) {
	body := &GetDomainsResponseBody{}

	response, err := s.connection.Get("/ltaas/v1/domains", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetDomain retrieves a single domain by id
func (s *Service) GetDomain(domainID string) (Domain, error) {
	body, err := s.getDomainResponseBody(domainID)

	return body.Data, err
}

func (s *Service) getDomainResponseBody(domainID string) (*GetDomainResponseBody, error) {
	body := &GetDomainResponseBody{}

	if domainID == "" {
		return body, fmt.Errorf("invalid domain id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ltaas/v1/domains/%s", domainID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{ID: domainID}
		}

		return nil
	})
}

// CreateDomain creates a new domain
func (s *Service) CreateDomain(req CreateDomainRequest) (string, error) {
	body, err := s.createDomainResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createDomainResponseBody(req CreateDomainRequest) (*GetDomainResponseBody, error) {
	body := &GetDomainResponseBody{}

	response, err := s.connection.Post("/ltaas/v1/domains", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// DeleteDomain removes a domain
func (s *Service) DeleteDomain(domainID string) error {
	_, err := s.deleteDomainResponseBody(domainID)

	return err
}

func (s *Service) deleteDomainResponseBody(domainID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainID == "" {
		return body, fmt.Errorf("invalid domain id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ltaas/v1/domains/%s", domainID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{ID: domainID}
		}

		return nil
	})
}

// VerifyDomainByFile verifies a domain by File method
func (s *Service) VerifyDomainByFile(domainID string) error {
	_, err := s.verifyDomainByFileResponseBody(domainID)

	return err
}

func (s *Service) verifyDomainByFileResponseBody(domainID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainID == "" {
		return body, fmt.Errorf("invalid domain id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ltaas/v1/domains/%s/verify-by-file", domainID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{ID: domainID}
		}

		return nil
	})
}

// VerifyDomainByDNS verifies a domain by DNS method
func (s *Service) VerifyDomainByDNS(domainID string) error {
	_, err := s.verifyDomainByDNSResponseBody(domainID)

	return err
}

func (s *Service) verifyDomainByDNSResponseBody(domainID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainID == "" {
		return body, fmt.Errorf("invalid domain id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ltaas/v1/domains/%s/verify-by-dns", domainID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{ID: domainID}
		}

		return nil
	})
}
