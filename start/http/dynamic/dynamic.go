package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func currnentTime(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "<h1>The current time is: %s</h1>", s)
}

func main() {
	http.HandleFunc("/", currnentTime)
	log.Println("Listening... port: 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
