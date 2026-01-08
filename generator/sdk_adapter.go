package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// convertSDKToTemplateData converts an SDK operation to template data
func convertSDKToTemplateData(op *SDKOperation, config *Config) TemplateData {
	// Normalize bundle name from tag
	bundle, rawBundle := normalizeBundleFromTag(op)
	apiBundleName := normalizeApiBundleName(rawBundle)

	// Build the command use string
	use := toKebabCase(op.OperationID, config)
	var args []string
	var argTypes []string

	// Extract path parameters
	for _, param := range op.Parameters {
		if param.In == "path" && param.Required {
			// Convert camelCase parameter names to snake_case for CLI consistency
			paramNameSnakeCase := toSnakeCaseSimple(param.Name)
			args = append(args, paramNameSnakeCase)
			use += fmt.Sprintf(" [%s]", paramNameSnakeCase)

			// Convert Go type to the type format used in templates
			argType := convertGoTypeToArgType(param.Type, op.OperationID, paramNameSnakeCase)
			argTypes = append(argTypes, argType)
		}
	}

	// Determine request body info
	hasRequestBody := op.RequestBody != nil
	requestBodyType := ""
	isOptionalParams := false

	if hasRequestBody {
		// Extract the type name from the full type string (e.g., "CreateActionConnectionRequest")
		requestBodyType = extractTypeName(op.RequestBody.Type)

		// Check if this operation uses optional parameters
		if optType, ok := config.OptionalParametersOperations[op.OperationID]; ok {
			isOptionalParams = true
			requestBodyType = optType
		}
	}

	// Determine response info
	hasResponse := op.HasResponse
	responseTypeGo := ""
	if hasResponse && op.ResponseType != "" {
		responseTypeGo = extractTypeName(op.ResponseType)
	}
	if responseTypeGo == "" {
		responseTypeGo = "interface{}"
	}

	// Apply response type overrides
	responseTypeOverrides := map[string]string{
		"AttachmentArray": "IncidentAttachmentsResponse",
	}
	if override, ok := responseTypeOverrides[responseTypeGo]; ok {
		responseTypeGo = override
	}

	// Extract resource type from the operation name (best source of truth)
	// E.g., "RegisterAppKey" -> "app_key_registration" (Register + AppKey = registration)
	// E.g., "GetAppKeyRegistration" -> "app_key_registration"
	resourceType := extractResourceTypeFromOperation(op.OperationID, bundle)

	// Compute aliases
	aliases := computeAliases(bundle, op.OperationID, config)

	// Build documentation URL
	docBundle := strings.ReplaceAll(strings.ToLower(rawBundle), " ", "-")
	docURL := fmt.Sprintf("https://docs.datadoghq.com/api/latest/%s/#%s", docBundle, toKebabCase(op.OperationID, config))

	// Remove trailing period from summary for consistency
	summary := strings.TrimSuffix(op.Summary, ".")

	return TemplateData{
		PackageName:      bundle,
		CommandName:      strings.ToUpper(op.OperationID[:1]) + op.OperationID[1:] + "Cmd",
		Use:              use,
		Short:            summary,
		Method:           strings.ToUpper(op.Method),
		Path:             op.Path,
		OperationID:      op.OperationID,
		BundleName:       rawBundle,
		ApiBundleName:    apiBundleName,
		Args:             args,
		ArgTypes:         argTypes,
		HasRequestBody:   hasRequestBody,
		RequestBodyType:  requestBodyType,
		IsOptionalParams: isOptionalParams,
		HasResponse:      hasResponse,
		ResourceType:     resourceType,
		ResponseTypeGo:   responseTypeGo,
		Aliases:          aliases,
		DocURL:           docURL,
	}
}

// normalizeBundleFromTag extracts bundle names from operation tags
func normalizeBundleFromTag(op *SDKOperation) (string, string) {
	rawBundle := "general"
	if len(op.Tags) > 0 {
		rawBundle = op.Tags[0]
	}
	bundle := strings.ToLower(rawBundle)
	bundle = strings.ReplaceAll(bundle, " ", "_")
	bundle = strings.ReplaceAll(bundle, "-", "_")
	return bundle, rawBundle
}

// convertGoTypeToArgType converts Go SDK types to the argument types used in templates
func convertGoTypeToArgType(goType, operationID, paramName string) string {
	// Remove any package prefix (e.g., "datadogV2.SomeType" -> "SomeType")
	typeName := filepath.Base(goType)

	// Apply custom overrides first
	if operationID == "GetSBOM" && paramName == "asset_type" {
		return "datadogV2.AssetType"
	}
	if operationID == "GetTeamSync" && paramName == "filter[source]" {
		return "datadogV2.TeamSyncAttributesSource"
	}
	if (operationID == "GetCostByOrg" || operationID == "GetHistoricalCostByOrg" || operationID == "GetMonthlyCostAttribution") && paramName == "start_month" {
		return "time.Time"
	}
	if (operationID == "GetHourlyUsage" || operationID == "GetUsageApplicationSecurityMonitoring" || operationID == "GetUsageLambdaTracedInvocations" || operationID == "GetUsageObservabilityPipelines") && paramName == "start_hr" {
		return "time.Time"
	}

	// Basic type mapping
	switch typeName {
	case "string":
		return "string"
	case "int", "int32", "int64":
		return "int64"
	case "float32", "float64":
		return "float64"
	case "bool":
		return "bool"
	case "[]string":
		return "[]string"
	default:
		// For complex types, keep them as-is
		return goType
	}
}

