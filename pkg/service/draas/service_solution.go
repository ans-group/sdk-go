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

// GetSolutionBackupResources retrieves a collection of backup resources for specified solution
func (s *Service) GetSolutionBackupResources(solutionID string, parameters connection.APIRequestParameters) ([]BackupResource, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[BackupResource], error) {
		return s.GetSolutionBackupResourcesPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionsPaginated retrieves a paginated list of solutions
func (s *Service) GetSolutionBackupResourcesPaginated(solutionID string, parameters connection.APIRequestParameters) (*connection.Paginated[BackupResource], error) {
	if solutionID == "" {
		return nil, fmt.Errorf("invalid solution id")
	}
	body, err := connection.Get[[]BackupResource](s.connection, fmt.Sprintf("/draas/v1/solutions/%s/backup-resources", solutionID), parameters, connection.NotFoundResponseHandler(&SolutionNotFoundError{ID: solutionID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[BackupResource], error) {
		return s.GetSolutionBackupResourcesPaginated(solutionID, p)
	}), err
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

// GetSolutionFailoverPlans retrieves a collection of failover plans for specified solution
func (s *Service) GetSolutionFailoverPlans(solutionID string, parameters connection.APIRequestParameters) ([]FailoverPlan, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[FailoverPlan], error) {
		return s.GetSolutionFailoverPlansPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionsPaginated retrieves a paginated list of solution failover plans
func (s *Service) GetSolutionFailoverPlansPaginated(solutionID string, parameters connection.APIRequestParameters) (*connection.Paginated[FailoverPlan], error) {
	if solutionID == "" {
		return nil, fmt.Errorf("invalid solution id")
	}
	body, err := connection.Get[[]FailoverPlan](s.connection, fmt.Sprintf("/draas/v1/solutions/%s/failover-plans", solutionID), parameters, connection.NotFoundResponseHandler(&SolutionNotFoundError{ID: solutionID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[FailoverPlan], error) {
		return s.GetSolutionFailoverPlansPaginated(solutionID, p)
	}), err
}

// GetSolutionFailoverPlan retrieves a single solution failover plan by id
func (s *Service) GetSolutionFailoverPlan(solutionID string, failoverPlanID string) (FailoverPlan, error) {
	if solutionID == "" {
		return FailoverPlan{}, fmt.Errorf("invalid solution id")
	}
	if failoverPlanID == "" {
		return FailoverPlan{}, fmt.Errorf("invalid failover plan id")
	}
	body, err := connection.Get[FailoverPlan](s.connection, fmt.Sprintf("/draas/v1/solutions/%s/failover-plans/%s", solutionID, failoverPlanID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&FailoverPlanNotFoundError{ID: failoverPlanID}))
	return body.Data, err
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

// GetSolutionComputeResources retrieves a collection of compute resources for specified solution
func (s *Service) GetSolutionComputeResources(solutionID string, parameters connection.APIRequestParameters) ([]ComputeResource, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[ComputeResource], error) {
		return s.GetSolutionComputeResourcesPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionComputeResourcesPaginated retrieves a paginated list of solution compute resources
func (s *Service) GetSolutionComputeResourcesPaginated(solutionID string, parameters connection.APIRequestParameters) (*connection.Paginated[ComputeResource], error) {
	if solutionID == "" {
		return nil, fmt.Errorf("invalid solution id")
	}
	body, err := connection.Get[[]ComputeResource](s.connection, fmt.Sprintf("/draas/v1/solutions/%s/compute-resources", solutionID), parameters, connection.NotFoundResponseHandler(&SolutionNotFoundError{ID: solutionID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[ComputeResource], error) {
		return s.GetSolutionComputeResourcesPaginated(solutionID, p)
	}), err
}

// GetSolutionComputeResource retrieves compute resources by id
func (s *Service) GetSolutionComputeResource(solutionID string, computeResourceID string) (ComputeResource, error) {
	if solutionID == "" {
		return ComputeResource{}, fmt.Errorf("invalid solution id")
	}
	if computeResourceID == "" {
		return ComputeResource{}, fmt.Errorf("invalid compute resource id")
	}
	body, err := connection.Get[ComputeResource](s.connection, fmt.Sprintf("/draas/v1/solutions/%s/compute-resources/%s", solutionID, computeResourceID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ComputeResourceNotFoundError{ID: computeResourceID}))
	return body.Data, err
}

// GetSolutionHardwarePlans retrieves a collection of hardware plans for specified solution
func (s *Service) GetSolutionHardwarePlans(solutionID string, parameters connection.APIRequestParameters) ([]HardwarePlan, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[HardwarePlan], error) {
		return s.GetSolutionHardwarePlansPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionHardwarePlansPaginated retrieves a paginated list of solution hardware plans
func (s *Service) GetSolutionHardwarePlansPaginated(solutionID string, parameters connection.APIRequestParameters) (*connection.Paginated[HardwarePlan], error) {
	if solutionID == "" {
		return nil, fmt.Errorf("invalid solution id")
	}
	body, err := connection.Get[[]HardwarePlan](s.connection, fmt.Sprintf("/draas/v1/solutions/%s/hardware-plans", solutionID), parameters, connection.NotFoundResponseHandler(&SolutionNotFoundError{ID: solutionID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[HardwarePlan], error) {
		return s.GetSolutionHardwarePlansPaginated(solutionID, p)
	}), err
}

// GetSolutionHardwarePlan retrieves hardware plans by id
func (s *Service) GetSolutionHardwarePlan(solutionID string, hardwarePlanID string) (HardwarePlan, error) {
	if solutionID == "" {
		return HardwarePlan{}, fmt.Errorf("invalid solution id")
	}
	if hardwarePlanID == "" {
		return HardwarePlan{}, fmt.Errorf("invalid hardware plan id")
	}
	body, err := connection.Get[HardwarePlan](s.connection, fmt.Sprintf("/draas/v1/solutions/%s/hardware-plans/%s", solutionID, hardwarePlanID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&HardwarePlanNotFoundError{ID: hardwarePlanID}))
	return body.Data, err
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
