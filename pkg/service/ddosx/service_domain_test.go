package ddosx

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ukfast/sdk-go/test"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/ukfast/sdk-go/pkg/connection"
	"github.com/ukfast/sdk-go/pkg/ptr"
	"github.com/ukfast/sdk-go/test/mocks"
)

func TestGetDomains(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"name\":\"testdomain1.co.uk\"}],\"meta\":{\"pagination\":{\"total_pages\":1}}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		domains, err := s.GetDomains(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, domains, 1)
		assert.Equal(t, "testdomain1.co.uk", domains[0].Name)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ddosx/v1/domains", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"name\":\"testdomain1.co.uk\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ddosx/v1/domains", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"name\":\"testdomain2.co.uk\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		domains, err := s.GetDomains(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, domains, 2)
		assert.Equal(t, "testdomain1.co.uk", domains[0].Name)
		assert.Equal(t, "testdomain2.co.uk", domains[1].Name)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetDomains(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetDomainsPaginated(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"name\":\"testdomain1.co.uk\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		domains, err := s.GetDomainsPaginated(connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Equal(t, "testdomain1.co.uk", domains[0].Name)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainsPaginated(connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetDomain(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"name\":\"testdomain1.co.uk\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		domain, err := s.GetDomain("testdomain1.co.uk")

		assert.Nil(t, err)
		assert.Equal(t, "testdomain1.co.uk", domain.Name)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomain("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomain("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomain("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.IsType(t, &DomainNotFoundError{}, err)
	})
}

func TestCreateDomain(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		createRequest := CreateDomainRequest{
			Name: "testdomain1.co.uk",
		}

		c.EXPECT().Post("/ddosx/v1/domains", gomock.Eq(&createRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.CreateDomain(createRequest)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.CreateDomain(CreateDomainRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestDeployDomain(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/deploy", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"name\":\"testdomain1.co.uk\"}}"))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.DeployDomain("testdomain1.co.uk")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/deploy", nil).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeployDomain("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeployDomain("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/deploy", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeployDomain("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.IsType(t, &DomainNotFoundError{}, err)
	})
}

func TestGetDomainRecords(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/records", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		records, err := s.GetDomainRecords("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, records, 1)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", records[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/records", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/records", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000001\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		records, err := s.GetDomainRecords("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, records, 2)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", records[0].ID)
		assert.Equal(t, "00000000-0000-0000-0000-000000000001", records[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/records", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetDomainRecords("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetDomainRecordsPaginated(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/records", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		records, err := s.GetDomainRecordsPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", records[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/records", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainRecordsPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainRecords("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/records", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainRecordsPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainNotFoundError{}, err)
	})
}

func TestGetDomainRecord(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/records/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		record, err := s.GetDomainRecord("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", record.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/records/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainRecord("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainRecord("", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRecordID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainRecord("testdomain1.co.uk", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid record ID", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/records/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainRecord("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &DomainRecordNotFoundError{}, err)
	})
}

func TestCreateDomainRecord(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		createRequest := CreateRecordRequest{
			Name: "testrecord.testdomain1.co.uk",
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/records", gomock.Eq(&createRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		id, err := s.CreateDomainRecord("testdomain1.co.uk", createRequest)

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", id)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.CreateDomainRecord("", CreateRecordRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/records", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateDomainRecord("testdomain1.co.uk", CreateRecordRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/records", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.CreateDomainRecord("testdomain1.co.uk", CreateRecordRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainNotFoundError{}, err)
	})
}

func TestPatchDomainRecord(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		patchRequest := PatchRecordRequest{
			Content: "1.2.3.4",
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/records/00000000-0000-0000-0000-000000000000", gomock.Eq(&patchRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PatchDomainRecord("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", patchRequest)

		assert.Nil(t, err)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainRecord("", "00000000-0000-0000-0000-000000000000", PatchRecordRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidDomainRecordID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainRecord("testdomain1.co.uk", "", PatchRecordRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid record ID", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/records/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchDomainRecord("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchRecordRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/records/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PatchDomainRecord("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchRecordRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainRecordNotFoundError{}, err)
	})
}

func TestDeleteDomainRecord(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/records/00000000-0000-0000-0000-000000000000", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.DeleteDomainRecord("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteDomainRecord("", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRecord_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteDomainRecord("testdomain1.co.uk", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid record ID", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/records/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteDomainRecord("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/records/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteDomainRecord("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &DomainRecordNotFoundError{}, err)
	})
}

func TestGetDomainProperties(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/properties", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		properties, err := s.GetDomainProperties("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, properties, 1)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", properties[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/properties", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/properties", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000001\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		properties, err := s.GetDomainProperties("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, properties, 2)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", properties[0].ID)
		assert.Equal(t, "00000000-0000-0000-0000-000000000001", properties[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/properties", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetDomainProperties("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetDomainPropertiesPaginated(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/properties", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		properties, err := s.GetDomainPropertiesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", properties[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/properties", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainPropertiesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainProperties("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/properties", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainPropertiesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainNotFoundError{}, err)
	})
}

func TestGetDomainProperty(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/properties/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		property, err := s.GetDomainProperty("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", property.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/properties/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainProperty("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainProperty("", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidPropertyID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainProperty("testdomain1.co.uk", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid property ID", err.Error())
	})

	t.Run("404_ReturnsDomainPropertyNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/properties/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainProperty("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &DomainPropertyNotFoundError{}, err)
	})
}

func TestPatchDomainProperty(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		patchRequest := PatchDomainPropertyRequest{
			Value: "testvalue",
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/properties/00000000-0000-0000-0000-000000000000", gomock.Eq(&patchRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PatchDomainProperty("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", patchRequest)

		assert.Nil(t, err)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainProperty("", "00000000-0000-0000-0000-000000000000", PatchDomainPropertyRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidDomainPropertyID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainProperty("testdomain1.co.uk", "", PatchDomainPropertyRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid property ID", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/properties/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchDomainProperty("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchDomainPropertyRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/properties/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PatchDomainProperty("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchDomainPropertyRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainPropertyNotFoundError{}, err)
	})
}

func TestGetDomainWAF(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"mode\":\"On\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		property, err := s.GetDomainWAF("testdomain1.co.uk")

		assert.Nil(t, err)
		assert.Equal(t, WAFModeOn, property.Mode)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainWAF("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainWAF("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainWAFNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainWAF("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.IsType(t, &DomainWAFNotFoundError{}, err)
	})
}

func TestCreateDomainWAF(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		expectedRequest := CreateWAFRequest{
			Mode:          WAFModeOn,
			ParanoiaLevel: WAFParanoiaLevelMedium,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/waf", gomock.Eq(&expectedRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 201,
			},
		}, nil).Times(1)

		err := s.CreateDomainWAF("testdomain1.co.uk", expectedRequest)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/waf", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.CreateDomainWAF("testdomain1.co.uk", CreateWAFRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.CreateDomainWAF("", CreateWAFRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/waf", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.CreateDomainWAF("testdomain1.co.uk", CreateWAFRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainNotFoundError{}, err)
	})
}

func TestPatchDomainWAF(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		expectedRequest := PatchWAFRequest{
			Mode:          WAFModeOn,
			ParanoiaLevel: WAFParanoiaLevelMedium,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/waf", gomock.Eq(&expectedRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PatchDomainWAF("testdomain1.co.uk", expectedRequest)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/waf", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchDomainWAF("testdomain1.co.uk", PatchWAFRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainWAF("", PatchWAFRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/waf", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PatchDomainWAF("testdomain1.co.uk", PatchWAFRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainWAFNotFoundError{}, err)
	})
}

func TestDeleteDomainWAF(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/waf", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.DeleteDomainWAF("testdomain1.co.uk")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/waf", nil).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteDomainWAF("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteDomainWAF("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/waf", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteDomainWAF("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.IsType(t, &DomainWAFNotFoundError{}, err)
	})
}

func TestGetDomainWAFRuleSets(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rulesets", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rulesets, err := s.GetDomainWAFRuleSets("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, rulesets, 1)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rulesets[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rulesets", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rulesets", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000001\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		rulesets, err := s.GetDomainWAFRuleSets("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, rulesets, 2)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rulesets[0].ID)
		assert.Equal(t, "00000000-0000-0000-0000-000000000001", rulesets[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rulesets", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetDomainWAFRuleSets("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetDomainWAFRuleSetsPaginated(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rulesets", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rulesets, err := s.GetDomainWAFRuleSetsPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rulesets[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rulesets", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainWAFRuleSetsPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainWAFRuleSets("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rulesets", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainWAFRuleSetsPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainWAFNotFoundError{}, err)
	})
}

func TestGetDomainWAFRuleSet(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rulesets/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		ruleset, err := s.GetDomainWAFRuleSet("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", ruleset.ID)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainWAFRuleSet("", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRuleSet_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainWAFRuleSet("testdomain1.co.uk", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule set ID", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rulesets/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainWAFRuleSet("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rulesets/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainWAFRuleSet("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &WAFRuleSetNotFoundError{}, err)
	})
}

func TestPatchDomainWAFRuleSet(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/waf/rulesets/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PatchDomainWAFRuleSet("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchWAFRuleSetRequest{})

		assert.Nil(t, err)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainWAFRuleSet("", "00000000-0000-0000-0000-000000000000", PatchWAFRuleSetRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRuleSet_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainWAFRuleSet("testdomain1.co.uk", "", PatchWAFRuleSetRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule set ID", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/waf/rulesets/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchDomainWAFRuleSet("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchWAFRuleSetRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/waf/rulesets/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PatchDomainWAFRuleSet("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchWAFRuleSetRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &WAFRuleSetNotFoundError{}, err)
	})
}

func TestGetDomainWAFRules(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rules", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rules, err := s.GetDomainWAFRules("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, rules, 1)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rules[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rules", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rules", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000001\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		rules, err := s.GetDomainWAFRules("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, rules, 2)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rules[0].ID)
		assert.Equal(t, "00000000-0000-0000-0000-000000000001", rules[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rules", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetDomainWAFRules("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetDomainWAFRulesPaginated(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rules", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rules, err := s.GetDomainWAFRulesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rules[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rules", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainWAFRulesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainWAFRules("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rules", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainWAFRulesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainWAFNotFoundError{}, err)
	})
}

func TestGetDomainWAFRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rule, err := s.GetDomainWAFRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rule.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainWAFRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainWAFRule("", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRuleID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainWAFRule("testdomain1.co.uk", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule ID", err.Error())
	})

	t.Run("404_WAFRuleNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainWAFRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &WAFRuleNotFoundError{}, err)
	})
}

func TestCreateDomainWAFRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		createRequest := CreateWAFRuleRequest{
			URI: "test.html",
			IP:  "1.2.3.4",
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/waf/rules", gomock.Eq(&createRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 201,
			},
		}, nil).Times(1)

		id, err := s.CreateDomainWAFRule("testdomain1.co.uk", createRequest)

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", id)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.CreateDomainWAFRule("", CreateWAFRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/waf/rules", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateDomainWAFRule("testdomain1.co.uk", CreateWAFRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/waf/rules", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.CreateDomainWAFRule("testdomain1.co.uk", CreateWAFRuleRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainNotFoundError{}, err)
	})
}

func TestPatchDomainWAFRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/waf/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PatchDomainWAFRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchWAFRuleRequest{})

		assert.Nil(t, err)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainWAFRule("", "00000000-0000-0000-0000-000000000000", PatchWAFRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRule_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainWAFRule("testdomain1.co.uk", "", PatchWAFRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule ID", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/waf/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchDomainWAFRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchWAFRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/waf/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PatchDomainWAFRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchWAFRuleRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &WAFRuleNotFoundError{}, err)
	})
}

func TestDeleteDomainWAFRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/waf/rules/00000000-0000-0000-0000-000000000000", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.DeleteDomainWAFRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteDomainWAFRule("", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRule_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteDomainWAFRule("testdomain1.co.uk", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule ID", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/waf/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteDomainWAFRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/waf/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteDomainWAFRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &WAFRuleNotFoundError{}, err)
	})
}

