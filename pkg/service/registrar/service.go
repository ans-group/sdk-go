package registrar

import "github.com/ukfast/sdk-go/pkg/connection"

// RegistrarService is an interface for managing the registrar service
type RegistrarService interface {
	GetDomains(parameters connection.APIRequestParameters) ([]Domain, error)
	GetDomainsPaginated(parameters connection.APIRequestParameters) ([]Domain, error)
	GetDomain(domainName string) (Domain, error)
	GetDomainNameservers(domainName string) ([]Nameserver, error)
}

// Service implements RegistrarService for managing
// registrar services via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of Service
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
