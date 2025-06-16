// Part 1, Lesson 1: Hello, World!
//
// Author: dunamismax
// Date: 06-15-2025
//
// This is your very first Go program and lesson. The lesson is taught
// through the comments in this file. Read them from top to bottom
// to understand what's happening.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

Welcome to your first lesson in Go! The journey to building simple, reliable,
and efficient software starts here.

WHAT IS A COMPILER?
Go is a COMPILED language. Unlike Python (an interpreted language), a Go program
is first translated into machine code by a COMPILER. This process creates a
single, standalone executable file that can be run on its own.

The `go run` command we will use is a convenient shortcut that compiles and runs
the program in one step without leaving the executable file behind. This makes
development feel fast and interactive.
*/

// --- The Main Package and Function ---

// Every runnable Go program must have a `package main`. Packages are how Go
// organizes and reuses code. The `main` package is special: it tells the Go
// compiler that this package should be compiled into an executable program.
package main

// We `import` other packages to use their code. The `fmt` package (short for
// format) is part of Go's standard library and contains functions for
// formatted I/O (input/output), similar to Python's `print`.
import "fmt"

// The `main` function is the entry point of our program. When the program is
// run, the code inside this function is what executes first.
func main() {
	// `Println` is a FUNCTION from the `fmt` package that prints its
	// arguments to the console, followed by a new line.
	fmt.Println("Hello, World!")
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

HOW TO RUN THIS CODE:

1.  First, you need to install the Go toolchain from the official website.
2.  Open a terminal or command prompt.
3.  Navigate to the directory where you saved this file.
4.  Use the `go run` command to compile and execute the file:
    `go run 1_hello_world.go`

You should see "Hello, World!" printed to your console. Congratulations!
*/
