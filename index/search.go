package index

import (
	"fmt"

	"github.com/blevesearch/bleve/v2"
)

func Search(indexName string, searchTerm string) ([]Document, error) {
	// open the index
	index, err := bleve.Open(indexName)
	if err != nil {
		return nil, fmt.Errorf("Error opening index %s: %v", indexName, err)
	}
	defer index.Close()

	// perform search
	query := bleve.NewTermQuery(searchTerm)
	query.FieldVal = "Text"
	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.Fields = []string{"*"}
	searchResults, err := index.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("Error searching index for results: %v", err)
	}

	return parseDocuments(searchResults), nil

}

func parseDocuments(searchResults *bleve.SearchResult) []Document {
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
