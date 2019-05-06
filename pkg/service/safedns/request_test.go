package safedns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateZoneRequest_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := CreateZoneRequest{
			Name: "testdomain1.com",
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := CreateZoneRequest{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}

func TestPatchZoneRequest_Validate_NoError(t *testing.T) {
	r := PatchZoneRequest{}

	err := r.Validate()

	assert.Nil(t, err)
}

func TestPatchRecordRequest_Validate_NoError(t *testing.T) {
	r := PatchRecordRequest{}

	err := r.Validate()

	assert.Nil(t, err)
}

func TestCreateRecordRequest_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := CreateRecordRequest{
			Name:    "www.testdomain1.com",
			Type:    "A",
			Content: "1.2.3.4",
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := CreateRecordRequest{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}

func TestCreateNoteRequest_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := CreateNoteRequest{
			Notes: "testnote",
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := CreateNoteRequest{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}

func TestPatchTemplateRequest_Validate_NoError(t *testing.T) {
	r := PatchTemplateRequest{}

	err := r.Validate()

	assert.Nil(t, err)
}

func TestCreateTemplateRequest_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := CreateTemplateRequest{
			Name: "testtemplate1",
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := CreateTemplateRequest{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}
