package safedns

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRecordTTL_Time_Expected(t *testing.T) {
	var ttl RecordTTL = 300

	expectedTime := time.Now().Add(time.Second * 300)
	ttlTime := ttl.Time()

	assert.Equal(t, expectedTime.Second(), ttlTime.Second())
}

func TestRecordTTL_Duration_Expected(t *testing.T) {
	var ttl RecordTTL = 300

	expectedDuration := time.Second * 300
	ttlDuration := ttl.Duration()

	assert.Equal(t, expectedDuration, ttlDuration)
}

func TestRecordTTL_String_Expected(t *testing.T) {
	var ttl RecordTTL = 300

	s := ttl.String()

	assert.Equal(t, "300", s)
}

func TestRecord_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := Record{
			Name:    "www.testdomain1.com",
			Content: "1.2.3.4",
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := Record{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}

func TestTemplate_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := Template{
			Name: "www.testdomain1.com",
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := Template{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}
