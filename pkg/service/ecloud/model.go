package ecloud

import (
	"github.com/ans-group/sdk-go/pkg/connection"
)

type VirtualMachineStatus string

const (
	VirtualMachineStatusComplete   VirtualMachineStatus = "Complete"
	VirtualMachineStatusFailed     VirtualMachineStatus = "Failed"
	VirtualMachineStatusBeingBuilt VirtualMachineStatus = "Being Built"
)

type VirtualMachineDiskType string

const (
	VirtualMachineDiskTypeStandard VirtualMachineDiskType = "Standard"
	VirtualMachineDiskTypeCluster  VirtualMachineDiskType = "Cluster"
)

type VirtualMachinePowerStatus string

func (s VirtualMachinePowerStatus) String() string {
	return string(s)
}

const (
	VirtualMachinePowerStatusOnline  VirtualMachinePowerStatus = "Online"
	VirtualMachinePowerStatusOffline VirtualMachinePowerStatus = "Offline"
)

var VirtualMachinePowerStatusEnum connection.Enum[VirtualMachinePowerStatus] = []VirtualMachinePowerStatus{
	VirtualMachinePowerStatusOnline,
	VirtualMachinePowerStatusOffline,
}

type DatastoreStatus string

const (
	DatastoreStatusCompleted DatastoreStatus = "Completed"
	DatastoreStatusFailed    DatastoreStatus = "Failed"
	DatastoreStatusExpanding DatastoreStatus = "Expanding"
	DatastoreStatusQueued    DatastoreStatus = "Queued"
)

type SolutionEnvironment string

const (
	SolutionEnvironmentHybrid  SolutionEnvironment = "Hybrid"
	SolutionEnvironmentPrivate SolutionEnvironment = "Private"
)

type FirewallRole string

const (
	FirewallRoleNA     FirewallRole = "N/A"
	FirewallRoleMaster FirewallRole = "Master"
	FirewallRoleSlave  FirewallRole = "Slave"
)

// VirtualMachine represents an eCloud Virtual Machine
type VirtualMachine struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Hostname     string `json:"hostname"`
	ComputerName string `json:"computername"`
	// Count in Cores
	CPU int `json:"cpu"`
	// Size in GB
	RAM int `json:"ram"`
	// Size in GB
	HDD         int                  `json:"hdd"`
	IPInternal  connection.IPAddress `json:"ip_internal"`
	IPExternal  connection.IPAddress `json:"ip_external"`
	Platform    string               `json:"platform"`
	Template    string               `json:"template"`
	Backup      bool                 `json:"backup"`
	Support     bool                 `json:"support"`
	Environment string               `json:"environment"`
	SolutionID  int                  `json:"solution_id"`
	Status      VirtualMachineStatus `json:"status"`
	PowerStatus string               `json:"power_status"`
	ToolsStatus string               `json:"tools_status"`
	Disks       []VirtualMachineDisk `json:"hdd_disks"`
	Encrypted   bool                 `json:"encrypted"`
	Role        string               `json:"role"`
	GPUProfile  string               `json:"gpu_profile"`
}

// VirtualMachineDisk represents an eCloud Virtual Machine disk
type VirtualMachineDisk struct {
	UUID string                 `json:"uuid"`
	Name string                 `json:"name"`
	Type VirtualMachineDiskType `json:"type"`
	Key  int                    `json:"key"`

	// Size in GB
	Capacity int `json:"capacity"`
}

// Tag represents an eCloud tag
type Tag struct {
	Key       string              `json:"key"`
	Value     string              `json:"value"`
	CreatedAt connection.DateTime `json:"created_at"`
}

// Solution represents an eCloud solution
type Solution struct {
	ID                int                 `json:"id"`
	Name              string              `json:"name"`
	Environment       SolutionEnvironment `json:"environment"`
	PodID             int                 `json:"pod_id"`
	EncryptionEnabled bool                `json:"encryption_enabled"`
	EncryptionDefault bool                `json:"encryption_default"`
}

// Site represents an eCloud site
type Site struct {
	ID         int    `json:"id"`
	State      string `json:"state"`
	SolutionID int    `json:"solution_id"`
	PodID      int    `json:"pod_id"`
}

