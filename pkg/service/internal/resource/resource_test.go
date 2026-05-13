package resource_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
	"github.com/ans-group/sdk-go/test/mocks"
)

type testModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type testStringModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type testNotFoundError struct {
	ID int
}

func (e *testNotFoundError) Error() string {
	return fmt.Sprintf("test model not found with id [%d]", e.ID)
}

type testStringNotFoundError struct {
	ID string
}

func (e *testStringNotFoundError) Error() string {
	return fmt.Sprintf("test model not found with id [%s]", e.ID)
}

func newIntResource(c connection.Connection) *resource.Resource[testModel, int] {
	return resource.NewIntResource[testModel](c, "/test/v1/models", "model",
		func(id int) error { return &testNotFoundError{ID: id} })
}

func newStringResource(c connection.Connection) *resource.Resource[testStringModel, string] {
	return resource.NewStringResource[testStringModel](c, "/test/v1/models", "model",
		func(id string) error { return &testStringNotFoundError{ID: id} })
}

func apiResponseJSON(statusCode int, body string) *connection.APIResponse {
	return &connection.APIResponse{
		Response: &http.Response{
			Body:       io.NopCloser(bytes.NewReader([]byte(body))),
			StatusCode: statusCode,
		},
	}
}

// --- List ---

func TestResource_List(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Get("/test/v1/models", gomock.Any()).Return(
			apiResponseJSON(200, `{"data":[{"id":1,"name":"foo"}],"meta":{"pagination":{"total_pages":1}}}`), nil).Times(1)

		items, err := newIntResource(c).List(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, items, 1)
		assert.Equal(t, 1, items[0].ID)
	})

	t.Run("MultiPage", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		gomock.InOrder(
			c.EXPECT().Get("/test/v1/models", gomock.Any()).Return(
				apiResponseJSON(200, `{"data":[{"id":1}],"meta":{"pagination":{"total_pages":2}}}`), nil),
			c.EXPECT().Get("/test/v1/models", gomock.Any()).Return(
				apiResponseJSON(200, `{"data":[{"id":2}],"meta":{"pagination":{"total_pages":2}}}`), nil),
		)

		items, err := newIntResource(c).List(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, items, 2)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Get("/test/v1/models", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error"))

		_, err := newIntResource(c).List(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error", err.Error())
	})
}

// --- ListPaginated ---

func TestResource_ListPaginated(t *testing.T) {
	t.Run("ReturnsPaginated", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Get("/test/v1/models", gomock.Any()).Return(
			apiResponseJSON(200, `{"data":[{"id":1}],"meta":{"pagination":{"total_pages":1}}}`), nil).Times(1)

		page, err := newIntResource(c).ListPaginated(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, page.Items(), 1)
		assert.Equal(t, 1, page.Items()[0].ID)
	})
}

// --- Get (int ID) ---

func TestResource_Get_IntID(t *testing.T) {
	t.Run("Valid_ReturnsItem", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Get("/test/v1/models/123", gomock.Any()).Return(
			apiResponseJSON(200, `{"data":{"id":123,"name":"bar"}}`), nil).Times(1)

		item, err := newIntResource(c).Get(123)

		assert.Nil(t, err)
		assert.Equal(t, 123, item.ID)
	})

	t.Run("ZeroID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)

		_, err := newIntResource(c).Get(0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid model id", err.Error())
	})

	t.Run("NegativeID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)

		_, err := newIntResource(c).Get(-1)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid model id", err.Error())
	})

	t.Run("NotFound_ReturnsNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Get("/test/v1/models/123", gomock.Any()).Return(
			apiResponseJSON(404, ""), nil).Times(1)

		_, err := newIntResource(c).Get(123)

		assert.NotNil(t, err)
		assert.IsType(t, &testNotFoundError{}, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Get("/test/v1/models/123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error"))

		_, err := newIntResource(c).Get(123)

		assert.NotNil(t, err)
		assert.Equal(t, "test error", err.Error())
	})
}

// --- Get (string ID) ---

func TestResource_Get_StringID(t *testing.T) {
	t.Run("Valid_ReturnsItem", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Get("/test/v1/models/abc-123", gomock.Any()).Return(
			apiResponseJSON(200, `{"data":{"id":"abc-123","name":"baz"}}`), nil).Times(1)

		item, err := newStringResource(c).Get("abc-123")

		assert.Nil(t, err)
		assert.Equal(t, "abc-123", item.ID)
	})

	t.Run("EmptyID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)

		_, err := newStringResource(c).Get("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid model id", err.Error())
	})

	t.Run("NotFound_ReturnsNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Get("/test/v1/models/abc-123", gomock.Any()).Return(
			apiResponseJSON(404, ""), nil).Times(1)

		_, err := newStringResource(c).Get("abc-123")

		assert.NotNil(t, err)
		assert.IsType(t, &testStringNotFoundError{}, err)
	})
}

// --- Create ---

func TestResource_Create(t *testing.T) {
	t.Run("ReturnsCreatedModel", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Post("/test/v1/models", gomock.Any()).Return(
			apiResponseJSON(200, `{"data":{"id":99,"name":"new"}}`), nil).Times(1)

		item, err := newIntResource(c).Create(struct{ Name string }{Name: "new"})

		assert.Nil(t, err)
		assert.Equal(t, 99, item.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Post("/test/v1/models", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error"))

		_, err := newIntResource(c).Create(struct{}{})

		assert.NotNil(t, err)
	})
}

// --- Patch ---

func TestResource_Patch(t *testing.T) {
	t.Run("Valid_ReturnsNil", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Patch("/test/v1/models/123", gomock.Any()).Return(
			apiResponseJSON(204, `{}`), nil).Times(1)

		err := newIntResource(c).Patch(123, struct{ Name string }{Name: "updated"})

		assert.Nil(t, err)
	})

	t.Run("InvalidID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)

		err := newIntResource(c).Patch(0, struct{}{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid model id", err.Error())
	})

	t.Run("NotFound_ReturnsNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Patch("/test/v1/models/123", gomock.Any()).Return(
			apiResponseJSON(404, ""), nil).Times(1)

		err := newIntResource(c).Patch(123, struct{}{})

		assert.NotNil(t, err)
		assert.IsType(t, &testNotFoundError{}, err)
	})
}

// --- Delete ---

func TestResource_Delete(t *testing.T) {
	t.Run("Valid_ReturnsNil", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Delete("/test/v1/models/123", gomock.Any()).Return(
			apiResponseJSON(204, `{}`), nil).Times(1)

		err := newIntResource(c).Delete(123)

		assert.Nil(t, err)
	})

	t.Run("InvalidID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)

		err := newIntResource(c).Delete(0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid model id", err.Error())
	})

	t.Run("EmptyStringID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)

		err := newStringResource(c).Delete("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid model id", err.Error())
	})

	t.Run("NotFound_ReturnsNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Delete("/test/v1/models/123", gomock.Any()).Return(
			apiResponseJSON(404, ""), nil).Times(1)

		err := newIntResource(c).Delete(123)

		assert.NotNil(t, err)
		assert.IsType(t, &testNotFoundError{}, err)
	})
}
