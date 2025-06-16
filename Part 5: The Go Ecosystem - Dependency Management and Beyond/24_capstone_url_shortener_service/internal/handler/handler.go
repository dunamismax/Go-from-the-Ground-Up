package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"

	// --- CORRECTED IMPORT PATHS ---
	"github.com/dunamismax/Go-from-the-Ground-Up/Part5_The_Go_Ecosystem/24_capstone_url_shortener_service/internal/shortener"
	"github.com/dunamismax/Go-from-the-Ground-Up/Part5_The_Go_Ecosystem/24_capstone_url_shortener_service/internal/store"
)

/*
This is the handler package. It is responsible for the API layer of our application.
It contains the `http.HandlerFunc` implementations that parse requests, call the
necessary business logic (from the store and shortener packages), and write JSON or
redirect responses.
*/

// Handler is a struct that holds the dependencies for our HTTP handlers.
type Handler struct {
	logger  *log.Logger
	store   *store.URLStore
	baseURL string
}

// NewHandler is a constructor that creates a new Handler with its dependencies.
func NewHandler(logger *log.Logger, store *store.URLStore, baseURL string) *Handler {
	return &Handler{
		logger:  logger,
		store:   store,
		baseURL: baseURL,
	}
}

// ShortenURLRequest defines the expected structure of the JSON request body.
type ShortenURLRequest struct {
	URL string `json:"url"`
}

// ShortenURLResponse defines the structure of the JSON response body.
type ShortenURLResponse struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}

// ShortenURLHandler handles requests to create a new short URL.
func (h *Handler) ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if _, err := url.ParseRequestURI(req.URL); err != nil {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	if code, found := h.store.GetCodeForURL(req.URL); found {
		h.logger.Printf("Found existing code '%s' for URL '%s'", code, req.URL)
		h.respondWithJSON(w, http.StatusOK, ShortenURLResponse{
			OriginalURL: req.URL,
			ShortURL:    h.baseURL + "/" + code,
		})
		return
	}

	var code string
	for {
		code = shortener.GenerateShortCode()
		if _, found := h.store.Get(code); !found {
			break
		}
	}

	h.store.Set(code, req.URL)
	h.logger.Printf("Created new code '%s' for URL '%s'", code, req.URL)

	h.respondWithJSON(w, http.StatusCreated, ShortenURLResponse{
		OriginalURL: req.URL,
		ShortURL:    h.baseURL + "/" + code,
	})
}

// RedirectHandler handles redirecting a short URL to its original destination.
func (h *Handler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	code := strings.TrimPrefix(r.URL.Path, "/")
	if code == "" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Go URL Shortener! Use POST /api/shorten to create a link."))
		return
	}

	originalURL, found := h.store.Get(code)
	if !found {
		http.NotFound(w, r)
		return
	}

	h.logger.Printf("Redirecting code '%s' to '%s'", code, originalURL)
	http.Redirect(w, r, originalURL, http.StatusFound)
}

// respondWithJSON is a helper to write JSON responses.
func (h *Handler) respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
