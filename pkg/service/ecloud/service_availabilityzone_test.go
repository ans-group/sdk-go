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

func TestGetAvailabilityZones(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/availability-zones", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"az-abcdef12\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		azs, err := s.GetAvailabilityZones(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, azs, 1)
		assert.Equal(t, "az-abcdef12", azs[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/availability-zones", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetAvailabilityZones(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetAvailabilityZone(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/availability-zones/az-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"az-abcdef12\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		az, err := s.GetAvailabilityZone("az-abcdef12")

		assert.Nil(t, err)
		assert.Equal(t, "az-abcdef12", az.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/availability-zones/az-abcdef12", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetAvailabilityZone("az-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidAvailabilityZoneID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetAvailabilityZone("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid az id", err.Error())
	})

	t.Run("404_ReturnsAvailabilityZoneNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/availability-zones/az-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetAvailabilityZone("az-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &AvailabilityZoneNotFoundError{}, err)
	})
}

func TestGetAvailabilityZoneIOPSTiers(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/availability-zones/az-abcdef12/iops", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"iops-abcdef12\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil)

		tiers, err := s.GetAvailabilityZoneIOPSTiers("az-abcdef12", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, tiers, 1)
		assert.Equal(t, "iops-abcdef12", tiers[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/availability-zones/az-abcdef12/iops", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetAvailabilityZoneIOPSTiers("az-abcdef12", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidAvailabilityZoneID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetAvailabilityZoneIOPSTiers("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid az id", err.Error())
	})

	t.Run("404_ReturnsAvailabilityZoneNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/availability-zones/az-abcdef12/iops", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		_, err := s.GetAvailabilityZoneIOPSTiers("az-abcdef12", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &AvailabilityZoneNotFoundError{}, err)
	})
}
