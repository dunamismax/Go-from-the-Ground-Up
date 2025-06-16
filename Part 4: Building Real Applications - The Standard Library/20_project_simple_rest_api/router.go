package main

import (
	"net/http"
	"strings"
)

// Router is a simple HTTP router. It maps HTTP methods and URL paths to handler functions.
// This teaches the core concept of routing without the complexity of a third-party library.
type Router struct {
	// The map keys are HTTP methods (e.g., "GET"), and the values are another map
	// where keys are URL paths and values are the handlers.
	rules map[string]map[string]http.HandlerFunc
}

// NewRouter creates and returns a new Router instance.
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

// HandleFunc registers a new handler for a given method and path.
func (r *Router) HandleFunc(method, path string, handler http.HandlerFunc) {
	// Initialize the inner map if it doesn't exist for this method.
	if _, ok := r.rules[method]; !ok {
		r.rules[method] = make(map[string]http.HandlerFunc)
	}
	r.rules[method][path] = handler
}

// FindHandler tries to find a handler for the given method and path.
// It supports simple path parameters like `/contacts/:id`.
func (r *Router) FindHandler(method, path string) (http.HandlerFunc, bool) {
	// First, try for an exact match.
	pathMap, ok := r.rules[method]
	if !ok {
		return nil, false
	}

	handler, ok := pathMap[path]
	if ok {
		return handler, true
	}

	// If no exact match, check for paths with parameters (e.g., "/contacts/:id").
	// This is a simplified approach for demonstration.
	for registeredPath, registeredHandler := range pathMap {
		if strings.Contains(registeredPath, ":") {
			// Split both paths by '/'
			pathParts := strings.Split(path, "/")
			registeredPathParts := strings.Split(registeredPath, "/")

			if len(pathParts) == len(registeredPathParts) {
				match := true
				for i, part := range registeredPathParts {
					if strings.HasPrefix(part, ":") {
						continue // It's a parameter, so it's a match for this part.
					}
					if part != pathParts[i] {
						match = false
						break
					}
				}
				if match {
					return registeredHandler, true
				}
			}
		}
	}

	return nil, false
}

// ServeHTTP makes our Router satisfy the http.Handler interface.
// This method is called for every incoming HTTP request.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, found := r.FindHandler(req.Method, req.URL.Path)

	if !found {
		// If no handler is found, respond with a 404 Not Found.
		http.NotFound(w, req)
		return
	}

	// If a handler is found, call it.
	handler(w, req)
}
