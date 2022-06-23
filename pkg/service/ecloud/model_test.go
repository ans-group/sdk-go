package ecloud

import (
	"testing"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/stretchr/testify/assert"
)

func TestVirtualMachineStatus_String_Expected(t *testing.T) {
	v := VirtualMachineStatusFailed

	s := v.String()

	assert.Equal(t, "Failed", s)
}

func TestVirtualMachineDiskType_String_Expected(t *testing.T) {
	v := VirtualMachineDiskTypeCluster

	s := v.String()

	assert.Equal(t, "Cluster", s)
}

func TestVirtualMachinePowerStatus_String_Expected(t *testing.T) {
	v := VirtualMachinePowerStatusOffline

	s := v.String()

	assert.Equal(t, "Offline", s)
}

func TestParseVirtualMachinePowerStatus(t *testing.T) {
	t.Run("Valid_ReturnsEnum", func(t *testing.T) {
		v := "Online"
		s, err := ParseVirtualMachinePowerStatus(v)

		assert.Nil(t, err)
		assert.Equal(t, VirtualMachinePowerStatusOnline, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalid"
		_, err := ParseVirtualMachinePowerStatus(v)

		assert.NotNil(t, err)
		assert.IsType(t, &connection.ErrInvalidEnumValue{}, err)
	})
}

func TestDatastoreStatus_String_Expected(t *testing.T) {
	v := DatastoreStatusFailed

	s := v.String()

	assert.Equal(t, "Failed", s)
}

func TestSolutionEnvironment_String_Expected(t *testing.T) {
	v := SolutionEnvironmentHybrid

	s := v.String()

	assert.Equal(t, "Hybrid", s)
}

func TestFirewallRole_String_Expected(t *testing.T) {
	v := FirewallRoleMaster

	s := v.String()

	assert.Equal(t, "Master", s)
}

func TestParseTemplateType(t *testing.T) {
	t.Run("Valid_ReturnsEnum", func(t *testing.T) {
		v := "pod"
		s, err := ParseTemplateType(v)

		assert.Nil(t, err)
		assert.Equal(t, TemplateTypePod, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalid"
		_, err := ParseTemplateType(v)

		assert.NotNil(t, err)
		assert.IsType(t, &connection.ErrInvalidEnumValue{}, err)
	})
}

func TestParseTaskStatus(t *testing.T) {
	t.Run("Valid_ReturnsEnum", func(t *testing.T) {
		v := "complete"
		s, err := ParseTaskStatus(v)

		assert.Nil(t, err)
		assert.Equal(t, TaskStatusComplete, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalid"
		_, err := ParseTaskStatus(v)

		assert.NotNil(t, err)
		assert.IsType(t, &connection.ErrInvalidEnumValue{}, err)
	})
}
