package ddosx

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDomainStatus_String_Expected(t *testing.T) {
	v := DomainStatusConfigured

	s := v.String()

	assert.Equal(t, "Configured", s)
}

func TestDomainPropertyName_String_Expected(t *testing.T) {
	v := DomainPropertyNameSecureOrigin

	s := v.String()

	assert.Equal(t, "secure_origin", s)
}

func TestParseDomainPropertyName(t *testing.T) {
	t.Run("Valid_ReturnsEnum", func(t *testing.T) {
		v := "secure_origin"
		s, err := ParseDomainPropertyName(v)

		assert.Nil(t, err)
		assert.Equal(t, DomainPropertyNameSecureOrigin, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalid"
		_, err := ParseDomainPropertyName(v)

		assert.NotNil(t, err)
		assert.Equal(t, "Invalid ddosx.DomainPropertyName. Valid values: client_max_body_size, proxy_timeout, ipv6_enabled, secure_origin", err.Error())
	})
}

func TestRecordType_String_Expected(t *testing.T) {
	v := RecordTypeAAAA

	s := v.String()

	assert.Equal(t, "AAAA", s)
}

func TestWAFMode_String_Expected(t *testing.T) {
	v := WAFModeOn

	s := v.String()

	assert.Equal(t, "On", s)
}

func TestParseWAFMode(t *testing.T) {
	t.Run("ParsesOn", func(t *testing.T) {
		v := "on"
		s, err := ParseWAFMode(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFModeOn, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalidmode"
		_, err := ParseWAFMode(v)

		assert.NotNil(t, err)
		assert.Equal(t, "Invalid ddosx.WAFMode. Valid values: On, Off, DetectionOnly", err.Error())
	})
}

func TestWAFParanoiaLevel_String_Expected(t *testing.T) {
	v := WAFParanoiaLevelHigh

	s := v.String()

	assert.Equal(t, "High", s)
}

func TestParseWAFParanoiaLevel(t *testing.T) {
	t.Run("ParsesLow", func(t *testing.T) {
		v := "low"
		s, err := ParseWAFParanoiaLevel(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFParanoiaLevelLow, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalidlevel"
		_, err := ParseWAFParanoiaLevel(v)

		assert.NotNil(t, err)
		assert.Equal(t, "Invalid ddosx.WAFParanoiaLevel. Valid values: Low, Medium, High, Highest", err.Error())
	})
}

func TestWAFRuleSetName_String_Expected(t *testing.T) {
	v := WAFRuleSetNameIPRepution

	s := v.String()

	assert.Equal(t, "IP Reputation", s)
}

func TestWAFAdvancedRuleSection_String_Expected(t *testing.T) {
	v := WAFAdvancedRuleSectionRequestURI

	s := v.String()

	assert.Equal(t, "REQUEST_URI", s)
}

func TestParseWAFAdvancedRuleSection(t *testing.T) {
	t.Run("ParsesArgs", func(t *testing.T) {
		v := "ARGS"
		s, err := ParseWAFAdvancedRuleSection(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFAdvancedRuleSectionArgs, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalidsection"
		_, err := ParseWAFAdvancedRuleSection(v)

		assert.NotNil(t, err)
		assert.Equal(t, "Invalid ddosx.WAFAdvancedRuleSection. Valid values: ARGS, MATCHED_VARS, REMOTE_HOST, REQUEST_BODY, REQUEST_COOKIES, REQUEST_HEADERS, REQUEST_URI", err.Error())
	})
}

func TestWAFAdvancedRuleModifier_String_Expected(t *testing.T) {
	v := WAFAdvancedRuleModifierContains

	s := v.String()

	assert.Equal(t, "contains", s)
}

func TestParseWAFAdvancedRuleModifier(t *testing.T) {
	t.Run("ParsesBeginsWith", func(t *testing.T) {
		v := "beginswith"
		s, err := ParseWAFAdvancedRuleModifier(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFAdvancedRuleModifierBeginsWith, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalidmodifier"
		_, err := ParseWAFAdvancedRuleModifier(v)

		assert.NotNil(t, err)
		assert.Equal(t, "Invalid ddosx.WAFAdvancedRuleModifier. Valid values: beginsWith, endsWith, contains, containsWord", err.Error())
	})
}

func TestACLIPMode_String_Expected(t *testing.T) {
	v := ACLIPModeAllow

	s := v.String()

	assert.Equal(t, "Allow", s)
}

func TestParseACLIPMode(t *testing.T) {
	t.Run("ParsesAllow", func(t *testing.T) {
		v := "allow"
		s, err := ParseACLIPMode(v)

		assert.Nil(t, err)
		assert.Equal(t, ACLIPModeAllow, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalidmode"
		_, err := ParseACLIPMode(v)

		assert.NotNil(t, err)
		assert.Equal(t, "Invalid ddosx.ACLIPMode. Valid values: Allow, Deny", err.Error())
	})
}

func TestACLGeoIPRulesMode_String_Expected(t *testing.T) {
	v := ACLGeoIPRulesModeWhitelist

	s := v.String()

	assert.Equal(t, "Whitelist", s)
}

func TestParseACLGeoIPRulesMode(t *testing.T) {
	t.Run("ParsesWhitelist", func(t *testing.T) {
		v := "whitelist"
		s, err := ParseACLGeoIPRulesMode(v)

		assert.Nil(t, err)
		assert.Equal(t, ACLGeoIPRulesModeWhitelist, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalidmode"
		_, err := ParseACLGeoIPRulesMode(v)

		assert.NotNil(t, err)
		assert.Equal(t, "Invalid ddosx.ACLGeoIPRulesMode. Valid values: Whitelist, Blacklist", err.Error())
	})
}

func TestCDNRuleCacheControl_String_Expected(t *testing.T) {
	v := CDNRuleCacheControlCustom

	s := v.String()

	assert.Equal(t, "Custom", s)
}

func TestParseCDNRuleCacheControl(t *testing.T) {
	t.Run("Valid_ReturnsEnum", func(t *testing.T) {
		v := "custom"
		s, err := ParseCDNRuleCacheControl(v)

		assert.Nil(t, err)
		assert.Equal(t, CDNRuleCacheControlCustom, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalid"
		_, err := ParseCDNRuleCacheControl(v)

		assert.NotNil(t, err)
		assert.Equal(t, "Invalid ddosx.CDNRuleCacheControl. Valid values: Custom, Origin", err.Error())
	})
}

func TestCDNRuleType_String_Expected(t *testing.T) {
	v := CDNRuleTypeGlobal

	s := v.String()

	assert.Equal(t, "global", s)
}

func TestParseCDNRuleType(t *testing.T) {
	t.Run("Valid_ReturnsEnum", func(t *testing.T) {
		v := "global"
		s, err := ParseCDNRuleType(v)

		assert.Nil(t, err)
		assert.Equal(t, CDNRuleTypeGlobal, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalid"
		_, err := ParseCDNRuleType(v)

		assert.NotNil(t, err)
		assert.Equal(t, "Invalid ddosx.CDNRuleType. Valid values: global, per-uri", err.Error())
	})
}

func TestHSTSRuleType_String_Expected(t *testing.T) {
	v := HSTSRuleTypeDomain

	s := v.String()

	assert.Equal(t, "domain", s)
}

func TestParseHSTSRuleType(t *testing.T) {
	t.Run("Valid_ReturnsEnum", func(t *testing.T) {
		v := "domain"
		s, err := ParseHSTSRuleType(v)

		assert.Nil(t, err)
		assert.Equal(t, HSTSRuleTypeDomain, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalid"
		_, err := ParseHSTSRuleType(v)

		assert.NotNil(t, err)
		assert.Equal(t, "Invalid ddosx.HSTSRuleType. Valid values: domain, record", err.Error())
	})
}

func TestCDNRuleCacheControlDuration_Duration_ReturnsExpected(t *testing.T) {
	d := CDNRuleCacheControlDuration{
		Years:   1,
		Months:  2,
		Days:    3,
		Hours:   4,
		Minutes: 5,
	}

	duration := d.Duration()
	str := duration.Round(time.Minute).String()

	assert.Equal(t, "10276h5m0s", str)
}

func TestCDNRuleCacheControlDuration_String_ReturnsExpected(t *testing.T) {
	d := CDNRuleCacheControlDuration{
		Years:   1,
		Months:  2,
		Days:    3,
		Hours:   4,
		Minutes: 5,
	}

	str := d.String()

	assert.Equal(t, "1y2mo3d4h5m", str)
}

func TestParseCDNRuleCacheControlDuration(t *testing.T) {
	t.Run("Valid_ReturnsParsedDuration", func(t *testing.T) {
		v := "4h1m"
		d, err := ParseCDNRuleCacheControlDuration(v)

		assert.Nil(t, err)
		assert.Equal(t, 4, d.Hours)
		assert.Equal(t, 1, d.Minutes)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalid"
		_, err := ParseCDNRuleCacheControlDuration(v)

		assert.NotNil(t, err)
		assert.Equal(t, "Digit not supplied for unit 'invalid'", err.Error())
	})
}
