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

func TestCreateRequest(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		expectedRequest := CreateRequestRequest{
			Author: Author{
				ID: 1,
			},
			Priority: RequestPriorityNormal,
			Subject:  "testsubject",
			Details:  "testdetails",
		}

		c.EXPECT().Post("/pss/v1/requests", &expectedRequest).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":123}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		id, err := s.CreateRequest(expectedRequest)

		assert.Nil(t, err)
		assert.Equal(t, 123, id)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/pss/v1/requests", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateRequest(CreateRequestRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetRequests(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v1/requests", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		requests, err := s.GetRequests(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, requests, 1)
		assert.Equal(t, 123, requests[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/pss/v1/requests", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/pss/v1/requests", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":456}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		requests, err := s.GetRequests(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, requests, 2)
		assert.Equal(t, 123, requests[0].ID)
		assert.Equal(t, 456, requests[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v1/requests", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetRequests(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetRequest(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v1/requests/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":123}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		request, err := s.GetRequest(123)

		assert.Nil(t, err)
		assert.Equal(t, 123, request.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v1/requests/123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetRequest(123)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidRequestID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetRequest(0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid request id", err.Error())
	})

	t.Run("404_ReturnsRequestNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v1/requests/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetRequest(123)

		assert.NotNil(t, err)
		assert.IsType(t, &RequestNotFoundError{}, err)
	})
}

func TestPatchRequest(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		expectedRequest := PatchRequestRequest{
			Priority: RequestPriorityHigh,
		}

		c.EXPECT().Patch("/pss/v1/requests/123", &expectedRequest).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":123}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PatchRequest(123, expectedRequest)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/pss/v1/requests/123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchRequest(123, PatchRequestRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidRequestID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchRequest(0, PatchRequestRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid request id", err.Error())
	})

	t.Run("404_ReturnsRequestNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/pss/v1/requests/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PatchRequest(123, PatchRequestRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &RequestNotFoundError{}, err)
	})
}

func TestCreateRequestReply(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		expectedRequest := CreateReplyRequest{
			Author: Author{
				ID: 456,
			},
			Description: "testdescription",
		}

		c.EXPECT().Post("/pss/v1/requests/123/replies", &expectedRequest).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"abc123\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		id, err := s.CreateRequestReply(123, expectedRequest)

		assert.Nil(t, err)
		assert.Equal(t, "abc123", id)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/pss/v1/requests/123/replies", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateRequestReply(123, CreateReplyRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidRequestID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.CreateRequestReply(0, CreateReplyRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid request id", err.Error())
	})

	t.Run("404_ReturnsRequestNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/pss/v1/requests/123/replies", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.CreateRequestReply(123, CreateReplyRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &RequestNotFoundError{}, err)
	})
}

func TestGetRequestConversation(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v1/requests/123/conversation", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"description\":\"test 1\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		replies, err := s.GetRequestConversation(123, connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Equal(t, "test 1", replies[0].Description)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v1/requests/123/conversation", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetRequestConversation(123, connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidRequestID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetRequestConversation(0, connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid request id", err.Error())
	})

	t.Run("404_ReturnsRequestNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v1/requests/123/conversation", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetRequestConversation(123, connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &RequestNotFoundError{}, err)
	})
}

func TestCreateRequestFeedback(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		expectedRequest := CreateFeedbackRequest{
			ContactID: 1,
		}

		c.EXPECT().Post("/pss/v1/requests/123/feedback", &expectedRequest).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":123}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		id, err := s.CreateRequestFeedback(123, expectedRequest)

		assert.Nil(t, err)
		assert.Equal(t, 123, id)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/pss/v1/requests/123/feedback", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateRequestFeedback(123, CreateFeedbackRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidRequestID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.CreateRequestFeedback(0, CreateFeedbackRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid request id", err.Error())
	})

	t.Run("404_ReturnsRequestNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/pss/v1/requests/123/feedback", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.CreateRequestFeedback(123, CreateFeedbackRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &RequestNotFoundError{}, err)
	})
}
