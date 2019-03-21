package registrar

import "github.com/ukfast/sdk-go/pkg/connection"

// RegistrarService is an interface for managing the registrar service
type RegistrarService interface {
}

// Service implements RegistrarService for managing
// registrar services via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of Service
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
