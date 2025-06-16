package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// --- Helper Functions for Responses ---

// respondWithError is a helper to send a JSON error message with a specific status code.
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// respondWithJSON is a helper to marshal a payload to JSON and write it to the ResponseWriter.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		// If marshaling fails, it's a server-side problem.
		respondWithError(w, http.StatusInternalServerError, "Failed to encode response")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// --- Handler Functions ---
// These are methods on ContactStore so they have access to the data store.

// handleGetContacts processes GET requests to retrieve all contacts.
func (s *ContactStore) handleGetContacts(w http.ResponseWriter, r *http.Request) {
	contacts := s.GetAll()
	respondWithJSON(w, http.StatusOK, contacts)
}

// handleCreateContact processes POST requests to create a new contact.
func (s *ContactStore) handleCreateContact(w http.ResponseWriter, r *http.Request) {
	var newContact Contact
	err := json.NewDecoder(r.Body).Decode(&newContact)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Basic Validation
	if newContact.Name == "" || newContact.Email == "" {
		respondWithError(w, http.StatusBadRequest, "Name and Email are required fields")
		return
	}

	createdContact := s.Create(newContact)
	respondWithJSON(w, http.StatusCreated, createdContact)
}

// handleGetContactByID processes GET requests for a single contact.
func (s *ContactStore) handleGetContactByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromURL(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid contact ID")
		return
	}

	contact, err := s.GetByID(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, contact)
}

// handleUpdateContact processes PUT requests to update a contact.
func (s *ContactStore) handleUpdateContact(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromURL(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid contact ID")
		return
	}

	var updatedContact Contact
	err = json.NewDecoder(r.Body).Decode(&updatedContact)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if updatedContact.Name == "" || updatedContact.Email == "" {
		respondWithError(w, http.StatusBadRequest, "Name and Email are required fields")
		return
	}

	contact, err := s.Update(id, updatedContact)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, contact)
}

// handleDeleteContact processes DELETE requests.
func (s *ContactStore) handleDeleteContact(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromURL(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid contact ID")
		return
	}

	err = s.Delete(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	// A 204 No Content response is standard for a successful DELETE with no body.
	w.WriteHeader(http.StatusNoContent)
}

// getIDFromURL is a utility function to parse the integer ID from the request URL.
func getIDFromURL(r *http.Request) (int, error) {
	parts := strings.Split(r.URL.Path, "/")
	// Expected path: /contacts/{id}, so parts would be ["", "contacts", "id"]
	if len(parts) < 3 {
		return 0, new("missing contact ID")
	}
	return strconv.Atoi(parts[2])
}
