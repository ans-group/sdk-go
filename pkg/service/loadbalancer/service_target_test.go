package loadbalancer

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

func TestGetTargets(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/loadbalancers/v2/targets", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		targets, err := s.GetTargets(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, targets, 1)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", targets[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/loadbalancers/v2/targets", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetTargets(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetTarget(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/loadbalancers/v2/targets/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		target, err := s.GetTarget("00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", target.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/loadbalancers/v2/targets/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetTarget("00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidTargetID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetTarget("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid target id", err.Error())
	})

	t.Run("404_ReturnsTargetNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/loadbalancers/v2/targets/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetTarget("00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &TargetNotFoundError{}, err)
	})
}

func TestPatchTarget(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		req := PatchTargetRequest{
			Port: 8080,
		}

		c.EXPECT().Patch("/loadbalancers/v2/targets/00000000-0000-0000-0000-000000000000", &req).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PatchTarget("00000000-0000-0000-0000-000000000000", req)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/loadbalancers/v2/targets/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchTarget("00000000-0000-0000-0000-000000000000", PatchTargetRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidTargetID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchTarget("", PatchTargetRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid target id", err.Error())
	})

	t.Run("404_ReturnsTargetNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/loadbalancers/v2/targets/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PatchTarget("00000000-0000-0000-0000-000000000000", PatchTargetRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &TargetNotFoundError{}, err)
	})
}

func TestDeleteTarget(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/loadbalancers/v2/targets/00000000-0000-0000-0000-000000000000", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.DeleteTarget("00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/loadbalancers/v2/targets/00000000-0000-0000-0000-000000000000", nil).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteTarget("00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidTargetID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteTarget("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid target id", err.Error())
	})

	t.Run("404_ReturnsTargetNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/loadbalancers/v2/targets/00000000-0000-0000-0000-000000000000", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteTarget("00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &TargetNotFoundError{}, err)
	})
}
