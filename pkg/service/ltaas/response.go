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
