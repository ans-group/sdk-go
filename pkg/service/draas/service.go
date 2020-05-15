package draas

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// DRaaSService is an interface for managing the DRaaS service
type DRaaSService interface {
	GetSolutions(parameters connection.APIRequestParameters) ([]Solution, error)
	GetSolutionsPaginated(parameters connection.APIRequestParameters) (*PaginatedSolution, error)
	GetSolution(solutionID string) (Solution, error)
	PatchSolution(solutionID string, req PatchSolutionRequest) error

	GetSolutionBackupResources(parameters connection.APIRequestParameters, solutionID string) ([]BackupResource, error)
	GetSolutionBackupResourcesPaginated(parameters connection.APIRequestParameters, solutionID string) (*PaginatedBackupResource, error)

	GetSolutionBackupService(solutionID string) (BackupService, error)
	ResetSolutionBackupServiceCredentials(solutionID string, req ResetBackupServiceCredentialsRequest) error

	GetSolutionFailoverPlans(parameters connection.APIRequestParameters, solutionID string) ([]FailoverPlan, error)
	GetSolutionFailoverPlansPaginated(parameters connection.APIRequestParameters, solutionID string) (*PaginatedFailoverPlan, error)
	GetSolutionFailoverPlan(solutionID string, failoverPlanID string) (FailoverPlan, error)
	StartSolutionFailoverPlan(solutionID string, failoverPlanID string) error
	StopSolutionFailoverPlan(solutionID string, failoverPlanID string) error

	GetSolutionComputeResources(parameters connection.APIRequestParameters, solutionID string) ([]ComputeResource, error)
	GetSolutionComputeResourcesPaginated(parameters connection.APIRequestParameters, solutionID string) (*PaginatedComputeResource, error)
	GetSolutionComputeResource(solutionID string, computeResourcesID string) (ComputeResource, error)

	GetSolutionHardwarePlans(parameters connection.APIRequestParameters, solutionID string) ([]HardwarePlan, error)
	GetSolutionHardwarePlansPaginated(parameters connection.APIRequestParameters, solutionID string) (*PaginatedHardwarePlan, error)
	GetSolutionHardwarePlan(solutionID string, hardwarePlanID string) (HardwarePlan, error)
	GetSolutionHardwarePlanReplicas(parameters connection.APIRequestParameters, solutionID string, hardwarePlanID string) ([]Replica, error)

	GetIOPSTiers(parameters connection.APIRequestParameters) ([]IOPSTier, error)
	GetIOPSTier(iopsTierID string) (IOPSTier, error)

	GetBillingTypes(parameters connection.APIRequestParameters) ([]BillingType, error)
	GetBillingTypesPaginated(parameters connection.APIRequestParameters) (*PaginatedBillingType, error)
	GetBillingType(billingTypeID string) (BillingType, error)
}

// Service implements DRaaSService for managing
// DRaaS certificates via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of DRaaSService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
