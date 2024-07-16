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

func TestGetCases(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/cases", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"PRB123456\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil)

		cases, err := s.GetCases(connection.APIRequestParameters{})

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

		c.EXPECT().Get("/pss/v2/cases", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetCases(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetCase(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/cases/PRB123456", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"PRB123456\"}}"))),
				StatusCode: 200,
			},
		}, nil)

		testcase, err := s.GetCase("PRB123456")

		assert.Nil(t, err)
		assert.Equal(t, "PRB123456", testcase.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/cases/PRB123456", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetCase("PRB123456")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetCase("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid case id", err.Error())
	})

	t.Run("404_ReturnsNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/cases/PRB123456", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		_, err := s.GetCase("PRB123456")

		assert.NotNil(t, err)
		assert.IsType(t, &CaseNotFoundError{}, err)
	})
}

func TestCreateIncidentCase(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		createRequest := CreateIncidentCaseRequest{
			Title:    "testincident",
			CaseType: CaseTypeIncident,
		}

		c.EXPECT().Post("/pss/v2/cases", gomock.Eq(&createRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"INC123456\"}}"))),
				StatusCode: 202,
			},
		}, nil)

		id, err := s.CreateIncidentCase(createRequest)

		assert.Nil(t, err)
		assert.Equal(t, "INC123456", id)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/pss/v2/cases", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.CreateIncidentCase(CreateIncidentCaseRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

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
		}, nil)

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
		}, nil)

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

		c.EXPECT().Get("/pss/v2/cases/INC123456", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

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
		}, nil)

		_, err := s.GetIncidentCase("INC123456")

		assert.NotNil(t, err)
		assert.IsType(t, &CaseNotFoundError{}, err)
	})
}

func TestCloseIncidentCase(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/pss/v2/cases/INC123456/close", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"INC123456\"}}"))),
				StatusCode: 200,
			},
		}, nil)

		incident, err := s.CloseIncidentCase("INC123456", CloseIncidentCaseRequest{})

		assert.Nil(t, err)
		assert.Equal(t, "INC123456", incident)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/pss/v2/cases/INC123456/close", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.CloseIncidentCase("INC123456", CloseIncidentCaseRequest{})

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

		_, err := s.CloseIncidentCase("", CloseIncidentCaseRequest{})

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

		c.EXPECT().Post("/pss/v2/cases/INC123456/close", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		_, err := s.CloseIncidentCase("INC123456", CloseIncidentCaseRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &CaseNotFoundError{}, err)
	})
}

func TestCreateChangeCase(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		createRequest := CreateChangeCaseRequest{
			Title:    "testchange",
			CaseType: CaseTypeChange,
		}

		c.EXPECT().Post("/pss/v2/cases", gomock.Eq(&createRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"CHG123456\"}}"))),
				StatusCode: 202,
			},
		}, nil)

		id, err := s.CreateChangeCase(createRequest)

		assert.Nil(t, err)
		assert.Equal(t, "CHG123456", id)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/pss/v2/cases", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.CreateChangeCase(CreateChangeCaseRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
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
		}, nil)

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

func TestGetChangeCase(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/cases/CHG123456", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"CHG123456\"}}"))),
				StatusCode: 200,
			},
		}, nil)

		change, err := s.GetChangeCase("CHG123456")

		assert.Nil(t, err)
		assert.Equal(t, "CHG123456", change.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/cases/CHG123456", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetChangeCase("CHG123456")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidChangeID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetChangeCase("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid change id", err.Error())
	})

	t.Run("404_ReturnsChangeNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/cases/CHG123456", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		_, err := s.GetChangeCase("CHG123456")

		assert.NotNil(t, err)
		assert.IsType(t, &CaseNotFoundError{}, err)
	})
}

func TestApproveChangeCase(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/pss/v2/cases/CHG123456/approve", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"CHG123456\"}}"))),
				StatusCode: 200,
			},
		}, nil)

		change, err := s.ApproveChangeCase("CHG123456", ApproveChangeCaseRequest{})

		assert.Nil(t, err)
		assert.Equal(t, "CHG123456", change)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/pss/v2/cases/CHG123456/approve", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.ApproveChangeCase("CHG123456", ApproveChangeCaseRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidChangeID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.ApproveChangeCase("", ApproveChangeCaseRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid change id", err.Error())
	})

	t.Run("404_ReturnsChangeNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/pss/v2/cases/CHG123456/approve", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		_, err := s.ApproveChangeCase("CHG123456", ApproveChangeCaseRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &CaseNotFoundError{}, err)
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
		}, nil)

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

func TestGetProblemCase(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/cases/PRB123456", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"PRB123456\"}}"))),
				StatusCode: 200,
			},
		}, nil)

		problem, err := s.GetProblemCase("PRB123456")

		assert.Nil(t, err)
		assert.Equal(t, "PRB123456", problem.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/cases/PRB123456", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetProblemCase("PRB123456")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidProblemID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetProblemCase("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid problem id", err.Error())
	})

	t.Run("404_ReturnsProblemNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/cases/PRB123456", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		_, err := s.GetProblemCase("PRB123456")

		assert.NotNil(t, err)
		assert.IsType(t, &CaseNotFoundError{}, err)
	})
}

func TestGetCaseUpdates(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/cases/PRB123456/updates", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"f99fad6e-cacf-4a47-9258-26cb37108e46\"},{\"id\":\"796473eb-e7eb-4b7c-a39a-6da704ccc41b\"}]}"))),
				StatusCode: 200,
			},
		}, nil)

		updates, err := s.GetCaseUpdates("PRB123456", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, updates, 2)
		assert.Equal(t, "f99fad6e-cacf-4a47-9258-26cb37108e46", updates[0].ID)
		assert.Equal(t, "796473eb-e7eb-4b7c-a39a-6da704ccc41b", updates[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/cases/PRB123456/updates", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetCaseUpdates("PRB123456", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetCaseUpdates("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid case id", err.Error())
	})

	t.Run("404_ReturnsNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/cases/PRB123456/updates", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		_, err := s.GetCaseUpdates("PRB123456", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &CaseNotFoundError{}, err)
	})
}
