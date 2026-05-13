package ecloud

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) backupGatewaySpecificationRes() *resource.Resource[BackupGatewaySpecification, string] {
	return resource.NewStringResource[BackupGatewaySpecification](s.connection, "/ecloud/v2/backup-gateway-specs", "backup gateway specification", func(id string) error {
		return &BackupGatewaySpecificationNotFoundError{ID: id}
	})
}

// GetBackupGatewaySpecifications retrieves a list of Backup gateway specifications
func (s *Service) GetBackupGatewaySpecifications(parameters connection.APIRequestParameters) ([]BackupGatewaySpecification, error) {
	return s.backupGatewaySpecificationRes().List(parameters)
}

// GetBackupGatewaySpecificationsPaginated retrieves a paginated list of Backup gateway specifications
func (s *Service) GetBackupGatewaySpecificationsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[BackupGatewaySpecification], error) {
	return s.backupGatewaySpecificationRes().ListPaginated(parameters)
}

// GetBackupGatewaySpecification retrieves a single Backup gateway specification by ID
func (s *Service) GetBackupGatewaySpecification(specificationID string) (BackupGatewaySpecification, error) {
	return s.backupGatewaySpecificationRes().Get(specificationID)
}
