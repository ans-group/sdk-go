package cloudflare

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetTotalSpendMonthToDate(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/cloudflare/v1/total-spend/month-to-date", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"spend_plan_amount\":1.23, \"total_spend\": 2.34}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		spend, err := s.GetTotalSpendMonthToDate()

		assert.Nil(t, err)
		assert.Equal(t, float32(1.23), spend.SpendPlanAmount)
		assert.Equal(t, float32(2.34), spend.TotalSpend)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/cloudflare/v1/total-spend/month-to-date", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetTotalSpendMonthToDate()

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}
