package connection

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/ans-group/sdk-go/test"
	"github.com/stretchr/testify/assert"
)

func TestAPIConnection_composeURI(t *testing.T) {
	t.Run("NoQuery_NoQuerySeparator", func(t *testing.T) {
		c := NewAPIKeyCredentialsAPIConnection("testkey")
		uri := c.composeURI("/some/test/resource", APIRequestPagination{}, APIRequestSorting{}, []APIRequestFiltering{})

		assert.True(t, strings.HasSuffix(uri, "/some/test/resource"))
	})

	t.Run("Query_ExpectedQuerySeparator", func(t *testing.T) {
		c := NewAPIKeyCredentialsAPIConnection("testkey")
		uri := c.composeURI("/some/test/resource", APIRequestPagination{PerPage: 71}, APIRequestSorting{}, []APIRequestFiltering{})

		assert.True(t, strings.HasSuffix(uri, "?per_page=71"))
	})
}

func TestAPIConnection_hydratePaginationQuery_PopulatesQuery(t *testing.T) {
	c := APIConnection{}

	q := &url.Values{}
	p := APIRequestPagination{PerPage: 71, Page: 6}

	c.hydratePaginationQuery(q, p)

	assert.Len(t, *q, 2)
	assert.Equal(t, "71", q.Get("per_page"))
	assert.Equal(t, "6", q.Get("page"))
}

func TestAPIConnection_hydrateFilteringQuery_PopulatesQuery(t *testing.T) {
	t.Run("SingleQuery", func(t *testing.T) {
		q := &url.Values{}
		filtering := []APIRequestFiltering{
			APIRequestFiltering{
				Property: "testproperty",
				Operator: EQOperator,
				Value:    []string{"testvalue"},
			},
		}

		c := APIConnection{}
		c.hydrateFilteringQuery(q, filtering)

		assert.Len(t, *q, 1)
		assert.Equal(t, "testvalue", q.Get("testproperty:eq"))
	})

	t.Run("MultipleQueries", func(t *testing.T) {
		q := &url.Values{}
		filtering := []APIRequestFiltering{
			APIRequestFiltering{
				Property: "testproperty",
				Operator: EQOperator,
				Value:    []string{"testvalue"},
			},
		}

		c := APIConnection{}
		c.hydrateFilteringQuery(q, filtering)

		assert.Len(t, *q, 1)
		assert.Equal(t, "testvalue", q.Get("testproperty:eq"))
	})

	t.Run("SingleQueryMultipleValue", func(t *testing.T) {
		q := &url.Values{}
		filtering := []APIRequestFiltering{
			APIRequestFiltering{
				Property: "testproperty",
				Operator: INOperator,
				Value:    []string{"testvalue1", "testvalue2"},
			},
		}

		c := APIConnection{}
		c.hydrateFilteringQuery(q, filtering)

		assert.Len(t, *q, 1)
		assert.Equal(t, "testvalue1,testvalue2", q.Get("testproperty:in"))
	})
}

func TestAPIConnection_hydrateSortingQuery_PopulatesQuery(t *testing.T) {

	t.Run("Ascending", func(t *testing.T) {
		q := &url.Values{}
		sorting := APIRequestSorting{
			Property: "testproperty",
		}

		c := APIConnection{}
		c.hydrateSortingQuery(q, sorting)

		assert.Len(t, *q, 1)
		assert.Equal(t, "testproperty:asc", q.Get("sort"))
	})

	t.Run("Descending", func(t *testing.T) {
		q := &url.Values{}
		sorting := APIRequestSorting{
			Property:   "testproperty",
			Descending: true,
		}

		c := APIConnection{}
		c.hydrateSortingQuery(q, sorting)

		assert.Len(t, *q, 1)
		assert.Equal(t, "testproperty:desc", q.Get("sort"))
	})
}

func TestAPIConnection_Get_ExpectedMethod(t *testing.T) {
	c := NewAPIKeyCredentialsAPIConnection("testkey")
	c.HTTPClient = test.NewTestClient(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "GET", req.Method)

		return &http.Response{}, nil
	})

	_, err := c.Get("/some/test/resource", APIRequestParameters{})

	assert.Nil(t, err)
}

func TestAPIConnection_Post_ExpectedMethod(t *testing.T) {
	c := NewAPIKeyCredentialsAPIConnection("testkey")
	c.HTTPClient = test.NewTestClient(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "POST", req.Method)

		return &http.Response{}, nil
	})

	_, err := c.Post("/some/test/resource", nil)

	assert.Nil(t, err)
}

func TestAPIConnection_Put_ExpectedMethod(t *testing.T) {
	c := NewAPIKeyCredentialsAPIConnection("testkey")
	c.HTTPClient = test.NewTestClient(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "PUT", req.Method)

		return &http.Response{}, nil
	})

	_, err := c.Put("/some/test/resource", nil)

	assert.Nil(t, err)
}

func TestAPIConnection_Patch_ExpectedMethod(t *testing.T) {
	c := NewAPIKeyCredentialsAPIConnection("testkey")
	c.HTTPClient = test.NewTestClient(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "PATCH", req.Method)

		return &http.Response{}, nil
	})

	_, err := c.Patch("/some/test/resource", nil)

	assert.Nil(t, err)
}

