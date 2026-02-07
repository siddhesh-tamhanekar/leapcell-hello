package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Get port from environment, default to 8080
	port := "8080"

	// Simple handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	log.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
		log.Fatal(err)
	}
}
