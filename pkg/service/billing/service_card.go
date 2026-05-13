package billing

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetCards retrieves a list of cards
func (s *Service) GetCards(parameters connection.APIRequestParameters) ([]Card, error) {
	return connection.InvokeRequestAll(s.GetCardsPaginated, parameters)
}

// GetCardsPaginated retrieves a paginated list of cards
func (s *Service) GetCardsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Card], error) {
	body, err := connection.Get[[]Card](s.connection, "/billing/v1/cards", parameters)
	return connection.NewPaginated(body, parameters, s.GetCardsPaginated), err
}

// GetCard retrieves a single card by id
func (s *Service) GetCard(cardID int) (Card, error) {
	if cardID < 1 {
		return Card{}, fmt.Errorf("invalid card id")
	}
	body, err := connection.Get[Card](s.connection, fmt.Sprintf("/billing/v1/cards/%d", cardID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&CardNotFoundError{ID: cardID}))
	return body.Data, err
}

// CreateCard creates a new card
func (s *Service) CreateCard(req CreateCardRequest) (int, error) {
	body, err := connection.Post[Card](s.connection, "/billing/v1/cards", &req)
	return body.Data.ID, err
}

// PatchCard patches a card
func (s *Service) PatchCard(cardID int, patch PatchCardRequest) error {
	if cardID < 1 {
		return fmt.Errorf("invalid card id")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/billing/v1/cards/%d", cardID), &patch, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&CardNotFoundError{ID: cardID}))
}

// DeleteCard removes a card
func (s *Service) DeleteCard(cardID int) error {
	if cardID < 1 {
		return fmt.Errorf("invalid card id")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/billing/v1/cards/%d", cardID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&CardNotFoundError{ID: cardID}))
}