// V1Network represents an eCloud v1 network
type V1Network struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// V1Host represents an eCloud v1 host
type V1Host struct {
	ID         int     `json:"id"`
	SolutionID int     `json:"solution_id"`
	PodID      int     `json:"pod_id"`
	Name       string  `json:"name"`
	CPU        HostCPU `json:"cpu"`
	RAM        HostRAM `json:"ram"`
}

// HostCPU represents an eCloud host's CPU resources
type HostCPU struct {
	Quantity int    `json:"qty"`
	Cores    int    `json:"cores"`
	Speed    string `json:"speed"`
}

// HostRAM represents an eCloud host's RAM resources
type HostRAM struct {
	// Size in GB
	Capacity int `json:"capacity"`
	// Size in GB
	Reserved int `json:"reserved"`
	// Size in GB
	Allocated int `json:"allocated"`
	// Size in GB
	Available int `json:"available"`
}

// Datastore represents an eCloud datastore
type Datastore struct {
	ID         int             `json:"id"`
	SolutionID int             `json:"solution_id"`
	SiteID     int             `json:"site_id"`
	Name       string          `json:"name"`
	Status     DatastoreStatus `json:"status"`
	// Size in GB
	Capacity int `json:"capacity"`
	// Size in GB
	Allocated int `json:"allocated"`
	// Size in GB
	Available int `json:"available"`
}

// Firewall represents an eCloud firewall
type Firewall struct {
	ID       int                  `json:"id"`
	Name     string               `json:"name"`
	Hostname string               `json:"hostname"`
	IP       connection.IPAddress `json:"ip"`
	Role     FirewallRole         `json:"role"`
}

// FirewallConfig represents an eCloud firewall config
type FirewallConfig struct {
	Config string `json:"config"`
}

// Template represents an eCloud template
type Template struct {
	Name string `json:"name"`
	// Count in Cores
	CPU int `json:"cpu"`
	// Size in GB
	RAM int `json:"ram"`
	// Size in GB
	HDD             int                  `json:"hdd"`
	Disks           []VirtualMachineDisk `json:"hdd_disks"`
	Platform        string               `json:"platform"`
	OperatingSystem string               `json:"operating_system"`
	SolutionID      int                  `json:"solution_id"`
}

// Pod represents an eCloud pod
type Pod struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Services struct {
		Public     bool `json:"public"`
		Burst      bool `json:"burst"`
		Appliances bool `json:"appliances"`
	} `json:"services"`
}

// Appliance represents an eCloud appliance
type Appliance struct {
	ID               string              `json:"id"`
	Name             string              `json:"name"`
	LogoURI          string              `json:"logo_uri"`
	Description      string              `json:"description"`
	DocumentationURI string              `json:"documentation_uri"`
	Publisher        string              `json:"publisher"`
	CreatedAt        connection.DateTime `json:"created_at"`
}

// ApplianceParameter represents an eCloud appliance parameter
type ApplianceParameter struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Key            string `json:"key"`
	Type           string `json:"type"`
	Description    string `json:"description"`
	Required       bool   `json:"required"`
	ValidationRule string `json:"validation_rule"`
}

// ActiveDirectoryDomain represents an eCloud active directory domain
type ActiveDirectoryDomain struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TemplateType string

func (s TemplateType) String() string {
	return string(s)
}

const (
	TemplateTypeSolution TemplateType = "solution"
	TemplateTypePod      TemplateType = "pod"
)

var TemplateTypeEnum connection.Enum[TemplateType] = []TemplateType{
	TemplateTypeSolution,
	TemplateTypePod,
}

// ConsoleSession represents an eCloud Virtual Machine console session
type ConsoleSession struct {
	URL string `json:"url"`
}

type SyncStatus string

const (
	SyncStatusComplete   SyncStatus = "complete"
	SyncStatusFailed     SyncStatus = "failed"
	SyncStatusInProgress SyncStatus = "in-progress"
)

func (s SyncStatus) String() string {
	return string(s)
}

type SyncType string

const (
	SyncTypeUpdate SyncType = "update"
	SyncTypeDelete SyncType = "delete"
)

func (s SyncType) String() string {
	return string(s)
}

// ResourceSync represents the sync status of a resource
type ResourceSync struct {
	Status SyncStatus `json:"status"`
	Type   SyncType   `json:"type"`
}

// ResourceTask represents the task status of a resource
type ResourceTask struct {
	InProgress bool `json:"in_progress"`
}

