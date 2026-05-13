package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) listenerBindRes() *resource.SubResource[Bind, int, int] {
	return resource.NewIntIntSubResource[Bind](s.connection,
		func(listenerID int) string { return fmt.Sprintf("/loadbalancers/v2/listeners/%d/binds", listenerID) },
		"listener", "id", func(listenerID int) error { return &ListenerNotFoundError{ID: listenerID} },
		"bind", "id", func(_, bindID int) error { return &BindNotFoundError{ID: bindID} })
}

// GetListenerBinds retrieves a list of binds
func (s *Service) GetListenerBinds(listenerID int, parameters connection.APIRequestParameters) ([]Bind, error) {
	return s.listenerBindRes().List(listenerID, parameters)
}

// GetListenerBindsPaginated retrieves a paginated list of binds
func (s *Service) GetListenerBindsPaginated(listenerID int, parameters connection.APIRequestParameters) (*connection.Paginated[Bind], error) {
	return s.listenerBindRes().ListPaginated(listenerID, parameters)
}

// GetListenerBind retrieves a single bind by id
func (s *Service) GetListenerBind(listenerID int, bindID int) (Bind, error) {
	return s.listenerBindRes().Get(listenerID, bindID)
}

// CreateListenerBind creates an bind
func (s *Service) CreateListenerBind(listenerID int, req CreateBindRequest) (int, error) {
	bind, err := s.listenerBindRes().Create(listenerID, &req)
	return bind.ID, err
}

// PatchListenerBind patches an bind
func (s *Service) PatchListenerBind(listenerID int, bindID int, req PatchBindRequest) error {
	return s.listenerBindRes().Patch(listenerID, bindID, &req)
}

// DeleteListenerBind deletes a bind
func (s *Service) DeleteListenerBind(listenerID int, bindID int) error {
	return s.listenerBindRes().Delete(listenerID, bindID)
}
