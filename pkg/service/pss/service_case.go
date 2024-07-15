package pss

import (
	"fmt"
	"net/url"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// CreateIncidentCase creates a incident case
func (s *Service) CreateIncidentCase(req CreateIncidentCaseRequest) (string, error) {
	body, err := s.createIncidentCaseResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createIncidentCaseResponseBody(req CreateIncidentCaseRequest) (*connection.APIResponseBodyData[IncidentCase], error) {
	req.CaseType = CaseTypeIncident

	return connection.Post[IncidentCase](s.connection, "/pss/v2/cases", &req)
}

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

// CloseIncidentCase approves a incident case by id
func (s *Service) CloseIncidentCase(incidentID string, req CloseIncidentCaseRequest) (string, error) {
	body, err := s.closeIncidentCaseResponseBody(incidentID, req)

	return body.Data.ID, err
}

func (s *Service) closeIncidentCaseResponseBody(incidentID string, req CloseIncidentCaseRequest) (*connection.APIResponseBodyData[IncidentCase], error) {
	if incidentID == "" {
		return &connection.APIResponseBodyData[IncidentCase]{}, fmt.Errorf("invalid incident id")
	}

	return connection.Post[IncidentCase](s.connection, fmt.Sprintf("/pss/v2/cases/%s/close", incidentID), &req, connection.NotFoundResponseHandler(&IncidentCaseNotFoundError{ID: incidentID}))
}

// CreateChangeCase creates a change case
func (s *Service) CreateChangeCase(req CreateChangeCaseRequest) (string, error) {
	body, err := s.createChangeCaseResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createChangeCaseResponseBody(req CreateChangeCaseRequest) (*connection.APIResponseBodyData[ChangeCase], error) {
	req.CaseType = CaseTypeChange

	return connection.Post[ChangeCase](s.connection, "/pss/v2/cases", &req)
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

// GetChangeCase retrieves a single instance case by id
func (s *Service) GetChangeCase(changeID string) (ChangeCase, error) {
	body, err := s.getChangeCaseResponseBody(changeID)

	return body.Data, err
}

func (s *Service) getChangeCaseResponseBody(changeID string) (*connection.APIResponseBodyData[ChangeCase], error) {
	if changeID == "" {
		return &connection.APIResponseBodyData[ChangeCase]{}, fmt.Errorf("invalid change id")
	}

	return connection.Get[ChangeCase](s.connection, fmt.Sprintf("/pss/v2/cases/%s", changeID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ChangeCaseNotFoundError{ID: changeID}))
}

// ApproveChangeCase approves a change case by id
func (s *Service) ApproveChangeCase(changeID string, req ApproveChangeCaseRequest) (string, error) {
	body, err := s.approveChangeCaseResponseBody(changeID, req)

	return body.Data.ID, err
}

func (s *Service) approveChangeCaseResponseBody(changeID string, req ApproveChangeCaseRequest) (*connection.APIResponseBodyData[ChangeCase], error) {
	if changeID == "" {
		return &connection.APIResponseBodyData[ChangeCase]{}, fmt.Errorf("invalid change id")
	}

	return connection.Post[ChangeCase](s.connection, fmt.Sprintf("/pss/v2/cases/%s/approve", changeID), &req, connection.NotFoundResponseHandler(&ChangeCaseNotFoundError{ID: changeID}))
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

// GetProblemCase retrieves a single instance case by id
func (s *Service) GetProblemCase(problemID string) (ProblemCase, error) {
	body, err := s.getProblemCaseResponseBody(problemID)

	return body.Data, err
}

func (s *Service) getProblemCaseResponseBody(problemID string) (*connection.APIResponseBodyData[ProblemCase], error) {
	if problemID == "" {
		return &connection.APIResponseBodyData[ProblemCase]{}, fmt.Errorf("invalid problem id")
	}

	return connection.Get[ProblemCase](s.connection, fmt.Sprintf("/pss/v2/cases/%s", problemID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ProblemCaseNotFoundError{ID: problemID}))
}

// GetCaseUpdates retrieves a list of problem case updates
func (s *Service) GetCaseUpdates(caseID string, parameters connection.APIRequestParameters) ([]CaseUpdate, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[CaseUpdate], error) {
		return s.GetCaseUpdatesPaginated(caseID, p)
	}, parameters)
}

// GetCaseUpdatesPaginated retrieves a paginated list of case updates
func (s *Service) GetCaseUpdatesPaginated(caseID string, parameters connection.APIRequestParameters) (*connection.Paginated[CaseUpdate], error) {
	body, err := s.getCaseUpdatesPaginatedResponseBody(caseID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[CaseUpdate], error) {
		return s.GetCaseUpdatesPaginated(caseID, p)
	}), err
}

func (s *Service) getCaseUpdatesPaginatedResponseBody(caseID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]CaseUpdate], error) {
	if caseID == "" {
		return &connection.APIResponseBodyData[[]CaseUpdate]{}, fmt.Errorf("invalid case id")
	}

	return connection.Get[[]CaseUpdate](s.connection, fmt.Sprintf("/pss/v2/cases/%s/updates", caseID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ChangeCaseNotFoundError{ID: changeID}))

}
