package registrar

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetWhois retrieves WHOIS information for a single domain
func (s *Service) GetWhois(domainName string) (Whois, error) {
	if domainName == "" {
		return Whois{}, fmt.Errorf("invalid domain name")
	}
	body, err := connection.Get[Whois](s.connection, fmt.Sprintf("/registrar/v1/whois/%s", domainName), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
	return body.Data, err
}

// GetWhoisRaw retrieves raw WHOIS information for a single domain
func (s *Service) GetWhoisRaw(domainName string) (string, error) {
	if domainName == "" {
		return "", fmt.Errorf("invalid domain name")
	}
	body, err := connection.Get[string](s.connection, fmt.Sprintf("/registrar/v1/whois/%s/raw", domainName), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
	return body.Data, err
}
