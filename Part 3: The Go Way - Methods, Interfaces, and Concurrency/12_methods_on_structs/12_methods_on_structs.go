// Part 3, Lesson 12: Methods on Structs
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file serves as the lesson and demonstration for methods.
// It explains the core concepts through structured comments and
// provides a runnable example of their implementation.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

WHAT ARE METHODS?

In previous lessons, we created structs to hold related data. But what if we
want to add BEHAVIOR to that data?

A METHOD is a function that is attached to a specific type. [2, 11] It's a way of saying,
"This function belongs to this struct and operates on its data." This is a
cornerstone of object-oriented style programming in Go, as it allows us to group
data and the operations on that data together. [5]

The type a method is attached to is called the RECEIVER. [1] The receiver is like a
special parameter that comes before the function name.

We will explore two types of receivers: VALUE receivers and POINTER receivers. [3]
*/

package main

import "fmt"

// We'll define a simple `Rectangle` struct to work with.
type Rectangle struct {
	Width  float64
	Height float64
}

// --- Part 1: Value Receiver Methods ---

// A VALUE RECEIVER method works on a COPY of the struct.
// This means that the method can access the struct's data, but it cannot
// CHANGE or MUTATE the original struct instance. [4]
//
// Value receivers are great for read-only operations. [3]

// The `(r Rectangle)` part before the function name `Area()` is the RECEIVER.
// It declares that `Area` is a method on the `Rectangle` type.
// By convention, the receiver variable is a short, lowercase name (e.g., `r` for `Rectangle`).
func (r Rectangle) Area() float64 {
	// Inside the method, `r` refers to the specific instance of the Rectangle
	// that the method was called on.
	return r.Width * r.Height
}

// --- Part 2: Pointer Receiver Methods ---

// A POINTER RECEIVER method works on a POINTER to the original struct.
// This means it gets a memory address, not a copy.
// Therefore, this method CAN change and mutate the original struct's values. [4]
//
// This is also more efficient for large structs, as it avoids copying
// a potentially large amount of data every time the method is called. [4]

// The `(r *Rectangle)` syntax indicates this is a pointer receiver.
// It receives a pointer to a `Rectangle` instance.
func (r *Rectangle) Scale(factor float64) {
	// We are modifying the fields of the original struct that `r` points to.
	r.Width = r.Width * factor
	r.Height = r.Height * factor
}

func main() {
	// --- Part 3: Using the Methods ---

	// Let's create an instance of our Rectangle.
	rect := Rectangle{Width: 10, Height: 5}

	fmt.Printf("Initial Rectangle: %+v\n", rect)

	// Call the value receiver method `Area`.
	// We use the dot notation, just like accessing a field.
	area := rect.Area()
	fmt.Printf("Area calculated with value receiver: %.2f\n", area)
	fmt.Printf("Rectangle after calling Area(): %+v (Note: it is unchanged)\n", rect)

	fmt.Println("\n--- Scaling the rectangle ---")

	// Now, let's call the pointer receiver method `Scale`.
	// Go conveniently allows us to call this method directly on the `rect` value.
	// It automatically converts `rect` to `&rect` (a pointer) for us. [7]
	rect.Scale(2) // This is syntactic sugar for (&rect).Scale(2)

	fmt.Printf("Rectangle after calling Scale(2): %+v (Note: it has been modified)\n", rect)

	// We can now call Area() again to see the new area.
	newArea := rect.Area()
	fmt.Printf("New area after scaling: %.2f\n", newArea)
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

HOW TO RUN THIS CODE:

1.  Open a terminal or command prompt.
2.  Navigate to the directory where you saved this file.
3.  Use the `go run` command to compile and execute the file:
    `go run 12_methods_on_structs.go`

KEY TAKEAWAYS:
- Methods add behavior to your data types (structs).
- VALUE receivers `(t MyType)` operate on a copy and cannot change the original.
- POINTER receivers `(t *MyType)` operate on the original data and can change it.
- Choose a pointer receiver if the method needs to modify the receiver, or if the
  struct is very large and you want to avoid the cost of copying. [4]
*/
