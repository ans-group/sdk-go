package billing

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetCards retrieves a list of cards
func (s *Service) GetCards(parameters connection.APIRequestParameters) ([]Card, error) {
	var cards []Card

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetCardsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, card := range response.(*PaginatedCard).Items {
			cards = append(cards, card)
		}
	}

	return cards, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetCardsPaginated retrieves a paginated list of cards
func (s *Service) GetCardsPaginated(parameters connection.APIRequestParameters) (*PaginatedCard, error) {
	body, err := s.getCardsPaginatedResponseBody(parameters)

	return NewPaginatedCard(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetCardsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getCardsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetCardSliceResponseBody, error) {
	body := &GetCardSliceResponseBody{}

	response, err := s.connection.Get("/billing/v1/cards", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetCard retrieves a single card by id
func (s *Service) GetCard(cardID int) (Card, error) {
	body, err := s.getCardResponseBody(cardID)

	return body.Data, err
}

func (s *Service) getCardResponseBody(cardID int) (*GetCardResponseBody, error) {
	body := &GetCardResponseBody{}

	if cardID < 1 {
		return body, fmt.Errorf("invalid card id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/billing/v1/cards/%d", cardID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &CardNotFoundError{ID: cardID}
		}

		return nil
	})
}
