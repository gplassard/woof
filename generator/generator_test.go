package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

func TestGenerationNoChanges(t *testing.T) {
	// 1. Save original working directory and change to project root
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

	// 2. Run the generator
	if err := RunGenerate(nil); err != nil {
		t.Fatalf("Failed to run generator: %v", err)
	}

	// 3. Check if git reports any changes in the cmd directory
	gitStatus := exec.Command("git", "status", "--porcelain", "cmd/")
	output, err := gitStatus.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run git status: %v\nOutput: %s", err, string(output))
	}

	if len(output) > 0 {
		t.Errorf("Generated files in cmd/ have changed:\n%s", string(output))
	}
}
