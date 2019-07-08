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

func TestGetFirewalls(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/firewalls", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		firewalls, err := s.GetFirewalls(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, firewalls, 1)
		assert.Equal(t, 123, firewalls[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ecloud/v1/firewalls", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ecloud/v1/firewalls", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":456}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		firewalls, err := s.GetFirewalls(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, firewalls, 2)
		assert.Equal(t, 123, firewalls[0].ID)
		assert.Equal(t, 456, firewalls[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/firewalls", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetFirewalls(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetFirewall(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/firewalls/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":123}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		firewall, err := s.GetFirewall(123)

		assert.Nil(t, err)
		assert.Equal(t, 123, firewall.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/firewalls/123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetFirewall(123)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidFirewallID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetFirewall(0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid firewall id", err.Error())
	})

	t.Run("404_ReturnsFirewallNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/firewalls/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetFirewall(123)

		assert.NotNil(t, err)
		assert.IsType(t, &FirewallNotFoundError{}, err)
	})
}

func TestGetFirewallConfig(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/firewalls/123/config", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"config\":\"testconfig1\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		config, err := s.GetFirewallConfig(123)

		assert.Nil(t, err)
		assert.Equal(t, "testconfig1", config.Config)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/firewalls/123/config", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetFirewallConfig(123)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidFirewallID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetFirewallConfig(0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid firewall id", err.Error())
	})

	t.Run("404_ReturnsFirewallNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/firewalls/123/config", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetFirewallConfig(123)

		assert.NotNil(t, err)
		assert.IsType(t, &FirewallNotFoundError{}, err)
	})
}
