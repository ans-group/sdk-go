package draas

import "github.com/ukfast/sdk-go/pkg/connection"

// GetSolutionsResponseBody represents the API response body from the GetSolutions resource
type GetSolutionsResponseBody struct {
	connection.APIResponseBody

	Data []Solution `json:"data"`
}

// GetSolutionResponseBody represents the API response body from the GetSolution resource
type GetSolutionResponseBody struct {
	connection.APIResponseBody

	Data Solution `json:"data"`
}

// GetBackupResourcesResponseBody represents the API response body from the GetBackupResources resource
type GetBackupResourcesResponseBody struct {
	connection.APIResponseBody

	Data []BackupResource `json:"data"`
}

// GetIOPSTiersResponseBody represents the API response body from the GetIOPSTiers resource
type GetIOPSTiersResponseBody struct {
	connection.APIResponseBody

	Data []IOPSTier `json:"data"`
}

// GetIOPSTierResponseBody represents the API response body from the GetIOPSTier resource
type GetIOPSTierResponseBody struct {
	connection.APIResponseBody

	Data IOPSTier `json:"data"`
}

// GetBackupServiceResponseBody represents the API response body from the GetBackupService resource
type GetBackupServiceResponseBody struct {
	connection.APIResponseBody

	Data BackupService `json:"data"`
}

// GetFailoverPlansResponseBody represents the API response body from the GetFailoverPlanss resource
type GetFailoverPlansResponseBody struct {
	connection.APIResponseBody

	Data []FailoverPlan `json:"data"`
}

// GetFailoverPlanResponseBody represents the API response body from the GetFailoverPlans resource
type GetFailoverPlanResponseBody struct {
	connection.APIResponseBody

	Data FailoverPlan `json:"data"`
}

// GetComputeResourcesResponseBody represents the API response body from the GetComputeResources resource
type GetComputeResourcesResponseBody struct {
	connection.APIResponseBody

	Data []ComputeResource `json:"data"`
}

// GetComputeResourceResponseBody represents the API response body from the GetComputeResource resource
type GetComputeResourceResponseBody struct {
	connection.APIResponseBody

	Data ComputeResource `json:"data"`
}

// GetHardwarePlansResponseBody represents the API response body from the GetHardwarePlans resource
type GetHardwarePlansResponseBody struct {
	connection.APIResponseBody

	Data []HardwarePlan `json:"data"`
}

// GetHardwarePlanResponseBody represents the API response body from the GetHardwarePlan resource
type GetHardwarePlanResponseBody struct {
	connection.APIResponseBody

	Data HardwarePlan `json:"data"`
}

// GetReplicasResponseBody represents the API response body from the GetReplicas resource
type GetReplicasResponseBody struct {
	connection.APIResponseBody

	Data []Replica `json:"data"`
}

// GetBillingTypesResponseBody represents the API response body from the GetBillingTypes resource
type GetBillingTypesResponseBody struct {
	connection.APIResponseBody

	Data []BillingType `json:"data"`
}

// GetBillingTypeResponseBody represents the API response body from the GetBillingType resource
type GetBillingTypeResponseBody struct {
	connection.APIResponseBody

	Data BillingType `json:"data"`
}
