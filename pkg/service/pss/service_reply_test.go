package pss

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetReply(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v1/replies/abc", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"abc\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		reply, err := s.GetReply("abc")

		assert.Nil(t, err)
		assert.Equal(t, "abc", reply.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v1/replies/abc", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetReply("abc")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidReplyID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetReply("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid reply id", err.Error())
	})

	t.Run("404_ReturnsReplyNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v1/replies/abc", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetReply("abc")

		assert.NotNil(t, err)
		assert.IsType(t, &ReplyNotFoundError{}, err)
	})
}

func TestDownloadReplyAttachmentStream(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		response := &http.Response{
			Body:       io.NopCloser(bytes.NewReader([]byte("test content"))),
			StatusCode: 200,
		}

		c.EXPECT().Get("/pss/v1/replies/abc/attachments/test.txt", gomock.Any()).Return(&connection.APIResponse{
			Response: response,
		}, nil)

		contentStream, err := s.DownloadReplyAttachmentStream("abc", "test.txt")
		assert.Nil(t, err)

		content, err := io.ReadAll(contentStream)
		assert.Nil(t, err)

		assert.Equal(t, "test content", string(content))
	})

	t.Run("InvalidReplyID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.DownloadReplyAttachmentStream("", "test.txt")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid reply id", err.Error())
	})

	t.Run("InvalidAttachmentName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.DownloadReplyAttachmentStream("abc", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid attachment name", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v1/replies/abc/attachments/test.txt", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.DownloadReplyAttachmentStream("abc", "test.txt")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsAttachmentNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v1/replies/abc/attachments/test.txt", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.DownloadReplyAttachmentStream("abc", "test.txt")

		assert.NotNil(t, err)
		assert.IsType(t, &AttachmentNotFoundError{}, err)
	})
}

func TestUploadReplyAttachmentStream(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		fileStream := io.NopCloser(bytes.NewReader([]byte("test content")))

		c.EXPECT().Post("/pss/v1/replies/abc/attachments/test.txt", fileStream).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"abc\"}}"))),
				StatusCode: 200,
			},
		}, nil)

		err := s.UploadReplyAttachmentStream("abc", "test.txt", fileStream)
		assert.Nil(t, err)
	})

	t.Run("InvalidReplyID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.UploadReplyAttachmentStream("", "test.txt", io.NopCloser(bytes.NewReader([]byte("test content"))))

		assert.NotNil(t, err)
		assert.Equal(t, "invalid reply id", err.Error())
	})

	t.Run("InvalidAttachmentName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.UploadReplyAttachmentStream("abc", "", io.NopCloser(bytes.NewReader([]byte("test content"))))

		assert.NotNil(t, err)
		assert.Equal(t, "invalid attachment name", err.Error())
	})

	t.Run("InvalidStream_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.UploadReplyAttachmentStream("abc", "test.txt", nil)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid stream", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/pss/v1/replies/abc/attachments/test.txt", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		err := s.UploadReplyAttachmentStream("abc", "test.txt", io.NopCloser(bytes.NewReader([]byte("test content"))))

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsAttachmentNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/pss/v1/replies/abc/attachments/test.txt", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"abc\"}}"))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.UploadReplyAttachmentStream("abc", "test.txt", io.NopCloser(bytes.NewReader([]byte("test content"))))

		assert.NotNil(t, err)
		assert.IsType(t, &ReplyNotFoundError{}, err)
	})
}

func TestDeleteReplyAttachment(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Delete("/pss/v1/replies/abc/attachments/test.txt", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.DeleteReplyAttachment("abc", "test.txt")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Delete("/pss/v1/replies/abc/attachments/test.txt", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteReplyAttachment("abc", "test.txt")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidReplyID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		err := s.DeleteReplyAttachment("", "test.txt")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid reply id", err.Error())
	})

	t.Run("InvalidAttachmentName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		err := s.DeleteReplyAttachment("abc", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid attachment name", err.Error())
	})

	t.Run("404_ReturnsReplyAttachmentNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Delete("/pss/v1/replies/abc/attachments/test.txt", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteReplyAttachment("abc", "test.txt")

		assert.NotNil(t, err)
		assert.IsType(t, &AttachmentNotFoundError{}, err)
	})
}
