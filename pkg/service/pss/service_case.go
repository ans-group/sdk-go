package pss

import (
	"fmt"
	"net/url"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) caseRes() *resource.Resource[Case, string] {
	return resource.NewStringResource[Case](s.connection, "/pss/v2/cases", "case",
		func(id string) error { return &CaseNotFoundError{ID: id} })
}

func (s *Service) incidentCaseRes() *resource.Resource[IncidentCase, string] {
	return resource.NewStringResource[IncidentCase](s.connection, "/pss/v2/cases", "incident",
		func(id string) error { return &CaseNotFoundError{ID: id} })
}

func (s *Service) changeCaseRes() *resource.Resource[ChangeCase, string] {
	return resource.NewStringResource[ChangeCase](s.connection, "/pss/v2/cases", "change",
		func(id string) error { return &CaseNotFoundError{ID: id} })
}

func (s *Service) problemCaseRes() *resource.Resource[ProblemCase, string] {
	return resource.NewStringResource[ProblemCase](s.connection, "/pss/v2/cases", "problem",
		func(id string) error { return &CaseNotFoundError{ID: id} })
}

// GetCases retrieves a list of cases
func (s *Service) GetCases(parameters connection.APIRequestParameters) ([]Case, error) {
	return s.caseRes().List(parameters)
}

// GetCasesPaginated retrieves a paginated list of cases
func (s *Service) GetCasesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Case], error) {
	return s.caseRes().ListPaginated(parameters)
}

// GetCase retrieves a single instance case by id
func (s *Service) GetCase(caseID string) (Case, error) {
	return s.caseRes().Get(caseID)
}

// CreateIncidentCase creates a incident case
func (s *Service) CreateIncidentCase(req CreateIncidentCaseRequest) (string, error) {
	req.CaseType = CaseTypeIncident
	body, err := connection.Post[IncidentCase](s.connection, "/pss/v2/cases", &req)
	return body.Data.ID, err
}

// GetIncidentCases retrieves a list of incident cases
func (s *Service) GetIncidentCases(parameters connection.APIRequestParameters) ([]IncidentCase, error) {
	return connection.InvokeRequestAll(s.GetIncidentCasesPaginated, parameters)
}

// GetIncidentCasesPaginated retrieves a paginated list of incident cases
func (s *Service) GetIncidentCasesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[IncidentCase], error) {
	body := &connection.APIResponseBodyData[[]IncidentCase]{}
	request := connection.APIRequest{
		Method:     "GET",
		Resource:   "/pss/v2/cases",
		Query:      url.Values{"case_type": []string{"incident"}},
		Parameters: parameters,
	}
	response, err := s.connection.Invoke(request)
	if err != nil {
		return connection.NewPaginated(body, parameters, s.GetIncidentCasesPaginated), err
	}
	return connection.NewPaginated(body, parameters, s.GetIncidentCasesPaginated), response.HandleResponse(body)
}

// GetIncidentCase retrieves a single instance case by id
func (s *Service) GetIncidentCase(incidentID string) (IncidentCase, error) {
	return s.incidentCaseRes().Get(incidentID)
}

// CloseIncidentCase approves a incident case by id
func (s *Service) CloseIncidentCase(incidentID string, req CloseIncidentCaseRequest) (string, error) {
	if incidentID == "" {
		return "", fmt.Errorf("invalid incident id")
	}
	body, err := connection.Post[IncidentCase](s.connection, fmt.Sprintf("/pss/v2/cases/%s/close", incidentID), &req, connection.NotFoundResponseHandler(&CaseNotFoundError{ID: incidentID}))
	return body.Data.ID, err
}

// CreateChangeCase creates a change case
func (s *Service) CreateChangeCase(req CreateChangeCaseRequest) (string, error) {
	req.CaseType = CaseTypeChange
	body, err := connection.Post[ChangeCase](s.connection, "/pss/v2/cases", &req)
	return body.Data.ID, err
}

// GetChangeCases retrieves a list of change cases
func (s *Service) GetChangeCases(parameters connection.APIRequestParameters) ([]ChangeCase, error) {
	return connection.InvokeRequestAll(s.GetChangeCasesPaginated, parameters)
}

