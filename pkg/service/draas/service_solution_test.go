package draas

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

func TestGetSolutions(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		solutions, err := s.GetSolutions(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, solutions, 1)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", solutions[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

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

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		solution, err := s.GetSolution("00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", solution.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetSolution("00000000-0000-0000-0000-000000000000")

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

		_, err := s.GetSolution("")

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

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetSolution("00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &SolutionNotFoundError{}, err)
	})
}

func TestPatchSolution(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		req := PatchSolutionRequest{}

		c.EXPECT().Patch("/draas/v1/solutions/00000000-0000-0000-0000-000000000000", &req).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PatchSolution("00000000-0000-0000-0000-000000000000", req)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/draas/v1/solutions/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchSolution("00000000-0000-0000-0000-000000000000", PatchSolutionRequest{})

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

		err := s.PatchSolution("", PatchSolutionRequest{})

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

		c.EXPECT().Patch("/draas/v1/solutions/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PatchSolution("00000000-0000-0000-0000-000000000000", PatchSolutionRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &SolutionNotFoundError{}, err)
	})
}

func TestGetSolutionBackupResources(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/backup-resources", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		resources, err := s.GetSolutionBackupResources("00000000-0000-0000-0000-000000000000", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, resources, 1)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", resources[0].ID)
	})

	t.Run("InvalidSolutionID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetSolutionBackupResources("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid solution id", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/backup-resources", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetSolutionBackupResources("00000000-0000-0000-0000-000000000000", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetSolutionBackupService(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/backup-service", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"service\":\"testservice\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		solution, err := s.GetSolutionBackupService("00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
		assert.Equal(t, "testservice", solution.Service)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/backup-service", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetSolutionBackupService("00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidSolutionBackupServiceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetSolutionBackupService("")

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

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/backup-service", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetSolutionBackupService("00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &SolutionNotFoundError{}, err)
	})
}

func TestResetSolutionBackupServiceCredentials(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		req := ResetBackupServiceCredentialsRequest{}

		c.EXPECT().Post("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/backup-service/reset-credentials", &req).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.ResetSolutionBackupServiceCredentials("00000000-0000-0000-0000-000000000000", req)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/backup-service/reset-credentials", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.ResetSolutionBackupServiceCredentials("00000000-0000-0000-0000-000000000000", ResetBackupServiceCredentialsRequest{})

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

		err := s.ResetSolutionBackupServiceCredentials("", ResetBackupServiceCredentialsRequest{})

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

		c.EXPECT().Post("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/backup-service/reset-credentials", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.ResetSolutionBackupServiceCredentials("00000000-0000-0000-0000-000000000000", ResetBackupServiceCredentialsRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &SolutionNotFoundError{}, err)
	})
}

func TestGetSolutionFailoverPlans(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/failover-plans", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		resources, err := s.GetSolutionFailoverPlans("00000000-0000-0000-0000-000000000000", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, resources, 1)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", resources[0].ID)
	})

	t.Run("InvalidSolutionID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetSolutionFailoverPlans("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid solution id", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/failover-plans", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetSolutionFailoverPlans("00000000-0000-0000-0000-000000000000", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetSolutionFailoverPlan(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/failover-plans/00000000-0000-0000-0000-000000000001", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000001\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		plan, err := s.GetSolutionFailoverPlan("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001")

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000001", plan.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/failover-plans/00000000-0000-0000-0000-000000000001", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetSolutionFailoverPlan("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001")

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

		_, err := s.GetSolutionFailoverPlan("", "00000000-0000-0000-0000-000000000001")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid solution id", err.Error())
	})

	t.Run("InvalidSolutionFailoverPlanID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetSolutionFailoverPlan("00000000-0000-0000-0000-000000000000", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid failover plan id", err.Error())
	})

	t.Run("404_ReturnsFailoverPlanNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/failover-plans/00000000-0000-0000-0000-000000000001", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetSolutionFailoverPlan("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001")

		assert.NotNil(t, err)
		assert.IsType(t, &FailoverPlanNotFoundError{}, err)
	})
}

func TestStartSolutionFailoverPlan(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		req := StartFailoverPlanRequest{
			StartDate: "2020-05-19T08:34:45.031Z",
		}

		c.EXPECT().Post("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/failover-plans/00000000-0000-0000-0000-000000000001/start", &req).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000001\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.StartSolutionFailoverPlan("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001", req)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/failover-plans/00000000-0000-0000-0000-000000000001/start", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.StartSolutionFailoverPlan("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001", StartFailoverPlanRequest{})

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

		err := s.StartSolutionFailoverPlan("", "00000000-0000-0000-0000-000000000001", StartFailoverPlanRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid solution id", err.Error())
	})

	t.Run("InvalidSolutionFailoverPlanID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.StartSolutionFailoverPlan("00000000-0000-0000-0000-000000000000", "", StartFailoverPlanRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid failover plan id", err.Error())
	})

	t.Run("404_ReturnsFailoverPlanNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/failover-plans/00000000-0000-0000-0000-000000000001/start", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.StartSolutionFailoverPlan("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001", StartFailoverPlanRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &FailoverPlanNotFoundError{}, err)
	})
}

