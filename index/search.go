package index

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/blevesearch/bleve/v2"
)

func (h *HttpHandler) SearchIndex(w http.ResponseWriter, r *http.Request) {
	// get search term from query params
	searchTerm := r.URL.Query().Get("q")
	if strings.TrimSpace(searchTerm) == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	// open the index
	index, err := bleve.Open(h.IndexName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error opening index: %v", err), http.StatusInternalServerError)
		return
	}
	defer index.Close()

	// perform search
	query := bleve.NewWildcardQuery(fmt.Sprintf("*%s*", searchTerm))
	query.FieldVal = "Text"
	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.Fields = []string{"*"}
	searchResults, err := index.Search(searchRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error searching index for results: %v", err), http.StatusInternalServerError)
		return
	}

	parsedDocuments := parseDocument(searchResults)
	json, err := json.Marshal(parsedDocuments)
	if err != nil {
		http.Error(w, fmt.Sprintf("JSON error: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", string(json))
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
