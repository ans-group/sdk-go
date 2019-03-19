package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContactType_String_Expected(t *testing.T) {
	ct := ContactTypeAccounts

	s := ct.String()

	assert.Equal(t, "Accounts", s)
}
