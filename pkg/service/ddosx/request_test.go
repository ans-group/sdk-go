package ddosx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ukfast/sdk-go/pkg/connection"
)

func TestCreateRecordRequest_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := CreateRecordRequest{
			Name:    "testrecord.testdomain1.com",
			Content: "1.2.3.4",
			Type:    RecordTypeA,
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

func TestCreateDomainRequest_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := CreateDomainRequest{
			Name: "testdomain1.com",
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := CreateDomainRequest{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}

func TestPatchDomainPropertyRequest_Validate_NoError(t *testing.T) {
	req := PatchDomainPropertyRequest{}

	err := req.Validate()

	assert.Nil(t, err)
}

func TestPatchRecordRequest_Validate_NoError(t *testing.T) {
	req := PatchRecordRequest{}

	err := req.Validate()

	assert.Nil(t, err)
}

func TestCreateACLGeoIPRuleRequest_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := CreateACLGeoIPRuleRequest{
			Code: "GB",
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := CreateACLGeoIPRuleRequest{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}

func TestPatchACLGeoIPRuleRequest_Validate_NoError(t *testing.T) {
	req := PatchACLGeoIPRuleRequest{}

	err := req.Validate()

	assert.Nil(t, err)
}

func TestCreateACLIPRuleRequest_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := CreateACLIPRuleRequest{
			IP:   connection.IPAddress("1.2.3.4"),
			Mode: ACLIPModeAllow,
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := CreateACLIPRuleRequest{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}

func TestPatchACLIPRuleRequest_Validate_NoError(t *testing.T) {
	req := PatchACLIPRuleRequest{}

	err := req.Validate()

	assert.Nil(t, err)
}

func TestPatchACLGeoIPRulesModeRequest_Validate_NoError(t *testing.T) {
	req := PatchACLGeoIPRulesModeRequest{}

	err := req.Validate()

	assert.Nil(t, err)
}

func TestCreateWAFRequest_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := CreateWAFRequest{
			Mode:          WAFModeOn,
			ParanoiaLevel: WAFParanoiaLevelMedium,
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := CreateWAFRequest{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}

func TestPatchWAFRequest_Validate_NoError(t *testing.T) {
	req := PatchWAFRequest{}

	err := req.Validate()

	assert.Nil(t, err)
}

func TestPatchWAFRuleSetRequest_Validate_NoError(t *testing.T) {
	req := PatchWAFRuleSetRequest{}

	err := req.Validate()

	assert.Nil(t, err)
}

func TestCreateWAFRuleRequest_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := CreateWAFRuleRequest{
			URI: "/some/uri",
			IP:  connection.IPAddress("1.2.3.4"),
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := CreateWAFRuleRequest{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}

func TestPatchWAFRuleRequest_Validate_NoError(t *testing.T) {
	req := PatchWAFRuleRequest{}

	err := req.Validate()

	assert.Nil(t, err)
}

func TestCreateWAFAdvancedRuleRequest_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := CreateWAFAdvancedRuleRequest{
			Section:  WAFAdvancedRuleSectionRequestURI,
			Modifier: WAFAdvancedRuleModifierContains,
			Phrase:   "test",
			IP:       connection.IPAddress("1.2.3.4"),
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := CreateWAFAdvancedRuleRequest{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}

func TestPatchWAFAdvancedRuleRequest_Validate_NoError(t *testing.T) {
	req := PatchWAFAdvancedRuleRequest{}

	err := req.Validate()

	assert.Nil(t, err)
}

func TestCreateSSLRequest_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := CreateSSLRequest{
			FriendlyName: "testcert1",
			UKFastSSLID:  123,
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("MissingUKFastSSLID_KeyRequired", func(t *testing.T) {
		c := CreateSSLRequest{
			FriendlyName: "testcert1",
			Certificate:  "abc",
		}

		err := c.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "Key must be provided when UKFastSSLID isn't provided", err.Error())
	})

	t.Run("MissingUKFastSSLID_CertificateRequired", func(t *testing.T) {
		c := CreateSSLRequest{
			FriendlyName: "testcert1",
			Key:          "key",
		}

		err := c.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "Certificate must be provided when UKFastSSLID isn't provided", err.Error())
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := CreateSSLRequest{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}

func TestPatchSSLRequest_Validate_NoError(t *testing.T) {
	req := PatchSSLRequest{}

	err := req.Validate()

	assert.Nil(t, err)
}

func TestCreateCDNRuleRequest_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := CreateCDNRuleRequest{
			URI:          "/testuri",
			CacheControl: CDNRuleCacheControlCustom,
			MimeTypes:    []string{"application/test"},
			Type:         CDNRuleTypePerURI,
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := CreateCDNRuleRequest{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}

func TestPatchCDNRuleRequest_Validate_NoError(t *testing.T) {
	req := PatchCDNRuleRequest{}

	err := req.Validate()

	assert.Nil(t, err)
}

func TestPurgeCDNRequest_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := PurgeCDNRequest{
			RecordName: "testrecord.com",
			URI:        "/someuri",
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := PurgeCDNRequest{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}

func TestPatchHSTSRuleRequest_Validate_NoError(t *testing.T) {
	req := PatchHSTSRuleRequest{}

	err := req.Validate()

	assert.Nil(t, err)
}

func TestCreateHSTSRuleRequest_Validate(t *testing.T) {
	t.Run("Valid_NoError", func(t *testing.T) {
		c := CreateHSTSRuleRequest{
			Type: HSTSRuleTypeDomain,
		}

		err := c.Validate()

		assert.Nil(t, err)
	})

	t.Run("TypeRecordMissingRecordName_ReturnsValidationError", func(t *testing.T) {
		c := CreateHSTSRuleRequest{
			Type: HSTSRuleTypeRecord,
		}

		err := c.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "RecordName must be specified with Type 'HSTSRuleTypeRecord'", err.Error())
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		c := CreateHSTSRuleRequest{}

		err := c.Validate()

		assert.NotNil(t, err)
	})
}
