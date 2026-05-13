package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetBackupGatewaySpecifications retrieves a list of Backup gateway specifications
func (s *Service) GetBackupGatewaySpecifications(parameters connection.APIRequestParameters) ([]BackupGatewaySpecification, error) {
	return connection.InvokeRequestAll(s.GetBackupGatewaySpecificationsPaginated, parameters)
}

// GetBackupGatewaySpecificationsPaginated retrieves a paginated list of Backup gateway specifications
func (s *Service) GetBackupGatewaySpecificationsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[BackupGatewaySpecification], error) {
	body, err := connection.Get[[]BackupGatewaySpecification](s.connection, "/ecloud/v2/backup-gateway-specs", parameters)
	return connection.NewPaginated(body, parameters, s.GetBackupGatewaySpecificationsPaginated), err
}

// GetBackupGatewaySpecification retrieves a single Backup gateway specification by ID
func (s *Service) GetBackupGatewaySpecification(specificationID string) (BackupGatewaySpecification, error) {
	if specificationID == "" {
		return BackupGatewaySpecification{}, fmt.Errorf("invalid backup gateway specification id")
	}
	body, err := connection.Get[BackupGatewaySpecification](
		s.connection,
		fmt.Sprintf("/ecloud/v2/backup-gateway-specs/%s", specificationID),
		connection.APIRequestParameters{},
		connection.NotFoundResponseHandler(&BackupGatewaySpecificationNotFoundError{ID: specificationID}),
	)
	return body.Data, err
}
