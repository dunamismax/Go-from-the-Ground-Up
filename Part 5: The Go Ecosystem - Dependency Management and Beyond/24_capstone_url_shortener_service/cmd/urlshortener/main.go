package main

import (
	"log"
	"net/http"
	"os"

	// --- CORRECTED IMPORT PATHS ---
	// These paths now reflect the full module path defined in the root go.mod file.
	// This allows the Go toolchain to find our internal packages correctly.
	"github.com/dunamismax/Go-from-the-Ground-Up/Part5_The_Go_Ecosystem/24_capstone_url_shortener_service/internal/handler"
	"github.com/dunamismax/Go-from-the-Ground-Up/Part5_The_Go_Ecosystem/24_capstone_url_shortener_service/internal/store"
)

/*
=====================================================================================
|                               - CAPSTONE PROJECT -                                |
|                              URL Shortener Service                                |
=====================================================================================

This is main.go, the entry point for our entire application.

Its primary responsibility is "wiring up" the application:
1.  CONFIGURATION: Setting up server configuration (like the address).
2.  DEPENDENCY CREATION: Initializing all the core components (logger, data store, handlers).
3.  ROUTING: Mapping URL paths to their corresponding handler functions.
4.  SERVER STARTUP: Starting the HTTP server to listen for requests.

This separation of concerns—where `main` handles setup and other packages handle
the logic—is a hallmark of professional Go applications.
*/

// Config holds the application's configuration values.
type Config struct {
	Addr    string
	BaseURL string
}

func main() {
	// --- 1. Configuration ---
	cfg := Config{
		Addr:    ":8080",
		BaseURL: "http://localhost:8080",
	}

	// --- 2. Dependency Creation ---
	logger := log.New(os.Stdout, "[INFO] ", log.LstdFlags)
	urlStore := store.NewURLStore()
	h := handler.NewHandler(logger, urlStore, cfg.BaseURL)

	// --- 3. Routing ---
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.RedirectHandler)
	mux.HandleFunc("/api/shorten", h.ShortenURLHandler)

	// --- 4. Server Startup ---
	logger.Printf("Server starting on %s", cfg.BaseURL)
	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
