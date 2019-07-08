package pss

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetReply retrieves a single reply by id
func (s *Service) GetReply(replyID int) (Reply, error) {
	body, err := s.getReplyResponseBody(replyID)

	return body.Data, err
}

func (s *Service) getReplyResponseBody(replyID int) (*GetReplyResponseBody, error) {
	body := &GetReplyResponseBody{}

	if replyID < 1 {
		return body, fmt.Errorf("invalid reply id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/pss/v1/replies/%d", replyID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ReplyNotFoundError{ID: replyID}
		}

		return nil
	})
}
