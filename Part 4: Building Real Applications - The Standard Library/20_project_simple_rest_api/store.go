package main

import (
	"fmt"
	"sync" // The sync package provides synchronization primitives, like mutexes.
)

// --- The Data Model ---

// Contact defines the structure of a single contact with JSON tags.
// These tags tell the `encoding/json` package how to map struct fields to JSON keys.
type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// --- The Data Storage Layer ---

// ContactStore encapsulates all data and logic for managing contacts.
// It holds the data map, a mutex for safe concurrent access, and the ID counter.
// This is a common pattern for creating a thread-safe in-memory "database".
type ContactStore struct {
	sync.Mutex    // Embed the mutex directly. This provides Lock() and Unlock() methods.
	contacts      map[int]Contact
	nextContactID int
}

// NewContactStore is a constructor function that initializes and returns a new ContactStore.
func NewContactStore() *ContactStore {
	return &ContactStore{
		contacts:      make(map[int]Contact),
		nextContactID: 1, // Start IDs at 1.
	}
}

// --- CRUD Methods ---

// Create adds a new contact to the store and returns the created contact with its new ID.
func (s *ContactStore) Create(contact Contact) Contact {
	s.Lock() // Lock the store to prevent concurrent writes.
	defer s.Unlock()

	contact.ID = s.nextContactID
	s.contacts[contact.ID] = contact
	s.nextContactID++
	return contact
}

// GetAll returns a slice of all contacts in the store.
func (s *ContactStore) GetAll() []Contact {
	s.Lock()
	defer s.Unlock()

	// It's good practice to return a slice rather than the map itself,
	// as it's a more common format for JSON list responses.
	allContacts := make([]Contact, 0, len(s.contacts))
	for _, contact := range s.contacts {
		allContacts = append(allContacts, contact)
	}
	return allContacts
}

// GetByID retrieves a single contact by its ID. It returns the contact and an error.
// Returning an error is idiomatic Go for handling cases where the item isn't found.
func (s *ContactStore) GetByID(id int) (Contact, error) {
	s.Lock()
	defer s.Unlock()

	contact, ok := s.contacts[id]
	if !ok {
		return Contact{}, fmt.Errorf("contact with id %d not found", id)
	}
	return contact, nil
}

// Update modifies an existing contact. It takes an ID and the new contact data.
func (s *ContactStore) Update(id int, updatedContact Contact) (Contact, error) {
	s.Lock()
	defer s.Unlock()

	_, ok := s.contacts[id]
	if !ok {
		return Contact{}, fmt.Errorf("contact with id %d not found", id)
	}

	updatedContact.ID = id // Ensure the ID remains the same.
	s.contacts[id] = updatedContact
	return updatedContact, nil
}

// Delete removes a contact from the store by its ID.
func (s *ContactStore) Delete(id int) error {
	s.Lock()
	defer s.Unlock()

	_, ok := s.contacts[id]
	if !ok {
		return fmt.Errorf("contact with id %d not found", id)
	}

	delete(s.contacts, id)
	return nil
}
