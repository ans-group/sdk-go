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

func TestGetListenerCertificates(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/loadbalancers/v2/listeners/123/certs", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":456}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil)

		certificates, err := s.GetListenerCertificates(123, connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, certificates, 1)
		assert.Equal(t, 456, certificates[0].ID)
	})

	t.Run("InvalidListenerID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetListenerCertificates(0, connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid listener id", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/loadbalancers/v2/listeners/123/certs", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetListenerCertificates(123, connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetListenerCertificate(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/loadbalancers/v2/listeners/123/certs/456", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":456}}"))),
				StatusCode: 200,
			},
		}, nil)

		certificate, err := s.GetListenerCertificate(123, 456)

		assert.Nil(t, err)
		assert.Equal(t, 456, certificate.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/loadbalancers/v2/listeners/123/certs/456", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetListenerCertificate(123, 456)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidListenerID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetListenerCertificate(0, 456)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid listener id", err.Error())
	})

	t.Run("InvalidCertificateID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetListenerCertificate(123, 0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid certificate id", err.Error())
	})

	t.Run("404_ReturnsListenerCertificateNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/loadbalancers/v2/listeners/123/certs/456", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		_, err := s.GetListenerCertificate(123, 456)

		assert.NotNil(t, err)
		assert.IsType(t, &CertificateNotFoundError{}, err)
	})
}

func TestCreateListenerCertificate(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		req := CreateCertificateRequest{
			Key: "test",
		}

		c.EXPECT().Post("/loadbalancers/v2/listeners/123/certs", &req).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":456}}"))),
				StatusCode: 200,
			},
		}, nil)

		id, err := s.CreateListenerCertificate(123, req)

		assert.Nil(t, err)
		assert.Equal(t, 456, id)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/loadbalancers/v2/listeners/123/certs", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.CreateListenerCertificate(123, CreateCertificateRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidListenerID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.CreateListenerCertificate(0, CreateCertificateRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid listener id", err.Error())
	})

	t.Run("404_ReturnsListenerCertificateNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/loadbalancers/v2/listeners/123/certs", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		_, err := s.CreateListenerCertificate(123, CreateCertificateRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &CertificateNotFoundError{}, err)
	})
}

func TestPatchListenerCertificate(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		req := PatchCertificateRequest{
			Key: "test",
		}

		c.EXPECT().Patch("/loadbalancers/v2/listeners/123/certs/456", &req).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
				StatusCode: 200,
			},
		}, nil)

		err := s.PatchListenerCertificate(123, 456, req)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/loadbalancers/v2/listeners/123/certs/456", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		err := s.PatchListenerCertificate(123, 456, PatchCertificateRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidListenerID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchListenerCertificate(0, 456, PatchCertificateRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid listener id", err.Error())
	})

	t.Run("InvalidCertificateID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchListenerCertificate(123, 0, PatchCertificateRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid certificate id", err.Error())
	})

	t.Run("404_ReturnsListenerCertificateNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/loadbalancers/v2/listeners/123/certs/456", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		err := s.PatchListenerCertificate(123, 456, PatchCertificateRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &CertificateNotFoundError{}, err)
	})
}

func TestDeleteListenerCertificate(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/loadbalancers/v2/listeners/123/certs/456", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
				StatusCode: 200,
			},
		}, nil)

		err := s.DeleteListenerCertificate(123, 456)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/loadbalancers/v2/listeners/123/certs/456", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		err := s.DeleteListenerCertificate(123, 456)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidListenerID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteListenerCertificate(0, 456)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid listener id", err.Error())
	})

	t.Run("InvalidCertificateID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteListenerCertificate(123, 0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid certificate id", err.Error())
	})

	t.Run("404_ReturnsListenerCertificateNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/loadbalancers/v2/listeners/123/certs/456", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		err := s.DeleteListenerCertificate(123, 456)

		assert.NotNil(t, err)
		assert.IsType(t, &CertificateNotFoundError{}, err)
	})
}
