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

// GetRouterFirewallPolicies retrieves a list of firewall rule policies
func (s *Service) GetRouterFirewallPolicies(routerID string, parameters connection.APIRequestParameters) ([]FirewallPolicy, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[FirewallPolicy], error) {
		return s.GetRouterFirewallPoliciesPaginated(routerID, p)
	}, parameters)
}

// GetRouterFirewallPoliciesPaginated retrieves a paginated list of firewall rule policies
func (s *Service) GetRouterFirewallPoliciesPaginated(routerID string, parameters connection.APIRequestParameters) (*connection.Paginated[FirewallPolicy], error) {
	if routerID == "" {
		return nil, fmt.Errorf("invalid router id")
	}
	body, err := connection.Get[[]FirewallPolicy](s.connection, fmt.Sprintf("/ecloud/v2/routers/%s/firewall-policies", routerID), parameters, connection.NotFoundResponseHandler(&RouterNotFoundError{ID: routerID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[FirewallPolicy], error) {
		return s.GetRouterFirewallPoliciesPaginated(routerID, p)
	}), err
}

// GetRouterNetworks retrieves a list of router networks
func (s *Service) GetRouterNetworks(routerID string, parameters connection.APIRequestParameters) ([]Network, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Network], error) {
		return s.GetRouterNetworksPaginated(routerID, p)
	}, parameters)
}

// GetRouterNetworksPaginated retrieves a paginated list of router networks
func (s *Service) GetRouterNetworksPaginated(routerID string, parameters connection.APIRequestParameters) (*connection.Paginated[Network], error) {
	if routerID == "" {
		return nil, fmt.Errorf("invalid router id")
	}
	body, err := connection.Get[[]Network](s.connection, fmt.Sprintf("/ecloud/v2/routers/%s/networks", routerID), parameters, connection.NotFoundResponseHandler(&RouterNotFoundError{ID: routerID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Network], error) {
		return s.GetRouterNetworksPaginated(routerID, p)
	}), err
}

// GetRouterVPNs retrieves a list of router VPNs
func (s *Service) GetRouterVPNs(routerID string, parameters connection.APIRequestParameters) ([]VPN, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[VPN], error) {
		return s.GetRouterVPNsPaginated(routerID, p)
	}, parameters)
}

// GetRouterVPNsPaginated retrieves a paginated list of router VPNs
func (s *Service) GetRouterVPNsPaginated(routerID string, parameters connection.APIRequestParameters) (*connection.Paginated[VPN], error) {
	if routerID == "" {
		return nil, fmt.Errorf("invalid router id")
	}
	body, err := connection.Get[[]VPN](s.connection, fmt.Sprintf("/ecloud/v2/routers/%s/vpns", routerID), parameters, connection.NotFoundResponseHandler(&RouterNotFoundError{ID: routerID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[VPN], error) {
		return s.GetRouterVPNsPaginated(routerID, p)
	}), err
}

// DeployRouterDefaultFirewallPolicies deploys default firewall policy resources for specified router
func (s *Service) DeployRouterDefaultFirewallPolicies(routerID string) error {
	if routerID == "" {
		return fmt.Errorf("invalid router id")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/ecloud/v2/routers/%s/configure-default-policies", routerID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&RouterNotFoundError{ID: routerID}))
}

// GetRouterTasks retrieves a list of Router tasks
func (s *Service) GetRouterTasks(routerID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetRouterTasksPaginated(routerID, p)
	}, parameters)
}

// GetRouterTasksPaginated retrieves a paginated list of Router tasks
func (s *Service) GetRouterTasksPaginated(routerID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	if routerID == "" {
		return nil, fmt.Errorf("invalid router id")
	}
	body, err := connection.Get[[]Task](s.connection, fmt.Sprintf("/ecloud/v2/routers/%s/tasks", routerID), parameters, connection.NotFoundResponseHandler(&RouterNotFoundError{ID: routerID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetRouterTasksPaginated(routerID, p)
	}), err
}
