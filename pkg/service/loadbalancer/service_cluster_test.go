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
	"github.com/ukfast/sdk-go/pkg/ptr"
	"github.com/ukfast/sdk-go/test/mocks"
)

func TestGetClusters(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/loadbalancers/v2/clusters", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		clusters, err := s.GetClusters(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, clusters, 1)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", clusters[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/loadbalancers/v2/clusters", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetClusters(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetCluster(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/loadbalancers/v2/clusters/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		cluster, err := s.GetCluster("00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", cluster.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/loadbalancers/v2/clusters/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetCluster("00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidClusterID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetCluster("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid cluster id", err.Error())
	})

	t.Run("404_ReturnsClusterNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/loadbalancers/v2/clusters/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetCluster("00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &ClusterNotFoundError{}, err)
	})
}

func TestPatchCluster(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		req := PatchClusterRequest{
			Name: ptr.String("somecluster"),
		}

		c.EXPECT().Patch("/loadbalancers/v2/clusters/00000000-0000-0000-0000-000000000000", &req).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PatchCluster("00000000-0000-0000-0000-000000000000", req)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/loadbalancers/v2/clusters/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchCluster("00000000-0000-0000-0000-000000000000", PatchClusterRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidClusterID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchCluster("", PatchClusterRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid cluster id", err.Error())
	})

	t.Run("404_ReturnsClusterNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/loadbalancers/v2/clusters/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PatchCluster("00000000-0000-0000-0000-000000000000", PatchClusterRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &ClusterNotFoundError{}, err)
	})
}

func TestDeleteCluster(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/loadbalancers/v2/clusters/00000000-0000-0000-0000-000000000000", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.DeleteCluster("00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/loadbalancers/v2/clusters/00000000-0000-0000-0000-000000000000", nil).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteCluster("00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidClusterID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteCluster("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid cluster id", err.Error())
	})

	t.Run("404_ReturnsClusterNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/loadbalancers/v2/clusters/00000000-0000-0000-0000-000000000000", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteCluster("00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &ClusterNotFoundError{}, err)
	})
}
