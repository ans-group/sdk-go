package ecloud

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

func TestGetInstances(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"i-abcdef12\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		instances, err := s.GetInstances(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, instances, 1)
		assert.Equal(t, "i-abcdef12", instances[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetInstances(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetInstance(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"i-abcdef12\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		instance, err := s.GetInstance("i-abcdef12")

		assert.Nil(t, err)
		assert.Equal(t, "i-abcdef12", instance.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetInstance("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})

	t.Run("404_ReturnsInstanceNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &InstanceNotFoundError{}, err)
	})
}

func TestCreateInstance(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		createRequest := CreateInstanceRequest{
			Name:        "test instance 1",
			VCPUCores:   2,
			RAMCapacity: 4,
		}

		c.EXPECT().Post("/ecloud/v2/instances", gomock.Eq(&createRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"i-abcdef12\"}}"))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		id, err := s.CreateInstance(createRequest)

		assert.Nil(t, err)
		assert.Equal(t, "i-abcdef12", id)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v2/instances", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateInstance(CreateInstanceRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestPatchInstance(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		patchRequest := PatchInstanceRequest{
			Name:        "test instance 1",
			VCPUCores:   2,
			RAMCapacity: 4,
		}

		c.EXPECT().Patch("/ecloud/v2/instances/i-abcdef12", gomock.Eq(&patchRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
				StatusCode: 202,
			},
		}, nil)

		err := s.PatchInstance("i-abcdef12", patchRequest)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ecloud/v2/instances/i-abcdef12", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchInstance("i-abcdef12", PatchInstanceRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchInstance("", PatchInstanceRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})

	t.Run("404_ReturnsInstanceNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ecloud/v2/instances/i-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		err := s.PatchInstance("i-abcdef12", PatchInstanceRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &InstanceNotFoundError{}, err)
	})
}

func TestDeleteInstance(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v2/instances/i-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		err := s.DeleteInstance("i-abcdef12")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v2/instances/i-abcdef12", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteInstance("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})

	t.Run("404_ReturnsInstanceNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v2/instances/i-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &InstanceNotFoundError{}, err)
	})
}

func TestLockInstance(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/lock", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		err := s.LockInstance("i-abcdef12")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/lock", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.LockInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.LockInstance("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})

	t.Run("404_ReturnsInstanceNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/lock", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.LockInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &InstanceNotFoundError{}, err)
	})
}

func TestUnlockInstance(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/unlock", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		err := s.UnlockInstance("i-abcdef12")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/unlock", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.UnlockInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.UnlockInstance("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})

	t.Run("404_ReturnsInstanceNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/unlock", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.UnlockInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &InstanceNotFoundError{}, err)
	})
}

func TestPowerOnInstance(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/power-on", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"task_id\":\"task-abcdef12\"},\"meta\":{\"location\":\"\"}}"))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		taskID, err := s.PowerOnInstance("i-abcdef12")

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

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/power-on", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.PowerOnInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.PowerOnInstance("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})

	t.Run("404_ReturnsInstanceNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/power-on", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.PowerOnInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &InstanceNotFoundError{}, err)
	})
}

func TestPowerOffInstance(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/power-off", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"task_id\":\"task-abcdef12\"},\"meta\":{\"location\":\"\"}}"))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		taskID, err := s.PowerOffInstance("i-abcdef12")

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

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/power-off", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.PowerOffInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.PowerOffInstance("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})

	t.Run("404_ReturnsInstanceNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/power-off", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.PowerOffInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &InstanceNotFoundError{}, err)
	})
}

func TestPowerResetInstance(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/power-reset", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"task_id\":\"task-abcdef12\"},\"meta\":{\"location\":\"\"}}"))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		taskID, err := s.PowerResetInstance("i-abcdef12")

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

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/power-reset", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.PowerResetInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.PowerResetInstance("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})

	t.Run("404_ReturnsInstanceNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/power-reset", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.PowerResetInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &InstanceNotFoundError{}, err)
	})
}

