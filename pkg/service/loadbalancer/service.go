package loadbalancer

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// LoadBalancerService is an interface for managing the LoadBalancer service
type LoadBalancerService interface {
	// Cluster
	GetClusters(parameters connection.APIRequestParameters) ([]Cluster, error)
	GetClustersPaginated(parameters connection.APIRequestParameters) (*PaginatedCluster, error)
	GetCluster(clusterID string) (Cluster, error)
	PatchCluster(clusterID string, req PatchClusterRequest) error
	DeleteCluster(clusterID string) error

	// Target
	GetTargets(parameters connection.APIRequestParameters) ([]Target, error)
	GetTargetsPaginated(parameters connection.APIRequestParameters) (*PaginatedTarget, error)
	GetTarget(targetID string) (Target, error)
	CreateTarget(req CreateTargetRequest) (string, error)
	PatchTarget(targetID string, req PatchTargetRequest) error
	DeleteTarget(targetID string) error

	// Target Group
	GetTargetGroups(parameters connection.APIRequestParameters) ([]TargetGroup, error)
	GetTargetGroupsPaginated(parameters connection.APIRequestParameters) (*PaginatedTargetGroup, error)
	GetTargetGroup(groupID string) (TargetGroup, error)
	CreateTargetGroup(req CreateTargetGroupRequest) (string, error)
	PatchTargetGroup(groupID string, req PatchTargetGroupRequest) error
	DeleteTargetGroup(groupID string) error

	// VIP
	GetVIPs(parameters connection.APIRequestParameters) ([]VIP, error)
	GetVIPsPaginated(parameters connection.APIRequestParameters) (*PaginatedVIP, error)
	GetVIP(vipID string) (VIP, error)
	CreateVIP(req CreateVIPRequest) (string, error)
	PatchVIP(vipID string, req PatchVIPRequest) error
	DeleteVIP(vipID string) error

	// ErrorPage
	GetErrorPages(parameters connection.APIRequestParameters) ([]ErrorPage, error)
	GetErrorPagesPaginated(parameters connection.APIRequestParameters) (*PaginatedErrorPage, error)
	GetErrorPage(errorPageID string) (ErrorPage, error)
	CreateErrorPage(req CreateErrorPageRequest) (string, error)
	PatchErrorPage(errorPageID string, req PatchErrorPageRequest) error
	DeleteErrorPage(errorPageID string) error

	// CustomOption
	GetCustomOptions(parameters connection.APIRequestParameters) ([]CustomOption, error)
	GetCustomOptionsPaginated(parameters connection.APIRequestParameters) (*PaginatedCustomOption, error)
	GetCustomOption(optionID string) (CustomOption, error)
	CreateCustomOption(req CreateCustomOptionRequest) (string, error)
	PatchCustomOption(optionID string, req PatchCustomOptionRequest) error
	DeleteCustomOption(optionID string) error

	// Listener
	GetListeners(parameters connection.APIRequestParameters) ([]Listener, error)
	GetListenersPaginated(parameters connection.APIRequestParameters) (*PaginatedListener, error)
	GetListener(listenerID string) (Listener, error)
	CreateListener(req CreateListenerRequest) (string, error)
	PatchListener(listenerID string, req PatchListenerRequest) error
	DeleteListener(listenerID string) error

	// Listener Access
	GetListenerAccesses(listenerID string, parameters connection.APIRequestParameters) ([]Access, error)
	GetListenerAccessesPaginated(listenerID string, parameters connection.APIRequestParameters) (*PaginatedAccess, error)
	GetListenerAccess(listenerID string, accessID string) (Access, error)
	CreateListenerAccess(listenerID string, req CreateAccessRequest) (string, error)
	PatchListenerAccess(listenerID string, accessID string, req PatchAccessRequest) error
	DeleteListenerAccess(listenerID string, accessID string) error

	// Listener Bind
	GetListenerBinds(listenerID string, parameters connection.APIRequestParameters) ([]Bind, error)
	GetListenerBindsPaginated(listenerID string, parameters connection.APIRequestParameters) (*PaginatedBind, error)
	GetListenerBind(listenerID string, bindID string) (Bind, error)
	CreateListenerBind(listenerID string, req CreateBindRequest) (string, error)
	PatchListenerBind(listenerID string, bindID string, req PatchBindRequest) error
	DeleteListenerBind(listenerID string, bindID string) error

	// Listener Certificate
	GetListenerCertificates(listenerID string, parameters connection.APIRequestParameters) ([]ListenerCertificate, error)
	GetListenerCertificatesPaginated(listenerID string, parameters connection.APIRequestParameters) (*PaginatedListenerCertificate, error)
	GetListenerCertificate(listenerID string, certificateID string) (ListenerCertificate, error)
	CreateListenerCertificate(listenerID string, req CreateListenerCertificateRequest) (string, error)
	PatchListenerCertificate(listenerID string, certificateID string, req PatchListenerCertificateRequest) error
	DeleteListenerCertificate(listenerID string, certificateID string) error

	// Listener ErrorPage
	GetListenerErrorPages(listenerID string, parameters connection.APIRequestParameters) ([]ListenerErrorPage, error)
	GetListenerErrorPagesPaginated(listenerID string, parameters connection.APIRequestParameters) (*PaginatedListenerErrorPage, error)
	GetListenerErrorPage(listenerID string, errorPageID string) (ListenerErrorPage, error)
	CreateListenerErrorPage(listenerID string, req CreateListenerErrorPageRequest) (string, error)
	PatchListenerErrorPage(listenerID string, errorPageID string, req PatchListenerErrorPageRequest) error
	DeleteListenerErrorPage(listenerID string, errorPageID string) error

	// Listener SSL
	GetListenerSSLs(listenerID string, parameters connection.APIRequestParameters) ([]SSL, error)
	GetListenerSSLsPaginated(listenerID string, parameters connection.APIRequestParameters) (*PaginatedSSL, error)
	GetListenerSSL(listenerID string, sslID string) (SSL, error)
	CreateListenerSSL(listenerID string, req CreateSSLRequest) (string, error)
	PatchListenerSSL(listenerID string, sslID string, req PatchSSLRequest) error
	DeleteListenerSSL(listenerID string, sslID string) error
}

// Service implements LoadBalancerService for managing
// LoadBalancer certificates via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of LoadBalancerService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
