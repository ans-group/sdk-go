package ltaas

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// LTaaSService is an interface for managing LTaaS
type LTaaSService interface {
	GetDomains(parameters connection.APIRequestParameters) ([]Domain, error)
	GetDomainsPaginated(parameters connection.APIRequestParameters) (*PaginatedDomain, error)
	GetDomain(domainID string) (Domain, error)
	CreateDomain(req CreateDomainRequest) (string, error)
	DeleteDomain(domainID string) error
	VerifyDomainFile(domainID string) error
	VerifyDomainDNS(domainID string) error

	GetTests(parameters connection.APIRequestParameters) ([]Test, error)
	GetTestsPaginated(parameters connection.APIRequestParameters) (*PaginatedTest, error)
	GetTest(testID string) (Test, error)
	CreateTest(req CreateTestRequest) (string, error)
	DeleteTest(testID string) error
	CreateTestJob(testID string, req CreateTestJobRequest) (string, error)

	GetJobs(parameters connection.APIRequestParameters) ([]Job, error)
	GetJobsPaginated(parameters connection.APIRequestParameters) (*PaginatedJob, error)
	GetJob(testID string) (Job, error)
	GetJobResults(jobID string) (JobResults, error)
	GetJobSettings(jobID string) (JobSettings, error)
	CreateJob(req CreateJobRequest) (string, error)
	DeleteJob(jobID string) error
	StopJob(jobID string) error

	GetThresholds(parameters connection.APIRequestParameters) ([]Threshold, error)
	GetThresholdsPaginated(parameters connection.APIRequestParameters) (*PaginatedThreshold, error)
	GetThreshold(thresholdID string) (Threshold, error)

	GetScenarios(parameters connection.APIRequestParameters) ([]Scenario, error)
	GetScenariosPaginated(parameters connection.APIRequestParameters) (*PaginatedScenario, error)

	GetLatestAgreement(agreementType AgreementType) (Agreement, error)
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
