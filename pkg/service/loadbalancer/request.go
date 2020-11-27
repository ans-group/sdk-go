package loadbalancer

import "github.com/ukfast/sdk-go/pkg/connection"

// PatchClusterRequest represents a request to patch a cluster
type PatchClusterRequest struct {
	Name *string `json:"name,omitempty"`
}

// CreateTargetRequest represents a request to create a target
type CreateTargetRequest struct {
	TargetGroupID string               `json:"targetgroup_id"`
	IP            connection.IPAddress `json:"ip"`
	Port          int                  `json:"port"`
	Weight        int                  `json:"weight"`
	Backup        bool                 `json:"backup"`
	CheckInterval int                  `json:"check_interval"`
	CheckSSL      bool                 `json:"check_ssl"`
	CheckRise     int                  `json:"check_rise"`
	CheckFall     int                  `json:"check_fall"`
	DisableHTTP2  bool                 `json:"disable_http2"`
	HTTP2Only     bool                 `json:"http2_only"`
}

// PatchTargetRequest represents a request to patch a target
type PatchTargetRequest struct {
	TargetGroupID string               `json:"targetgroup_id,omitempty"`
	IP            connection.IPAddress `json:"ip,omitempty"`
	Port          int                  `json:"port,omitempty"`
	Weight        int                  `json:"weight,omitempty"`
	Backup        *bool                `json:"backup,omitempty"`
	CheckInterval int                  `json:"check_interval,omitempty"`
	CheckSSL      *bool                `json:"check_ssl,omitempty"`
	CheckRise     int                  `json:"check_rise,omitempty"`
	CheckFall     int                  `json:"check_fall,omitempty"`
	DisableHTTP2  *bool                `json:"disable_http2,omitempty"`
	HTTP2Only     *bool                `json:"http2_only,omitempty"`
}
