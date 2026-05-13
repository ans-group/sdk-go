package ddosx

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) recordRes() *resource.Resource[Record, string] {
	return resource.NewStringResource[Record](s.connection, "/ddosx/v1/records", "record",
		func(id string) error { return &RecordNotFoundError{ID: id} })
}

// GetRecords retrieves a list of records
func (s *Service) GetRecords(parameters connection.APIRequestParameters) ([]Record, error) {
	return s.recordRes().List(parameters)
}

// GetRecordsPaginated retrieves a paginated list of domains
func (s *Service) GetRecordsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Record], error) {
	return s.recordRes().ListPaginated(parameters)
}
