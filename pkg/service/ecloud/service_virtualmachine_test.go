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
	"github.com/ukfast/sdk-go/pkg/ptr"
	"github.com/ukfast/sdk-go/test/mocks"
)

func TestGetVirtualMachines(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/vms", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		vms, err := s.GetVirtualMachines(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, vms, 1)
		assert.Equal(t, 123, vms[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ecloud/v1/vms", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ecloud/v1/vms", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":456}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		vms, err := s.GetVirtualMachines(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, vms, 2)
		assert.Equal(t, 123, vms[0].ID)
		assert.Equal(t, 456, vms[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/vms", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetVirtualMachines(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetVirtualMachine(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/vms/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":123}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		vm, err := s.GetVirtualMachine(123)

		assert.Nil(t, err)
		assert.Equal(t, 123, vm.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/vms/123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetVirtualMachine(123)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVirtualMachineID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetVirtualMachine(0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid virtual machine id", err.Error())
	})

	t.Run("404_ReturnsVirtualMachineNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/vms/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetVirtualMachine(123)

		assert.NotNil(t, err)
		assert.IsType(t, &VirtualMachineNotFoundError{}, err)
	})
}

func TestDeleteVirtualMachine(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v1/vms/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		err := s.DeleteVirtualMachine(123)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v1/vms/123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteVirtualMachine(123)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVirtualMachineID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteVirtualMachine(0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid virtual machine id", err.Error())
	})

	t.Run("404_ReturnsVirtualMachineNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v1/vms/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteVirtualMachine(123)

		assert.NotNil(t, err)
		assert.IsType(t, &VirtualMachineNotFoundError{}, err)
	})
}

func TestCreateVirtualMachine(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		createRequest := CreateVirtualMachineRequest{
			Name: "test vm 1",
			CPU:  2,
			RAM:  4,
		}

		c.EXPECT().Post("/ecloud/v1/vms", gomock.Eq(&createRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":123}}"))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		id, err := s.CreateVirtualMachine(createRequest)

		assert.Nil(t, err)
		assert.Equal(t, 123, id)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v1/vms", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateVirtualMachine(CreateVirtualMachineRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestPatchVirtualMachine(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		patch := PatchVirtualMachineRequest{
			Name: ptr.String("new test vm name 1"),
		}

		c.EXPECT().Patch("/ecloud/v1/vms/123", gomock.Eq(&patch)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		err := s.PatchVirtualMachine(123, patch)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ecloud/v1/vms/123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchVirtualMachine(123, PatchVirtualMachineRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVirtualMachineID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchVirtualMachine(0, PatchVirtualMachineRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid virtual machine id", err.Error())
	})

	t.Run("404_ReturnsVirtualMachineNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ecloud/v1/vms/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PatchVirtualMachine(123, PatchVirtualMachineRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &VirtualMachineNotFoundError{}, err)
	})
}

func TestCloneVirtualMachine(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		cloneRequest := CloneVirtualMachineRequest{
			Name: "test vm 1",
		}

		c.EXPECT().Post("/ecloud/v1/vms/123/clone", gomock.Eq(&cloneRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":456}}"))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		id, err := s.CloneVirtualMachine(123, cloneRequest)

		assert.Nil(t, err)
		assert.Equal(t, 456, id)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v1/vms/123/clone", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CloneVirtualMachine(123, CloneVirtualMachineRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVirtualMachineID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.CloneVirtualMachine(0, CloneVirtualMachineRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid virtual machine id", err.Error())
	})

	t.Run("404_ReturnsVirtualMachineNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v1/vms/123/clone", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.CloneVirtualMachine(123, CloneVirtualMachineRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &VirtualMachineNotFoundError{}, err)
	})
}

func TestPowerOnVirtualMachine(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v1/vms/123/power-on", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.PowerOnVirtualMachine(123)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v1/vms/123/power-on", nil).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PowerOnVirtualMachine(123)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVirtualMachineID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PowerOnVirtualMachine(0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid virtual machine id", err.Error())
	})

	t.Run("404_ReturnsVirtualMachineNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v1/vms/123/power-on", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PowerOnVirtualMachine(123)

		assert.NotNil(t, err)
		assert.IsType(t, &VirtualMachineNotFoundError{}, err)
	})
}

