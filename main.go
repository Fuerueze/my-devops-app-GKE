package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request: ", r.URL.Path)
	fmt.Fprintf(w, "Hello, GKE, trigger test!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
