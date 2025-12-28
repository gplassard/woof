package main

import (
	"testing"
)

func TestToKebabCase(t *testing.T) {
	config, err := loadConfig()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	tests := []struct {
		input string
		want  string
	}{
		{"ListAPIs", "list-apis"},
		{"CreateCSMThreat", "create-csm-threat"},
		{"GetSLO", "get-slo"},
		{"GetSBOM", "get-sbom"},
		{"GetUC", "get-uc"},
		{"UpdateUser", "update-user"},
		{"Simple", "simple"},
		{"MultipleWords", "multiple-words"},
		{"Already-Kebab", "already-kebab"},
		{"CSMThreat", "csm-threat"},
		{"ThreatCSM", "threat-csm"},
		{"CSM", "csm"},
		{"SLOAndCSM", "slo-and-csm"},
		{"APIKey", "api-key"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := toKebabCase(tt.input, config)
			if got != tt.want {
				t.Errorf("toKebabCase(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestToSnakeCase(t *testing.T) {
	config, err := loadConfig()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	tests := []struct {
		input string
		want  string
	}{
		{"ListAPIs", "list_apis"},
		{"CreateCSMThreat", "create_csm_threat"},
		{"GetSLO", "get_slo"},
		{"GetSBOM", "get_sbom"},
		{"GetUC", "get_uc"},
		{"UpdateUser", "update_user"},
		{"Simple", "simple"},
		{"MultipleWords", "multiple_words"},
		{"Already_Snake", "already_snake"},
		{"CSMThreat", "csm_threat"},
		{"ThreatCSM", "threat_csm"},
		{"CSM", "csm"},
		{"SLOAndCSM", "slo_and_csm"},
		{"APIKey", "api_key"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := toSnakeCase(tt.input, config)
			if got != tt.want {
				t.Errorf("toSnakeCase(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestCasesWithoutConfig(t *testing.T) {
	tests := []struct {
		input string
		kebab string
		snake string
	}{
		{"ListAPIs", "list-a-p-is", "list_a_p_is"},
		{"CreateCSMThreat", "create-c-s-m-threat", "create_c_s_m_threat"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			gotKebab := toKebabCase(tt.input, nil)
			if gotKebab != tt.kebab {
				t.Errorf("toKebabCase(%q, nil) = %q, want %q", tt.input, gotKebab, tt.kebab)
			}
			gotSnake := toSnakeCase(tt.input, nil)
			if gotSnake != tt.snake {
				t.Errorf("toSnakeCase(%q, nil) = %q, want %q", tt.input, gotSnake, tt.snake)
			}
		})
	}
}
