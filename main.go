package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func serveJSONFiles(baseDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Remove leading slash and sanitize path
		relativePath := filepath.Clean(r.URL.Path[1:])

		// Prevent directory traversal attacks
		fullPath := filepath.Join(baseDir, relativePath)
		if !strings.HasPrefix(fullPath, baseDir) {
			http.Error(w, "Invalid path", http.StatusBadRequest)
			return
		}

		// Add .json extension if not present
		if filepath.Ext(fullPath) == "" {
			fullPath += ".json"
		}

		// Read file contents
		fileContent, err := os.ReadFile(fullPath)
		if err != nil {
			if os.IsNotExist(err) {
				http.NotFound(w, r)
				return
			}
			http.Error(w, "Error reading file: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Validate JSON
		var jsonData interface{}
		if err := json.Unmarshal(fileContent, &jsonData); err != nil {
			http.Error(w, "Invalid JSON file: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Set content type and write response
		w.Header().Set("Content-Type", "application/json")
		w.Write(fileContent)
	}
}

func main() {
	// Use current working directory if no argument provided
	baseDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Could not determine current directory:", err)
		os.Exit(1)
	}

	// Allow optional directory argument
	if len(os.Args) > 1 {
		baseDir = os.Args[1]
	}

	// Validate directory exists
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		fmt.Printf("Directory does not exist: %s\n", baseDir)
		os.Exit(1)
	}

	// Create handler
	http.HandleFunc("/", serveJSONFiles(baseDir))

	// Start server
	port := 9999
	fmt.Printf("Starting server on port %d, serving files from %s\n", port, baseDir)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