func TestPowerShutdownInstance(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/power-shutdown", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"task_id\":\"task-abcdef12\"},\"meta\":{\"location\":\"\"}}"))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		taskID, err := s.PowerShutdownInstance("i-abcdef12")

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

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/power-shutdown", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.PowerShutdownInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.PowerShutdownInstance("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})

	t.Run("404_ReturnsInstanceNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/power-shutdown", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.PowerShutdownInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &InstanceNotFoundError{}, err)
	})
}

func TestPowerRestartInstance(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/power-restart", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"task_id\":\"task-abcdef12\"},\"meta\":{\"location\":\"\"}}"))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		taskID, err := s.PowerRestartInstance("i-abcdef12")

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

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/power-restart", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.PowerRestartInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.PowerRestartInstance("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})

	t.Run("404_ReturnsInstanceNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/power-restart", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.PowerRestartInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &InstanceNotFoundError{}, err)
	})
}

func TestMigrateInstance(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		req := MigrateInstanceRequest{
			HostGroupID: "hg-abcdef12",
		}

		c.EXPECT().Post("/ecloud/v2/instances/i-abcdef12/migrate", &req).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"task_id\":\"task-abcdef12\"},\"meta\":{\"location\":\"\"}}"))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		taskID, err := s.MigrateInstance("i-abcdef12", req)

		assert.Nil(t, err)
		assert.Equal(t, "task-abcdef12", taskID)
	})

	t.Run("ValidResourceTier", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		req := MigrateInstanceRequest{
			ResourceTierID: "rt-abcdef12",
		}

		c.EXPECT().Post("/ecloud/v2/instances/i-abcdef12/migrate", &req).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"task_id\":\"task-abcdef12\"},\"meta\":{\"location\":\"\"}}"))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		taskID, err := s.MigrateInstance("i-abcdef12", req)

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

		c.EXPECT().Post("/ecloud/v2/instances/i-abcdef12/migrate", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.MigrateInstance("i-abcdef12", MigrateInstanceRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.MigrateInstance("", MigrateInstanceRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})

	t.Run("404_ReturnsInstanceNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v2/instances/i-abcdef12/migrate", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.MigrateInstance("i-abcdef12", MigrateInstanceRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &InstanceNotFoundError{}, err)
	})
}

func TestGetInstanceVolumes(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12/volumes", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"vol-abcdef12\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil)

		instances, err := s.GetInstanceVolumes("i-abcdef12", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, instances, 1)
		assert.Equal(t, "vol-abcdef12", instances[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12/volumes", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetInstanceVolumes("i-abcdef12", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetInstanceVolumes("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})

	t.Run("404_ReturnsInstanceNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12/volumes", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		_, err := s.GetInstanceVolumes("i-abcdef12", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &InstanceNotFoundError{}, err)
	})
}

func TestGetInstanceCredentials(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12/credentials", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"vol-abcdef12\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil)

		instances, err := s.GetInstanceCredentials("i-abcdef12", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, instances, 1)
		assert.Equal(t, "vol-abcdef12", instances[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12/credentials", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetInstanceCredentials("i-abcdef12", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetInstanceCredentials("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})

	t.Run("404_ReturnsInstanceNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12/credentials", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		_, err := s.GetInstanceCredentials("i-abcdef12", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &InstanceNotFoundError{}, err)
	})
}

func TestGetInstanceNICs(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12/nics", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"vol-abcdef12\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil)

		instances, err := s.GetInstanceNICs("i-abcdef12", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, instances, 1)
		assert.Equal(t, "vol-abcdef12", instances[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12/nics", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetInstanceNICs("i-abcdef12", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetInstanceNICs("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})

	t.Run("404_ReturnsInstanceNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12/nics", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		_, err := s.GetInstanceNICs("i-abcdef12", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &InstanceNotFoundError{}, err)
	})
}

func TestCreateInstanceConsoleSession(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v2/instances/i-abcdef12/console-session", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"url\":\"ukfast.co.uk\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		session, err := s.CreateInstanceConsoleSession("i-abcdef12")

		assert.Nil(t, err)
		assert.Equal(t, "ukfast.co.uk", session.URL)
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.CreateInstanceConsoleSession("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v2/instances/i-abcdef12/console-session", nil).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateInstanceConsoleSession("i-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsInstanceNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v2/instances/i-abcdef12/console-session", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.CreateInstanceConsoleSession("i-abcdef12")

		assert.NotNil(t, err)
		assert.IsType(t, &InstanceNotFoundError{}, err)
	})
}

