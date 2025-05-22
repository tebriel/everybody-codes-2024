// Package utils wraps things we commonly do
package utils

import (
	"os"
	"strings"
)

// ReadFile reads the content of a file and returns it as a string.
func ReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}

	// convert bytes to string
	return strings.TrimSuffix(string(data), "\n"), nil
}
