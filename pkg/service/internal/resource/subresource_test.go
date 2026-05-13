package resource_test

import (
	"errors"
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
	"github.com/ans-group/sdk-go/test/mocks"
)

// childModel and error types used in sub-resource tests

type childModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type parentNotFoundError struct {
	ParentID string
}

func (e *parentNotFoundError) Error() string {
	return fmt.Sprintf("parent not found [%s]", e.ParentID)
}

type childNotFoundError struct {
	ChildID int
}

func (e *childNotFoundError) Error() string {
	return fmt.Sprintf("child not found [%d]", e.ChildID)
}

// helpers to build reusable SubResourceList and SubResource instances

func newStringSubResourceList(c connection.Connection) *resource.SubResourceList[childModel, string] {
	return resource.NewStringSubResourceList[childModel](c,
		func(parentID string) string { return fmt.Sprintf("/test/v1/parents/%s/children", parentID) },
		"parent", "id",
		func(parentID string) error { return &parentNotFoundError{ParentID: parentID} })
}

func newUncheckedSubResourceList(c connection.Connection) *resource.SubResourceList[childModel, string] {
	return resource.NewUncheckedStringSubResourceList[childModel](c,
		func(parentID string) string { return fmt.Sprintf("/test/v1/parents/%s/children", parentID) })
}

func newStringIntSubResource(c connection.Connection) *resource.SubResource[childModel, string, int] {
	return resource.NewStringIntSubResource[childModel](c,
		func(parentID string) string { return fmt.Sprintf("/test/v1/parents/%s/children", parentID) },
		"parent", "id",
		func(parentID string) error { return &parentNotFoundError{ParentID: parentID} },
		"child", "id",
		func(_ string, childID int) error { return &childNotFoundError{ChildID: childID} })
}

// --- SubResourceList ---

func TestSubResourceList_List(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Get("/test/v1/parents/p1/children", gomock.Any()).Return(
			apiResponseJSON(200, `{"data":[{"id":1,"name":"foo"}],"meta":{"pagination":{"total_pages":1}}}`), nil).Times(1)

		items, err := newStringSubResourceList(c).List("p1", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, items, 1)
		assert.Equal(t, 1, items[0].ID)
	})

	t.Run("MultiPage", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		gomock.InOrder(
			c.EXPECT().Get("/test/v1/parents/p1/children", gomock.Any()).Return(
				apiResponseJSON(200, `{"data":[{"id":1}],"meta":{"pagination":{"total_pages":2}}}`), nil),
			c.EXPECT().Get("/test/v1/parents/p1/children", gomock.Any()).Return(
				apiResponseJSON(200, `{"data":[{"id":2}],"meta":{"pagination":{"total_pages":2}}}`), nil),
		)

		items, err := newStringSubResourceList(c).List("p1", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, items, 2)
	})
}

func TestSubResourceList_ListPaginated(t *testing.T) {
	t.Run("Valid_ReturnsPaginated", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Get("/test/v1/parents/p1/children", gomock.Any()).Return(
			apiResponseJSON(200, `{"data":[{"id":1}],"meta":{"pagination":{"total_pages":1}}}`), nil).Times(1)

		page, err := newStringSubResourceList(c).ListPaginated("p1", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, page.Items(), 1)
	})

	t.Run("EmptyParentID_ReturnsValidationError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)

		_, err := newStringSubResourceList(c).ListPaginated("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid parent id", err.Error())
	})

	t.Run("NotFound_ReturnsParentNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Get("/test/v1/parents/p1/children", gomock.Any()).Return(
			apiResponseJSON(404, ""), nil).Times(1)

		_, err := newStringSubResourceList(c).ListPaginated("p1", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &parentNotFoundError{}, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Get("/test/v1/parents/p1/children", gomock.Any()).Return(
			&connection.APIResponse{}, errors.New("test error"))

		_, err := newStringSubResourceList(c).ListPaginated("p1", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error", err.Error())
	})
}

func TestSubResourceList_Unchecked(t *testing.T) {
	t.Run("EmptyParentID_DoesNotValidate", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Get("/test/v1/parents//children", gomock.Any()).Return(
			apiResponseJSON(200, `{"data":[],"meta":{"pagination":{"total_pages":1}}}`), nil).Times(1)

		_, err := newUncheckedSubResourceList(c).ListPaginated("", connection.APIRequestParameters{})

		assert.Nil(t, err)
	})
}

