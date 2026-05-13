package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) routerRes() *resource.Resource[Router, string] {
	return resource.NewStringResource[Router](s.connection, "/ecloud/v2/routers", "router", func(id string) error {
		return &RouterNotFoundError{ID: id}
	})
}

// GetRouters retrieves a list of routers
func (s *Service) GetRouters(parameters connection.APIRequestParameters) ([]Router, error) {
	return s.routerRes().List(parameters)
}

// GetRoutersPaginated retrieves a paginated list of routers
func (s *Service) GetRoutersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Router], error) {
	return s.routerRes().ListPaginated(parameters)
}

// GetRouter retrieves a single router by id
func (s *Service) GetRouter(routerID string) (Router, error) {
	return s.routerRes().Get(routerID)
}

// CreateRouter creates a new Router
func (s *Service) CreateRouter(req CreateRouterRequest) (string, error) {
	data, err := s.routerRes().Create(&req)
	return data.ID, err
}

// PatchRouter patches a Router
func (s *Service) PatchRouter(routerID string, req PatchRouterRequest) error {
	return s.routerRes().Patch(routerID, &req)
}

// DeleteRouter deletes a Router
func (s *Service) DeleteRouter(routerID string) error {
	return s.routerRes().Delete(routerID)
}

func (s *Service) routerFirewallPolicyRes() *resource.SubResourceList[FirewallPolicy, string] {
	return resource.NewStringSubResourceList[FirewallPolicy](s.connection,
		func(routerID string) string { return fmt.Sprintf("/ecloud/v2/routers/%s/firewall-policies", routerID) },
		"router", "id", func(routerID string) error { return &RouterNotFoundError{ID: routerID} })
}

// GetRouterFirewallPolicies retrieves a list of firewall rule policies
func (s *Service) GetRouterFirewallPolicies(routerID string, parameters connection.APIRequestParameters) ([]FirewallPolicy, error) {
	return s.routerFirewallPolicyRes().List(routerID, parameters)
}

// GetRouterFirewallPoliciesPaginated retrieves a paginated list of firewall rule policies
func (s *Service) GetRouterFirewallPoliciesPaginated(routerID string, parameters connection.APIRequestParameters) (*connection.Paginated[FirewallPolicy], error) {
	return s.routerFirewallPolicyRes().ListPaginated(routerID, parameters)
}

func (s *Service) routerNetworkRes() *resource.SubResourceList[Network, string] {
	return resource.NewStringSubResourceList[Network](s.connection,
		func(routerID string) string { return fmt.Sprintf("/ecloud/v2/routers/%s/networks", routerID) },
		"router", "id", func(routerID string) error { return &RouterNotFoundError{ID: routerID} })
}

// GetRouterNetworks retrieves a list of router networks
func (s *Service) GetRouterNetworks(routerID string, parameters connection.APIRequestParameters) ([]Network, error) {
	return s.routerNetworkRes().List(routerID, parameters)
}

// GetRouterNetworksPaginated retrieves a paginated list of router networks
func (s *Service) GetRouterNetworksPaginated(routerID string, parameters connection.APIRequestParameters) (*connection.Paginated[Network], error) {
	return s.routerNetworkRes().ListPaginated(routerID, parameters)
}

func (s *Service) routerVPNRes() *resource.SubResourceList[VPN, string] {
	return resource.NewStringSubResourceList[VPN](s.connection,
		func(routerID string) string { return fmt.Sprintf("/ecloud/v2/routers/%s/vpns", routerID) },
		"router", "id", func(routerID string) error { return &RouterNotFoundError{ID: routerID} })
}

// GetRouterVPNs retrieves a list of router VPNs
func (s *Service) GetRouterVPNs(routerID string, parameters connection.APIRequestParameters) ([]VPN, error) {
	return s.routerVPNRes().List(routerID, parameters)
}

// GetRouterVPNsPaginated retrieves a paginated list of router VPNs
func (s *Service) GetRouterVPNsPaginated(routerID string, parameters connection.APIRequestParameters) (*connection.Paginated[VPN], error) {
	return s.routerVPNRes().ListPaginated(routerID, parameters)
}

// DeployRouterDefaultFirewallPolicies deploys default firewall policy resources for specified router
func (s *Service) DeployRouterDefaultFirewallPolicies(routerID string) error {
	if routerID == "" {
		return fmt.Errorf("invalid router id")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/ecloud/v2/routers/%s/configure-default-policies", routerID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&RouterNotFoundError{ID: routerID}))
}

func (s *Service) routerTasksRes() *resource.SubResourceList[Task, string] {
	return resource.NewStringSubResourceList[Task](s.connection,
		func(routerID string) string { return fmt.Sprintf("/ecloud/v2/routers/%s/tasks", routerID) },
		"router", "id", func(routerID string) error { return &RouterNotFoundError{ID: routerID} })
}

// GetRouterTasks retrieves a list of Router tasks
func (s *Service) GetRouterTasks(routerID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return s.routerTasksRes().List(routerID, parameters)
}

// GetRouterTasksPaginated retrieves a paginated list of Router tasks
func (s *Service) GetRouterTasksPaginated(routerID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	return s.routerTasksRes().ListPaginated(routerID, parameters)
}
