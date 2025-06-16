// Part 1, Lesson 3: Packages and Imports
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file serves as the lesson and demonstration for using packages.
// It explains the core concepts through structured comments and
// provides a runnable example of their implementation.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

WHAT ARE PACKAGES?
So far, we've only used the `fmt` package. But what is a package? A PACKAGE is a
way to organize and reuse code in Go. Think of it as a toolbox that contains
a collection of related tools (FUNCTIONS, TYPES, and VARIABLES).

THE STANDARD LIBRARY
Go comes with a rich STANDARD LIBRARY, which is a collection of core packages
that are available to you without having to install anything extra. [8] These packages
provide tools for a huge range of common tasks, from performing mathematical
calculations to building web servers. Today, we will explore two more very
useful packages: `math` and `strings`.

HOW IMPORTS WORK
The `import` keyword tells Go which packages your program needs to use. When you
import a package, you gain access to its EXPORTED members. An exported member
(like a function) starts with a CAPITAL letter (e.g., `fmt.Println`, `math.Sqrt`).
If a member starts with a lowercase letter, it is private to that package and
cannot be used outside of it.
*/

// --- Importing Multiple Packages ---

// To import multiple packages, you can use a single `import` keyword followed
// by parentheses. This is the standard, idiomatic way to do it in Go.
package main

import (
	"fmt"
	"math"    // Contains mathematical functions and constants. [2]
	"strings" // Contains functions for string manipulation. [1, 3]
)

// The `main` function is the entry point of our program.
func main() {
	// --- Part 1: Using the `math` Package ---

	// The `math` package provides basic constants and mathematical functions. [2]
	// To use a function from a package, you use "dot notation": `<PackageName>.<FunctionName>()`
	// For example, let's calculate the square root of a number.
	var number float64 = 81
	squareRoot := math.Sqrt(number)
	fmt.Printf("The square root of %.1f is %.1f\n", number, squareRoot)

	// The `math` package also has useful constants, like Pi.
	// Note that `Pi` is a constant from the package, so it doesn't have `()` at the end. [2]
	circleRadius := 10.0
	circleCircumference := 2 * math.Pi * circleRadius
	fmt.Printf("A circle with a radius of %.1f has a circumference of %f\n", circleRadius, circleCircumference)

	// Let's use `Ceil` to round a number UP to the nearest integer,
	// and `Floor` to round a number DOWN.
	someFloat := 4.78
	roundedUp := math.Ceil(someFloat)
	roundedDown := math.Floor(someFloat)
	fmt.Printf("Original: %f, Rounded Up: %.1f, Rounded Down: %.1f\n", someFloat, roundedUp, roundedDown)

	// --- Part 2: Using the `strings` Package ---

	// The `strings` package provides many useful functions for working with strings. [5]
	// Let's declare a string to work with.
	originalString := "Go is an amazing and powerful language."

	// `ToUpper` converts a string to all uppercase.
	// `HasPrefix` checks if a string starts with a certain substring. It returns a bool. [3]
	fmt.Println("Uppercase:", strings.ToUpper(originalString))
	fmt.Println("Does it start with 'Go'?", strings.HasPrefix(originalString, "Go"))

	// `ReplaceAll` finds all instances of a substring and replaces them.
	// It takes three arguments: the original string, the substring to find, and the replacement.
	replacedString := strings.ReplaceAll(originalString, "amazing", "incredibly")
	fmt.Println("After replacement:", replacedString)

	// `Join` is used to concatenate (join) elements of a string slice into a single string.
	// It takes a slice of strings and a separator.
	words := []string{"Let's", "build", "something", "cool!"}
	sentence := strings.Join(words, " ") // Join with a space separator.
	fmt.Println("Joined sentence:", sentence)
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

KEY TAKEAWAYS:
- PACKAGES are collections of reusable Go code.
- Go's STANDARD LIBRARY provides many powerful packages for common tasks.
- Use the `import` keyword to make a package's code available in your file.
- Access members of a package using dot notation (e.g., `math.Pi`, `strings.ToUpper`).
- A package member must be EXPORTED (start with a capital letter) to be accessible.

HOW TO RUN THIS CODE:

1.  Open a terminal or command prompt.
2.  Navigate to the directory where you saved this file.
3.  Use the `go run` command to compile and execute the file:
    `go run 3_packages_and_imports.go`

Explore the official Go documentation to see what other functions are available in
the `math` and `strings` packages!
*/
