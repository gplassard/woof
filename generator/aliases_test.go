package main

import (
	"reflect"
	"testing"
)

func TestComputeAliases(t *testing.T) {
	config, err := loadConfig()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	tests := []struct {
		name        string
		bundle      string
		operationID string
		want        []string
	}{
		{
			name:        "bulk-delete-datastore-items",
			bundle:      "actions_datastores",
			operationID: "BulkDeleteDatastoreItems",
			want:        []string{"bulk-delete-items"},
		},
		{
			name:        "create-action-connection",
			bundle:      "action_connection",
			operationID: "CreateActionConnection",
			want:        []string{"create"},
		},
		{
			name:        "list-spans-get",
			bundle:      "spans",
			operationID: "ListSpansGet",
			want:        []string{"list-get"},
		},
		{
			name:        "list-powerpacks",
			bundle:      "powerpack",
			operationID: "ListPowerpacks",
			want:        []string{"list"},
		},
		{
			name:        "create-datastore",
			bundle:      "actions_datastores",
			operationID: "CreateDatastore",
			want:        []string{"create"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := computeAliases(tt.bundle, tt.operationID, config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("computeAliases() = %v, want %v", got, tt.want)
			}
		})
	}
}
