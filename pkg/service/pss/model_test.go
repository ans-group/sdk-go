package pss

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorType_String_Expected(t *testing.T) {
	v := AuthorTypeAuto

	s := v.String()

	assert.Equal(t, "Auto", s)
}

func TestRequestPriority_String_Expected(t *testing.T) {
	v := RequestPriorityNormal

	s := v.String()

	assert.Equal(t, "Normal", s)
}

func TestRequestStatus_String_Expected(t *testing.T) {
	v := RequestStatusCompleted

	s := v.String()

	assert.Equal(t, "Completed", s)
}
