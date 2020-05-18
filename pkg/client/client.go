package client

import (
	"github.com/ukfast/sdk-go/pkg/connection"
	"github.com/ukfast/sdk-go/pkg/service/account"
	"github.com/ukfast/sdk-go/pkg/service/ddosx"
	"github.com/ukfast/sdk-go/pkg/service/draas"
	"github.com/ukfast/sdk-go/pkg/service/ecloud"
	"github.com/ukfast/sdk-go/pkg/service/ltaas"
	"github.com/ukfast/sdk-go/pkg/service/pss"
	"github.com/ukfast/sdk-go/pkg/service/registrar"
	"github.com/ukfast/sdk-go/pkg/service/safedns"
	"github.com/ukfast/sdk-go/pkg/service/ssl"
	"github.com/ukfast/sdk-go/pkg/service/storage"
)

type Client interface {
	AccountService() account.AccountService
	DDoSXService() ddosx.DDoSXService
	DRaaSService() draas.DRaaSService
	ECloudService() ecloud.ECloudService
	LTaaSService() ltaas.LTaaSService
	PSSService() pss.PSSService
	RegistrarService() registrar.RegistrarService
	SafeDNSService() safedns.SafeDNSService
	SSLService() ssl.SSLService
	StorageService() storage.StorageService
}

type UKFastClient struct {
	connection connection.Connection
}

func NewClient(connection connection.Connection) *UKFastClient {
	return &UKFastClient{
		connection: connection,
	}
}

func (c *UKFastClient) AccountService() account.AccountService {
	return account.NewService(c.connection)
}

func (c *UKFastClient) DDoSXService() ddosx.DDoSXService {
	return ddosx.NewService(c.connection)
}

func (c *UKFastClient) DRaaSService() draas.DRaaSService {
	return draas.NewService(c.connection)
}

func (c *UKFastClient) ECloudService() ecloud.ECloudService {
	return ecloud.NewService(c.connection)
}

func (c *UKFastClient) LTaaSService() ltaas.LTaaSService {
	return ltaas.NewService(c.connection)
}

func (c *UKFastClient) PSSService() pss.PSSService {
	return pss.NewService(c.connection)
}

func (c *UKFastClient) RegistrarService() registrar.RegistrarService {
	return registrar.NewService(c.connection)
}

func (c *UKFastClient) SafeDNSService() safedns.SafeDNSService {
	return safedns.NewService(c.connection)
}

func (c *UKFastClient) SSLService() ssl.SSLService {
	return ssl.NewService(c.connection)
}

func (c *UKFastClient) StorageService() storage.StorageService {
	return storage.NewService(c.connection)
}
