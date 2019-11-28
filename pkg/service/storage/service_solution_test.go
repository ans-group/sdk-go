package storage

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

func TestGetSolutions(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ukfast-storage/v1/solutions", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		solutions, err := s.GetSolutions(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, solutions, 1)
		assert.Equal(t, 123, solutions[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ukfast-storage/v1/solutions", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ukfast-storage/v1/solutions", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":456}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		solutions, err := s.GetSolutions(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, solutions, 2)
		assert.Equal(t, 123, solutions[0].ID)
		assert.Equal(t, 456, solutions[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ukfast-storage/v1/solutions", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetSolutions(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetSolution(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ukfast-storage/v1/solutions/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":123}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		solution, err := s.GetSolution(123)

		assert.Nil(t, err)
		assert.Equal(t, 123, solution.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ukfast-storage/v1/solutions/123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetSolution(123)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidSolutionID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetSolution(0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid solution id", err.Error())
	})

	t.Run("404_ReturnsSolutionNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ukfast-storage/v1/solutions/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetSolution(123)

		assert.NotNil(t, err)
		assert.IsType(t, &SolutionNotFoundError{}, err)
	})
}