func TestGetInstanceTasks(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12/tasks", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"task-abcdef12\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		tasks, err := s.GetInstanceTasks("i-abcdef12", connection.APIRequestParameters{})

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

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12/tasks", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetInstanceTasks("i-abcdef12", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetInstanceTasks("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})

	t.Run("404_ReturnsRouterNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12/tasks", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetInstanceTasks("i-abcdef12", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &InstanceNotFoundError{}, err)
	})
}

func TestAttachInstanceVolume(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		patchRequest := AttachDetachInstanceVolumeRequest{
			VolumeID: "vol-abcdef12",
		}

		c.EXPECT().Post("/ecloud/v2/instances/i-abcdef12/volume-attach", gomock.Eq(&patchRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"task_id\":\"task-abcdef12\"},\"meta\":{\"location\":\"\"}}"))),
				StatusCode: 202,
			},
		}, nil)

		taskID, err := s.AttachInstanceVolume("i-abcdef12", patchRequest)

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

		c.EXPECT().Post("/ecloud/v2/instances/i-abcdef12/volume-attach", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.AttachInstanceVolume("i-abcdef12", AttachDetachInstanceVolumeRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.AttachInstanceVolume("", AttachDetachInstanceVolumeRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})

	t.Run("404_ReturnsInstanceNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v2/instances/i-abcdef12/volume-attach", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		_, err := s.AttachInstanceVolume("i-abcdef12", AttachDetachInstanceVolumeRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &InstanceNotFoundError{}, err)
	})
}

func TestDetachInstanceVolume(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		patchRequest := AttachDetachInstanceVolumeRequest{
			VolumeID: "vol-abcdef12",
		}

		c.EXPECT().Post("/ecloud/v2/instances/i-abcdef12/volume-detach", gomock.Eq(&patchRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"task_id\":\"task-abcdef12\"},\"meta\":{\"location\":\"\"}}"))),
				StatusCode: 202,
			},
		}, nil)

		taskID, err := s.DetachInstanceVolume("i-abcdef12", patchRequest)

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

		c.EXPECT().Post("/ecloud/v2/instances/i-abcdef12/volume-detach", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.DetachInstanceVolume("i-abcdef12", AttachDetachInstanceVolumeRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.DetachInstanceVolume("", AttachDetachInstanceVolumeRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})

	t.Run("404_ReturnsInstanceNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v2/instances/i-abcdef12/volume-detach", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		_, err := s.DetachInstanceVolume("i-abcdef12", AttachDetachInstanceVolumeRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &InstanceNotFoundError{}, err)
	})
}

func TestGetInstanceFloatingIPs(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12/floating-ips", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"vol-abcdef12\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil)

		instances, err := s.GetInstanceFloatingIPs("i-abcdef12", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, instances, 1)
		assert.Equal(t, "vol-abcdef12", instances[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12/floating-ips", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetInstanceFloatingIPs("i-abcdef12", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetInstanceFloatingIPs("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})

	t.Run("404_ReturnsInstanceNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12/floating-ips", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		_, err := s.GetInstanceFloatingIPs("i-abcdef12", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &InstanceNotFoundError{}, err)
	})
}

func TestCreateInstanceImage(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		req := CreateInstanceImageRequest{
			Name: "testimage",
		}

		c.EXPECT().Post("/ecloud/v2/instances/i-abcdef12/create-image", gomock.Eq(&req)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"task_id\":\"task-abcdef12\"},\"meta\":{\"location\":\"\"}}"))),
				StatusCode: 202,
			},
		}, nil)

		task, err := s.CreateInstanceImage("i-abcdef12", req)

		assert.Nil(t, err)
		assert.Equal(t, "task-abcdef12", task.TaskID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v2/instances/i-abcdef12/create-image", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateInstanceImage("i-abcdef12", CreateInstanceImageRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.CreateInstanceImage("", CreateInstanceImageRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})
}

