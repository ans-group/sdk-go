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

func TestGetLoadBalancerNetworks(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/load-balancer-networks", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"lbn-abcdef12\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		networks, err := s.GetLoadBalancerNetworks(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, networks, 1)
		assert.Equal(t, "lbn-abcdef12", networks[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/load-balancer-networks", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetLoadBalancerNetworks(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetLoadBalancerNetwork(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/load-balancer-networks/lbn-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"lbn-abcdef12\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		network, err := s.GetLoadBalancerNetwork("lbn-abcdef12")

		assert.Nil(t, err)
		assert.Equal(t, "lbn-abcdef12", network.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/load-balancer-networks/lbn-abcdef12", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetLoadBalancerNetwork("lbn-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidLoadBalancerNetworkID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetLoadBalancerNetwork("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid load balancer network id", err.Error())
	})

	t.Run("404_ReturnsLoadBalancerNetworkNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/load-balancer-networks/lbn-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetLoadBalancerNetwork("lbn-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &LoadBalancerNetworkNotFoundError{}, err)
	})
}

func TestCreateLoadBalancerNetwork(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		req := CreateLoadBalancerNetworkRequest{
			Name: "test",
		}

		c.EXPECT().Post("/ecloud/v2/load-balancer-networks", &req).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"lbn-abcdef12\",\"task_id\":\"task-abcdef12\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		taskRef, err := s.CreateLoadBalancerNetwork(req)

		assert.Nil(t, err)
		assert.Equal(t, "lbn-abcdef12", taskRef.ResourceID)
		assert.Equal(t, "task-abcdef12", taskRef.TaskID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v2/load-balancer-networks", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateLoadBalancerNetwork(CreateLoadBalancerNetworkRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestPatchLoadBalancerNetwork(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		req := PatchLoadBalancerNetworkRequest{
			Name: "somenetwork",
		}

		c.EXPECT().Patch("/ecloud/v2/load-balancer-networks/lbn-abcdef12", &req).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"lbn-abcdef12\",\"task_id\":\"task-abcdef12\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		task, err := s.PatchLoadBalancerNetwork("lbn-abcdef12", req)

		assert.Nil(t, err)
		assert.Equal(t, "lbn-abcdef12", task.ResourceID)
		assert.Equal(t, "task-abcdef12", task.TaskID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ecloud/v2/load-balancer-networks/lbn-abcdef12", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.PatchLoadBalancerNetwork("lbn-abcdef12", PatchLoadBalancerNetworkRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidLoadBalancerNetworkID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.PatchLoadBalancerNetwork("", PatchLoadBalancerNetworkRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid load balancer network id", err.Error())
	})

	t.Run("404_ReturnsLoadBalancerNetworkNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ecloud/v2/load-balancer-networks/lbn-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.PatchLoadBalancerNetwork("lbn-abcdef12", PatchLoadBalancerNetworkRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &LoadBalancerNetworkNotFoundError{}, err)
	})
}

func TestDeleteLoadBalancerNetwork(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v2/load-balancer-networks/lbn-abcdef12", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"task_id\":\"task-abcdef12\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		taskID, err := s.DeleteLoadBalancerNetwork("lbn-abcdef12")

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

		c.EXPECT().Delete("/ecloud/v2/load-balancer-networks/lbn-abcdef12", nil).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.DeleteLoadBalancerNetwork("lbn-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidLoadBalancerNetworkID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.DeleteLoadBalancerNetwork("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid load balancer network id", err.Error())
	})

	t.Run("404_ReturnsLoadBalancerNetworkNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v2/load-balancer-networks/lbn-abcdef12", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.DeleteLoadBalancerNetwork("lbn-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &LoadBalancerNetworkNotFoundError{}, err)
	})
}