// --- SubResource ---

func TestSubResource_ListPaginated(t *testing.T) {
	t.Run("Valid_ReturnsPaginated", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Get("/test/v1/parents/p1/children", gomock.Any()).Return(
			apiResponseJSON(200, `{"data":[{"id":1}],"meta":{"pagination":{"total_pages":1}}}`), nil).Times(1)

		page, err := newStringIntSubResource(c).ListPaginated("p1", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, page.Items(), 1)
	})

	t.Run("EmptyParentID_ReturnsValidationError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)

		_, err := newStringIntSubResource(c).ListPaginated("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid parent id", err.Error())
	})

	t.Run("NotFound_ReturnsParentNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Get("/test/v1/parents/p1/children", gomock.Any()).Return(
			apiResponseJSON(404, ""), nil).Times(1)

		_, err := newStringIntSubResource(c).ListPaginated("p1", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &parentNotFoundError{}, err)
	})
}

func TestSubResource_Get(t *testing.T) {
	t.Run("Valid_ReturnsItem", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Get("/test/v1/parents/p1/children/42", gomock.Any()).Return(
			apiResponseJSON(200, `{"data":{"id":42,"name":"child"}}`), nil).Times(1)

		item, err := newStringIntSubResource(c).Get("p1", 42)

		assert.Nil(t, err)
		assert.Equal(t, 42, item.ID)
	})

	t.Run("EmptyParentID_ReturnsValidationError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)

		_, err := newStringIntSubResource(c).Get("", 42)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid parent id", err.Error())
	})

	t.Run("ZeroChildID_ReturnsValidationError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)

		_, err := newStringIntSubResource(c).Get("p1", 0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid child id", err.Error())
	})

	t.Run("NotFound_ReturnsChildNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Get("/test/v1/parents/p1/children/42", gomock.Any()).Return(
			apiResponseJSON(404, ""), nil).Times(1)

		_, err := newStringIntSubResource(c).Get("p1", 42)

		assert.NotNil(t, err)
		assert.IsType(t, &childNotFoundError{}, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Get("/test/v1/parents/p1/children/42", gomock.Any()).Return(
			&connection.APIResponse{}, errors.New("test error"))

		_, err := newStringIntSubResource(c).Get("p1", 42)

		assert.NotNil(t, err)
		assert.Equal(t, "test error", err.Error())
	})
}

func TestSubResource_Create(t *testing.T) {
	t.Run("Valid_ReturnsCreated", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Post("/test/v1/parents/p1/children", gomock.Any()).Return(
			apiResponseJSON(201, `{"data":{"id":99,"name":"new"}}`), nil).Times(1)

		item, err := newStringIntSubResource(c).Create("p1", struct{}{})

		assert.Nil(t, err)
		assert.Equal(t, 99, item.ID)
	})

	t.Run("EmptyParentID_ReturnsValidationError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)

		_, err := newStringIntSubResource(c).Create("", struct{}{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid parent id", err.Error())
	})

	t.Run("NotFound_ReturnsParentNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Post("/test/v1/parents/p1/children", gomock.Any()).Return(
			apiResponseJSON(404, ""), nil).Times(1)

		_, err := newStringIntSubResource(c).Create("p1", struct{}{})

		assert.NotNil(t, err)
		assert.IsType(t, &parentNotFoundError{}, err)
	})
}

func TestSubResource_Patch(t *testing.T) {
	t.Run("Valid_ReturnsNil", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Patch("/test/v1/parents/p1/children/42", gomock.Any()).Return(
			apiResponseJSON(200, ""), nil).Times(1)

		err := newStringIntSubResource(c).Patch("p1", 42, struct{}{})

		assert.Nil(t, err)
	})

	t.Run("EmptyParentID_ReturnsValidationError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)

		err := newStringIntSubResource(c).Patch("", 42, struct{}{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid parent id", err.Error())
	})

	t.Run("ZeroChildID_ReturnsValidationError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)

		err := newStringIntSubResource(c).Patch("p1", 0, struct{}{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid child id", err.Error())
	})

	t.Run("NotFound_ReturnsChildNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Patch("/test/v1/parents/p1/children/42", gomock.Any()).Return(
			apiResponseJSON(404, ""), nil).Times(1)

		err := newStringIntSubResource(c).Patch("p1", 42, struct{}{})

		assert.NotNil(t, err)
		assert.IsType(t, &childNotFoundError{}, err)
	})
}

