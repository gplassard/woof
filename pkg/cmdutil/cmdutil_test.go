package cmdutil

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrintJSON(t *testing.T) {
	tests := []struct {
		name       string
		payload    string
		filterType string
		expected   string
		factory    func() interface{}
	}{
		{
			name:       "No filter",
			payload:    "testdata/no_filter_payload.json",
			filterType: "",
			expected:   "testdata/no_filter_expected.json",
			factory:    func() interface{} { return &datadogV2.UsersResponse{} },
		},
		{
			name:       "Filter users",
			payload:    "testdata/filter_users_payload.json",
			filterType: "users",
			expected:   "testdata/filter_users_expected.json",
			factory:    func() interface{} { return &datadogV2.UsersResponse{} },
		},
		{
			name:       "Filter roles",
			payload:    "testdata/filter_roles_payload.json",
			filterType: "roles",
			expected:   "testdata/filter_roles_expected.json",
			factory:    func() interface{} { return &datadogV2.RolesResponse{} },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload := tt.factory()

			content, err := os.ReadFile(tt.payload)
			require.NoError(t, err, "failed to read payload file")
			err = json.Unmarshal(content, payload)
			require.NoError(t, err, "failed to unmarshal payload")

			content, err = os.ReadFile(tt.expected)
			require.NoError(t, err, "failed to read expected file")
			var expectedObj interface{}
			err = json.Unmarshal(content, &expectedObj)
			require.NoError(t, err, "failed to unmarshal expected")
			expectedBytes, err := json.MarshalIndent(expectedObj, "", "  ")
			require.NoError(t, err, "failed to marshal expected for normalization")

			output := FormatJSON(payload, tt.filterType)
			actualBytes := []byte(output)

			assert.Equal(t, string(expectedBytes), string(actualBytes), "FormatJSON() mismatch")
		})
	}
}
