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

	t.Run("ParsesOff", func(t *testing.T) {
		v := "off"
		s, err := ParseWAFMode(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFModeOff, s)
	})

	t.Run("ParsesDetectionOnly", func(t *testing.T) {
		v := "detectiononly"
		s, err := ParseWAFMode(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFModeDetectionOnly, s)
	})

	t.Run("MixedCase_Parses", func(t *testing.T) {
		v := "On"
		s, err := ParseWAFMode(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFModeOn, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalidmode"
		_, err := ParseWAFMode(v)

		assert.NotNil(t, err)
		assert.Equal(t, "Invalid WAF mode", err.Error())
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
	t.Run("ParsesMedium", func(t *testing.T) {
		v := "medium"
		s, err := ParseWAFParanoiaLevel(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFParanoiaLevelMedium, s)
	})
	t.Run("ParsesHigh", func(t *testing.T) {
		v := "high"
		s, err := ParseWAFParanoiaLevel(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFParanoiaLevelHigh, s)
	})
	t.Run("ParsesHighest", func(t *testing.T) {
		v := "highest"
		s, err := ParseWAFParanoiaLevel(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFParanoiaLevelHighest, s)
	})

	t.Run("MixedCase_Parses", func(t *testing.T) {
		v := "Low"
		s, err := ParseWAFParanoiaLevel(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFParanoiaLevelLow, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalidlevel"
		_, err := ParseWAFParanoiaLevel(v)

		assert.NotNil(t, err)
		assert.Equal(t, "Invalid WAF paranoia level", err.Error())
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

	t.Run("ParsesMatchedVars", func(t *testing.T) {
		v := "MATCHED_VARS"
		s, err := ParseWAFAdvancedRuleSection(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFAdvancedRuleSectionMatchedVars, s)
	})

	t.Run("ParsesRemoteHost", func(t *testing.T) {
		v := "REMOTE_HOST"
		s, err := ParseWAFAdvancedRuleSection(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFAdvancedRuleSectionRemoteHost, s)
	})

	t.Run("ParsesRequestBody", func(t *testing.T) {
		v := "REQUEST_BODY"
		s, err := ParseWAFAdvancedRuleSection(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFAdvancedRuleSectionRequestBody, s)
	})

	t.Run("ParsesRequestCookies", func(t *testing.T) {
		v := "REQUEST_COOKIES"
		s, err := ParseWAFAdvancedRuleSection(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFAdvancedRuleSectionRequestCookies, s)
	})

	t.Run("ParsesRequestHeaders", func(t *testing.T) {
		v := "REQUEST_HEADERS"
		s, err := ParseWAFAdvancedRuleSection(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFAdvancedRuleSectionRequestHeaders, s)
	})

	t.Run("ParsesRequestURI", func(t *testing.T) {
		v := "REQUEST_URI"
		s, err := ParseWAFAdvancedRuleSection(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFAdvancedRuleSectionRequestURI, s)
	})

	t.Run("MixedCase_Parses", func(t *testing.T) {
		v := "request_URI"
		s, err := ParseWAFAdvancedRuleSection(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFAdvancedRuleSectionRequestURI, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalidsection"
		_, err := ParseWAFAdvancedRuleSection(v)

		assert.NotNil(t, err)
		assert.Equal(t, "Invalid advanced rule section", err.Error())
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

	t.Run("ParsesContains", func(t *testing.T) {
		v := "contains"
		s, err := ParseWAFAdvancedRuleModifier(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFAdvancedRuleModifierContains, s)
	})

	t.Run("ParsesEndsWith", func(t *testing.T) {
		v := "endswith"
		s, err := ParseWAFAdvancedRuleModifier(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFAdvancedRuleModifierEndsWith, s)
	})

	t.Run("ParsesContainsWord", func(t *testing.T) {
		v := "containsword"
		s, err := ParseWAFAdvancedRuleModifier(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFAdvancedRuleModifierContainsWord, s)
	})

	t.Run("MixedCase_Parses", func(t *testing.T) {
		v := "Contains"
		s, err := ParseWAFAdvancedRuleModifier(v)

		assert.Nil(t, err)
		assert.Equal(t, WAFAdvancedRuleModifierContains, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalidmodifier"
		_, err := ParseWAFAdvancedRuleModifier(v)

		assert.NotNil(t, err)
		assert.Equal(t, "Invalid advanced rule modifier", err.Error())
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

	t.Run("ParsesDeny", func(t *testing.T) {
		v := "deny"
		s, err := ParseACLIPMode(v)

		assert.Nil(t, err)
		assert.Equal(t, ACLIPModeDeny, s)
	})

	t.Run("MixedCase_Parses", func(t *testing.T) {
		v := "Deny"
		s, err := ParseACLIPMode(v)

		assert.Nil(t, err)
		assert.Equal(t, ACLIPModeDeny, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalidmode"
		_, err := ParseACLIPMode(v)

		assert.NotNil(t, err)
		assert.Equal(t, "Invalid ACL IP mode", err.Error())
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

	t.Run("ParsesBlacklist", func(t *testing.T) {
		v := "blacklist"
		s, err := ParseACLGeoIPRulesMode(v)

		assert.Nil(t, err)
		assert.Equal(t, ACLGeoIPRulesModeBlacklist, s)
	})

	t.Run("MixedCase_Parses", func(t *testing.T) {
		v := "Blacklist"
		s, err := ParseACLGeoIPRulesMode(v)

		assert.Nil(t, err)
		assert.Equal(t, ACLGeoIPRulesModeBlacklist, s)
	})

	t.Run("Invalid_ReturnsError", func(t *testing.T) {
		v := "invalidmode"
		_, err := ParseACLGeoIPRulesMode(v)

		assert.NotNil(t, err)
		assert.Equal(t, "Invalid ACL GeoIP rules filtering mode", err.Error())
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

func TestNewCDNRuleCacheControlDurationFromDuration(t *testing.T) {
	d := time.Duration(time.Hour * 24 * 365)
	d = d + time.Duration(time.Hour*24*30*2)
	d = d + time.Duration(time.Hour*24*3)
	d = d + time.Duration(time.Hour*4)
	d = d + time.Duration(time.Minute*5)

	a := NewCDNRuleCacheControlDurationFromDuration(d)

	assert.Equal(t, 1, a.Years)
	assert.Equal(t, 2, a.Months)
	assert.Equal(t, 3, a.Days)
	assert.Equal(t, 4, a.Hours)
	assert.Equal(t, 5, a.Minutes)
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

	assert.Equal(t, "10276h5m0s", str)
}
