// Part 4, Lesson 18: Handling JSON
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file serves as the lesson and demonstration for handling JSON.
// It explains the core concepts through structured comments and
// provides a runnable example of their implementation.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

WHAT IS JSON?
JSON (JavaScript Object Notation) is a lightweight, text-based format for sharing
data. [2, 4, 7, 8] It is the most common format used by APIs and web services to send and
receive information. Because it is simple for both humans to read and machines
to parse, it has become the universal language for data exchange on the internet. [2, 10]

Go has fantastic, built-in support for JSON in its `encoding/json` package. [3]
There are two main operations we will learn:

1.  MARSHALING: The process of converting a Go data structure (like a struct or a
    map) INTO a JSON string. [5] Think of it as "packaging up" your Go data to
    send over the network.

2.  UNMARSHALING: The process of converting a JSON string INTO a Go data
    structure. This is "un-packaging" the data you receive so you can work
    with it in your Go program. [1, 5]
*/

// The `encoding/json` package contains all the functions we need.
package main

import (
	"encoding/json"
	"fmt"
)

// --- Part 1: Defining a Go Struct with JSON Tags ---

// When working with JSON, we typically want to map the JSON data directly to
// a Go `struct`. This gives us the benefit of Go's static typing.
//
// To control how the `encoding/json` package handles our struct fields, we use
// STRUCT TAGS. A struct tag is a string of metadata attached to a field. [15, 19]
//
// For JSON, the tag `json:"fieldName"` tells the package what the corresponding
// key name should be in the JSON string. [3, 5] This is how we can have idiomatic Go
// field names (e.g., `UserName`, capitalized and exported) while using idiomatic
// JSON key names (e.g., `userName`, camelCase).
type User struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	IsActive bool     `json:"isActive"`
	Courses  []string `json:"courses"`

	// This field has the tag `omitempty`. This means if the field has its
	// zero value (0 for an int), it will be completely left out of the
	// JSON output. This is useful for optional fields. [9, 16]
	LoginAttempts int `json:"loginAttempts,omitempty"`

	// This field has a `-` tag. This tells the JSON package to always
	// ignore this field. It will never be marshaled or unmarshaled. [9, 16]
	// This is perfect for internal data that should not be exposed.
	internalSecret string `json:"-"`
}

func main() {
	// --- Part 2: Marshaling (Go Struct to JSON) ---

	// First, let's create an instance of our User struct.
	user1 := User{
		Name:          "Alice",
		Age:           30,
		IsActive:      true,
		Courses:       []string{"Go Foundations", "REST APIs in Go"},
		LoginAttempts: 2, // This will be included because it's not the zero value.
	}

	// We'll create another user, but leave LoginAttempts as 0.
	user2 := User{
		Name:     "Bob",
		Age:      25,
		IsActive: false,
		Courses:  []string{"Go Concurrency"},
		// LoginAttempts is 0 (the zero value for int), so it will be omitted.
	}

	// The `json.Marshal` function takes our Go data structure and returns
	// a byte slice (`[]byte`) representing the JSON, and an error. [9]
	// NOTE: Only EXPORTED fields (starting with a capital letter) are encoded. [3]
	fmt.Println("--- Marshaling a single user ---")
	user1Json, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	// We convert the byte slice to a string to print it.
	fmt.Println(string(user1Json))

	// The output from `json.Marshal` is compact. For human-readable output,
	// it's better to use `json.MarshalIndent`. It adds newlines and indentation. [14]
	fmt.Println("\n--- Marshaling with Indentation ---")
	users := []User{user1, user2}
	prettyJson, err := json.MarshalIndent(users, "", "  ") // "" prefix, "  " for indentation
	if err != nil {
		fmt.Println("Error marshaling indented JSON:", err)
		return
	}
	fmt.Println(string(prettyJson))

	// --- Part 3: Unmarshaling (JSON to Go Struct) ---

	// Now for the reverse. Let's take a JSON string and turn it into a Go struct.
	// We use a raw string literal (``) to avoid having to escape the double quotes.
	jsonString := `
	{
		"name": "Charlie",
		"age": 42,
		"isActive": true,
		"courses": ["Advanced Go", "gRPC Fundamentals"]
	}`

	// First, we need a variable of the target type to hold the decoded data.
	var charlie User

	// `json.Unmarshal` takes the JSON data as a byte slice and a POINTER to the
	// variable where the data should be stored. [1, 5, 13]
	// It needs a pointer so it can modify the `charlie` variable directly.
	err = json.Unmarshal([]byte(jsonString), &charlie)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println("\n--- Unmarshaling into a Go Struct ---")
	// The `%+v` format verb is great for printing structs, as it includes field names.
	fmt.Printf("Unmarshaled user: %+v\n", charlie)
	fmt.Printf("Charlie's first course is: %s\n", charlie.Courses[0])
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

HOW TO RUN THIS CODE:

1.  Open a terminal or command prompt.
2.  Navigate to the directory where you saved this file.
3.  Use the `go run` command to compile and execute the file:
    `go run 18_handling_json.go`

You have now learned how to seamlessly convert data between Go's typed structs
and the universal JSON format. This skill is the backbone of modern web
development in Go.
*/
