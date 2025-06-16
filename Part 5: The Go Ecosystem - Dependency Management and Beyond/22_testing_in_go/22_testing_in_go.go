// Part 5, Lesson 22: Testing in Go
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file serves as the lesson and demonstration for Go's built-in testing.
// The functions in this file will be tested by the code in `22_testing_in_go_test.go`.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

Writing code is only half the battle. How do we know our code actually works
correctly now and in the future? The answer is AUTOMATED TESTING.

Go has a first-class, built-in testing framework that is simple, lightweight,
and deeply integrated into the language's tooling. You don't need any third-party
libraries to write effective tests.

THE CORE CONCEPTS OF GO TESTING
1.  **Test Files:** Test code lives in files that end with the `_test.go` suffix.
    For this lesson, our tests will be in `22_testing_in_go_test.go`.

2.  **Test Functions:** A test function begins with the word `Test` and takes a single
    argument: `t *testing.T`. The `t` variable is our toolkit for reporting
    test failures.

3.  **The `go test` Command:** You run your tests from the terminal using the simple
    `go test` command. Go automatically finds all the `_test.go` files in the
    current directory and runs them.

For this lesson, we have two files: this one, which contains the functions we
want to test, and the `_test.go` file, which contains the actual tests.
*/

package main

import (
	"errors"
	"fmt"
)

// --- Part 1: The Functions to Be Tested ---

// Here are two simple functions that we will write tests for. The goal of our
// tests will be to verify that these functions produce the correct output
// for a given set of inputs.

// Add returns the sum of two integers. It's a "pure" function, making it
// very easy to test.
func Add(a, b int) int {
	return a + b
}

// Divide returns the result of a division. It also returns an error if the
// caller attempts to divide by zero, giving us a chance to test error conditions.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// --- Part 2: Main Execution Block ---

// This main function allows the file to be run directly with `go run`. It's not
// part of the testing workflow, but it helps demonstrate what our functions do.
// The real verification happens when we run `go test`.
func main() {
	fmt.Println("This file contains functions to be tested.")
	fmt.Println("Run `go test -v` in your terminal to see the test results.")
	fmt.Println("\n--- Demonstrating Add function ---")
	sum := Add(5, 10)
	fmt.Printf("Add(5, 10) = %d\n", sum)

	fmt.Println("\n--- Demonstrating Divide function ---")
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Divide(10, 2) = %f\n", result)
	}

	result, err = Divide(10, 0)
	if err != nil {
		fmt.Printf("Divide(10, 0) produced an error as expected: %v\n", err)
	}
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

HOW TO RUN THE TESTS:

1.  **Save Both Files:** Make sure you have saved both `22_testing_in_go.go`
    and `22_testing_in_go_test.go` in the same directory.

2.  **Open a Terminal:** Navigate to the directory where you saved the files.

3.  **Run the Tests:** Execute the `go test` command. This command finds and
    runs all functions starting with `Test` in all files ending with `_test.go`.

    `go test`

    If all tests pass, you'll see a simple `ok` message.

4.  **Run with Verbose Output:** To see which tests are being run and get more
    detailed output, use the `-v` (verbose) flag. This is highly recommended.

    `go test -v`

    You will now see each test function name and whether it PASSED, along with
    any log messages from the tests.
*/
