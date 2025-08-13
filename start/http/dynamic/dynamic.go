package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

var tmpl = template.Must(template.New("time").Parse(`<h1>The current time is: {{.}}</h1>`))

func currentTime(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	s := time.Now().Format("2006-01-02 15:04:05")
	tmpl.Execute(w, s)
}

func main() {
	http.HandleFunc("/", currentTime)
	log.Println("Listening... port: 3000")
	// Use HTTPS in production: http.ListenAndServeTLS(":3000", "cert.pem", "key.pem", nil)
	log.Fatal(http.ListenAndServe(":3000", nil))

}