type TaskStatus string

func (s TaskStatus) String() string {
	return string(s)
}

const (
	TaskStatusComplete   TaskStatus = "complete"
	TaskStatusFailed     TaskStatus = "failed"
	TaskStatusInProgress TaskStatus = "in-progress"
)

var TaskStatusEnum connection.Enum[TaskStatus] = []TaskStatus{
	TaskStatusComplete,
	TaskStatusFailed,
	TaskStatusInProgress,
}

// VPC represents an eCloud VPC
type VPC struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	RegionID           string              `json:"region_id"`
	ClientID           *int                `json:"client_id"`
	Sync               ResourceSync        `json:"sync"`
	Task               ResourceTask        `json:"task"`
	SupportEnabled     bool                `json:"support_enabled"`
	ConsoleEnabled     bool                `json:"console_enabled"`
	AdvancedNetworking bool                `json:"advanced_networking"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// AvailabilityZone represents an eCloud availability zone
type AvailabilityZone struct {
	ID               string `json:"id"`
	Code             string `json:"code"`
	Name             string `json:"name"`
	DatacentreSiteID int    `json:"datacentre_site_id"`
	RegionID         string `json:"region_id"`
}

// Network represents an eCloud network
type Network struct {
	ID        string              `json:"id"`
	Name      string              `json:"name"`
	RouterID  string              `json:"router_id"`
	Subnet    string              `json:"subnet"`
	Sync      ResourceSync        `json:"sync"`
	Task      ResourceTask        `json:"task"`
	CreatedAt connection.DateTime `json:"created_at"`
	UpdatedAt connection.DateTime `json:"updated_at"`
}

// DHCP represents an eCloud DHCP server/policy
type DHCP struct {
	ID                 string              `json:"id"`
	VPCID              string              `json:"vpc_id"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// VPN represents an eCloud VPN
type VPN struct {
	ID        string              `json:"id"`
	RouterID  string              `json:"router_id"`
	CreatedAt connection.DateTime `json:"created_at"`
	UpdatedAt connection.DateTime `json:"updated_at"`
}

// Instance represents an eCloud instance
type Instance struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	VPCID              string              `json:"vpc_id"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	ImageID            string              `json:"image_id"`
	VCPUCores          int                 `json:"vcpu_cores"`
	VCPUSockets        int                 `json:"vcpu_sockets"`
	VCPUCoresPerSocket int                 `json:"vcpu_cores_per_socket"`
	RAMCapacity        int                 `json:"ram_capacity"`
	Locked             bool                `json:"locked"`
	BackupEnabled      bool                `json:"backup_enabled"`
	BackupGatewayID    string              `json:"backup_gateway_id"`
	BackupAgentEnabled bool                `json:"secure_backup"` // TODO: Change tag to 'backup_agent_enabled' when ADO#34659 released
	IsEncrypted        bool                `json:"is_encrypted"`
	Platform           string              `json:"platform"`
	VolumeCapacity     int                 `json:"volume_capacity"`
	VolumeGroupID      string              `json:"volume_group_id"`
	HostGroupID        string              `json:"host_group_id"`
	ResourceTierID     string              `json:"resource_tier_id"`
	Sync               ResourceSync        `json:"sync"`
	Task               ResourceTask        `json:"task"`
	Online             *bool               `json:"online"`
	AgentRunning       *bool               `json:"agent_running"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// FloatingIP represents an eCloud floating IP address
type FloatingIP struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	VPCID              string              `json:"vpc_id"`
	IPAddress          string              `json:"ip_address"`
	ResourceID         string              `json:"resource_id"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	Sync               ResourceSync        `json:"sync"`
	Task               ResourceTask        `json:"task"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// FirewallPolicy represents an eCloud firewall policy
type FirewallPolicy struct {
	ID        string              `json:"id"`
	Name      string              `json:"name"`
	RouterID  string              `json:"router_id"`
	Sequence  int                 `json:"sequence"`
	Sync      ResourceSync        `json:"sync"`
	Task      ResourceTask        `json:"task"`
	CreatedAt connection.DateTime `json:"created_at"`
	UpdatedAt connection.DateTime `json:"updated_at"`
}

type FirewallRuleAction string

func (s FirewallRuleAction) String() string {
	return string(s)
}

