package ltaas

import "github.com/ukfast/sdk-go/pkg/connection"

// GetDomainsResponseBody represents the API response body from the GetDomains resource
type GetDomainsResponseBody struct {
	connection.APIResponseBody

	Data []Domain `json:"data"`
}

// GetDomainResponseBody represents the API response body from the GetDomain resource
type GetDomainResponseBody struct {
	connection.APIResponseBody

	Data Domain `json:"data"`
}

// GetTestsResponseBody represents the API response body from the GetTests resource
type GetTestsResponseBody struct {
	connection.APIResponseBody

	Data []Test `json:"data"`
}

// GetTestResponseBody represents the API response body from the GetTest resource
type GetTestResponseBody struct {
	connection.APIResponseBody

	Data Test `json:"data"`
}

// GetJobsResponseBody represents the API response body from the GetJobs resource
type GetJobsResponseBody struct {
	connection.APIResponseBody

	Data []Job `json:"data"`
}

// GetJobResponseBody represents the API response body from the GetJob resource
type GetJobResponseBody struct {
	connection.APIResponseBody

	Data Job `json:"data"`
}

// GetJobResultsResponseBody represents the API response body from the GetJobResults resource
type GetJobResultsResponseBody struct {
	connection.APIResponseBody

	Data JobResults `json:"data"`
}

// GetJobSettingsResponseBody represents the API response body from the GetJobSettings resource
type GetJobSettingsResponseBody struct {
	connection.APIResponseBody

	Data JobSettings `json:"data"`
}

// GetThresholdsResponseBody represents the API response body from the GetThresholds resource
type GetThresholdsResponseBody struct {
	connection.APIResponseBody

	Data []Threshold `json:"data"`
}

// GetThresholdResponseBody represents the API response body from the GetThreshold resource
type GetThresholdResponseBody struct {
	connection.APIResponseBody

	Data Threshold `json:"data"`
}

// GetScenariosResponseBody represents the API response body from the GetScenarios resource
type GetScenariosResponseBody struct {
	connection.APIResponseBody

	Data []Scenario `json:"data"`
}

// GetAgreementResponseBody represents the API response body from the GetAgreement resource
type GetAgreementResponseBody struct {
	connection.APIResponseBody

	Data Agreement `json:"data"`
}

// GetAccountResponseBody represents the API response body from the GetAccount resource
type GetAccountResponseBody struct {
	connection.APIResponseBody

	Data Account `json:"data"`
}
