//go:generate go run ../../gen/model_response/main.go -package loadbalancer -source model.go -destination model_response_generated.go
//go:generate go run ../../gen/model_paginated/main.go -package loadbalancer -source model.go -destination model_paginated_generated.go

package loadbalancer

import "github.com/ukfast/sdk-go/pkg/connection"

// Target represents a target
// +genie:model_response
// +genie:model_paginated
type Target struct {
	ID            int                  `json:"id"`
	TargetGroupID int                  `json:"targetgroup_id"`
	Name          string               `json:"name"`
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
	ID                   int                 `json:"id"`
	ClusterID            int                 `json:"cluster_id"`
	Name                 string              `json:"name"`
	Balance              string              `json:"balance"`
	Mode                 string              `json:"mode"`
	Close                bool                `json:"close"`
	Sticky               bool                `json:"sticky"`
	CookieOpts           string              `json:"cookie_opts"`
	Source               string              `json:"source"`
	TimeoutsConnect      int                 `json:"timeouts_connect"`
	TimeoutServer        int                 `json:"timeouts_server"`
	CustomOptions        string              `json:"custom_options"`
	MonitorURL           string              `json:"monitor_url"`
	MonitorMethod        string              `json:"monitor_method"`
	MonitorHost          string              `json:"monitor_host"`
	MonitorHTTPVersion   string              `json:"monitor_http_version"`
	MonitorExpect        string              `json:"monitor_expect"`
	MonitorTCPMonitoring bool                `json:"monitor_tcp_monitoring"`
	CheckPort            int                 `json:"check_port"`
	SendProxy            bool                `json:"send_proxy"`
	SendProxyV2          bool                `json:"send_proxy_v2"`
	CreatedAt            connection.DateTime `json:"created_at"`
	UpdatedAt            connection.DateTime `json:"updated_at"`
}

// Cluster represents a cluster
// +genie:model_response
// +genie:model_paginated
type Cluster struct {
	ID         int                 `json:"id"`
	Name       string              `json:"name"`
	Deployed   bool                `json:"deployed"`
	DeployedAt connection.DateTime `json:"deployed_at"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

// VIP represents a target virtual IP address
// +genie:model_response
// +genie:model_paginated
type VIP struct {
	ID           int                 `json:"id"`
	ClusterID    int                 `json:"cluster_id"`
	InternalCIDR string              `json:"internal_cidr"`
	ExternalCIDR string              `json:"external_cidr"`
	MACAddress   string              `json:"mac_address"`
	CreatedAt    connection.DateTime `json:"created_at"`
	UpdatedAt    connection.DateTime `json:"updated_at"`
}

type ListenerMode string

const (
	ListenerModeHTTP ListenerMode = "http"
	ListenerModeTCP  ListenerMode = "tcp"
)

var ListenerModeEnum connection.EnumSlice = []connection.Enum{
	ListenerModeHTTP,
	ListenerModeTCP,
}

// ParseListenerMode attempts to parse a ListenerMode from string
func ParseListenerMode(s string) (ListenerMode, error) {
	e, err := connection.ParseEnum(s, ListenerModeEnum)
	if err != nil {
		return "", err
	}

	return e.(ListenerMode), err
}

func (s ListenerMode) String() string {
	return string(s)
}

// Listener represents a listener / frontend
// +genie:model_response
// +genie:model_paginated
type Listener struct {
	ID                   int                 `json:"id"`
	Name                 string              `json:"name"`
	ClusterID            int                 `json:"cluster_id"`
	HSTSEnabled          bool                `json:"hsts_enabled"`
	Mode                 ListenerMode        `json:"mode"`
	HSTSMaxAge           int                 `json:"hsts_maxage"`
	Close                bool                `json:"close"`
	RedirectHTTPS        bool                `json:"redirect_https"`
	DefaultTargetGroupID int                 `json:"default_targetgroup_id"`
	AllowTLSV1           bool                `json:"allow_tlsv1"`
	AllowTLSV11          bool                `json:"allow_tlsv11"`
	DisableTLSV12        bool                `json:"disable_tlsv12"`
	DisableHTTP2         bool                `json:"disable_http2"`
	HTTP2Only            bool                `json:"http2_only"`
	CustomCiphers        string              `json:"custom_ciphers"`
	CreatedAt            connection.DateTime `json:"created_at"`
	UpdatedAt            connection.DateTime `json:"updated_at"`
}

// AccessIP represents an access IP
// +genie:model_response
// +genie:model_paginated
type AccessIP struct {
	ID        int                  `json:"id"`
	IP        connection.IPAddress `json:"ip"`
	CreatedAt connection.DateTime  `json:"created_at"`
	UpdatedAt connection.DateTime  `json:"updated_at"`
}

// Bind represents a bind
// +genie:model_response
// +genie:model_paginated
type Bind struct {
	ID         int                 `json:"id"`
	ListenerID int                 `json:"listener_id"`
	VIPID      int                 `json:"vip_id"`
	Port       int                 `json:"port"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

// Certificate represents a certificate
// +genie:model_response
// +genie:model_paginated
type Certificate struct {
	ID          int                 `json:"id"`
	ListenerID  int                 `json:"listener_id"`
	Name        string              `json:"name"`
	Key         string              `json:"key"`
	Certificate string              `json:"certificate"`
	CABundle    string              `json:"ca_bundle"`
	CreatedAt   connection.DateTime `json:"created_at"`
	UpdatedAt   connection.DateTime `json:"updated_at"`
}

// Header represents a header
// +genie:model_response
// +genie:model_paginated
type Header struct {
	Header string `json:"header"`
}

// ACL represents an ACL
// +genie:model_response
// +genie:model_paginated
type ACL struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	ListenerID    int            `json:"listener_id"`
	TargetGroupID int            `json:"target_group_id"`
	Conditions    []ACLCondition `json:"conditions"`
}

// ACLCondition represents an ACL condition
// +genie:model_response
// +genie:model_paginated
type ACLCondition struct {
	Name      string                 `json:"name"`
	Arguments []ACLConditionArgument `json:"arguments"`
}

// ACLConditionArgument represents an ACL condition argument
// +genie:model_response
// +genie:model_paginated
type ACLConditionArgument struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
