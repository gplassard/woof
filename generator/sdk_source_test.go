package main

import "testing"

func TestParseSDKIncludesKnownOperation(t *testing.T) {
	config, err := loadConfig()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}
	input, err := loadFromSDK(config)
	if err != nil {
		t.Fatalf("failed to load from sdk: %v", err)
	}

	var found *OperationModel
	for i := range input.Operations {
		if input.Operations[i].OperationID == "DeleteActionConnection" {
			found = &input.Operations[i]
			break
		}
	}
	if found == nil {
		t.Fatalf("DeleteActionConnection was not extracted")
	}
	if found.Bundle != "action_connection" {
		t.Fatalf("unexpected bundle %q", found.Bundle)
	}
	if found.Method != "Delete" {
		t.Fatalf("unexpected method %q", found.Method)
	}
	if found.Path != "/api/v2/actions/connections/{connection_id}" {
		t.Fatalf("unexpected path %q", found.Path)
	}
	if len(found.Parameters) != 1 || found.Parameters[0].Name != "connection_id" {
		t.Fatalf("unexpected params %#v", found.Parameters)
	}
}
