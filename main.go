package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/aaustinparker/file-indexer/index"
)

// TODO - uncomment and figure out how to inject configs in other files
var (
	port           = flag.String("port", ":8080", "Port to listen on")
	dataDir        = flag.String("dataDir", "data", "Source directory for files")
	indexName      = flag.String("indexName", "index.bleve", "Name of the search index")
	deleteExisting = flag.Bool("deleteExisting", true, "Whether to delete the existing index on server start")
)

func main() {
	// parse command line flags
	flag.Parse()

	// create httpHandler with injected dependencies
	httpHandler := &index.HttpHandler{
		DataDir:        *dataDir,
		IndexName:      *indexName,
		DeleteExisting: *deleteExisting,
	}

	// define handlers for each route
	http.HandleFunc("/create", httpHandler.CreateIndex)
	http.HandleFunc("/search", httpHandler.SearchIndex)

	// start the server
	log.Printf("Listening on %s", *port)
	log.Fatal(http.ListenAndServe(*port, nil))
}
