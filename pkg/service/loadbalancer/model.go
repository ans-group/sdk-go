package loadbalancer

import "github.com/ans-group/sdk-go/pkg/connection"

// Target represents a target
type Target struct {
	ID                 int                  `json:"id"`
	TargetGroupID      int                  `json:"target_group_id"`
	Name               string               `json:"name"`
	IP                 connection.IPAddress `json:"ip"`
	Port               int                  `json:"port"`
	Weight             int                  `json:"weight"`
	Backup             bool                 `json:"backup"`
	CheckInterval      int                  `json:"check_interval"`
	CheckSSL           bool                 `json:"check_ssl"`
	CheckRise          int                  `json:"check_rise"`
	CheckFall          int                  `json:"check_fall"`
	DisableHTTP2       bool                 `json:"disable_http2"`
	HTTP2Only          bool                 `json:"http2_only"`
	Active             bool                 `json:"active"`
	SessionCookieValue string               `json:"session_cookie_value"`
	CreatedAt          connection.DateTime  `json:"created_at"`
	UpdatedAt          connection.DateTime  `json:"updated_at"`
}

type TargetGroupBalance string

const (
	TargetGroupBalanceRoundRobin TargetGroupBalance = "roundrobin"
	TargetGroupBalanceStaticRR   TargetGroupBalance = "static-rr"
	TargetGroupBalanceLeastConn  TargetGroupBalance = "leastconn"
	TargetGroupBalanceFirst      TargetGroupBalance = "first"
	TargetGroupBalanceRDPCookie  TargetGroupBalance = "rdp-cookie"
	TargetGroupBalanceURI        TargetGroupBalance = "uri"
	TargetGroupBalanceHDR        TargetGroupBalance = "hdr"
	TargetGroupBalanceURLParam   TargetGroupBalance = "url_param"
	TargetGroupBalanceSource     TargetGroupBalance = "source"
)

var TargetGroupBalanceEnum connection.EnumSlice = []connection.Enum{
	TargetGroupBalanceRoundRobin,
	TargetGroupBalanceStaticRR,
	TargetGroupBalanceLeastConn,
	TargetGroupBalanceFirst,
	TargetGroupBalanceRDPCookie,
	TargetGroupBalanceURI,
	TargetGroupBalanceHDR,
	TargetGroupBalanceURLParam,
	TargetGroupBalanceSource,
}

// ParseTargetGroupBalance attempts to parse a TargetGroupBalance from string
func ParseTargetGroupBalance(s string) (TargetGroupBalance, error) {
	e, err := connection.ParseEnum(s, TargetGroupBalanceEnum)
	if err != nil {
		return "", err
	}

	return e.(TargetGroupBalance), err
}

func (s TargetGroupBalance) String() string {
	return string(s)
}

type TargetGroupMonitorMethod string

const (
	TargetGroupMonitorMethodGET     TargetGroupMonitorMethod = "GET"
	TargetGroupMonitorMethodHEAD    TargetGroupMonitorMethod = "HEAD"
	TargetGroupMonitorMethodOPTIONS TargetGroupMonitorMethod = "OPTIONS"
)

var TargetGroupMonitorMethodEnum connection.EnumSlice = []connection.Enum{
	TargetGroupMonitorMethodGET,
	TargetGroupMonitorMethodHEAD,
	TargetGroupMonitorMethodOPTIONS,
}

// ParseTargetGroupMonitorMethod attempts to parse a TargetGroupMonitorMethod from string
func ParseTargetGroupMonitorMethod(s string) (TargetGroupMonitorMethod, error) {
	e, err := connection.ParseEnum(s, TargetGroupMonitorMethodEnum)
	if err != nil {
		return "", err
	}

	return e.(TargetGroupMonitorMethod), err
}

func (s TargetGroupMonitorMethod) String() string {
	return string(s)
}

