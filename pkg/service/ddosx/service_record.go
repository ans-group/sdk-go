package ddosx

import (
	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetRecords retrieves a list of records
func (s *Service) GetRecords(parameters connection.APIRequestParameters) ([]Record, error) {
	return connection.InvokeRequestAll(s.GetRecordsPaginated, parameters)
}

// GetRecordsPaginated retrieves a paginated list of domains
func (s *Service) GetRecordsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Record], error) {
	body, err := connection.Get[[]Record](s.connection, "/ddosx/v1/records", parameters)
	return connection.NewPaginated(body, parameters, s.GetRecordsPaginated), err
}