func TestGetDomainWAFAdvancedRules(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/advanced-rules", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rules, err := s.GetDomainWAFAdvancedRules("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, rules, 1)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rules[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/advanced-rules", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/advanced-rules", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000001\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		rules, err := s.GetDomainWAFAdvancedRules("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, rules, 2)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rules[0].ID)
		assert.Equal(t, "00000000-0000-0000-0000-000000000001", rules[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/advanced-rules", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetDomainWAFAdvancedRules("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetDomainWAFAdvancedRulesPaginated(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/advanced-rules", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rules, err := s.GetDomainWAFAdvancedRulesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rules[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/advanced-rules", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainWAFAdvancedRulesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainWAFAdvancedRules("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/advanced-rules", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainWAFAdvancedRulesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainWAFNotFoundError{}, err)
	})
}

func TestGetDomainWAFAdvancedRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/advanced-rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rule, err := s.GetDomainWAFAdvancedRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rule.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/advanced-rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainWAFAdvancedRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainWAFAdvancedRule("", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRuleID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainWAFAdvancedRule("testdomain1.co.uk", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule ID", err.Error())
	})

	t.Run("404_WAFAdvancedRuleNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/waf/advanced-rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainWAFAdvancedRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &WAFAdvancedRuleNotFoundError{}, err)
	})
}

func TestCreateDomainWAFAdvancedRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		createRequest := CreateWAFAdvancedRuleRequest{
			Section:  "REQUEST_URI",
			Modifier: WAFAdvancedRuleModifierContains,
			Phrase:   "testphrase",
			IP:       "1.2.3.4",
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/waf/advanced-rules", gomock.Eq(&createRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 201,
			},
		}, nil).Times(1)

		id, err := s.CreateDomainWAFAdvancedRule("testdomain1.co.uk", createRequest)

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", id)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.CreateDomainWAFAdvancedRule("", CreateWAFAdvancedRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/waf/advanced-rules", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateDomainWAFAdvancedRule("testdomain1.co.uk", CreateWAFAdvancedRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/waf/advanced-rules", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.CreateDomainWAFAdvancedRule("testdomain1.co.uk", CreateWAFAdvancedRuleRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainNotFoundError{}, err)
	})
}

func TestPatchDomainWAFAdvancedRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/waf/advanced-rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PatchDomainWAFAdvancedRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchWAFAdvancedRuleRequest{})

		assert.Nil(t, err)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainWAFAdvancedRule("", "00000000-0000-0000-0000-000000000000", PatchWAFAdvancedRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRule_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainWAFAdvancedRule("testdomain1.co.uk", "", PatchWAFAdvancedRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule ID", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/waf/advanced-rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchDomainWAFAdvancedRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchWAFAdvancedRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/waf/advanced-rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PatchDomainWAFAdvancedRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchWAFAdvancedRuleRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &WAFAdvancedRuleNotFoundError{}, err)
	})
}

func TestDeleteDomainWAFAdvancedRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/waf/advanced-rules/00000000-0000-0000-0000-000000000000", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.DeleteDomainWAFAdvancedRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteDomainWAFAdvancedRule("", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRule_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteDomainWAFAdvancedRule("testdomain1.co.uk", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule ID", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/waf/advanced-rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteDomainWAFAdvancedRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/waf/advanced-rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteDomainWAFAdvancedRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &WAFAdvancedRuleNotFoundError{}, err)
	})
}

func TestGetDomainACLGeoIPRules(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rules, err := s.GetDomainACLGeoIPRules("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, rules, 1)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rules[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000001\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		rules, err := s.GetDomainACLGeoIPRules("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, rules, 2)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rules[0].ID)
		assert.Equal(t, "00000000-0000-0000-0000-000000000001", rules[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetDomainACLGeoIPRules("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetDomainACLGeoIPRulesPaginated(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rules, err := s.GetDomainACLGeoIPRulesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rules[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainACLGeoIPRulesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainACLGeoIPRules("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainACLGeoIPRulesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainNotFoundError{}, err)
	})
}

func TestGetDomainACLGeoIPRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rule, err := s.GetDomainACLGeoIPRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rule.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainACLGeoIPRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainACLGeoIPRule("", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRuleID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainACLGeoIPRule("testdomain1.co.uk", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule ID", err.Error())
	})

	t.Run("404_ACLGeoIPRuleNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainACLGeoIPRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &ACLGeoIPRuleNotFoundError{}, err)
	})
}

func TestCreateDomainACLGeoIPRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		createRequest := CreateACLGeoIPRuleRequest{
			Code: "GB",
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips", gomock.Eq(&createRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 201,
			},
		}, nil).Times(1)

		id, err := s.CreateDomainACLGeoIPRule("testdomain1.co.uk", createRequest)

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", id)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.CreateDomainACLGeoIPRule("", CreateACLGeoIPRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateDomainACLGeoIPRule("testdomain1.co.uk", CreateACLGeoIPRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.CreateDomainACLGeoIPRule("testdomain1.co.uk", CreateACLGeoIPRuleRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainNotFoundError{}, err)
	})
}

func TestPatchDomainACLGeoIPRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		patchRequest := PatchACLGeoIPRuleRequest{
			Code: "GB",
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips/00000000-0000-0000-0000-000000000000", gomock.Eq(&patchRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PatchDomainACLGeoIPRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", patchRequest)

		assert.Nil(t, err)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainACLGeoIPRule("", "00000000-0000-0000-0000-000000000000", PatchACLGeoIPRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRule_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainACLGeoIPRule("testdomain1.co.uk", "", PatchACLGeoIPRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule ID", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchDomainACLGeoIPRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchACLGeoIPRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PatchDomainACLGeoIPRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchACLGeoIPRuleRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &ACLGeoIPRuleNotFoundError{}, err)
	})
}

func TestDeleteDomainACLGeoIPRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips/00000000-0000-0000-0000-000000000000", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.DeleteDomainACLGeoIPRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteDomainACLGeoIPRule("", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRule_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteDomainACLGeoIPRule("testdomain1.co.uk", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule ID", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteDomainACLGeoIPRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteDomainACLGeoIPRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &ACLGeoIPRuleNotFoundError{}, err)
	})
}

func TestGetDomainACLGeoIPRulesMode(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips/mode", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"mode\":\"Whitelist\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		mode, err := s.GetDomainACLGeoIPRulesMode("testdomain1.co.uk")

		assert.Nil(t, err)
		assert.Equal(t, ACLGeoIPRulesModeWhitelist, mode)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips/mode", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainACLGeoIPRulesMode("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainACLGeoIPRulesMode("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips/mode", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainACLGeoIPRulesMode("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.IsType(t, &DomainNotFoundError{}, err)
	})
}

func TestPatchDomainACLGeoIPRulesMode(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		patchRequest := PatchACLGeoIPRulesModeRequest{
			Mode: ACLGeoIPRulesModeWhitelist,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips/mode", gomock.Eq(&patchRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PatchDomainACLGeoIPRulesMode("testdomain1.co.uk", patchRequest)

		assert.Nil(t, err)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainACLGeoIPRulesMode("", PatchACLGeoIPRulesModeRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips/mode", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchDomainACLGeoIPRulesMode("testdomain1.co.uk", PatchACLGeoIPRulesModeRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/acls/geo-ips/mode", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PatchDomainACLGeoIPRulesMode("testdomain1.co.uk", PatchACLGeoIPRulesModeRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainNotFoundError{}, err)
	})
}

func TestGetDomainACLIPRules(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/ips", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rules, err := s.GetDomainACLIPRules("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, rules, 1)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rules[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/ips", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/ips", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000001\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		rules, err := s.GetDomainACLIPRules("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, rules, 2)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rules[0].ID)
		assert.Equal(t, "00000000-0000-0000-0000-000000000001", rules[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/ips", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetDomainACLIPRules("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetDomainACLIPRulesPaginated(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/ips", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rules, err := s.GetDomainACLIPRulesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rules[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/ips", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainACLIPRulesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainACLIPRulesPaginated("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/ips", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainACLIPRulesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainWAFNotFoundError{}, err)
	})
}

func TestGetDomainACLIPRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/ips/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rule, err := s.GetDomainACLIPRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rule.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/ips/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainACLIPRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainACLIPRule("", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRuleID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainACLIPRule("testdomain1.co.uk", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule ID", err.Error())
	})

	t.Run("404_ACLIPRuleNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/acls/ips/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainACLIPRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &ACLIPRuleNotFoundError{}, err)
	})
}

func TestCreateDomainACLIPRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		createRequest := CreateACLIPRuleRequest{
			IP: "1.2.3.4",
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/acls/ips", gomock.Eq(&createRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 201,
			},
		}, nil).Times(1)

		id, err := s.CreateDomainACLIPRule("testdomain1.co.uk", createRequest)

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", id)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.CreateDomainACLIPRule("", CreateACLIPRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/acls/ips", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateDomainACLIPRule("testdomain1.co.uk", CreateACLIPRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/acls/ips", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.CreateDomainACLIPRule("testdomain1.co.uk", CreateACLIPRuleRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainNotFoundError{}, err)
	})
}

func TestPatchDomainACLIPRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		patchRequest := PatchACLIPRuleRequest{
			IP: "1.2.3.4",
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/acls/ips/00000000-0000-0000-0000-000000000000", gomock.Eq(&patchRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PatchDomainACLIPRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", patchRequest)

		assert.Nil(t, err)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainACLIPRule("", "00000000-0000-0000-0000-000000000000", PatchACLIPRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRule_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainACLIPRule("testdomain1.co.uk", "", PatchACLIPRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule ID", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/acls/ips/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchDomainACLIPRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchACLIPRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/acls/ips/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PatchDomainACLIPRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchACLIPRuleRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &ACLIPRuleNotFoundError{}, err)
	})
}

func TestDeleteDomainACLIPRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/acls/ips/00000000-0000-0000-0000-000000000000", nil).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.DeleteDomainACLIPRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteDomainACLIPRule("", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRule_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteDomainACLIPRule("testdomain1.co.uk", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule ID", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/acls/ips/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteDomainACLIPRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/acls/ips/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteDomainACLIPRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &ACLIPRuleNotFoundError{}, err)
	})
}

func TestDownloadDomainVerificationFile(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		response := &http.Response{
			Body:       ioutil.NopCloser(bytes.NewReader([]byte("test content"))),
			StatusCode: 200,
		}
		response.Header = http.Header{}
		response.Header.Set("Content-Disposition", "attachment; filename=\"testfile1.txt\"")

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/verify/file-upload", gomock.Any()).Return(&connection.APIResponse{
			Response: response,
		}, nil)

		content, filename, err := s.DownloadDomainVerificationFile("testdomain1.co.uk")

		assert.Nil(t, err)
		assert.Equal(t, "test content", content)
		assert.Equal(t, "testfile1.txt", filename)
	})

	t.Run("InvalidDomain_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, _, err := s.DownloadDomainVerificationFile("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("StreamError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		errorCloser := &test.TestReadCloser{
			ReadError: errors.New("test read error"),
		}

		response := &http.Response{
			Body:       errorCloser,
			StatusCode: 200,
		}
		response.Header = http.Header{}
		response.Header.Set("Content-Disposition", "attachment; filename=\"testfile1.txt\"")

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/verify/file-upload", gomock.Any()).Return(&connection.APIResponse{
			Response: response,
		}, nil)

		_, _, err := s.DownloadDomainVerificationFile("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.Equal(t, "test read error", err.Error())
	})

	t.Run("MissingContentDispositionHeader_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}
		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/verify/file-upload", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("test content"))),
				StatusCode: 200,
			},
		}, nil)

		_, _, err := s.DownloadDomainVerificationFile("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.Equal(t, "mime: no media type", err.Error())
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/verify/file-upload", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, _, err := s.DownloadDomainVerificationFile("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/verify/file-upload", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, _, err := s.DownloadDomainVerificationFile("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.IsType(t, &DomainNotFoundError{}, err)
	})
}

func TestAddDomainCDNConfiguration(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/cdn", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
				StatusCode: 201,
			},
		}, nil).Times(1)

		err := s.AddDomainCDNConfiguration("testdomain1.co.uk")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/cdn", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.AddDomainCDNConfiguration("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.AddDomainCDNConfiguration("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/cdn", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.AddDomainCDNConfiguration("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.IsType(t, &DomainNotFoundError{}, err)
	})
}

func TestDeleteDomainCDNConfiguration(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/cdn", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.DeleteDomainCDNConfiguration("testdomain1.co.uk")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/cdn", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteDomainCDNConfiguration("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteDomainCDNConfiguration("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainCDNConfigurationNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/cdn", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteDomainCDNConfiguration("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.IsType(t, &DomainCDNConfigurationNotFoundError{}, err)
	})
}

func TestCreateDomainCDNRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		expectedRequest := CreateCDNRuleRequest{
			URI:          "/testuri",
			CacheControl: CDNRuleCacheControlCustom,
			MimeTypes:    []string{"application/test"},
			Type:         CDNRuleTypePerURI,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/cdn/rules", gomock.Eq(&expectedRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 201,
			},
		}, nil).Times(1)

		id, err := s.CreateDomainCDNRule("testdomain1.co.uk", expectedRequest)

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", id)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/cdn/rules", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateDomainCDNRule("testdomain1.co.uk", CreateCDNRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.CreateDomainCDNRule("", CreateCDNRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainCDNConfigurationNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/cdn/rules", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.CreateDomainCDNRule("testdomain1.co.uk", CreateCDNRuleRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainCDNConfigurationNotFoundError{}, err)
	})
}

func TestGetDomainCDNRules(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/cdn/rules", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rules, err := s.GetDomainCDNRules("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, rules, 1)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rules[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/cdn/rules", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/cdn/rules", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000001\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		rules, err := s.GetDomainCDNRules("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, rules, 2)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rules[0].ID)
		assert.Equal(t, "00000000-0000-0000-0000-000000000001", rules[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/cdn/rules", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetDomainCDNRules("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetDomainCDNRulesPaginated(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/cdn/rules", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rules, err := s.GetDomainCDNRulesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rules[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/cdn/rules", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainCDNRulesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainCDNRules("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainCDNConfigurationNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/cdn/rules", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainCDNRulesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainCDNConfigurationNotFoundError{}, err)
	})
}

func TestGetDomainCDNRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/cdn/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rule, err := s.GetDomainCDNRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rule.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/cdn/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainCDNRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainCDNRule("", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRuleID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainCDNRule("testdomain1.co.uk", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule ID", err.Error())
	})

	t.Run("404_ReturnsCDNRuleNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/cdn/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainCDNRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &CDNRuleNotFoundError{}, err)
	})
}

func TestPatchDomainCDNRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		expectedRequest := PatchCDNRuleRequest{
			URI: "/testuri",
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/cdn/rules/00000000-0000-0000-0000-000000000000", gomock.Eq(&expectedRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PatchDomainCDNRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", expectedRequest)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/cdn/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchDomainCDNRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchCDNRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainCDNRule("", "00000000-0000-0000-0000-000000000000", PatchCDNRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRuleID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainCDNRule("testdomain1.co.uk", "", PatchCDNRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule ID", err.Error())
	})

	t.Run("404_ReturnsCDNRuleNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/cdn/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PatchDomainCDNRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchCDNRuleRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &CDNRuleNotFoundError{}, err)
	})
}

func TestDeleteDomainCDNRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/cdn/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.DeleteDomainCDNRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/cdn/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteDomainCDNRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteDomainCDNRule("", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRuleID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteDomainCDNRule("testdomain1.co.uk", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule ID", err.Error())
	})

	t.Run("404_ReturnsCDNRuleNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/cdn/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteDomainCDNRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &CDNRuleNotFoundError{}, err)
	})
}

func TestPurgeDomainCDN(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		expectedRequest := PurgeCDNRequest{
			RecordName: "sub.example.com",
			URI:        "/testuri",
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/cdn/purge", gomock.Eq(&expectedRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{}}"))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.PurgeDomainCDN("testdomain1.co.uk", expectedRequest)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/cdn/purge", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PurgeDomainCDN("testdomain1.co.uk", PurgeCDNRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PurgeDomainCDN("", PurgeCDNRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainCDNConfigurationNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/cdn/purge", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PurgeDomainCDN("testdomain1.co.uk", PurgeCDNRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainCDNConfigurationNotFoundError{}, err)
	})
}

func TestGetDomainHSTSConfiguration(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/hsts", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"enabled\":true}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		configuration, err := s.GetDomainHSTSConfiguration("testdomain1.co.uk")

		assert.Nil(t, err)
		assert.Equal(t, true, configuration.Enabled)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/hsts", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainHSTSConfiguration("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainHSTSConfiguration("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_DomainHSTSConfigurationNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/hsts", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainHSTSConfiguration("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.IsType(t, &DomainHSTSConfigurationNotFoundError{}, err)
	})
}

func TestAddDomainHSTSConfiguration(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/hsts", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.AddDomainHSTSConfiguration("testdomain1.co.uk")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/hsts", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.AddDomainHSTSConfiguration("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.AddDomainHSTSConfiguration("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/hsts", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.AddDomainHSTSConfiguration("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.IsType(t, &DomainNotFoundError{}, err)
	})
}

func TestDeleteDomainHSTSConfiguration(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/hsts", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.DeleteDomainHSTSConfiguration("testdomain1.co.uk")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/hsts", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteDomainHSTSConfiguration("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteDomainHSTSConfiguration("")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainHSTSConfigurationNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/hsts", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteDomainHSTSConfiguration("testdomain1.co.uk")

		assert.NotNil(t, err)
		assert.IsType(t, &DomainHSTSConfigurationNotFoundError{}, err)
	})
}

func TestCreateDomainHSTSRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		expectedRequest := CreateHSTSRuleRequest{
			Type: HSTSRuleTypeDomain,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/hsts/rules", gomock.Eq(&expectedRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 201,
			},
		}, nil).Times(1)

		id, err := s.CreateDomainHSTSRule("testdomain1.co.uk", expectedRequest)

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", id)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/hsts/rules", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.CreateDomainHSTSRule("testdomain1.co.uk", CreateHSTSRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.CreateDomainHSTSRule("", CreateHSTSRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainHSTSConfigurationNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Post("/ddosx/v1/domains/testdomain1.co.uk/hsts/rules", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.CreateDomainHSTSRule("testdomain1.co.uk", CreateHSTSRuleRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainHSTSConfigurationNotFoundError{}, err)
	})
}

