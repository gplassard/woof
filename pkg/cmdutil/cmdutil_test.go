package cmdutil

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/spf13/cobra"
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

func TestUnmarshalPayload(t *testing.T) {
	type TestBody struct {
		Name string `json:"name"`
	}

	tests := []struct {
		name         string
		setupFlags   func() *cobra.Command
		setupFile    func() (string, func())
		expectedBody TestBody
		expectError  bool
	}{
		{
			name: "JSON input from --payload",
			setupFlags: func() *cobra.Command {
				cmd := &cobra.Command{}
				cmd.Flags().String("payload", `{"name": "arg-json"}`, "")
				cmd.Flags().String("payload-file", "", "")
				return cmd
			},
			expectedBody: TestBody{Name: "arg-json"},
		},
		{
			name: "Input from --payload-file flag",
			setupFlags: func() *cobra.Command {
				cmd := &cobra.Command{}
				cmd.Flags().String("payload", "", "")
				cmd.Flags().String("payload-file", "test_UnmarshalPayload.json", "")
				return cmd
			},
			setupFile: func() (string, func()) {
				err := os.WriteFile("test_UnmarshalPayload.json", []byte(`{"name": "file-json"}`), 0644)
				if err != nil {
					t.Fatal(err)
				}
				return "test_UnmarshalPayload.json", func() {
					os.Remove("test_UnmarshalPayload.json")
				}
			},
			expectedBody: TestBody{Name: "file-json"},
		},
		{
			name: "Missing payload error",
			setupFlags: func() *cobra.Command {
				cmd := &cobra.Command{}
				cmd.Flags().String("payload", "", "")
				cmd.Flags().String("payload-file", "", "")
				return cmd
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := tt.setupFlags()
			if tt.setupFile != nil {
				_, cleanup := tt.setupFile()
				defer cleanup()
			}

			var body TestBody
			err := UnmarshalPayload(cmd, &body)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, body)
			}
		})
	}
}
