package ssl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCertificateStatus_String_Expected(t *testing.T) {
	v := CertificateStatusCompleted

	s := v.String()

	assert.Equal(t, "Completed", s)
}
