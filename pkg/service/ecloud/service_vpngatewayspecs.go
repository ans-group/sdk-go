package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetVPNGatewaySpecifications retrieves a list of VPN gateway specifications
func (s *Service) GetVPNGatewaySpecifications(parameters connection.APIRequestParameters) ([]VPNGatewaySpecification, error) {
	return connection.InvokeRequestAll(s.GetVPNGatewaySpecificationsPaginated, parameters)
}

// GetVPNGatewaySpecificationsPaginated retrieves a paginated list of VPN gateway specifications
func (s *Service) GetVPNGatewaySpecificationsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VPNGatewaySpecification], error) {
	body, err := connection.Get[[]VPNGatewaySpecification](s.connection, "/ecloud/v2/vpn-gateway-specifications", parameters)
	return connection.NewPaginated(body, parameters, s.GetVPNGatewaySpecificationsPaginated), err
}

// GetVPNGatewaySpecification retrieves a single VPN gateway specification by ID
func (s *Service) GetVPNGatewaySpecification(specificationID string) (VPNGatewaySpecification, error) {
	if specificationID == "" {
		return VPNGatewaySpecification{}, fmt.Errorf("invalid vpn gateway specification id")
	}
	body, err := connection.Get[VPNGatewaySpecification](s.connection, fmt.Sprintf("/ecloud/v2/vpn-gateway-specifications/%s", specificationID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&VPNGatewaySpecificationNotFoundError{ID: specificationID}))
	return body.Data, err
}

// GetVPNGatewaySpecificationAvailabilityZones retrieves a list of availability zones for a vpn gateway specification
func (s *Service) GetVPNGatewaySpecificationAvailabilityZones(specificationID string, parameters connection.APIRequestParameters) ([]AvailabilityZone, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[AvailabilityZone], error) {
		return s.GetVPNGatewaySpecificationAvailabilityZonesPaginated(specificationID, p)
	}, parameters)
}

// GetVPNGatewaySpecificationAvailabilityZonesPaginated retrieves a paginated list of availability zones for a vpn gateway specification
func (s *Service) GetVPNGatewaySpecificationAvailabilityZonesPaginated(specificationID string, parameters connection.APIRequestParameters) (*connection.Paginated[AvailabilityZone], error) {
	if specificationID == "" {
		return nil, fmt.Errorf("invalid vpn gateway specification id")
	}
	body, err := connection.Get[[]AvailabilityZone](s.connection, fmt.Sprintf("/ecloud/v2/vpn-gateway-specifications/%s/availability-zones", specificationID), parameters, connection.NotFoundResponseHandler(&VPNGatewaySpecificationNotFoundError{ID: specificationID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[AvailabilityZone], error) {
		return s.GetVPNGatewaySpecificationAvailabilityZonesPaginated(specificationID, p)
	}), err
}
