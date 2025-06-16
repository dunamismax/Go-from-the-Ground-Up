// Part 5, Lesson 21: Go Modules and Dependencies
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file serves as the lesson and demonstration for Go Modules.
// It explains the core concepts through structured comments and
// provides a runnable example of their implementation using a third-party package.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

Welcome to Part 5! Until now, we have exclusively used Go's STANDARD LIBRARY
(packages like `fmt`, `net/http`, `os`, etc.). The standard library is powerful,
but the true strength of a modern programming language is its ecosystem—the vast
collection of open-source code written by the community.

WHAT IS A DEPENDENCY?
A DEPENDENCY is code from an external, third-party package that your project
needs to function. For example, you might need a package for advanced database
access, complex data visualization, or, in this case, a more powerful way to
compare data structures.

WHAT ARE GO MODULES?
GO MODULES is Go's built-in system for managing dependencies. A MODULE is a
collection of related Go packages that are versioned together as a single unit.
The modules system handles:
-   Defining your project as a module.
-   Tracking which versions of which dependencies your project needs.
-   Downloading the correct dependency versions automatically.

This lesson is different. You will need to run a few commands in your terminal
BEFORE you run this Go file. Follow the instructions at the bottom carefully.
*/

// --- Part 1: Importing a Third-Party Package ---

package main

import (
	// We still use the `fmt` package from the standard library.
	"fmt"

	// Now, we are importing the `cmp` package from Google. This package is
	// widely used, especially in tests, to provide a much safer and more
	// powerful way to compare if two Go values are equal. [2, 3]
	// When the Go toolchain sees this import path, it knows it needs to find and
	// download the `github.com/google/go-cmp` module if it hasn't already.
	"github.com/google/go-cmp/cmp"
)

// Let's define a simple struct we can use for our comparisons.
type User struct {
	ID   int
	Name string
}

func main() {
	// --- Part 2: Using the Third-Party Code ---

	// The `go mod` commands (which you will run) will have downloaded the
	// `go-cmp` package. Now we can use its functions.
	//
	// The `==` operator in Go works for simple structs, but it has limitations.
	// It doesn't work for slices or maps, and it can be tricky with other
	// complex types. The `go-cmp` package provides a robust alternative. [1, 5]

	// Create two identical instances of our User struct.
	userA := User{ID: 1, Name: "Alice"}
	userB := User{ID: 1, Name: "Alice"}

	// Create a different instance.
	userC := User{ID: 2, Name: "Bob"}

	// Use `cmp.Equal` to compare the two identical structs.
	areAandBEqual := cmp.Equal(userA, userB)
	fmt.Printf("Are userA and userB equal? %v\n", areAandBEqual)

	// Now compare two different structs.
	areAandCEqual := cmp.Equal(userA, userC)
	fmt.Printf("Are userA and userC equal? %v\n", areAandCEqual)

	// The real power of `go-cmp` comes from its ability to show a "diff" —
	// a human-readable report of exactly what is different between two values.
	// This is incredibly useful for debugging.
	fmt.Println("\n--- Finding the Difference ---")
	diff := cmp.Diff(userA, userC)
	if diff != "" {
		fmt.Printf("Found a difference between userA and userC:\n%s", diff)
	}
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

HOW TO RUN THIS CODE (VERY IMPORTANT):

This lesson requires you to interact with the Go Modules system using commands
in your terminal. The `go run` command alone is not enough to start.

1.  **Create a New Directory:** Create a new, empty folder on your computer for
    this lesson. For example:
    `mkdir go_modules_project && cd go_modules_project`

2.  **Save This File:** Save the code above into that new directory as
    `21_go_modules_and_dependencies.go`.

3.  **Initialize the Module (`go mod init`):**
    In your terminal, inside the new directory, run the `go mod init` command.
    This command creates a new `go.mod` file, which officially turns your
    directory into a Go module and tracks its dependencies. You must give your
    module a name. For this lesson, let's call it `myprogram`.

    `go mod init myprogram`

    You will see a message: `go: creating new go.mod: module myprogram`.
    The `go.mod` file is a simple text file that now exists in your directory.

4.  **Add Dependencies (`go mod tidy`):**
    Now that you have a module file and a `.go` file with imports, you can ask
    the Go toolchain to find, download, and track the necessary dependencies.
    The `go mod tidy` command is the standard way to do this. It looks at your
    `import` statements, finds what's missing, and automatically downloads them.

    `go mod tidy`

    You will see output as Go downloads the `github.com/google/go-cmp` module
    and any modules it depends on. This also creates a `go.sum` file, which
    contains cryptographic hashes of the dependencies to ensure they haven't been
    tampered with.

5.  **Run the Program:**
    Now that the dependencies are downloaded and tracked in your `go.mod` file,
    you can run your program as usual:

    `go run .`

    You should see the comparison results and the diff printed to your console.
    You have successfully managed a dependency and run a program that uses
    powerful third-party code!
*/
