// Part 3, Lesson 16: Project: Concurrent Web Checker
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file is a project that demonstrates the power of goroutines and channels
// by building a tool that checks the status of a list of websites concurrently.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

PROJECT BRIEF: CONCURRENT WEBSITE STATUS CHECKER

In our last two lessons, we learned about GOROUTINES for running tasks at the
same time and CHANNELS for communicating safely between them. Now, we'll put
that knowledge to work in a practical, real-world project.

THE GOAL:
Build a command-line tool that takes a list of website URLs and checks the status
of each one as quickly as possible.

THE CHALLENGE:
If we check the websites one by one (sequentially), the total time taken will be
the sum of all the individual check times. Some websites might be slow to respond,
holding up the entire process.

THE SOLUTION:
We will use goroutines to check each website in its own lightweight thread of
execution. This means all the checks will happen concurrently. We'll use a
channel to collect the results from these goroutines as they finish.

This project will combine:
-   GOROUTINES: To run each check concurrently.
-   CHANNELS: To receive the status results back in the main function.
-   SLICES: To hold our list of websites.
-   PACKAGES: `fmt`, `net/http` (to make the web requests), and `time` (to time our program).
*/

package main

import (
	"fmt"
	"net/http"
)

// --- Part 1: The Worker Function ---

// This function will be run in its own goroutine for each website.
// It takes the URL to check and a channel to send the result back on.
func checkWebsite(url string, c chan string) {
	// The `http.Get` function makes an HTTP request to the given URL.
	// This is a "blocking" call; it will wait until the request is complete.
	resp, err := http.Get(url)

	// If there's an error (e.g., the domain doesn't exist or we can't connect),
	// the `err` variable will not be `nil`.
	if err != nil {
		// We send a string describing the error back through the channel.
		c <- fmt.Sprintf("%s might be down! Error: %s", url, err.Error())
		return
	}

	// It's very important to close the response body to free up the network
	// connection. `defer` ensures this happens at the end of the function.
	defer resp.Body.Close()

	// If we successfully get a response, we send a string with the site's
	// status code back through the channel.
	c <- fmt.Sprintf("%s is OK - Status: %s", url, resp.Status)
}

func main() {
	// --- Part 2: Setting Up the Work ---

	fmt.Println("Starting concurrent website check...")

	// Here is our list of websites to check.
	websites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"https://thiswebsitedoesnotexist.com", // This one should fail.
		"http://twitter.com",
	}

	// Create a channel to receive the string results from our goroutines.
	channel := make(chan string)

	// --- Part 3: Launching the Goroutines ---

	// We use a `for...range` loop to iterate over our slice of websites.
	for _, url := range websites {
		// For each URL, we launch a NEW GOROUTINE using the `go` keyword.
		// We pass the URL and the channel to our worker function.
		// The `main` function does NOT wait; it immediately starts the next one.
		go checkWebsite(url, channel)
	}

	// --- Part 4: Receiving the Results ---

	// Now we need to collect the results from the channel.
	// We know exactly how many results to expect: one for each website.
	// So we will loop that many times, receiving one result per iteration.
	for i := 0; i < len(websites); i++ {
		// The `<-channel` expression BLOCKS and waits until a value is sent
		// to the channel. As soon as a goroutine finishes its check and sends
		// a result, this line will receive it and print it.
		//
		// Notice that the results will print in the order they are completed,
		// NOT in the order we started them. This proves they are running concurrently!
		fmt.Println(<-channel)
	}

	fmt.Println("All websites have been checked.")
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

HOW TO RUN THIS CODE:

1.  Open a terminal or command prompt.
2.  Navigate to the directory where you saved this file.
3.  Use the `go run` command to compile and execute the file:
    `go run 16_project_concurrent_web_checker.go`

Congratulations! You've just built a high-performance, concurrent tool. If you
had checked these websites sequentially, the total time would be the sum of all
the request times. By doing it concurrently, the total time is roughly the time
it takes for the SLOWEST website to respond. This is the power of Go's concurrency
model in action.
*/
