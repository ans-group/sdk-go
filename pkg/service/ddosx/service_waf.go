package ddosx

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetWAFLogs retrieves a list of ssls
func (s *Service) GetWAFLogs(parameters connection.APIRequestParameters) ([]WAFLog, error) {
	var ssls []WAFLog

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetWAFLogsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, ssl := range response.(*PaginatedWAFLog).Items {
			ssls = append(ssls, ssl)
		}
	}

	return ssls, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetWAFLogsPaginated retrieves a paginated list of ssls
func (s *Service) GetWAFLogsPaginated(parameters connection.APIRequestParameters) (*PaginatedWAFLog, error) {
	body, err := s.getWAFLogsPaginatedResponseBody(parameters)

	return NewPaginatedWAFLog(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetWAFLogsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getWAFLogsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetWAFLogArrayResponseBody, error) {
	body := &GetWAFLogArrayResponseBody{}

	response, err := s.connection.Get("/ddosx/v1/ssls", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetWAFLog retrieves a single ssl by id
func (s *Service) GetWAFLog(sslID string) (WAFLog, error) {
	body, err := s.getWAFLogResponseBody(sslID)

	return body.Data, err
}

func (s *Service) getWAFLogResponseBody(sslID string) (*GetWAFLogResponseBody, error) {
	body := &GetWAFLogResponseBody{}

	if sslID == "" {
		return body, fmt.Errorf("invalid ssl id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/ssls/%s", sslID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &WAFLogNotFoundError{ID: sslID}
		}

		return nil
	})
}
