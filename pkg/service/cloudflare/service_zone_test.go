package cloudflare

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

func TestGetZones(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/cloudflare/v1/zones", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		zones, err := s.GetZones(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, zones, 1)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", zones[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/cloudflare/v1/zones", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetZones(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetZone(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/cloudflare/v1/zones/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		zone, err := s.GetZone("00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", zone.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/cloudflare/v1/zones/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetZone("00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidZoneID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetZone("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid zone id", err.Error())
	})

	t.Run("404_ReturnsZoneNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/cloudflare/v1/zones/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetZone("00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &ZoneNotFoundError{}, err)
	})
}

func TestCreateZone(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		createRequest := CreateZoneRequest{
			Name: "testzone1.co.uk",
		}

		c.EXPECT().Post("/cloudflare/v1/zones", gomock.Eq(&createRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		zoneID, err := s.CreateZone(createRequest)

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", zoneID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/cloudflare/v1/zones", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateZone(CreateZoneRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestPatchZone(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		patchRequest := PatchZoneRequest{
			SubscriptionID: "e69a8445-b224-4555-9785-8de9238f7eaf",
		}

		c.EXPECT().Post("/cloudflare/v1/zones/00000000-0000-0000-0000-000000000000", gomock.Eq(&patchRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PatchZone("00000000-0000-0000-0000-000000000000", patchRequest)

		assert.Nil(t, err)
	})

	t.Run("InvalidZoneID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchZone("", PatchZoneRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid zone id", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/cloudflare/v1/zones/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchZone("00000000-0000-0000-0000-000000000000", PatchZoneRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestDeleteZone(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/cloudflare/v1/zones/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.DeleteZone("00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/cloudflare/v1/zones/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteZone("00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidZoneID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteZone("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid zone id", err.Error())
	})

	t.Run("404_ReturnsZoneNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/cloudflare/v1/zones/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteZone("00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &ZoneNotFoundError{}, err)
	})
}
