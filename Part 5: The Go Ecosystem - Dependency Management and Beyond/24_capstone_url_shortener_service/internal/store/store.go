// Package store manages the data persistence for the URL shortener.
// For this capstone, we are using a simple in-memory map.
package store

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sync"
)

var (
	// ErrNotFound is returned when a short code is not found in the store.
	ErrNotFound = errors.New("short code not found")
)

// URLStore is a thread-safe in-memory key-value store for URLs.
type URLStore struct {
	mu   sync.Mutex
	urls map[string]string // a map from shortCode to originalURL
}

// NewURLStore creates and returns a new URLStore.
func NewURLStore() *URLStore {
	return &URLStore{
		urls: make(map[string]string),
	}
}

// Save saves a URL and returns a unique short code for it.
func (s *URLStore) Save(url string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Generate a short, random code.
	shortCode, err := generateShortCode(6)
	if err != nil {
		return "", err
	}

	s.urls[shortCode] = url
	return shortCode, nil
}

// Get retrieves the original URL for a given short code.
// It returns ErrNotFound if the code is not in the store.
func (s *URLStore) Get(shortCode string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	url, ok := s.urls[shortCode]
	if !ok {
		return "", ErrNotFound
	}
	return url, nil
}

// generateShortCode creates a random, hex-encoded string of a given length.
func generateShortCode(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
