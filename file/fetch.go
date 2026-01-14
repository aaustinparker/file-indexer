package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func Fetch(dataDir string, fileName string) (string, error) {
	directoryPath, _ := filepath.Abs(dataDir)
	filePath := filepath.Join(directoryPath, fileName)

	// make sure file exists in the data directory
	if _, err := os.Stat(filePath); err != nil {
		return "", fmt.Errorf("Can't open file at location %s: %v", filePath, err)
	}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("Error reading file at location %s: %v", filePath, err)
	}

	return string(fileContent), nil
}
