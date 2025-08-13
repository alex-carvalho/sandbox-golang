package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func secureFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	path := filepath.Clean(r.URL.Path)
	if strings.Contains(path, "..") {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	
	if path == "/" {
		path = "/index.html"
	}
	
	http.ServeFile(w, r, filepath.Join("public", path))
}

func main() {
	http.HandleFunc("/", secureFileHandler)

	log.Println("Listening... port: 3000")
	// Use HTTPS in production: http.ListenAndServeTLS(":3000", "cert.pem", "key.pem", nil)
	log.Fatal(http.ListenAndServe(":3000", nil))

}
