package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func resolvePath(path string) string {
	if path == "" {
		return path
	}

	// Expand environment variables and clean the path
	expPath := os.ExpandEnv(path)
	cleanPath := filepath.Clean(expPath)

	// Resolve the absolute path
	absPath, err := filepath.Abs(cleanPath)
	if err != nil {
		fmt.Printf("Error expanding path: %v\n", err)
		os.Exit(1)
	}

	return absPath
}
