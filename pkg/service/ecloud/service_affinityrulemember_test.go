package ecloud

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

func TestGetAffinityRuleMembers(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/affinity-rules/ar-abcdef12/members", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"arm-abcdef12\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		members, err := s.GetAffinityRuleMembers("ar-abcdef12", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, members, 1)
		assert.Equal(t, "arm-abcdef12", members[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/affinity-rules/ar-abcdef12/members", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetAffinityRuleMembers("ar-abcdef12", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetAffinityRuleMember(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/affinity-rules/ar-abcdef12/members/arm-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"arm-abcdef12\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		member, err := s.GetAffinityRuleMember("ar-abcdef12", "arm-abcdef12")

		assert.Nil(t, err)
		assert.Equal(t, "arm-abcdef12", member.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/affinity-rules/ar-abcdef12/members/arm-abcdef12", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetAffinityRuleMember("ar-abcdef12", "arm-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidAffinityRuleMemberID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetAffinityRuleMember("ar-abcdef12", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid affinity rule member id", err.Error())
	})

	t.Run("404_ReturnsAffinityRuleMemberNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/affinity-rules/ar-abcdef12/members/arm-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetAffinityRuleMember("ar-abcdef12", "arm-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &AffinityRuleMemberNotFoundError{}, err)
	})
}

func TestCreateAffinityRuleMember(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		req := CreateAffinityRuleMemberRequest{
			InstanceID: "i-abcdef12",
		}

		c.EXPECT().Post("/ecloud/v2/affinity-rules/ar-abcdef12/members", &req).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"arm-abcdef12\",\"task_id\":\"task-abcdef12\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		taskRef, err := s.CreateAffinityRuleMember("ar-abcdef12", req)

		assert.Nil(t, err)
		assert.Equal(t, "arm-abcdef12", taskRef.ResourceID)
		assert.Equal(t, "task-abcdef12", taskRef.TaskID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v2/affinity-rules/ar-abcdef12/members", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateAffinityRuleMember("ar-abcdef12", CreateAffinityRuleMemberRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestDeleteAffinityRuleMember(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v2/affinity-rules/ar-abcdef12/members/arm-abcdef12", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"task_id\":\"task-abcdef12\"}}"))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		taskID, err := s.DeleteAffinityRuleMember("ar-abcdef12", "arm-abcdef12")

		assert.Nil(t, err)
		assert.Equal(t, "task-abcdef12", taskID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v2/affinity-rules/ar-abcdef12/members/arm-abcdef12", nil).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.DeleteAffinityRuleMember("ar-abcdef12", "arm-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidAffinityRuleMemberID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.DeleteAffinityRuleMember("ar-abcdef12", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid affinity rule member id", err.Error())
	})

	t.Run("404_ReturnsAffinityRuleMemberNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v2/affinity-rules/ar-abcdef12/members/arm-abcdef12", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.DeleteAffinityRuleMember("ar-abcdef12", "arm-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &AffinityRuleMemberNotFoundError{}, err)
	})
}
