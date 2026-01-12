package index

import (
	"fmt"
	"html"
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

	// search for some text
	query := bleve.NewWildcardQuery(fmt.Sprintf("*%s*", searchTerm))
	query.FieldVal = "Text"
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error searching index for results: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Found, %q", html.EscapeString(searchResults.String()))
}
