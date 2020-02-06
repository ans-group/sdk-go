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

type testEnum string

var testEnumOne testEnum = "One"
var testEnumTwo testEnum = "Two"
var testEnums EnumSlice = []Enum{testEnumOne, testEnumTwo}

func (t testEnum) String() string {
	return string(t)
}

func TestEnum_ParseEnum(t *testing.T) {
	t.Run("ParsesExact", func(t *testing.T) {
		v := "One"
		s, err := ParseEnum(v, testEnums)

		assert.Nil(t, err)
		assert.Equal(t, testEnumOne, s)
	})

	t.Run("ParsesMixedCase", func(t *testing.T) {
		v := "OnE"
		s, err := ParseEnum(v, testEnums)

		assert.Nil(t, err)
		assert.Equal(t, testEnumOne, s)
	})

	t.Run("NoEnumsSupplied_ReturnsError", func(t *testing.T) {
		v := "OnE"
		_, err := ParseEnum(v, []Enum{})

		assert.NotNil(t, err)
		assert.Equal(t, "Must provide at least one enum", err.Error())
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalid"
		_, err := ParseEnum(v, testEnums)

		assert.NotNil(t, err)
		assert.IsType(t, &ErrInvalidEnumValue{}, err)
	})
}

func TestEnumSlice_StringSlice_ReturnsExpected(t *testing.T) {
	s := testEnums.StringSlice()

	assert.Equal(t, s, []string{"One", "Two"})
}

func TestEnumSlice_String_ReturnsExpected(t *testing.T) {
	s := testEnums.String()

	assert.Equal(t, s, "One, Two")
}
