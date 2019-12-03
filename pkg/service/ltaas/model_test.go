package ltaas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainStatus_String_Expected(t *testing.T) {
	v := DomainStatusVerified

	s := v.String()

	assert.Equal(t, "Verified", s)
}
