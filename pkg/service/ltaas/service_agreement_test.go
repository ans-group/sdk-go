package ltaas_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/ukfast/sdk-go/pkg/connection"
	"github.com/ukfast/sdk-go/pkg/service/ltaas"
	"github.com/ukfast/sdk-go/test/mocks"
)

func TestGetAgreement(t *testing.T) {
	t.Run("SuccessfulRequest_ReturnsExpected", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := ltaas.NewService(c)

		c.EXPECT().Get("/ltaas/v1/agreements/latest/single", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"version\":\"v1.0\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		agreement, err := s.GetLatestAgreement(ltaas.AgreementTypeSingle)

		assert.Nil(t, err)
		assert.Equal(t, "v1.0", agreement.Version)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := ltaas.NewService(c)

		c.EXPECT().Get("/ltaas/v1/agreements/latest/single", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetLatestAgreement(ltaas.AgreementTypeSingle)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidAgreementID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := ltaas.NewService(c)

		_, err := s.GetLatestAgreement(ltaas.AgreementType(""))

		assert.NotNil(t, err)
		assert.Equal(t, "invalid agreement id", err.Error())
	})

	t.Run("404_ReturnsAgreementNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := ltaas.NewService(c)

		c.EXPECT().Get("/ltaas/v1/agreements/latest/invalid", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetLatestAgreement(ltaas.AgreementType("invalid"))

		assert.NotNil(t, err)
		assert.IsType(t, &ltaas.AgreementNotFoundError{}, err)
	})
}
