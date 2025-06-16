// Part 3, Lesson 14: Goroutines
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file serves as the lesson and demonstration for goroutines.
// It explains the core concepts through structured comments and
// provides a runnable example of their implementation.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

Welcome to one of Go's most powerful and famous features: CONCURRENCY.

CONCURRENCY vs. PARALLELISM
Concurrency is the ability to handle multiple tasks at the same time. [12] It's about
structuring your program to deal with many things at once. Parallelism is doing
multiple tasks at the same time. Concurrency can exist on a single CPU core,
while parallelism requires multiple cores. Go makes concurrency simple.

WHAT IS A GOROUTINE?
A GOROUTINE is a lightweight thread of execution managed by the Go runtime. [4, 5] They are
incredibly cheap compared to traditional operating system threads, costing only
a few kilobytes of memory. [11, 21] This means you can easily have thousands, or even
millions, of goroutines running at once.

The `main()` function itself runs in a special, main goroutine. [9] Every other
goroutine is started from there.

How do you start one? With one simple keyword: `go`. [3, 7]
*/

package main

import (
	"fmt"
	"time"
)

// --- Part 1: A Simple Function to Run Concurrently ---

// We will use this simple function to demonstrate how goroutines work.
// It prints a message multiple times with a small delay.
func count(name string, limit int) {
	for i := 1; i <= limit; i++ {
		fmt.Printf("%s: %d\n", name, i)
		// We add a small sleep to simulate work and make the concurrent
		// nature easier to observe.
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Printf("--- '%s' goroutine finished ---\n", name)
}

func main() {
	// --- Part 2: Starting a Goroutine ---

	// First, let's call the function normally (synchronously).
	// The program will execute this function completely before moving on.
	fmt.Println("Starting synchronous call...")
	count("sync_call", 3)
	fmt.Println("Synchronous call finished.")
	fmt.Println("------------------------------------")

	// Now, let's start the same function as a GOROUTINE.
	// We just add the `go` keyword before the function call.
	fmt.Println("Starting goroutine...")
	go count("goroutine_1", 3)

	// WHAT HAPPENS NOW?
	// The `go` keyword starts the `count` function in a new goroutine. [1]
	// Crucially, the `main` function DOES NOT wait for it to finish. [2, 18] It
	// immediately continues to the next line of code.

	// --- Part 3: The Problem of the Exiting Main Function ---

	// If the `main` function finishes, the entire program exits, and any
	// running goroutines are terminated immediately. [17, 18]

	// If you run the code as-is (by commenting out the sleep below), you will
	// likely see "Main goroutine finished" and the program will exit before
	// "goroutine_1" gets a chance to print anything.

	// To see our goroutine in action, we need to make the main goroutine
	// wait. For this lesson, we will use a simple, but crude, method:
	// `time.Sleep()`. This is NOT a good or reliable way to coordinate
	// goroutines in a real application. It's just a tool to demonstrate
	// the concept for now.
	//
	// In the next lesson, we will learn about CHANNELS, which are the proper,
	// idiomatic Go way to communicate with and wait for goroutines.

	fmt.Println("Main goroutine is now waiting...")
	time.Sleep(300 * time.Millisecond)

	fmt.Println("Main goroutine finished. Program will now exit.")
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

HOW TO RUN THIS CODE:

1.  Open a terminal or command prompt.
2.  Navigate to the directory where you saved this file.
3.  Use the `go run` command to compile and execute the file:
    `go run 14_goroutines.go`

You will see the output of the synchronous call, then the output from the
goroutine interleaved with the messages from the main function. Try commenting
out the `time.Sleep` in `main` and run it again to see how the program exits
without waiting for the goroutine.

KEY TAKEAWAYS:
- Concurrency is about dealing with many things at once.
- A goroutine is a lightweight, cheap thread of execution. [3, 13]
- You start a goroutine with the `go` keyword.
- The `main` function (and the program) will NOT wait for goroutines to complete
  on its own. You must explicitly synchronize them.
*/