// GetChangeCasesPaginated retrieves a paginated list of change cases
func (s *Service) GetChangeCasesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ChangeCase], error) {
	body := &connection.APIResponseBodyData[[]ChangeCase]{}
	request := connection.APIRequest{
		Method:     "GET",
		Resource:   "/pss/v2/cases",
		Query:      url.Values{"case_type": []string{"change"}},
		Parameters: parameters,
	}
	response, err := s.connection.Invoke(request)
	if err != nil {
		return connection.NewPaginated(body, parameters, s.GetChangeCasesPaginated), err
	}
	return connection.NewPaginated(body, parameters, s.GetChangeCasesPaginated), response.HandleResponse(body)
}

// GetChangeCase retrieves a single instance case by id
func (s *Service) GetChangeCase(changeID string) (ChangeCase, error) {
	return s.changeCaseRes().Get(changeID)
}

// ApproveChangeCase approves a change case by id
func (s *Service) ApproveChangeCase(changeID string, req ApproveChangeCaseRequest) (string, error) {
	if changeID == "" {
		return "", fmt.Errorf("invalid change id")
	}
	body, err := connection.Post[ChangeCase](s.connection, fmt.Sprintf("/pss/v2/cases/%s/approve", changeID), &req, connection.NotFoundResponseHandler(&CaseNotFoundError{ID: changeID}))
	return body.Data.ID, err
}

// GetProblemCases retrieves a list of problem cases
func (s *Service) GetProblemCases(parameters connection.APIRequestParameters) ([]ProblemCase, error) {
	return connection.InvokeRequestAll(s.GetProblemCasesPaginated, parameters)
}

// GetProblemCasesPaginated retrieves a paginated list of problem cases
func (s *Service) GetProblemCasesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ProblemCase], error) {
	body := &connection.APIResponseBodyData[[]ProblemCase]{}
	request := connection.APIRequest{
		Method:     "GET",
		Resource:   "/pss/v2/cases",
		Query:      url.Values{"case_type": []string{"problem"}},
		Parameters: parameters,
	}
	response, err := s.connection.Invoke(request)
	if err != nil {
		return connection.NewPaginated(body, parameters, s.GetProblemCasesPaginated), err
	}
	return connection.NewPaginated(body, parameters, s.GetProblemCasesPaginated), response.HandleResponse(body)
}

// GetProblemCase retrieves a single instance case by id
func (s *Service) GetProblemCase(problemID string) (ProblemCase, error) {
	return s.problemCaseRes().Get(problemID)
}

// GetCaseUpdates retrieves a list of problem case updates
func (s *Service) GetCaseUpdates(caseID string, parameters connection.APIRequestParameters) ([]CaseUpdate, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[CaseUpdate], error) {
		return s.GetCaseUpdatesPaginated(caseID, p)
	}, parameters)
}

// GetCaseUpdatesPaginated retrieves a paginated list of case updates
func (s *Service) GetCaseUpdatesPaginated(caseID string, parameters connection.APIRequestParameters) (*connection.Paginated[CaseUpdate], error) {
	if caseID == "" {
		return nil, fmt.Errorf("invalid case id")
	}
	body, err := connection.Get[[]CaseUpdate](s.connection, fmt.Sprintf("/pss/v2/cases/%s/updates", caseID), parameters, connection.NotFoundResponseHandler(&CaseNotFoundError{ID: caseID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[CaseUpdate], error) {
		return s.GetCaseUpdatesPaginated(caseID, p)
	}), err
}

// GetProblemCase retrieves a single instance case by id
func (s *Service) GetCaseUpdate(caseID string, updateID string) (CaseUpdate, error) {
	if caseID == "" {
		return CaseUpdate{}, fmt.Errorf("invalid case id")
	}
	if updateID == "" {
		return CaseUpdate{}, fmt.Errorf("invalid case update id")
	}
	body, err := connection.Get[CaseUpdate](s.connection, fmt.Sprintf("/pss/v2/cases/%s/updates/%s", caseID, updateID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&CaseUpdateNotFoundError{ID: updateID}))
	return body.Data, err
}

// CreateCaseUpdate creates a change case
func (s *Service) CreateCaseUpdate(caseID string, req CreateCaseUpdateRequest) (string, error) {
	if caseID == "" {
		return "", fmt.Errorf("invalid case id")
	}
	body, err := connection.Post[CaseUpdate](s.connection, fmt.Sprintf("/pss/v2/cases/%s/updates", caseID), &req)
	return body.Data.ID, err
}
