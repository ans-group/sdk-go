package safedns

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

func TestGetSettings(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/settings", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":123}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		settings, err := s.GetSettings()

		assert.Nil(t, err)
		assert.Equal(t, 123, settings.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := NewService(c)

		c.EXPECT().Get("/safedns/v1/settings", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetSettings()

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}
