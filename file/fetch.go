package file

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const matchedLineBuffer int = 5

func Fetch(dataDir string, fileName string, lineNumber int) (string, error) {
	directoryPath, _ := filepath.Abs(dataDir)
	filePath := filepath.Join(directoryPath, fileName)

	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("Can't open file at location %s: %v", filePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	startLine := max(1, lineNumber-matchedLineBuffer)
	endLine := lineNumber + matchedLineBuffer

	currentLine := 1
	var matchedLines strings.Builder
	for scanner.Scan() && currentLine <= endLine {
		if currentLine >= startLine {
			lineText := scanner.Text()
			fmt.Fprintf(&matchedLines, "%d.\t%s\n", currentLine, lineText)
		}
		currentLine++
	}

	return matchedLines.String(), nil
}
