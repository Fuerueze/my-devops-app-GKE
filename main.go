package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("Starting server...")
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	platform := os.Getenv("PLATFORM")
	if platform == "" {
		platform = "Unknown"
	}

	fmt.Fprintf(
		w,
		"Hello! This application is running on %s\n",
		platform,
	)
}
