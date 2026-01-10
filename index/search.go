package index

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/blevesearch/bleve/v2"
)

func (h *HttpHandler) SearchIndex(w http.ResponseWriter, r *http.Request) {
	// open the index
	index, err := bleve.Open(h.IndexName)
	if err != nil {
		log.Printf("Error opening index: %v", err)
		return
	}
	defer index.Close()

	// search for some text
	query := bleve.NewWildcardQuery("*random*")
	query.FieldVal = "Text"
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		log.Printf("Error searching index for results: %v", err)
		return
	}

	fmt.Fprintf(w, "Found, %q", html.EscapeString(searchResults.String()))
}
