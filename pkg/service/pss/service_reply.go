package pss

import (
	"fmt"
	"io"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetReply retrieves a single reply by id
func (s *Service) GetReply(replyID string) (Reply, error) {
	if replyID == "" {
		return Reply{}, fmt.Errorf("invalid reply id")
	}
	body, err := connection.Get[Reply](s.connection, fmt.Sprintf("/pss/v1/replies/%s", replyID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ReplyNotFoundError{ID: replyID}))
	return body.Data, err
}

// DownloadReplyAttachmentStream downloads the provided attachment, returning
// a stream of the file contents and an error
func (s *Service) DownloadReplyAttachmentStream(replyID string, attachmentName string) (contentStream io.ReadCloser, err error) {
	response, err := s.downloadReplyAttachmentResponse(replyID, attachmentName)
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

func (s *Service) downloadReplyAttachmentResponse(replyID string, attachmentName string) (*connection.APIResponse, error) {
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

	return response, response.HandleResponse(nil)
}

// UploadReplyAttachmentStream uploads the provided attachment
func (s *Service) UploadReplyAttachmentStream(replyID string, attachmentName string, stream io.Reader) (err error) {
	if replyID == "" {
		return fmt.Errorf("invalid reply id")
	}
	if attachmentName == "" {
		return fmt.Errorf("invalid attachment name")
	}
	if stream == nil {
		return fmt.Errorf("invalid stream")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/pss/v1/replies/%s/attachments/%s", replyID, attachmentName), stream, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&ReplyNotFoundError{ID: replyID}))
}

// DeleteReplyAttachment removes a reply attachment
func (s *Service) DeleteReplyAttachment(replyID string, attachmentName string) error {
	if replyID == "" {
		return fmt.Errorf("invalid reply id")
	}
	if attachmentName == "" {
		return fmt.Errorf("invalid attachment name")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/pss/v1/replies/%s/attachments/%s", replyID, attachmentName), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&AttachmentNotFoundError{Name: attachmentName}))
}
