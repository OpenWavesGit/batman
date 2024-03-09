package fileReader

import (
	"fmt"
	"os"
)

func ReadFile(filePath string) (string, error) {
	// Read the contents of the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(data), nil
}
