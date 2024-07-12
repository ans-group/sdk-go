package pss

import (
	"fmt"
	"net/url"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetIncidentCases retrieves a list of incident cases
func (s *Service) GetIncidentCases(parameters connection.APIRequestParameters) ([]IncidentCase, error) {
	return connection.InvokeRequestAll(s.GetIncidentCasesPaginated, parameters)
}

// GetIncidentCasesPaginated retrieves a paginated list of incident cases
func (s *Service) GetIncidentCasesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[IncidentCase], error) {
	body, err := s.getIncidentCasesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetIncidentCasesPaginated), err
}

func (s *Service) getIncidentCasesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]IncidentCase], error) {
	body := &connection.APIResponseBodyData[[]IncidentCase]{}

	request := connection.APIRequest{
		Method:     "GET",
		Resource:   "/pss/v2/cases",
		Query:      url.Values{"case_type": []string{"incident"}},
		Parameters: parameters,
	}

	response, err := s.connection.Invoke(request)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body)
}

// GetIncidentCase retrieves a single instance case by id
func (s *Service) GetIncidentCase(incidentID string) (IncidentCase, error) {
	body, err := s.getIncidentCaseResponseBody(incidentID)

	return body.Data, err
}

func (s *Service) getIncidentCaseResponseBody(incidentID string) (*connection.APIResponseBodyData[IncidentCase], error) {
	if incidentID == "" {
		return &connection.APIResponseBodyData[IncidentCase]{}, fmt.Errorf("invalid incident id")
	}

	return connection.Get[IncidentCase](s.connection, fmt.Sprintf("/pss/v2/cases/%s", incidentID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&IncidentCaseNotFoundError{ID: incidentID}))
}

// GetChangeCases retrieves a list of change cases
func (s *Service) GetChangeCases(parameters connection.APIRequestParameters) ([]ChangeCase, error) {
	return connection.InvokeRequestAll(s.GetChangeCasesPaginated, parameters)
}

// GetChangeCasesPaginated retrieves a paginated list of change cases
func (s *Service) GetChangeCasesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ChangeCase], error) {
	body, err := s.getChangeCasesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetChangeCasesPaginated), err
}

func (s *Service) getChangeCasesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]ChangeCase], error) {
	body := &connection.APIResponseBodyData[[]ChangeCase]{}

	request := connection.APIRequest{
		Method:     "GET",
		Resource:   "/pss/v2/cases",
		Query:      url.Values{"case_type": []string{"change"}},
		Parameters: parameters,
	}

	response, err := s.connection.Invoke(request)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body)
}

// GetProblemCases retrieves a list of problem cases
func (s *Service) GetProblemCases(parameters connection.APIRequestParameters) ([]ProblemCase, error) {
	return connection.InvokeRequestAll(s.GetProblemCasesPaginated, parameters)
}

// GetProblemCasesPaginated retrieves a paginated list of problem cases
func (s *Service) GetProblemCasesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ProblemCase], error) {
	body, err := s.getProblemCasesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetProblemCasesPaginated), err
}

func (s *Service) getProblemCasesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]ProblemCase], error) {
	body := &connection.APIResponseBodyData[[]ProblemCase]{}

	request := connection.APIRequest{
		Method:     "GET",
		Resource:   "/pss/v2/cases",
		Query:      url.Values{"case_type": []string{"problem"}},
		Parameters: parameters,
	}

	response, err := s.connection.Invoke(request)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body)
}
