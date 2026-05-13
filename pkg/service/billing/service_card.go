package billing

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) cardRes() *resource.Resource[Card, int] {
	return resource.NewIntResource[Card](s.connection, "/billing/v1/cards", "card",
		func(id int) error { return &CardNotFoundError{ID: id} })
}

// GetCards retrieves a list of cards
func (s *Service) GetCards(parameters connection.APIRequestParameters) ([]Card, error) {
	return s.cardRes().List(parameters)
}

// GetCardsPaginated retrieves a paginated list of cards
func (s *Service) GetCardsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Card], error) {
	return s.cardRes().ListPaginated(parameters)
}

// GetCard retrieves a single card by id
func (s *Service) GetCard(cardID int) (Card, error) {
	return s.cardRes().Get(cardID)
}

// CreateCard creates a new card
func (s *Service) CreateCard(req CreateCardRequest) (int, error) {
	data, err := s.cardRes().Create(&req)
	return data.ID, err
}

// PatchCard patches a card
func (s *Service) PatchCard(cardID int, patch PatchCardRequest) error {
	return s.cardRes().Patch(cardID, &patch)
}

// DeleteCard removes a card
func (s *Service) DeleteCard(cardID int) error {
	return s.cardRes().Delete(cardID)
}