func TestEncryptInstance(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/encrypt", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"task_id\":\"task-abcdef12\"},\"meta\":{\"location\":\"\"}}"))),
				StatusCode: 202,
			},
		}, nil)

		taskID, err := s.EncryptInstance("i-abcdef12")

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

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/encrypt", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.EncryptInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.EncryptInstance("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})
}

func TestDecryptInstance(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/decrypt", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"task_id\":\"task-abcdef12\"},\"meta\":{\"location\":\"\"}}"))),
				StatusCode: 202,
			},
		}, nil)

		taskID, err := s.DecryptInstance("i-abcdef12")

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

		c.EXPECT().Put("/ecloud/v2/instances/i-abcdef12/decrypt", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.DecryptInstance("i-abcdef12")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.DecryptInstance("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})
}

func TestExecuteInstanceScript(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v2/instances/i-abcdef12/user-script", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"task_id\":\"task-abcdef12\"},\"meta\":{\"location\":\"\"}}"))),
				StatusCode: 202,
			},
		}, nil)

		executeRequest := ExecuteInstanceScriptRequest{
			Script:   "test script 1",
			Username: "test_user",
			Password: "test_password",
		}

		taskID, err := s.ExecuteInstanceScript("i-abcdef12", executeRequest)

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

		c.EXPECT().Post("/ecloud/v2/instances/i-abcdef12/user-script", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		executeRequest := ExecuteInstanceScriptRequest{
			Script:   "test script 1",
			Username: "test_user",
			Password: "test_password",
		}

		_, err := s.ExecuteInstanceScript("i-abcdef12", executeRequest)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidInstanceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		executeRequest := ExecuteInstanceScriptRequest{
			Script:   "test script 1",
			Username: "test_user",
			Password: "test_password",
		}

		_, err := s.ExecuteInstanceScript("", executeRequest)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid instance id", err.Error())
	})
}

func TestGetInstancesWithTags(t *testing.T) {
	t.Run("SingleInstanceWithTags", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"i-abcdef12\",\"tags\":[{\"id\":\"tag-abc123\",\"name\":\"production\",\"scope\":\"vpc\"}]}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		instances, err := s.GetInstances(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, instances, 1)
		assert.Equal(t, "i-abcdef12", instances[0].ID)
		assert.Len(t, instances[0].Tags, 1)
		assert.Equal(t, "tag-abc123", instances[0].Tags[0].ID)
		assert.Equal(t, "production", instances[0].Tags[0].Name)
		assert.Equal(t, "vpc", instances[0].Tags[0].Scope)
	})

	t.Run("MultipleInstancesWithDifferentTags", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"i-abcdef12\",\"tags\":[{\"id\":\"tag-abc123\",\"name\":\"production\",\"scope\":\"vpc\"}]},{\"id\":\"i-def456\",\"tags\":[{\"id\":\"tag-ghi789\",\"name\":\"development\",\"scope\":\"instance\"}]}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		instances, err := s.GetInstances(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, instances, 2)

		// First instance
		assert.Equal(t, "i-abcdef12", instances[0].ID)
		assert.Len(t, instances[0].Tags, 1)
		assert.Equal(t, "tag-abc123", instances[0].Tags[0].ID)
		assert.Equal(t, "production", instances[0].Tags[0].Name)
		assert.Equal(t, "vpc", instances[0].Tags[0].Scope)

		// Second instance
		assert.Equal(t, "i-def456", instances[1].ID)
		assert.Len(t, instances[1].Tags, 1)
		assert.Equal(t, "tag-ghi789", instances[1].Tags[0].ID)
		assert.Equal(t, "development", instances[1].Tags[0].Name)
		assert.Equal(t, "instance", instances[1].Tags[0].Scope)
	})
}

