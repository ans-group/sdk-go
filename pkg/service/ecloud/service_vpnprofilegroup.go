package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetVPNProfileGroups retrieves a list of VPN profile groups
func (s *Service) GetVPNProfileGroups(parameters connection.APIRequestParameters) ([]VPNProfileGroup, error) {
	return connection.InvokeRequestAll(s.GetVPNProfileGroupsPaginated, parameters)
}

// GetVPNProfileGroupsPaginated retrieves a paginated list of VPN profile groups
func (s *Service) GetVPNProfileGroupsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VPNProfileGroup], error) {
	body, err := connection.Get[[]VPNProfileGroup](s.connection, "/ecloud/v2/vpn-profile-groups", parameters)
	return connection.NewPaginated(body, parameters, s.GetVPNProfileGroupsPaginated), err
}

// GetVPNProfileGroup retrieves a single VPN profile group by id
func (s *Service) GetVPNProfileGroup(profileGroupID string) (VPNProfileGroup, error) {
	if profileGroupID == "" {
		return VPNProfileGroup{}, fmt.Errorf("invalid vpn profile group id")
	}
	body, err := connection.Get[VPNProfileGroup](s.connection, fmt.Sprintf("/ecloud/v2/vpn-profile-groups/%s", profileGroupID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&VPNProfileGroupNotFoundError{ID: profileGroupID}))
	return body.Data, err
}
