package ltaas

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetJobs retrieves a list of jobs
func (s *Service) GetJobs(parameters connection.APIRequestParameters) ([]Job, error) {
	var sites []Job

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetJobsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, site := range response.(*PaginatedJob).Items {
			sites = append(sites, site)
		}
	}

	return sites, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetJobsPaginated retrieves a paginated list of jobs
func (s *Service) GetJobsPaginated(parameters connection.APIRequestParameters) (*PaginatedJob, error) {
	body, err := s.getJobsPaginatedResponseBody(parameters)

	return NewPaginatedJob(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetJobsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getJobsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetJobsResponseBody, error) {
	body := &GetJobsResponseBody{}

	response, err := s.connection.Get("/ltaas/v1/jobs", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetJob retrieves a single job by id
func (s *Service) GetJob(jobID string) (Job, error) {
	body, err := s.getJobResponseBody(jobID)

	return body.Data, err
}

func (s *Service) getJobResponseBody(jobID string) (*GetJobResponseBody, error) {
	body := &GetJobResponseBody{}

	if jobID == "" {
		return body, fmt.Errorf("invalid job id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ltaas/v1/jobs/%s", jobID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &JobNotFoundError{ID: jobID}
		}

		return nil
	})
}

// GetJobResults retrieves the results of a single job by id
func (s *Service) GetJobResults(jobID string) (JobResults, error) {
	body, err := s.getJobResultsResponseBody(jobID)

	return body.Data, err
}

func (s *Service) getJobResultsResponseBody(jobID string) (*GetJobResultsResponseBody, error) {
	body := &GetJobResultsResponseBody{}

	if jobID == "" {
		return body, fmt.Errorf("invalid job id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ltaas/v1/jobs/%s/results", jobID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &JobNotFoundError{ID: jobID}
		}

		return nil
	})
}

// CreateJob creates a new job
func (s *Service) CreateJob(req CreateJobRequest) (string, error) {
	body, err := s.createJobResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createJobResponseBody(req CreateJobRequest) (*GetJobResponseBody, error) {
	body := &GetJobResponseBody{}

	response, err := s.connection.Post("/ltaas/v1/jobs", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// DeleteJob removes a job
func (s *Service) DeleteJob(jobID string) error {
	_, err := s.deleteJobResponseBody(jobID)

	return err
}

func (s *Service) deleteJobResponseBody(jobID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if jobID == "" {
		return body, fmt.Errorf("invalid job id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ltaas/v1/jobs/%s", jobID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &JobNotFoundError{ID: jobID}
		}

		return nil
	})
}
