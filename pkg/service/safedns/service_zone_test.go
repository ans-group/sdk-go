package safedns

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/test/mocks"
)

func TestGetZones(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/zones", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"name\":\"testdomain1.co.uk\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		zones, err := s.GetZones(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, zones, 1)
		assert.Equal(t, "testdomain1.co.uk", zones[0].Name)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		gomock.InOrder(
			c.EXPECT().Get("/safedns/v1/zones", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"name\":\"testdomain1.co.uk\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/safedns/v1/zones", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"name\":\"testdomain2.co.uk\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		zones, err := s.GetZones(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, zones, 2)
		assert.Equal(t, "testdomain1.co.uk", zones[0].Name)
		assert.Equal(t, "testdomain2.co.uk", zones[1].Name)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/zones", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

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

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/zones/testdomain1.co.uk", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"name\":\"testdomain1.co.uk\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		zone, err := s.GetZone("testdomain1.co.uk")

		assert.Nil(t, err)
		assert.Equal(t, "testdomain1.co.uk", zone.Name)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/zones/testdomain1.co.uk", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetZone("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidZoneName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		_, err := s.GetZone("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid zone name", err.Error())
	})

	t.Run("404_ReturnsZoneNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/zones/testdomain1.co.uk", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetZone("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.IsType(t, &ZoneNotFoundError{}, err)
	})
}

func TestCreateZone(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		createRequest := CreateZoneRequest{
			Name: "testdomain1.co.uk",
		}

		c.EXPECT().Post("/safedns/v1/zones", gomock.Eq(&createRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":123}}"))),
				StatusCode: 201,
			},
		}, nil).Times(1)

		err := s.CreateZone(createRequest)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Post("/safedns/v1/zones", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.CreateZone(CreateZoneRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestPatchZone(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		patchRequest := PatchZoneRequest{
			Description: "testdomain1.co.uk",
		}

		c.EXPECT().Patch("/safedns/v1/zones/testdomain1.co.uk", gomock.Eq(&patchRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PatchZone("testdomain1.co.uk", patchRequest)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Patch("/safedns/v1/zones/testdomain1.co.uk", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchZone("testdomain1.co.uk", PatchZoneRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidZoneName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		err := s.PatchZone("", PatchZoneRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid zone name", err.Error())
	})
}

func TestDeleteZone(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Delete("/safedns/v1/zones/testdomain1.co.uk", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.DeleteZone("testdomain1.co.uk")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Delete("/safedns/v1/zones/testdomain1.co.uk", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteZone("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidZoneName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		err := s.DeleteZone("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid zone name", err.Error())
	})

	t.Run("404_ReturnsZoneNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Delete("/safedns/v1/zones/testdomain1.co.uk", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteZone("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.IsType(t, &ZoneNotFoundError{}, err)
	})
}

func TestGetZoneRecords(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/zones/testdomain1.co.uk/records", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		records, err := s.GetZoneRecords("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, records, 1)
		assert.Equal(t, 123, records[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		gomock.InOrder(
			c.EXPECT().Get("/safedns/v1/zones/testdomain1.co.uk/records", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/safedns/v1/zones/testdomain1.co.uk/records", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":456}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		records, err := s.GetZoneRecords("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, records, 2)
		assert.Equal(t, 123, records[0].ID)
		assert.Equal(t, 456, records[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/zones/testdomain1.co.uk/records", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetZoneRecords("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetZoneRecord(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/zones/testdomain1.co.uk/records/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":123}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		record, err := s.GetZoneRecord("testdomain1.co.uk", 123)

		assert.Nil(t, err)
		assert.Equal(t, 123, record.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/zones/testdomain1.co.uk/records/123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetZoneRecord("testdomain1.co.uk", 123)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidZoneName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		_, err := s.GetZoneRecord("", 123)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid zone name", err.Error())
	})

	t.Run("InvalidRecordID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		_, err := s.GetZoneRecord("testdomain1.co.uk", 0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid record id", err.Error())
	})

	t.Run("404_ReturnsZoneRecordNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/zones/testdomain1.co.uk/records/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetZoneRecord("testdomain1.co.uk", 123)

		assert.NotNil(t, err)
		assert.IsType(t, &ZoneRecordNotFoundError{}, err)
	})
}

func TestCreateZoneRecord(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		createReq := CreateRecordRequest{
			Name: "test.testdomain1.co.uk",
		}

		c.EXPECT().Post("/safedns/v1/zones/testdomain1.co.uk/records", gomock.Eq(&createReq)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":123}}"))),
				StatusCode: 201,
			},
		}, nil).Times(1)

		id, err := s.CreateZoneRecord("testdomain1.co.uk", createReq)

		assert.Nil(t, err)
		assert.Equal(t, 123, id)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Post("/safedns/v1/zones/testdomain1.co.uk/records", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateZoneRecord("testdomain1.co.uk", CreateRecordRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidZoneName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		_, err := s.CreateZoneRecord("", CreateRecordRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid zone name", err.Error())
	})

	t.Run("404_ReturnsZoneNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Post("/safedns/v1/zones/testdomain1.co.uk/records", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.CreateZoneRecord("testdomain1.co.uk", CreateRecordRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &ZoneNotFoundError{}, err)
	})
}

func TestUpdateZoneRecord(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		record := Record{
			ID:      123,
			Content: "1.2.3.4",
		}

		c.EXPECT().Put("/safedns/v1/zones/testdomain1.co.uk/records/123", gomock.Eq(&record)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		_, err := s.UpdateZoneRecord("testdomain1.co.uk", record)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Put("/safedns/v1/zones/testdomain1.co.uk/records/123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.UpdateZoneRecord("testdomain1.co.uk", Record{ID: 123})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidZoneName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		_, err := s.UpdateZoneRecord("", Record{ID: 123})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid zone name", err.Error())
	})

	t.Run("InvalidRecordID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		_, err := s.UpdateZoneRecord("testdomain1.co.uk", Record{ID: 0})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid record id", err.Error())
	})

	t.Run("404_ReturnsZoneRecordNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Put("/safedns/v1/zones/testdomain1.co.uk/records/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.UpdateZoneRecord("testdomain1.co.uk", Record{ID: 123})

		assert.NotNil(t, err)
		assert.IsType(t, &ZoneRecordNotFoundError{}, err)
	})
}

