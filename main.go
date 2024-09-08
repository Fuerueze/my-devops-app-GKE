package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s", r.URL.Path)
	fmt.Fprintf(w, "Hello, GKE, TEST,LOL,LOL")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
