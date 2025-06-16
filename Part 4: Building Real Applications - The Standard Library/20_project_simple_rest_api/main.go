// Part 4, Lesson 20: Project: Simple REST API
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file is a multi-file project. It creates a simple REST API
// for a contact book, combining our knowledge of net/http, JSON,
// structs, and maps.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

PROJECT BRIEF: SIMPLE REST API

In this project, we will build a simple but functional REST API (Representational
State Transfer API). A REST API is an architectural style for web services that
uses standard HTTP methods (like GET, POST, PUT, DELETE) to allow clients to
interact with resources.

Our API will manage a "contact book" and will expose one resource: `/contacts`.
We will implement two HTTP methods for this resource:
- `GET /contacts`: To retrieve a list of all contacts.
- `POST /contacts`: To add a new contact to the list.

To do this, we'll store our data in memory using a map, and we will introduce
a `sync.Mutex` to make our data access safe for concurrent requests.
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync" // The sync package provides synchronization primitives, like mutexes.
)

// --- Part 1: The Data Model and Storage ---

// Contact defines the structure of a single contact with JSON tags.
type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// We'll use an in-memory database (a map) to store our contacts.
// The key will be the contact's ID.
var contactStore = make(map[int]Contact)

// This variable will act as a simple auto-incrementing primary key.
var nextContactID = 1

// A Mutex is a MUTUAL EXCLUSION lock. We need it because multiple requests
// (goroutines) could try to access `contactStore` at the same time, leading to
// a RACE CONDITION. The mutex ensures that only one goroutine can access the
// map at any given time.
var storeMutex = &sync.Mutex{}

// --- Part 2: The API Handler ---

// contactsHandler will handle all requests to the `/contacts` endpoint.
func contactsHandler(w http.ResponseWriter, r *http.Request) {
	// We use a `switch` statement on the request's HTTP method.
	switch r.Method {
	case http.MethodGet:
		handleGetContacts(w, r)
	case http.MethodPost:
		handlePostContact(w, r)
	default:
		// If any other method is used, we send a "Method Not Allowed" response.
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleGetContacts processes GET requests to retrieve all contacts.
func handleGetContacts(w http.ResponseWriter, r *http.Request) {
	// Lock the mutex to ensure safe reading of the map. `defer` ensures
	// Unlock() is called right before the function returns.
	storeMutex.Lock()
	defer storeMutex.Unlock()

	// A JSON array is more common for lists than a JSON object.
	// We'll convert our map of contacts into a slice for marshaling.
	contacts := make([]Contact, 0, len(contactStore))
	for _, contact := range contactStore {
		contacts = append(contacts, contact)
	}

	// Set the Content-Type header so the client knows to expect JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts) // Encode the slice directly to the ResponseWriter.
}

// handlePostContact processes POST requests to create a new contact.
func handlePostContact(w http.ResponseWriter, r *http.Request) {
	var newContact Contact
	// Use json.NewDecoder to read the request body and decode the JSON into our struct.
	// This is more efficient than reading the whole body into memory first.
	err := json.NewDecoder(r.Body).Decode(&newContact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Lock the mutex for safe writing.
	storeMutex.Lock()
	defer storeMutex.Unlock()

	// Assign a new ID and store the contact.
	newContact.ID = nextContactID
	contactStore[newContact.ID] = newContact
	nextContactID++

	// Respond to the client with the newly created contact (including its new ID).
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // Set the status code to 201 Created.
	json.NewEncoder(w).Encode(newContact)
}

func main() {
	// --- Part 3: Initializing Data and Starting the Server ---

	// Pre-populate our store with some initial data.
	contactStore[1] = Contact{ID: 1, Name: "Alice", Email: "alice@example.com", Phone: "111-111-1111"}
	contactStore[2] = Contact{ID: 2, Name: "Bob", Email: "bob@example.com", Phone: "222-222-2222"}
	nextContactID = 3

	// Register our handler for the `/contacts` route.
	http.HandleFunc("/contacts", contactsHandler)

	fmt.Println("Server starting on port :8080...")
	fmt.Println("See README.md for instructions on how to use the API.")

	// Start the server.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

HOW TO RUN THIS CODE:

This is a multi-file project. Please see the `README.md` file in this directory
for detailed instructions on how to run the server and interact with it using
tools like `curl`.
*/
