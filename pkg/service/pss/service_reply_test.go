package pss

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/ukfast/sdk-go/pkg/connection"
	"github.com/ukfast/sdk-go/test/mocks"
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
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"abc\"}}"))),
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
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetReply("abc")

		assert.NotNil(t, err)
		assert.IsType(t, &ReplyNotFoundError{}, err)
	})
}

func TestDownloadReplyAttachmentFileStream(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		response := &http.Response{
			Body:       ioutil.NopCloser(bytes.NewReader([]byte("test content"))),
			StatusCode: 200,
		}

		c.EXPECT().Get("/pss/v1/replies/abc/attachments/test.txt", gomock.Any()).Return(&connection.APIResponse{
			Response: response,
		}, nil)

		contentStream, err := s.DownloadReplyAttachmentFileStream("abc", "test.txt")
		assert.Nil(t, err)

		content, err := ioutil.ReadAll(contentStream)
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

		_, err := s.DownloadReplyAttachmentFileStream("", "test.txt")

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

		_, err := s.DownloadReplyAttachmentFileStream("abc", "")

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

		_, err := s.DownloadReplyAttachmentFileStream("abc", "test.txt")

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
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.DownloadReplyAttachmentFileStream("abc", "test.txt")

		assert.NotNil(t, err)
		assert.IsType(t, &AttachmentNotFoundError{}, err)
	})
}
