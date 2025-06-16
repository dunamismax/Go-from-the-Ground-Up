// Part 1, Lesson 7: Error Handling
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file serves as the lesson and demonstration for idiomatic error handling in Go.
// It explains the core concepts through structured comments and
// provides a runnable example of their implementation.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

HOW GO HANDLES ERRORS
Many programming languages (like Java, Python, or C#) use a mechanism called
"exceptions" and `try...catch` blocks to handle errors. Go takes a very different
and explicit approach.

In Go, ERRORS ARE VALUES.

There is a built-in interface type called `error`. When a function might fail,
instead of "throwing an exception," it will return its normal result PLUS a second
value of type `error`. [1] This is the idiomatic `(value, error)` return pattern we
saw briefly in the functions lesson.

- If the function succeeds, the `error` value it returns will be `nil`.
- If the function fails, it will return a non-nil `error` value describing what went wrong.

`nil` is a special value in Go representing the zero value for pointers, interfaces,
maps, slices, channels, and function types. For our purposes today, just know that
`if err == nil`, it means "no error occurred."

This approach forces you, the programmer, to explicitly check for and handle
potential failures right where they happen, leading to more robust and reliable code. [3]
*/

package main

import (
	"errors" // A package to create basic error values.
	"fmt"
)

// --- Part 1: Creating a Function That Can Fail ---

// Let's create a `divide` function. Division can fail if the divisor is zero.
// This function follows the idiomatic Go pattern: it returns the result (`float64`)
// and an `error`.
func divide(dividend float64, divisor float64) (float64, error) {
	// This is the condition we need to check for.
	if divisor == 0.0 {
		// If the condition for failure is met, we return a zero value for our
		// result type (0.0 for float64) and a NEW error.
		// `errors.New()` creates a simple error with a given message.
		return 0.0, errors.New("cannot divide by zero")
	}

	// If we get here, it means the function can succeed.
	// We perform the calculation and return the result, along with `nil`
	// to indicate that there was no error.
	result := dividend / divisor
	return result, nil
}

// --- Part 2: Creating Formatted Errors ---

// The `errors.New` function is simple. For more context, the `fmt` package
// provides `fmt.Errorf`, which works just like `fmt.Printf` but returns an error value.
func checkTemperature(temp float64) error {
	if temp > 100.0 {
		// Create a more descriptive error with dynamic values.
		return fmt.Errorf("temperature is %.1f, which is dangerously high", temp)
	}
	if temp < -50.0 {
		return fmt.Errorf("temperature is %.1f, which is dangerously low", temp)
	}

	// No problems, so we return nil.
	return nil
}

func main() {
	fmt.Println("Starting the Error Handling Lesson.")

	// --- Part 3: Handling Errors - The Success Case ---
	fmt.Println("\n--- Handling a successful operation ---")

	// We call our function and receive TWO values: the result and the error.
	result, err := divide(100.0, 5.0)

	// THE IDIOMATIC GO ERROR CHECK:
	// Immediately after a function call that can fail, check if `err` is not `nil`. [12]
	if err != nil {
		// This block only runs if an error occurred.
		fmt.Println("An error occurred:", err)
	} else {
		// This block runs if `err` was `nil` (i.e., the call was successful).
		fmt.Printf("Success! The result of the division is: %.2f\n", result)
	}

	// --- Part 4: Handling Errors - The Failure Case ---
	fmt.Println("\n--- Handling a failed operation ---")

	// Let's call the function with values that will cause it to fail.
	result, err = divide(100.0, 0.0)

	// We perform the exact same check.
	if err != nil {
		// This time, `err` is NOT nil, so this block will execute.
		fmt.Println("An error occurred:", err)
		// When an error happens, the `result` variable still exists but its value
		// (0.0 in this case) should be considered invalid and not used.
		fmt.Printf("The returned result was: %f (this value should not be trusted)\n", result)
	} else {
		fmt.Printf("Success! The result of the division is: %.2f\n", result)
	}

	// --- Part 5: Using Formatted Errors ---
	fmt.Println("\n--- Using formatted errors from fmt.Errorf ---")
	err = checkTemperature(150.5)
	if err != nil {
		fmt.Println("Error checking temperature:", err)
	}

	err = checkTemperature(-100.2)
	if err != nil {
		fmt.Println("Error checking temperature:", err)
	}

	err = checkTemperature(25.0)
	if err == nil {
		fmt.Println("Temperature 25.0 is within the safe range.")
	}
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

KEY TAKEAWAYS:
- In Go, ERRORS ARE VALUES of the built-in `error` type.
- Go does not use exceptions or `try/catch`.
- Functions that can fail should return a result and an error: `(value, error)`. [11]
- If an error is `nil`, the function succeeded. If it is not `nil`, it failed.
- The most common pattern in Go is to check for errors immediately after calling a
  function: `if err != nil { ... }`. [12]
- Use `errors.New("message")` for simple, static error messages.
- Use `fmt.Errorf("format string", args...)` for more detailed, dynamic errors.

HOW TO RUN THIS CODE:

1.  Open a terminal or command prompt.
2.  Navigate to the directory where you saved this file.
3.  Use the `go run` command to compile and execute the file:
    `go run 7_error_handling.go`
*/