func TestAPIConnection_Delete_ExpectedMethod(t *testing.T) {
	c := NewAPIKeyCredentialsAPIConnection("testkey")
	c.HTTPClient = test.NewTestClient(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "DELETE", req.Method)

		return &http.Response{}, nil
	})

	_, err := c.Delete("/some/test/resource", nil)

	assert.Nil(t, err)
}

func TestAPIConnection_Invoke_WithReader_ExpectedBody(t *testing.T) {
	testRequestBody := ioutil.NopCloser(bytes.NewReader([]byte("test content")))

	c := NewAPIKeyCredentialsAPIConnection("testkey")
	c.HTTPClient = test.NewTestClient(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, testRequestBody, req.Body)

		return &http.Response{}, nil
	})

	_, err := c.Invoke(APIRequest{
		Body: testRequestBody,
	})

	assert.Nil(t, err)
}

type unencodableModel struct {
	Content chan int
}

func (m *unencodableModel) Validate() *ValidationError {
	return nil
}

type modelValidationTest struct {
	ValidationError *ValidationError
}

func (m *modelValidationTest) Validate() *ValidationError {
	return m.ValidationError
}

func TestAPIConnection_Invoke(t *testing.T) {
	c := NewAPIKeyCredentialsAPIConnection("testkey")

	t.Run("ExpectedMethod", func(t *testing.T) {
		c.HTTPClient = test.NewTestClient(func(req *http.Request) (*http.Response, error) {
			assert.Equal(t, "POST", req.Method)

			return &http.Response{}, nil
		})

		_, err := c.Invoke(APIRequest{
			Method: "POST",
		})

		assert.Nil(t, err)
	})

	t.Run("NewRequestError_ReturnsError", func(t *testing.T) {
		c := NewAPIKeyCredentialsAPIConnection("testkey1")
		_, err := c.Invoke(APIRequest{
			Method: "POST",
			Body: &modelValidationTest{
				ValidationError: &ValidationError{
					Message: "test validation error",
				},
			},
		})

		assert.NotNil(t, err)
		assert.Equal(t, "test validation error", err.Error())
	})
}

func TestAPIConnection_NewRequest(t *testing.T) {
	t.Run("InvalidModel_ReturnsError", func(t *testing.T) {
		c := NewAPIKeyCredentialsAPIConnection("testkey1")
		_, err := c.NewRequest(APIRequest{
			Method: "POST",
			Body: &modelValidationTest{
				ValidationError: &ValidationError{
					Message: "test validation error",
				},
			},
		})

		assert.NotNil(t, err)
		assert.Equal(t, "test validation error", err.Error())
	})

	t.Run("JSONEncodingError_ReturnsError", func(t *testing.T) {
		c := NewAPIKeyCredentialsAPIConnection("testkey1")
		_, err := c.NewRequest(APIRequest{
			Method: "POST",
			Body:   &unencodableModel{},
		})

		assert.NotNil(t, err)
		assert.Equal(t, "json: unsupported type: chan int", err.Error())
	})

	t.Run("AddsExpectedHeaders", func(t *testing.T) {
		c := NewAPIKeyCredentialsAPIConnection("testkey1")
		c.UserAgent = "testuseragent1"
		req, err := c.NewRequest(APIRequest{
			Method: "POST",
		})

		assert.Nil(t, err)
		assert.Len(t, req.Header, 4)
		assert.Equal(t, "application/json", req.Header["Content-Type"][0])
		assert.Equal(t, "application/json", req.Header["Accept"][0])
		assert.Equal(t, "testuseragent1", req.Header["User-Agent"][0])
		assert.Equal(t, "testkey1", req.Header["Authorization"][0])
	})

	t.Run("AddsAdditionalHeaders", func(t *testing.T) {
		c := NewAPIKeyCredentialsAPIConnection("testkey1")
		c.Headers = http.Header{}
		c.Headers["X-Test-Header"] = []string{"value1"}
		c.Headers["X-Test-Header-Multi"] = []string{"value1", "value2"}
		req, err := c.NewRequest(APIRequest{
			Method: "POST",
		})

		assert.Nil(t, err)
		assert.Equal(t, "value1", req.Header["X-Test-Header"][0])
		assert.Equal(t, "value1", req.Header["X-Test-Header-Multi"][0])
		assert.Equal(t, "value2", req.Header["X-Test-Header-Multi"][1])
	})

	t.Run("HTTPRequestNewRequestError_ReturnsError", func(t *testing.T) {
		c := NewAPIKeyCredentialsAPIConnection("testkey1")
		c.APIScheme = "invali!d"
		c.UserAgent = "testuseragent1"
		_, err := c.NewRequest(APIRequest{
			Method: "POST",
		})

		assert.NotNil(t, err)
	})
}

func TestAPIConnection_InvokeRequest(t *testing.T) {
	c := NewAPIKeyCredentialsAPIConnection("testkey")

	httpErr := errors.New("test error")

	t.Run("HTTPClientDoError_ReturnsError", func(t *testing.T) {
		c.HTTPClient = test.NewTestClient(func(req *http.Request) (*http.Response, error) {
			return &http.Response{}, httpErr
		})

		url, _ := url.Parse("https://localhost")
		_, err := c.InvokeRequest(&http.Request{URL: url})

		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, httpErr))
	})
}
