package ecloud

import (
	"testing"

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
	t.Run("ExactOffline_Parses", func(t *testing.T) {
		status, err := ParseVirtualMachinePowerStatus("Offline")

		assert.Nil(t, err)
		assert.Equal(t, VirtualMachinePowerStatus("Offline"), status)
	})

	t.Run("ExactOnline_Parses", func(t *testing.T) {
		status, err := ParseVirtualMachinePowerStatus("Online")

		assert.Nil(t, err)
		assert.Equal(t, VirtualMachinePowerStatus("Online"), status)
	})

	t.Run("MixedCase_Parses", func(t *testing.T) {
		status, err := ParseVirtualMachinePowerStatus("OfFlInE")

		assert.Nil(t, err)
		assert.Equal(t, VirtualMachinePowerStatus("Offline"), status)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		_, err := ParseVirtualMachinePowerStatus("InvalidStatus")

		assert.NotNil(t, err)
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
