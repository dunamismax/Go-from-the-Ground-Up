// Part 4, Lesson 20: Project: Simple REST API
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file is the entry point for our multi-file REST API project.
// It initializes dependencies and starts the web server.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

PROJECT BRIEF: A COMPLETE REST API

In this enhanced project, we will build a complete and well-structured REST API.
This version improves upon our initial concept by:
 1. SEPARATING CONCERNS: We split our code into multiple files (`main.go`,
    `handlers.go`, `store.go`, `router.go`) to make it organized and maintainable.
 2. FULL CRUD: We implement all standard CRUD (Create, Read, Update, Delete)
    operations for our `/contacts` resource.
 3. CUSTOM ROUTER: We build a simple router from scratch to understand how
    incoming requests are matched to handler functions.
 4. STRUCTURED JSON RESPONSES: We create helper functions to ensure all API
    responses, including errors, are in a consistent JSON format.

This structure is much closer to what you would see in a professional Go application.
The `main.go` file now has a single responsibility: starting the application.
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// --- Part 1: Initialization ---

	// Create an instance of our contact store. This encapsulates all our
	// data and the mutex for safe concurrent access.
	store := NewContactStore()

	// Pre-populate our store with some initial data for demonstration.
	store.Create(Contact{Name: "Alice", Email: "alice@example.com", Phone: "111-111-1111"})
	store.Create(Contact{Name: "Bob", Email: "bob@example.com", Phone: "222-222-2222"})

	// Create a new instance of our custom router.
	router := NewRouter()

	// --- Part 2: Registering API Routes ---

	// We register our handlers for the various endpoints and HTTP methods.
	// This clear, declarative style makes it easy to see all available API routes.
	router.HandleFunc(http.MethodGet, "/contacts", store.handleGetContacts)
	router.HandleFunc(http.MethodPost, "/contacts", store.handleCreateContact)

	// For routes with an ID, our custom router knows how to handle them.
	router.HandleFunc(http.MethodGet, "/contacts/:id", store.handleGetContactByID)
	router.HandleFunc(http.MethodPut, "/contacts/:id", store.handleUpdateContact)
	router.HandleFunc(http.MethodDelete, "/contacts/:id", store.handleDeleteContact)

	// --- Part 3: Starting the Server ---

	port := ":8080"
	fmt.Printf("Server starting on port %s...\n", port)
	fmt.Println("See README.md for instructions on how to use the API.")

	// We pass our custom router as the handler for the server. The router's
	// ServeHTTP method will now be responsible for directing all incoming traffic.
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

HOW TO RUN THIS CODE:

1.  Open a terminal and navigate to this project's directory.
2.  Run the server using the `go run .` command, which compiles and runs all
    .go files in the directory.
    `go run .`
3.  Follow the instructions in `README.md` to interact with the API using `curl`.
*/
