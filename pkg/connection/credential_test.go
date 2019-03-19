package connection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIKeyCredentials_GetAuthHeaders_CorrectHeaders(t *testing.T) {
	credentials := APIKeyCredentials{
		APIKey: "testapikey",
	}

	h := credentials.GetAuthHeaders()

	assert.Len(t, h, 1)
	assert.Equal(t, "testapikey", h["Authorization"])
}