const (
	FirewallRuleActionAllow  FirewallRuleAction = "ALLOW"
	FirewallRuleActionDrop   FirewallRuleAction = "DROP"
	FirewallRuleActionReject FirewallRuleAction = "REJECT"
)

var FirewallRuleActionEnum connection.Enum[FirewallRuleAction] = []FirewallRuleAction{
	FirewallRuleActionAllow,
	FirewallRuleActionDrop,
	FirewallRuleActionReject,
}

type FirewallRuleDirection string

func (s FirewallRuleDirection) String() string {
	return string(s)
}

const (
	FirewallRuleDirectionIn    FirewallRuleDirection = "IN"
	FirewallRuleDirectionOut   FirewallRuleDirection = "OUT"
	FirewallRuleDirectionInOut FirewallRuleDirection = "IN_OUT"
)

var FirewallRuleDirectionEnum connection.Enum[FirewallRuleDirection] = []FirewallRuleDirection{
	FirewallRuleDirectionIn,
	FirewallRuleDirectionOut,
	FirewallRuleDirectionInOut,
}

// FirewallRule represents an eCloud firewall rule
type FirewallRule struct {
	ID               string                `json:"id"`
	Name             string                `json:"name"`
	FirewallPolicyID string                `json:"firewall_policy_id"`
	Sequence         int                   `json:"sequence"`
	Source           string                `json:"source"`
	Destination      string                `json:"destination"`
	Action           FirewallRuleAction    `json:"action"`
	Direction        FirewallRuleDirection `json:"direction"`
	Enabled          bool                  `json:"enabled"`
	CreatedAt        connection.DateTime   `json:"created_at"`
	UpdatedAt        connection.DateTime   `json:"updated_at"`
}

type FirewallRulePortProtocol string

func (s FirewallRulePortProtocol) String() string {
	return string(s)
}

const (
	FirewallRulePortProtocolTCP    FirewallRulePortProtocol = "TCP"
	FirewallRulePortProtocolUDP    FirewallRulePortProtocol = "UDP"
	FirewallRulePortProtocolICMPv4 FirewallRulePortProtocol = "ICMPv4"
)

var FirewallRulePortProtocolEnum connection.Enum[FirewallRulePortProtocol] = []FirewallRulePortProtocol{
	FirewallRulePortProtocolTCP,
	FirewallRulePortProtocolUDP,
	FirewallRulePortProtocolICMPv4,
}

// FirewallRulePort represents an eCloud firewall rule port
type FirewallRulePort struct {
	ID             string                   `json:"id"`
	Name           string                   `json:"name"`
	FirewallRuleID string                   `json:"firewall_rule_id"`
	Protocol       FirewallRulePortProtocol `json:"protocol"`
	Source         string                   `json:"source"`
	Destination    string                   `json:"destination"`
	CreatedAt      connection.DateTime      `json:"created_at"`
	UpdatedAt      connection.DateTime      `json:"updated_at"`
}

// Region represents an eCloud region
type Region struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Router represents an eCloud router
type Router struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	VPCID              string              `json:"vpc_id"`
	RouterThroughputID string              `json:"router_throughput_id"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	Sync               ResourceSync        `json:"sync"`
	Task               ResourceTask        `json:"task"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// Credential represents an eCloud credential
