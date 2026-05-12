package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetDatastores retrieves a list of datastores
func (s *Service) GetDatastores(parameters connection.APIRequestParameters) ([]Datastore, error) {
	return connection.InvokeRequestAll(s.GetDatastoresPaginated, parameters)
}

// GetDatastoresPaginated retrieves a paginated list of datastores
func (s *Service) GetDatastoresPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Datastore], error) {
	body, err := connection.Get[[]Datastore](s.connection, "/ecloud/v1/datastores", parameters)
	return connection.NewPaginated(body, parameters, s.GetDatastoresPaginated), err
}

// GetDatastore retrieves a single datastore by ID
func (s *Service) GetDatastore(datastoreID int) (Datastore, error) {
	if datastoreID < 1 {
		return Datastore{}, fmt.Errorf("invalid datastore id")
	}
	body, err := connection.Get[Datastore](s.connection, fmt.Sprintf("/ecloud/v1/datastores/%d", datastoreID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&DatastoreNotFoundError{ID: datastoreID}))
	return body.Data, err
}
