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

// CreateTargetGroupRequest represents a request to create a target group
type CreateTargetGroupRequest struct {
	ClusterID       string `json:"cluster_id"`
	Name            string `json:"name,omitempty"`
	Balance         string `json:"balance,omitempty"`
	Mode            string `json:"mode,omitempty"`
	Close           *bool  `json:"close,omitempty"`
	Sticky          *bool  `json:"sticky,omitempty"`
	CookieOpts      string `json:"cookie_opts,omitempty"`
	Source          string `json:"source,omitempty"`
	TimeoutsConnect string `json:"timeouts_connect,omitempty"`
	TimeoutServer   string `json:"timeouts_server,omitempty"`
	CustomOptions   string `json:"custom_options,omitempty"`
	MonitorURL      string `json:"monitor_url,omitempty"`
	MonitorHost     string `json:"monitor_host,omitempty"`
}

// PatchTargetGroupRequest represents a request to patch a target group
type PatchTargetGroupRequest struct {
	Name            string `json:"name,omitempty"`
	Balance         string `json:"balance,omitempty"`
	Mode            string `json:"mode,omitempty"`
	Close           *bool  `json:"close,omitempty"`
	Sticky          *bool  `json:"sticky,omitempty"`
	CookieOpts      string `json:"cookie_opts,omitempty"`
	Source          string `json:"source,omitempty"`
	TimeoutsConnect string `json:"timeouts_connect,omitempty"`
	TimeoutServer   string `json:"timeouts_server,omitempty"`
	CustomOptions   string `json:"custom_options,omitempty"`
	MonitorURL      string `json:"monitor_url,omitempty"`
	MonitorHost     string `json:"monitor_host,omitempty"`
}

// CreateVIPRequest represents a request to create a target group
type CreateVIPRequest struct {
	ClusterID string `json:"cluster_id"`
	Type      string `json:"type"`
	CIDR      string `json:"cidr"`
}

// PatchVIPRequest represents a request to patch a target group
type PatchVIPRequest struct {
	Type string `json:"type,omitempty"`
	CIDR string `json:"cidr,omitempty"`
}

// CreateErrorPageRequest represents a request to create an error page
type CreateErrorPageRequest struct {
	ListenerID string `json:"listener_id"`
	VIPsID     string `json:"vips_id"`
	Port       int    `json:"port"`
}

// PatchErrorPageRequest represents a request to patch an error page
type PatchErrorPageRequest struct {
	VIPsID string `json:"vips_id,omitempty"`
	Port   int    `json:"port,omitempty"`
}

// CreateListenerRequest represents a request to create a listener
type CreateListenerRequest struct {
	Name                 string `json:"name"`
	ClusterID            string `json:"cluster_id"`
	VIPsID               string `json:"vips_id"`
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
	VIPsID               string `json:"vips_id,omitempty"`
	Port                 int    `json:"port,omitempty"`
	HSTSEnabled          *bool  `json:"hsts_enabled,omitempty"`
	Mode                 string `json:"mode,omitempty"`
	HSTSMaxAge           int    `json:"hsts_maxage,omitempty"`
	Close                *bool  `json:"close,omitempty"`
	RedirectHTTPS        *bool  `json:"redirect_https,omitempty"`
	DefaultTargetGroupID string `json:"default_targetgroup_id,omitempty"`
}

// CreateAccessRequest represents a request to create an access rule
type CreateAccessRequest struct {
	Whitelist bool `json:"whitelist"`
}

// PatchAccessRequest represents a request to patch an access rule
type PatchAccessRequest struct {
	Whitelist *bool `json:"whitelist,omitempty"`
}

// CreateBindRequest represents a request to create a bind
type CreateBindRequest struct {
	VIPsID string `json:"vips_id"`
	Port   int    `json:"port"`
}

// PatchBindRequest represents a request to patch a bind
type PatchBindRequest struct {
	VIPsID string `json:"vips_id,omitempty"`
	Port   int    `json:"port,omitempty"`
}

// CreateListenerCertificateRequest represents a request to create a certificate
type CreateListenerCertificateRequest struct {
	CertsName string `json:"certs_name"`
	CertsPEM  string `json:"certs_pem"`
}

// PatchListenerCertificateRequest represents a request to patch a certificate
type PatchListenerCertificateRequest struct {
	CertsName string `json:"certs_name,omitempty"`
	CertsPEM  string `json:"certs_pem,omitempty"`
}

// CreateListenerErrorPageRequest represents a request to create a listener error page relationship
type CreateListenerErrorPageRequest struct {
	ErrorPageID string `json:"error_page_id"`
}

// PatchListenerErrorPageRequest represents a request to patch a listener error page relationship
type PatchListenerErrorPageRequest struct {
	ErrorPageID string `json:"error_page_id,omitempty"`
}

// CreateSSLRequest represents a request to create SSL options
type CreateSSLRequest struct {
	BindsID            string `json:"binds_id"`
	Enabled            bool   `json:"enabled"`
	AllowTLSv1         bool   `json:"allow_tlsv1"`
	AllowTLSv11        bool   `json:"allow_tlsv11"`
	DisableHTTP2       bool   `json:"disable_http2"`
	HTTP2Only          bool   `json:"http2_only"`
	CustomCiphers      string `json:"custom_ciphers"`
	CustomTLS13Ciphers string `json:"custom_tls13_ciphers"`
}

// PatchSSLRequest represents a request to patch SSL options
type PatchSSLRequest struct {
	BindsID            string `json:"binds_id,omitempty"`
	Enabled            *bool  `json:"enabled,omitempty"`
	AllowTLSv1         *bool  `json:"allow_tlsv1,omitempty"`
	AllowTLSv11        *bool  `json:"allow_tlsv11,omitempty"`
	DisableHTTP2       *bool  `json:"disable_http2,omitempty"`
	HTTP2Only          *bool  `json:"http2_only,omitempty"`
	CustomCiphers      string `json:"custom_ciphers,omitempty"`
	CustomTLS13Ciphers string `json:"custom_tls13_ciphers,omitempty"`
}

// CreateCustomOptionRequest represents a request to create a custom option
type CreateCustomOptionRequest struct {
	ListenerID    string `json:"listener_id,omitempty"`
	TargetGroupID string `json:"targetgroup_id,omitempty"`
	TargetID      string `json:"target_id,omitempty"`
	String        string `json:"string"`
}

// PatchCustomOptionRequest represents a request to patch a custom option
type PatchCustomOptionRequest struct {
	ListenerID    string `json:"listener_id,omitempty"`
	TargetGroupID string `json:"targetgroup_id,omitempty"`
	TargetID      string `json:"target_id,omitempty"`
	String        string `json:"string,omitempty"`
}
