package pss

import (
	"io"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// PSSService is an interface for managing PSS
type PSSService interface {
	CreateRequest(req CreateRequestRequest) (int, error)
	GetRequests(parameters connection.APIRequestParameters) ([]Request, error)
	GetRequestsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Request], error)
	GetRequest(requestID int) (Request, error)
	PatchRequest(requestID int, req PatchRequestRequest) error

	GetRequestFeedback(requestID int) (Feedback, error)
	CreateRequestFeedback(requestID int, req CreateFeedbackRequest) (int, error)

	CreateRequestReply(requestID int, req CreateReplyRequest) (string, error)
	GetRequestReplies(solutionID int, parameters connection.APIRequestParameters) ([]Reply, error)
	GetRequestRepliesPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[Reply], error)
	GetRequestConversation(requestID int, parameters connection.APIRequestParameters) ([]Reply, error)
	GetRequestConversationPaginated(requestID int, parameters connection.APIRequestParameters) (*connection.Paginated[Reply], error)

	GetReply(replyID string) (Reply, error)

	DownloadReplyAttachmentStream(replyID string, attachmentName string) (contentStream io.ReadCloser, err error)
	UploadReplyAttachmentStream(replyID string, attachmentName string, fileStream io.Reader) (err error)
	DeleteReplyAttachment(replyID string, attachmentName string) error

	GetCaseCategories(parameters connection.APIRequestParameters) ([]CaseCategory, error)
	GetCaseCategoriesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseCategory], error)

	GetChangeImpactCaseOptions(parameters connection.APIRequestParameters) ([]CaseOption, error)
	GetChangeImpactCaseOptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseOption], error)
	GetChangeRiskCaseOptions(parameters connection.APIRequestParameters) ([]CaseOption, error)
	GetChangeRiskCaseOptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseOption], error)
	GetIncidentImpactCaseOptions(parameters connection.APIRequestParameters) ([]CaseOption, error)
	GetIncidentImpactCaseOptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseOption], error)
	GetIncidentTypeCaseOptions(parameters connection.APIRequestParameters) ([]CaseOption, error)
	GetIncidentTypeCaseOptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseOption], error)

	GetIncidentCases(parameters connection.APIRequestParameters) ([]IncidentCase, error)
	GetIncidentCasesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[IncidentCase], error)
	GetIncidentCase(incidentID string) (IncidentCase, error)

	GetChangeCases(parameters connection.APIRequestParameters) ([]ChangeCase, error)
	GetChangeCasesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ChangeCase], error)
	GetProblemCases(parameters connection.APIRequestParameters) ([]ProblemCase, error)
	GetProblemCasesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ProblemCase], error)
}

// Service implements PSSService for managing
// PSS via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of PSSService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
