package connection

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ans-group/sdk-go/pkg/logging"
)

// APIResponse represents the base API response
type APIResponse struct {
	*http.Response
}

// APIResponseBody represents the base API response body
type APIResponseBody struct {
	APIResponseBodyError

	Metadata APIResponseMetadata `json:"meta"`
}

type APIResponseBodyError struct {
	Errors  []APIResponseBodyErrorItem `json:"errors"`
	Message string                     `json:"message"`
}

func (e *APIResponseBodyError) Error() string {
	var errArr []string

	// Message will be populated in certain circumstances, populate error array if exists
	if e.Message != "" {
		errArr = append(errArr, fmt.Sprintf("message=\"%s\"", e.Message))
	}

	for _, err := range e.Errors {
		errArr = append(errArr, err.Error())
	}

	return strings.Join(errArr, "; ")
}

// APIResponseBodyErrorItem represents an API response error
type APIResponseBodyErrorItem struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
	Source string `json:"source"`
}

func (a *APIResponseBodyErrorItem) Error() string {
	return fmt.Sprintf("title=\"%s\", detail=\"%s\", status=\"%d\", source=\"%s\"", a.Title, a.Detail, a.Status, a.Source)
}

// APIResponseMetadata represents the API response metadata
type APIResponseMetadata struct {
	Pagination APIResponseMetadataPagination `json:"pagination"`
}

// APIResponseMetadataPagination represents the API response pagination data
type APIResponseMetadataPagination struct {
	Total      int                                `json:"total"`
	Count      int                                `json:"count"`
	PerPage    int                                `json:"per_page"`
	TotalPages int                                `json:"total_pages"`
	Links      APIResponseMetadataPaginationLinks `json:"links"`
}

// APIResponseMetadataPaginationLinks represents the links returned within the API response pagination data
type APIResponseMetadataPaginationLinks struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	First    string `json:"first"`
	Last     string `json:"last"`
}

// ValidateStatusCode validates the API response
func (r *APIResponse) ValidateStatusCode(codes ...int) bool {
	if len(codes) > 0 {
		for _, code := range codes {
			if r.StatusCode == code {
				return true
			}
		}
	} else {
		if r.StatusCode >= 200 && r.StatusCode <= 299 {
			return true
		}
	}

	return false
}

type ResponseHandler func(resp *APIResponse) error

func StatusCodeResponseHandler(code int, err error) ResponseHandler {
	return func(resp *APIResponse) error {
		if resp.StatusCode == code {
			return err
		}

		return nil
	}
}

func NotFoundResponseHandler(err error) ResponseHandler {
	return StatusCodeResponseHandler(404, err)
}

type ResponseDeserializer interface {
	Deserialize(r *APIResponse) error
}

func APIResponseJSONDeserializer(r *APIResponse, out interface{}) error {
	defer r.Response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(r.Response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body with response status code %d: %s", r.StatusCode, err)
	}

	logging.Debugf("Response body: %s", string(bodyBytes))

	if len(bodyBytes) > 0 {
		return json.Unmarshal(bodyBytes, out)
	}

	return nil
}

// HandleResponse deserializes the response body into provided respBody, and validates
// the response using the optionally provided ResponseHandler handlers
func (r *APIResponse) HandleResponse(respBody interface{}, handlers ...ResponseHandler) error {
	if respBody != nil {
		if respBodyDeserializer, ok := respBody.(ResponseDeserializer); ok {
			err := respBodyDeserializer.Deserialize(r)
			if err != nil {
				return err
			}
		} else {
			err := APIResponseJSONDeserializer(r, respBody)
			if err != nil {
				return err
			}
		}
	}

	for _, handler := range handlers {
		if handler != nil {
			err := handler(r)
			if err != nil {
				return err
			}
		}
	}

	if !r.ValidateStatusCode() {
		errStr := fmt.Sprintf("unexpected status code (%d)", r.StatusCode)
		err, ok := respBody.(error)
		if ok {
			return fmt.Errorf("%s: %w", errStr, err)
		}

		return errors.New(errStr)
	}

	return nil
}

// APIResponseBodyStringData represents the API response body containing generic data
type APIResponseBodyData[T any] struct {
	APIResponseBody

	Data T `json:"data"`
}

// APIResponseBodyStringData represents the API response body containing string data
type APIResponseBodyStringData struct {
	APIResponseBody

	Data string `json:"data"`
}
