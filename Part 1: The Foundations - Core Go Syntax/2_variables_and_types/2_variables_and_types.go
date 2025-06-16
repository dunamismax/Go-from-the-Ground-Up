// Part 1, Lesson 2: Variables and Types
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file serves as the lesson and demonstration for variables and basic types.
// It explains the core concepts through structured comments and
// provides a runnable example of their implementation.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

Welcome to your second lesson! In the last lesson, we printed a fixed message.
To write more useful programs, we need a way to store and manage information
that can change. We do this with VARIABLES.

WHAT IS A VARIABLE?
A VARIABLE is a named storage location for data. Think of it like a labeled box
where you can keep a piece of information, like a name, an age, or a price.

STATIC TYPING
Go is a STATICALLY TYPED language. This means that when you create a variable,
you must tell the compiler what TYPE of data it will hold (like text, a whole
number, a decimal number, etc.). Once a variable's type is set, it cannot be
changed. [1] This helps prevent many common bugs before the program even runs.
*/

package main

import "fmt"

func main() {
	// --- Part 1: Declaring Variables with `var` ---

	// The `var` keyword is the fundamental way to declare a variable.
	// The pattern is: `var <variableName> <type>`
	// Here, we declare a variable named `courseName` that will hold a `string` type.
	var courseName string

	// Now we can assign a value to it using the `=` operator.
	courseName = "Go From The Ground Up"

	// We can print the value stored in the variable.
	fmt.Println("The name of this course is:", courseName)

	// You can also declare a variable and assign its initial value in one line.
	// The pattern is: `var <variableName> <type> = <value>`
	var lessonNumber int = 2
	fmt.Println("We are on lesson number:", lessonNumber)

	// --- Part 2: Go's Basic Data Types ---

	// Go has several built-in data types. Let's look at the most common ones. [15]

	// STRING: A sequence of text characters. String values are enclosed in double quotes.
	var aGreeting string = "Hello there!"

	// INT: An integer, which is a whole number (no decimal point).
	var userAge int = 30

	// FLOAT64: A floating-point number, which can have a decimal point.
	// It's called `float64` because it uses 64 bits of memory for higher precision. [11]
	var itemPrice float64 = 49.95

	// BOOL: A boolean value, which can only be `true` or `false`.
	var isLoggedIn bool = true

	fmt.Println(aGreeting, "Your age is", userAge, "and the item costs", itemPrice)
	fmt.Println("Is the user logged in?", isLoggedIn)

	// --- Part 3: The `:=` Short Declaration Operator ---

	// Inside a function, Go provides a shorter way to declare and initialize a
	// variable using the `:=` operator. [8] This is the most common way you'll see
	// variables created in Go.

	// This one line does two things:
	// 1. It declares a new variable named `projectName`.
	// 2. It infers the type from the value (in this case, `string`) and assigns it.
	// This is called TYPE INFERENCE. [8]
	projectName := "Go-from-the-Ground-Up"

	// You can only use `:=` when declaring a NEW variable. [2] If the variable already
	// exists, you must use the standard assignment operator `=`.
	projectName = "A new project name!" // This is valid.

	// The `:=` operator can only be used inside functions, not at the package level. [14]
	currentYear := 2025
	piApproximation := 3.14159
	isComplete := false

	// `Printf` is another function from `fmt`. The 'f' stands for "formatted".
	// The `%T` verb is a special placeholder that prints the TYPE of a variable.
	fmt.Printf("Variable 'projectName' has value '%s' and type %T\n", projectName, projectName)
	fmt.Printf("Variable 'currentYear' has value '%d' and type %T\n", currentYear, currentYear)
	fmt.Printf("Variable 'piApproximation' has value '%f' and type %T\n", piApproximation, piApproximation)
	fmt.Printf("Variable 'isComplete' has value '%t' and type %T\n", isComplete, isComplete)

	// --- Part 4: Zero Values ---

	// What happens if you declare a variable but don't give it a value?
	// In Go, it's automatically given a "ZERO VALUE". [3] This prevents bugs from
	// uninitialized variables holding random data. [4]

	// Let's declare some variables without initializing them.
	var defaultName string   // Zero value for string is "" (an empty string) [4]
	var defaultAge int       // Zero value for int is 0 [4]
	var defaultPrice float64 // Zero value for float64 is 0.0 [4]
	var defaultStatus bool   // Zero value for bool is false [4]

	// Let's print them to see their zero values.
	// Note the `%q` verb for the string, which prints it with quotes.
	fmt.Printf("Zero Values: Name: %q, Age: %d, Price: %f, Status: %t\n", defaultName, defaultAge, defaultPrice, defaultStatus)
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

KEY TAKEAWAYS:
- VARIABLES are named places to store data.
- Go is STATICALLY TYPED; a variable's type is fixed upon creation.
- Use `var <name> <type> = <value>` for full declaration.
- Use `<name> := <value>` for short declaration with type inference inside functions.
- Every type has a ZERO VALUE (e.g., 0, "", false), which is its default if
  no value is assigned. [9, 20]

HOW TO RUN THIS CODE:

1.  Open a terminal or command prompt.
2.  Navigate to the directory where you saved this file.
3.  Use the `go run` command to compile and execute the file:
    `go run 2_variables_and_types.go`

Experiment by changing the values or creating new variables of different types!
*/
