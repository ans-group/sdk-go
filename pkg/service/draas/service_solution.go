package draas

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) solutionRes() *resource.Resource[Solution, string] {
	return resource.NewStringResource[Solution](s.connection, "/draas/v1/solutions", "solution",
		func(id string) error { return &SolutionNotFoundError{ID: id} })
}

// GetSolutions retrieves a list of solutions
func (s *Service) GetSolutions(parameters connection.APIRequestParameters) ([]Solution, error) {
	return s.solutionRes().List(parameters)
}

// GetSolutionsPaginated retrieves a paginated list of solutions
func (s *Service) GetSolutionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Solution], error) {
	return s.solutionRes().ListPaginated(parameters)
}

// GetSolution retrieves a single solution by id
func (s *Service) GetSolution(solutionID string) (Solution, error) {
	return s.solutionRes().Get(solutionID)
}

// PatchSolution patches a solution by ID
func (s *Service) PatchSolution(solutionID string, req PatchSolutionRequest) error {
	return s.solutionRes().Patch(solutionID, &req)
}

func (s *Service) solutionBackupResourceRes() *resource.SubResourceList[BackupResource, string] {
	return resource.NewStringSubResourceList[BackupResource](s.connection,
		func(solutionID string) string {
			return fmt.Sprintf("/draas/v1/solutions/%s/backup-resources", solutionID)
		},
		"solution", "id", func(solutionID string) error { return &SolutionNotFoundError{ID: solutionID} })
}

// GetSolutionBackupResources retrieves a collection of backup resources for specified solution
func (s *Service) GetSolutionBackupResources(solutionID string, parameters connection.APIRequestParameters) ([]BackupResource, error) {
	return s.solutionBackupResourceRes().List(solutionID, parameters)
}

// GetSolutionBackupResourcesPaginated retrieves a paginated list of solution backup resources
func (s *Service) GetSolutionBackupResourcesPaginated(solutionID string, parameters connection.APIRequestParameters) (*connection.Paginated[BackupResource], error) {
	return s.solutionBackupResourceRes().ListPaginated(solutionID, parameters)
}

// GetSolutionBackupService retrieves the backup service for the specified solution
func (s *Service) GetSolutionBackupService(solutionID string) (BackupService, error) {
	if solutionID == "" {
		return BackupService{}, fmt.Errorf("invalid solution id")
	}
	body, err := connection.Get[BackupService](s.connection, fmt.Sprintf("/draas/v1/solutions/%s/backup-service", solutionID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&SolutionNotFoundError{ID: solutionID}))
	return body.Data, err
}