// TargetGroup represents a target group
type TargetGroup struct {
	ID                       int                      `json:"id"`
	ClusterID                int                      `json:"cluster_id"`
	Name                     string                   `json:"name"`
	Balance                  TargetGroupBalance       `json:"balance"`
	Mode                     Mode                     `json:"mode"`
	Close                    bool                     `json:"close"`
	Sticky                   bool                     `json:"sticky"`
	CookieOpts               string                   `json:"cookie_opts"`
	Source                   string                   `json:"source"`
	TimeoutsConnect          int                      `json:"timeouts_connect"`
	TimeoutsServer           int                      `json:"timeouts_server"`
	TimeoutsHTTPRequest      int                      `json:"timeouts_http_request"`
	TimeoutsCheck            int                      `json:"timeouts_check"`
	TimeoutsTunnel           int                      `json:"timeouts_tunnel"`
	CustomOptions            string                   `json:"custom_options"`
	MonitorURL               string                   `json:"monitor_url"`
	MonitorMethod            TargetGroupMonitorMethod `json:"monitor_method"`
	MonitorHost              string                   `json:"monitor_host"`
	MonitorHTTPVersion       string                   `json:"monitor_http_version"`
	MonitorExpect            string                   `json:"monitor_expect"`
	MonitorExpectString      string                   `json:"monitor_expect_string"`
	MonitorExpectStringRegex bool                     `json:"monitor_expect_string_regex"`
	MonitorTCPMonitoring     bool                     `json:"monitor_tcp_monitoring"`
	CheckPort                int                      `json:"check_port"`
	SendProxy                bool                     `json:"send_proxy"`
	SendProxyV2              bool                     `json:"send_proxy_v2"`
	SSL                      bool                     `json:"ssl"`
	SSLVerify                bool                     `json:"ssl_verify"`
	SNI                      bool                     `json:"sni"`
	CreatedAt                connection.DateTime      `json:"created_at"`
	UpdatedAt                connection.DateTime      `json:"updated_at"`
}