// extractTypeName extracts the type name from a full Go type string
// E.g., "CreateActionConnectionRequest" -> "CreateActionConnectionRequest"
// E.g., "datadogV2.CreateActionConnectionRequest" -> "CreateActionConnectionRequest"
// E.g., "*CreateActionConnectionRequest" -> "CreateActionConnectionRequest"
func extractTypeName(typeStr string) string {
	// Remove pointer prefix
	typeStr = strings.TrimPrefix(typeStr, "*")

	// Remove package prefix if present
	parts := strings.Split(typeStr, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}

	return typeStr
}

// extractResourceTypeFromOperation extracts the resource type from an operation name
// E.g., "RegisterAppKey" -> "app_key_registration" (Register operation on AppKey)
// E.g., "GetAppKeyRegistration" -> "app_key_registration"
// E.g., "CreateActionConnection" -> "action_connection"
func extractResourceTypeFromOperation(operationID, fallback string) string {
	// Convert operation ID to snake_case to work with
	opSnake := toSnakeCaseSimple(operationID)

	// Remove common verb prefixes
	verbs := []string{
		"get_", "list_", "create_", "update_", "delete_", "search_", "query_",
		"register_", "unregister_", "activate_", "deactivate_",
		"enable_", "disable_", "validate_", "execute_", "run_", "cancel_", "retry_",
		"bulk_",
	}

	for _, verb := range verbs {
		if strings.HasPrefix(opSnake, verb) {
			resourceName := strings.TrimPrefix(opSnake, verb)

			// Strip plural suffix to get singular resource name
			resourceName = strings.TrimSuffix(resourceName, "s")

			// For register/unregister operations, if the resource name doesn't already
			// end with "_registration", append it
			if (verb == "register_" || verb == "unregister_") && !strings.HasSuffix(resourceName, "_registration") {
				return resourceName + "_registration"
			}

			// Otherwise return the resource name as-is
			if resourceName != "" {
				return resourceName
			}
		}
	}

	// If we couldn't extract a resource, use fallback
	return fallback
}

// extractResourceTypeFromResponseType extracts the resource type from a response type name
// E.g., "GetAppKeyRegistrationResponse" -> "app_key_registration"
// E.g., "CreateActionConnectionResponse" -> "action_connection"
// E.g., "ListIncidentsResponse" -> "incident"
// E.g., "RegisterAppKeyResponse" from "RegisterAppKey" op -> "app_key_registration"
func extractResourceTypeFromResponseType(responseType, operationID, fallback string) string {
	if responseType == "" || responseType == "interface{}" {
		return fallback
	}

	// Remove "Response" suffix first
	name := strings.TrimSuffix(responseType, "Response")

	// Try to remove CRUD prefixes
	crudPrefixes := []string{"Get", "List", "Create", "Update", "Delete", "Search", "Query", "Bulk"}
	for _, prefix := range crudPrefixes {
		if strings.HasPrefix(name, prefix) {
			name = strings.TrimPrefix(name, prefix)
			// Convert to snake_case
			resourceType := toSnakeCaseSimple(name)
			if resourceType != "" {
				return resourceType
			}
		}
	}

	// Try to remove action verb prefixes (Register, Activate, etc.)
	actionVerbs := []string{"Register", "Unregister", "Activate", "Deactivate",
		"Enable", "Disable", "Validate", "Execute", "Run", "Cancel", "Retry"}
	for _, verb := range actionVerbs {
		if strings.HasPrefix(name, verb) {
			resourceName := strings.TrimPrefix(name, verb)
			// Convert to snake_case
			resourceType := toSnakeCaseSimple(resourceName)

			// If we extracted a resource, try to find its full name from the operation ID
			// E.g., "RegisterAppKey" operation + "AppKey" extracted = "app_key_registration"
			// by looking for "AppKey" in the operation name
			if resourceType != "" {
				// Check if there's a more specific resource name in other operations
				// by looking at the operation ID for hints
				fullResourceType := findFullResourceName(resourceType, operationID)
				return fullResourceType
			}
		}
	}

	// If we couldn't extract anything, use fallback
	return fallback
}

// findFullResourceName attempts to find the full resource name when given a shortened version
// E.g., "app_key" from "RegisterAppKey" operation -> "app_key_registration"
// This works by inferring that if the operation is "RegisterAppKey", it's likely operating
// on an "AppKeyRegistration" resource (registration is a common pattern for verbs like Register)
func findFullResourceName(shortResourceType, operationID string) string {
	// For operations with action verbs (Register, Unregister, etc.),
	// the resource is often "<resource>_registration" or "<resource>_configuration"
	operationLower := strings.ToLower(operationID)

	// If it's a Register/Unregister operation, the resource is likely a "registration"
	if strings.HasPrefix(operationLower, "register") || strings.HasPrefix(operationLower, "unregister") {
		return shortResourceType + "_registration"
	}

	// If it's an Activate/Deactivate operation, the resource might stay as-is
	// or could be a "configuration"
	if strings.HasPrefix(operationLower, "activate") || strings.HasPrefix(operationLower, "deactivate") {
		// Keep as-is for now, can add "_configuration" if needed
		return shortResourceType
	}

	// For other operations, return as-is
	return shortResourceType
}
