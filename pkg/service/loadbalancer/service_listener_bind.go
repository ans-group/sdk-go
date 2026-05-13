package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetListenerBinds retrieves a list of binds
func (s *Service) GetListenerBinds(listenerID int, parameters connection.APIRequestParameters) ([]Bind, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Bind], error) {
		return s.GetListenerBindsPaginated(listenerID, p)
	}, parameters)
}

// GetListenerBindsPaginated retrieves a paginated list of binds
func (s *Service) GetListenerBindsPaginated(listenerID int, parameters connection.APIRequestParameters) (*connection.Paginated[Bind], error) {
	if listenerID < 1 {
		return nil, fmt.Errorf("invalid listener id")
	}
	body, err := connection.Get[[]Bind](s.connection, fmt.Sprintf("/loadbalancers/v2/listeners/%d/binds", listenerID), parameters)
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Bind], error) {
		return s.GetListenerBindsPaginated(listenerID, p)
	}), err
}

// GetListenerBind retrieves a single bind by id
func (s *Service) GetListenerBind(listenerID int, bindID int) (Bind, error) {
	if listenerID < 1 {
		return Bind{}, fmt.Errorf("invalid listener id")
	}
	if bindID < 1 {
		return Bind{}, fmt.Errorf("invalid bind id")
	}
	body, err := connection.Get[Bind](s.connection, fmt.Sprintf("/loadbalancers/v2/listeners/%d/binds/%d", listenerID, bindID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&BindNotFoundError{ID: listenerID}))
	return body.Data, err
}

// CreateListenerBind creates an bind
func (s *Service) CreateListenerBind(listenerID int, req CreateBindRequest) (int, error) {
	if listenerID < 1 {
		return 0, fmt.Errorf("invalid listener id")
	}
	body, err := connection.Post[Bind](s.connection, fmt.Sprintf("/loadbalancers/v2/listeners/%d/binds", listenerID), &req, connection.NotFoundResponseHandler(&BindNotFoundError{ID: listenerID}))
	return body.Data.ID, err
}

// PatchListenerBind patches an bind
func (s *Service) PatchListenerBind(listenerID int, bindID int, req PatchBindRequest) error {
	if listenerID < 1 {
		return fmt.Errorf("invalid listener id")
	}
	if bindID < 1 {
		return fmt.Errorf("invalid bind id")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/loadbalancers/v2/listeners/%d/binds/%d", listenerID, bindID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&BindNotFoundError{ID: listenerID}))
}

// DeleteListenerBind deletes a bind
func (s *Service) DeleteListenerBind(listenerID int, bindID int) error {
	if listenerID < 1 {
		return fmt.Errorf("invalid listener id")
	}
	if bindID < 1 {
		return fmt.Errorf("invalid bind id")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/loadbalancers/v2/listeners/%d/binds/%d", listenerID, bindID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&BindNotFoundError{ID: listenerID}))
}
