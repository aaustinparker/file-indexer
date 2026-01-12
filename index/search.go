package index

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/blevesearch/bleve/v2"
)

func Search(indexName string, searchTerm string) (string, error) {
	// get search term from query params
	if strings.TrimSpace(searchTerm) == "" {
		return "", fmt.Errorf("Query parameter 'q' is required")
	}

	// open the index
	index, err := bleve.Open(indexName)
	if err != nil {
		return "", fmt.Errorf("Error opening index %s: %v", indexName, err)
	}
	defer index.Close()

	// perform search
	query := bleve.NewWildcardQuery(fmt.Sprintf("*%s*", searchTerm))
	query.FieldVal = "Text"
	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.Fields = []string{"*"}
	searchResults, err := index.Search(searchRequest)
	if err != nil {
		return "", fmt.Errorf("Error searching index for results: %v", err)
	}

	parsedDocuments := parseDocument(searchResults)
	json, err := json.Marshal(parsedDocuments)
	if err != nil {
		return "", fmt.Errorf("JSON parsing error: %v", err)
	}

	return string(json), nil
}

func parseDocument(searchResults *bleve.SearchResult) []Document {
	parsedDocuments := []Document{}
	for _, result := range searchResults.Hits {
		doc := Document{
			FileName:   result.Fields["FileName"].(string),
			LineNumber: int(result.Fields["LineNumber"].(float64)),
			Text:       result.Fields["Text"].(string),
		}
		parsedDocuments = append(parsedDocuments, doc)
	}
	return parsedDocuments
}
