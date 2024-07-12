package pss

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetIncidentCases(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		p := connection.APIRequestParameters{
			Pagination: connection.APIRequestPagination{
				Page: 1,
			},
		}
		request := connection.APIRequest{
			Method:     "GET",
			Resource:   "/pss/v2/cases",
			Query:      url.Values{"case_type": []string{"incident"}},
			Parameters: p,
		}

		c.EXPECT().Invoke(gomock.Eq(request)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"INC123456\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		cases, err := s.GetIncidentCases(p)

		assert.Nil(t, err)
		assert.Len(t, cases, 1)
		assert.Equal(t, "INC123456", cases[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		p := connection.APIRequestParameters{
			Pagination: connection.APIRequestPagination{
				Page: 1,
			},
		}
		request := connection.APIRequest{
			Method:     "GET",
			Resource:   "/pss/v2/cases",
			Query:      url.Values{"case_type": []string{"incident"}},
			Parameters: p,
		}

		c.EXPECT().Invoke(request).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetIncidentCases(p)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetIncidentCase(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/cases/INC123456", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"INC123456\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		incident, err := s.GetIncidentCase("INC123456")

		assert.Nil(t, err)
		assert.Equal(t, "INC123456", incident.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/cases/INC123456", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetIncidentCase("INC123456")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidIncidentID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetIncidentCase("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid incident id", err.Error())
	})

	t.Run("404_ReturnsIncidentNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/cases/INC123456", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetIncidentCase("INC123456")

		assert.NotNil(t, err)
		assert.IsType(t, &IncidentCaseNotFoundError{}, err)
	})
}

func TestGetChangeCases(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		p := connection.APIRequestParameters{
			Pagination: connection.APIRequestPagination{
				Page: 1,
			},
		}
		request := connection.APIRequest{
			Method:     "GET",
			Resource:   "/pss/v2/cases",
			Query:      url.Values{"case_type": []string{"change"}},
			Parameters: p,
		}

		c.EXPECT().Invoke(gomock.Eq(request)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"CHG123456\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		cases, err := s.GetChangeCases(p)

		assert.Nil(t, err)
		assert.Len(t, cases, 1)
		assert.Equal(t, "CHG123456", cases[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		p := connection.APIRequestParameters{
			Pagination: connection.APIRequestPagination{
				Page: 1,
			},
		}
		request := connection.APIRequest{
			Method:     "GET",
			Resource:   "/pss/v2/cases",
			Query:      url.Values{"case_type": []string{"change"}},
			Parameters: p,
		}

		c.EXPECT().Invoke(request).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetChangeCases(p)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetProblemCases(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		p := connection.APIRequestParameters{
			Pagination: connection.APIRequestPagination{
				Page: 1,
			},
		}
		request := connection.APIRequest{
			Method:     "GET",
			Resource:   "/pss/v2/cases",
			Query:      url.Values{"case_type": []string{"problem"}},
			Parameters: p,
		}

		c.EXPECT().Invoke(gomock.Eq(request)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"PRB123456\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		cases, err := s.GetProblemCases(p)

		assert.Nil(t, err)
		assert.Len(t, cases, 1)
		assert.Equal(t, "PRB123456", cases[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		p := connection.APIRequestParameters{
			Pagination: connection.APIRequestPagination{
				Page: 1,
			},
		}
		request := connection.APIRequest{
			Method:     "GET",
			Resource:   "/pss/v2/cases",
			Query:      url.Values{"case_type": []string{"problem"}},
			Parameters: p,
		}

		c.EXPECT().Invoke(request).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetProblemCases(p)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}
