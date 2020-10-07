package ecloud

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// PatchTagRequest represents an eCloud tag patch request
type PatchTagRequest struct {
	Value string `json:"value,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchTagRequest) Validate() *connection.ValidationError {
	return nil
}

// CreateTagRequest represents a request to create an eCloud tag
type CreateTagRequest struct {
	connection.APIRequestBodyDefaultValidator

	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateTagRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// CreateVirtualMachineRequest represents a request to create an eCloud virtual machine
type CreateVirtualMachineRequest struct {
	connection.APIRequestBodyDefaultValidator

	Environment      string `json:"environment" validate:"required"`
	Template         string `json:"template,omitempty"`
	ApplianceID      string `json:"appliance_id,omitempty"`
	TemplatePassword string `json:"template_password,omitempty"`
	// Count in Cores
	CPU int `json:"cpu" validate:"required"`
	// Size in GB
	RAM int `json:"ram" validate:"required"`
	// Size in GB
	HDD                     int                                    `json:"hdd,omitempty"`
	Disks                   []CreateVirtualMachineRequestDisk      `json:"hdd_disks,omitempty"`
	Name                    string                                 `json:"name,omitempty"`
	ComputerName            string                                 `json:"computername,omitempty"`
	Tags                    []CreateTagRequest                     `json:"tags,omitempty"`
	Backup                  bool                                   `json:"backup"`
	Support                 bool                                   `json:"support"`
	Monitoring              bool                                   `json:"monitoring"`
	MonitoringContacts      []int                                  `json:"monitoring_contacts"`
	SolutionID              int                                    `json:"solution_id,omitempty"`
	DatastoreID             int                                    `json:"datastore_id,omitempty"`
	SiteID                  int                                    `json:"site_id,omitempty"`
	NetworkID               int                                    `json:"network_id,omitempty"`
	ExternalIPRequired      bool                                   `json:"external_ip_required"`
	SSHKeys                 []string                               `json:"ssh_keys,omitempty"`
	Parameters              []CreateVirtualMachineRequestParameter `json:"parameters,omitempty"`
	Encrypt                 *bool                                  `json:"encrypt,omitempty"`
	Role                    string                                 `json:"role,omitempty"`
	BootstrapScript         string                                 `json:"bootstrap_script,omitempty"`
	ActiveDirectoryDomainID int                                    `json:"ad_domain_id,omitempty"`
	PodID                   int                                    `json:"pod_id,omitempty"`
}

// CreateVirtualMachineRequestDisk represents a request to create an eCloud virtual machine disk
type CreateVirtualMachineRequestDisk struct {
	Name string `json:"name,omitempty"`
	// Size in GB
	Capacity int `json:"capacity" validate:"required"`
}

// CreateVirtualMachineRequestParameter represents an eCloud virtual machine parameter
type CreateVirtualMachineRequestParameter struct {
	connection.APIRequestBodyDefaultValidator

	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateVirtualMachineRequest) Validate() *connection.ValidationError {
	if c.HDD == 0 && (c.Disks == nil || len(c.Disks) < 1) {
		return connection.NewValidationError("HDD or Disks must be provided")
	}

	if c.Template == "" && c.ApplianceID == "" {
		return connection.NewValidationError("Template or ApplianceID must be provided")
	}

	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// PatchSolutionRequest represents an eCloud solution patch request
type PatchSolutionRequest struct {
	Name              *string `json:"name,omitempty"`
	EncryptionDefault *bool   `json:"encryption_default,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchSolutionRequest) Validate() *connection.ValidationError {
	return nil
}

// RenameTemplateRequest represents an eCloud template rename request
type RenameTemplateRequest struct {
	connection.APIRequestBodyDefaultValidator

	Destination string `json:"destination" validate:"required"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *RenameTemplateRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// CloneVirtualMachineRequest represents a request to clone an eCloud virtual machine
type CloneVirtualMachineRequest struct {
	connection.APIRequestBodyDefaultValidator

	Name string `json:"name" validate:"required"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CloneVirtualMachineRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// PatchVirtualMachineRequest represents an eCloud virtual machine patch request
type PatchVirtualMachineRequest struct {
	Name *string `json:"name,omitempty"`
	// Count in Cores
	CPU int `json:"cpu,omitempty"`
	// Size in GB
	RAM int `json:"ram,omitempty"`
	// KV map of hard disks, key being hard disk name, value being size in GB
	Disks []PatchVirtualMachineRequestDisk `json:"hdd_disks,omitempty"`
	Role  string                           `json:"role,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchVirtualMachineRequest) Validate() *connection.ValidationError {
	return nil
}

type PatchVirtualMachineRequestDiskState string

const (
	PatchVirtualMachineRequestDiskStatePresent PatchVirtualMachineRequestDiskState = "present"
	PatchVirtualMachineRequestDiskStateAbsent  PatchVirtualMachineRequestDiskState = "absent"
)

func (s PatchVirtualMachineRequestDiskState) String() string {
	return string(s)
}

// PatchVirtualMachineRequestDisk represents an eCloud virtual machine patch request disk
type PatchVirtualMachineRequestDisk struct {
	UUID string `json:"uuid,omitempty"`
	// Size in GB
	Capacity int                                 `json:"capacity,omitempty"`
	State    PatchVirtualMachineRequestDiskState `json:"state,omitempty"`
}

// CreateVirtualMachineTemplateRequest represents a request to clone an eCloud virtual machine template
type CreateVirtualMachineTemplateRequest struct {
	connection.APIRequestBodyDefaultValidator

	TemplateName string       `json:"template_name" validate:"required"`
	TemplateType TemplateType `json:"template_type"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateVirtualMachineTemplateRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// CreateVPCRequest represents a request to create a VPC
type CreateVPCRequest struct {
	Name     *string `json:"name,omitempty"`
	RegionID string  `json:"region_id"`
}

// PatchVPCRequest represents a request to patch a VPC
type PatchVPCRequest struct {
	Name *string `json:"name,omitempty"`
}

// CreateNetworkRequest represents a request to create a network
type CreateNetworkRequest struct {
	Name     *string `json:"name,omitempty"`
	RouterID string  `json:"router_id"`
}

// PatchNetworkRequest represents a request to patch a network
type PatchNetworkRequest struct {
	Name *string `json:"name,omitempty"`
}

// CreateRouterRequest represents a request to create a router
type CreateRouterRequest struct {
	Name               *string `json:"name,omitempty"`
	VPCID              string  `json:"vpc_id"`
	AvailabilityZoneID string  `json:"availability_zone_id"`
}

// PatchRouterRequest represents a request to patch a router
type PatchRouterRequest struct {
	Name *string `json:"name,omitempty"`
}

// CreateVPNRequest represents a request to create a VPN
type CreateVPNRequest struct {
	RouterID string `json:"router_id"`
}

// CreateLoadBalancerClusterRequest represents a request to create a load balancer cluster
type CreateLoadBalancerClusterRequest struct {
	Name               *string `json:"name,omitempty"`
	VPCID              string  `json:"vpc_id"`
	AvailabilityZoneID string  `json:"availability_zone_id"`
	Nodes              int     `json:"nodes"`
}

// PatchLoadBalancerClusterRequest represents a request to patch a load balancer cluster
type PatchLoadBalancerClusterRequest struct {
	Name *string `json:"name,omitempty"`
}