func TestPatchZoneRecord(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		patch := PatchRecordRequest{
			Content: "1.2.3.4",
		}

		c.EXPECT().Put("/safedns/v1/zones/testdomain1.co.uk/records/123", gomock.Eq(&patch)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		_, err := s.PatchZoneRecord("testdomain1.co.uk", 123, patch)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Put("/safedns/v1/zones/testdomain1.co.uk/records/123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.PatchZoneRecord("testdomain1.co.uk", 123, PatchRecordRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidZoneName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		_, err := s.PatchZoneRecord("", 123, PatchRecordRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid zone name", err.Error())
	})

	t.Run("InvalidRecordID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		_, err := s.PatchZoneRecord("testdomain1.co.uk", 0, PatchRecordRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid record id", err.Error())
	})

	t.Run("404_ReturnsZoneRecordNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Put("/safedns/v1/zones/testdomain1.co.uk/records/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.PatchZoneRecord("testdomain1.co.uk", 123, PatchRecordRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &ZoneRecordNotFoundError{}, err)
	})
}

func TestDeleteZoneRecord(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Delete("/safedns/v1/zones/testdomain1.co.uk/records/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.DeleteZoneRecord("testdomain1.co.uk", 123)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Delete("/safedns/v1/zones/testdomain1.co.uk/records/123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteZoneRecord("testdomain1.co.uk", 123)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidZoneName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		err := s.DeleteZoneRecord("", 123)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid zone name", err.Error())
	})

	t.Run("InvalidRecordID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		err := s.DeleteZoneRecord("testdomain1.co.uk", 0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid record id", err.Error())
	})

	t.Run("404_ReturnsZoneRecordNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Delete("/safedns/v1/zones/testdomain1.co.uk/records/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteZoneRecord("testdomain1.co.uk", 123)

		assert.NotNil(t, err)
		assert.IsType(t, &ZoneRecordNotFoundError{}, err)
	})
}

func TestGetZoneNotes(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/zones/testdomain1.co.uk/notes", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		notes, err := s.GetZoneNotes("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, notes, 1)
		assert.Equal(t, 123, notes[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		gomock.InOrder(
			c.EXPECT().Get("/safedns/v1/zones/testdomain1.co.uk/notes", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/safedns/v1/zones/testdomain1.co.uk/notes", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":456}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		notes, err := s.GetZoneNotes("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, notes, 2)
		assert.Equal(t, 123, notes[0].ID)
		assert.Equal(t, 456, notes[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/zones/testdomain1.co.uk/notes", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetZoneNotes("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetZoneNote(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/zones/testdomain1.co.uk/notes/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":123}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		note, err := s.GetZoneNote("testdomain1.co.uk", 123)

		assert.Nil(t, err)
		assert.Equal(t, 123, note.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/zones/testdomain1.co.uk/notes/123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetZoneNote("testdomain1.co.uk", 123)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidZoneName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		_, err := s.GetZoneNote("", 123)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid zone name", err.Error())
	})

	t.Run("InvalidNoteID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		_, err := s.GetZoneNote("testdomain1.co.uk", 0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid note id", err.Error())
	})

	t.Run("404_ReturnsZoneNoteNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/zones/testdomain1.co.uk/notes/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetZoneNote("testdomain1.co.uk", 123)

		assert.NotNil(t, err)
		assert.IsType(t, &ZoneNoteNotFoundError{}, err)
	})
}

func TestCreateZoneNote(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		createReq := CreateNoteRequest{
			Notes: "test note 1",
		}

		c.EXPECT().Post("/safedns/v1/zones/testdomain1.co.uk/notes", gomock.Eq(&createReq)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":123}}"))),
				StatusCode: 201,
			},
		}, nil).Times(1)

		id, err := s.CreateZoneNote("testdomain1.co.uk", createReq)

		assert.Nil(t, err)
		assert.Equal(t, 123, id)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Post("/safedns/v1/zones/testdomain1.co.uk/notes", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateZoneNote("testdomain1.co.uk", CreateNoteRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidZoneName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		_, err := s.CreateZoneNote("", CreateNoteRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid zone name", err.Error())
	})

	t.Run("404_ReturnsZoneNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Post("/safedns/v1/zones/testdomain1.co.uk/notes", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.CreateZoneNote("testdomain1.co.uk", CreateNoteRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &ZoneNotFoundError{}, err)
	})
}
