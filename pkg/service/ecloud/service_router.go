package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetRouters retrieves a list of routers
func (s *Service) GetRouters(parameters connection.APIRequestParameters) ([]Router, error) {
	return connection.InvokeRequestAll(s.GetRoutersPaginated, parameters)
}

// GetRoutersPaginated retrieves a paginated list of routers
func (s *Service) GetRoutersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Router], error) {
	body, err := connection.Get[[]Router](s.connection, "/ecloud/v2/routers", parameters)
	return connection.NewPaginated(body, parameters, s.GetRoutersPaginated), err
}

// GetRouter retrieves a single router by id
func (s *Service) GetRouter(routerID string) (Router, error) {
	if routerID == "" {
		return Router{}, fmt.Errorf("invalid router id")
	}
	body, err := connection.Get[Router](s.connection, fmt.Sprintf("/ecloud/v2/routers/%s", routerID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&RouterNotFoundError{ID: routerID}))
	return body.Data, err
}

// CreateRouter creates a new Router
func (s *Service) CreateRouter(req CreateRouterRequest) (string, error) {
	body, err := connection.Post[Router](s.connection, "/ecloud/v2/routers", &req)
	return body.Data.ID, err
}

// PatchRouter patches a Router
func (s *Service) PatchRouter(routerID string, req PatchRouterRequest) error {
	if routerID == "" {
		return fmt.Errorf("invalid router id")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/ecloud/v2/routers/%s", routerID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&RouterNotFoundError{ID: routerID}))
}

// DeleteRouter deletes a Router
func (s *Service) DeleteRouter(routerID string) error {
	if routerID == "" {
		return fmt.Errorf("invalid router id")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ecloud/v2/routers/%s", routerID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&RouterNotFoundError{ID: routerID}))
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
