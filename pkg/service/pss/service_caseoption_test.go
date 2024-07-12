package pss

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

func TestGetChangeImpactCaseOptions(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/case-options/change-impacts", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"option\":\"Major\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		options, err := s.GetChangeImpactCaseOptions(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, options, 1)
		assert.Equal(t, "Major", options[0].Option)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/case-options/change-impacts", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetChangeImpactCaseOptions(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetChangeRiskCaseOptions(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/case-options/change-risks", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"option\":\"Major\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		options, err := s.GetChangeRiskCaseOptions(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, options, 1)
		assert.Equal(t, "Major", options[0].Option)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/case-options/change-risks", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetChangeRiskCaseOptions(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetIncidentImpactCaseOptions(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/case-options/incident-impacts", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"option\":\"Major\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		options, err := s.GetIncidentImpactCaseOptions(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, options, 1)
		assert.Equal(t, "Major", options[0].Option)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/case-options/incident-impacts", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetIncidentImpactCaseOptions(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetIncidentTypeCaseOptions(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/case-options/incident-types", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"option\":\"Major\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		options, err := s.GetIncidentTypeCaseOptions(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, options, 1)
		assert.Equal(t, "Major", options[0].Option)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/pss/v2/case-options/incident-types", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetIncidentTypeCaseOptions(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}
