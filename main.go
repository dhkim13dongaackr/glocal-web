package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// VersionMetadata defines the structure for version info
type VersionMetadata struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Path        string `json:"path"` // Relative URL path to the file
	Snapshot    string `json:"snapshot"`
	Date        string `json:"date"`
}

func main() {
	port := "20080"

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Preview Dashboard
	http.HandleFunc("/preview", previewHandler)

	// Root redirect
	http.HandleFunc("/", rootHandler)

	fmt.Printf("Starting server on http://localhost:%s\n", port)
	fmt.Printf("View versions at http://localhost:%s/preview\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func getVersions() ([]VersionMetadata, error) {
	jsonFile, err := ioutil.ReadFile("./static/versions/metadata.json")
	if err != nil {
		return nil, err
	}

	var versions []VersionMetadata
	err = json.Unmarshal(jsonFile, &versions)
	if err != nil {
		return nil, err
	}
	return versions, nil
}

func previewHandler(w http.ResponseWriter, r *http.Request) {
	versions, err := getVersions()
	if err != nil {
		http.Error(w, "Could not load version metadata", http.StatusInternalServerError)
		return
	}

	// Parse the template file
	t, err := template.ParseFiles("templates/preview.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, versions)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	versions, err := getVersions()
	if err != nil || len(versions) == 0 {
		http.Redirect(w, r, "/preview", http.StatusFound)
		return
	}

	// Redirect to the first (latest) version in the list
	http.Redirect(w, r, versions[0].Path, http.StatusFound)
}
