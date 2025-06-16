// Part 2, Lesson 8: Arrays and Slices
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file serves as the lesson and demonstration for arrays and slices.
// It explains the core concepts through structured comments and
// provides a runnable example of their implementation.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

So far, we've worked with single variables (`int`, `string`, etc.). But what if
we need to work with a collection of items, like a list of names or a set of scores?
This is where ARRAYS and SLICES come in.

ARRAYS
An ARRAY is a numbered sequence of elements of a specific length. The key feature
of an array is that its size is FIXED and defined at the time of creation. [1] An array
of 5 integers is a different type from an array of 6 integers. Because of this
rigidity, arrays are not as common in Go as their more flexible cousin, the slice.

SLICES
A SLICE is the real workhorse of Go. [2] It's a more flexible, powerful, and common way
to work with sequences of data. A slice is a lightweight data structure that provides
a "view" into a portion of an underlying array. Unlike arrays, slices can be
resized dynamically. [3] Most of the time when you're working with lists of things in Go,
you'll be using a slice.
*/

package main

import "fmt"

func main() {
	// --- Part 1: Arrays ---

	// An array is declared with its size inside the brackets: `var name [size]type`.
	// This creates an array that can hold exactly 3 strings.
	var favoriteFruits [3]string

	// Like other variables, an array is initialized with the zero value for its type.
	// For strings, the zero value is "".
	fmt.Printf("Initial empty array: %q\n", favoriteFruits)

	// We access and assign values using zero-based indexing.
	favoriteFruits[0] = "Apple"
	favoriteFruits[1] = "Banana"
	favoriteFruits[2] = "Cherry"
	fmt.Printf("Array after assignment: %q\n", favoriteFruits)

	// The length of an array is part of its type. This rigidity makes them less flexible.
	// favoriteFruits[3] = "Mango" // This would cause a compile-time error: index out of bounds.

	// You can declare and initialize an array in one line.
	primeNumbers := [5]int{2, 3, 5, 7, 11}
	fmt.Printf("Prime numbers array: %v\n", primeNumbers)

	// The `len()` function returns the size of the array.
	fmt.Printf("The primeNumbers array has %d elements.\n", len(primeNumbers))

	// --- Part 2: Slices - The Flexible Way ---

	// A slice is declared without a size in the brackets: `name := []type{...}`.
	// This creates a slice of strings.
	studentNames := []string{"Alice", "Bob", "Charlie"}
	fmt.Println("\n--- Working with Slices ---")
	fmt.Printf("Initial studentNames slice: %v\n", studentNames)

	// The `len()` function also works on slices, giving the number of elements.
	fmt.Printf("The slice has %d students.\n", len(studentNames))

	// --- Part 3: Manipulating Slices with `append` ---

	// The real power of slices is that they are dynamic. We can add to them using
	// the built-in `append` function.
	// `append` returns a NEW slice. You must assign the result back to the same variable.
	studentNames = append(studentNames, "Diana")
	fmt.Printf("After appending 'Diana': %v\n", studentNames)
	fmt.Printf("The slice now has %d students.\n", len(studentNames))

	// You can append multiple items at once.
	studentNames = append(studentNames, "Eve", "Frank")
	fmt.Printf("After appending two more names: %v\n", studentNames)

	// To append another slice, you must "unfurl" it with the `...` operator.
	newStudents := []string{"Grace", "Heidi"}
	studentNames = append(studentNames, newStudents...)
	fmt.Printf("After appending another slice: %v\n", studentNames)

	// --- Part 4: Creating a Slice from a Slice ---

	// You can create a new slice that "views" a portion of an existing one.
	// This is done with the syntax `aSlice[low:high]`.
	// It includes the element at `low` index and goes up to, but does not include,
	// the element at `high` index.
	fmt.Println("\n--- Creating sub-slices ---")

	// Get a slice of the first three students (indices 0, 1, 2).
	firstThree := studentNames[0:3]
	fmt.Printf("First three students: %v\n", firstThree)

	// Get a slice of the middle students (indices 3, 4).
	middleStudents := studentNames[3:5]
	fmt.Printf("Middle students: %v\n", middleStudents)

	// IMPORTANT: Slicing does NOT copy the underlying data. The new slice
	// points to the same backing array. Modifying the sub-slice will modify
	// the original.
	fmt.Println("\n--- Slices share underlying data ---")
	fmt.Printf("Original slice before change: %v\n", studentNames)
	middleStudents[0] = "EVE-WAS-HERE" // Change the first element of `middleStudents`
	fmt.Printf("Original slice AFTER change: %v\n", studentNames)

	// --- Part 5: Creating Slices with `make` ---

	// Sometimes you want to create a slice with a certain initial size, but you
	// don't have the values yet. The `make` function is used for this.
	// `make([]type, length, capacity)`
	// LENGTH is the number of elements the slice contains (initialized to zero values).
	// CAPACITY is the size of the underlying array that is allocated.
	scores := make([]int, 5, 10) // Create a slice of ints with length 5 and capacity 10.
	fmt.Println("\n--- Using the `make` function ---")
	fmt.Printf("Made slice: %v\n", scores)
	fmt.Printf("Length: %d, Capacity: %d\n", len(scores), cap(scores))
	scores[0] = 100
	scores[4] = 95
	fmt.Printf("Made slice after assignment: %v\n", scores)
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

KEY TAKEAWAYS:
- ARRAYS have a fixed size that is part of their type (e.g., `[4]int`).
- SLICES are dynamic and their size can change. This makes them far more common. [8]
- Use `[]type{}` to create a slice with initial values.
- Use the built-in `append()` function to add elements to a slice. Remember to assign
  the result back: `mySlice = append(mySlice, newValue)`.
- Use the `[low:high]` syntax to create a sub-slice. Be aware that this new slice
  SHARES data with the original.
- Use `make([]type, len, cap)` to create a slice with a specific initial length
  and capacity. [12]

HOW TO RUN THIS CODE:

1.  Open a terminal or command prompt.
2.  Navigate to the directory where you saved this file.
3.  Use the `go run` command to compile and execute the file:
    `go run 8_arrays_and_slices.go`
*/
