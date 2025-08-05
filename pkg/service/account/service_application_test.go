package account

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetApplications(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/account/v1/applications", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"test-id-123\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		applications, err := s.GetApplications(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, applications, 1)
		assert.Equal(t, "test-id-123", applications[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/account/v1/applications", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"test-id-123\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/account/v1/applications", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"test-id-456\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		applications, err := s.GetApplications(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, applications, 2)
		assert.Equal(t, "test-id-123", applications[0].ID)
		assert.Equal(t, "test-id-456", applications[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/account/v1/applications", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetApplications(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetApplication(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/account/v1/applications/test-id-123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"test-id-123\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		application, err := s.GetApplication("test-id-123")

		assert.Nil(t, err)
		assert.Equal(t, "test-id-123", application.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/account/v1/applications/test-id-123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetApplication("test-id-123")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidApplicationID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetApplication("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid application id", err.Error())
	})

	t.Run("404_ReturnsApplicationNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/account/v1/applications/test-id-123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetApplication("test-id-123")

		assert.NotNil(t, err)
		assert.IsType(t, &ApplicationNotFoundError{}, err)
	})
}

func TestGetServices(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/account/v1/services", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"test-id-123\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		services, err := s.GetServices(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, services, 1)
		assert.Equal(t, "test-id-123", services[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/account/v1/services", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"test-id-123\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/account/v1/services", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"test-id-456\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		services, err := s.GetServices(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, services, 2)
		assert.Equal(t, "test-id-123", services[0].ID)
		assert.Equal(t, "test-id-456", services[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/account/v1/services", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetServices(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestCreateApplication(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		createRequest := CreateApplicationRequest{}

		c.EXPECT().Post("/account/v1/applications", gomock.Eq(&createRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"test-id-123\"}}"))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		id, err := s.CreateApplication(createRequest)

		assert.Nil(t, err)
		assert.Equal(t, "test-id-123", id.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/account/v1/applications", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateApplication(CreateApplicationRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestUpdateApplication(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		updateRequest := UpdateApplicationRequest{}

		c.EXPECT().Patch("/account/v1/applications/test-id-123", gomock.Eq(&updateRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"test-id-123\"}}"))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		err := s.UpdateApplication("test-id-123", updateRequest)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/account/v1/applications/test-id-123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.UpdateApplication("test-id-123", UpdateApplicationRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidApplicationID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.UpdateApplication("", UpdateApplicationRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid application id", err.Error())
	})

	t.Run("404_ReturnsApplicationNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/account/v1/applications/test-id-123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.UpdateApplication("test-id-123", UpdateApplicationRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &ApplicationNotFoundError{}, err)
	})
}

func TestGetApplicationServices(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/account/v1/applications/test-id-123/services", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"scopes\":[{\"service\":\"accounts\",\"roles\":[\"read\",\"write\"]}]}}"))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		id, err := s.GetApplicationServices("test-id-123")

		assert.Nil(t, err)
		assert.Equal(t, "accounts", id.Scopes[0].Service)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/account/v1/applications/test-id-123/services", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetApplicationServices("test-id-123")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidApplicationID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetApplicationServices("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid application id", err.Error())
	})

	t.Run("404_ReturnsApplicationNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/account/v1/applications/test-id-123/services", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetApplicationServices("test-id-123")

		assert.NotNil(t, err)
		assert.IsType(t, &ApplicationNotFoundError{}, err)
	})
}

func TestSetApplicationServices(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		serviceRequest := SetServiceRequest{}

		c.EXPECT().Put("/account/v1/applications/test-id-123/services", gomock.Eq(&serviceRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"test-id-123\"}}"))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		err := s.SetApplicationServices("test-id-123", serviceRequest)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/account/v1/applications/test-id-123/services", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.SetApplicationServices("test-id-123", SetServiceRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidApplicationID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.SetApplicationServices("", SetServiceRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid application id", err.Error())
	})

	t.Run("404_ReturnsApplicationNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/account/v1/applications/test-id-456/services", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.SetApplicationServices("test-id-456", SetServiceRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &ApplicationNotFoundError{}, err)
	})
}

func TestDeleteApplicationServices(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/account/v1/applications/test-id-123/services", gomock.Eq(SetServiceRequest{Scopes: []ApplicationServiceScope{}})).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"test-id-123\"}}"))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		err := s.DeleteApplicationServices("test-id-123")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/account/v1/applications/test-id-123/services", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteApplicationServices("test-id-123")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidApplicationID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteApplicationServices("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid application id", err.Error())
	})

	t.Run("404_ReturnsApplicationNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/account/v1/applications/test-id-456/services", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteApplicationServices("test-id-456")

		assert.NotNil(t, err)
		assert.IsType(t, &ApplicationNotFoundError{}, err)
	})
}

func TestGetApplicationRestrictions(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/account/v1/applications/test-id-123/ip-restrictions", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"ip_restriction_type\":\"allowlist\",\"ip_ranges\":[\"8.8.8.8\"]},\"meta\":{}}"))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		id, err := s.GetApplicationRestrictions("test-id-123")

		assert.Nil(t, err)
		assert.Equal(t, "allowlist", id.IPRestrictionType)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/account/v1/applications/test-id-123/ip-restrictions", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetApplicationRestrictions("test-id-123")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidApplicationID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetApplicationRestrictions("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid application id", err.Error())
	})

	t.Run("404_ReturnsApplicationNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/account/v1/applications/test-id-123/ip-restrictions", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetApplicationRestrictions("test-id-123")

		assert.NotNil(t, err)
		assert.IsType(t, &ApplicationNotFoundError{}, err)
	})
}

func TestSetApplicationRestrictions(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		restrictionRequest := SetRestrictionRequest{}

		c.EXPECT().Put("/account/v1/applications/test-id-123/ip-restrictions", gomock.Eq(&restrictionRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		err := s.SetApplicationRestrictions("test-id-123", restrictionRequest)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/account/v1/applications/test-id-123/ip-restrictions", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.SetApplicationRestrictions("test-id-123", SetRestrictionRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidApplicationID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.SetApplicationRestrictions("", SetRestrictionRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid application id", err.Error())
	})

	t.Run("404_ReturnsApplicationNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/account/v1/applications/test-id-456/ip-restrictions", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.SetApplicationRestrictions("test-id-456", SetRestrictionRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &ApplicationNotFoundError{}, err)
	})
}

func TestDeleteApplicationRestriction(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/account/v1/applications/test-id-123/ip-restrictions", gomock.Eq(connection.APIRequestParameters{})).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.DeleteApplicationRestrictions("test-id-123")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/account/v1/applications/test-id-123/ip-restrictions", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteApplicationRestrictions("test-id-123")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidApplicationID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteApplicationRestrictions("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid application id", err.Error())
	})

	t.Run("404_ReturnsApplicationNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/account/v1/applications/test-id-456/ip-restrictions", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteApplicationRestrictions("test-id-456")

		assert.NotNil(t, err)
		assert.IsType(t, &ApplicationNotFoundError{}, err)
	})
}

func TestDeleteApplication(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/account/v1/applications/test-id-123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		err := s.DeleteApplication("test-id-123")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/account/v1/applications/test-id-123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteApplication("test-id-123")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidClientID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteApplication("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid application id", err.Error())
	})

	t.Run("404_ReturnsClientNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/account/v1/applications/test-id-123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteApplication("test-id-123")

		assert.NotNil(t, err)
		assert.IsType(t, &ApplicationNotFoundError{}, err)
	})
}
