package index

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/blevesearch/bleve/v2"
)

func (h *HttpHandler) CreateIndex(w http.ResponseWriter, r *http.Request) {
	// delete old index if needed
	if h.DeleteExisting {
		if err := deleteIndex(h.IndexName); err != nil {
			log.Printf("No existing index to delete")
		}
	}

	// create a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(h.IndexName, mapping)
	if err != nil {
		log.Printf("Error creating new index: %v", err)
		return
	}

	defer index.Close()

	// iterate through files in the data directory
	dirEntries, err := os.ReadDir(h.DataDir)
	if err != nil {
		log.Printf("Error reading the data directory: %v", err)
		return
	}

	for _, file := range dirEntries {
		if file.Type().IsDir() || filepath.Ext(file.Name()) != ".txt" {
			continue
		}

		if err := indexFile(index, file.Name()); err != nil {
			log.Printf("%v", err)
		}
	}
}

func indexFile(index bleve.Index, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Failed to open file %s: %w", fileName, err)
	}
	defer file.Close()

	lineNumber := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if err := indexDocument(index, file.Name(), lineNumber, line); err != nil {
			return fmt.Errorf("Failed to index file %s at line %d: %w", fileName, lineNumber, err)
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Failed to parse file %s at line %d: %w", fileName, lineNumber, err)
	}

	return nil
}

func indexDocument(index bleve.Index, fileName string, lineNumber int, text string) error {
	document := struct {
		FileName   string
		LineNumber int
		Text       string
	}{
		FileName:   fileName,
		LineNumber: lineNumber,
		Text:       text,
	}

	documentId := fmt.Sprintf("%s_%d", document.FileName, document.LineNumber)
	return index.Index(documentId, document)

}

func deleteIndex(path string) error {
	idx, err := bleve.Open(path)
	if err != nil {
		return err
	}
	if err := idx.Close(); err != nil {
		return err
	}
	return os.RemoveAll(path)
}
