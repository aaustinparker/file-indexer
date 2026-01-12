package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/aaustinparker/file-indexer/index"
)

// TODO - uncomment and figure out how to inject configs in other files
var (
	port      = flag.String("port", ":8080", "Port to listen on")
	dataDir   = flag.String("dataDir", "data", "Source directory for files")
	indexName = flag.String("indexName", "index.bleve", "Name of the search index")
)

func main() {
	// parse command line flags
	flag.Parse()

	// create httpHandler with injected configs
	httpHandler := &index.HttpHandler{
		DataDir:   *dataDir,
		IndexName: *indexName,
	}

	// create index at startup
	httpHandler.CreateIndex()

	// create hooks for each server route
	http.HandleFunc("/", httpHandler.Home)
	http.HandleFunc("/search", httpHandler.SearchIndex)

	// start the server
	log.Printf("Listening on %s", *port)
	log.Fatal(http.ListenAndServe(*port, nil))
}
