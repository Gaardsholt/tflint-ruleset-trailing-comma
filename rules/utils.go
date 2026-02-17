package rules

import (
	"os"
	"path/filepath"
	"strings"
)

func isWhitespace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}

// isFileInCurrentModule checks if the file belongs to the current module context.
func isFileInCurrentModule(filename string) bool {
	dir := filepath.Dir(filename)
	if dir == "." {
		return true
	}

	cwd, err := os.Getwd()
	if err != nil {
		return true
	}

	return strings.HasSuffix(cwd, dir)
}
