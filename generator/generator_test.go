package main

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestGenerationFromSDK(t *testing.T) {
	origDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	_, filename, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(filename), "..")
	err = os.Chdir(root)
	if err != nil {
		t.Fatalf("Failed to change directory to root: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(origDir)
	})

	if err := RunGenerate(nil); err != nil {
		t.Fatalf("Failed to run generator: %v", err)
	}

	mustExist := []string{
		"cmd/root.gen.go",
		"cmd/action_connection/entrypoint.gen.go",
		"cmd/action_connection/delete_action_connection.gen.go",
	}
	for _, file := range mustExist {
		if _, err := os.Stat(file); err != nil {
			t.Fatalf("expected generated file %s: %v", file, err)
		}
	}
}
