package loadbalancer

import "github.com/ukfast/sdk-go/pkg/connection"

// GetTargetSliceResponseBody represents an API response body containing []Target data
type GetTargetSliceResponseBody struct {
	connection.APIResponseBody
	Data []Target `json:"data"`
}

// GetTargetResponseBody represents an API response body containing Target data
type GetTargetResponseBody struct {
	connection.APIResponseBody
	Data Target `json:"data"`
}

// GetClusterSliceResponseBody represents an API response body containing []Cluster data
type GetClusterSliceResponseBody struct {
	connection.APIResponseBody
	Data []Cluster `json:"data"`
}

// GetClusterResponseBody represents an API response body containing Cluster data
type GetClusterResponseBody struct {
	connection.APIResponseBody
	Data Cluster `json:"data"`
}
