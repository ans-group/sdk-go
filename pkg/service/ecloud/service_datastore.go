package ecloud

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) datastoreRes() *resource.Resource[Datastore, int] {
	return resource.NewIntResource[Datastore](s.connection, "/ecloud/v1/datastores", "datastore", func(id int) error {
		return &DatastoreNotFoundError{ID: id}
	})
}

// GetDatastores retrieves a list of datastores
func (s *Service) GetDatastores(parameters connection.APIRequestParameters) ([]Datastore, error) {
	return s.datastoreRes().List(parameters)
}

// GetDatastoresPaginated retrieves a paginated list of datastores
func (s *Service) GetDatastoresPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Datastore], error) {
	return s.datastoreRes().ListPaginated(parameters)
}

// GetDatastore retrieves a single datastore by ID
func (s *Service) GetDatastore(datastoreID int) (Datastore, error) {
	return s.datastoreRes().Get(datastoreID)
}
