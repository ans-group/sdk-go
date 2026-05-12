package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetSites retrieves a list of sites
func (s *Service) GetSites(parameters connection.APIRequestParameters) ([]Site, error) {
	return connection.InvokeRequestAll(s.GetSitesPaginated, parameters)
}

// GetSitesPaginated retrieves a paginated list of sites
func (s *Service) GetSitesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Site], error) {
	body, err := connection.Get[[]Site](s.connection, "/ecloud/v1/sites", parameters)
	return connection.NewPaginated(body, parameters, s.GetSitesPaginated), err
}

// GetSite retrieves a single site by ID
func (s *Service) GetSite(siteID int) (Site, error) {
	if siteID < 1 {
		return Site{}, fmt.Errorf("invalid site id")
	}
	body, err := connection.Get[Site](s.connection, fmt.Sprintf("/ecloud/v1/sites/%d", siteID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&SiteNotFoundError{ID: siteID}))
	return body.Data, err
}
