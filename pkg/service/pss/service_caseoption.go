package pss

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) changeImpactCaseOptionRes() *resource.Resource[CaseOption, string] {
	return resource.NewStringResource[CaseOption](s.connection, "/pss/v2/case-options/change-impacts", "change-impact-case-option",
		func(id string) error { return fmt.Errorf("case option %s not found", id) })
}

func (s *Service) changeRiskCaseOptionRes() *resource.Resource[CaseOption, string] {
	return resource.NewStringResource[CaseOption](s.connection, "/pss/v2/case-options/change-risks", "change-risk-case-option",
		func(id string) error { return fmt.Errorf("case option %s not found", id) })
}

func (s *Service) incidentImpactCaseOptionRes() *resource.Resource[CaseOption, string] {
	return resource.NewStringResource[CaseOption](s.connection, "/pss/v2/case-options/incident-impacts", "incident-impact-case-option",
		func(id string) error { return fmt.Errorf("case option %s not found", id) })
}

func (s *Service) incidentTypeCaseOptionRes() *resource.Resource[CaseOption, string] {
	return resource.NewStringResource[CaseOption](s.connection, "/pss/v2/case-options/incident-types", "incident-type-case-option",
		func(id string) error { return fmt.Errorf("case option %s not found", id) })
}

// GetCaseOptions retrieves a list of change impact case options
func (s *Service) GetChangeImpactCaseOptions(parameters connection.APIRequestParameters) ([]CaseOption, error) {
	return s.changeImpactCaseOptionRes().List(parameters)
}

// GetCaseOptionsPaginated retrieves a paginated list of change impact case options
func (s *Service) GetChangeImpactCaseOptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseOption], error) {
	return s.changeImpactCaseOptionRes().ListPaginated(parameters)
}

// GetCaseOptions retrieves a list of change risk case options
func (s *Service) GetChangeRiskCaseOptions(parameters connection.APIRequestParameters) ([]CaseOption, error) {
	return s.changeRiskCaseOptionRes().List(parameters)
}

// GetCaseOptionsPaginated retrieves a paginated list of change risk case options
func (s *Service) GetChangeRiskCaseOptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseOption], error) {
	return s.changeRiskCaseOptionRes().ListPaginated(parameters)
}

// GetCaseOptions retrieves a list of incident impact case options
func (s *Service) GetIncidentImpactCaseOptions(parameters connection.APIRequestParameters) ([]CaseOption, error) {
	return s.incidentImpactCaseOptionRes().List(parameters)
}

// GetCaseOptionsPaginated retrieves a paginated list of incident impact case options
func (s *Service) GetIncidentImpactCaseOptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseOption], error) {
	return s.incidentImpactCaseOptionRes().ListPaginated(parameters)
}

// GetCaseOptions retrieves a list of incident type case options
func (s *Service) GetIncidentTypeCaseOptions(parameters connection.APIRequestParameters) ([]CaseOption, error) {
	return s.incidentTypeCaseOptionRes().List(parameters)
}

// GetCaseOptionsPaginated retrieves a paginated list of incident type case options
func (s *Service) GetIncidentTypeCaseOptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseOption], error) {
	return s.incidentTypeCaseOptionRes().ListPaginated(parameters)
}
