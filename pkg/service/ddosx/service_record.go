package ddosx

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetRecords retrieves a list of records
func (s *Service) GetRecords(parameters connection.APIRequestParameters) ([]Record, error) {
	var records []Record

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRecordsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, record := range response.(*PaginatedRecords).Records {
			records = append(records, record)
		}
	}

	return records, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetRecordsPaginated retrieves a paginated list of domains
func (s *Service) GetRecordsPaginated(parameters connection.APIRequestParameters) (*PaginatedRecords, error) {
	body, err := s.getRecordsPaginatedResponseBody(parameters)

	return NewPaginatedRecords(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRecordsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getRecordsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetRecordsResponseBody, error) {
	body := &GetRecordsResponseBody{}

	response, err := s.connection.Get("/ddosx/v1/records", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