func TestGetDomainHSTSRules(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/hsts/rules", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rules, err := s.GetDomainHSTSRules("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, rules, 1)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rules[0].ID)
	})

	t.Run("Multiple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		gomock.InOrder(
			c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/hsts/rules", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
			c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/hsts/rules", gomock.Any()).Return(&connection.APIResponse{
				Response: &http.Response{
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000001\"}],\"meta\":{\"pagination\":{\"total_pages\":2}}}"))),
					StatusCode: 200,
				},
			}, nil),
		)

		rules, err := s.GetDomainHSTSRules("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Len(t, rules, 2)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rules[0].ID)
		assert.Equal(t, "00000000-0000-0000-0000-000000000001", rules[1].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/hsts/rules", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1"))

		_, err := s.GetDomainHSTSRules("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})
}

func TestGetDomainHSTSRulesPaginated(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/hsts/rules", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":[{\"id\":\"00000000-0000-0000-0000-000000000000\"}]}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rules, err := s.GetDomainHSTSRulesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rules[0].ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/hsts/rules", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainHSTSRulesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainHSTSRules("", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("404_ReturnsDomainHSTSConfigurationNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/hsts/rules", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainHSTSRulesPaginated("testdomain1.co.uk", connection.APIRequestParameters{})

		assert.NotNil(t, err)
		assert.IsType(t, &DomainHSTSConfigurationNotFoundError{}, err)
	})
}

func TestGetDomainHSTSRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/hsts/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		rule, err := s.GetDomainHSTSRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", rule.ID)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/hsts/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		_, err := s.GetDomainHSTSRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainHSTSRule("", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRuleID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		_, err := s.GetDomainHSTSRule("testdomain1.co.uk", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule ID", err.Error())
	})

	t.Run("404_ReturnsHSTSRuleNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Get("/ddosx/v1/domains/testdomain1.co.uk/hsts/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		_, err := s.GetDomainHSTSRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &HSTSRuleNotFoundError{}, err)
	})
}

func TestPatchDomainHSTSRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		expectedRequest := PatchHSTSRuleRequest{
			MaxAge: ptr.Int(300),
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/hsts/rules/00000000-0000-0000-0000-000000000000", gomock.Eq(&expectedRequest)).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 200,
			},
		}, nil).Times(1)

		err := s.PatchDomainHSTSRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", expectedRequest)

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/hsts/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.PatchDomainHSTSRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchHSTSRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainHSTSRule("", "00000000-0000-0000-0000-000000000000", PatchHSTSRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRuleID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.PatchDomainHSTSRule("testdomain1.co.uk", "", PatchHSTSRuleRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule ID", err.Error())
	})

	t.Run("404_ReturnsHSTSRuleNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Patch("/ddosx/v1/domains/testdomain1.co.uk/hsts/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.PatchDomainHSTSRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000", PatchHSTSRuleRequest{})

		assert.NotNil(t, err)
		assert.IsType(t, &HSTSRuleNotFoundError{}, err)
	})
}

func TestDeleteDomainHSTSRule(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/hsts/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("{\"data\":{\"id\":\"00000000-0000-0000-0000-000000000000\"}}"))),
				StatusCode: 204,
			},
		}, nil).Times(1)

		err := s.DeleteDomainHSTSRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.Nil(t, err)
	})

	t.Run("ConnectionError_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/hsts/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{}, errors.New("test error 1")).Times(1)

		err := s.DeleteDomainHSTSRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
	})

	t.Run("InvalidDomainName_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteDomainHSTSRule("", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid domain name", err.Error())
	})

	t.Run("InvalidRuleID_ReturnsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		err := s.DeleteDomainHSTSRule("testdomain1.co.uk", "")

		assert.NotNil(t, err)
		assert.Equal(t, "invalid rule ID", err.Error())
	})

	t.Run("404_ReturnsHSTSRuleNotFoundError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c := mocks.NewMockConnection(mockCtrl)

		s := Service{
			connection: c,
		}

		c.EXPECT().Delete("/ddosx/v1/domains/testdomain1.co.uk/hsts/rules/00000000-0000-0000-0000-000000000000", gomock.Any()).Return(&connection.APIResponse{
			Response: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				StatusCode: 404,
			},
		}, nil).Times(1)

		err := s.DeleteDomainHSTSRule("testdomain1.co.uk", "00000000-0000-0000-0000-000000000000")

		assert.NotNil(t, err)
		assert.IsType(t, &HSTSRuleNotFoundError{}, err)
	})
}
