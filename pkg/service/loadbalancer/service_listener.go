package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) listenerRes() *resource.Resource[Listener, int] {
	return resource.NewIntResource[Listener](s.connection, "/loadbalancers/v2/listeners", "listener",
		func(id int) error { return &ListenerNotFoundError{ID: id} })
}

// GetListeners retrieves a list of listeners
func (s *Service) GetListeners(parameters connection.APIRequestParameters) ([]Listener, error) {
	return s.listenerRes().List(parameters)
}

// GetListenersPaginated retrieves a paginated list of listeners
func (s *Service) GetListenersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Listener], error) {
	return s.listenerRes().ListPaginated(parameters)
}

// GetListener retrieves a single listener by id
func (s *Service) GetListener(listenerID int) (Listener, error) {
	return s.listenerRes().Get(listenerID)
}

// CreateListener creates a listener
func (s *Service) CreateListener(req CreateListenerRequest) (int, error) {
	data, err := s.listenerRes().Create(&req)
	return data.ID, err
}

// PatchListener patches a listener
func (s *Service) PatchListener(listenerID int, req PatchListenerRequest) error {
	return s.listenerRes().Patch(listenerID, &req)
}

// DisableListenerGeoIP patches a listener
func (s *Service) DisableListenerGeoIP(listenerID int) error {
	if listenerID < 1 {
		return fmt.Errorf("invalid listener id")
	}
	disableGeoIPReq := struct {
		GeoIP interface{} `json:"geoip"`
	}{}
	_, err := connection.Patch[interface{}](s.connection, fmt.Sprintf("/loadbalancers/v2/listeners/%d", listenerID), &disableGeoIPReq, connection.NotFoundResponseHandler(&ListenerNotFoundError{ID: listenerID}))
	return err
}

// DeleteListener deletes a listener
func (s *Service) DeleteListener(listenerID int) error {
	return s.listenerRes().Delete(listenerID)
}
