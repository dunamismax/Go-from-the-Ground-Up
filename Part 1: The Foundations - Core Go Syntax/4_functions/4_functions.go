// Part 1, Lesson 4: Functions
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file serves as the lesson and demonstration for functions.
// It explains the core concepts through structured comments and
// provides a runnable example of their implementation.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

WHAT IS A FUNCTION?
A FUNCTION is a reusable block of code that performs a specific task. [7] Functions are
the fundamental building blocks of a Go program. They help us organize code,
avoid repetition, and make our programs easier to read and maintain.

We have already been using functions! `fmt.Println()` is a function from the `fmt`
package, and our entire program runs inside the special `main` function.

Today, we'll learn how to create our own functions.
*/

package main

import "fmt"

// --- Part 1: Defining and Calling a Simple Function ---

// Here, we define a simple function called `printWelcomeMessage`.
// The `func` keyword starts the function declaration. [3]
// `printWelcomeMessage` is the name of our function.
// The `()` parentheses are for parameters (arguments), which this function doesn't have. [7]
// The `{}` curly braces contain the "body" of the function—the code that runs
// when the function is called.
func printWelcomeMessage() {
	fmt.Println("--------------------------------")
	fmt.Println("Welcome to the Function Lesson!")
	fmt.Println("--------------------------------")
}

// --- Part 2: Functions with Parameters ---

// Functions become much more powerful when they can accept input data.
// These inputs are called PARAMETERS.
// This function, `greetUser`, takes one parameter: `name` of type `string`.
// The syntax is `<parameterName> <type>`.
func greetUser(name string) {
	fmt.Printf("Hello, %s! It's great to have you here.\n", name)
}

// --- Part 3: Functions with Return Values ---

// Functions can also perform a task and return a result.
// To do this, we must declare the type of the data being returned.
// The syntax is `func functionName(params) <returnType>`. [2]
// This function takes two integers and returns a single integer.
func add(a int, b int) int {
	// The `return` keyword specifies the value to send back to the caller.
	sum := a + b
	return sum
}

// --- Part 4: Functions with Multiple Return Values ---

// A unique and powerful feature of Go is that functions can return more than one value. [4]
// This is extremely common in Go, especially for returning a result and an error status.
// The return types are listed inside parentheses `(type1, type2)`. [4]
// This function attempts to divide two numbers and returns the result AND a boolean
// indicating if the operation was successful.
func divide(dividend float64, divisor float64) (float64, bool) {
	// It's impossible to divide by zero.
	if divisor == 0.0 {
		// Return a zero value for the result and `false` for success.
		return 0.0, false
	}
	// If the divisor is not zero, perform the division and return the result
	// along with `true` to indicate success.
	result := dividend / divisor
	return result, true
}

// The `main` function is where our program execution begins.
// We will CALL our other functions from here to see them in action.
func main() {
	// Call the simple function.
	printWelcomeMessage()

	fmt.Println("\n--- Calling function with parameters ---")
	// Call the function and provide an ARGUMENT (the actual value) for the parameter.
	greetUser("Alice")
	greetUser("Bob")

	fmt.Println("\n--- Calling function with a return value ---")
	// Call the `add` function and store its return value in a variable.
	total := add(15, 27)
	fmt.Println("The result of adding 15 and 27 is:", total)

	fmt.Println("\n--- Calling function with multiple return values ---")
	// When calling a function with multiple returns, we can assign the results
	// to multiple variables at once.
	quotient, success := divide(100.0, 5.0)

	// Check the boolean `success` flag to see if the division worked.
	if success {
		fmt.Printf("Result of 100.0 / 5.0 is: %.2f\n", quotient)
	} else {
		fmt.Println("Division failed.")
	}

	// Let's try the failure case.
	quotient, success = divide(100.0, 0.0)
	if success {
		fmt.Printf("Result of 100.0 / 0.0 is: %.2f\n", quotient)
	} else {
		fmt.Println("Attempt to divide by zero failed as expected.")
	}
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

KEY TAKEAWAYS:
- FUNCTIONS are reusable blocks of code defined with the `func` keyword.
- PARAMETERS are variables listed in a function's definition that accept input.
- ARGUMENTS are the actual values passed to a function when it is called.
- Functions can `return` one or more values. Declaring the `return type` is mandatory.
- Returning multiple values is a core feature of Go, often used for returning a
  result and an error/status. [8]

HOW TO RUN THIS CODE:

1.  Open a terminal or command prompt.
2.  Navigate to the directory where you saved this file.
3.  Use the `go run` command to compile and execute the file:
    `go run 4_functions.go`
*/
