// Package handler contains the HTTP handlers for our service.
package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"urlshortener/internal/store" // Import the internal store package
)

// URLStorer is an interface that defines the methods our handlers need from a store.
// Using an interface here makes our handlers more flexible and easier to test.
type URLStorer interface {
	Save(url string) (string, error)
	Get(shortCode string) (string, error)
}

// --- Request/Response Structs ---

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortCode string `json:"short_code"`
}

// RegisterRoutes sets up the HTTP handlers for the service.
// It takes a URLStorer (our `store.URLStore` will satisfy this interface).
func RegisterRoutes(mux *http.ServeMux, s URLStorer) {
	mux.HandleFunc("/shorten", handleShorten(s))
	mux.HandleFunc("/", handleRedirect(s))
}

// handleShorten is the handler for the POST /shorten endpoint.
func handleShorten(s URLStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req ShortenRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if req.URL == "" {
			http.Error(w, "URL must not be empty", http.StatusBadRequest)
			return
		}

		shortCode, err := s.Save(req.URL)
		if err != nil {
			http.Error(w, "Failed to create short code", http.StatusInternalServerError)
			return
		}

		resp := ShortenResponse{ShortCode: shortCode}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	}
}

// handleRedirect is the handler for the GET /{shortCode} endpoint.
func handleRedirect(s URLStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Trim the leading "/" from the request path to get the short code.
		shortCode := strings.TrimPrefix(r.URL.Path, "/")
		if shortCode == "" {
			http.NotFound(w, r)
			return
		}

		originalURL, err := s.Get(shortCode)
		if err != nil {
			if errors.Is(err, store.ErrNotFound) {
				http.Error(w, "Short code not found", http.StatusNotFound)
			} else {
				http.Error(w, "Server error", http.StatusInternalServerError)
			}
			return
		}

		// Perform the redirect.
		http.Redirect(w, r, originalURL, http.StatusFound)
	}
}
