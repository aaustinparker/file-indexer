package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/aaustinparker/file-indexer/index"
)

// TODO - uncomment and figure out how to inject configs in other files
var (
	port = flag.String("port", ":8080", "Port to listen on")
	// dataDir = flag.String("dataDir", "data", "Source directory for files")
	// indexName = flag.String("indexName", "example.bleve", "Name of the search index")
	// deleteExisting = flag.Bool("deleteExisting", true, "Whether to delete existing index on server startup")
)

func main() {
	// parse command line flags
	flag.Parse()

	// define handlers for each route
	http.HandleFunc("/create", index.CreateIndexHandler)
	http.HandleFunc("/search", index.SearchIndexHandler)

	// start the server
	log.Printf("Listening on %s", *port)
	log.Fatal(http.ListenAndServe(*port, nil))
}
