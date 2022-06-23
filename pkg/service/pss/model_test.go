package pss

import (
	"testing"

	"github.com/ans-group/sdk-go/pkg/connection"
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

func TestParseRequestPriority(t *testing.T) {
	t.Run("Valid_ReturnsEnum", func(t *testing.T) {
		v := "high"
		s, err := ParseRequestPriority(v)

		assert.Nil(t, err)
		assert.Equal(t, RequestPriorityHigh, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalid"
		_, err := ParseRequestPriority(v)

		assert.NotNil(t, err)
		assert.IsType(t, &connection.ErrInvalidEnumValue{}, err)
	})
}

func TestRequestStatus_String_Expected(t *testing.T) {
	v := RequestStatusCompleted

	s := v.String()

	assert.Equal(t, "Completed", s)
}
