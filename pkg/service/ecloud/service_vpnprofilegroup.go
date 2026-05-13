package ecloud

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) vpnProfileGroupRes() *resource.Resource[VPNProfileGroup, string] {
	return resource.NewStringResource[VPNProfileGroup](s.connection, "/ecloud/v2/vpn-profile-groups", "vpn profile group", func(id string) error {
		return &VPNProfileGroupNotFoundError{ID: id}
	})
}

// GetVPNProfileGroups retrieves a list of VPN profile groups
func (s *Service) GetVPNProfileGroups(parameters connection.APIRequestParameters) ([]VPNProfileGroup, error) {
	return s.vpnProfileGroupRes().List(parameters)
}

// GetVPNProfileGroupsPaginated retrieves a paginated list of VPN profile groups
func (s *Service) GetVPNProfileGroupsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VPNProfileGroup], error) {
	return s.vpnProfileGroupRes().ListPaginated(parameters)
}

// GetVPNProfileGroup retrieves a single VPN profile group by id
func (s *Service) GetVPNProfileGroup(profileGroupID string) (VPNProfileGroup, error) {
	return s.vpnProfileGroupRes().Get(profileGroupID)
}
