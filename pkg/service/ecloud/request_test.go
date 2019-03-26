package ecloud

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatchTagRequest_Validate_NoError(t *testing.T) {
	r := PatchTagRequest{}

	err := r.Validate()

	assert.Nil(t, err)
}

func TestCreateTagRequest_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := CreateTagRequest{
			Key:   "testkey1",
			Value: "testvalue1",
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := CreateTagRequest{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}

func TestCreateVirtualMachineRequest_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := CreateVirtualMachineRequest{
			Environment: "testenvironment",
			Template:    "testtemplate",
			CPU:         2,
			RAM:         4,
			HDD:         20,
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := CreateVirtualMachineRequest{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}

func TestPatchSolutionRequest_Validate_NoError(t *testing.T) {
	r := PatchSolutionRequest{}

	err := r.Validate()

	assert.Nil(t, err)
}

func TestRenameTemplateRequest_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := RenameTemplateRequest{
			Destination: "testtemplate1",
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := RenameTemplateRequest{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}

func TestCloneVirtualMachineRequest_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := CloneVirtualMachineRequest{
			Name: "testvm1",
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := CloneVirtualMachineRequest{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}

func TestPatchVirtualMachineRequest_Validate_NoError(t *testing.T) {
	r := PatchVirtualMachineRequest{}

	err := r.Validate()

	assert.Nil(t, err)
}

func TestPatchVirtualMachineRequestDiskState_String_Expected(t *testing.T) {
	v := PatchVirtualMachineRequestDiskStatePresent

	s := v.String()

	assert.Equal(t, "present", s)
}
