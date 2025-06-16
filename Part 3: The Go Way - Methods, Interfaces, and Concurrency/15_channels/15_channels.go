// Part 3, Lesson 15: Channels
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file serves as the lesson and demonstration for channels.
// It explains the core concepts through structured comments and
// provides a runnable example of their implementation.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

THE SOLUTION TO GOROUTINE SYNCHRONIZATION: CHANNELS

In the last lesson, we saw that the `main` goroutine exits without waiting for
other goroutines to finish. We used a "bad" `time.Sleep()` to work around this.
Now, we will learn the correct, idiomatic Go way to solve this problem: CHANNELS.

WHAT IS A CHANNEL?
A channel is a typed "pipe" or "conduit" that allows goroutines to communicate
with each other safely. [4] You can send values into a channel from one goroutine
and receive those values in another goroutine. [15] This mechanism is central to Go's
concurrency model and is expressed in the Go proverb:

"Do not communicate by sharing memory; instead, share memory by communicating."

Channels provide synchronization because send and receive operations BLOCK by
default until the other side is ready. This allows us to coordinate goroutines
without complex locks. [5, 7, 15]
*/

package main

import (
	"fmt"
	"time"
)

// --- Part 1: Creating, Sending, and Receiving ---

// A channel is created using the built-in `make()` function.
// Channels are strongly typed; you must specify the type of data
// that the channel is allowed to transport.
//
// SYNTAX:
// `myChannel := make(chan DataType)`
//
// The arrow `<-` is the channel operator.
//
// SEND a value into a channel:
// `myChannel <- "some value"`
//
// RECEIVE a value from a channel:
// `myValue := <-myChannel`

// --- Part 2: Using a Channel for Synchronization ---

// This function simulates a task that takes some time to complete.
// It accepts a channel as an argument, which it will use to signal
// when it is done.
func worker(done chan bool) {
	fmt.Println("Worker: Starting work...")
	time.Sleep(1 * time.Second) // Simulate a time-consuming task.
	fmt.Println("Worker: Work complete.")

	// Send a value (in this case, `true`) into the `done` channel.
	// This send operation will block until the main goroutine is ready
	// to receive it.
	done <- true
}

// --- Part 3: Using a Channel to Pass Data ---

// This function performs a "calculation" and sends the result
// back through a channel.
func calculator(resultChan chan string) {
	fmt.Println("Calculator: Performing complex calculation...")
	time.Sleep(2 * time.Second)
	result := "The answer is 42."

	// Send the final result back to the main goroutine.
	resultChan <- result
}

func main() {
	// --- Example 1: Basic Synchronization ---
	fmt.Println("--- Example 1: Synchronization ---")

	// Create a channel that will transport boolean values.
	done := make(chan bool)

	// Start the worker in a new goroutine, passing it the channel.
	go worker(done)

	// The `main` goroutine will now BLOCK on this line. It will wait
	// until a value is sent into the `done` channel. [11] This is how we
	// wait for the goroutine to finish its work without using `time.Sleep()`.
	<-done

	fmt.Println("Main: Received 'done' signal. Worker has finished.")
	fmt.Println()

	// --- Example 2: Passing Data Back ---
	fmt.Println("--- Example 2: Passing Data ---")

	// Create a channel that will transport a string.
	resultChan := make(chan string)

	go calculator(resultChan)

	// Block and wait for the result from the `calculator` goroutine.
	// The value sent into the channel is assigned to the `result` variable.
	result := <-resultChan
	fmt.Printf("Main: Received result from calculator: '%s'\n", result)
	fmt.Println()

	// --- Example 3: Ranging Over a Channel ---
	fmt.Println("--- Example 3: Ranging and Closing ---")
	tasks := make(chan string)

	go func() {
		// This goroutine sends multiple "tasks" to the channel.
		tasks <- "Task 1: Process data"
		tasks <- "Task 2: Send email"
		tasks <- "Task 3: Write to log"

		// It is important for the SENDER to CLOSE the channel to signal
		// that no more values will be sent. [3, 6, 23]
		close(tasks)
	}()

	// We can use a `for range` loop to receive values from a channel
	// until it is closed. [2, 8] The loop automatically breaks when the channel
	// is closed and all values have been received.
	fmt.Println("Main: Waiting to receive tasks...")
	for task := range tasks {
		fmt.Printf("Main: Received task - %s\n", task)
		time.Sleep(200 * time.Millisecond) // Simulate processing the task.
	}

	fmt.Println("Main: All tasks received. Channel was closed.")
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

HOW TO RUN THIS CODE:

1.  Open a terminal or command prompt.
2.  Navigate to the directory where you saved this file.
3.  Use the `go run` command to compile and execute the file:
    `go run 15_channels.go`

KEY TAKEAWAYS:
- Channels are the preferred way to communicate and synchronize between goroutines. [12, 15]
- Sending (`ch <- val`) and receiving (`val := <-ch`) are BLOCKING operations on unbuffered channels. [7]
- This blocking behavior is what allows for powerful, yet simple, synchronization. [5, 11]
- The sender should be the one to `close()` a channel to signal that no more data is coming.
- You can iterate over a channel with `for range` until it is closed. [2]
*/
