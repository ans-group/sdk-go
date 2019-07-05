package ecloud

import (
	"errors"
	"strings"

	"github.com/ukfast/sdk-go/pkg/connection"
)

type VirtualMachineStatus string

const (
	VirtualMachineStatusComplete   VirtualMachineStatus = "Complete"
	VirtualMachineStatusFailed     VirtualMachineStatus = "Failed"
	VirtualMachineStatusBeingBuilt VirtualMachineStatus = "Being Built"
)

func (s VirtualMachineStatus) String() string {
	return string(s)
}

type VirtualMachineDiskType string

func (e VirtualMachineDiskType) String() string {
	return string(e)
}

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

// ParseVirtualMachinePowerStatus attempts to parse a VirtualMachinePowerStatus from string
func ParseVirtualMachinePowerStatus(s string) (VirtualMachinePowerStatus, error) {
	switch strings.ToUpper(s) {
	case "ONLINE":
		return VirtualMachinePowerStatusOnline, nil
	case "OFFLINE":
		return VirtualMachinePowerStatusOffline, nil
	}

	return "", errors.New("Invalid power status")
}

type DatastoreStatus string

func (s DatastoreStatus) String() string {
	return string(s)
}

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

func (s SolutionEnvironment) String() string {
	return string(s)
}

type FirewallRole string

func (r FirewallRole) String() string {
	return string(r)
}

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
}

// PaginatedVirtualMachines represents a paginated collection of vms
type PaginatedVirtualMachines struct {
	*connection.PaginatedBase

	VirtualMachines []VirtualMachine
}

// NewPaginatedVirtualMachines returns a pointer to an initialized PaginatedVirtualMachines struct
func NewPaginatedVirtualMachines(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, vms []VirtualMachine) *PaginatedVirtualMachines {
	return &PaginatedVirtualMachines{
		VirtualMachines: vms,
		PaginatedBase:   connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
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

// PaginatedTags represents a paginated collection of tags
type PaginatedTags struct {
	*connection.PaginatedBase

	Tags []Tag
}

// NewPaginatedTags returns a pointer to an initialized PaginatedTags struct
func NewPaginatedTags(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, tags []Tag) *PaginatedTags {
	return &PaginatedTags{
		Tags:          tags,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// Solution represents an eCloud solution
type Solution struct {
	ID          int                 `json:"id"`
	Name        string              `json:"name"`
	Environment SolutionEnvironment `json:"environment"`
	PodID       int                 `json:"pod_id"`
}

// PaginatedSolutions represents a paginated collection of solutions
type PaginatedSolutions struct {
	*connection.PaginatedBase

	Solutions []Solution
}

// NewPaginatedSolutions returns a pointer to an initialized PaginatedSolutions struct
func NewPaginatedSolutions(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, solutions []Solution) *PaginatedSolutions {
	return &PaginatedSolutions{
		Solutions:     solutions,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// Site represents an eCloud site
type Site struct {
	ID         int    `json:"id"`
	State      string `json:"state"`
	SolutionID int    `json:"solution_id"`
	PodID      int    `json:"pod_id"`
}

// PaginatedSites represents a paginated collection of sites
type PaginatedSites struct {
	*connection.PaginatedBase

	Sites []Site
}

// NewPaginatedSites returns a pointer to an initialized PaginatedSites struct
func NewPaginatedSites(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, sites []Site) *PaginatedSites {
	return &PaginatedSites{
		Sites:         sites,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// Network represents an eCloud network
type Network struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// PaginatedNetworks represents a paginated collection of networks
type PaginatedNetworks struct {
	*connection.PaginatedBase

	Networks []Network
}

// NewPaginatedNetworks returns a pointer to an initialized PaginatedNetworks struct
func NewPaginatedNetworks(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, networks []Network) *PaginatedNetworks {
	return &PaginatedNetworks{
		Networks:      networks,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// Host represents an eCloud host
type Host struct {
	ID         int     `json:"id"`
	SolutionID int     `json:"solution_id"`
	PodID      int     `json:"pod_id"`
	Name       string  `json:"name"`
	CPU        HostCPU `json:"cpu"`
	RAM        HostRAM `json:"ram"`
}

// PaginatedHosts represents a paginated collection of hosts
type PaginatedHosts struct {
	*connection.PaginatedBase

	Hosts []Host
}

// NewPaginatedHosts returns a pointer to an initialized PaginatedHosts struct
func NewPaginatedHosts(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, hosts []Host) *PaginatedHosts {
	return &PaginatedHosts{
		Hosts:         hosts,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
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

// PaginatedDatastores represents a paginated collection of datastores
type PaginatedDatastores struct {
	*connection.PaginatedBase

	Datastores []Datastore
}

// NewPaginatedDatastores returns a pointer to an initialized PaginatedDatastores struct
func NewPaginatedDatastores(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, datastores []Datastore) *PaginatedDatastores {
	return &PaginatedDatastores{
		Datastores:    datastores,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// Firewall represents an eCloud firewall
type Firewall struct {
	ID       int                  `json:"id"`
	Name     string               `json:"name"`
	Hostname string               `json:"hostname"`
	IP       connection.IPAddress `json:"ip"`
	Role     FirewallRole         `json:"role"`
}

// PaginatedFirewalls represents a paginated collection of firewalls
type PaginatedFirewalls struct {
	*connection.PaginatedBase

	Firewalls []Firewall
}

// NewPaginatedFirewalls returns a pointer to an initialized PaginatedFirewalls struct
func NewPaginatedFirewalls(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, firewalls []Firewall) *PaginatedFirewalls {
	return &PaginatedFirewalls{
		Firewalls:     firewalls,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
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

// PaginatedTemplates represents a paginated collection of templates
type PaginatedTemplates struct {
	*connection.PaginatedBase

	Templates []Template
}

// NewPaginatedTemplates returns a pointer to an initialized PaginatedTemplates struct
func NewPaginatedTemplates(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, templates []Template) *PaginatedTemplates {
	return &PaginatedTemplates{
		Templates:     templates,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// Pod represents an eCloud pod
type Pod struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// PaginatedPods represents a paginated collection of pods
type PaginatedPods struct {
	*connection.PaginatedBase

	Pods []Pod
}

// NewPaginatedPods returns a pointer to an initialized PaginatedPods struct
func NewPaginatedPods(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, pods []Pod) *PaginatedPods {
	return &PaginatedPods{
		Pods:          pods,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
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

// PaginatedAppliances represents a paginated collection of appliances
type PaginatedAppliances struct {
	*connection.PaginatedBase

	Appliances []Appliance
}

// NewPaginatedAppliances returns a pointer to an initialized PaginatedAppliances struct
func NewPaginatedAppliances(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, appliances []Appliance) *PaginatedAppliances {
	return &PaginatedAppliances{
		Appliances:    appliances,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
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

// PaginatedApplianceParameters represents a paginated collection of parameters
type PaginatedApplianceParameters struct {
	*connection.PaginatedBase

	ApplianceParameters []ApplianceParameter
}

// NewPaginatedApplianceParameters returns a pointer to an initialized PaginatedApplianceParameters struct
func NewPaginatedApplianceParameters(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, appParameters []ApplianceParameter) *PaginatedApplianceParameters {
	return &PaginatedApplianceParameters{
		ApplianceParameters: appParameters,
		PaginatedBase:       connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