// Cluster represents a cluster
type Cluster struct {
	ID         int                 `json:"id"`
	Name       string              `json:"name"`
	Deployed   bool                `json:"deployed"`
	DeployedAt connection.DateTime `json:"deployed_at"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

// VIP represents a target virtual IP address
type VIP struct {
	ID           int                 `json:"id"`
	ClusterID    int                 `json:"cluster_id"`
	InternalCIDR string              `json:"internal_cidr"`
	ExternalCIDR string              `json:"external_cidr"`
	MACAddress   string              `json:"mac_address"`
	CreatedAt    connection.DateTime `json:"created_at"`
	UpdatedAt    connection.DateTime `json:"updated_at"`
}

type Mode string

const (
	ModeHTTP Mode = "http"
	ModeTCP  Mode = "tcp"
)

var ModeEnum connection.EnumSlice = []connection.Enum{
	ModeHTTP,
	ModeTCP,
}

// ParseMode attempts to parse a Mode from string
func ParseMode(s string) (Mode, error) {
	e, err := connection.ParseEnum(s, ModeEnum)
	if err != nil {
		return "", err
	}

	return e.(Mode), err
}

func (s Mode) String() string {
	return string(s)
}

// Listener represents a listener / frontend
type Listener struct {
	ID                   int                 `json:"id"`
	Name                 string              `json:"name"`
	ClusterID            int                 `json:"cluster_id"`
	HSTSEnabled          bool                `json:"hsts_enabled"`
	Mode                 Mode                `json:"mode"`
	HSTSMaxAge           int                 `json:"hsts_maxage"`
	Close                bool                `json:"close"`
	RedirectHTTPS        bool                `json:"redirect_https"`
	DefaultTargetGroupID int                 `json:"default_target_group_id"`
	AccessIsAllowList    bool                `json:"access_is_allow_list"`
	AllowTLSV1           bool                `json:"allow_tlsv1"`
	AllowTLSV11          bool                `json:"allow_tlsv11"`
	DisableTLSV12        bool                `json:"disable_tlsv12"`
	DisableHTTP2         bool                `json:"disable_http2"`
	HTTP2Only            bool                `json:"http2_only"`
	CustomCiphers        string              `json:"custom_ciphers"`
	CustomOptions        string              `json:"custom_options"`
	TimeoutsClient       int                 `json:"timeouts_client"`
	GeoIP                *ListenerGeoIP      `json:"geoip"`
	CreatedAt            connection.DateTime `json:"created_at"`
	UpdatedAt            connection.DateTime `json:"updated_at"`
}
type ListenerGeoIPRestriction string

const (
	ListenerGeoIPRestrictionAllow ListenerGeoIPRestriction = "allow"
	ListenerGeoIPRestrictionDeny  ListenerGeoIPRestriction = "deny"
)

// ParseListenerGeoIPRestriction attempts to parse a ListenerGeoIPRestriction from string
func ParseListenerGeoIPRestriction(s string) (ListenerGeoIPRestriction, error) {
	e, err := connection.ParseEnum(s, ListenerGeoIPRestrictionEnum)
	if err != nil {
		return "", err
	}

	return e.(ListenerGeoIPRestriction), err
}

func (s ListenerGeoIPRestriction) String() string {
	return string(s)
}

var ListenerGeoIPRestrictionEnum connection.EnumSlice = []connection.Enum{
	ListenerGeoIPRestrictionAllow,
	ListenerGeoIPRestrictionDeny,
}

type ListenerGeoIP struct {
	Restriction   ListenerGeoIPRestriction `json:"restriction"`
	Continents    []string                 `json:"continents"`
	Countries     []string                 `json:"countries"`
	EuropeanUnion bool                     `json:"european_union"`
}

// AccessIP represents an access IP
type AccessIP struct {
	ID        int                  `json:"id"`
	IP        connection.IPAddress `json:"ip"`
	CreatedAt connection.DateTime  `json:"created_at"`
	UpdatedAt connection.DateTime  `json:"updated_at"`
}

// Bind represents a bind
type Bind struct {
	ID         int                 `json:"id"`
	ListenerID int                 `json:"listener_id"`
	VIPID      int                 `json:"vip_id"`
	Port       int                 `json:"port"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

// Certificate represents a certificate
type Certificate struct {
	ID         int                 `json:"id"`
	ListenerID int                 `json:"listener_id"`
	Name       string              `json:"name"`
	ExpiresAt  connection.DateTime `json:"expires_at"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

// ACL represents an ACL
type ACL struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	ListenerID    int            `json:"listener_id"`
	TargetGroupID int            `json:"target_group_id"`
	Conditions    []ACLCondition `json:"conditions"`
	Actions       []ACLAction    `json:"actions"`
}

// ACLArgument represents an ACL condition/action argument
type ACLArgument struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// ACLCondition represents an ACL condition
type ACLCondition struct {
	Name      string                 `json:"name"`
	Inverted  bool                   `json:"inverted"`
	Arguments map[string]ACLArgument `json:"arguments"`
}

// ACLAction represents an ACL action
type ACLAction struct {
	Name      string                 `json:"name"`
	Arguments map[string]ACLArgument `json:"arguments"`
}

// ACLTemplates represents a collection of ACL condition/action templates
type ACLTemplates struct {
	Conditions []ACLTemplateCondition `json:"conditions"`
	Actions    []ACLTemplateAction    `json:"actions"`
}

type ACLTemplateCondition struct {
	Name         string                `json:"name"`
	FriendlyName string                `json:"friendly_name"`
	Description  string                `json:"description"`
	Arguments    []ACLTemplateArgument `json:"arguments"`
}

type ACLTemplateAction struct {
	Name         string                `json:"name"`
	FriendlyName string                `json:"friendly_name"`
	Description  string                `json:"description"`
	Arguments    []ACLTemplateArgument `json:"arguments"`
}

type ACLTemplateArgument struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Example     interface{} `json:"example"`
	Values      []string    `json:"values"`
}

// Deployment represents a load balancer deployment
type Deployment struct {
	ID              int                 `json:"id"`
	ClusterID       int                 `json:"cluster_id"`
	Successful      bool                `json:"successful"`
	RequestedByType string              `json:"requested_by_type"`
	RequestedByID   string              `json:"requested_by_id"`
	PSSID           int                 `json:"pss_id"`
	CreatedAt       connection.DateTime `json:"created_at"`
	UpdatedAt       connection.DateTime `json:"updated_at"`
}
