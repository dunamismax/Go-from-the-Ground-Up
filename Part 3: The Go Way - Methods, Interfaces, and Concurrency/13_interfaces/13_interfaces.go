// Part 3, Lesson 13: Interfaces
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file serves as the lesson and demonstration for interfaces.
// It explains the core concepts through structured comments and
// provides a runnable example of their implementation.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

WHAT IS AN INTERFACE?

An interface in Go is a type that defines a set of method signatures. It's like a
contract or a blueprint. [15, 9] If a type (like a struct) has all the methods
listed in the interface, it is said to "satisfy" or "implement" that interface. [4]

This is the key to POLYMORPHISM in Go. [5] Polymorphism allows us to write functions
that can operate on objects of different types, as long as they all share the
same behavior defined by the interface. [19]

THE GO WAY: IMPLICIT IMPLEMENTATION

Unlike other languages (like Java or C#), Go's interfaces are satisfied IMPLICITLY. [10, 24]
You do not need to use a keyword like `implements` to declare that your struct
adheres to an interface. [11, 23] If your type has the required methods, it automatically
fulfills the contract. [9] This makes Go's code less verbose and more flexible. [20]
*/

package main

import (
	"fmt"
	"math"
)

// --- Part 1: Defining an Interface ---

// Let's define an interface for a geometric shape.
// A "Shape" is anything that has a defined area and perimeter.
// An interface is created with the `type` and `interface` keywords.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// --- Part 2: Implementing the Interface with Structs ---

// Now, let's create two different structs: `Rectangle` and `Circle`.

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

// To make `Rectangle` satisfy the `Shape` interface, we must implement
// BOTH `Area()` and `Perimeter()` methods for it.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// To make `Circle` satisfy the `Shape` interface, we must also implement
// BOTH `Area()` and `Perimeter()` methods for it.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// --- Part 3: Using Interfaces for Polymorphism ---

// This is where the power of interfaces shines. We can write a function
// that takes a `Shape` as an argument. It doesn't need to know or care
// whether it's receiving a `Rectangle` or a `Circle`. [1] It only knows that
// whatever it receives, it will have `Area()` and `Perimeter()` methods.
func PrintShapeDetails(s Shape) {
	fmt.Printf("Shape details: %+v\n", s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
	fmt.Println("---")
}

func main() {
	// Let's create instances of our structs.
	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 7}

	// Because BOTH `Rectangle` and `Circle` have the methods defined in the
	// `Shape` interface, we can pass instances of them to our generic function.
	fmt.Println("Printing details for a Rectangle:")
	PrintShapeDetails(rect)

	fmt.Println("Printing details for a Circle:")
	PrintShapeDetails(circ)

	// We can also create collections (like a slice) of interfaces. This allows
	// us to store different concrete types in the same data structure.
	fmt.Println("\nIterating over a slice of Shapes:")
	shapes := []Shape{rect, circ}
	for _, shape := range shapes {
		// Inside the loop, `shape` is of type `Shape`, but its underlying
		// concrete type is either a `Rectangle` or a `Circle`.
		PrintShapeDetails(shape)
	}
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

HOW TO RUN THIS CODE:

1.  Open a terminal or command prompt.
2.  Navigate to the directory where you saved this file.
3.  Use the `go run` command to compile and execute the file:
    `go run 13_interfaces.go`

KEY TAKEAWAYS:
- Interfaces define BEHAVIOR (what a type can DO) by listing method signatures.
- A type satisfies an interface IMPLICITLY if it implements all of the interface's methods.
- Interfaces allow for POLYMORPHISM, which means writing flexible functions that can
  work with multiple different types in a uniform way.
- This helps to DECOUPLE your code, meaning components don't need to know about
  the specific implementations of other components, only the behavior they provide.
*/