// ResetSolutionBackupServiceCredentials resets the credentials for the solution backup service
func (s *Service) ResetSolutionBackupServiceCredentials(solutionID string, req ResetBackupServiceCredentialsRequest) error {
	if solutionID == "" {
		return fmt.Errorf("invalid solution id")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/draas/v1/solutions/%s/backup-service/reset-credentials", solutionID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&SolutionNotFoundError{ID: solutionID}))
}

func (s *Service) solutionFailoverPlanRes() *resource.SubResource[FailoverPlan, string, string] {
	return resource.NewStringStringSubResource[FailoverPlan](s.connection,
		func(solutionID string) string {
			return fmt.Sprintf("/draas/v1/solutions/%s/failover-plans", solutionID)
		},
		"solution", "id", func(solutionID string) error { return &SolutionNotFoundError{ID: solutionID} },
		"failover plan", "id", func(_, failoverPlanID string) error { return &FailoverPlanNotFoundError{ID: failoverPlanID} })
}

// GetSolutionFailoverPlans retrieves a collection of failover plans for specified solution
func (s *Service) GetSolutionFailoverPlans(solutionID string, parameters connection.APIRequestParameters) ([]FailoverPlan, error) {
	return s.solutionFailoverPlanRes().List(solutionID, parameters)
}

// GetSolutionFailoverPlansPaginated retrieves a paginated list of solution failover plans
func (s *Service) GetSolutionFailoverPlansPaginated(solutionID string, parameters connection.APIRequestParameters) (*connection.Paginated[FailoverPlan], error) {
	return s.solutionFailoverPlanRes().ListPaginated(solutionID, parameters)
}

// GetSolutionFailoverPlan retrieves a single solution failover plan by id
func (s *Service) GetSolutionFailoverPlan(solutionID string, failoverPlanID string) (FailoverPlan, error) {
	return s.solutionFailoverPlanRes().Get(solutionID, failoverPlanID)
}

// StartSolutionFailoverPlan starts the specified failover plan
func (s *Service) StartSolutionFailoverPlan(solutionID string, failoverPlanID string, req StartFailoverPlanRequest) error {
	if solutionID == "" {
		return fmt.Errorf("invalid solution id")
	}
	if failoverPlanID == "" {
		return fmt.Errorf("invalid failover plan id")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/draas/v1/solutions/%s/failover-plans/%s/start", solutionID, failoverPlanID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&FailoverPlanNotFoundError{ID: failoverPlanID}))
}

// StopSolutionFailoverPlan stops the specified failover plan
func (s *Service) StopSolutionFailoverPlan(solutionID string, failoverPlanID string) error {
	if solutionID == "" {
		return fmt.Errorf("invalid solution id")
	}
	if failoverPlanID == "" {
		return fmt.Errorf("invalid failover plan id")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/draas/v1/solutions/%s/failover-plans/%s/stop", solutionID, failoverPlanID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&FailoverPlanNotFoundError{ID: failoverPlanID}))
}

func (s *Service) solutionComputeResourceRes() *resource.SubResource[ComputeResource, string, string] {
	return resource.NewStringStringSubResource[ComputeResource](s.connection,
		func(solutionID string) string {
			return fmt.Sprintf("/draas/v1/solutions/%s/compute-resources", solutionID)
		},
		"solution", "id", func(solutionID string) error { return &SolutionNotFoundError{ID: solutionID} },
		"compute resource", "id", func(_, computeResourceID string) error {
			return &ComputeResourceNotFoundError{ID: computeResourceID}
		})
}

// GetSolutionComputeResources retrieves a collection of compute resources for specified solution
func (s *Service) GetSolutionComputeResources(solutionID string, parameters connection.APIRequestParameters) ([]ComputeResource, error) {
	return s.solutionComputeResourceRes().List(solutionID, parameters)
}

// GetSolutionComputeResourcesPaginated retrieves a paginated list of solution compute resources
func (s *Service) GetSolutionComputeResourcesPaginated(solutionID string, parameters connection.APIRequestParameters) (*connection.Paginated[ComputeResource], error) {
	return s.solutionComputeResourceRes().ListPaginated(solutionID, parameters)
}

// GetSolutionComputeResource retrieves compute resources by id
func (s *Service) GetSolutionComputeResource(solutionID string, computeResourceID string) (ComputeResource, error) {
	return s.solutionComputeResourceRes().Get(solutionID, computeResourceID)
}

func (s *Service) solutionHardwarePlanRes() *resource.SubResource[HardwarePlan, string, string] {
	return resource.NewStringStringSubResource[HardwarePlan](s.connection,
		func(solutionID string) string {
			return fmt.Sprintf("/draas/v1/solutions/%s/hardware-plans", solutionID)
		},
		"solution", "id", func(solutionID string) error { return &SolutionNotFoundError{ID: solutionID} },
		"hardware plan", "id", func(_, hardwarePlanID string) error { return &HardwarePlanNotFoundError{ID: hardwarePlanID} })
}

// GetSolutionHardwarePlans retrieves a collection of hardware plans for specified solution
func (s *Service) GetSolutionHardwarePlans(solutionID string, parameters connection.APIRequestParameters) ([]HardwarePlan, error) {
	return s.solutionHardwarePlanRes().List(solutionID, parameters)
}

// GetSolutionHardwarePlansPaginated retrieves a paginated list of solution hardware plans
func (s *Service) GetSolutionHardwarePlansPaginated(solutionID string, parameters connection.APIRequestParameters) (*connection.Paginated[HardwarePlan], error) {
	return s.solutionHardwarePlanRes().ListPaginated(solutionID, parameters)
}

// GetSolutionHardwarePlan retrieves hardware plans by id
func (s *Service) GetSolutionHardwarePlan(solutionID string, hardwarePlanID string) (HardwarePlan, error) {
	return s.solutionHardwarePlanRes().Get(solutionID, hardwarePlanID)
}

// GetSolutionHardwarePlanReplicas retrieves a collection of hardware plans for specified solution
func (s *Service) GetSolutionHardwarePlanReplicas(solutionID string, hardwarePlanID string, parameters connection.APIRequestParameters) ([]Replica, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Replica], error) {
		return s.GetSolutionHardwarePlanReplicasPaginated(solutionID, hardwarePlanID, p)
	}, parameters)
}

// GetSolutionHardwarePlanReplicasPaginated retrieves a paginated list of solution hardware plans
func (s *Service) GetSolutionHardwarePlanReplicasPaginated(solutionID string, hardwarePlanID string, parameters connection.APIRequestParameters) (*connection.Paginated[Replica], error) {
	if solutionID == "" {
		return nil, fmt.Errorf("invalid solution id")
	}
	if hardwarePlanID == "" {
		return nil, fmt.Errorf("invalid hardware plan id")
	}
	body, err := connection.Get[[]Replica](s.connection, fmt.Sprintf("/draas/v1/solutions/%s/hardware-plans/%s/replicas", solutionID, hardwarePlanID), parameters, connection.NotFoundResponseHandler(&SolutionNotFoundError{ID: solutionID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Replica], error) {
		return s.GetSolutionHardwarePlanReplicasPaginated(solutionID, hardwarePlanID, p)
	}), err
}

// UpdateSolutionReplicaIOPS updates a solution replica by ID
func (s *Service) UpdateSolutionReplicaIOPS(solutionID string, replicaID string, req UpdateReplicaIOPSRequest) error {
	if solutionID == "" {
		return fmt.Errorf("invalid solution id")
	}
	if replicaID == "" {
		return fmt.Errorf("invalid replica id")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/draas/v1/solutions/%s/replicas/%s/iops", solutionID, replicaID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&ReplicaNotFoundError{ID: replicaID}))
}
