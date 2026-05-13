package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) vpnGatewaySpecificationRes() *resource.Resource[VPNGatewaySpecification, string] {
	return resource.NewStringResource[VPNGatewaySpecification](s.connection, "/ecloud/v2/vpn-gateway-specifications", "vpn gateway specification", func(id string) error {
		return &VPNGatewaySpecificationNotFoundError{ID: id}
	})
}

// GetVPNGatewaySpecifications retrieves a list of VPN gateway specifications
func (s *Service) GetVPNGatewaySpecifications(parameters connection.APIRequestParameters) ([]VPNGatewaySpecification, error) {
	return s.vpnGatewaySpecificationRes().List(parameters)
}

// GetVPNGatewaySpecificationsPaginated retrieves a paginated list of VPN gateway specifications
func (s *Service) GetVPNGatewaySpecificationsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VPNGatewaySpecification], error) {
	return s.vpnGatewaySpecificationRes().ListPaginated(parameters)
}

// GetVPNGatewaySpecification retrieves a single VPN gateway specification by ID
func (s *Service) GetVPNGatewaySpecification(specificationID string) (VPNGatewaySpecification, error) {
	return s.vpnGatewaySpecificationRes().Get(specificationID)
}

func (s *Service) vpnGatewaySpecAvailabilityZoneRes() *resource.SubResourceList[AvailabilityZone, string] {
	return resource.NewStringSubResourceList[AvailabilityZone](s.connection,
		func(specificationID string) string {
			return fmt.Sprintf("/ecloud/v2/vpn-gateway-specifications/%s/availability-zones", specificationID)
		},
		"vpn gateway specification", "id", func(specificationID string) error {
			return &VPNGatewaySpecificationNotFoundError{ID: specificationID}
		})
}

// GetVPNGatewaySpecificationAvailabilityZones retrieves a list of availability zones for a vpn gateway specification
func (s *Service) GetVPNGatewaySpecificationAvailabilityZones(specificationID string, parameters connection.APIRequestParameters) ([]AvailabilityZone, error) {
	return s.vpnGatewaySpecAvailabilityZoneRes().List(specificationID, parameters)
}

// GetVPNGatewaySpecificationAvailabilityZonesPaginated retrieves a paginated list of availability zones for a vpn gateway specification
func (s *Service) GetVPNGatewaySpecificationAvailabilityZonesPaginated(specificationID string, parameters connection.APIRequestParameters) (*connection.Paginated[AvailabilityZone], error) {
	return s.vpnGatewaySpecAvailabilityZoneRes().ListPaginated(specificationID, parameters)
}
