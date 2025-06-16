// Part 4, Lesson 19: Intro to net/http
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file serves as the lesson and demonstration for building a web server.
// It explains the core concepts through structured comments and
// provides a runnable example of their implementation.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

Welcome to the world of web services in Go! One of Go's biggest strengths is
its `net/http` package, which provides all the tools you need to build a
fast, production-ready web server WITHOUT any external frameworks or libraries. [6]

THE CLIENT-SERVER MODEL
A web server's job is to listen for incoming HTTP REQUESTS from a CLIENT (like a
web browser or another program). When it receives a request, the server processes
it and sends back an HTTP RESPONSE.

In this lesson, we will build a very simple server that listens for requests and
responds with text. The two key components we will use are:
- `http.HandleFunc`: To tell the server which function should handle a request
  for a specific URL path.
- `http.ListenAndServe`: To start the server and make it listen for connections. [3]
*/

package main

import (
	"fmt"
	"log"
	"net/http" // The `net/http` package contains all we need for HTTP clients and servers. [18]
)

// --- Part 1: Defining Handler Functions ---

// A HANDLER is a function that receives HTTP requests and is responsible for
// writing a response. In Go, a handler function must have a specific signature. [10]
// It must accept two arguments:
//
// 1. `http.ResponseWriter`: An interface that our function uses to build and
//    send the response back to the client. We "write" our response to it. [1, 11]
//
// 2. `*http.Request`: A pointer to a struct that contains all the information
//    about the client's incoming request, such as the URL, headers, and any
//    data being sent. [1, 7]

// `rootHandler` will handle requests for the main "/" path of our server.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// We can inspect the request. Here, we'll just log the path that was requested.
	fmt.Printf("Received request for path: %s\n", r.URL.Path)

	// We use `fmt.Fprintf` to write a formatted string to the ResponseWriter.
	// This sends the text back to the client as the body of the HTTP response.
	fmt.Fprintf(w, "Welcome! You've reached the root of our simple Go web server.")
}

// `helloHandler` will handle requests for the "/hello" path.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request for path: %s\n", r.URL.Path)

	// A handler can do more than just send text. It can check for things.
	// For example, if the user tries to go to any sub-path of /hello/, we can
	// tell them the page doesn't exist.
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return // Important to return after sending a response.
	}

	fmt.Fprintf(w, "Hello, web! This is a different page.")
}

func main() {
	// --- Part 2: Registering Handlers and Starting the Server ---

	// Before we can start the server, we need to tell it which handler function
	// to use for which URL path. This is called ROUTING or "multiplexing".
	// The `http.HandleFunc` function registers a handler for a given path pattern. [1, 5]
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)

	// Now we define the address and port for our server to listen on.
	// ":8080" means the server will listen on port 8080 on all available
	// network interfaces of the machine.
	port := ":8080"
	fmt.Printf("Starting server on port %s\n", port)
	fmt.Println("Open your browser and navigate to http://localhost:8080")

	// `http.ListenAndServe` starts the server. [3, 7]
	// The first argument is the address/port. The second is the handler.
	// We pass `nil` for the handler to use the default server multiplexer,
	// which is the one we just registered our handlers on with `http.HandleFunc`. [3, 4]
	//
	// This function is BLOCKING. It will run forever, listening for requests,
	// and will only return if an unexpected error occurs (like the port
	// already being in use).
	err := http.ListenAndServe(port, nil)
	if err != nil {
		// Using `log.Fatal` is a common pattern here. If the server can't start,
		// it will print the error and exit the program.
		log.Fatal("ListenAndServe: ", err)
	}
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

HOW TO RUN THIS CODE:

1.  Open a terminal or command prompt.
2.  Navigate to the directory where you saved this file.
3.  Use the `go run` command to compile and execute the file:
    `go run 19_intro_to_net_http.go`

4.  Your terminal will show "Starting server on port :8080...". The program
    will appear to "hang" or pause. This is CORRECT! It's the server actively
    listening for connections.

5.  Open a web browser (like Chrome, Firefox, or Safari) and go to these URLs:
    -   `http://localhost:8080` (to see the rootHandler in action)
    -   `http://localhost:8080/hello` (to see the helloHandler in action)

6.  To stop the server, go back to your terminal and press `Ctrl+C`.
*/