func TestStopSolutionFailoverPlan(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/failover-plans/00000000-0000-0000-0000-000000000001/stop", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000001\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.StopSolutionFailoverPlan("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/failover-plans/00000000-0000-0000-0000-000000000001/stop", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.StopSolutionFailoverPlan("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001")

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

		err := s.StopSolutionFailoverPlan("", "00000000-0000-0000-0000-000000000001")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid solution id", err.Error())
	})

	t.Run("InvalidSolutionFailoverPlanID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.StopSolutionFailoverPlan("00000000-0000-0000-0000-000000000000", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid failover plan id", err.Error())
	})

	t.Run("404_ReturnsFailoverPlanNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/failover-plans/00000000-0000-0000-0000-000000000001/stop", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.StopSolutionFailoverPlan("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001")

		assert.NotNil(t, err)
		assert.IsType(t, &FailoverPlanNotFoundError{}, err)
	})
}

func TestGetSolutionComputeResources(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/compute-resources", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		resources, err := s.GetSolutionComputeResources("00000000-0000-0000-0000-000000000000", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, resources, 1)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", resources[0].ID)
	})

	t.Run("InvalidSolutionID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetSolutionComputeResources("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid solution id", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/compute-resources", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetSolutionComputeResources("00000000-0000-0000-0000-000000000000", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetSolutionComputeResource(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/compute-resources/00000000-0000-0000-0000-000000000001", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000001\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		plan, err := s.GetSolutionComputeResource("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001")

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000001", plan.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/compute-resources/00000000-0000-0000-0000-000000000001", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetSolutionComputeResource("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001")

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

		_, err := s.GetSolutionComputeResource("", "00000000-0000-0000-0000-000000000001")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid solution id", err.Error())
	})

	t.Run("InvalidSolutionComputeResourceID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetSolutionComputeResource("00000000-0000-0000-0000-000000000000", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid compute resource id", err.Error())
	})

	t.Run("404_ReturnsComputeResourceNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/compute-resources/00000000-0000-0000-0000-000000000001", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetSolutionComputeResource("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001")

		assert.NotNil(t, err)
		assert.IsType(t, &ComputeResourceNotFoundError{}, err)
	})
}

func TestGetSolutionHardwarePlans(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/hardware-plans", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		resources, err := s.GetSolutionHardwarePlans("00000000-0000-0000-0000-000000000000", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, resources, 1)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", resources[0].ID)
	})

	t.Run("InvalidSolutionID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetSolutionHardwarePlans("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid solution id", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/hardware-plans", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetSolutionHardwarePlans("00000000-0000-0000-0000-000000000000", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetSolutionHardwarePlan(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/hardware-plans/00000000-0000-0000-0000-000000000001", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000001\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		plan, err := s.GetSolutionHardwarePlan("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001")

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000001", plan.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/hardware-plans/00000000-0000-0000-0000-000000000001", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetSolutionHardwarePlan("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001")

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

		_, err := s.GetSolutionHardwarePlan("", "00000000-0000-0000-0000-000000000001")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid solution id", err.Error())
	})

	t.Run("InvalidSolutionHardwarePlanID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetSolutionHardwarePlan("00000000-0000-0000-0000-000000000000", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid hardware plan id", err.Error())
	})

	t.Run("404_ReturnsHardwarePlanNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/hardware-plans/00000000-0000-0000-0000-000000000001", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetSolutionHardwarePlan("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001")

		assert.NotNil(t, err)
		assert.IsType(t, &HardwarePlanNotFoundError{}, err)
	})
}

func TestGetSolutionHardwarePlanReplicas(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/hardware-plans/00000000-0000-0000-0000-000000000001/replicas", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000002\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		replicas, err := s.GetSolutionHardwarePlanReplicas("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, replicas, 1)
		assert.Equal(t, "00000000-0000-0000-0000-000000000002", replicas[0].ID)
	})

	t.Run("InvalidSolutionID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetSolutionHardwarePlanReplicas("", "00000000-0000-0000-0000-000000000001", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid solution id", err.Error())
	})

	t.Run("InvalidHardwarePlanID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetSolutionHardwarePlanReplicas("00000000-0000-0000-0000-000000000000", "", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid hardware plan id", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/hardware-plans/00000000-0000-0000-0000-000000000001/replicas", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetSolutionHardwarePlanReplicas("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestUpdateSolutionReplicaIOPS(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		req := UpdateReplicaIOPSRequest{
			IOPSTierID: "someid",
		}

		c.EXPECT().Post("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/replicas/00000000-0000-0000-0000-000000000001/iops", &req).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.UpdateSolutionReplicaIOPS("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001", req)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/replicas/00000000-0000-0000-0000-000000000001/iops", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.UpdateSolutionReplicaIOPS("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001", UpdateReplicaIOPSRequest{})

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

		err := s.UpdateSolutionReplicaIOPS("", "00000000-0000-0000-0000-000000000001", UpdateReplicaIOPSRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid solution id", err.Error())
	})

	t.Run("InvalidReplicaID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.UpdateSolutionReplicaIOPS("00000000-0000-0000-0000-000000000000", "", UpdateReplicaIOPSRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid replica id", err.Error())
	})

	t.Run("404_ReturnsReplicaNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/draas/v1/solutions/00000000-0000-0000-0000-000000000000/replicas/00000000-0000-0000-0000-000000000001/iops", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.UpdateSolutionReplicaIOPS("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001", UpdateReplicaIOPSRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &ReplicaNotFoundError{}, err)
	})
}
