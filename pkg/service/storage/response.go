package storage

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

// GetVolumesResponseBody represents the API response body from the GetVolumes resource
type GetVolumesResponseBody struct {
	connection.APIResponseBody

	Data []Volume `json:"data"`
}

// GetVolumeResponseBody represents the API response body from the GetVolume resource
type GetVolumeResponseBody struct {
	connection.APIResponseBody

	Data Volume `json:"data"`
}

// GetHostsResponseBody represents the API response body from the GetHosts resource
type GetHostsResponseBody struct {
	connection.APIResponseBody

	Data []Host `json:"data"`
}

// GetHostResponseBody represents the API response body from the GetHost resource
type GetHostResponseBody struct {
	connection.APIResponseBody

	Data Host `json:"data"`
}
