package connection

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ukfast/sdk-go/test"
)

func TestAPIResponse_DeserializeResponseBody(t *testing.T) {
	t.Run("ExpectedOutput", func(t *testing.T) {
		type testout struct {
			TestProperty1 string `json:"testproperty1"`
		}

		b := ioutil.NopCloser(bytes.NewReader([]byte("{\"testproperty1\":\"testvalue1\"}")))
		resp := APIResponse{
			Response: &http.Response{
				Body: b,
			},
		}

		out := testout{}

		err := resp.DeserializeResponseBody(&out)

		assert.Nil(t, err)
		assert.Equal(t, "testvalue1", out.TestProperty1)
	})

	t.Run("NoBody_ReturnsNil", func(t *testing.T) {
		type testout struct{}

		b := ioutil.NopCloser(bytes.NewReader([]byte{}))
		resp := APIResponse{
			Response: &http.Response{
				Body: b,
			},
		}

		out := testout{}

		err := resp.DeserializeResponseBody(&out)

		assert.Nil(t, err)
	})

	t.Run("ReaderError_ReturnsError", func(t *testing.T) {
		type testout struct{}

		b := test.TestReadCloser{
			ReadError: errors.New("test reader error 1"),
		}

		resp := APIResponse{
			Response: &http.Response{
				StatusCode: 500,
				Body:       &b,
			},
		}

		out := testout{}

		err := resp.DeserializeResponseBody(&out)

		assert.NotNil(t, err)
		assert.Equal(t, "failed to read response body with response status code 500: test reader error 1", err.Error())
	})
}

func TestAPIResponse_ValidateStatusCode(t *testing.T) {
	t.Run("SingleCodeExpected_NoError", func(t *testing.T) {
		resp := APIResponse{
			Response: &http.Response{
				StatusCode: 200,
			},
		}

		err := resp.ValidateStatusCode([]int{200}, &APIResponseBody{})

		assert.Nil(t, err)
	})

	t.Run("MultipleCodesExpected_NoError", func(t *testing.T) {
		resp := APIResponse{
			Response: &http.Response{
				StatusCode: 201,
			},
		}

		err := resp.ValidateStatusCode([]int{200, 201}, &APIResponseBody{})

		assert.Nil(t, err)
	})

	t.Run("SingleCodeUnexpected_Error", func(t *testing.T) {
		resp := APIResponse{
			Response: &http.Response{
				StatusCode: 500,
			},
		}

		err := resp.ValidateStatusCode([]int{200}, &APIResponseBody{})

		assert.NotNil(t, err)
	})

	t.Run("MultipleCodesUnexpected_Error", func(t *testing.T) {
		resp := APIResponse{
			Response: &http.Response{
				StatusCode: 500,
			},
		}

		err := resp.ValidateStatusCode([]int{200, 201}, &APIResponseBody{})

		assert.NotNil(t, err)
	})

	t.Run("NoCodeExpected_ValidStatusCode_NoError", func(t *testing.T) {
		resp := APIResponse{
			Response: &http.Response{
				StatusCode: 200,
			},
		}

		err := resp.ValidateStatusCode([]int{}, &APIResponseBody{})

		assert.Nil(t, err)
	})

	t.Run("NoCodeExpected_InvalidStatusCode_Error", func(t *testing.T) {
		resp := APIResponse{
			Response: &http.Response{
				StatusCode: 500,
			},
		}

		err := resp.ValidateStatusCode([]int{}, &APIResponseBody{})

		assert.NotNil(t, err)
	})
}

func TestAPIResponse_HandleResponse(t *testing.T) {
	t.Run("ValidResponseCode_NoError", func(t *testing.T) {
		resp := APIResponse{
			Response: &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
			},
		}

		err := resp.HandleResponse(&APIResponseBody{}, nil)

		assert.Nil(t, err)
	})

	t.Run("DeserializeResponseBodyError_ReturnsError", func(t *testing.T) {
		resp := APIResponse{
			Response: &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("invalidjson"))),
			},
		}

		err := resp.HandleResponse(&APIResponseBody{}, nil)

		assert.NotNil(t, err)
	})

	t.Run("InvalidResponseCode_ReturnsError", func(t *testing.T) {
		resp := APIResponse{
			Response: &http.Response{
				StatusCode: 500,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
			},
		}

		err := resp.HandleResponse(&APIResponseBody{}, nil)

		assert.NotNil(t, err)
	})

	t.Run("HandlerError_ReturnsError", func(t *testing.T) {
		resp := APIResponse{
			Response: &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
			},
		}

		err := resp.HandleResponse(&APIResponseBody{}, func(resp *APIResponse) error {
			return errors.New("test error")
		})

		assert.NotNil(t, err)
		assert.Equal(t, "test error", err.Error())
	})
}

func TestAPIResponseError_String(t *testing.T) {
	apiError := APIResponseError{
		Detail: "test detail 1",
		Source: "test source 1",
		Status: 500,
		Title:  "test title 1",
	}
	expected := "title=\"test title 1\", detail=\"test detail 1\", status=\"500\", source=\"test source 1\""

	t.Run("String_Expected", func(t *testing.T) {
		s := apiError.String()

		assert.Equal(t, expected, s)
	})

	t.Run("Error_Expected", func(t *testing.T) {
		s := apiError.Error()

		assert.Equal(t, expected, s.Error())
	})
}

func TestAPIResponseBody_ErrorString(t *testing.T) {
	t.Run("SingleError", func(t *testing.T) {
		b := APIResponseBody{
			Errors: []APIResponseError{
				APIResponseError{
					Detail: "test detail 1",
					Source: "test source 1",
					Status: 500,
					Title:  "test title 1",
				},
			},
		}

		msg := b.ErrorString()

		assert.Equal(t, "title=\"test title 1\", detail=\"test detail 1\", status=\"500\", source=\"test source 1\"", msg)
	})

	t.Run("MultipleErrors", func(t *testing.T) {
		b := APIResponseBody{
			Errors: []APIResponseError{
				APIResponseError{
					Detail: "test detail 1",
					Source: "test source 1",
					Status: 500,
					Title:  "test title 1",
				},
				APIResponseError{
					Detail: "test detail 2",
					Source: "test source 2",
					Status: 501,
					Title:  "test title 2",
				},
			},
		}

		msg := b.ErrorString()

		assert.Equal(t, "title=\"test title 1\", detail=\"test detail 1\", status=\"500\", source=\"test source 1\"; title=\"test title 2\", detail=\"test detail 2\", status=\"501\", source=\"test source 2\"", msg)
	})

	t.Run("WithMessage", func(t *testing.T) {
		b := APIResponseBody{
			Errors: []APIResponseError{
				APIResponseError{
					Detail: "test detail 1",
					Source: "test source 1",
					Status: 500,
					Title:  "test title 1",
				},
			},
			Message: "test message 1",
		}

		msg := b.ErrorString()

		assert.Equal(t, "message=\"test message 1\"; title=\"test title 1\", detail=\"test detail 1\", status=\"500\", source=\"test source 1\"", msg)
	})
}

func TestAPIResponseBody_Pagination_ReturnsPagination(t *testing.T) {
	b := APIResponseBody{
		Metadata: APIResponseMetadata{
			Pagination: APIResponseMetadataPagination{
				Count:      1,
				PerPage:    2,
				TotalPages: 3,
			},
		},
	}

	pagination := b.Pagination()

	assert.Equal(t, 1, pagination.Count)
	assert.Equal(t, 2, pagination.PerPage)
	assert.Equal(t, 3, pagination.TotalPages)
}
