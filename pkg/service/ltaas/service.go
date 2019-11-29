package ltaas

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// LTaaSService is an interface for managing LTaaS
type LTaaSService interface {
	GetDomains(parameters connection.APIRequestParameters) ([]Domain, error)
	GetDomainsPaginated(parameters connection.APIRequestParameters) (*PaginatedDomain, error)
	GetDomain(domainID int) (Domain, error)
}

// Service implements LTaaSService for managing
// LTaaS via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of LTaaSService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
