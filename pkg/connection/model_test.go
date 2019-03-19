package connection

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDate_Time_Parses(t *testing.T) {
	t.Run("Parses", func(t *testing.T) {
		var d Date = "2018-12-13"

		timeV := d.Time()

		assert.Equal(t, 2018, timeV.Year())
		assert.Equal(t, time.December, timeV.Month())
		assert.Equal(t, 13, timeV.Day())
	})

	t.Run("InvalidDateDefaultTime", func(t *testing.T) {
		var d Date = "2018-13-13"

		timeV := d.Time()

		assert.Equal(t, 1, timeV.Year())
		assert.Equal(t, time.January, timeV.Month())
		assert.Equal(t, 1, timeV.Day())
	})
}

func TestDate_String_Expected(t *testing.T) {
	var d Date = "2018-12-13"

	s := d.String()

	assert.Equal(t, "2018-12-13", s)
}

func TestDateTime_Time_Parses(t *testing.T) {
	t.Run("Parses", func(t *testing.T) {
		var d DateTime = "2018-12-13T14:18:31+0000"

		timeV := d.Time()

		assert.Equal(t, 2018, timeV.Year())
		assert.Equal(t, time.December, timeV.Month())
		assert.Equal(t, 13, timeV.Day())
	})

	t.Run("InvalidDateDefaultTime", func(t *testing.T) {
		var d DateTime = "2018-13-13T10:58:55+0000"

		timeV := d.Time()

		assert.Equal(t, 1, timeV.Year())
		assert.Equal(t, time.January, timeV.Month())
		assert.Equal(t, 1, timeV.Day())
	})
}

func TestDateTime_String_Expected(t *testing.T) {
	var d DateTime = "2018-12-13T10:58:55+0000"

	s := d.String()

	assert.Equal(t, "2018-12-13T10:58:55+0000", s)
}

func TestIPAddress_IP_Parses(t *testing.T) {
	t.Run("Parses", func(t *testing.T) {
		var i IPAddress = "1.2.3.4"

		ip := i.IP()

		assert.Equal(t, uint8(0x1), ip[12])
		assert.Equal(t, uint8(0x2), ip[13])
		assert.Equal(t, uint8(0x3), ip[14])
		assert.Equal(t, uint8(0x4), ip[15])
	})
}

func TestIPAddress_String_Expected(t *testing.T) {
	var i IPAddress = "1.2.3.4"

	s := i.String()

	assert.Equal(t, "1.2.3.4", s)
}
