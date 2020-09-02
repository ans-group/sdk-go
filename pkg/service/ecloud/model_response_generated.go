package ecloud

import "github.com/ukfast/sdk-go/pkg/connection"

// GetVirtualMachineSliceResponseBody represents an API response body containing []VirtualMachine data
type GetVirtualMachineSliceResponseBody struct {
	connection.APIResponseBody
	Data []VirtualMachine `json:"data"`
}

// GetVirtualMachineResponseBody represents an API response body containing VirtualMachine data
type GetVirtualMachineResponseBody struct {
	connection.APIResponseBody
	Data VirtualMachine `json:"data"`
}

// GetTagSliceResponseBody represents an API response body containing []Tag data
type GetTagSliceResponseBody struct {
	connection.APIResponseBody
	Data []Tag `json:"data"`
}

// GetTagResponseBody represents an API response body containing Tag data
type GetTagResponseBody struct {
	connection.APIResponseBody
	Data Tag `json:"data"`
}

// GetSolutionSliceResponseBody represents an API response body containing []Solution data
type GetSolutionSliceResponseBody struct {
	connection.APIResponseBody
	Data []Solution `json:"data"`
}

// GetSolutionResponseBody represents an API response body containing Solution data
type GetSolutionResponseBody struct {
	connection.APIResponseBody
	Data Solution `json:"data"`
}

// GetSiteSliceResponseBody represents an API response body containing []Site data
type GetSiteSliceResponseBody struct {
	connection.APIResponseBody
	Data []Site `json:"data"`
}

// GetSiteResponseBody represents an API response body containing Site data
type GetSiteResponseBody struct {
	connection.APIResponseBody
	Data Site `json:"data"`
}

// GetNetworkSliceResponseBody represents an API response body containing []Network data
type GetNetworkSliceResponseBody struct {
	connection.APIResponseBody
	Data []Network `json:"data"`
}

// GetNetworkResponseBody represents an API response body containing Network data
type GetNetworkResponseBody struct {
	connection.APIResponseBody
	Data Network `json:"data"`
}

// GetHostSliceResponseBody represents an API response body containing []Host data
type GetHostSliceResponseBody struct {
	connection.APIResponseBody
	Data []Host `json:"data"`
}

// GetHostResponseBody represents an API response body containing Host data
type GetHostResponseBody struct {
	connection.APIResponseBody
	Data Host `json:"data"`
}

// GetDatastoreSliceResponseBody represents an API response body containing []Datastore data
type GetDatastoreSliceResponseBody struct {
	connection.APIResponseBody
	Data []Datastore `json:"data"`
}

// GetDatastoreResponseBody represents an API response body containing Datastore data
type GetDatastoreResponseBody struct {
	connection.APIResponseBody
	Data Datastore `json:"data"`
}

// GetFirewallSliceResponseBody represents an API response body containing []Firewall data
type GetFirewallSliceResponseBody struct {
	connection.APIResponseBody
	Data []Firewall `json:"data"`
}

// GetFirewallResponseBody represents an API response body containing Firewall data
type GetFirewallResponseBody struct {
	connection.APIResponseBody
	Data Firewall `json:"data"`
}

// GetFirewallConfigSliceResponseBody represents an API response body containing []FirewallConfig data
type GetFirewallConfigSliceResponseBody struct {
	connection.APIResponseBody
	Data []FirewallConfig `json:"data"`
}

// GetFirewallConfigResponseBody represents an API response body containing FirewallConfig data
type GetFirewallConfigResponseBody struct {
	connection.APIResponseBody
	Data FirewallConfig `json:"data"`
}

// GetTemplateSliceResponseBody represents an API response body containing []Template data
type GetTemplateSliceResponseBody struct {
	connection.APIResponseBody
	Data []Template `json:"data"`
}

// GetTemplateResponseBody represents an API response body containing Template data
type GetTemplateResponseBody struct {
	connection.APIResponseBody
	Data Template `json:"data"`
}

// GetPodSliceResponseBody represents an API response body containing []Pod data
type GetPodSliceResponseBody struct {
	connection.APIResponseBody
	Data []Pod `json:"data"`
}

// GetPodResponseBody represents an API response body containing Pod data
type GetPodResponseBody struct {
	connection.APIResponseBody
	Data Pod `json:"data"`
}

// GetApplianceSliceResponseBody represents an API response body containing []Appliance data
type GetApplianceSliceResponseBody struct {
	connection.APIResponseBody
	Data []Appliance `json:"data"`
}

// GetApplianceResponseBody represents an API response body containing Appliance data
type GetApplianceResponseBody struct {
	connection.APIResponseBody
	Data Appliance `json:"data"`
}

// GetApplianceParameterSliceResponseBody represents an API response body containing []ApplianceParameter data
type GetApplianceParameterSliceResponseBody struct {
	connection.APIResponseBody
	Data []ApplianceParameter `json:"data"`
}

// GetApplianceParameterResponseBody represents an API response body containing ApplianceParameter data
type GetApplianceParameterResponseBody struct {
	connection.APIResponseBody
	Data ApplianceParameter `json:"data"`
}

// GetActiveDirectoryDomainSliceResponseBody represents an API response body containing []ActiveDirectoryDomain data
type GetActiveDirectoryDomainSliceResponseBody struct {
	connection.APIResponseBody
	Data []ActiveDirectoryDomain `json:"data"`
}

// GetActiveDirectoryDomainResponseBody represents an API response body containing ActiveDirectoryDomain data
type GetActiveDirectoryDomainResponseBody struct {
	connection.APIResponseBody
	Data ActiveDirectoryDomain `json:"data"`
}

// GetConsoleSessionSliceResponseBody represents an API response body containing []ConsoleSession data
type GetConsoleSessionSliceResponseBody struct {
	connection.APIResponseBody
	Data []ConsoleSession `json:"data"`
}

// GetConsoleSessionResponseBody represents an API response body containing ConsoleSession data
type GetConsoleSessionResponseBody struct {
	connection.APIResponseBody
	Data ConsoleSession `json:"data"`
}
