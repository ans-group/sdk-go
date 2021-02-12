package ecloud

import (
	"github.com/ukfast/sdk-go/pkg/connection"
	"github.com/ukfast/sdk-go/pkg/service/account"
)

// ECloudService is an interface for managing eCloud
type ECloudService interface {
	// Virtual Machine
	GetVirtualMachines(parameters connection.APIRequestParameters) ([]VirtualMachine, error)
	GetVirtualMachinesPaginated(parameters connection.APIRequestParameters) (*PaginatedVirtualMachine, error)
	GetVirtualMachine(vmID int) (VirtualMachine, error)
	CreateVirtualMachine(req CreateVirtualMachineRequest) (int, error)
	PatchVirtualMachine(vmID int, patch PatchVirtualMachineRequest) error
	CloneVirtualMachine(vmID int, req CloneVirtualMachineRequest) (int, error)
	DeleteVirtualMachine(vmID int) error
	PowerOnVirtualMachine(vmID int) error
	PowerOffVirtualMachine(vmID int) error
	PowerResetVirtualMachine(vmID int) error
	PowerShutdownVirtualMachine(vmID int) error
	PowerRestartVirtualMachine(vmID int) error
	CreateVirtualMachineTemplate(vmID int, req CreateVirtualMachineTemplateRequest) error
	GetVirtualMachineTags(vmID int, parameters connection.APIRequestParameters) ([]Tag, error)
	GetVirtualMachineTagsPaginated(vmID int, parameters connection.APIRequestParameters) (*PaginatedTag, error)
	GetVirtualMachineTag(vmID int, tagKey string) (Tag, error)
	CreateVirtualMachineTag(vmID int, req CreateTagRequest) error
	PatchVirtualMachineTag(vmID int, tagKey string, patch PatchTagRequest) error
	DeleteVirtualMachineTag(vmID int, tagKey string) error
	CreateVirtualMachineConsoleSession(vmID int) (ConsoleSession, error)

	// Solution
	GetSolutions(parameters connection.APIRequestParameters) ([]Solution, error)
	GetSolutionsPaginated(parameters connection.APIRequestParameters) (*PaginatedSolution, error)
	GetSolution(solutionID int) (Solution, error)
	PatchSolution(solutionID int, patch PatchSolutionRequest) (int, error)
	GetSolutionVirtualMachines(solutionID int, parameters connection.APIRequestParameters) ([]VirtualMachine, error)
	GetSolutionVirtualMachinesPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedVirtualMachine, error)
	GetSolutionSites(solutionID int, parameters connection.APIRequestParameters) ([]Site, error)
	GetSolutionSitesPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedSite, error)
	GetSolutionDatastores(solutionID int, parameters connection.APIRequestParameters) ([]Datastore, error)
	GetSolutionDatastoresPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedDatastore, error)
	GetSolutionHosts(solutionID int, parameters connection.APIRequestParameters) ([]Host, error)
	GetSolutionHostsPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedHost, error)
	GetSolutionNetworks(solutionID int, parameters connection.APIRequestParameters) ([]V1Network, error)
	GetSolutionNetworksPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedV1Network, error)
	GetSolutionFirewalls(solutionID int, parameters connection.APIRequestParameters) ([]Firewall, error)
	GetSolutionFirewallsPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedFirewall, error)
	GetSolutionTemplates(solutionID int, parameters connection.APIRequestParameters) ([]Template, error)
	GetSolutionTemplatesPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedTemplate, error)
	GetSolutionTemplate(solutionID int, templateName string) (Template, error)
	DeleteSolutionTemplate(solutionID int, templateName string) error
	RenameSolutionTemplate(solutionID int, templateName string, req RenameTemplateRequest) error
	GetSolutionTags(solutionID int, parameters connection.APIRequestParameters) ([]Tag, error)
	GetSolutionTagsPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedTag, error)
	GetSolutionTag(solutionID int, tagKey string) (Tag, error)
	CreateSolutionTag(solutionID int, req CreateTagRequest) error
	PatchSolutionTag(solutionID int, tagKey string, patch PatchTagRequest) error
	DeleteSolutionTag(solutionID int, tagKey string) error

	// Site
	GetSites(parameters connection.APIRequestParameters) ([]Site, error)
	GetSitesPaginated(parameters connection.APIRequestParameters) (*PaginatedSite, error)
	GetSite(siteID int) (Site, error)

	// Host
	GetHosts(parameters connection.APIRequestParameters) ([]Host, error)
	GetHostsPaginated(parameters connection.APIRequestParameters) (*PaginatedHost, error)
	GetHost(hostID int) (Host, error)

	// Datastore
	GetDatastores(parameters connection.APIRequestParameters) ([]Datastore, error)
	GetDatastoresPaginated(parameters connection.APIRequestParameters) (*PaginatedDatastore, error)
	GetDatastore(datastoreID int) (Datastore, error)

	// Firewall
	GetFirewalls(parameters connection.APIRequestParameters) ([]Firewall, error)
	GetFirewallsPaginated(parameters connection.APIRequestParameters) (*PaginatedFirewall, error)
	GetFirewall(firewallID int) (Firewall, error)
	GetFirewallConfig(firewallID int) (FirewallConfig, error)

	// Pod
	GetPods(parameters connection.APIRequestParameters) ([]Pod, error)
	GetPodsPaginated(parameters connection.APIRequestParameters) (*PaginatedPod, error)
	GetPod(podID int) (Pod, error)
	GetPodTemplates(podID int, parameters connection.APIRequestParameters) ([]Template, error)
	GetPodTemplatesPaginated(podID int, parameters connection.APIRequestParameters) (*PaginatedTemplate, error)
	GetPodTemplate(podID int, templateName string) (Template, error)
	RenamePodTemplate(podID int, templateName string, req RenameTemplateRequest) error
	DeletePodTemplate(podID int, templateName string) error
	GetPodAppliances(podID int, parameters connection.APIRequestParameters) ([]Appliance, error)
	GetPodAppliancesPaginated(podID int, parameters connection.APIRequestParameters) (*PaginatedAppliance, error)
	PodConsoleAvailable(podID int) (bool, error)

	// Appliance
	GetAppliances(parameters connection.APIRequestParameters) ([]Appliance, error)
	GetAppliancesPaginated(parameters connection.APIRequestParameters) (*PaginatedAppliance, error)
	GetAppliance(applianceID string) (Appliance, error)
	GetApplianceParameters(applianceID string, reqParameters connection.APIRequestParameters) ([]ApplianceParameter, error)
	GetApplianceParametersPaginated(applianceID string, parameters connection.APIRequestParameters) (*PaginatedApplianceParameter, error)

	// Credit
	GetCredits(parameters connection.APIRequestParameters) ([]account.Credit, error)

	GetActiveDirectoryDomains(parameters connection.APIRequestParameters) ([]ActiveDirectoryDomain, error)
	GetActiveDirectoryDomainsPaginated(parameters connection.APIRequestParameters) (*PaginatedActiveDirectoryDomain, error)
	GetActiveDirectoryDomain(domainID int) (ActiveDirectoryDomain, error)

	// V2

	// VPC
	GetVPCs(parameters connection.APIRequestParameters) ([]VPC, error)
	GetVPCsPaginated(parameters connection.APIRequestParameters) (*PaginatedVPC, error)
	GetVPC(vpcID string) (VPC, error)
	CreateVPC(req CreateVPCRequest) (string, error)
	PatchVPC(vpcID string, patch PatchVPCRequest) error
	DeleteVPC(vpcID string) error
	DeployVPCDefaults(vpcID string) error
	GetVPCVolumes(vpcID string, parameters connection.APIRequestParameters) ([]Volume, error)
	GetVPCVolumesPaginated(vpcID string, parameters connection.APIRequestParameters) (*PaginatedVolume, error)
	GetVPCInstances(vpcID string, parameters connection.APIRequestParameters) ([]Instance, error)
	GetVPCInstancesPaginated(vpcID string, parameters connection.APIRequestParameters) (*PaginatedInstance, error)

	// Availability zone
	GetAvailabilityZones(parameters connection.APIRequestParameters) ([]AvailabilityZone, error)
	GetAvailabilityZonesPaginated(parameters connection.APIRequestParameters) (*PaginatedAvailabilityZone, error)
	GetAvailabilityZone(azID string) (AvailabilityZone, error)

	// Network
	GetNetworks(parameters connection.APIRequestParameters) ([]Network, error)
	GetNetworksPaginated(parameters connection.APIRequestParameters) (*PaginatedNetwork, error)
	GetNetwork(networkID string) (Network, error)
	CreateNetwork(req CreateNetworkRequest) (string, error)
	PatchNetwork(networkID string, patch PatchNetworkRequest) error
	DeleteNetwork(networkID string) error
	GetNetworkNICs(networkID string, parameters connection.APIRequestParameters) ([]NIC, error)
	GetNetworkNICsPaginated(networkID string, parameters connection.APIRequestParameters) (*PaginatedNIC, error)

	// DHCP
	GetDHCPs(parameters connection.APIRequestParameters) ([]DHCP, error)
	GetDHCPsPaginated(parameters connection.APIRequestParameters) (*PaginatedDHCP, error)
	GetDHCP(dhcpID string) (DHCP, error)

	// VPN
	GetVPNs(parameters connection.APIRequestParameters) ([]VPN, error)
	GetVPNsPaginated(parameters connection.APIRequestParameters) (*PaginatedVPN, error)
	GetVPN(vpnID string) (VPN, error)
	CreateVPN(req CreateVPNRequest) (string, error)
	DeleteVPN(vpcID string) error

	// Instance
	GetInstances(parameters connection.APIRequestParameters) ([]Instance, error)
	GetInstancesPaginated(parameters connection.APIRequestParameters) (*PaginatedInstance, error)
	GetInstance(instanceID string) (Instance, error)
	CreateInstance(req CreateInstanceRequest) (string, error)
	PatchInstance(instanceID string, req PatchInstanceRequest) error
	DeleteInstance(instanceID string) error
	LockInstance(instanceID string) error
	UnlockInstance(instanceID string) error
	PowerOnInstance(instanceID string) error
	PowerOffInstance(instanceID string) error
	PowerResetInstance(instanceID string) error
	PowerShutdownInstance(instanceID string) error
	PowerRestartInstance(instanceID string) error
	GetInstanceVolumes(instanceID string, parameters connection.APIRequestParameters) ([]Volume, error)
	GetInstanceVolumesPaginated(instanceID string, parameters connection.APIRequestParameters) (*PaginatedVolume, error)
	GetInstanceCredentials(instanceID string, parameters connection.APIRequestParameters) ([]Credential, error)
	GetInstanceCredentialsPaginated(instanceID string, parameters connection.APIRequestParameters) (*PaginatedCredential, error)
	GetInstanceNICs(instanceID string, parameters connection.APIRequestParameters) ([]NIC, error)
	GetInstanceNICsPaginated(instanceID string, parameters connection.APIRequestParameters) (*PaginatedNIC, error)

	// Floating IP
	GetFloatingIPs(parameters connection.APIRequestParameters) ([]FloatingIP, error)
	GetFloatingIPsPaginated(parameters connection.APIRequestParameters) (*PaginatedFloatingIP, error)
	GetFloatingIP(fipID string) (FloatingIP, error)
	CreateFloatingIP(req CreateFloatingIPRequest) (string, error)
	PatchFloatingIP(fipID string, req PatchFloatingIPRequest) error
	DeleteFloatingIP(fipID string) error
	AssignFloatingIP(fipID string, req AssignFloatingIPRequest) error
	UnassignFloatingIP(fipID string) error

	// Firewall Policy
	GetFirewallPolicies(parameters connection.APIRequestParameters) ([]FirewallPolicy, error)
	GetFirewallPoliciesPaginated(parameters connection.APIRequestParameters) (*PaginatedFirewallPolicy, error)
	GetFirewallPolicy(policyID string) (FirewallPolicy, error)
	CreateFirewallPolicy(req CreateFirewallPolicyRequest) (string, error)
	PatchFirewallPolicy(policyID string, req PatchFirewallPolicyRequest) error
	DeleteFirewallPolicy(policyID string) error
	GetFirewallPolicyFirewallRules(policyID string, parameters connection.APIRequestParameters) ([]FirewallRule, error)
	GetFirewallPolicyFirewallRulesPaginated(policyID string, parameters connection.APIRequestParameters) (*PaginatedFirewallRule, error)

	// Firewall Rule
	GetFirewallRules(parameters connection.APIRequestParameters) ([]FirewallRule, error)
	GetFirewallRulesPaginated(parameters connection.APIRequestParameters) (*PaginatedFirewallRule, error)
	GetFirewallRule(ruleID string) (FirewallRule, error)
	CreateFirewallRule(req CreateFirewallRuleRequest) (string, error)
	PatchFirewallRule(ruleID string, req PatchFirewallRuleRequest) error
	DeleteFirewallRule(ruleID string) error
	GetFirewallRuleFirewallRulePorts(firewallRuleID string, parameters connection.APIRequestParameters) ([]FirewallRulePort, error)
	GetFirewallRuleFirewallRulePortsPaginated(firewallRuleID string, parameters connection.APIRequestParameters) (*PaginatedFirewallRulePort, error)

	// Firewall Rule Ports
	GetFirewallRulePorts(parameters connection.APIRequestParameters) ([]FirewallRulePort, error)
	GetFirewallRulePortsPaginated(parameters connection.APIRequestParameters) (*PaginatedFirewallRulePort, error)
	GetFirewallRulePort(ruleID string) (FirewallRulePort, error)
	CreateFirewallRulePort(req CreateFirewallRulePortRequest) (string, error)
	PatchFirewallRulePort(ruleID string, req PatchFirewallRulePortRequest) error
	DeleteFirewallRulePort(ruleID string) error

	// Router
	GetRouters(parameters connection.APIRequestParameters) ([]Router, error)
	GetRoutersPaginated(parameters connection.APIRequestParameters) (*PaginatedRouter, error)
	GetRouter(routerID string) (Router, error)
	CreateRouter(req CreateRouterRequest) (string, error)
	PatchRouter(routerID string, patch PatchRouterRequest) error
	DeleteRouter(routerID string) error
	GetRouterFirewallPolicies(routerID string, parameters connection.APIRequestParameters) ([]FirewallPolicy, error)
	GetRouterFirewallPoliciesPaginated(routerID string, parameters connection.APIRequestParameters) (*PaginatedFirewallPolicy, error)
	GetRouterNetworks(routerID string, parameters connection.APIRequestParameters) ([]Network, error)
	GetRouterNetworksPaginated(routerID string, parameters connection.APIRequestParameters) (*PaginatedNetwork, error)
	GetRouterVPNs(routerID string, parameters connection.APIRequestParameters) ([]VPN, error)
	GetRouterVPNsPaginated(routerID string, parameters connection.APIRequestParameters) (*PaginatedVPN, error)
	DeployRouterDefaultFirewallPolicies(routerID string) error

	// Region
	GetRegions(parameters connection.APIRequestParameters) ([]Region, error)
	GetRegionsPaginated(parameters connection.APIRequestParameters) (*PaginatedRegion, error)
	GetRegion(regionID string) (Region, error)

	// Load balancers
	GetLoadBalancerClusters(parameters connection.APIRequestParameters) ([]LoadBalancerCluster, error)
	GetLoadBalancerClustersPaginated(parameters connection.APIRequestParameters) (*PaginatedLoadBalancerCluster, error)
	GetLoadBalancerCluster(lbcID string) (LoadBalancerCluster, error)
	CreateLoadBalancerCluster(req CreateLoadBalancerClusterRequest) (string, error)
	PatchLoadBalancerCluster(lbcID string, patch PatchLoadBalancerClusterRequest) error
	DeleteLoadBalancerCluster(lbcID string) error

	// Volumes
	GetVolumes(parameters connection.APIRequestParameters) ([]Volume, error)
	GetVolumesPaginated(parameters connection.APIRequestParameters) (*PaginatedVolume, error)
	GetVolume(volumeID string) (Volume, error)
	PatchVolume(volumeID string, patch PatchVolumeRequest) error
	DeleteVolume(volumeID string) error
	GetVolumeInstances(volumeID string, parameters connection.APIRequestParameters) ([]Instance, error)
	GetVolumeInstancesPaginated(volumeID string, parameters connection.APIRequestParameters) (*PaginatedInstance, error)

	// NICs
	GetNICs(parameters connection.APIRequestParameters) ([]NIC, error)
	GetNICsPaginated(parameters connection.APIRequestParameters) (*PaginatedNIC, error)
	GetNIC(nicID string) (NIC, error)

	// Billing metrics
	GetBillingMetrics(parameters connection.APIRequestParameters) ([]BillingMetric, error)
	GetBillingMetricsPaginated(parameters connection.APIRequestParameters) (*PaginatedBillingMetric, error)
	GetBillingMetric(metricID string) (BillingMetric, error)

	// Router throughputs
	GetRouterThroughputs(parameters connection.APIRequestParameters) ([]RouterThroughput, error)
	GetRouterThroughputsPaginated(parameters connection.APIRequestParameters) (*PaginatedRouterThroughput, error)
	GetRouterThroughput(metricID string) (RouterThroughput, error)
}

// Service implements ECloudService for managing
// eCloud via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of eCloud Service
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
