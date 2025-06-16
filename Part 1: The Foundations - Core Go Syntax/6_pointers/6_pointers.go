// Part 1, Lesson 6: Pointers
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file serves as the lesson and demonstration for pointers.
// It explains the core concepts through structured comments and
// provides a runnable example of their implementation.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

Welcome to one of the most powerful, and often misunderstood, concepts in
programming: POINTERS. This lesson will introduce them gently.

WHAT IS MEMORY?
When your program runs, every variable you create is stored in your computer's
memory (RAM). Think of memory as a giant collection of boxes, where each box has
a unique ADDRESS.

WHAT IS A POINTER?
A normal variable (like an `int`) holds a direct value, such as `42`.
A POINTER is a special type of variable that doesn't hold a value itself, but
instead holds the MEMORY ADDRESS of another variable. It "points to" the location
where the actual data is stored.

WHY USE POINTERS?
1.  EFFICIENCY: When you pass data to a function, Go typically creates a copy. If
    you have a very large piece of data, copying it can be slow and use a lot
    of memory. Passing a pointer (just the tiny address) is much faster.

2.  MODIFICATION: Because functions work on copies of data, they cannot normally
    change the original variables you pass to them. If you want a function to
    be able to modify the original variable, you must give it a pointer to that
    variable. This is the most common reason we use them.
*/

package main

import "fmt"

// --- Part 1: Functions and Value Copies ---

// Let's first demonstrate how functions in Go work on copies.
// This function takes an integer, tries to change it, but the change
// will NOT be reflected outside this function. This is known as PASS-BY-VALUE.
func changeValueByCopy(value int) {
	fmt.Printf("  Inside changeValueByCopy, initial value is: %d\n", value)
	value = 100 // We change the COPY of the variable, not the original.
	fmt.Printf("  Inside changeValueByCopy, changed value to: %d\n", value)
}

// --- Part 2: Functions with Pointers for Modification ---

// Now, let's create a function that CAN modify the original value.
// It accepts a POINTER to an integer. The type `*int` means "a pointer to an int".
func changeValueViaPointer(ptr *int) {
	fmt.Printf("  Inside changeValueViaPointer, the pointer holds address: %p\n", ptr)
	// The `*` operator here is used for DEREFERENCING. It means "go to the
	// memory address this pointer is holding and modify the value stored there".
	*ptr = 100
	fmt.Println("  Inside changeValueViaPointer, we changed the value at that address to 100.")
}

func main() {
	// --- Part 3: Getting a Variable's Memory Address ---

	// Let's declare a simple variable. It's stored somewhere in memory.
	myNumber := 42
	fmt.Printf("Original variable 'myNumber' -> Value: %d, Type: %T\n", myNumber, myNumber)

	// We can see the memory address of `myNumber` using the `&` operator.
	// `&` should be read as "address of".
	// The address will be printed in hexadecimal format (e.g., 0xc000018030).
	// This specific address will be different every time you run the program.
	fmt.Printf("Memory address of 'myNumber' -> Address: %p\n", &myNumber)

	// --- Part 4: Declaring and Using a Pointer ---

	// Let's declare a variable that can hold the address of `myNumber`.
	// Its type is `*int` (a pointer to an integer).
	var numberPointer *int

	// Now we assign the memory address of `myNumber` to our pointer variable.
	numberPointer = &myNumber

	fmt.Printf("Pointer variable 'numberPointer' -> Value (an address): %p, Type: %T\n", numberPointer, numberPointer)

	// To get the actual value that the pointer is pointing to, we DEREFERENCE it
	// using the `*` operator.
	// `*numberPointer` should be read as "the value at the address stored in numberPointer".
	fmt.Printf("Value at the address 'numberPointer' points to -> Value: %d\n", *numberPointer)

	// --- Part 5: Putting It All Together in Functions ---

	fmt.Println("\n--- Demonstrating Pass-by-Value ---")
	originalValue := 50
	fmt.Printf("Before calling function, 'originalValue' is: %d\n", originalValue)
	changeValueByCopy(originalValue)
	fmt.Printf("After calling function, 'originalValue' is STILL: %d\n", originalValue)
	fmt.Println("(The original value was not changed because the function got a copy.)")

	fmt.Println("\n--- Demonstrating Pass-by-Pointer ---")
	fmt.Printf("Before calling function, 'originalValue' is: %d\n", originalValue)
	// We pass the MEMORY ADDRESS of `originalValue` to the function.
	changeValueViaPointer(&originalValue)
	fmt.Printf("After calling function, 'originalValue' is NOW: %d\n", originalValue)
	fmt.Println("(The original value was changed because the function had its address.)")

}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

KEY TAKEAWAYS:
- Every variable has a value and a MEMORY ADDRESS.
- A POINTER is a variable that stores a memory address.
- The `&` operator gives you the memory address of a variable. (`&x` -> "address of x")
- The `*` operator has two jobs:
    1. In a type declaration (`*int`), it means "pointer to".
    2. With a pointer variable (`*p`), it DEREFERENCES the pointer, giving you the
       value stored at that address.
- Passing a pointer to a function allows that function to modify the original variable's data.

HOW TO RUN THIS CODE:

1.  Open a terminal or command prompt.
2.  Navigate to the directory where you saved this file.
3.  Use the `go run` command to compile and execute the file:
    `go run 6_pointers.go`
*/
