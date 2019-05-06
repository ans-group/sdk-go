package pss

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ukfast/sdk-go/pkg/connection"
)

func TestNewService_InjectsConnection(t *testing.T) {
	c := connection.NewAPIKeyCredentialsAPIConnection("testkey")
	service := NewService(c)

	assert.True(t, (service.connection == c), "expected connection to be the same ptr")
}
