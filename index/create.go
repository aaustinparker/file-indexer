package index

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/blevesearch/bleve/v2"
)

func Create(indexName string, dataDir string) {
	// delete old index
	index, err := bleve.Open(indexName)
	if err == nil {
		log.Printf("Deleting old index")
		index.Close()
		os.RemoveAll(indexName)
	}

	// create new index
	mapping := bleve.NewIndexMapping()
	index, err = bleve.New(indexName, mapping)
	if err != nil {
		log.Printf("Error creating new index: %v", err)
		return
	}

	log.Printf("New index created at %s", indexName)
	defer index.Close()

	// iterate through files in the data directory
	files, err := os.ReadDir(dataDir)
	if err != nil {
		log.Printf("Error reading the data directory: %v", err)
		return
	}

	for _, file := range files {
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
	document := &Document{
		FileName:   fileName,
		LineNumber: lineNumber,
		Text:       text,
	}

	documentId := fmt.Sprintf("%s_%d", document.FileName, document.LineNumber)
	return index.Index(documentId, document)

}
