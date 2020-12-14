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

// GetTargetGroupSliceResponseBody represents an API response body containing []TargetGroup data
type GetTargetGroupSliceResponseBody struct {
	connection.APIResponseBody
	Data []TargetGroup `json:"data"`
}

// GetTargetGroupResponseBody represents an API response body containing TargetGroup data
type GetTargetGroupResponseBody struct {
	connection.APIResponseBody
	Data TargetGroup `json:"data"`
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

// GetVIPSliceResponseBody represents an API response body containing []VIP data
type GetVIPSliceResponseBody struct {
	connection.APIResponseBody
	Data []VIP `json:"data"`
}

// GetVIPResponseBody represents an API response body containing VIP data
type GetVIPResponseBody struct {
	connection.APIResponseBody
	Data VIP `json:"data"`
}

// GetListenerSliceResponseBody represents an API response body containing []Listener data
type GetListenerSliceResponseBody struct {
	connection.APIResponseBody
	Data []Listener `json:"data"`
}

// GetListenerResponseBody represents an API response body containing Listener data
type GetListenerResponseBody struct {
	connection.APIResponseBody
	Data Listener `json:"data"`
}

// GetErrorPageSliceResponseBody represents an API response body containing []ErrorPage data
type GetErrorPageSliceResponseBody struct {
	connection.APIResponseBody
	Data []ErrorPage `json:"data"`
}

// GetErrorPageResponseBody represents an API response body containing ErrorPage data
type GetErrorPageResponseBody struct {
	connection.APIResponseBody
	Data ErrorPage `json:"data"`
}

// GetAccessSliceResponseBody represents an API response body containing []Access data
type GetAccessSliceResponseBody struct {
	connection.APIResponseBody
	Data []Access `json:"data"`
}

// GetAccessResponseBody represents an API response body containing Access data
type GetAccessResponseBody struct {
	connection.APIResponseBody
	Data Access `json:"data"`
}

// GetBindSliceResponseBody represents an API response body containing []Bind data
type GetBindSliceResponseBody struct {
	connection.APIResponseBody
	Data []Bind `json:"data"`
}

// GetBindResponseBody represents an API response body containing Bind data
type GetBindResponseBody struct {
	connection.APIResponseBody
	Data Bind `json:"data"`
}

// GetListenerCertificateSliceResponseBody represents an API response body containing []ListenerCertificate data
type GetListenerCertificateSliceResponseBody struct {
	connection.APIResponseBody
	Data []ListenerCertificate `json:"data"`
}

// GetListenerCertificateResponseBody represents an API response body containing ListenerCertificate data
type GetListenerCertificateResponseBody struct {
	connection.APIResponseBody
	Data ListenerCertificate `json:"data"`
}

// GetListenerErrorPageSliceResponseBody represents an API response body containing []ListenerErrorPage data
type GetListenerErrorPageSliceResponseBody struct {
	connection.APIResponseBody
	Data []ListenerErrorPage `json:"data"`
}

// GetListenerErrorPageResponseBody represents an API response body containing ListenerErrorPage data
type GetListenerErrorPageResponseBody struct {
	connection.APIResponseBody
	Data ListenerErrorPage `json:"data"`
}

// GetSSLSliceResponseBody represents an API response body containing []SSL data
type GetSSLSliceResponseBody struct {
	connection.APIResponseBody
	Data []SSL `json:"data"`
}

// GetSSLResponseBody represents an API response body containing SSL data
type GetSSLResponseBody struct {
	connection.APIResponseBody
	Data SSL `json:"data"`
}

// GetCustomOptionSliceResponseBody represents an API response body containing []CustomOption data
type GetCustomOptionSliceResponseBody struct {
	connection.APIResponseBody
	Data []CustomOption `json:"data"`
}

// GetCustomOptionResponseBody represents an API response body containing CustomOption data
type GetCustomOptionResponseBody struct {
	connection.APIResponseBody
	Data CustomOption `json:"data"`
}
