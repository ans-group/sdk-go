package loadbalancer

import "github.com/ukfast/sdk-go/pkg/connection"

// PatchClusterRequest represents a request to patch a cluster
type PatchClusterRequest struct {
	Name string `json:"name,omitempty"`
}

// CreateTargetRequest represents a request to create a target
type CreateTargetRequest struct {
	TargetGroupID int                  `json:"targetgroup_id"`
	Name          string               `json:"name,omitempty"`
	IP            connection.IPAddress `json:"ip"`
	Port          int                  `json:"port,omitempty"`
	Weight        int                  `json:"weight,omitempty"`
	Backup        bool                 `json:"backup"`
	CheckInterval int                  `json:"check_interval,omitempty"`
	CheckSSL      bool                 `json:"check_ssl"`
	CheckRise     int                  `json:"check_rise,omitempty"`
	CheckFall     int                  `json:"check_fall,omitempty"`
	DisableHTTP2  bool                 `json:"disable_http2"`
	HTTP2Only     bool                 `json:"http2_only"`
}

// PatchTargetRequest represents a request to patch a target
type PatchTargetRequest struct {
	TargetGroupID int                  `json:"targetgroup_id,omitempty"`
	Name          string               `json:"name,omitempty"`
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

// CreateTargetGroupRequest represents a request to create a target group
type CreateTargetGroupRequest struct {
	ClusterID            int    `json:"cluster_id"`
	Name                 string `json:"name,omitempty"`
	Balance              string `json:"balance,omitempty"`
	Mode                 string `json:"mode,omitempty"`
	Close                bool   `json:"close"`
	Sticky               bool   `json:"sticky"`
	CookieOpts           string `json:"cookie_opts,omitempty"`
	Source               string `json:"source,omitempty"`
	TimeoutsConnect      int    `json:"timeouts_connect,omitempty"`
	TimeoutServer        int    `json:"timeouts_server,omitempty"`
	CustomOptions        string `json:"custom_options,omitempty"`
	MonitorURL           string `json:"monitor_url,omitempty"`
	MonitorMethod        string `json:"monitor_method,omitempty"`
	MonitorHost          string `json:"monitor_host,omitempty"`
	MonitorHTTPVersion   string `json:"monitor_http_version,omitempty"`
	MonitorExpect        string `json:"monitor_expect,omitempty"`
	MonitorTCPMonitoring bool   `json:"monitor_tcp_monitoring"`
	CheckPort            int    `json:"check_port,omitempty"`
	SendProxy            bool   `json:"send_proxy"`
	SendProxyV2          bool   `json:"send_proxy_v2"`
}

// PatchTargetGroupRequest represents a request to patch a target group
type PatchTargetGroupRequest struct {
	ClusterID            int    `json:"cluster_id,omitempty"`
	Name                 string `json:"name,omitempty"`
	Balance              string `json:"balance,omitempty"`
	Mode                 string `json:"mode,omitempty"`
	Close                *bool  `json:"close,omitempty"`
	Sticky               *bool  `json:"sticky,omitempty"`
	CookieOpts           string `json:"cookie_opts,omitempty"`
	Source               string `json:"source,omitempty"`
	TimeoutsConnect      int    `json:"timeouts_connect,omitempty"`
	TimeoutServer        int    `json:"timeouts_server,omitempty"`
	CustomOptions        string `json:"custom_options,omitempty"`
	MonitorURL           string `json:"monitor_url,omitempty"`
	MonitorMethod        string `json:"monitor_method,omitempty"`
	MonitorHost          string `json:"monitor_host,omitempty"`
	MonitorHTTPVersion   string `json:"monitor_http_version,omitempty"`
	MonitorExpect        string `json:"monitor_expect,omitempty"`
	MonitorTCPMonitoring *bool  `json:"monitor_tcp_monitoring,omitempty"`
	CheckPort            int    `json:"check_port,omitempty"`
	SendProxy            *bool  `json:"send_proxy,omitempty"`
	SendProxyV2          *bool  `json:"send_proxy_v2,omitempty"`
}

// CreateVIPRequest represents a request to create a target group
type CreateVIPRequest struct {
	ClusterID int    `json:"cluster_id"`
	Type      string `json:"type"`
	CIDR      string `json:"cidr"`
}

// PatchVIPRequest represents a request to patch a target group
type PatchVIPRequest struct {
	Type string `json:"type,omitempty"`
	CIDR string `json:"cidr,omitempty"`
}

// CreateListenerRequest represents a request to create a listener
type CreateListenerRequest struct {
	Name                 string `json:"name"`
	ClusterID            int    `json:"cluster_id"`
	VIPID                int    `json:"vips_id"`
	Port                 int    `json:"port"`
	HSTSEnabled          bool   `json:"hsts_enabled"`
	Mode                 string `json:"mode"`
	HSTSMaxAge           int    `json:"hsts_maxage"`
	Close                bool   `json:"close"`
	RedirectHTTPS        bool   `json:"redirect_https"`
	DefaultTargetGroupID string `json:"default_targetgroup_id"`
}

// PatchListenerRequest represents a request to patch a listener
type PatchListenerRequest struct {
	Name                 string `json:"name,omitempty"`
	VIPID                int    `json:"vips_id,omitempty"`
	Port                 int    `json:"port,omitempty"`
	HSTSEnabled          *bool  `json:"hsts_enabled,omitempty"`
	Mode                 string `json:"mode,omitempty"`
	HSTSMaxAge           int    `json:"hsts_maxage,omitempty"`
	Close                *bool  `json:"close,omitempty"`
	RedirectHTTPS        *bool  `json:"redirect_https,omitempty"`
	DefaultTargetGroupID int    `json:"default_targetgroup_id,omitempty"`
}

// CreateAccessIPRequest represents a request to create an access rule
type CreateAccessIPRequest struct {
	IP connection.IPAddress `json:"ip"`
}

// PatchAccessIPRequest represents a request to patch an access rule
type PatchAccessIPRequest struct {
	IP connection.IPAddress `json:"ip,omitempty"`
}

// CreateBindRequest represents a request to create a bind
type CreateBindRequest struct {
	ListenerID int `json:"listener_id"`
	VIPID      int `json:"vip_id"`
	Port       int `json:"port"`
}

// PatchBindRequest represents a request to patch a bind
type PatchBindRequest struct {
	ListenerID int `json:"listener_id,omitempty"`
	VIPID      int `json:"vip_id,omitempty"`
	Port       int `json:"port,omitempty"`
}

// CreateCertificateRequest represents a request to create a certificate
type CreateCertificateRequest struct {
	ListenerID  int    `json:"listener_id"`
	Name        string `json:"name"`
	Key         string `json:"key"`
	Certificate string `json:"certificate"`
	CABundle    string `json:"ca_bundle"`
}

// PatchListenerCertificateRequest represents a request to patch a certificate
type PatchCertificateRequest struct {
	ListenerID  int    `json:"listener_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Key         string `json:"key,omitempty"`
	Certificate string `json:"certificate,omitempty"`
	CABundle    string `json:"ca_bundle,omitempty"`
}

// CreateACLRequest represents a request to create a ACL
type CreateACLRequest struct {
	CertsName string `json:"certs_name"`
	CertsPEM  string `json:"certs_pem"`
}

// PatchListenerACLRequest represents a request to patch a ACL
type PatchACLRequest struct {
	CertsName string `json:"certs_name,omitempty"`
	CertsPEM  string `json:"certs_pem,omitempty"`
}
