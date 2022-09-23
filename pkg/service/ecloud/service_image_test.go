package ecloud

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetImages(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/images", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"img-abcdef12\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		images, err := s.GetImages(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, images, 1)
		assert.Equal(t, "img-abcdef12", images[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ecloud/v2/images", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"img-abcdef12\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ecloud/v2/images", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"img-abcdef23\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		images, err := s.GetImages(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, images, 2)
		assert.Equal(t, "img-abcdef12", images[0].ID)
		assert.Equal(t, "img-abcdef23", images[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/images", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetImages(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetImage(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/images/img-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"img-abcdef12\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		image, err := s.GetImage("img-abcdef12")

		assert.Nil(t, err)
		assert.Equal(t, "img-abcdef12", image.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/images/img-abcdef12", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetImage("img-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidImageID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetImage("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid image id", err.Error())
	})

	t.Run("404_ReturnsImageNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/images/img-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetImage("img-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &ImageNotFoundError{}, err)
	})
}

func TestUpdateImage(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		req := UpdateImageRequest{
			Name: "testimage",
		}

		c.EXPECT().Patch("/ecloud/v2/images/img-abcdef12", gomock.Eq(&req)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"task_id\":\"task-abcdef12\"},\"meta\":{\"location\":\"\"}}"))),
				StatusCode: 202,
			},
		}, nil)

		task, err := s.UpdateImage("img-abcdef12", req)

		assert.Nil(t, err)
		assert.Equal(t, "task-abcdef12", task.TaskID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ecloud/v2/images/img-abcdef12", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.UpdateImage("img-abcdef12", UpdateImageRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.UpdateImage("", UpdateImageRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid image id", err.Error())
	})
}

func TestDeleteImage(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v2/images/img-abcdef12", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"task_id\":\"task-abcdef12\"},\"meta\":{\"location\":\"\"}}"))),
				StatusCode: 202,
			},
		}, nil)

		taskID, err := s.DeleteImage("img-abcdef12")

		assert.Nil(t, err)
		assert.Equal(t, "task-abcdef12", taskID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v2/images/img-abcdef12", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.DeleteImage("img-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.DeleteImage("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid image id", err.Error())
	})
}

func TestGetImageParameters(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/images/img-abcdef12/parameters", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"key\":\"testkey\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		parameters, err := s.GetImageParameters("img-abcdef12", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, parameters, 1)
		assert.Equal(t, "testkey", parameters[0].Key)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ecloud/v2/images/img-abcdef12/parameters", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"key\":\"testkey1\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ecloud/v2/images/img-abcdef12/parameters", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"key\":\"testkey2\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		parameters, err := s.GetImageParameters("img-abcdef12", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, parameters, 2)
		assert.Equal(t, "testkey1", parameters[0].Key)
		assert.Equal(t, "testkey2", parameters[1].Key)
	})

	t.Run("connection.APIResponseBodyData[ImagesPaginated]Error_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/images/img-abcdef12/parameters", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetImageParameters("img-abcdef12", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}
