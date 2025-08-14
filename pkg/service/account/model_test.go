package account

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContactType_String_Expected(t *testing.T) {
	ct := ContactTypeAccounts

	s := ct.String()

	assert.Equal(t, "Accounts", s)
}

func TestApplicationRestriction_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name        string
		jsonData    string
		expected    ApplicationRestriction
		expectError bool
	}{
		{
			name:     "Empty array returns empty restriction",
			jsonData: `[]`,
			expected: ApplicationRestriction{
				IPRestrictionType: "",
				IPRanges:          nil,
			},
			expectError: false,
		},
		{
			name:     "Valid object with allowlist unmarshals correctly",
			jsonData: `{"ip_restriction_type":"allowlist","ip_ranges":["198.51.100.1","198.51.100.2"]}`,
			expected: ApplicationRestriction{
				IPRestrictionType: "allowlist",
				IPRanges:          []string{"198.51.100.1", "198.51.100.2"},
			},
			expectError: false,
		},
		{
			name:     "Valid object with denylist unmarshals correctly",
			jsonData: `{"ip_restriction_type":"denylist","ip_ranges":["198.51.100.3"]}`,
			expected: ApplicationRestriction{
				IPRestrictionType: "denylist",
				IPRanges:          []string{"198.51.100.3"},
			},
			expectError: false,
		},
		{
			name:     "Empty object unmarshals correctly",
			jsonData: `{"ip_restriction_type":"","ip_ranges":[]}`,
			expected: ApplicationRestriction{
				IPRestrictionType: "",
				IPRanges:          []string{},
			},
			expectError: false,
		},
		{
			name:     "Object with null ip_ranges unmarshals correctly",
			jsonData: `{"ip_restriction_type":"allowlist","ip_ranges":null}`,
			expected: ApplicationRestriction{
				IPRestrictionType: "allowlist",
				IPRanges:          nil,
			},
			expectError: false,
		},
		{
			name:        "Invalid JSON returns error",
			jsonData:    `{invalid json}`,
			expected:    ApplicationRestriction{},
			expectError: true,
		},
		{
			name:     "Array with content treated as empty restriction",
			jsonData: `["some","content"]`,
			expected: ApplicationRestriction{
				IPRestrictionType: "",
				IPRanges:          nil,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var restriction ApplicationRestriction
			err := json.Unmarshal([]byte(tt.jsonData), &restriction)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.IPRestrictionType, restriction.IPRestrictionType)
				assert.Equal(t, tt.expected.IPRanges, restriction.IPRanges)
			}
		})
	}
}
