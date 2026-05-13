package pss

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) requestRes() *resource.Resource[Request, int] {
	return resource.NewIntResource[Request](s.connection, "/pss/v1/requests", "request",
		func(id int) error { return &RequestNotFoundError{ID: id} })
}

// CreateRequest creates a new request
func (s *Service) CreateRequest(req CreateRequestRequest) (int, error) {
	data, err := s.requestRes().Create(&req)
	return data.ID, err
}

// GetRequests retrieves a list of requests
func (s *Service) GetRequests(parameters connection.APIRequestParameters) ([]Request, error) {
	return s.requestRes().List(parameters)
}

// GetRequestsPaginated retrieves a paginated list of requests
func (s *Service) GetRequestsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Request], error) {
	return s.requestRes().ListPaginated(parameters)
}

// GetRequest retrieves a single request by id
func (s *Service) GetRequest(requestID int) (Request, error) {
	return s.requestRes().Get(requestID)
}

// PatchRequest patches a request
func (s *Service) PatchRequest(requestID int, req PatchRequestRequest) error {
	return s.requestRes().Patch(requestID, &req)
}

// CreateRequestReply creates a new request reply
func (s *Service) CreateRequestReply(requestID int, req CreateReplyRequest) (string, error) {
	if requestID < 1 {
		return "", fmt.Errorf("invalid request id")
	}
	body, err := connection.Post[Reply](s.connection, fmt.Sprintf("/pss/v1/requests/%d/replies", requestID), &req, connection.NotFoundResponseHandler(&RequestNotFoundError{ID: requestID}))
	return body.Data.ID, err
}

// GetRequestReplies is an alias for GetRequestConversation
func (s *Service) GetRequestReplies(solutionID int, parameters connection.APIRequestParameters) ([]Reply, error) {
	return s.GetRequestConversation(solutionID, parameters)
}

// GetRequestRepliesPaginated is an alias for GetRequestConversationPaginated
func (s *Service) GetRequestRepliesPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[Reply], error) {
	return s.GetRequestConversationPaginated(solutionID, parameters)
}

func (s *Service) requestConversationRes() *resource.SubResourceList[Reply, int] {
	return resource.NewIntSubResourceList[Reply](s.connection,
		func(requestID int) string { return fmt.Sprintf("/pss/v1/requests/%d/conversation", requestID) },
		"request", "id", func(requestID int) error { return &RequestNotFoundError{ID: requestID} })
}

// GetRequestConversation retrieves a list of replies
func (s *Service) GetRequestConversation(solutionID int, parameters connection.APIRequestParameters) ([]Reply, error) {
	return s.requestConversationRes().List(solutionID, parameters)
}

// GetRequestConversationPaginated retrieves a paginated list of domains
func (s *Service) GetRequestConversationPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[Reply], error) {
	return s.requestConversationRes().ListPaginated(solutionID, parameters)
}

// GetRequestFeedback retrieves feedback for a request
func (s *Service) GetRequestFeedback(requestID int) (Feedback, error) {
	if requestID < 1 {
		return Feedback{}, fmt.Errorf("invalid request id")
	}
	body, err := connection.Get[Feedback](s.connection, fmt.Sprintf("/pss/v1/requests/%d/feedback", requestID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&RequestFeedbackNotFoundError{RequestID: requestID}))
	return body.Data, err
}

// CreateRequestFeedback creates a new request feedback
func (s *Service) CreateRequestFeedback(requestID int, req CreateFeedbackRequest) (int, error) {
	if requestID < 1 {
		return 0, fmt.Errorf("invalid request id")
	}
	body, err := connection.Post[Feedback](s.connection, fmt.Sprintf("/pss/v1/requests/%d/feedback", requestID), &req, connection.NotFoundResponseHandler(&RequestNotFoundError{ID: requestID}))
	return body.Data.ID, err
}