func TestPowerOffVirtualMachine(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v1/vms/123/power-off", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.PowerOffVirtualMachine(123)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v1/vms/123/power-off", nil).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PowerOffVirtualMachine(123)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVirtualMachineID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PowerOffVirtualMachine(0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid virtual machine id", err.Error())
	})

	t.Run("404_ReturnsVirtualMachineNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v1/vms/123/power-off", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PowerOffVirtualMachine(123)

		assert.NotNil(t, err)
		assert.IsType(t, &VirtualMachineNotFoundError{}, err)
	})
}

func TestPowerResetVirtualMachine(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v1/vms/123/power-reset", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PowerResetVirtualMachine(123)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v1/vms/123/power-reset", nil).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PowerResetVirtualMachine(123)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVirtualMachineID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PowerResetVirtualMachine(0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid virtual machine id", err.Error())
	})

	t.Run("404_ReturnsVirtualMachineNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v1/vms/123/power-reset", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PowerResetVirtualMachine(123)

		assert.NotNil(t, err)
		assert.IsType(t, &VirtualMachineNotFoundError{}, err)
	})
}

func TestPowerShutdownVirtualMachine(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v1/vms/123/power-shutdown", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PowerShutdownVirtualMachine(123)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v1/vms/123/power-shutdown", nil).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PowerShutdownVirtualMachine(123)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVirtualMachineID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PowerShutdownVirtualMachine(0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid virtual machine id", err.Error())
	})

	t.Run("404_ReturnsVirtualMachineNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v1/vms/123/power-shutdown", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PowerShutdownVirtualMachine(123)

		assert.NotNil(t, err)
		assert.IsType(t, &VirtualMachineNotFoundError{}, err)
	})
}

func TestPowerRestartVirtualMachine(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v1/vms/123/power-restart", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PowerRestartVirtualMachine(123)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v1/vms/123/power-restart", nil).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PowerRestartVirtualMachine(123)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVirtualMachineID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PowerRestartVirtualMachine(0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid virtual machine id", err.Error())
	})

	t.Run("404_ReturnsVirtualMachineNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v1/vms/123/power-restart", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PowerRestartVirtualMachine(123)

		assert.NotNil(t, err)
		assert.IsType(t, &VirtualMachineNotFoundError{}, err)
	})
}

func TestCreateVirtualMachineTemplate(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		expectedRequest := CreateVirtualMachineTemplateRequest{
			TemplateName: "test template 1",
		}

		c.EXPECT().Post("/ecloud/v1/vms/123/clone-to-template", gomock.Eq(&expectedRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 202,
			},
		}, nil).Times(1)

		err := s.CreateVirtualMachineTemplate(123, expectedRequest)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v1/vms/123/clone-to-template", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.CreateVirtualMachineTemplate(123, CreateVirtualMachineTemplateRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVirtualMachineID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.CreateVirtualMachineTemplate(0, CreateVirtualMachineTemplateRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid virtual machine id", err.Error())
	})

	t.Run("404_ReturnsVirtualMachineNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v1/vms/123/clone-to-template", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.CreateVirtualMachineTemplate(123, CreateVirtualMachineTemplateRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &VirtualMachineNotFoundError{}, err)
	})
}

func TestGetVirtualMachineTags(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/vms/123/tags", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"key\":\"testkey1\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		tags, err := s.GetVirtualMachineTags(123, connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, tags, 1)
		assert.Equal(t, "testkey1", tags[0].Key)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ecloud/v1/vms/123/tags", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"key\":\"testkey1\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ecloud/v1/vms/123/tags", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"key\":\"testkey2\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		tags, err := s.GetVirtualMachineTags(123, connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, tags, 2)
		assert.Equal(t, "testkey1", tags[0].Key)
		assert.Equal(t, "testkey2", tags[1].Key)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/vms/123/tags", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetVirtualMachineTags(123, connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetVirtualMachineTag(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/vms/123/tags/testkey1", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"key\":\"testkey1\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		vm, err := s.GetVirtualMachineTag(123, "testkey1")

		assert.Nil(t, err)
		assert.Equal(t, "testkey1", vm.Key)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/vms/123/tags/testkey1", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetVirtualMachineTag(123, "testkey1")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVirtualMachineID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetVirtualMachineTag(0, "testkey1")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid virtual machine id", err.Error())
	})

	t.Run("InvalidTagKey_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetVirtualMachineTag(123, "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid tag key", err.Error())
	})

	t.Run("404_ReturnsVirtualMachineNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/vms/123/tags/testkey1", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetVirtualMachineTag(123, "testkey1")

		assert.NotNil(t, err)
		assert.IsType(t, &TagNotFoundError{}, err)
	})
}