func TestGetInstanceWithTags(t *testing.T) {
	t.Run("SingleTag", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"i-abcdef12\",\"tags\":[{\"id\":\"tag-def456\",\"name\":\"development\",\"scope\":\"instance\"}]}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		instance, err := s.GetInstance("i-abcdef12")

		assert.Nil(t, err)
		assert.Equal(t, "i-abcdef12", instance.ID)
		assert.Len(t, instance.Tags, 1)
		assert.Equal(t, "tag-def456", instance.Tags[0].ID)
		assert.Equal(t, "development", instance.Tags[0].Name)
		assert.Equal(t, "instance", instance.Tags[0].Scope)
	})

	t.Run("MultipleTags", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"i-abcdef12\",\"tags\":[{\"id\":\"tag-vpc123\",\"name\":\"production\",\"scope\":\"vpc\"},{\"id\":\"tag-inst456\",\"name\":\"web-server\",\"scope\":\"instance\"}]}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		instance, err := s.GetInstance("i-abcdef12")

		assert.Nil(t, err)
		assert.Equal(t, "i-abcdef12", instance.ID)
		assert.Len(t, instance.Tags, 2)
		assert.Equal(t, "tag-vpc123", instance.Tags[0].ID)
		assert.Equal(t, "production", instance.Tags[0].Name)
		assert.Equal(t, "vpc", instance.Tags[0].Scope)
		assert.Equal(t, "tag-inst456", instance.Tags[1].ID)
		assert.Equal(t, "web-server", instance.Tags[1].Name)
		assert.Equal(t, "instance", instance.Tags[1].Scope)
	})

	t.Run("EmptyTagsArray", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"i-abcdef12\",\"tags\":[]}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		instance, err := s.GetInstance("i-abcdef12")

		assert.Nil(t, err)
		assert.Equal(t, "i-abcdef12", instance.ID)
		assert.Len(t, instance.Tags, 0)
	})

	t.Run("MissingTagsField", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v2/instances/i-abcdef12", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"i-abcdef12\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		instance, err := s.GetInstance("i-abcdef12")

		assert.Nil(t, err)
		assert.Equal(t, "i-abcdef12", instance.ID)
		assert.Len(t, instance.Tags, 0)
	})
}

func TestCreateInstanceWithTags(t *testing.T) {
	t.Run("WithTagIDs", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		createRequest := CreateInstanceRequest{
			Name:        "test instance 1",
			VCPUCores:   2,
			RAMCapacity: 4,
			TagIDs:      []string{"tag-abc123", "tag-def456"},
		}

		c.EXPECT().Post("/ecloud/v2/instances", gomock.Eq(&createRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"i-abcdef12\"}}"))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		id, err := s.CreateInstance(createRequest)

		assert.Nil(t, err)
		assert.Equal(t, "i-abcdef12", id)
	})

	t.Run("WithEmptyTagIDs", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		createRequest := CreateInstanceRequest{
			Name:        "test instance 1",
			VCPUCores:   2,
			RAMCapacity: 4,
			TagIDs:      []string{},
		}

		c.EXPECT().Post("/ecloud/v2/instances", gomock.Eq(&createRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"i-abcdef12\"}}"))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		id, err := s.CreateInstance(createRequest)

		assert.Nil(t, err)
		assert.Equal(t, "i-abcdef12", id)
	})
}

func TestPatchInstanceWithTags(t *testing.T) {
	t.Run("WithTagIDs", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		tagIDs := []string{"tag-abc123", "tag-def456"}
		patchRequest := PatchInstanceRequest{
			Name:        "test instance 1",
			VCPUCores:   2,
			RAMCapacity: 4,
			TagIDs:      &tagIDs,
		}

		c.EXPECT().Patch("/ecloud/v2/instances/i-abcdef12", gomock.Eq(&patchRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
				StatusCode: 202,
			},
		}, nil)

		err := s.PatchInstance("i-abcdef12", patchRequest)

		assert.Nil(t, err)
	})

	t.Run("RemoveAllTags", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		var emptyTagsIDs []string
		patchRequest := PatchInstanceRequest{
			Name:        "test instance 1",
			VCPUCores:   2,
			RAMCapacity: 4,
			TagIDs:      &emptyTagsIDs,
		}

		c.EXPECT().Patch("/ecloud/v2/instances/i-abcdef12", gomock.Eq(&patchRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
				StatusCode: 202,
			},
		}, nil)

		err := s.PatchInstance("i-abcdef12", patchRequest)

		assert.Nil(t, err)
	})
}