type Credential struct {
	ID         string              `json:"id"`
	Name       string              `json:"name"`
	ResourceID string              `json:"resource_id"`
	Host       string              `json:"host"`
	Username   string              `json:"username"`
	Password   string              `json:"password"`
	Port       int                 `json:"port"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

type VolumeType string

const (
	VolumeTypeOS   VolumeType = "os"
	VolumeTypeData VolumeType = "data"
)

func (s VolumeType) String() string {
	return string(s)
}

// Volume represents an eCloud volume
type Volume struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	VPCID              string              `json:"vpc_id"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	Capacity           int                 `json:"capacity"`
	IOPS               int                 `json:"iops"`
	Attached           bool                `json:"attached"`
	Type               VolumeType          `json:"type"`
	VolumeGroupID      string              `json:"volume_group_id"`
	IsShared           bool                `json:"is_shared"`
	IsEncrypted        bool                `json:"is_encrypted"`
	Port               int                 `json:"port"`
	Sync               ResourceSync        `json:"sync"`
	Task               ResourceTask        `json:"task"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// NIC represents an eCloud network interface card
type NIC struct {
	ID         string              `json:"id"`
	MACAddress string              `json:"mac_address"`
	InstanceID string              `json:"instance_id"`
	NetworkID  string              `json:"network_id"`
	IPAddress  string              `json:"ip_address"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

// RouterThroughput represents an eCloud router throughput
type RouterThroughput struct {
	ID                 string              `json:"id"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	Name               string              `json:"name"`
	CommittedBandwidth int                 `json:"committed_bandwidth"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// DiscountPlan represents an eCloud discount plan
type DiscountPlan struct {
	ID                       string              `json:"id"`
	ResellerID               int                 `json:"reseller_id"`
	ContactID                int                 `json:"contact_id"`
	Name                     string              `json:"name"`
	CommitmentAmount         float32             `json:"commitment_amount"`
	CommitmentBeforeDiscount float32             `json:"commitment_before_discount"`
	DiscountRate             float32             `json:"discount_rate"`
	TermLength               int                 `json:"term_length"`
	TermStartDate            connection.DateTime `json:"term_start_date"`
	TermEndDate              connection.DateTime `json:"term_end_date"`
	Status                   string              `json:"status"`
	ResponseDate             connection.DateTime `json:"response_date"`
	CreatedAt                connection.DateTime `json:"created_at"`
	UpdatedAt                connection.DateTime `json:"updated_at"`
}

// BillingMetric represents an eCloud billing metric
type BillingMetric struct {
	ID         string              `json:"id"`
	ResourceID string              `json:"resource_id"`
	VPCID      string              `json:"vpc_id"`
	Key        string              `json:"key"`
	Value      string              `json:"value"`
	Start      connection.DateTime `json:"start"`
	End        connection.DateTime `json:"end"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

// Image represents an eCloud image
type Image struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	LogoURI            string              `json:"logo_uri"`
	Description        string              `json:"description"`
	DocumentationURI   string              `json:"documentation_uri"`
	Platform           string              `json:"platform"`
	Visibility         string              `json:"visibility"`
	VPCID              string              `json:"vpc_id"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	Sync               ResourceSync        `json:"sync"`
	Task               ResourceTask        `json:"task"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// ImageParameter represents an eCloud image parameter
type ImageParameter struct {
	ID             string              `json:"id"`
	Name           string              `json:"name"`
	Key            string              `json:"key"`
	Type           string              `json:"type"`
	Description    string              `json:"description"`
	Required       bool                `json:"required"`
	ValidationRule string              `json:"validation_rule"`
	CreatedAt      connection.DateTime `json:"created_at"`
	UpdatedAt      connection.DateTime `json:"updated_at"`
}

// ImageMetadata represents eCloud image metadata
type ImageMetadata struct {
	Key       string              `json:"key"`
	Value     string              `json:"value"`
	CreatedAt connection.DateTime `json:"created_at"`
	UpdatedAt connection.DateTime `json:"updated_at"`
}

// HostGroup represents an eCloud host group
type HostGroup struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	VPCID              string              `json:"vpc_id"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	WindowsEnabled     bool                `json:"windows_enabled"`
	HostSpecID         string              `json:"host_spec_id"`
	Sync               ResourceSync        `json:"sync"`
	Task               ResourceTask        `json:"task"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// HostSpec represents an eCloud host spec
type HostSpec struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	CPUSockets    int    `json:"cpu_sockets"`
	CPUType       string `json:"cpu_type"`
	CPUCores      int    `json:"cpu_cores"`
	CPUClockSpeed int    `json:"cpu_clock_speed"`
	RAMCapacity   int    `json:"ram_capacity"`
}

// Host represents an eCloud v2 host
type Host struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	HostGroupID string              `json:"host_group_id"`
	Sync        ResourceSync        `json:"sync"`
	Task        ResourceTask        `json:"task"`
	CreatedAt   connection.DateTime `json:"created_at"`
	UpdatedAt   connection.DateTime `json:"updated_at"`
}

// SSHKeyPair represents an eCloud SSH key pair
type SSHKeyPair struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	PublicKey string `json:"public_key"`
}

// Task represents a task against an eCloud resource
type Task struct {
	ID         string              `json:"id"`
	ResourceID string              `json:"resource_id"`
	Name       string              `json:"name"`
	Status     TaskStatus          `json:"status"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

// TaskReference represents a reference to an on-going task
type TaskReference struct {
	TaskID     string `json:"task_id"`
	ResourceID string `json:"id"`
}

// NetworkPolicy represents an eCloud network policy
type NetworkPolicy struct {
	ID        string              `json:"id"`
	Name      string              `json:"name"`
	NetworkID string              `json:"network_id"`
	VPCID     string              `json:"vpc_id"`
	Sync      ResourceSync        `json:"sync"`
	Task      ResourceTask        `json:"task"`
	CreatedAt connection.DateTime `json:"created_at"`
	UpdatedAt connection.DateTime `json:"updated_at"`
}

type NetworkRuleAction string

func (s NetworkRuleAction) String() string {
	return string(s)
}

const (
	NetworkRuleActionAllow  NetworkRuleAction = "ALLOW"
	NetworkRuleActionDrop   NetworkRuleAction = "DROP"
	NetworkRuleActionReject NetworkRuleAction = "REJECT"
)

var NetworkRuleActionEnum connection.Enum[NetworkRuleAction] = []NetworkRuleAction{
	NetworkRuleActionAllow,
	NetworkRuleActionDrop,
	NetworkRuleActionReject,
}

type NetworkRuleDirection string

func (s NetworkRuleDirection) String() string {
	return string(s)
}

const (
	NetworkRuleDirectionIn    NetworkRuleDirection = "IN"
	NetworkRuleDirectionOut   NetworkRuleDirection = "OUT"
	NetworkRuleDirectionInOut NetworkRuleDirection = "IN_OUT"
)

var NetworkRuleDirectionEnum connection.Enum[NetworkRuleDirection] = []NetworkRuleDirection{
	NetworkRuleDirectionIn,
	NetworkRuleDirectionOut,
	NetworkRuleDirectionInOut,
}

// NetworkRule represents an eCloud network rule
type NetworkRule struct {
	ID              string               `json:"id"`
	Name            string               `json:"name"`
	NetworkPolicyID string               `json:"network_policy_id"`
	Sequence        int                  `json:"sequence"`
	Source          string               `json:"source"`
	Destination     string               `json:"destination"`
	Type            string               `json:"type"`
	Action          NetworkRuleAction    `json:"action"`
	Direction       NetworkRuleDirection `json:"direction"`
	Enabled         bool                 `json:"enabled"`
	CreatedAt       connection.DateTime  `json:"created_at"`
	UpdatedAt       connection.DateTime  `json:"updated_at"`
}

type NetworkRulePortProtocol string

func (s NetworkRulePortProtocol) String() string {
	return string(s)
}

const (
	NetworkRulePortProtocolTCP    NetworkRulePortProtocol = "TCP"
	NetworkRulePortProtocolUDP    NetworkRulePortProtocol = "UDP"
	NetworkRulePortProtocolICMPv4 NetworkRulePortProtocol = "ICMPv4"
)

var NetworkRulePortProtocolEnum connection.Enum[NetworkRulePortProtocol] = []NetworkRulePortProtocol{
	NetworkRulePortProtocolTCP,
	NetworkRulePortProtocolUDP,
	NetworkRulePortProtocolICMPv4,
}

// NetworkRulePort represents an eCloud network rule port
type NetworkRulePort struct {
	ID            string                  `json:"id"`
	Name          string                  `json:"name"`
	NetworkRuleID string                  `json:"network_rule_id"`
	Protocol      NetworkRulePortProtocol `json:"protocol"`
	Source        string                  `json:"source"`
	Destination   string                  `json:"destination"`
	CreatedAt     connection.DateTime     `json:"created_at"`
	UpdatedAt     connection.DateTime     `json:"updated_at"`
}

// VolumeGroup represents an eCloud volume group resource
type VolumeGroup struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	VPCID              string              `json:"vpc_id"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	Usage              VolumeGroupUsage    `json:"usage"`
	Sync               ResourceSync        `json:"sync"`
	Task               ResourceTask        `json:"task"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

type VolumeGroupUsage struct {
	Volumes int `json:"volumes"`
}

// VPNProfileGroup represents an eCloud VPN profile group
type VPNProfileGroup struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	Description        string              `json:"description"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// VPNService represents an eCloud VPN service
type VPNService struct {
	ID        string              `json:"id"`
	Name      string              `json:"name"`
	RouterID  string              `json:"router_id"`
	VPCID     string              `json:"vpc_id"`
	Sync      ResourceSync        `json:"sync"`
	Task      ResourceTask        `json:"task"`
	CreatedAt connection.DateTime `json:"created_at"`
	UpdatedAt connection.DateTime `json:"updated_at"`
}

// VPNEndpoint represents an eCloud VPN endpoint
type VPNEndpoint struct {
	ID           string              `json:"id"`
	Name         string              `json:"name"`
	VPNServiceID string              `json:"vpn_service_id"`
	FloatingIPID string              `json:"floating_ip_id"`
	Sync         ResourceSync        `json:"sync"`
	Task         ResourceTask        `json:"task"`
	CreatedAt    connection.DateTime `json:"created_at"`
	UpdatedAt    connection.DateTime `json:"updated_at"`
}

// VPNSession represents an eCloud VPN session
type VPNSession struct {
	ID                string                   `json:"id"`
	Name              string                   `json:"name"`
	VPNProfileGroupID string                   `json:"vpn_profile_group_id"`
	VPNServiceID      string                   `json:"vpn_service_id"`
	VPNEndpointID     string                   `json:"vpn_endpoint_id"`
	RemoteIP          connection.IPAddress     `json:"remote_ip"`
	RemoteNetworks    string                   `json:"remote_networks"`
	LocalNetworks     string                   `json:"local_networks"`
	TunnelDetails     *VPNSessionTunnelDetails `json:"tunnel_details"`
	Sync              ResourceSync             `json:"sync"`
	Task              ResourceTask             `json:"task"`
	CreatedAt         connection.DateTime      `json:"created_at"`
	UpdatedAt         connection.DateTime      `json:"updated_at"`
}

type VPNSessionTunnelDetails struct {
	SessionState     string                       `json:"session_state"`
	TunnelStatistics []VPNSessionTunnelStatistics `json:"tunnel_statistics"`
}

type VPNSessionTunnelStatistics struct {
	TunnelStatus     string `json:"tunnel_status"`
	TunnelDownReason string `json:"tunnel_down_reason"`
	LocalSubnet      string `json:"local_subnet"`
	PeerSubnet       string `json:"peer_subnet"`
}

// VPNSessionPreSharedKey represents an eCloud VPN session pre-shared key
type VPNSessionPreSharedKey struct {
	PSK string `json:"psk"`
}

// LoadBalancer represents an eCloud loadbalancer
type LoadBalancer struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	VPCID              string              `json:"vpc_id"`
	LoadBalancerSpecID string              `json:"load_balancer_spec_id"`
	Sync               ResourceSync        `json:"sync"`
	Task               ResourceTask        `json:"task"`
	ConfigID           int                 `json:"config_id"`
	Nodes              int                 `json:"nodes"`
	NetworkID          string              `json:"network_id"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// LoadBalancerSpec represents an eCloud loadbalancer specification
type LoadBalancerSpec struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	CreatedAt   connection.DateTime `json:"created_at"`
	UpdatedAt   connection.DateTime `json:"updated_at"`
}

// VIP represents an eCloud load balancer VIP
type VIP struct {
	ID             string              `json:"id"`
	Name           string              `json:"name"`
	LoadBalancerID string              `json:"load_balancer_id"`
	IPAddressID    string              `json:"ip_address_id"`
	ConfigID       int                 `json:"config_id"`
	Sync           ResourceSync        `json:"sync"`
	Task           ResourceTask        `json:"task"`
	CreatedAt      connection.DateTime `json:"created_at"`
	UpdatedAt      connection.DateTime `json:"updated_at"`
}

// IP Address represents an eCloud VPC Network IP Address
type IPAddress struct {
	ID        string               `json:"id"`
	Name      string               `json:"name"`
	IPAddress connection.IPAddress `json:"ip_address"`
	NetworkID string               `json:"network_id"`
	Type      string               `json:"type"`
	Sync      ResourceSync         `json:"sync"`
	Task      ResourceTask         `json:"task"`
	CreatedAt connection.DateTime  `json:"created_at"`
	UpdatedAt connection.DateTime  `json:"updated_at"`
}

type AffinityRuleType string

func (s AffinityRuleType) String() string {
	return string(s)
}

const (
	Affinity     AffinityRuleType = "affinity"
	AntiAffinity AffinityRuleType = "anti-affinity"
)

var AffinityRuleTypeEnum connection.Enum[AffinityRuleType] = []AffinityRuleType{
	Affinity,
	AntiAffinity,
}

// AffinityRule represents an eCloud Affinity or Anti-Affinity Rule
type AffinityRule struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	VPCID              string              `json:"vpc_id"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	Type               AffinityRuleType    `json:"type"`
	Sync               ResourceSync        `json:"sync"`
	Task               ResourceTask        `json:"task"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// AffinityRuleMember represents an eCloud Affinity or Anti-Affinity rule member
type AffinityRuleMember struct {
	ID             string              `json:"id"`
	AffinityRuleID string              `json:"affinity_rule_id"`
	InstanceID     string              `json:"instance_id"`
	Sync           ResourceSync        `json:"sync"`
	Task           ResourceTask        `json:"task"`
	CreatedAt      connection.DateTime `json:"created_at"`
	UpdatedAt      connection.DateTime `json:"updated_at"`
}

type ResourceTier struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	AvailabilityZoneID string `json:"availability_zone_id"`
}

type NATOverloadRuleAction string

func (s NATOverloadRuleAction) String() string {
	return string(s)
}

const (
	NATOverloadRuleActionAllow NATOverloadRuleAction = "allow"
	NATOverloadRuleActionDeny  NATOverloadRuleAction = "deny"
)

var NATOverloadRuleActionEnum connection.Enum[NATOverloadRuleAction] = []NATOverloadRuleAction{
	NATOverloadRuleActionAllow,
	NATOverloadRuleActionDeny,
}

// NATOverloadRule represents an eCloud NAT overload rule
type NATOverloadRule struct {
	ID           string                `json:"id"`
	Name         string                `json:"name"`
	NetworkID    string                `json:"network_id"`
	FloatingIPID string                `json:"floating_ip_id"`
	Subnet       string                `json:"subnet"`
	Action       NATOverloadRuleAction `json:"action"`
	Sync         ResourceSync          `json:"sync"`
	Task         ResourceTask          `json:"task"`
	CreatedAt    connection.DateTime   `json:"created_at"`
	UpdatedAt    connection.DateTime   `json:"updated_at"`
}

// IOPSTier represents an eCloud IOPS tier
type IOPSTier struct {
	ID        string              `json:"id"`
	Name      string              `json:"name"`
	Level     int                 `json:"level"`
	CreatedAt connection.DateTime `json:"created_at"`
	UpdatedAt connection.DateTime `json:"updated_at"`
}

// VPNGateway represents a VPN gateway
type VPNGateway struct {
	ID              string              `json:"id"`
	Name            string              `json:"name"`
	RouterID        string              `json:"router_id"`
	SpecificationID string              `json:"specification_id"`
	Hostname        string              `json:"hostname"`
	FQDN            string              `json:"fqdn"`
	Sync            ResourceSync        `json:"sync"`
	Task            ResourceTask        `json:"task"`
	CreatedAt       connection.DateTime `json:"created_at"`
	UpdatedAt       connection.DateTime `json:"updated_at"`
}

// VPNGatewaySpecification represents a VPN gateway specification
type VPNGatewaySpecification struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	CreatedAt   connection.DateTime `json:"created_at"`
	UpdatedAt   connection.DateTime `json:"updated_at"`
}

// VPNGatewayUser represents a VPN gateway user
type VPNGatewayUser struct {
	ID           string              `json:"id"`
	Name         string              `json:"name"`
	VPNGatewayID string              `json:"vpn_gateway_id"`
	Username     string              `json:"username,omitempty"`
	Sync         ResourceSync        `json:"sync"`
	Task         ResourceTask        `json:"task"`
	CreatedAt    connection.DateTime `json:"created_at"`
	UpdatedAt    connection.DateTime `json:"updated_at"`
}

// BackupGatewaySpecification represents a Backup Gateway specification
type BackupGatewaySpecification struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// BackupGateway represents a Backup Gateway
type BackupGateway struct {
	ID                 string              `json:"id"`
	VPCID              string              `json:"vpc_id"`
	Name               string              `json:"name"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	GatewaySpecID      string              `json:"gateway_spec_id"`
	Sync               ResourceSync        `json:"sync"`
	Task               ResourceTask        `json:"task"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}
