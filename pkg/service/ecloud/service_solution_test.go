package ecloud

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/ptr"
	"github.com/ans-group/sdk-go/test/mocks"
)

func TestGetSolutions(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
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
			c.EXPECT().Get("/ecloud/v1/solutions", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ecloud/v1/solutions", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":456}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
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

		c.EXPECT().Get("/ecloud/v1/solutions", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

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

		c.EXPECT().Get("/ecloud/v1/solutions/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":123}}"))),
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

		c.EXPECT().Get("/ecloud/v1/solutions/123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

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

		c.EXPECT().Get("/ecloud/v1/solutions/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetSolution(123)

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

		patch := PatchSolutionRequest{
			Name: ptr.String("new test name 1"),
		}

		c.EXPECT().Patch("/ecloud/v1/solutions/123", gomock.Eq(&patch)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		_, err := s.PatchSolution(123, patch)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ecloud/v1/solutions/123", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.PatchSolution(123, PatchSolutionRequest{})

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

		_, err := s.PatchSolution(0, PatchSolutionRequest{})

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

		c.EXPECT().Patch("/ecloud/v1/solutions/123", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.PatchSolution(123, PatchSolutionRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &SolutionNotFoundError{}, err)
	})
}

func TestGetSolutionVirtualMachines(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/vms", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		vms, err := s.GetSolutionVirtualMachines(123, connection.APIRequestParameters{})

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
			c.EXPECT().Get("/ecloud/v1/solutions/123/vms", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ecloud/v1/solutions/123/vms", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":456}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		vms, err := s.GetSolutionVirtualMachines(123, connection.APIRequestParameters{})

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

		c.EXPECT().Get("/ecloud/v1/solutions/123/vms", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetSolutionVirtualMachines(123, connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetSolutionSites(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/sites", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		sites, err := s.GetSolutionSites(123, connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, sites, 1)
		assert.Equal(t, 123, sites[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ecloud/v1/solutions/123/sites", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ecloud/v1/solutions/123/sites", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":456}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		sites, err := s.GetSolutionSites(123, connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, sites, 2)
		assert.Equal(t, 123, sites[0].ID)
		assert.Equal(t, 456, sites[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/sites", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetSolutionSites(123, connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetSolutionDatastores(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/datastores", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		datastores, err := s.GetSolutionDatastores(123, connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, datastores, 1)
		assert.Equal(t, 123, datastores[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ecloud/v1/solutions/123/datastores", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ecloud/v1/solutions/123/datastores", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":456}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		datastores, err := s.GetSolutionDatastores(123, connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, datastores, 2)
		assert.Equal(t, 123, datastores[0].ID)
		assert.Equal(t, 456, datastores[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/datastores", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetSolutionDatastores(123, connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetSolutionHosts(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/hosts", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		hosts, err := s.GetSolutionHosts(123, connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, hosts, 1)
		assert.Equal(t, 123, hosts[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ecloud/v1/solutions/123/hosts", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ecloud/v1/solutions/123/hosts", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":456}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		hosts, err := s.GetSolutionHosts(123, connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, hosts, 2)
		assert.Equal(t, 123, hosts[0].ID)
		assert.Equal(t, 456, hosts[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/hosts", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetSolutionHosts(123, connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetSolutionNetworks(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/networks", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		networks, err := s.GetSolutionNetworks(123, connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, networks, 1)
		assert.Equal(t, 123, networks[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ecloud/v1/solutions/123/networks", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ecloud/v1/solutions/123/networks", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":456}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		networks, err := s.GetSolutionNetworks(123, connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, networks, 2)
		assert.Equal(t, 123, networks[0].ID)
		assert.Equal(t, 456, networks[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/networks", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetSolutionNetworks(123, connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetSolutionFirewalls(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/firewalls", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		firewalls, err := s.GetSolutionFirewalls(123, connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, firewalls, 1)
		assert.Equal(t, 123, firewalls[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ecloud/v1/solutions/123/firewalls", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":123}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ecloud/v1/solutions/123/firewalls", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":456}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		firewalls, err := s.GetSolutionFirewalls(123, connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, firewalls, 2)
		assert.Equal(t, 123, firewalls[0].ID)
		assert.Equal(t, 456, firewalls[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/firewalls", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetSolutionFirewalls(123, connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetSolutionTemplates(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/templates", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"name\":\"testtemplate1\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		templates, err := s.GetSolutionTemplates(123, connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, templates, 1)
		assert.Equal(t, "testtemplate1", templates[0].Name)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ecloud/v1/solutions/123/templates", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"name\":\"testtemplate1\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ecloud/v1/solutions/123/templates", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"name\":\"testtemplate2\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		templates, err := s.GetSolutionTemplates(123, connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, templates, 2)
		assert.Equal(t, "testtemplate1", templates[0].Name)
		assert.Equal(t, "testtemplate2", templates[1].Name)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/templates", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetSolutionTemplates(123, connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetSolutionTemplate(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/templates/testname1", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"name\":\"testname1\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		template, err := s.GetSolutionTemplate(123, "testname1")

		assert.Nil(t, err)
		assert.Equal(t, "testname1", template.Name)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/templates/testname1", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetSolutionTemplate(123, "testname1")

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

		_, err := s.GetSolutionTemplate(0, "testname1")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid solution id", err.Error())
	})

	t.Run("InvalidTemplateName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetSolutionTemplate(123, "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid template name", err.Error())
	})

	t.Run("404_ReturnsSolutionNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/templates/testname1", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetSolutionTemplate(123, "testname1")

		assert.NotNil(t, err)
		assert.IsType(t, &TemplateNotFoundError{}, err)
	})
}

func TestRenameSolutionTemplate(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		expectedRequest := RenameTemplateRequest{
			Destination: "testname2",
		}

		c.EXPECT().Post("/ecloud/v1/solutions/123/templates/testname1/move", gomock.Eq(&expectedRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte{})),
				StatusCode: 202,
			},
		}, nil)

		err := s.RenameSolutionTemplate(123, "testname1", expectedRequest)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v1/solutions/123/templates/testname1/move", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.RenameSolutionTemplate(123, "testname1", RenameTemplateRequest{})

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

		err := s.RenameSolutionTemplate(0, "testname1", RenameTemplateRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid solution id", err.Error())
	})

	t.Run("InvalidTemplateName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.RenameSolutionTemplate(123, "", RenameTemplateRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid template name", err.Error())
	})

	t.Run("404_ReturnsSolutionNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v1/solutions/123/templates/testname1/move", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		err := s.RenameSolutionTemplate(123, "testname1", RenameTemplateRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &TemplateNotFoundError{}, err)
	})
}

func TestDeleteSolutionTemplate(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v1/solutions/123/templates/testname1", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte{})),
				StatusCode: 202,
			},
		}, nil)

		err := s.DeleteSolutionTemplate(123, "testname1")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v1/solutions/123/templates/testname1", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteSolutionTemplate(123, "testname1")

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

		err := s.DeleteSolutionTemplate(0, "testname1")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid solution id", err.Error())
	})

	t.Run("InvalidTemplateName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteSolutionTemplate(123, "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid template name", err.Error())
	})

	t.Run("404_ReturnsSolutionNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v1/solutions/123/templates/testname1", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil)

		err := s.DeleteSolutionTemplate(123, "testname1")

		assert.NotNil(t, err)
		assert.IsType(t, &TemplateNotFoundError{}, err)
	})
}

func TestGetSolutionTags(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/tags", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"key\":\"testkey1\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		tags, err := s.GetSolutionTags(123, connection.APIRequestParameters{})

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
			c.EXPECT().Get("/ecloud/v1/solutions/123/tags", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"key\":\"testkey1\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ecloud/v1/solutions/123/tags", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"key\":\"testkey2\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		tags, err := s.GetSolutionTags(123, connection.APIRequestParameters{})

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

		c.EXPECT().Get("/ecloud/v1/solutions/123/tags", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetSolutionTags(123, connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetSolutionTag(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/tags/testkey1", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte("{\"data\":{\"key\":\"testkey1\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		solution, err := s.GetSolutionTag(123, "testkey1")

		assert.Nil(t, err)
		assert.Equal(t, "testkey1", solution.Key)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/tags/testkey1", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetSolutionTag(123, "testkey1")

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

		_, err := s.GetSolutionTag(0, "testkey1")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid solution id", err.Error())
	})

	t.Run("InvalidTagKey_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetSolutionTag(123, "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid tag key", err.Error())
	})

	t.Run("404_ReturnsSolutionNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ecloud/v1/solutions/123/tags/testkey1", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetSolutionTag(123, "testkey1")

		assert.NotNil(t, err)
		assert.IsType(t, &TagV1NotFoundError{}, err)
	})
}

func TestCreateSolutionTag(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v1/solutions/123/tags", gomock.Eq(&CreateTagV1Request{Key: "testkey1", Value: "test value 1"})).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 201,
			},
		}, nil).Times(1)

		err := s.CreateSolutionTag(123, CreateTagV1Request{Key: "testkey1", Value: "test value 1"})

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ecloud/v1/solutions/123/tags", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.CreateSolutionTag(123, CreateTagV1Request{Key: "testkey1", Value: "test value 1"})

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

		err := s.CreateSolutionTag(0, CreateTagV1Request{Key: "testkey1", Value: "test value 1"})

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

		c.EXPECT().Post("/ecloud/v1/solutions/123/tags", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.CreateSolutionTag(123, CreateTagV1Request{Key: "testkey1", Value: "test value 1"})

		assert.NotNil(t, err)
		assert.IsType(t, &SolutionNotFoundError{}, err)
	})
}

func TestPatchSolutionTag(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		patch := PatchTagV1Request{
			Value: "new test value 1",
		}

		c.EXPECT().Patch("/ecloud/v1/solutions/123/tags/testkey1", gomock.Eq(&patch)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PatchSolutionTag(123, "testkey1", patch)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ecloud/v1/solutions/123/tags/testkey1", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchSolutionTag(123, "testkey1", PatchTagV1Request{})

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

		err := s.PatchSolutionTag(0, "testkey1", PatchTagV1Request{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid solution id", err.Error())
	})

	t.Run("InvalidTagKey_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchSolutionTag(123, "", PatchTagV1Request{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid tag key", err.Error())
	})

	t.Run("404_ReturnsSolutionNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ecloud/v1/solutions/123/tags/testkey1", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PatchSolutionTag(123, "testkey1", PatchTagV1Request{})

		assert.NotNil(t, err)
		assert.IsType(t, &TagV1NotFoundError{}, err)
	})
}

func TestDeleteSolutionTag(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v1/solutions/123/tags/testkey1", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.DeleteSolutionTag(123, "testkey1")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v1/solutions/123/tags/testkey1", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteSolutionTag(123, "testkey1")

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

		err := s.DeleteSolutionTag(0, "testkey1")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid solution id", err.Error())
	})

	t.Run("InvalidTagKey_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteSolutionTag(123, "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid tag key", err.Error())
	})

	t.Run("404_ReturnsSolutionNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ecloud/v1/solutions/123/tags/testkey1", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteSolutionTag(123, "testkey1")

		assert.NotNil(t, err)
		assert.IsType(t, &TagV1NotFoundError{}, err)
	})
}
