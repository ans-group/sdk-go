package safedns

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

func TestGetTemplates(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/templates", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		templates, err := s.GetTemplates(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, templates, 1)
		assert.Equal(t, 123, templates[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		gomock.InOrder(
			c.EXPECT().Get("/safedns/v1/templates", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/safedns/v1/templates", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":456}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		templates, err := s.GetTemplates(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, templates, 2)
		assert.Equal(t, 123, templates[0].ID)
		assert.Equal(t, 456, templates[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/templates", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetTemplates(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetTemplate(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/templates/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":123}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		template, err := s.GetTemplate(123)

		assert.Nil(t, err)
		assert.Equal(t, 123, template.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/templates/123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetTemplate(123)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidTemplateID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		_, err := s.GetTemplate(0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid template id", err.Error())
	})

	t.Run("404_ReturnsTemplateNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/templates/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetTemplate(123)

		assert.NotNil(t, err)
		assert.IsType(t, &TemplateNotFoundError{}, err)
	})
}

func TestCreateTemplate(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		createRequest := CreateTemplateRequest{
			Name: "test template 1",
		}

		c.EXPECT().Post("/safedns/v1/templates", gomock.Eq(&createRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":123}}"))),
				StatusCode: 201,
			},
		}, nil).Times(1)

		id, err := s.CreateTemplate(createRequest)

		assert.Nil(t, err)
		assert.Equal(t, 123, id)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Post("/safedns/v1/templates", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateTemplate(CreateTemplateRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestPatchTemplate(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		patch := PatchTemplateRequest{
			Name: "new test template name 1",
		}

		c.EXPECT().Patch("/safedns/v1/templates/123", gomock.Eq(&patch)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":123}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		id, err := s.PatchTemplate(123, patch)

		assert.Nil(t, err)
		assert.Equal(t, 123, id)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Patch("/safedns/v1/templates/123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.PatchTemplate(123, PatchTemplateRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidTemplateID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		_, err := s.PatchTemplate(0, PatchTemplateRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid template id", err.Error())
	})

	t.Run("404_ReturnsTemplateNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Patch("/safedns/v1/templates/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.PatchTemplate(123, PatchTemplateRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &TemplateNotFoundError{}, err)
	})
}

func TestDeleteTemplate(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Delete("/safedns/v1/templates/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.DeleteTemplate(123)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Delete("/safedns/v1/templates/123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteTemplate(123)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidTemplateID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		err := s.DeleteTemplate(0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid template id", err.Error())
	})

	t.Run("404_ReturnsTemplateNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Delete("/safedns/v1/templates/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteTemplate(123)

		assert.NotNil(t, err)
		assert.IsType(t, &TemplateNotFoundError{}, err)
	})
}

func TestGetTemplateRecords(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/templates/123/records", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":456}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		records, err := s.GetTemplateRecords(123, connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, records, 1)
		assert.Equal(t, 456, records[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		gomock.InOrder(
			c.EXPECT().Get("/safedns/v1/templates/123/records", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":456}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/safedns/v1/templates/123/records", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":789}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		records, err := s.GetTemplateRecords(123, connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, records, 2)
		assert.Equal(t, 456, records[0].ID)
		assert.Equal(t, 789, records[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/templates/123/records", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetTemplateRecords(123, connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetTemplateRecord(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/templates/123/records/456", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":456}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		record, err := s.GetTemplateRecord(123, 456)

		assert.Nil(t, err)
		assert.Equal(t, 456, record.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/templates/123/records/456", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetTemplateRecord(123, 456)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidTemplateID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		_, err := s.GetTemplateRecord(0, 456)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid template id", err.Error())
	})

	t.Run("InvalidRecordID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		_, err := s.GetTemplateRecord(123, 0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid record id", err.Error())
	})

	t.Run("404_ReturnsTemplateRecordNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/templates/123/records/456", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetTemplateRecord(123, 456)

		assert.NotNil(t, err)
		assert.IsType(t, &TemplateRecordNotFoundError{}, err)
	})
}

func TestCreateTemplateRecord(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		createReq := CreateRecordRequest{
			Name: "test.123",
		}

		c.EXPECT().Post("/safedns/v1/templates/123/records", gomock.Eq(&createReq)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":456}}"))),
				StatusCode: 201,
			},
		}, nil).Times(1)

		id, err := s.CreateTemplateRecord(123, createReq)

		assert.Nil(t, err)
		assert.Equal(t, 456, id)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Post("/safedns/v1/templates/123/records", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateTemplateRecord(123, CreateRecordRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidTemplateID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		_, err := s.CreateTemplateRecord(0, CreateRecordRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid template id", err.Error())
	})

	t.Run("404_ReturnsTemplateNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Post("/safedns/v1/templates/123/records", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.CreateTemplateRecord(123, CreateRecordRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &TemplateNotFoundError{}, err)
	})
}

func TestPatchTemplateRecord(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		patch := PatchRecordRequest{
			Content: "1.2.3.4",
		}

		c.EXPECT().Patch("/safedns/v1/templates/123/records/456", gomock.Eq(&patch)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		_, err := s.PatchTemplateRecord(123, 456, patch)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Patch("/safedns/v1/templates/123/records/456", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.PatchTemplateRecord(123, 456, PatchRecordRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidTemplateID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		_, err := s.PatchTemplateRecord(0, 456, PatchRecordRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid template id", err.Error())
	})

	t.Run("InvalidRecordID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		_, err := s.PatchTemplateRecord(123, 0, PatchRecordRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid record id", err.Error())
	})

	t.Run("404_ReturnsTemplateRecordNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Patch("/safedns/v1/templates/123/records/456", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.PatchTemplateRecord(123, 456, PatchRecordRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &TemplateRecordNotFoundError{}, err)
	})
}

func TestDeleteTemplateRecord(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Delete("/safedns/v1/templates/123/records/456", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.DeleteTemplateRecord(123, 456)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Delete("/safedns/v1/templates/123/records/456", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteTemplateRecord(123, 456)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidTemplateID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		err := s.DeleteTemplateRecord(0, 456)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid template id", err.Error())
	})

	t.Run("InvalidRecordID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		err := s.DeleteTemplateRecord(123, 0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid record id", err.Error())
	})

	t.Run("404_ReturnsTemplateRecordNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Delete("/safedns/v1/templates/123/records/456", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteTemplateRecord(123, 456)

		assert.NotNil(t, err)
		assert.IsType(t, &TemplateRecordNotFoundError{}, err)
	})
}
