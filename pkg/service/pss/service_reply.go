package pss

import (
	"fmt"
	"io"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetReply retrieves a single reply by id
func (s *Service) GetReply(replyID string) (Reply, error) {
	body, err := s.getReplyResponseBody(replyID)

	return body.Data, err
}

func (s *Service) getReplyResponseBody(replyID string) (*GetReplyResponseBody, error) {
	body := &GetReplyResponseBody{}

	if replyID == "" {
		return body, fmt.Errorf("invalid reply id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/pss/v1/replies/%s", replyID), connection.APIRequestParameters{})
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

// DownloadReplyAttachmentFileStream downloads the provided attachment, returning
// a stream of the file contents and an error
func (s *Service) DownloadReplyAttachmentFileStream(replyID string, attachmentName string) (contentStream io.ReadCloser, err error) {
	response, err := s.downloadReplyAttachmentResponse(replyID, attachmentName)
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

func (s *Service) downloadReplyAttachmentResponse(replyID string, attachmentName string) (*connection.APIResponse, error) {
	body := &connection.APIResponseBody{}
	response := &connection.APIResponse{}

	if replyID == "" {
		return response, fmt.Errorf("invalid reply id")
	}
	if attachmentName == "" {
		return response, fmt.Errorf("invalid attachment name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/pss/v1/replies/%s/attachments/%s", replyID, attachmentName), connection.APIRequestParameters{})
	if err != nil {
		return response, err
	}

	if response.StatusCode == 404 {
		return response, &AttachmentNotFoundError{Name: attachmentName}
	}

	return response, response.ValidateStatusCode([]int{}, body)
}
