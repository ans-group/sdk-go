package ltaas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainStatus_String_Expected(t *testing.T) {
	v := DomainStatusCompleted

	s := v.String()

	assert.Equal(t, "Completed", s)
}
