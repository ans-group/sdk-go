package ltaas_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/ukfast/sdk-go/pkg/connection"
	"github.com/ukfast/sdk-go/pkg/service/ltaas"
)

func TestDomainVerificationMethod_String_Expected(t *testing.T) {
	v := ltaas.DomainVerificationMethodDNS

	s := v.String()

	assert.Equal(t, "DNS", s)
}

func TestParseDomainVerificationMethod(t *testing.T) {
	t.Run("Valid_ReturnsEnum", func(t *testing.T) {
		v := "DNS"
		s, err := ltaas.ParseDomainVerificationMethod(v)

		assert.Nil(t, err)
		assert.Equal(t, ltaas.DomainVerificationMethodDNS, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalid"
		_, err := ltaas.ParseDomainVerificationMethod(v)

		assert.NotNil(t, err)
		assert.IsType(t, &connection.ErrInvalidEnumValue{}, err)
	})
}

func TestDomainStatus_String_Expected(t *testing.T) {
	v := ltaas.DomainStatusVerified

	s := v.String()

	assert.Equal(t, "Verified", s)
}

func TestTestProtocol_String_Expected(t *testing.T) {
	v := ltaas.TestProtocolHTTP

	s := v.String()

	assert.Equal(t, "http", s)
}

func TestTestRecurringType_String_Expected(t *testing.T) {
	v := ltaas.TestRecurringTypeWeekly

	s := v.String()

	assert.Equal(t, "Weekly", s)
}

func TestJobStatus_String_Expected(t *testing.T) {
	v := ltaas.JobStatusRunning

	s := v.String()

	assert.Equal(t, "Running", s)
}

func TestJobFailType_String_Expected(t *testing.T) {
	v := ltaas.JobFailTypeTest

	s := v.String()

	assert.Equal(t, "Test", s)
}

func TestParseTestDuration(t *testing.T) {
	t.Run("Valid_Parses", func(t *testing.T) {
		testDuration, err := ltaas.ParseTestDuration("1h2m3s")

		assert.Nil(t, err)
		assert.Equal(t, "01:02:03", string(*testDuration))
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		_, err := ltaas.ParseTestDuration("invalid")

		assert.NotNil(t, err)
	})
}

func TestTestDuration_Duration(t *testing.T) {
	t.Run("Valid_ReturnsExpectedDuration", func(t *testing.T) {
		testDuration := ltaas.TestDuration("01:02:03")

		d := testDuration.Duration()

		assert.Equal(t, (time.Hour*1 + time.Minute*2 + time.Second*3), d)
	})

	t.Run("LessThat8Chars_ReturnsZeroValueDuration", func(t *testing.T) {
		testDuration := ltaas.TestDuration("01")

		d := testDuration.Duration()

		assert.Equal(t, time.Duration(0), d)
	})

	t.Run("Invalid_ReturnsZeroValueDuration", func(t *testing.T) {
		testDuration := ltaas.TestDuration("invalidduration")

		d := testDuration.Duration()

		assert.Equal(t, time.Duration(0), d)
	})
}

func TestParseAgreementType(t *testing.T) {
	t.Run("Valid_ReturnsEnum", func(t *testing.T) {
		v := "single"
		s, err := ltaas.ParseAgreementType(v)

		assert.Nil(t, err)
		assert.Equal(t, ltaas.AgreementTypeSingle, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalid"
		_, err := ltaas.ParseAgreementType(v)

		assert.NotNil(t, err)
		assert.IsType(t, &connection.ErrInvalidEnumValue{}, err)
	})
}
