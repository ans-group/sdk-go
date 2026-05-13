package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetListeners retrieves a list of listeners
func (s *Service) GetListeners(parameters connection.APIRequestParameters) ([]Listener, error) {
	return connection.InvokeRequestAll(s.GetListenersPaginated, parameters)
}

// GetListenersPaginated retrieves a paginated list of listeners
func (s *Service) GetListenersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Listener], error) {
	body, err := connection.Get[[]Listener](s.connection, "/loadbalancers/v2/listeners", parameters)
	return connection.NewPaginated(body, parameters, s.GetListenersPaginated), err
}

// GetListener retrieves a single listener by id
func (s *Service) GetListener(listenerID int) (Listener, error) {
	if listenerID < 1 {
		return Listener{}, fmt.Errorf("invalid listener id")
	}
	body, err := connection.Get[Listener](s.connection, fmt.Sprintf("/loadbalancers/v2/listeners/%d", listenerID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ListenerNotFoundError{ID: listenerID}))
	return body.Data, err
}

// CreateListener creates a listener
func (s *Service) CreateListener(req CreateListenerRequest) (int, error) {
	body, err := connection.Post[Listener](s.connection, "/loadbalancers/v2/listeners", &req)
	return body.Data.ID, err
}

// PatchListener patches a listener
func (s *Service) PatchListener(listenerID int, req PatchListenerRequest) error {
	if listenerID < 1 {
		return fmt.Errorf("invalid listener id")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/loadbalancers/v2/listeners/%d", listenerID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&ListenerNotFoundError{ID: listenerID}))
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
	if listenerID < 1 {
		return fmt.Errorf("invalid listener id")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/loadbalancers/v2/listeners/%d", listenerID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&ListenerNotFoundError{ID: listenerID}))
}
