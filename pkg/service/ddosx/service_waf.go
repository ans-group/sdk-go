package ddosx

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetWAFLogs retrieves a list of logs
func (s *Service) GetWAFLogs(parameters connection.APIRequestParameters) ([]WAFLog, error) {
	return connection.InvokeRequestAll(s.GetWAFLogsPaginated, parameters)
}

// GetWAFLogsPaginated retrieves a paginated list of logs
func (s *Service) GetWAFLogsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[WAFLog], error) {
	body, err := connection.Get[[]WAFLog](s.connection, "/ddosx/v1/waf/logs", parameters)
	return connection.NewPaginated(body, parameters, s.GetWAFLogsPaginated), err
}

// GetWAFLog retrieves a single log by id
func (s *Service) GetWAFLog(requestID string) (WAFLog, error) {
	if requestID == "" {
		return WAFLog{}, fmt.Errorf("invalid request id")
	}
	body, err := connection.Get[WAFLog](s.connection, fmt.Sprintf("/ddosx/v1/waf/logs/%s", requestID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&WAFLogNotFoundError{ID: requestID}))
	return body.Data, err
}

// GetWAFLogMatches retrieves a list of log matches
func (s *Service) GetWAFLogMatches(parameters connection.APIRequestParameters) ([]WAFLogMatch, error) {
	return connection.InvokeRequestAll(s.GetWAFLogMatchesPaginated, parameters)
}

// GetWAFLogMatchesPaginated retrieves a paginated list of log matches
func (s *Service) GetWAFLogMatchesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[WAFLogMatch], error) {
	body, err := connection.Get[[]WAFLogMatch](s.connection, "/ddosx/v1/waf/logs/matches", parameters)
	return connection.NewPaginated(body, parameters, s.GetWAFLogMatchesPaginated), err
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
