package ecloud

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) siteRes() *resource.Resource[Site, int] {
	return resource.NewIntResource[Site](s.connection, "/ecloud/v1/sites", "site", func(id int) error {
		return &SiteNotFoundError{ID: id}
	})
}

// GetSites retrieves a list of sites
func (s *Service) GetSites(parameters connection.APIRequestParameters) ([]Site, error) {
	return s.siteRes().List(parameters)
}

// GetSitesPaginated retrieves a paginated list of sites
func (s *Service) GetSitesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Site], error) {
	return s.siteRes().ListPaginated(parameters)
}

// GetSite retrieves a single site by ID
func (s *Service) GetSite(siteID int) (Site, error) {
	return s.siteRes().Get(siteID)
}
