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

func TestRecordType_String_Expected(t *testing.T) {
	v := RecordTypeAAAA

	s := v.String()

	assert.Equal(t, "AAAA", s)
}
