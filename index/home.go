package index

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func (h *HttpHandler) Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index/search.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to prepare page: %v", err), http.StatusInternalServerError)
		return
	}

	// pass some values to the HTML template
	directoryPath, _ := filepath.Abs(h.DataDir)
	templateValues := struct {
		DirectoryPath string
	}{
		DirectoryPath: directoryPath,
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, templateValues)
}
