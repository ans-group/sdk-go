package ecloud

import "fmt"

// VirtualMachineNotFoundError indicates a virtual machine was not found within eCloud
type VirtualMachineNotFoundError struct {
	ID int
}

func (e *VirtualMachineNotFoundError) Error() string {
	return fmt.Sprintf("virtual machine not found with ID [%d]", e.ID)
}

// TagNotFoundError indicates a tag was not found within eCloud
type TagNotFoundError struct {
	Key string
}

func (e *TagNotFoundError) Error() string {
	return fmt.Sprintf("tag not found with key [%s]", e.Key)
}

// SolutionNotFoundError indicates a solution was not found within eCloud
type SolutionNotFoundError struct {
	ID int
}

func (e *SolutionNotFoundError) Error() string {
	return fmt.Sprintf("solution not found with ID [%d]", e.ID)
}

// SiteNotFoundError indicates a site was not found within eCloud
type SiteNotFoundError struct {
	ID int
}

func (e *SiteNotFoundError) Error() string {
	return fmt.Sprintf("site not found with ID [%d]", e.ID)
}

// HostNotFoundError indicates a host was not found within eCloud
type HostNotFoundError struct {
	ID int
}

func (e *HostNotFoundError) Error() string {
	return fmt.Sprintf("host not found with ID [%d]", e.ID)
}

// DatastoreNotFoundError indicates a datastore was not found within eCloud
type DatastoreNotFoundError struct {
	ID int
}

func (e *DatastoreNotFoundError) Error() string {
	return fmt.Sprintf("datastore not found with ID [%d]", e.ID)
}

// TemplateNotFoundError indicates a template was not found within eCloud
type TemplateNotFoundError struct {
	Name string
}

func (e *TemplateNotFoundError) Error() string {
	return fmt.Sprintf("template not found with name [%s]", e.Name)
}

// FirewallNotFoundError indicates a firewall was not found within eCloud
type FirewallNotFoundError struct {
	ID int
}

func (e *FirewallNotFoundError) Error() string {
	return fmt.Sprintf("firewall not found with ID [%d]", e.ID)
}

// PodNotFoundError indicates a pod was not found within eCloud
type PodNotFoundError struct {
	ID int
}

func (e *PodNotFoundError) Error() string {
	return fmt.Sprintf("pod not found with ID [%d]", e.ID)
}

// ApplianceNotFoundError indicates an appliance was not found within eCloud
type ApplianceNotFoundError struct {
	ID string
}

func (e *ApplianceNotFoundError) Error() string {
	return fmt.Sprintf("appliance not found with ID [%s]", e.ID)
}

// ActiveDirectoryDomainNotFoundError indicates an Active Directory Domain was not found
type ActiveDirectoryDomainNotFoundError struct {
	ID int
}

func (e *ActiveDirectoryDomainNotFoundError) Error() string {
	return fmt.Sprintf("domain not found with ID [%d]", e.ID)
}

// VPCNotFoundError indicates a VPC was not found
type VPCNotFoundError struct {
	ID string
}

func (e *VPCNotFoundError) Error() string {
	return fmt.Sprintf("VPC not found with ID [%s]", e.ID)
}

// AvailabilityZoneNotFoundError indicates a VPC was not found
type AvailabilityZoneNotFoundError struct {
	ID string
}

func (e *AvailabilityZoneNotFoundError) Error() string {
	return fmt.Sprintf("Availability zone not found with ID [%s]", e.ID)
}

// NetworkNotFoundError indicates a network was not found
type NetworkNotFoundError struct {
	ID string
}

func (e *NetworkNotFoundError) Error() string {
	return fmt.Sprintf("Network not found with ID [%s]", e.ID)
}

// DHCPNotFoundError indicates a DHCP server/config was not found
type DHCPNotFoundError struct {
	ID string
}

func (e *DHCPNotFoundError) Error() string {
	return fmt.Sprintf("DHCP not found with ID [%s]", e.ID)
}

// VPNNotFoundError indicates a VPN was not found
type VPNNotFoundError struct {
	ID string
}

func (e *VPNNotFoundError) Error() string {
	return fmt.Sprintf("VPN not found with ID [%s]", e.ID)
}

// InstanceNotFoundError indicates an instance was not found
type InstanceNotFoundError struct {
	ID string
}

func (e *InstanceNotFoundError) Error() string {
	return fmt.Sprintf("Instance not found with ID [%s]", e.ID)
}
