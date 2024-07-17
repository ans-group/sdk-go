package ddosx

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCDNRuleCacheControlDuration_Duration_ReturnsExpected(t *testing.T) {
	d := CDNRuleCacheControlDuration{
		Years:   1,
		Months:  2,
		Days:    3,
		Hours:   4,
		Minutes: 5,
	}

	duration := d.Duration()
	str := duration.Round(time.Minute).String()

	assert.Equal(t, "10276h5m0s", str)
}

func TestCDNRuleCacheControlDuration_String_ReturnsExpected(t *testing.T) {
	d := CDNRuleCacheControlDuration{
		Years:   1,
		Months:  2,
		Days:    3,
		Hours:   4,
		Minutes: 5,
	}

	str := d.String()

	assert.Equal(t, "1y2mo3d4h5m", str)
}

func TestParseCDNRuleCacheControlDuration(t *testing.T) {
	t.Run("Valid_ReturnsParsedDuration", func(t *testing.T) {
		v := "4h1m"
		d, err := ParseCDNRuleCacheControlDuration(v)

		assert.Nil(t, err)
		assert.Equal(t, 4, d.Hours)
		assert.Equal(t, 1, d.Minutes)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalid"
		_, err := ParseCDNRuleCacheControlDuration(v)

		assert.NotNil(t, err)
		assert.Equal(t, "Digit not supplied for unit 'invalid'", err.Error())
	})
}
