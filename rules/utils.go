package rules

import (
	"os"
	"path/filepath"
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

	// Convert both paths to absolute paths for proper comparison
	absFilePath := filepath.Join(cwd, filename)
	absFileDir := filepath.Dir(absFilePath)

	// Clean both paths to normalize them
	cleanCwd := filepath.Clean(cwd)
	cleanFileDir := filepath.Clean(absFileDir)

	// The file is in the current module if its directory matches or is the current working directory
	return cleanFileDir == cleanCwd
}