func TestSubResource_PatchReturning(t *testing.T) {
	t.Run("Valid_ReturnsItem", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Patch("/test/v1/parents/p1/children/42", gomock.Any()).Return(
			apiResponseJSON(200, `{"data":{"id":42,"name":"updated"}}`), nil).Times(1)

		item, err := newStringIntSubResource(c).PatchReturning("p1", 42, struct{}{})

		assert.Nil(t, err)
		assert.Equal(t, 42, item.ID)
		assert.Equal(t, "updated", item.Name)
	})

	t.Run("EmptyParentID_ReturnsValidationError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)

		_, err := newStringIntSubResource(c).PatchReturning("", 42, struct{}{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid parent id", err.Error())
	})

	t.Run("ZeroChildID_ReturnsValidationError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)

		_, err := newStringIntSubResource(c).PatchReturning("p1", 0, struct{}{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid child id", err.Error())
	})

	t.Run("NotFound_ReturnsChildNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Patch("/test/v1/parents/p1/children/42", gomock.Any()).Return(
			apiResponseJSON(404, ""), nil).Times(1)

		_, err := newStringIntSubResource(c).PatchReturning("p1", 42, struct{}{})

		assert.NotNil(t, err)
		assert.IsType(t, &childNotFoundError{}, err)
	})
}

func TestSubResource_Update(t *testing.T) {
	t.Run("Valid_ReturnsUpdated", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Put("/test/v1/parents/p1/children/42", gomock.Any()).Return(
			apiResponseJSON(200, `{"data":{"id":42,"name":"updated"}}`), nil).Times(1)

		item, err := newStringIntSubResource(c).Update("p1", 42, struct{}{})

		assert.Nil(t, err)
		assert.Equal(t, 42, item.ID)
	})

	t.Run("EmptyParentID_ReturnsValidationError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)

		_, err := newStringIntSubResource(c).Update("", 42, struct{}{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid parent id", err.Error())
	})

	t.Run("ZeroChildID_ReturnsValidationError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)

		_, err := newStringIntSubResource(c).Update("p1", 0, struct{}{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid child id", err.Error())
	})

	t.Run("NotFound_ReturnsChildNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Put("/test/v1/parents/p1/children/42", gomock.Any()).Return(
			apiResponseJSON(404, ""), nil).Times(1)

		_, err := newStringIntSubResource(c).Update("p1", 42, struct{}{})

		assert.NotNil(t, err)
		assert.IsType(t, &childNotFoundError{}, err)
	})
}

func TestSubResource_Delete(t *testing.T) {
	t.Run("Valid_ReturnsNil", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Delete("/test/v1/parents/p1/children/42", gomock.Any()).Return(
			apiResponseJSON(204, ""), nil).Times(1)

		err := newStringIntSubResource(c).Delete("p1", 42)

		assert.Nil(t, err)
	})

	t.Run("EmptyParentID_ReturnsValidationError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)

		err := newStringIntSubResource(c).Delete("", 42)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid parent id", err.Error())
	})

	t.Run("ZeroChildID_ReturnsValidationError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)

		err := newStringIntSubResource(c).Delete("p1", 0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid child id", err.Error())
	})

	t.Run("NotFound_ReturnsChildNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Delete("/test/v1/parents/p1/children/42", gomock.Any()).Return(
			apiResponseJSON(404, ""), nil).Times(1)

		err := newStringIntSubResource(c).Delete("p1", 42)

		assert.NotNil(t, err)
		assert.IsType(t, &childNotFoundError{}, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		c := mocks.NewMockConnection(mockCtrl)
		c.EXPECT().Delete("/test/v1/parents/p1/children/42", gomock.Any()).Return(
			&connection.APIResponse{}, errors.New("test error"))

		err := newStringIntSubResource(c).Delete("p1", 42)

		assert.NotNil(t, err)
		assert.Equal(t, "test error", err.Error())
	})
}
