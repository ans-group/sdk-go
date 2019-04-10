package elapsed_test

import (
	"testing"
	"time"

	"github.com/ukfast/sdk-go/internal/pkg/elapsed"
	"gopkg.in/go-playground/assert.v1"
)

func TestParseDuration_ReturnsExpected(t *testing.T) {
	day := time.Duration(time.Hour * 24)

	d := time.Duration(day * 365)
	d = d + time.Duration(day*60)
	d = d + time.Duration(day*3)
	d = d + time.Duration(time.Hour*4)
	d = d + time.Duration(time.Minute*5)
	d = d + time.Duration(time.Second*6)
	d = d + time.Duration(7)

	years, months, days, hours, minutes, seconds, nanoseconds := elapsed.ParseDuration(d)

	assert.Equal(t, 1, years)
	assert.Equal(t, 2, months)
	assert.Equal(t, 3, days)
	assert.Equal(t, 4, hours)
	assert.Equal(t, 5, minutes)
	assert.Equal(t, 6, seconds)
	assert.Equal(t, 7, nanoseconds)
}

func TestNewDuration_ReturnsExpected(t *testing.T) {
	d := elapsed.NewDuration(1, 2, 3, 4, 5, 6, 7)

	assert.Equal(t, time.Duration(36993906000000007), d)
}
