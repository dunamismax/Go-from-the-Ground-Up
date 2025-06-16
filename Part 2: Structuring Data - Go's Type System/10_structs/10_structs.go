// Part 2, Lesson 10: Structs
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file serves as the lesson and demonstration for structs.
// It explains the core concepts through structured comments and
// provides a runnable example of their implementation.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

WHAT IS A STRUCT?
So far, we've used basic types (`int`, `string`), and collections of those types
(`[]string`, `map[string]int`). But what if we need to group different types
of related data together? For example, a "user" isn't just a name or an age; a
user has a name, AND an age, AND an email, etc.

A STRUCT (short for structure) is a composite data type that allows you to group
together variables of different types under a single name. [1] It's a way to create
your own custom, complex data types. [2]

Think of a struct as a blueprint for creating a logical entity. Once you define
the blueprint, you can create multiple "instances" of it, each with its own data.
*/

package main

import "fmt"

// --- Part 1: Defining a Struct ---

// We define a new struct type using the `type` and `struct` keywords.
// By convention, struct names are often capitalized (e.g., `User`).
// This defines a blueprint for a `User`. It doesn't create one yet.
// Each piece of data inside the struct is called a FIELD. Each field has a
// name (e.g., `Username`) and a type (e.g., `string`).
type User struct {
	ID        int
	Username  string
	Email     string
	IsActive  bool
	Followers int
}

// Structs can also be composed of other structs. This is a key concept in Go
// called COMPOSITION. Let's define another struct.
type Address struct {
	Street     string
	City       string
	PostalCode string
}

// Now we can create a more complex struct that contains another struct.
type Customer struct {
	Name           string
	BillingAddress Address // This field's type is the `Address` struct we just defined.
	ContactInfo    User    // We can even embed the User struct.
}

func main() {
	// --- Part 2: Creating an Instance of a Struct ---

	// Now that we have our `User` blueprint, we can create a variable of that type.
	// This creates a `User` struct where all fields are set to their zero values.
	var user1 User
	fmt.Println("--- Creating and Using Structs ---")
	fmt.Printf("Zero-valued user1: %+v\n", user1) // %+v prints fields and values

	// We can assign values to the fields using dot notation.
	user1.ID = 1
	user1.Username = "dunamismax"
	user1.Email = "instructor@gofromthegroundup.com"
	user1.IsActive = true
	user1.Followers = 1337
	fmt.Printf("user1 after assignment: %+v\n", user1)

	// A more common way to create a struct instance is with a "struct literal".
	// This lets you set the field values at the time of creation.
	// Providing the field names makes the code more readable and robust. [9]
	user2 := User{
		ID:        2,
		Username:  "gopher",
		Email:     "gopher@golang.org",
		IsActive:  true,
		Followers: 99999,
	}
	fmt.Printf("user2 created with a struct literal: %+v\n", user2)

	// --- Part 3: Accessing Struct Fields ---

	// As we've seen, you access the data within a struct's fields using dot notation.
	fmt.Println("\n--- Accessing Struct Fields ---")
	fmt.Printf("Username for user2 is: %s\n", user2.Username)
	fmt.Printf("Is user1 active? %t\n", user1.IsActive)

	// --- Part 4: Structs as Function Arguments ---

	// You can pass entire structs to functions. Like other types in Go, structs are
	// PASSED BY VALUE. This means the function receives a COPY of the struct,
	// and any changes made inside the function will not affect the original.
	fmt.Println("\n--- Passing Structs to Functions ---")
	displayUser(user2)

	// To allow a function to modify the original struct, you would pass a POINTER
	// to it, just like we learned in the pointers lesson. We will see more of this
	// later when we learn about methods.

	// --- Part 5: Working with Embedded Structs ---

	// Creating an instance of a composed struct works the same way. You create
	// literals for the structs inside as well.
	customer1 := Customer{
		Name: "Jane Doe",
		BillingAddress: Address{
			Street:     "123 Go Lane",
			City:       "Gopherville",
			PostalCode: "12345",
		},
		ContactInfo: user1, // We can use a previously created struct.
	}

	fmt.Println("\n--- Working with Embedded Structs ---")
	fmt.Printf("Customer: %+v\n", customer1)

	// To access the fields of an embedded struct, you chain the dot notation.
	fmt.Printf("Customer lives in city: %s\n", customer1.BillingAddress.City)
	fmt.Printf("Customer's contact username is: %s\n", customer1.ContactInfo.Username)
}

// This function takes a User struct as an argument.
// It receives a COPY of the struct passed to it.
func displayUser(u User) {
	fmt.Printf("  Displaying user inside function: ID #%d, Username: %s\n", u.ID, u.Username)
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

KEY TAKEAWAYS:
- A STRUCT is a user-defined type that groups together fields of different types.
- Use the `type YourName struct { ... }` syntax to define the blueprint.
- Create instances using a struct literal: `instance := YourName{Field1: value1, ...}`.
- Access and modify fields using dot notation (e.g., `instance.Field1`).
- You can compose structs by having a field whose type is another struct.
- Structs are passed BY VALUE to functions, meaning the function gets a copy.

HOW TO RUN THIS CODE:

1.  Open a terminal or command prompt.
2.  Navigate to the directory where you saved this file.
3.  Use the `go run` command to compile and execute the file:
    `go run 10_structs.go`
*/
