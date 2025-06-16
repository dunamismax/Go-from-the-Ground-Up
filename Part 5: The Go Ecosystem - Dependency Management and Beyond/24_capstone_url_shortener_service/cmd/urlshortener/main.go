// The main package for the URL shortener executable.
package main

import (
	"fmt"
	"log"
	"net/http"
	"urlshortener/internal/handler" // Import our internal packages
	"urlshortener/internal/store"
)

func main() {
	// 1. Create an instance of our in-memory store.
	urlStore := store.NewURLStore()

	// 2. Create a new ServeMux (HTTP request router).
	mux := http.NewServeMux()

	// 3. Register our routes and handlers, injecting the store.
	handler.RegisterRoutes(mux, urlStore)

	port := ":8080"
	fmt.Printf("Starting URL shortener service on http://localhost%s\n", port)

	// 4. Start the HTTP server.
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
