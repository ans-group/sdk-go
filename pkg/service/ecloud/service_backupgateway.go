package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetBackupGateways retrieves a list of backup gateways
func (s *Service) GetBackupGateways(parameters connection.APIRequestParameters) ([]BackupGateway, error) {
	return connection.InvokeRequestAll(s.GetBackupGatewaysPaginated, parameters)
}

// GetBackupGatewaysPaginated retrieves a paginated list of backup gateways
func (s *Service) GetBackupGatewaysPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[BackupGateway], error) {
	body, err := connection.Get[[]BackupGateway](s.connection, "/ecloud/v2/backup-gateways", parameters)
	return connection.NewPaginated(body, parameters, s.GetBackupGatewaysPaginated), err
}

// GetBackupGateway retrieves a single backup gateway by ID
func (s *Service) GetBackupGateway(gatewayID string) (BackupGateway, error) {
	if gatewayID == "" {
		return BackupGateway{}, fmt.Errorf("invalid backup gateway id")
	}
	body, err := connection.Get[BackupGateway](s.connection, fmt.Sprintf("/ecloud/v2/backup-gateways/%s", gatewayID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&BackupGatewayNotFoundError{ID: gatewayID}))
	return body.Data, err
}

// CreateBackupGateway creates a new backup gateway
func (s *Service) CreateBackupGateway(req CreateBackupGatewayRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/backup-gateways", &req)
	return body.Data, err
}

// PatchBackupGateway patches a backup gateway
func (s *Service) PatchBackupGateway(gatewayID string, req PatchBackupGatewayRequest) (TaskReference, error) {
	if gatewayID == "" {
		return TaskReference{}, fmt.Errorf("invalid gateway id")
	}
	body, err := connection.Patch[TaskReference](
		s.connection,
		fmt.Sprintf("/ecloud/v2/backup-gateways/%s", gatewayID),
		&req,
		connection.NotFoundResponseHandler(&BackupGatewayNotFoundError{ID: gatewayID}),
	)
	return body.Data, err
}

// DeleteBackupGateway deletes a backup gateway
func (s *Service) DeleteBackupGateway(gatewayID string) (string, error) {
	if gatewayID == "" {
		return "", fmt.Errorf("invalid gateway id")
	}
	body, err := connection.Delete[TaskReference](
		s.connection,
		fmt.Sprintf("/ecloud/v2/backup-gateways/%s", gatewayID),
		nil,
		connection.NotFoundResponseHandler(&BackupGatewayNotFoundError{ID: gatewayID}),
	)
	return body.Data.TaskID, err
}
