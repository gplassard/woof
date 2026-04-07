# Generator Migration Notes (Spec -> SDK)

## What changed

- Removed OpenAPI-spec-driven generation.
- Generation now parses `datadogV2` Go SDK source directly.
- Removed legacy spec download/version wiring.

## Compatibility behavior

- Existing command templates and output structure are preserved.
- Existing skip/optional-operation config is preserved.
- Added bundle compatibility mapping for service names that differ from historical tag-derived bundle names:
  - `APM` -> `apm`
  - `APMRetentionFilters` -> `apm_retention_filters`
  - `AuthNMappings` -> `authn_mappings`
  - `HighAvailabilityMultiRegion` -> `high_availability_multiregion`
  - `OCIIntegration` -> `oci_integration`
  - `ServiceNowIntegration` -> `servicenow_integration`

## Intentional/accepted differences

- Minor naming drift can still happen where SDK metadata and previous tag-based naming diverge.
- Summaries come from SDK method comments (first line) instead of OpenAPI summary fields.
- Resource type inference now comes from SDK structs/enums traversal.

## Validation

- `go run ./generator` succeeds without OpenAPI input.
- `go test ./generator/...` passes.
- `go test ./...` passes with generated output.
