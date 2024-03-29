package ecloud

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

func TestGetVPNSessions(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/vpn-sessions", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"vpns-abcdef12\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		sessions, err := s.GetVPNSessions(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, sessions, 1)
		assert.Equal(t, "vpns-abcdef12", sessions[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/vpn-sessions", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetVPNSessions(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetVPNSession(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/vpn-sessions/vpns-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"vpns-abcdef12\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		session, err := s.GetVPNSession("vpns-abcdef12")

		assert.Nil(t, err)
		assert.Equal(t, "vpns-abcdef12", session.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/vpn-sessions/vpns-abcdef12", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetVPNSession("vpns-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVPNSessionID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetVPNSession("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid vpn session id", err.Error())
	})

	t.Run("404_ReturnsVPNSessionNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/vpn-sessions/vpns-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetVPNSession("vpns-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &VPNSessionNotFoundError{}, err)
	})
}

func TestCreateVPNSession(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		req := CreateVPNSessionRequest{
			Name: "test",
		}

		c.EXPECT().Post("/ecloud/v2/vpn-sessions", &req).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"vpns-abcdef12\",\"task_id\":\"task-abcdef12\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		taskRef, err := s.CreateVPNSession(req)

		assert.Nil(t, err)
		assert.Equal(t, "vpns-abcdef12", taskRef.ResourceID)
		assert.Equal(t, "task-abcdef12", taskRef.TaskID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v2/vpn-sessions", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateVPNSession(CreateVPNSessionRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestPatchVPNSession(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		req := PatchVPNSessionRequest{
			Name: "somesession",
		}

		c.EXPECT().Patch("/ecloud/v2/vpn-sessions/vpns-abcdef12", &req).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"vpns-abcdef12\",\"task_id\":\"task-abcdef12\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		task, err := s.PatchVPNSession("vpns-abcdef12", req)

		assert.Nil(t, err)
		assert.Equal(t, "vpns-abcdef12", task.ResourceID)
		assert.Equal(t, "task-abcdef12", task.TaskID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ecloud/v2/vpn-sessions/vpns-abcdef12", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.PatchVPNSession("vpns-abcdef12", PatchVPNSessionRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVPNSessionID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.PatchVPNSession("", PatchVPNSessionRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid vpn session id", err.Error())
	})

	t.Run("404_ReturnsVPNSessionNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ecloud/v2/vpn-sessions/vpns-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.PatchVPNSession("vpns-abcdef12", PatchVPNSessionRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &VPNSessionNotFoundError{}, err)
	})
}

func TestDeleteVPNSession(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v2/vpn-sessions/vpns-abcdef12", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"task_id\":\"task-abcdef12\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		taskID, err := s.DeleteVPNSession("vpns-abcdef12")

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

		c.EXPECT().Delete("/ecloud/v2/vpn-sessions/vpns-abcdef12", nil).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.DeleteVPNSession("vpns-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVPNSessionID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.DeleteVPNSession("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid vpn session id", err.Error())
	})

	t.Run("404_ReturnsVPNSessionNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v2/vpn-sessions/vpns-abcdef12", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.DeleteVPNSession("vpns-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &VPNSessionNotFoundError{}, err)
	})
}

func TestGetVPNSessionTasks(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/vpn-sessions/vpns-abcdef12/tasks", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"task-abcdef12\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		tasks, err := s.GetVPNSessionTasks("vpns-abcdef12", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, tasks, 1)
		assert.Equal(t, "task-abcdef12", tasks[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/vpn-sessions/vpns-abcdef12/tasks", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetVPNSessionTasks("vpns-abcdef12", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVPNSessionID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetVPNSessionTasks("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid vpn session id", err.Error())
	})

	t.Run("404_ReturnsRouterNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/vpn-sessions/vpns-abcdef12/tasks", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetVPNSessionTasks("vpns-abcdef12", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &VPNSessionNotFoundError{}, err)
	})
}

func TestGetVPNSessionPreSharedKey(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/vpn-sessions/vpns-abcdef12/pre-shared-key", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"psk\":\"testpsk\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		session, err := s.GetVPNSessionPreSharedKey("vpns-abcdef12")

		assert.Nil(t, err)
		assert.Equal(t, "testpsk", session.PSK)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/vpn-sessions/vpns-abcdef12/pre-shared-key", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetVPNSessionPreSharedKey("vpns-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVPNSessionPreSharedKeyID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetVPNSessionPreSharedKey("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid vpn session id", err.Error())
	})

	t.Run("404_ReturnsVPNSessionNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/vpn-sessions/vpns-abcdef12/pre-shared-key", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetVPNSessionPreSharedKey("vpns-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &VPNSessionNotFoundError{}, err)
	})
}

func TestUpdateVPNSessionPreSharedKey(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		req := UpdateVPNSessionPreSharedKeyRequest{
			PSK: "somepsk",
		}

		c.EXPECT().Put("/ecloud/v2/vpn-sessions/vpns-abcdef12/pre-shared-key", &req).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"vpns-abcdef12\",\"task_id\":\"task-abcdef12\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		task, err := s.UpdateVPNSessionPreSharedKey("vpns-abcdef12", req)

		assert.Nil(t, err)
		assert.Equal(t, "vpns-abcdef12", task.ResourceID)
		assert.Equal(t, "task-abcdef12", task.TaskID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/vpn-sessions/vpns-abcdef12/pre-shared-key", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.UpdateVPNSessionPreSharedKey("vpns-abcdef12", UpdateVPNSessionPreSharedKeyRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVPNSessionID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.UpdateVPNSessionPreSharedKey("", UpdateVPNSessionPreSharedKeyRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid vpn session id", err.Error())
	})

	t.Run("404_ReturnsVPNSessionNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/vpn-sessions/vpns-abcdef12/pre-shared-key", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.UpdateVPNSessionPreSharedKey("vpns-abcdef12", UpdateVPNSessionPreSharedKeyRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &VPNSessionNotFoundError{}, err)
	})
}
