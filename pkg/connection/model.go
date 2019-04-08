package connection

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// Date represents date string from API
type Date string

// Time returns Time struct for DateTime
func (c Date) Time() time.Time {
	t, _ := time.Parse("2006-01-02", c.String())

	return t
}

func (c Date) String() string {
	return string(c)
}

// DateTime represents datetime string from API
type DateTime string

// Time returns Time struct for DateTime
func (c DateTime) Time() time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05-0700", c.String())

	return t
}

func (c DateTime) String() string {
	return string(c)
}

// IPAddress represents ip address string from API
type IPAddress string

func (i IPAddress) IP() net.IP {
	return net.ParseIP(i.String())
}

func (i IPAddress) String() string {
	return string(i)
}

type Enum interface {
	String() string
}

// ParseEnum parses string s against array of enums, returning parsed enum and nil error, or nil with error
func ParseEnum(name string, s string, enums []Enum) (Enum, error) {
	var validValues []string
	for _, e := range enums {
		if strings.ToUpper(s) == strings.ToUpper(e.String()) {
			return e, nil
		}

		validValues = append(validValues, e.String())
	}

	return nil, fmt.Errorf("Invalid %s. Valid values: %s", name, strings.Join(validValues, ", "))
}
