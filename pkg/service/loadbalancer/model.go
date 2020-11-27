//go:generate go run ../../gen/model_response/main.go -package loadbalancer -source model.go -destination model_response_generated.go
//go:generate go run ../../gen/model_paginated/main.go -package loadbalancer -source model.go -destination model_paginated_generated.go

package loadbalancer

import "github.com/ukfast/sdk-go/pkg/connection"

// Target represents a target
// +genie:model_response
// +genie:model_paginated
type Target struct {
	ID            string               `json:"id"`
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
	CreatedAt     connection.DateTime  `json:"created_at"`
	UpdatedAt     connection.DateTime  `json:"updated_at"`
}

// Cluster represents a cluster
// +genie:model_response
// +genie:model_paginated
type Cluster struct {
	ID        string              `json:"id"`
	Name      string              `json:"name"`
	CreatedAt connection.DateTime `json:"created_at"`
	UpdatedAt connection.DateTime `json:"updated_at"`
}
