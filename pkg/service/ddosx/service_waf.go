package ddosx

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) wafLogRes() *resource.Resource[WAFLog, string] {
	return resource.NewStringResource[WAFLog](s.connection, "/ddosx/v1/waf/logs", "request",
		func(id string) error { return &WAFLogNotFoundError{ID: id} })
}

func (s *Service) wafLogMatchRes() *resource.Resource[WAFLogMatch, string] {
	return resource.NewStringResource[WAFLogMatch](s.connection, "/ddosx/v1/waf/logs/matches", "match",
		func(id string) error { return &WAFLogMatchNotFoundError{ID: id} })
}

// GetWAFLogs retrieves a list of logs
func (s *Service) GetWAFLogs(parameters connection.APIRequestParameters) ([]WAFLog, error) {
	return s.wafLogRes().List(parameters)
}

// GetWAFLogsPaginated retrieves a paginated list of logs
func (s *Service) GetWAFLogsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[WAFLog], error) {
	return s.wafLogRes().ListPaginated(parameters)
}

// GetWAFLog retrieves a single log by id
func (s *Service) GetWAFLog(requestID string) (WAFLog, error) {
	return s.wafLogRes().Get(requestID)
}

// GetWAFLogMatches retrieves a list of log matches
func (s *Service) GetWAFLogMatches(parameters connection.APIRequestParameters) ([]WAFLogMatch, error) {
	return s.wafLogMatchRes().List(parameters)
}

// GetWAFLogMatchesPaginated retrieves a paginated list of log matches
func (s *Service) GetWAFLogMatchesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[WAFLogMatch], error) {
	return s.wafLogMatchRes().ListPaginated(parameters)
}

// GetWAFLogRequestMatches retrieves a list of log matches for request
func (s *Service) GetWAFLogRequestMatches(requestID string, parameters connection.APIRequestParameters) ([]WAFLogMatch, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[WAFLogMatch], error) {
		return s.GetWAFLogRequestMatchesPaginated(requestID, p)
	}, parameters)
}

// GetWAFLogRequestMatchesPaginated retrieves a paginated list of matches for request
func (s *Service) GetWAFLogRequestMatchesPaginated(requestID string, parameters connection.APIRequestParameters) (*connection.Paginated[WAFLogMatch], error) {
	if requestID == "" {
		return nil, fmt.Errorf("invalid request id")
	}
	body, err := connection.Get[[]WAFLogMatch](s.connection, fmt.Sprintf("/ddosx/v1/waf/logs/%s/matches", requestID), parameters, connection.NotFoundResponseHandler(&WAFLogNotFoundError{ID: requestID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[WAFLogMatch], error) {
		return s.GetWAFLogRequestMatchesPaginated(requestID, p)
	}), err
}

// GetWAFLogRequestMatch retrieves a single waf log request match
func (s *Service) GetWAFLogRequestMatch(requestID string, matchID string) (WAFLogMatch, error) {
	if requestID == "" {
		return WAFLogMatch{}, fmt.Errorf("invalid request id")
	}
	if matchID == "" {
		return WAFLogMatch{}, fmt.Errorf("invalid match id")
	}
	body, err := connection.Get[WAFLogMatch](s.connection, fmt.Sprintf("/ddosx/v1/waf/logs/%s/matches/%s", requestID, matchID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&WAFLogMatchNotFoundError{ID: requestID}))
	return body.Data, err
}
