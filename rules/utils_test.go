package rules

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsWhitespace(t *testing.T) {
	tests := []struct {
		input    byte
		expected bool
	}{
		{' ', true},
		{'\t', true},
		{'\n', true},
		{'\r', true},
		{'a', false},
		{'1', false},
		{'.', false},
	}

	for _, test := range tests {
		if result := isWhitespace(test.input); result != test.expected {
			t.Errorf("isWhitespace(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsFileInCurrentModule(t *testing.T) {
	// Create a temporary directory structure for testing module contexts
	// structure:
	// root/
	//   main.tf
	//   modules/
	//     child/
	//       main.tf

	tmpDir, err := os.MkdirTemp("", "tflint-test-root")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir) //nolint:errcheck // We don't care about this error in our test

	modulesDir := filepath.Join(tmpDir, "modules", "child")
	if err := os.MkdirAll(modulesDir, 0755); err != nil {
		t.Fatal(err)
	}

	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(originalWd) //nolint:errcheck // We don't care about this error in our test

	// Test Case 1: Running in Root Module
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}

	// 1a. File in root should be IN (True)
	if !isFileInCurrentModule("main.tf") {
		t.Errorf("Root: main.tf should be in current module")
	}

	// 1b. File in submodule should be OUT (False) - THIS CATCHES THE BUG
	// Note: TFLint passes relative paths usually.
	if isFileInCurrentModule("modules/child/main.tf") {
		t.Errorf("Root: modules/child/main.tf should NOT be in current module")
	}

	// Test Case 2: Running in Submodule
	if err := os.Chdir(modulesDir); err != nil {
		t.Fatal(err)
	}

	// 2a. File in submodule (relative to root execution?)
	// When TFLint runs recursively, it changes CWD to the module dir.
	// But filenames might still be relative to project root in some contexts?
	// If existing logic assumes relative paths are relative to CWD:
	if !isFileInCurrentModule("main.tf") {
		t.Errorf("Submodule: main.tf should be in current module")
	}

	// 2b. File from parent (root) should be OUT (False)
	if isFileInCurrentModule("../../main.tf") {
		t.Errorf("Submodule: ../../main.tf should NOT be in current module")
	}
}
