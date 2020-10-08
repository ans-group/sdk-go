package loadbalancer

import "github.com/ukfast/sdk-go/pkg/connection"

// GetGroupSliceResponseBody represents an API response body containing []Group data
type GetGroupSliceResponseBody struct {
	connection.APIResponseBody
	Data []Group `json:"data"`
}

// GetGroupResponseBody represents an API response body containing Group data
type GetGroupResponseBody struct {
	connection.APIResponseBody
	Data Group `json:"data"`
}

// GetConfigurationSliceResponseBody represents an API response body containing []Configuration data
type GetConfigurationSliceResponseBody struct {
	connection.APIResponseBody
	Data []Configuration `json:"data"`
}

// GetConfigurationResponseBody represents an API response body containing Configuration data
type GetConfigurationResponseBody struct {
	connection.APIResponseBody
	Data Configuration `json:"data"`
}
