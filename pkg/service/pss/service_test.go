package pss

import (
	"testing"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/stretchr/testify/assert"
)

func TestNewService_InjectsConnection(t *testing.T) {
	c := connection.NewAPIKeyCredentialsAPIConnection("testkey")
	service := NewService(c)

	assert.True(t, (service.connection == c), "expected connection to be the same ptr")
}
