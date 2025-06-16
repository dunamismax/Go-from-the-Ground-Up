package store

import "sync"

/*
This is the store package. It is responsible for all data persistence logic.
By keeping data access logic separate, we can easily swap out the storage
backend in the future (e.g., from in-memory to a database like Redis or
PostgreSQL) without changing our HTTP handlers.
*/

// URLStore holds the data for our URL shortener and provides safe concurrent access.
type URLStore struct {
	// We use an RWMutex (Read-Write Mutex). It allows multiple "readers" (redirects)
	// to access the data at the same time, but only one "writer" (creating a new link).
	// This is a performance optimization since our service will have many more reads than writes.
	mu sync.RWMutex

	// urls maps a short code (e.g., "aB3dC") to its original, long URL.
	urls map[string]string

	// codes maps an original, long URL to its already-generated short code.
	// This makes our creation endpoint IDEMPOTENT: creating a short link for the
	// same long URL twice will return the same short code.
	codes map[string]string
}

// NewURLStore is a constructor function that creates and returns a new, initialized URLStore.
func NewURLStore() *URLStore {
	return &URLStore{
		urls:  make(map[string]string),
		codes: make(map[string]string),
	}
}

// Get retrieves the original URL for a given short code.
// It returns the URL and a boolean indicating if the code was found.
func (s *URLStore) Get(code string) (string, bool) {
	s.mu.RLock() // Acquire a read lock. Multiple goroutines can hold a read lock.
	defer s.mu.RUnlock()
	url, found := s.urls[code]
	return url, found
}

// Set saves a new short code and its corresponding original URL to the store.
func (s *URLStore) Set(code, url string) {
	s.mu.Lock() // Acquire a write lock. Only one goroutine can hold a write lock.
	defer s.mu.Unlock()
	s.urls[code] = url
	s.codes[url] = code
}

// GetCodeForURL checks if a short code already exists for a given original URL.
func (s *URLStore) GetCodeForURL(url string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	code, found := s.codes[url]
	return code, found
}
