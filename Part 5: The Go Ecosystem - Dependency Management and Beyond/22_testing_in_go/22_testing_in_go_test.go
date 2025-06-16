// This is the test file for 22_testing_in_go.go.
// It must be in the same package and have the `_test.go` suffix.

package main

import (
	"testing" // The testing package is essential for writing tests.
)

// --- Test 1: A Simple Test Function ---

// TestAdd is a test function for our `Add` function.
// The name `TestAdd` is important: it starts with `Test` and is followed by the
// name of the function it is testing (a common convention).
func TestAdd(t *testing.T) {
	// Arrange: Set up the inputs and expected output.
	a, b := 5, 10
	expected := 15

	// Act: Call the function we are testing.
	got := Add(a, b)

	// Assert: Check if the result is what we expected.
	if got != expected {
		// If the result is not correct, we fail the test by calling t.Errorf.
		// t.Errorf logs the error message and marks the test as failed, but
		// allows other tests to continue running.
		t.Errorf("Add(%d, %d) = %d; want %d", a, b, got, expected)
	}
}

// --- Test 2: Table-Driven Tests ---

// Testing every single case one-by-one can be repetitive. A common and powerful
// pattern in Go is the "table-driven test". We define a "table" (a slice) of
// test cases, and then iterate over them in a single test function.

// TestDivide uses a table-driven approach to test our `Divide` function.
func TestDivide(t *testing.T) {
	// Arrange: Define the table of test cases. Each item in the slice is a
	// struct containing the inputs and expected outputs for one test case.
	testCases := []struct {
		name        string  // A name for the test case.
		a, b        float64 // Inputs
		expected    float64 // Expected result
		expectError bool    // Whether we expect an error or not.
	}{
		// Case 1: Simple, valid division.
		{name: "simple division", a: 10, b: 2, expected: 5, expectError: false},
		// Case 2: Division resulting in a fraction.
		{name: "fractional result", a: 1, b: 4, expected: 0.25, expectError: false},
		// Case 3: Division by zero, which should produce an error.
		{name: "division by zero", a: 100, b: 0, expected: 0, expectError: true},
	}

	// Act & Assert: Loop through the table.
	for _, tc := range testCases {
		// `t.Run` creates a sub-test for each case in our table. This is great
		// because it gives us clear output about which specific case failed.
		t.Run(tc.name, func(t *testing.T) {
			got, err := Divide(tc.a, tc.b)

			// Check if we got an error when we weren't expecting one.
			if !tc.expectError && err != nil {
				// t.Fatalf logs the error and stops the execution of THIS sub-test immediately.
				t.Fatalf("Divide(%f, %f) produced an unexpected error: %v", tc.a, tc.b, err)
			}

			// Check if we did NOT get an error when we WERE expecting one.
			if tc.expectError && err == nil {
				t.Fatalf("Divide(%f, %f) expected an error but got none", tc.a, tc.b)
			}

			// Check if the result is correct, but only if we weren't expecting an error.
			if !tc.expectError && got != tc.expected {
				t.Errorf("Divide(%f, %f) = %f; want %f", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}
