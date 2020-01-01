package ltaas

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// LTaaSService is an interface for managing LTaaS
type LTaaSService interface {
	CreateDomain(req CreateDomainRequest) (string, error)
	GetDomains(parameters connection.APIRequestParameters) ([]Domain, error)
	GetDomainsPaginated(parameters connection.APIRequestParameters) (*PaginatedDomain, error)
	GetDomain(domainID string) (Domain, error)
	DeleteDomain(domainID string) error
	VerifyDomainByFile(domainID string) error
	VerifyDomainByDNS(domainID string) error

	GetTests(parameters connection.APIRequestParameters) ([]Test, error)
	GetTestsPaginated(parameters connection.APIRequestParameters) (*PaginatedTest, error)
	GetTest(testID string) (Test, error)

	GetJobs(parameters connection.APIRequestParameters) ([]Job, error)
	GetJobsPaginated(parameters connection.APIRequestParameters) (*PaginatedJob, error)
	GetJob(testID string) (Job, error)
	GetJobResults(jobID string) (JobResults, error)
	CreateJob(req CreateJobRequest) (string, error)
	DeleteJob(jobID string) error
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
