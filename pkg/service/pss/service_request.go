package pss

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// CreateRequest creates a new request
func (s *Service) CreateRequest(req CreateRequestRequest) (int, error) {
	body, err := connection.Post[Request](s.connection, "/pss/v1/requests", &req)
	return body.Data.ID, err
}

// GetRequests retrieves a list of requests
func (s *Service) GetRequests(parameters connection.APIRequestParameters) ([]Request, error) {
	return connection.InvokeRequestAll(s.GetRequestsPaginated, parameters)
}

// GetRequestsPaginated retrieves a paginated list of requests
func (s *Service) GetRequestsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Request], error) {
	body, err := connection.Get[[]Request](s.connection, "/pss/v1/requests", parameters)
	return connection.NewPaginated(body, parameters, s.GetRequestsPaginated), err
}

// GetRequest retrieves a single request by id
func (s *Service) GetRequest(requestID int) (Request, error) {
	if requestID < 1 {
		return Request{}, fmt.Errorf("invalid request id")
	}
	body, err := connection.Get[Request](s.connection, fmt.Sprintf("/pss/v1/requests/%d", requestID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&RequestNotFoundError{ID: requestID}))
	return body.Data, err
}

// PatchRequest patches a request
func (s *Service) PatchRequest(requestID int, req PatchRequestRequest) error {
	if requestID < 1 {
		return fmt.Errorf("invalid request id")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/pss/v1/requests/%d", requestID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&RequestNotFoundError{ID: requestID}))
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

// GetRequestConversation retrieves a list of replies
func (s *Service) GetRequestConversation(solutionID int, parameters connection.APIRequestParameters) ([]Reply, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Reply], error) {
		return s.GetRequestConversationPaginated(solutionID, p)
	}, parameters)
}

// GetRequestConversationPaginated retrieves a paginated list of domains
func (s *Service) GetRequestConversationPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[Reply], error) {
	if solutionID < 1 {
		return nil, fmt.Errorf("invalid request id")
	}
	body, err := connection.Get[[]Reply](s.connection, fmt.Sprintf("/pss/v1/requests/%d/conversation", solutionID), parameters, connection.NotFoundResponseHandler(&RequestNotFoundError{ID: solutionID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Reply], error) {
		return s.GetRequestConversationPaginated(solutionID, p)
	}), err
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
