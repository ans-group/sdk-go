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

// TargetGroup represents a target group
// +genie:model_response
// +genie:model_paginated
type TargetGroup struct {
	ID              string              `json:"id"`
	ClusterID       string              `json:"cluster_id"`
	Name            string              `json:"name"`
	Balance         string              `json:"balance"`
	Mode            string              `json:"mode"`
	Close           bool                `json:"close"`
	Sticky          bool                `json:"sticky"`
	CookieOpts      string              `json:"cookie_opts"`
	Source          string              `json:"source"`
	TimeoutsConnect int                 `json:"timeouts_connect"`
	TimeoutServer   int                 `json:"timeouts_server"`
	CustomOptions   string              `json:"custom_options"`
	MonitorURL      string              `json:"monitor_url"`
	MonitorHost     string              `json:"monitor_host"`
	CreatedAt       connection.DateTime `json:"created_at"`
	UpdatedAt       connection.DateTime `json:"updated_at"`
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

// VIP represents a target virtual IP address
// +genie:model_response
// +genie:model_paginated
type VIP struct {
	ID        string              `json:"id"`
	ClusterID string              `json:"cluster_id"`
	Type      string              `json:"type"`
	CIDR      string              `json:"cidr"`
	CreatedAt connection.DateTime `json:"created_at"`
	UpdatedAt connection.DateTime `json:"updated_at"`
}

// Listener represents a listener / frontend
// +genie:model_response
// +genie:model_paginated
type Listener struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ClusterID string `json:"cluster_id"`
	// TBC - should VIPsID and Port be removed here in favour of related binds?
	VIPsID               string              `json:"vips_id"`
	Port                 int                 `json:"port"`
	HSTSEnabled          bool                `json:"hsts_enabled"`
	Mode                 string              `json:"mode"`
	HSTSMaxAge           int                 `json:"hsts_maxage"`
	Close                bool                `json:"close"`
	RedirectHTTPS        bool                `json:"redirect_https"`
	DefaultTargetGroupID string              `json:"default_targetgroup_id"`
	CreatedAt            connection.DateTime `json:"created_at"`
	UpdatedAt            connection.DateTime `json:"updated_at"`
}

// ErrorPage represents an error page
// +genie:model_response
// +genie:model_paginated
type ErrorPage struct {
	ID         string              `json:"id"`
	StatusCode string              `json:"status_code"`
	Content    string              `json:"content"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

// Access represents an access rule
// +genie:model_response
// +genie:model_paginated
type Access struct {
	ID         string              `json:"id"`
	ListenerID string              `json:"listener_id"`
	Whitelist  bool                `json:"whitelist"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

// Bind represents a bind
// +genie:model_response
// +genie:model_paginated
type Bind struct {
	ID         string              `json:"id"`
	ListenerID string              `json:"listener_id"`
	VIPsID     string              `json:"vips_id"`
	Port       int                 `json:"port"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

// ListenerCertificate represents a listener certificate
// +genie:model_response
// +genie:model_paginated
type ListenerCertificate struct {
	ID         string              `json:"id"`
	ListenerID string              `json:"listener_id"`
	CertsName  string              `json:"certs_name"`
	CertsPEM   string              `json:"certs_pem"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

// ListenerErrorPage represents a listener error page relationship
// +genie:model_response
// +genie:model_paginated
type ListenerErrorPage struct {
	ID          string              `json:"id"`
	ListenerID  string              `json:"listener_id"`
	ErrorPageID string              `json:"error_page_id"`
	CreatedAt   connection.DateTime `json:"created_at"`
	UpdatedAt   connection.DateTime `json:"updated_at"`
}

// SSL represents a SSL options
// +genie:model_response
// +genie:model_paginated
type SSL struct {
	ID                 string              `json:"id"`
	BindsID            string              `json:"binds_id"`
	Enabled            bool                `json:"enabled"`
	AllowTLSv1         bool                `json:"allow_tlsv1"`
	AllowTLSv11        bool                `json:"allow_tlsv11"`
	DisableHTTP2       bool                `json:"disable_http2"`
	HTTP2Only          bool                `json:"http2_only"`
	CustomCiphers      string              `json:"custom_ciphers"`
	CustomTLS13Ciphers string              `json:"custom_tls13_ciphers"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// CustomOption represents a custom option
// +genie:model_response
// +genie:model_paginated
type CustomOption struct {
	ID            string              `json:"id"`
	ListenerID    string              `json:"listener_id"`
	TargetGroupID string              `json:"targetgroup_id"`
	TargetID      string              `json:"target_id"`
	String        string              `json:"string"`
	CreatedAt     connection.DateTime `json:"created_at"`
	UpdatedAt     connection.DateTime `json:"updated_at"`
}
