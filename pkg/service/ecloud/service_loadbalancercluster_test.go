package ecloud

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

func TestGetLoadBalancerClusters(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/lbcs", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"lbc-abcdef12\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		lbc, err := s.GetLoadBalancerClusters(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, lbc, 1)
		assert.Equal(t, "lbc-abcdef12", lbc[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/lbcs", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetLoadBalancerClusters(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetLoadBalancerCluster(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/lbcs/lbc-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"lbc-abcdef12\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		lbc, err := s.GetLoadBalancerCluster("lbc-abcdef12")

		assert.Nil(t, err)
		assert.Equal(t, "lbc-abcdef12", lbc.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/lbcs/lbc-abcdef12", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetLoadBalancerCluster("lbc-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidLoadBalancerClusterID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetLoadBalancerCluster("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid load balancer cluster id", err.Error())
	})

	t.Run("404_ReturnsLoadBalancerClusterNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/lbcs/lbc-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetLoadBalancerCluster("lbc-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &LoadBalancerClusterNotFoundError{}, err)
	})
}
