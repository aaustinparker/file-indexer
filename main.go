package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/aaustinparker/file-indexer/file"
	"github.com/aaustinparker/file-indexer/index"
)

var (
	port      = flag.String("port", ":8080", "Port to listen on")
	dataDir   = flag.String("dataDir", "data", "Source directory for files")
	indexName = flag.String("indexName", "index.bleve", "Name of the search index")
)

func main() {
	// parse command line flags
	flag.Parse()

	// create index at startup
	index.Create(*indexName, *dataDir)

	// create handlers for each route
	http.HandleFunc("/", renderPage)
	http.HandleFunc("/search", searchIndex)
	http.HandleFunc("/file", fetchFile)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// start the server
	log.Printf("Listening on %s", *port)
	log.Fatal(http.ListenAndServe(*port, nil))
}

func searchIndex(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.URL.Query().Get("q")

	documents, err := index.Search(*indexName, searchTerm)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
	} else {
		json, err := json.Marshal(documents)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to generate JSON: %v", err), http.StatusInternalServerError)
		} else {
			fmt.Fprintf(w, "%s", string(json))
		}
	}
}

func renderPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/search.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to prepare page: %v", err), http.StatusInternalServerError)
		return
	}

	// pass some values to the HTML template
	directoryPath, _ := filepath.Abs(*dataDir)
	templateValues := struct {
		DirectoryPath string
	}{
		DirectoryPath: directoryPath,
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, templateValues)
}

func fetchFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("fileName")

	fileContent, err := file.Fetch(*dataDir, fileName)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, "%s", fileContent)
	}
}