func TestCreateVirtualMachineTag(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v1/vms/123/tags", gomock.Eq(&CreateTagRequest{Key: "testkey1", Value: "test value 1"})).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 201,
			},
		}, nil).Times(1)

		err := s.CreateVirtualMachineTag(123, CreateTagRequest{Key: "testkey1", Value: "test value 1"})

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v1/vms/123/tags", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.CreateVirtualMachineTag(123, CreateTagRequest{Key: "testkey1", Value: "test value 1"})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVirtualMachineID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.CreateVirtualMachineTag(0, CreateTagRequest{Key: "testkey1", Value: "test value 1"})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid virtual machine id", err.Error())
	})

	t.Run("404_ReturnsVirtualMachineNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v1/vms/123/tags", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.CreateVirtualMachineTag(123, CreateTagRequest{Key: "testkey1", Value: "test value 1"})

		assert.NotNil(t, err)
		assert.IsType(t, &VirtualMachineNotFoundError{}, err)
	})
}

func TestPatchVirtualMachineTag(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		patch := PatchTagRequest{
			Value: "new test value 1",
		}

		c.EXPECT().Patch("/ecloud/v1/vms/123/tags/testkey1", gomock.Eq(&patch)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PatchVirtualMachineTag(123, "testkey1", patch)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ecloud/v1/vms/123/tags/testkey1", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchVirtualMachineTag(123, "testkey1", PatchTagRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVirtualMachineID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchVirtualMachineTag(0, "testkey1", PatchTagRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid virtual machine id", err.Error())
	})

	t.Run("InvalidTagKey_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchVirtualMachineTag(123, "", PatchTagRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid tag key", err.Error())
	})

	t.Run("404_ReturnsVirtualMachineNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ecloud/v1/vms/123/tags/testkey1", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PatchVirtualMachineTag(123, "testkey1", PatchTagRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &TagNotFoundError{}, err)
	})
}

func TestDeleteVirtualMachineTag(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v1/vms/123/tags/testkey1", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.DeleteVirtualMachineTag(123, "testkey1")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v1/vms/123/tags/testkey1", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteVirtualMachineTag(123, "testkey1")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidVirtualMachineID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteVirtualMachineTag(0, "testkey1")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid virtual machine id", err.Error())
	})

	t.Run("InvalidTagKey_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteVirtualMachineTag(123, "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid tag key", err.Error())
	})

	t.Run("404_ReturnsVirtualMachineNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v1/vms/123/tags/testkey1", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteVirtualMachineTag(123, "testkey1")

		assert.NotNil(t, err)
		assert.IsType(t, &TagNotFoundError{}, err)
	})
}

func TestCreateVirtualMachineConsoleSession(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v1/vms/123/console-session", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"url\":\"ukfast.co.uk\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		session, err := s.CreateVirtualMachineConsoleSession(123)

		assert.Nil(t, err)
		assert.Equal(t, "ukfast.co.uk", session.URL)
	})

	t.Run("InvalidVirtualMachineID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.CreateVirtualMachineConsoleSession(0)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid virtual machine id", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v1/vms/123/console-session", nil).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateVirtualMachineConsoleSession(123)

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsVirtualMachineNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Put("/ecloud/v1/vms/123/console-session", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.CreateVirtualMachineConsoleSession(123)

		assert.NotNil(t, err)
		assert.IsType(t, &VirtualMachineNotFoundError{}, err)
	})
}
