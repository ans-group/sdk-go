package pss

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) caseCategoryRes() *resource.Resource[CaseCategory, string] {
	return resource.NewStringResource[CaseCategory](s.connection, "/pss/v2/case-categories", "case-category",
		func(id string) error { return fmt.Errorf("case category %s not found", id) })
}

// GetCaseCategories retrieves a list of case categories
func (s *Service) GetCaseCategories(parameters connection.APIRequestParameters) ([]CaseCategory, error) {
	return s.caseCategoryRes().List(parameters)
}

// GetCaseCategoriesPaginated retrieves a paginated list of case categories
func (s *Service) GetCaseCategoriesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseCategory], error) {
	return s.caseCategoryRes().ListPaginated(parameters)
}
